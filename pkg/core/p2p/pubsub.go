package p2p

import (
	"context"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/mlayerprotocol/go-mlayer/entities"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type Channel struct {
	// Messages is a channel of messages received from other peers in the chat channel
	Messages chan entities.PubSubMessage

	Ctx   context.Context
	ps    *pubsub.PubSub
	Topic *pubsub.Topic
	sub   *pubsub.Subscription

	ChannelName string
	ID          peer.ID
	Wallet      string
}

func JoinChannel(ctx context.Context, ps *pubsub.PubSub, selfID peer.ID, walletAddress string, channelName string, channelBufferSize uint) (*Channel, error) {
	// join the pubsub topic
	topic, err := ps.Join(topicName(channelName))
	if err != nil {
		return nil, err
	}
	logger.Infof("Peer joined channel %s", channelName)

	
	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	cr := &Channel{
		Ctx:         ctx,
		ps:          ps,
		Topic:       topic,
		sub:         sub,
		ID:          selfID,
		Wallet:      walletAddress,
		ChannelName: channelName,
		Messages:    make(chan entities.PubSubMessage, channelBufferSize),
	}

	// start reading messages from the subscription in a loop
	go cr.readLoop()
	return cr, nil
}

// Publish sends a message to the pubsub topic.
func (cr *Channel) Publish(m entities.PubSubMessage) error {
	return cr.Topic.Publish(cr.Ctx, m.MsgPack())
}

func (cr *Channel) ListPeers() []peer.ID {
	return cr.ps.ListPeers(topicName(cr.ChannelName))
}

// readLoop pulls messages from the pubsub topic and pushes them onto the Messages channel.
func (cr *Channel) readLoop() {
	for {
		msg, err := cr.sub.Next(cr.Ctx)
		if err != nil {
			close(cr.Messages)
			panic(err)
		}
		// only forward messages delivered by others
		// if msg.ReceivedFrom == cr.ID {
		// 	continue
		// }
		pmsg, err := entities.UnpackPubSubMessage(msg.Data)
		
		if err != nil {
			logger.Errorf("Invalid pubsub message received: %v", err)
			continue
		}
		// logger.Infof("New Mesage %v \n\n %v", pmsg)
		// b, err := pmsg.EncodeBytes()
		// if(err != nil) {
		// 	logger.Error("Unable to encode msg %v", msg)
		// 	continue
		// }
		// signer, err := crypto.GetSignerECC(&b, &pmsg.Data.)
		// if err != nil {
		// 	logger.Error("Unable to get signer")
		// 	continue
		// }
		// logger.Infof("Pubsub message signer %s", signer)
		// TODO
		// get the stake contract for this signer and ensure they have enough Validator stake
		// if not, identify their IP and blacklist it. Ignore the message

		cr.Messages <- pmsg
	}
}

func topicName(channelName string) string {
	return "icm-channel:" + channelName
}
