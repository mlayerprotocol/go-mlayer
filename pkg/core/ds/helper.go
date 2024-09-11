package ds

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ipfs/go-datastore"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
)

const (
	SyncedBlockKey = "/syncedBlock"
)
func GetLastSyncedBlock(ctx *context.Context) (*big.Int, error) {
	systemStore, ok := (*ctx).Value(constants.SystemStore).(*Datastore)
	if !ok {
		return nil, fmt.Errorf("GetLastSyncBlock: unable to load systemStore from context")
	}
	lastBlockByte, err := systemStore.Get(*ctx, Key(SyncedBlockKey))
	if err != nil && err != datastore.ErrNotFound {
		return nil, fmt.Errorf("GetLastSyncedBlock: %v", err)
	}
	return new(big.Int).SetBytes(lastBlockByte), nil
}

func SetLastSyncedBlock(ctx *context.Context, value *big.Int) (error) {
	systemStore, ok := (*ctx).Value(constants.SystemStore).(*Datastore)
	if !ok {
		return  fmt.Errorf("GetLastSyncBlock: unable to load systemStore from context")
	}
	return systemStore.Set(*ctx, Key(SyncedBlockKey), value.Bytes(), true)
}