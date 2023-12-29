package channelpool

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/utils"
)

// transmits messages received from other nodes in p2p to daemon
var IncomingMessagesP2P2_D_c = make(chan *entities.ClientMessage)

// transmits messages sent through rpc, or other channels to daemon
var SentMessagesRPC_D_c = make(chan *entities.ClientMessage)

// transmits validated messages from Daemon to P2P to be broadcasted
var OutgoingMessagesD_P2P_c = make(chan *entities.ClientMessage)

// transmits new subscriptions from RPC to Daemon for processing
var SubscribersRPC_D_c = make(chan *entities.SubscriptionEvent)

// transmits valid subscriptions from Daemon to P2P for broadcasting
var SubscriptionD_P2P_C = make(chan *entities.SubscriptionEvent)
var ClientHandshakeC = make(chan *utils.ClientHandshake)
var IncomingDeliveryProofsC = make(chan *entities.DeliveryProof)
var OutgoingDeliveryProof_BlockC = make(chan *entities.Block)
var OutgoingDeliveryProofC = make(chan *entities.DeliveryProof)
var PubSubInputBlockC = make(chan *entities.Block)
var PubSubInputProofC = make(chan *entities.DeliveryProof)
var PublishedSubC = make(chan *entities.SubscriptionEvent)
