package notifee

import (
	"context"

	"github.com/libp2p/go-libp2p/core/peerstore"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
)
var logger = &log.Logger

type DiscoveryNotifee struct {
	Host host.Host
	Dht *dht.IpfsDHT
	HandleConnect func (*host.Host, *peer.AddrInfo)
}
func (n *DiscoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	logger.Infof("Discovered new peer in notifee %s\n", pi.ID.String())
	err := n.Host.Connect(context.Background(), pi)
	
	if err != nil {
		logger.Warningf("Unable to connect with peer: %s %o", pi.ID, err)
		return
	}
	if len(pi.ID) == 0 {
		return
	}
	n.Host.Peerstore().AddAddrs(pi.ID, pi.Addrs, peerstore.PermanentAddrTTL)
	n.HandleConnect(&n.Host, &pi)
}
