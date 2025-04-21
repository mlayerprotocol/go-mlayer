package p2p

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/ds"
)

var StateDhtSyncer dhtSyncer
type DhtPrefix string
type DhtSyncData struct {
    Key string
    Value []byte
}

const (
    AppDhtPrefix  DhtPrefix = "app"
    AppRefDhtPrefix  DhtPrefix = "appRef"
    ValDhtPrefix  DhtPrefix = "val"
    AppNullifierDhtPrefix  DhtPrefix = "apNul"
)
type dhtSyncer struct{
    store *ds.Datastore
    context *context.Context
}

// func init () {
//     StatedhtSyncer = NewdhtSyncer(stores.StateStore, context.Background())
// }


func NewDhtSyncer(store *ds.Datastore, ctx context.Context) dhtSyncer {
    dht := dhtSyncer{store, &ctx}
    // go dht.sync()
    return dht
}


// func (v *dhtSyncer) AddKeys(data []DhtSyncData) error {
//   batch :=  v.store.DB.NewWriteBatch()
//     for _, d := data {
//         batch.Set(ds.Key(d.Key).Bytes(), d.Value)
//     }
//   return  batch.Flush()
// }   


func (v *dhtSyncer) AddKey(dhtKeyPrefix DhtPrefix, dataKey string, value []byte) error {
    dhtKey := fmt.Sprintf("ml/%s/%s", dhtKeyPrefix,  dataKey)
  return   v.store.Set(context.Background(), datastore.NewKey(dhtKey), value, true)
}   

func (v *dhtSyncer) GetDhtValue(dhtKeyPrefix DhtPrefix, dataKey string) ([]byte, error) {
    dhtKey := fmt.Sprintf("ml/%s/%s", dhtKeyPrefix,  dataKey)
    data, err := idht.GetValue(*v.context, dhtKey, nil)
    if err == nil {
     return nil, err
    }
    return data, err
  } 

func (v *dhtSyncer) Sync() {
    for {
        if Initialized {
            notSynced,  err := v.store.Query(*v.context, query.Query{
                Prefix: "/dht",
            })
            if err != nil {
                continue
            }
           
            for result := range notSynced.Next() {
                parts := strings.Split(result.Entry.Key, "/")
                dhtKey := fmt.Sprintf("/ml/%s/%s", parts[len(parts)-2],  parts[len(parts)-1])
                err := idht.PutValue(*v.context, dhtKey, result.Entry.Value)
                if err == nil {
                    v.store.Delete(*v.context, ds.Key(result.Entry.Key) )
                } else {
                    logger.Errorf("FailedToSyncDHt: %v - %v", err, dhtKey)
                }
            }
        } else {
            logger.Infof("P2pNotInitialized...")
        }
        time.Sleep(1 * time.Second)
    }
}
