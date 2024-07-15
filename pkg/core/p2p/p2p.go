package p2p

import (
	"bufio"
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	// "github.com/gin-gonic/gin"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/channelpool"
	"github.com/mlayerprotocol/go-mlayer/internal/service"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/db"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/p2p/notifee"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"

	record "github.com/libp2p/go-libp2p-record"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
	"github.com/multiformats/go-multiaddr"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/routing"
	noise "github.com/libp2p/go-libp2p/p2p/security/noise"
	libp2ptls "github.com/libp2p/go-libp2p/p2p/security/tls"
	"github.com/sirupsen/logrus"
	// rest "messagingprotocol/pkg/core/rest"
	// dhtConfig "github.com/libp2p/go-libp2p-kad-dht/internal/config"
)

var logger = &log.Logger

// var config configs.MainConfiguration
type P2pChannelFlow int8
const (
	P2pChannelOut P2pChannelFlow = 1
	P2pChannelIn P2pChannelFlow = 2
)

var protocolId string
var privKey crypto.PrivKey
var config *configs.MainConfiguration
var handShakeProtocolId = "mlayer/handshake/1.0.0"
var P2pComChannels map[string]map[P2pChannelFlow]chan P2pPayload


func init() {
	P2pComChannels = make(map[string]map[P2pChannelFlow]chan P2pPayload)
} 


const (
	AuthorizationChannel string = "ml-authorization-channel"
	TopicChannel         string = "ml-topic-channel"
	SubnetChannel        string = "ml-sub-network-channel"
	WalletChannel        string = "ml-wallet-channel"
	MessageChannel       string = "ml-message-channel"
	SubscriptionChannel         = "ml-subscription-channel"
	// UnSubscribeChannel                = "ml-unsubscribe-channel"
	// ApproveSubscriptionChannel        = "ml-approve-subscription-channel"
	BatchChannel         = "ml-batch-channel"
	DeliveryProofChannel = "ml-delivery-proof"
)

// var PeerStreams = make(map[string]peer.ID)
var PeerPubKeys = make(map[peer.ID][]byte)
var DisconnectFromPeer = make(map[peer.ID]bool)
// var node *host.Host
var idht *dht.IpfsDHT




// defaultNick generates a nickname based on the $USER environment variable and
// the last 8 chars of a peer ID.
func defaultNick(p peer.ID) string {
	// TODO load name from flag/config
	return fmt.Sprintf("%s-%s", os.Getenv("USER"), shortID(p))
}

// shortID returns the last 8 chars of a base58-encoded peer id.
func shortID(p peer.ID) string {
	pretty := p.String()
	return pretty[len(pretty)-12:]
}

func discover(ctx context.Context, h host.Host, kdht *dht.IpfsDHT, rendezvous string, config *configs.MainConfiguration) {
	// kdht.PutValue(ctx, "user/name", []byte("femi"))
	// v, err := kdht.GetValue(ctx, "user/name")
	// if err != nil {
	// 	logger.Error("KDHTERROR", err)
	// }
	// logger.Infof("VALUEEEEE %s", string(v))
	routingDiscovery := drouting.NewRoutingDiscovery(kdht)
	dutil.Advertise(ctx, routingDiscovery, rendezvous)

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:

			peers, err := routingDiscovery.FindPeers(ctx, rendezvous)
			if err != nil {
				logger.Error(err)
				continue
			}
		 	
			for p := range peers {

				if p.ID == h.ID() {
					continue
				}

				if h.Network().Connectedness(p.ID) != network.Connected {
					_, err = h.Network().DialPeer(ctx, p.ID)
					if err != nil {
						logger.Debugf("Failed to connect to peer: %s \n%s", p.ID.String(), err.Error())
						h.Peerstore().RemovePeer(p.ID)
						kdht.ForceRefresh()
						continue
					}	
					if len(p.ID) == 0 {
						continue
					}				
					logger.Infof("Connected to discovered peer: %s at %s \n", p.ID.String(), p.Addrs)
					handleConnect(&h, &p)
				}
			}
		}
	}
}


func Run(mainCtx *context.Context) {
	// fmt.Printf("publicKey %s", privateKey)
	// The context governs the lifetime of the libp2p node.
	// Cancelling it will stop the the host.

	ctx, cancel := context.WithCancel(*mainCtx)
	defer cancel()

	cfg, ok := ctx.Value(constants.ConfigKey).(*configs.MainConfiguration)
	config = cfg
	if !ok {
		panic("Unable to load config from context")
	}

	p2pDataStore := db.New(&ctx,   string(constants.P2PDataStore))
	defer p2pDataStore.Close()

	if !ok {
		panic("Unable to load data store from context")
	}

	protocolId = config.ProtocolVersion

	

	// incomingAuthorizationC, ok := ctx.Value(constants.IncomingAuthorizationEventChId).(*chan *entities.Event)
	// if !ok {
	// 	panic(apperror.Internal("incomingAuthorizationC channel closed"))
	// }

	// incomingTopicEventC, ok := ctx.Value(constants.IncomingTopicEventChId).(*chan *entities.Event)
	// if !ok {
	// 	panic(apperror.Internal("incomingTopicEventC channel closed"))
	// }

	// incomingMessagesC, ok := ctx.Value(constants.IncomingMessageChId).(*chan *entities.ClientPayload)
	// if !ok {

	// }
	// outgoinMessageC, ok := ctx.Value(utils.OutgoingMessageDP2PChId).(*chan *entities.ClientPayload)
	// if !ok {

	// }

	// subscriptionC, ok := ctx.Value(constants.SubscriptionDP2PChId).(*chan *entities.Subscription)
	// if !ok {

	// }

	// outgoingDPBlockCh, ok := ctx.Value(constants.OutgoingDeliveryProof_BlockChId).(*chan *entities.Block)
	// outgoingProofCh, ok := ctx.Value(utils.OutgoingDeliveryProofCh).(*chan *utils.DeliveryProof)
	// publishedSubscriptionC, ok := ctx.Value(constants.SubscribeChId).(*chan *entities.Subscription)
	// if !ok {

	// }

	privKey, _, err := crypto.GenerateKeyPair(
		crypto.Ed25519, // Select your key type. Ed25519 are nice short
		2048,             // Select key length when possible (i.e. RSA).
	)
	if err != nil {
		panic(err)
	}
	

	// if len(config.NodePrivateKey) == 0 {
	// 	priv, _, err := crypto.GenerateKeyPair(
	// 		crypto.Ed25519, // Select your key type. Ed25519 are nice short
	// 		-1,             // Select key length when possible (i.e. RSA).
	// 	)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	privKey = priv
	// } else {
	// 	priv, err := crypto.UnmarshalECDSAPrivateKey(hexutil.MustDecode(config.NodePrivateKey))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	privKey = priv
	// }

	// conMgr := connmgr.NewConnManager(
	// 	100,         // Lowwater
	// 	400,         // HighWater,
	// 	time.Minute, // GracePeriod
	// )
	

	h, err := libp2p.New(
		// Use the keypair we generated
		libp2p.Identity(privKey),
		// Multiple listen addresses
		libp2p.ListenAddrStrings(config.ListenerAdresses...),
		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),
		// support noise connections
		libp2p.Security(noise.ID, noise.New),
		// support any other default transports (TCP)
		libp2p.DefaultTransports,
		// libp2p.Transport(ws.New),
		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.

		// libp2p.ConnectionManager(connmgr.NewConnManager(
		// 	100,         // Lowwater
		// 	400,         // HighWater,
		// 	time.Minute, // GracePeriod
		// )),

		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),
		// Let this host use the DHT to find other hosts
		
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {

			var bootstrapPeers []peer.AddrInfo
			
			for _, addr := range config.BootstrapPeers {
				addr, _ := multiaddr.NewMultiaddr(addr)
				pi, _ := peer.AddrInfoFromP2pAddr(addr)
				bootstrapPeers = append(bootstrapPeers, *pi)
			}
			var dhtOptions []dht.Option
			dhtOptions = append(dhtOptions,
				dht.BootstrapPeers(bootstrapPeers...),
			dht.ProtocolPrefix(protocol.ID(protocolId)),
			dht.ProtocolPrefix(protocol.ID(handShakeProtocolId)),
			dht.Datastore(p2pDataStore),
			dht.NamespacedValidator("pk", record.PublicKeyValidator{}),
				dht.NamespacedValidator("ipns", record.PublicKeyValidator{}),
				dht.NamespacedValidator("ml", &DhtValidator{}),
			)
			if !config.BootstrapNode {
				dhtOptions = append(dhtOptions, dht.Mode(dht.ModeServer))
			}
			// dhtOptions = append(dhtOptions,  dht.Datastore(syncDatastore))

			
			kdht, err := dht.New(ctx, h,  
				dhtOptions...,  )
			if err != nil {
				panic(err)
			}

			// validator = {
			// 	// Validate validates the given record, returning an error if it's
			// 	// invalid (e.g., expired, signed by the wrong key, etc.).
			// 	Validate(key string, value []byte) error

			// 	// Select selects the best record from the set of records (e.g., the
			// 	// newest).
			// 	//
			// 	// Decisions made by select should be stable.
			// 	Select(key string, values [][]byte) (int, error)
			// }
			// dhtOptions = append(dhtOptions, dht.NamespacedValidator("subsc", customValidator))

		
			

			//if cfg.BootstrapNode {
				if err = kdht.Bootstrap(ctx); err != nil {
					logger.Fatalf("Error starting bootstrap node %o", err)
					return nil, err
				}
			// }

			idht = kdht

			// for _, addr := range config.BootstrapPeers {
			// 	addr, _ := multiaddr.NewMultiaddr(addr)
			// 	pi, err := peer.AddrInfoFromP2pAddr(addr)
			// 	if err != nil {
			// 		logger.Warnf("Invalid boostrap peer address (%s): %s \n", addr, err)
			// 	} else {
			// 		error := h.Connect(ctx, *pi)
			// 		if error != nil {
			// 			logger.Debugf("Unable connect to boostrap peer (%s): %s \n", addr, err)
			// 			continue
			// 		}
			// 		logger.Debugf("Connected to boostrap peer (%s)", addr)
			// 		h.Peerstore().AddAddrs(pi.ID, pi.Addrs, peerstore.PermanentAddrTTL)
			// 		handleConnect(&h, pi)
			// 	}
			// }
			

			// routingOptions := routing.Options{
			// 	Expired: true,
			// 	Offline: true,
			// }
			// var	routingOptionsSlice []routing.Option;
			// routingOptionsSlice = append(routingOptionsSlice, routingOptions.ToOption())
			// key := "/$name/$first"
			// putErr := kdht.PutValue(ctx, key, []byte("femi"), routingOptions.ToOption())

			// if putErr != nil {
			// 	logger.Infof("Put the error %o", putErr)
			// }
			return idht, err
		}),
		// libp2p.Relay(options...),
		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		// libp2p.DefaultEnableRelay(),
		//libp2p.EnableAutoRelay(),
		// If you want to help other peers to figure out if they are behind
		// NATs, you can launch the server-side of AutoNAT too (AutoRelay
		// already runs the client)
		//
		// This service is highly rate-limited and should not cause any
		// performance issues.
		libp2p.EnableNATService(),
	)

	// gater := NetworkGater{host: h, config: config, blockPeers: make(map[peer.ID]struct{})}

    
	go discover(ctx, h, idht,  fmt.Sprintf("%s-%s", constants.NETWORK_NAME, config.AddressPrefix), config)
	if err != nil {
		panic(err)
	}
	h.Network().Notify(&notifee.ConnectionNotifee{Dht: idht})
	
	h.SetStreamHandler(protocol.ID(handShakeProtocolId), handleHandshake)
	h.SetStreamHandler(protocol.ID(protocolId), handlePayload)
	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}
	// setup local mDNS discovery
	err = setupDiscovery(h, fmt.Sprintf("%s-%s", constants.NETWORK_NAME, config.AddressPrefix))
	if err != nil {
		panic(err)
	}

	// node = &h

	// The last step to get fully up and running would be to connect to
	// bootstrap peers (or any other peers). We leave this commented as
	// this is an example and the peer will die as soon as it finishes, so
	// it is unnecessary to put strain on the network.

	logger.Infof("Host started with ID %s", h.ID())
	logger.Infof("Host Network: %s", protocolId)
	logger.Infof("Host Listening on: %s", h.Addrs())

	// Subscrbers
	authorizationPubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), AuthorizationChannel, config.ChannelMessageBufferSize)
	if err != nil {
		panic(err)
	}

	topicPubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), TopicChannel, config.ChannelMessageBufferSize)
	if err != nil {
		panic(err)
	}

	SubnetPubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), SubnetChannel, config.ChannelMessageBufferSize)
	if err != nil {
		panic(err)
	}

	WalletPubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), WalletChannel, config.ChannelMessageBufferSize)
	if err != nil {
		panic(err)
	}

	subscriptionPubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), SubscriptionChannel, config.ChannelMessageBufferSize)
	if err != nil {
		panic(err)
	}

	messagePubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), MessageChannel, config.ChannelMessageBufferSize)
	if err != nil {
		panic(err)
	}



	

	// unsubscribePubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), UnSubscribeChannel, config.ChannelMessageBufferSize)
	// if err != nil {
	// 	panic(err)
	// }

	// approveSubscriptionPubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), ApproveSubscriptionChannel, config.ChannelMessageBufferSize)
	// if err != nil {
	// 	panic(err)
	// }


	
	// Publishers
	go PublishChannelEventToNetwork(channelpool.AuthorizationEventPublishC, authorizationPubSub, mainCtx)
	go PublishChannelEventToNetwork(channelpool.TopicEventPublishC, topicPubSub, mainCtx)
	go PublishChannelEventToNetwork(channelpool.SubnetEventPublishC, SubnetPubSub, mainCtx)
	go PublishChannelEventToNetwork(channelpool.WalletEventPublishC, WalletPubSub, mainCtx)
	go PublishChannelEventToNetwork(channelpool.SubscriptionEventPublishC, subscriptionPubSub, mainCtx)
	go PublishChannelEventToNetwork(channelpool.MessageEventPublishC, messagePubSub, mainCtx)
	// go PublishChannelEventToNetwork(channelpool.UnSubscribeEventPublishC, unsubscribePubSub, mainCtx)
	// go PublishChannelEventToNetwork(channelpool.ApproveSubscribeEventPublishC, approveSubscriptionPubSub, mainCtx)

	// Subscribers
	

	go ProcessEventsReceivedFromOtherNodes(&entities.Authorization{}, authorizationPubSub, mainCtx, service.HandleNewPubSubAuthEvent)
	go ProcessEventsReceivedFromOtherNodes(&entities.Topic{}, topicPubSub, mainCtx, service.HandleNewPubSubTopicEvent)
	go ProcessEventsReceivedFromOtherNodes(&entities.Subnet{}, SubnetPubSub, mainCtx, service.HandleNewPubSubSubnetEvent)
	go ProcessEventsReceivedFromOtherNodes(&entities.Wallet{}, WalletPubSub, mainCtx, service.HandleNewPubSubWalletEvent)
	go ProcessEventsReceivedFromOtherNodes(&entities.Subscription{}, subscriptionPubSub, mainCtx, service.HandleNewPubSubSubscriptionEvent)
	go ProcessEventsReceivedFromOtherNodes(&entities.Message{}, messagePubSub, mainCtx, service.HandleNewPubSubMessageEvent)
	// go ProcessEventsReceivedFromOtherNodes(&entities.Subscription{}, unsubscribePubSub, mainCtx, service.HandleNewPubSubUnSubscribeEvent)
	// go ProcessEventsReceivedFromOtherNodes(&entities.Subscription{}, approveSubscriptionPubSub, mainCtx, service.HandleNewPubSubApproveSubscriptionEvent)

	// messagePubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), MessageChannel, config.ChannelMessageBufferSize)
	// if err != nil {
	// 	panic(err)
	// }

	// batchPubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), BatchChannel, config.ChannelMessageBufferSize)
	// if err != nil {
	// 	panic(err)
	//}
	// delieveryProofPubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), DeliveryProofChannel, config.ChannelMessageBufferSize)
	// if err != nil {
	// 	panic(err)
	// }

	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	for {
	// 		select {

	// 		case authEvent, ok := <-authorizationPubSub.Messages:
	// 			if !ok {
	// 				cancel()
	// 				logger.Fatalf("Primary Message channel closed. Please restart server to try or adjust buffer size in config")
	// 				return
	// 			}
	// 			// !validating message
	// 			// !if not a valid message continue
	// 			// _, err := inMessage.MsgPack()
	// 			// if err != nil {
	// 			// 	continue
	// 			// }
	// 			//TODO:
	// 			// if not a valid message, continue

	// 			logger.Infof("Received new message %s\n", authEvent.ToString())
	// 			cm := models.AuthorizationEvent{}
	// 			err = encoder.MsgPackUnpackStruct(authEvent.Data, cm)
	// 			if err != nil {

	// 			}
	// 			*incomingAuthorizationC <- &cm
	// 		case inMessage, ok := <-batchPubSub.Messages:
	// 			if !ok {
	// 				cancel()
	// 				logger.Fatalf("Primary Message channel closed. Please restart server to try or adjust buffer size in config")
	// 				return
	// 			}
	// 			// !validating message
	// 			// !if not a valid message continue
	// 			// _, err := inMessage.MsgPack()
	// 			// if err != nil {
	// 			// 	continue
	// 			// }
	// 			//TODO:
	// 			// if not a valid message, continue

	// 			logger.Infof("Received new message %s\n", inMessage.ToString())
	// 			cm, err := entities.MsgUnpackClientPayload(inMessage.Data)
	// 			if err != nil {

	// 			}
	// 			*incomingMessagesC <- &cm
	// 		case sub, ok := <-subscriptionPubSub.Messages:
	// 			if !ok {
	// 				cancel()
	// 				logger.Fatalf("Primary Message channel closed. Please restart server to try or adjust buffer size in config")
	// 				return
	// 			}
	// 			// logger.Info("Received new message %s\n", inMessage.Message.Body.Message)
	// 			cm, err := entities.UnpackSubscription(sub.Data)
	// 			if err != nil {

	// 			}
	// 			logger.Info("New subscription updates:::", string(cm.ToJSON()))
	// 			// *incomingMessagesC <- &cm
	// 			cm.Broadcasted = false
	// 			*publishedSubscriptionC <- &cm
	// 		}
	// 	}
	// }()
	if config.Validator {
		storeAddress(&ctx, &h)
	}
	defer forever()

}

func forever() {
	for {
		time.Sleep(time.Hour)
	}
}

// func handleHandshake(stream network.Stream) {
	
// 	// // if len(PeerStreams[stream.ID()]) == 0 {
// 	// // 	logger.Infof("No peer for stream %s Peer %s", stream.ID(), PeerStreams[stream.ID()])
// 	// // 	return
// 	// // }
// 	// logger.Infof("Got a new stream 2! %s Peer %s", stream.ID(), PeerStreams[stream.ID()])
// 	// // stream.SetReadDeadline()
// 	// // Create a buffer stream for non blocking read and write.
// 	// peer := stPeerStreams[stream.ID()]
// 	// rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
// 	// logger.Infof("Got a new stream 3! %s Peer %s", stream.ID(), PeerStreams[stream.ID()])
// 	host := idht.Host()
// 	verifyHandshake(&host, &stream)
// 	// go sendData(rw)

// }

func handleHandshake(stream network.Stream) {
	 // ctx, _ := context.WithCancel(context.Background())
	// config, _ := ctx.Value(constants.ConfigKey).(*configs.MainConfiguration)
	// defer delete(DisconnectFromPeer, p )
	peerId := (stream).Conn().RemotePeer()

	
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	for {
		hsData, err := rw.ReadBytes('\n')
		if err != nil   {
			logger.Errorf("Error reading from buffer %o", err)
			return
		}
		if hsData == nil {
			//break
			return
		}

 		handshake, err := UnpackNodeHandshake(hsData)
		
		if err != nil {
			logger.WithFields(logrus.Fields{ "data": handshake}).Warnf("Failed to parse handshake: %o", err)
			return
			// break
		}
		validHandshake := handshake.IsValid(config.AddressPrefix)
		
		logger.Infof("Validating peer %s", (stream).Conn().RemotePeer())
		if !validHandshake {
			disconnect((stream).Conn().RemotePeer())
			return
		}
		if handshake.NodeType == constants.ValidatorNodeType {
			// Validate stake as well

			// validStake := isValidStake(handshake, p, config)
			// if !validStake {
			// 	// disconnect(*node, p)
			// 	logger.WithFields(logrus.Fields{"address": handshake.Signer, "data": hsData}).Infof("Disconnecting from peer (%s) with inadequate stake in network", p)
			// 	return
			// }
		}
	

		// b, _ := hexutil.Decode(handshake.Signer)
		// PeerPubKeys[p] = b
		// break
		logger.WithFields(logrus.Fields{ "peer": peerId, "pubKey": handshake.Signer }).Info("Successfully connected peer with valid handshake")
		delete(DisconnectFromPeer, peerId)
	 }
}

func sendHandshake( stream  network.Stream, data []byte) {
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	_, err := rw.WriteString(fmt.Sprintf("%s\n", string(data)))
	if err != nil {
		logger.Warn("Error writing to to stream")
		return
	}
	
	err = rw.Flush()
	logger.Infof("Flushed data to stream %s", stream.ID())
	if err != nil {
		// fmt.Println("Error flushing buffer")
		// panic(err)
		logger.Error("Error flushing to to stream")
		return
	}
}

func storeAddress(ctx *context.Context, h *host.Host) {
	for {
		logger.Infof("Peers with address %d", (*h).Peerstore().PeersWithAddrs().Len())
		if (*h).Peerstore().PeersWithAddrs().Len() == 1 {
			time.Sleep(60 * time.Second)
			continue
		}
		time.Sleep(4 * time.Hour)		 
		 mad, err := NewNodeMultiAddressData(config, config.PrivateKey, getMultiAddresses(*h))
		if err != nil {
			logger.Error(err)
		}
		key := "/ml/val/"+config.OperatorPublicKey
		// v, err := idht.GetValue(*ctx, key)
		// 	if err != nil {
		// 		logger.Error("KDHT_GET_ERROR: ", err)
		// 	} else {
		// 		logger.Infof("VALURRRR %s", string(v))
		// 	}
		err = idht.PutValue(*ctx, key,  mad.MsgPack())
		if err != nil {
			logger.Error("KDHT_PUT_ERROR", err)
		}  else {
			logger.Infof("Successfully put value")
		}
		// else {
		// 	time.Sleep(2 * time.Second) 
		// 	v, err := idht.GetValue(ctx, key)
		// 	if err != nil {
		// 		logger.Error("KDHT_GET_ERROR", err)
		// 	} else {
		// 		logger.Infof("VALURRRR %s", string(v))
		// 	}
		// }
	}
}
// called when a peer connects
func handleConnect(h *host.Host, pairAddr *peer.AddrInfo) {
	// pi := *pa
	logger.Infof("My multiaddress: %s", getMultiAddresses(*h))
	if pairAddr == nil {
		return
	}
	DisconnectFromPeer[pairAddr.ID] = true
	
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	handshakeStream, err := (*h).NewStream(ctx, pairAddr.ID, protocol.ID(handShakeProtocolId))
		
	if err != nil {
		logger.Warningf("Unable to establish stream with peer: %s %o", pairAddr.ID, err)
	} else {
		nodeType := constants.RelayNodeType
		if config.Validator {
			nodeType = constants.ValidatorNodeType
		}
		hs, _ := NewNodeHandshake(config, handShakeProtocolId, config.PrivateKey, nodeType)
		// b, _ := hs.EncodeBytes()
		logger.Infof("Created handshake with salt %s", hs.Salt)
		logger.Infof("Created new stream %s with peer %s", handshakeStream.ID(), pairAddr.ID)
		defer func (pairID peer.ID) {
			time.Sleep(3 * time.Second)
			if DisconnectFromPeer[pairAddr.ID] {
				handshakeStream.Close()
				disconnect(pairAddr.ID)
			}
		}(pairAddr.ID)
		go sendHandshake(handshakeStream, (*hs).MsgPack())
		go handleHandshake(handshakeStream)

		host := idht.Host()
		networkStream, err := host.NewStream(idht.Context(), pairAddr.ID, protocol.ID(protocolId))
		if err != nil {
			(networkStream).Reset()
			return
		}
		go handlePayload(networkStream)
		
		

		// _, pub, _ := crypto.GenerateKeyPair(crypto.RSA, 2048)
		//time.Sleep(5 * time.Second) 
		// peerID, _ := peer.IDFromPublicKey(pub)
		logger.Infof("Waiting to send data to peer")
		time.Sleep(5 * time.Second)
		logger.Infof("Starting send process...")
		eventBytes := (&entities.EventPath{}).MsgPack()
		payload, err := NewP2pPayload(config, P2pActionGetEvent, eventBytes, config.PrivateKey)
		if err != nil {
			logger.Errorf("ERrror: %s", err)
			return
		}
		err = payload.Sign(config.PrivateKey)
		if err != nil {
			logger.Infof("Error SIgning: %v", err)
			return
		}
		logger.Infof("Payload data signed: %s", payload.Id)
		resp, err := payload.SendRequest(&ctx, config, multiaddr.StringCast("/ip4/127.0.0.1/tcp/6000/ws/p2p/12D3KooWSrBKX6uWXsxFdi1KFHEimiEYBDVMADtYXZXNsXChkccF"))
		if err != nil {
			logger.Errorf("Error :%v", err)
		}
		logger.Infof("Resopnse :%s", string(resp))
		
	}
}

func disconnect( id peer.ID) {
	idht.Host().Network().ClosePeer(id)
}

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func setupDiscovery(h host.Host, serviceName string) error {
	logger.Infof("Setting up Discovery on %s ....", serviceName)
	n := notifee.DiscoveryNotifee{Host: h, HandleConnect: handleConnect, Dht: idht}
	
	disc := mdns.NewMdnsService(h, serviceName, &n)
	return disc.Start()
}

func connectToNode(targetAddr multiaddr.Multiaddr, connectionId string, ctx context.Context) (*peer.AddrInfo, error) {
	targetInfo, err := peer.AddrInfoFromP2pAddr(targetAddr)
	if err != nil {
		logger.Errorf("Failed to get peer info: %v", err)
		return targetInfo, err
	}
	
	if P2pComChannels[connectionId] == nil {
		P2pComChannels[connectionId] = make(map[P2pChannelFlow]chan P2pPayload)
		P2pComChannels[targetInfo.ID.String()] = make(map[P2pChannelFlow]chan P2pPayload)
	  P2pComChannels[connectionId][P2pChannelIn] = make(chan P2pPayload)
	  P2pComChannels[targetInfo.ID.String()][P2pChannelOut] = make(chan P2pPayload)
	}
	h := idht.Host()
	// Add the target peer to the host's peerstore
	h.Peerstore().AddAddrs(targetInfo.ID, targetInfo.Addrs, peerstore.PermanentAddrTTL)
	err = h.Connect(ctx, *targetInfo)
	if err != nil {
		h.Peerstore().RemovePeer(targetInfo.ID)
		delete(P2pComChannels, connectionId)
		return nil, err
	}

	// Connect to the target node
	return targetInfo, h.Connect(ctx, *targetInfo);
}

func getMultiAddresses(h host.Host) []string {
	m := []string{}
		addrs := h.Addrs()
		
		for _, addr := range addrs {
			m = append(m, fmt.Sprintf("%s/p2p/%s\n", addr, h.ID().String()))
		}
		logger.Infof("MULTI %v", m)
	return m
}

func handlePayload(stream network.Stream) {
	go writePayload(stream)
	go readPayload(stream)
}

func writePayload(stream network.Stream) {   
	
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	peerId := stream.Conn().RemotePeer()
	delimeter :=  []byte{'\n'}
	logger.Infof("Remote Peer Id: %s", peerId.String())
	// go func () {
	// 	time.Sleep(7 * time.Second)
	// 	P2pComChannels[peerId.String()]["o"] <- []byte("Femi")
	// }()
	for {
		if P2pComChannels[peerId.String()] == nil {
			time.Sleep(1 * time.Second)
			continue
		}
		
		msg := <-P2pComChannels[peerId.String()][P2pChannelOut]
		
		logger.Infof("Received Data: %T", msg.IsValid(config.AddressPrefix))
	
		_, err := rw.Write(append(msg.MsgPack(), delimeter...))
			if err != nil {
				logger.Warn("Error writing to to stream")
				return
			}
			
			err = rw.Flush()
			logger.Infof("Flushed data to payload stream %s", stream.ID())
			if err != nil {
				// fmt.Println("Error flushing buffer")
				// panic(err)
				logger.Error("Error flushing to to stream")
				return
			}
		
	}
 }

func readPayload(stream network.Stream) {   
	peer := stream.Conn().RemotePeer()
   rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
  
   for {
	   pData, err := rw.ReadBytes('\n')
	   if err != nil  {
		   logger.Errorf("Error reading from buffer %o", err)
		   return
	   }
	   if pData == nil {
		   //break
		   return
	   }

	   payload, err := UnpackP2pPayload(pData)
	   logger.Infof("Received Data: %v", payload)
	   if err != nil {
		   logger.WithFields(logrus.Fields{ "data": payload}).Warnf("Failed to parse payload: %o", err)
		   return
		   // break
	   }
	   validPayload := payload.IsValid(config.AddressPrefix)
	   
	   
	   logger.Infof("Validating peer %s", (stream).Conn().RemotePeer())
	   if !validPayload {
		  
		   return
	   }
	   if payload.Action  == P2pActionResponse {
		P2pComChannels[payload.Id][P2pChannelIn] <- payload
		return
	   }

	   response, _ := NewP2pPayload(config, P2pActionResponse, []byte{}, config.PrivateKey)
	   delimeter :=  []byte{'\n'}
	   switch  payload.Action {
		case P2pActionGetEvent:
			// data is the event path
			logger.Infof("Getting event for peer... %s", peer.String())
			eventPath, err := entities.UnpackEventPath(payload.Data)
			if err != nil {
				response.ResponseCode = 500
				response.Error = err.Error()
			}
			event, err := query.GetEventFromPath(&eventPath)
			if err == query.ErrorNotFound {
				response.ResponseCode = 404
				response.Error = "Event not found"
			}
			if err == nil {
				response.Data = event.MsgPack()
			}
			
			response.Sign(config.PrivateKey)
			rw.Write(append(response.MsgPack(), delimeter...))

	   }
	  
	//    switch payload.Action {
	//    case P2pActionRequestProof:
	// 	batch, err := entities.UnpackRewardBatch(payload.Data)
	// 	if err != nil {
	// 		logger.Errorf("HandleP2pPayload: %v", err)
	// 	}
	// 	// loop through reward batch, and ensure no duplicate

	//    }
	}
}


// check the dht before going onchain
func GetCycleMessageCost(ctx *context.Context, cycle uint64) (*big.Int, error) {
	
	priceKey :=  fmt.Sprintf("/ml/cost/%d", cycle)
	priceByte, err := idht.GetValue(*ctx, priceKey)
	// 
	if err != nil {
		return nil, err
	}
	if len(priceByte) > 0 {
		priceData, err := UnpackMessagePrice(priceByte)
		if err != nil {
			return getAndSaveMessageCostFromChain(ctx, cycle)
		}
		return big.NewInt(0).SetBytes(priceData.Price), nil
	} else {
		return  getAndSaveMessageCostFromChain(ctx, cycle)
	}
}

func getAndSaveMessageCostFromChain(ctx *context.Context, cycle uint64) (*big.Int, error) {
	price, err := chain.API.GetMessageCost(cycle)
	if err != nil {
		return nil, err
	}
	priceKey :=  fmt.Sprintf("/ml/cost/%d", cycle)
	mp, err := NewMessagePrice(config, config.PrivateKey, price.Bytes())
	if err != nil {
		return price, err
	}
	err = idht.PutValue(*ctx, priceKey, mp.MsgPack())
	return price, err
}
