package db

import (
	"fmt"
	"io/ioutil"

	badger "github.com/dgraph-io/badger"
)

func Db()  {
	dir, err := ioutil.TempDir("", "badger-test")
	if err != nil {
		panic(err)
	}
	//defer ioutil.removeDir(dir)
	db, err := badger.Open(badger.DefaultOptions(dir))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte("key"))
		// We expect ErrKeyNotFound
		fmt.Println("Error", err)
		return nil
	})

	if err != nil {
		panic(err)
	}

	txn := db.NewTransaction(true) // Read-write txn
	err = txn.SetEntry(badger.NewEntry([]byte("key"), []byte("value set successfully")))
	if err != nil {
		panic(err)
	}
	err = txn.Commit()
	if err != nil {
		panic(err)
	}

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("key"))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		fmt.Printf("Valueeeee %s\n", string(val))
		return nil
	})

	if err != nil {
		panic(err)
	}
}
