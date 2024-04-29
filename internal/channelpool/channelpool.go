package channelpool

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
)

// channels for moving events received from other nodes through the pubsub channels
var AuthorizationEvent_SubscriptionC = make(chan *entities.Event)
var IncomingTopicEventSubscriptionC = make(chan *entities.Event)

// channels for broadcasting new events to other nodes
var AuthorizationEventPublishC = make(chan *entities.Event)
var TopicEventPublishC = make(chan *entities.Event)
var SubNetworkEventPublishC = make(chan *entities.Event)
var SubscriptionEventPublishC = make(chan *entities.Event)
var MessageEventPublishC = make(chan *entities.Event)
var UnSubscribeEventPublishC = make(chan *entities.Event)
var ApproveSubscribeEventPublishC = make(chan *entities.Event)

// CLEANUP
// most of these will be deleted
// transmits events received from other nodes in p2p to daemon
var IncomingMessageEvent_P2P_D_C = make(chan *entities.Event)

// transmits validated events from Daemon to P2P to be broadcasted
var OutgoingMessageEvents_D_P2P_C = make(chan *entities.Event)

// transmits messages sent through rpc, or other channels to daemon
var NewPayload_Cli_D_C = make(chan *entities.ClientPayload)

// transmits new subscriptions from RPC to Daemon for processing
var Subscribers_RPC_D_C = make(chan *entities.Subscription)

// transmits valid subscriptions from Daemon to P2P for broadcasting
var Subscription_D_P2P_C = make(chan *entities.Subscription)

var ClientHandshakeC = make(chan *entities.ClientHandshake)
var IncomingDeliveryProofsC = make(chan *entities.DeliveryProof)
var OutgoingDeliveryProof_BlockC = make(chan *entities.Block)
var OutgoingDeliveryProofC = make(chan *entities.DeliveryProof)
var PubSubInputBlockC = make(chan *entities.Block)
var PubSubInputProofC = make(chan *entities.DeliveryProof)
var PublishedSubC = make(chan *entities.Subscription)
