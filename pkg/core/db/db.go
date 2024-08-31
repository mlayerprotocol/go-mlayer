package db

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	ds "github.com/ipfs/go-datastore"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
)

func Key(key string) ds.Key {
	return ds.NewKey(key)
}

func New(mainCtx *context.Context, keyStore string) *Datastore {
	ctx, cancel := context.WithCancel(*mainCtx)
	defer cancel()
	cfg, ok := ctx.Value(constants.ConfigKey).(*configs.MainConfiguration)
	if !ok {
		return nil
	}
	dir := filepath.Join(cfg.DataDir, "store", "kv", keyStore)
	if !strings.HasPrefix(dir, "./") && !strings.HasPrefix(dir, "../") && !filepath.IsAbs(dir) {
		dir = "./" + dir
		if strings.HasPrefix(cfg.DataDir, "../") {
			dir = "." + dir
		}
	}
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	ds, err := NewDatastore(dir, &DefaultOptions)
	if err != nil {
		panic(err)
	}
	return ds

	// err = db.View(func(txn *badger.Txn) error {
	// 	_, err := txn.Get([]byte("key"))
	// 	// We expect ErrKeyNotFound
	// 	fmt.Println("Error", err)
	// 	return nil
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// txn := db.NewTransaction(true) // Read-write txn
	// err = txn.SetEntry(badger.NewEntry([]byte("key"), []byte("value set successfully")))
	// if err != nil {
	// 	panic(err)
	// }
	// err = txn.Commit()
	// if err != nil {
	// 	panic(err)
	// }

	// err = db.View(func(txn *badger.Txn) error {
	// 	item, err := txn.Get([]byte("key"))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	val, err := item.ValueCopy(nil)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Printf("Valueeeee %s\n", string(val))
	// 	return nil
	// })

	// if err != nil {
	// 	panic(err)
	// }
}
