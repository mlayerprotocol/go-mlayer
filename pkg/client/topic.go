package client

import (
	// "errors"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/channelpool"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/internal/service"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
)

// type TopicService struct {
// 	Ctx context.Context
// 	Cfg configs.MainConfiguration
// }

// func NewTopicService(mainCtx *context.Context) *TopicService {
// 	ctx := *mainCtx
// 	cfg, _ := ctx.Value(constants.ConfigKey).(*configs.MainConfiguration)
// 	return &TopicService{
// 		Ctx: ctx,
// 		Cfg: *cfg,
// 	}
// }

// func (p *TopicService) NewTopicSubscription(sub *entities.Subscription) error {
// 	// subscribersc, ok := p.Ctx.Value(utils.SubscribeChId).(*chan *entities.Subscription)

// 	// validate before storing
// 	if entities.IsValidSubscription(*sub, true) {
// 		topicSubscriberStore, ok := p.Ctx.Value(constants.NewTopicSubscriptionStore).(*db.Datastore)
// 		if !ok {
// 			return errors.New("Could not connect to subscription datastore")
// 		}
// 		error := topicSubscriberStore.Set(p.Ctx, db.Key(sub.Key()), sub.MsgPack(), false)
// 		if error != nil {
// 			return error
// 		}
// 	}
// 	return nil
// }

/*
	Validate and Process the topic request
*/
func CreateTopic(
	payload entities.ClientPayload, ctx *context.Context,
	) (*models.TopicEvent, error) {
	
	cfg, _ :=(*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	// check if client payload is valid
	
	// if err := payload.Validate(entities.PublicKeyString(cfg.NetworkPublicKey)); err != nil {
	// 	return nil, err
	// }
	authState, err := ValidateClientPayload(&payload)
	if authState == nil {
		// agent not authorized
	}

	
	payloadData := entities.Topic{} 
	d, _ := json.Marshal(payload.Data)
	e := json.Unmarshal(d, &payloadData)
	if e!=nil {
		logger.Errorf("UnmarshalError %v", e)
	}
	payload.Data = payloadData
	// data := payload
	var assocPrevEvent string 
	var assocAuthEvent string
	if payload.EventType == uint16(constants.CreateTopicEvent) {
		// dont worry validating the AuthHash for Authorization requests
		if (uint64(payloadData.Timestamp) > uint64(time.Now().UnixMilli()) + 15000)  {
			return nil,  errors.New("Authorization timestamp exceeded")
		}
	
		currentState, grantorAuthState, err := service.ValidateTopicData(&payloadData)
		if err != nil {
			return nil, err
		}
		
		// generate associations
		if currentState != nil {
			assocPrevEvent= currentState.EventHash
		
		}
		if grantorAuthState != nil {
			assocAuthEvent = grantorAuthState.EventHash
			// assocAuthEvent =  entities.EventPath{
			// 	Relationship: entities.AuthorizationEventAssoc,
			// 	Hash: grantorAuthState.EventHash,
			// 	Model: entities.AuthorizationEventModel,
			// }
		}

	}
	payloadHash, _ := payload.GetHash()
	
	// create event struct
	event :=  entities.Event{
			Payload: payload,
			Timestamp: uint64(time.Now().UnixMilli()),
			EventType: uint16(payload.EventType),
			Associations : []string{},
			PreviousEventHash: assocPrevEvent,
			AuthEventHash: assocAuthEvent,
			Synced: false,
			PayloadHash: hex.EncodeToString(payloadHash),
			Broadcasted : false,
			BlockNumber: chain.MLChainApi.GetCurrentBlockNumber(),
			Validator: entities.PublicKeyString(cfg.NetworkPublicKey),
		}
	
	logger.Infof("Validator: %s", event.Validator)
	b, err := event.EncodeBytes()
	
	event.Hash =  hex.EncodeToString(crypto.Sha256(b))
	_, event.Signature  = crypto.SignEDD(b, cfg.NetworkPrivateKey)
	
	
	eModel, created, err := query.SaveRecord(
		models.TopicEvent{
			Event: entities.Event{Hash: event.Hash},
		},
		models.TopicEvent{
			Event: event,
		}, false, nil)
	if err != nil {
		return nil, err
	}
	
	// channelpool.TopicEventPublishC <- &(eModel.Event)
	
	if created {
		channelpool.TopicEventPublishC <- &(eModel.Event)
	}
	return eModel, nil
}

// func ListenForNewTopicEventFromPubSub (mainCtx *context.Context) {
// 	ctx, cancel := context.WithCancel(*mainCtx)
// 	defer cancel()
	
// 	incomingTopicC, ok := (*mainCtx).Value(constants.IncomingTopicEventChId).(*chan *entities.Event)
// 	if !ok {
// 		logger.Errorf("incomingTopicC closed")
// 		return
// 	}
// 	for {
// 		event, ok :=  <-*incomingTopicC
// 		if !ok {
// 			logger.Fatal("incomingTopicC closed for read")
// 			return
// 		}
// 		go service.HandleNewPubSubTopicEvent(event, ctx)
// 	}
// }
