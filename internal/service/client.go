package service

import (
	"context"
	"errors"

	"sync"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/channelpool"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/db"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
)

var logger = &log.Logger

func ConnectClient(message []byte, protocol constants.Protocol,  client interface{}) (*entities.ClientHandshake, error) {
	verifiedRequest, _ := entities.UnpackClientHandshake(message)
	verifiedRequest.ClientSocket = &client
	verifiedRequest.Protocol = protocol;
	logger.Debug("VerifiedRequest.Message: ", verifiedRequest.Message)
	vByte := []byte(verifiedRequest.Message)
	if crypto.VerifySignatureECC(verifiedRequest.Signer, &vByte, verifiedRequest.Signature) {
		// verifiedConn = append(verifiedConn, c)
		logger.Debug("Verification was successful: ", verifiedRequest)
		return &verifiedRequest, nil
	}
	return nil,  errors.New("Invaliad handshake")
	
}

func ValidateEvent(event interface{}) error {
	e := event.(entities.Event)
	b, err := e.EncodeBytes()
	if err != nil {
		logger.Errorf("Invalid Encoding %v", err)
		return err
	}
	// logger.Errorf("Payload Validator: %s; Event Signer: %s; PublicKey: %s", e.Payload.Validator, e.Validator, )
	if e.GetValidator() != e.Payload.Validator {
		return apperror.Forbidden("Payload validator does not match event validator")
	}
	logger.Infof("EVENT%s %s", string(e.GetValidator()), e.GetSignature())
	valid, err  := crypto.VerifySignatureEDD(string(e.GetValidator()), &b, e.GetSignature())
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("4003: Invalid node signature")
	}
	// TODO check to ensure that signer is an active validator, if not drop the event
	return nil
}



// func ValidateMessageClient(
// 	ctx context.Context,
// 	connectedSubscribers *map[string]map[string][]interface{},
// 	clientHandshake *entities.ClientHandshake,
// 	channelSubscriberStore *db.Datastore) {
// 	// VALIDATE AND DISTRIBUTE
// 	logger.Debugf("Signer:  %s\n", clientHandshake.Signer)
// 	results, err := channelSubscriberStore.Query(ctx, dsQuery.Query{
// 		Prefix: "/" + clientHandshake.Signer,
// 	})
// 	if err != nil {
// 		logger.Errorf("Channel Subscriber Store Query Error %o", err)
// 		return
// 	}
// 	entries, _err := results.Rest()
// 	for i := 0; i < len(entries); i++ {
// 		_sub, _ := entities.SubscriptionFromBytes(entries[i].Value)
// 		if (*connectedSubscribers)[_sub.TopicId] == nil {
// 			(*connectedSubscribers)[_sub.TopicId] = make(map[string][]interface{})
// 		}
// 		(*connectedSubscribers)[_sub.TopicId][_sub.Subscriber] = append((*connectedSubscribers)[_sub.TopicId][_sub.Subscriber], clientHandshake.ClientSocket)
// 	}
// 	logger.Infof("results:  %s  -  %o\n", entries[0].Value, _err)
// }

func ValidateAndAddToDeliveryProofToBlock(ctx context.Context,
	proof *entities.DeliveryProof,
	deliveryProofStore *db.Datastore,
	channelSubscriberStore *db.Datastore,
	stateStore *db.Datastore,
	localBlockStore *db.Datastore,
	MaxBlockSize int,
	mutex *sync.RWMutex,
) {
	 
	err := deliveryProofStore.Set(ctx, db.Key(proof.Key()), proof.MsgPack(), true)
	if err == nil {
		// msg, err := validMessagesStore.Get(ctx, db.Key(fmt.Sprintf("/%s/%s", proof.MessageSender, proof.MessageHash)))
		// if err != nil {
		// 	// invalid proof or proof has been tampered with
		// 	return
		// }
		// get signer of proof
		b, err := proof.EncodeBytes()
		if err != nil {
			return
		}
		susbscriber, err := crypto.GetSignerECC(&b, &proof.Signature)
		if (err != nil) {
			// invalid proof or proof has been tampered with
			return
		}
		// check if the signer of the proof is a member of the channel
		isSubscriber, err := channelSubscriberStore.Has(ctx, db.Key("/"+susbscriber+"/"+proof.MessageHash))
		if isSubscriber {
			// proof is valid, so we should add to a new or existing batch
			var block *entities.Block
			var err error
			txn, err := stateStore.NewTransaction(ctx, false)
			if err != nil {
				logger.Errorf("State query errror %o", err)
				// invalid proof or proof has been tampered with
				return
			}
			blockData, err := txn.Get(ctx, db.Key(constants.CurrentDeliveryProofBlockStateKey))
			if err != nil {
				logger.Errorf("State query errror %o", err)
				// invalid proof or proof has been tampered with
				txn.Discard(ctx)
				return
			}
			if len(blockData) > 0 && block.Size < MaxBlockSize {
				block, err = entities.UnpackBlock(blockData)
				if err != nil {
					logger.Errorf("Invalid batch %o", err)
					// invalid proof or proof has been tampered with
					txn.Discard(ctx)
					return
				}
			} else {
				// generate a new batch
				block = entities.NewBlock()

			}
			block.Size += 1
			if block.Size >= MaxBlockSize {
				block.Closed = true
				block.NodeHeight = chain.MLChainApi.GetCurrentBlockNumber()
			}
			// save the proof and the batch
			block.Hash = hexutil.Encode(crypto.Keccak256Hash([]byte(proof.Signature + block.Hash)))
			err = txn.Put(ctx, db.Key(constants.CurrentDeliveryProofBlockStateKey), block.MsgPack())
			if err != nil {
				logger.Errorf("Unable to update State store error %o", err)
				txn.Discard(ctx)
				return
			}
			proof.Block = block.BlockId
			proof.Index = block.Size
			err = deliveryProofStore.Put(ctx, db.Key(proof.Key()), proof.MsgPack())
			if err != nil {
				txn.Discard(ctx)
				logger.Errorf("Unable to save proof to store error %o", err)
				return
			}
			err = localBlockStore.Put(ctx, db.Key(constants.CurrentDeliveryProofBlockStateKey), block.MsgPack())
			if err != nil {
				logger.Errorf("Unable to save batch error %o", err)
				txn.Discard(ctx)
				return
			}
			err = txn.Commit(ctx)
			if err != nil {
				logger.Errorf("Unable to commit state update transaction errror %o", err)
				txn.Discard(ctx)
				return
			}
			// dispatch the proof and the batch
			if block.Closed {
				channelpool.OutgoingDeliveryProof_BlockC <- block
			}
			channelpool.OutgoingDeliveryProofC <- proof

		}

	}

}