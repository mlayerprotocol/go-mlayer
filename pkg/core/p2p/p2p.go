package p2p

import (
	"bufio"
	"context"
	"fmt"
	"math"
	"os"
	"time"

	// "github.com/gin-gonic/gin"
	"github.com/ByteGum/go-icms/pkg/core/chain/evm"
	utils "github.com/ByteGum/go-icms/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	noise "github.com/libp2p/go-libp2p-noise"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	libp2ptls "github.com/libp2p/go-libp2p-tls"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/sirupsen/logrus"
	// rest "messagingprotocol/pkg/core/rest"
)

var logger = utils.Logger
var config utils.Configuration

var protocolId string

const DiscoveryServiceTag = "icm-network"
const (
	MessageChannel string = "icm-message-channel"
)

var peerStreams = make(map[string]peer.ID)
var peerPubKeys = make(map[peer.ID][]byte)
var node *host.Host

// defaultNick generates a nickname based on the $USER environment variable and
// the last 8 chars of a peer ID.
func defaultNick(p peer.ID) string {
	// TODO load name from flag/config
	return fmt.Sprintf("%s-%s", os.Getenv("USER"), shortID(p))
}

// shortID returns the last 8 chars of a base58-encoded peer id.
func shortID(p peer.ID) string {
	pretty := p.Pretty()
	return pretty[len(pretty)-12:]
}

func Run(mainCtx *context.Context) {
	// fmt.Printf("publicKey %s", privateKey)
	// The context governs the lifetime of the libp2p node.
	// Cancelling it will stop the the host.
	ctx, cancel := context.WithCancel(*mainCtx)
	cfg, ok := ctx.Value("Config").(utils.Configuration)
	if !ok {

	}
	config = cfg
	protocolId = config.Network

	incomingMessagesC, ok := ctx.Value("IncomingMessageC").(chan utils.ClientMessage)
	if !ok {

	}
	outgoinMessageC, ok := ctx.Value("IncomingMessageC").(chan utils.ClientMessage)
	if !ok {

	}

	defer cancel()

	// // To construct a simple host with all the default settings, just use `New`
	// h, err := libp2p.New(ctx)
	// if err != nil {
	// 	panic(err) s
	// }

	// r := gin.Default()
	// r = rest.SetupOriginatorRoutes(r)
	// r.Run("localhost:8080")

	// log.Printf("Hello World, my hosts ID is %s\n", h.ID())

	// Now, normally you do not just want a simple host, you want
	// that is fully configured to best support your p2p application.
	// Let's create a second host setting some more options.
	// Set your own keypaircsd
	priv, _, err := crypto.GenerateKeyPair(

		crypto.ECDSA, // Select your key type. Ed25519 are nice short
		-1,           // Select key length when possible (i.e. RSA).
	)
	// privK, _ := ethCrypto.HexToECDSA(privateKey)
	// priv, _, err := crypto.ECDSAKeyPairFromKey(privK)
	if err != nil {
		panic(err)
	}
	var idht *dht.IpfsDHT

	h, err := libp2p.New(
		// Use the keypair we generated
		libp2p.Identity(priv),
		// Multiple listen addresses
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/9000/ws",
			"/ip4/0.0.0.0/tcp/0",
		),
		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),
		// support noise connections
		libp2p.Security(noise.ID, noise.New),
		// support any other default transports (TCP)
		libp2p.DefaultTransports,
		// libp2p.Transport(ws.New),
		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.
		libp2p.ConnectionManager(connmgr.NewConnManager(
			100,         // Lowwater
			400,         // HighWater,
			time.Minute, // GracePeriod
		)),
		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),
		// Let this host use the DHT to find other hosts

		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			idht, err = dht.New(ctx, h)
			return idht, err
		}),

		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		libp2p.EnableAutoRelay(),
		// If you want to help other peers to figure out if they are behind
		// NATs, you can launch the server-side of AutoNAT too (AutoRelay
		// already runs the client)
		//
		// This service is highly rate-limited and should not cause any
		// performance issues.
		libp2p.EnableNATService(),
	)
	if err != nil {
		panic(err)
	}

	h.SetStreamHandler(protocol.ID(protocolId), handleStream)
	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}
	// setup local mDNS discovery
	err = setupDiscovery(ctx, h)
	if err != nil {
		panic(err)
	}
	node = &h

	// The last step to get fully up and running would be to connect to
	// bootstrap peers (or any other peers). We leave this commented as
	// this is an example and the peer will die as soon as it finishes, so
	// it is unnecessary to put strain on the network.

	// for _, addr := range dht.DefaultBootstrapPeers {
	// 	pi, _ := peer.AddrInfoFromP2pAddr(addr)
	// 	// We ignore errors as some bootstrap peers may be down
	// 	// and that is fine.
	// 	h.Connect(ctx, *pi)
	// }

	logger.Infof("Host started with ID is %s\n", h.ID())

	cr, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), MessageChannel, config.ChannelMessageBufferSize)
	if err != nil {
		panic(err)
	}
	logger.WithFields(logrus.Fields{"event": "JoinChannel", "peer": h.ID()}).Infof("Peer joined channel %s", cr.ChannelName)
	go func() {
		time.Sleep(10 * time.Second)
		for {
			select {
			case inMessage, ok := <-cr.Messages:
				if !ok {
					logger.Errorf("Primary Message channel closed. Please restart server to try or adjust buffer size in config")
					cancel()
					return
				}
				// msg, err := d.ToJSON()
				// if err != nil {
				// 	continue
				// }
				logger.Info("Received new message %s\n", inMessage.Message.Body.Text)
				incomingMessagesC <- *inMessage
			}
		}
	}()
	for {
		select {
		case outMessage, ok := <-outgoinMessageC:
			if !ok {
				logger.Errorf("Outgoing channel closed. Please restart server to try or adjust buffer size in config")
				return
			}
			err := cr.Publish(outMessage)
			if err != nil {
				logger.Errorf("Failed to publish message. Please restart server to try or adjust buffer size in config")

				return
			}
		}
	}

}

// printErr is like fmt.Printf, but writes to stderr.
// func printErr(m string, args ...interface{}) {
// 	fmt.Fprintf(os.Stderr, m, args...)
// }

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h host.Host
}

func handleStream(stream network.Stream) {
	// logger.Debugf("Got a new stream! %s", stream.ID())
	// stream.SetReadDeadline()
	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	go readData(peerStreams[stream.ID()], rw)
	// go sendData(rw)

}

func readData(p peer.ID, rw *bufio.ReadWriter) {
	for {
		hsString, err := rw.ReadString('\n')
		if err != nil {
			logger.Errorf("Error reading from buffer %w", err)
			panic(err)
		}
		if hsString == "" {
			break
		}
		logger.WithFields(logrus.Fields{"peer": p, "data": hsString}).Info("New Handshake data from peer")
		handshake, err := utils.HandshakeFromJSON(hsString)
		if err != nil {
			logger.WithFields(logrus.Fields{"peer": p, "data": hsString}).Warnf("Failed to parse handshake: %w", err)
			break
		}
		validHandshke := isValidHandshake(handshake, p)
		if !validHandshke {
			disconnect(*node, p)
			logger.WithFields(logrus.Fields{"peer": p, "data": hsString}).Warnf("Disconnecting from peer (%s) with invalid handshake", p)
			return
		}
		validStake := isValidStake(handshake, p)
		if !validStake {
			disconnect(*node, p)
			logger.WithFields(logrus.Fields{"address": handshake.Signer, "data": hsString}).Warnf("Disconnecting from peer (%s) with inadequate stake in network", p)
			return
		}
		b, _ := hexutil.Decode(handshake.Signer)
		peerPubKeys[p] = b
		break
	}
}

func isValidHandshake(handshake utils.Handshake, p peer.ID) bool {
	handshakeMessage := handshake.ToJSON()
	if math.Abs(float64(handshake.Data.Timestamp-int(time.Now().Unix()))) > utils.VALID_HANDSHAKE_SECONDS {
		logger.WithFields(logrus.Fields{"peer": p, "data": handshakeMessage}).Warnf("Hanshake Expired: %s", handshakeMessage)
		return false
	}
	message := handshake.Data.ToString()
	isValid := utils.VerifySignature(handshake.Signer, message, handshake.Signature)
	if !isValid {
		logger.WithFields(logrus.Fields{"message": message, "signature": handshake.Signature}).Warnf("Invalid signer %s", handshake.Signer)
		return false
	}
	logger.Infof("New Valid handshake from peer: %s", p)
	return true
}
func isValidStake(handshake utils.Handshake, p peer.ID) bool {
	stakeContract, err := evm.StakeContract(config.RPCUrl, config.StakeContract)
	if err != nil {
		logger.Errorf("RPC error %w", err)
	}
	level, err := stakeContract.GetAccountLevel(nil, evm.ToHexAddress(handshake.Signer))
	if level == utils.StandardAccountType {
		logger.WithFields(logrus.Fields{"address": handshake.Signer, "accountType": level}).Warnf("Inadequate stake balance for peer %s", p)
		return false
	}
	return true
}
func sendData(p peer.ID, rw *bufio.ReadWriter, data []byte) {

	// defer disconnect(*node, p)
	_, err := rw.WriteString(fmt.Sprintf("%s\n", string(data)))
	if err != nil {
		logger.Warn("Error writing to to stream")
		return
	}
	err = rw.Flush()
	if err != nil {
		// fmt.Println("Error flushing buffer")
		// panic(err)
		logger.Warn("Error flushing to to stream")
		return
	}
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	logger.Infof("Discovered new peer %s\n", pi.ID.Pretty())
	err := n.h.Connect(context.Background(), pi)

	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := n.h.NewStream(ctx, pi.ID, protocol.ID(protocolId))

	if err != nil {
		logger.Warningf("Unable to establish stream with peer: %s %w", pi.ID, err)
	} else {
		logger.Infof("Streaming to peer: %s", pi.ID)
		rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
		logger.Infof("New StreamID: %s", stream.ID())
		peerStreams[stream.ID()] = pi.ID
		hs := utils.CreateHandshake(defaultNick(n.h.ID()), protocolId, config.PrivateKey)
		go sendData(pi.ID, rw, (&hs).ToJSON())
		// go readData(rw)
	}
}

func disconnect(h host.Host, id peer.ID) {
	h.Network().ClosePeer(id)
}

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func setupDiscovery(ctx context.Context, h host.Host) error {
	n := discoveryNotifee{h: h}
	// n.h = make(chan peer.AddrInfo)
	// setup mDNS discovery to find local peers
	disc := mdns.NewMdnsService(h, DiscoveryServiceTag, &n)
	if err := disc.Start(); err != nil {
		panic(err)
	}
	// disc.RegisterNotifee(&n)
	return nil
}
