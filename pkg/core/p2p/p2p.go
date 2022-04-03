package p2p

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	// "github.com/gin-gonic/gin"

	utils "github.com/fero-tech/splanch/utils"
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
	// rest "messagingprotocol/pkg/core/rest"
)

var logger = utils.Logger()

// defaultNick generates a nickname based on the $USER environment variable and
// the last 8 chars of a peer ID.
func defaultNick(p peer.ID) string {
	return fmt.Sprintf("%s-%s", os.Getenv("USER"), shortID(p))
}

// shortID returns the last 8 chars of a base58-encoded peer id.
func shortID(p peer.ID) string {
	pretty := p.Pretty()
	return pretty[len(pretty)-8:]
}

func Run(protocolId string, privateKey string) {
	fmt.Printf("publicKey %s", privateKey)
	// The context governs the lifetime of the libp2p node.
	// Cancelling it will stop the the host.
	ctx, cancel := context.WithCancel(context.Background())
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
	logger.Debug("PrivKeysssss %s", priv.GetPublic())
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
	h.SetStreamHandler(protocol.ID(protocolId), handleStream)

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

	logger.Debug("Hello World, my host ID is %s\n", h.ID())
	// use the nickname from the cli flag, or a default if blank
	nickFlag := flag.String("nick", "", "nickname to use in chat. will be generated if empty")
	roomFlag := flag.String("room", "awesome-chat-room", "name of chat room to join")
	flag.Parse()

	nick := *nickFlag

	if len(nick) == 0 {
		nick = defaultNick(h.ID())
	}

	// join the room from the cli flag, or the flag default
	room := *roomFlag

	// join the chat room
	cr, err := JoinChannel(ctx, ps, h.ID(), nick, room)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s", cr.ChannelName)

	// draw the UI
	ui := NewChatUI(cr)
	if err = ui.Run(); err != nil {
		logger.Error("error running text UI: %s", err)
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
	logger.Debug("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	go readHandshake(rw)
	go writeData(rw)

	// 'stream' will stay open until you close it (or the other side closes it).
}
func readHandshake(rw *bufio.ReadWriter) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			logger.Error("Error reading from buffer")
			panic(err)
		}

		if str == "" {
			return
		}
		if str != "\n" {
			// Green console colour: 	\x1b[32m
			// Reset console colour: 	\x1b[0m
			fmt.Printf("\x1b[32m%s\x1b[0m> \n", str)
		}

	}
}
func writeData(rw *bufio.ReadWriter) {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		logger.Debug("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin")
			panic(err)
		}

		_, err = rw.WriteString(fmt.Sprintf("%s\n", sendData))
		if err != nil {
			fmt.Println("Error writing to buffer")
			panic(err)
		}
		err = rw.Flush()
		if err != nil {
			fmt.Println("Error flushing buffer")
			panic(err)
		}
	}
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	fmt.Printf("discovered new peer %s\n", pi.ID.Pretty())
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
}

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const DiscoveryServiceTag = "pubsub-chat-example"

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
