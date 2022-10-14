package p2p

import (
	"context"

	// "github.com/ByteGum/go-icms/utils"
	"github.com/libp2p/go-libp2p-core/peer"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// Channel represents a subscription to a single PubSub topic. Messages
// can be published to the topic with Channel.Publish, and received
// messages are pushed to the Messages channel.
// type PubSubMessage struct {
// 	Data interface{}
// 	// ClientMessage    utils.ClientMessage
// 	// Subscription     utils.Subscription
// 	// SubscribersCount utils.SubscriberCount
// }

// func (msg PubSubMessage) ToJSON() []byte {
// 	m, _ := json.Marshal(msg)
// 	return m
// }

// func PubSubMessageFromBytes(b []byte) (, error) {
// 	var message PubSubMessage
// 	// if err := json.Unmarshal(b, &message); err != nil {
// 	// 	panic(err)
// 	// }
// 	err := json.Unmarshal(b, &message)
// 	return message, err
// }

type Channel struct {
	// Messages is a channel of messages received from other peers in the chat channel
	Messages chan []byte

	Ctx   context.Context
	ps    *pubsub.PubSub
	Topic *pubsub.Topic
	sub   *pubsub.Subscription

	ChannelName string
	ID          peer.ID
	Wallet      string
}

// ChannelMessage gets converted to/from JSON and sent in the body of pubsub messages.
// type ChannelMessage struct {
// 	Message utils.NodeMessage
// }

// JoinChannel tries to subscribe to the PubSub topic for the channel name, returning
// a Channel on success.
func JoinChannel(ctx context.Context, ps *pubsub.PubSub, selfID peer.ID, walletAddress string, channelName string, channelBufferSize uint) (*Channel, error) {
	// join the pubsub topic
	topic, err := ps.Join(topicName(channelName))
	if err != nil {
		return nil, err
	}

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
		Messages:    make(chan []byte, channelBufferSize),
	}

	// start reading messages from the subscription in a loop
	go cr.readLoop()
	return cr, nil
}

// Publish sends a message to the pubsub topic.
func (cr *Channel) Publish(m []byte) error {
	// if err != nil {
	// 	return err
	// }
	return cr.Topic.Publish(cr.Ctx, m)
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
		logger.Info("NEW MSG FROM CHANNEL")
		// only forward messages delivered by others
		if msg.ReceivedFrom == cr.ID {
			continue
		}
		// cm,  := msg.Data
		// if err != nil {
		// 	continue
		// }
		// send valid messages onto the Messages channel
		cr.Messages <- msg.Data
	}
}

func topicName(channelName string) string {
	return "icm-channel:" + channelName
}
