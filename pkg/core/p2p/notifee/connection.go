package notifee

import (
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/multiformats/go-multiaddr"
)


type ConnectionNotifee struct {
	Dht *dht.IpfsDHT
}

// Listen is called when network starts listening on an addr
func (n *ConnectionNotifee) Listen(netw network.Network, ma multiaddr.Multiaddr) {
	logger.Debugf("Listening.....")
}

// ListenClose is called when network starts listening on an addr
func (n *ConnectionNotifee) ListenClose(netw network.Network, ma multiaddr.Multiaddr) {}

// Connected is called when a connection opened
func (n *ConnectionNotifee) Connected(netw network.Network, conn network.Conn) {
	
	//retain max 4 connections
	// if (len(netw.Conns()) > 4){
	// 	conn.Close()
	// 	fmt.Printf("Connection refused for peer: %v!\n", conn.RemotePeer().Pretty())
	// }a
}

// Disconnected is called when a connection closed
func (cn *ConnectionNotifee) Disconnected(netw network.Network, conn network.Conn) {
	id := conn.RemotePeer()
	
	cn.Dht.Host().Peerstore().RemovePeer(id)
}

// OpenedStream is called when a stream opened
func (cn *ConnectionNotifee) OpenedStream(netw network.Network, stream network.Stream) {}

// ClosedStream is called when a stream was closed
func (cn *ConnectionNotifee) ClosedStream(netw network.Network, stream network.Stream) {}


