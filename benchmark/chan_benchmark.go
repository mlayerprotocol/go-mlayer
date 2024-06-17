package benchmark

import (
	"fmt"
	"log"
	"sync"
	"time"

	badger "github.com/dgraph-io/badger/v4"
)

// User struct
type User struct {
    ID   int
    Name string
}

// Worker function to process users from the channel and write to BadgerDB
func worker(db *badger.DB, userChan <-chan User, wg *sync.WaitGroup) {
    defer wg.Done()
    for user := range userChan {
        err := db.Update(func(txn *badger.Txn) error {
            key := []byte(fmt.Sprintf("%012duser", user.ID))
            val := []byte(user.Name)
            return txn.Set(key, val)
        })
        if err != nil {
            log.Printf("Error writing to BadgerDB: %v", err)
        }
    }
}

func main() {
    // Open BadgerDB
    opts := badger.DefaultOptions("./data/benchmark").WithInMemory(false) // Using in-memory mode for testing
    db, err := badger.Open(opts)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    var wg sync.WaitGroup
    userChan := make(chan User)

    // Start worker goroutines
    numWorkers := 100
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go worker(db, userChan, &wg)
    }
    var sendWg sync.WaitGroup
    // Simulate multiple routines sending User structs
    numRoutines := 100000
    t := time.Now().UnixMilli()
    for i := 0; i < numRoutines; i++ {
        sendWg.Add(1)
        go func(id int) {
            defer sendWg.Done()
            user := User{
                ID:   id*10 + i,
                Name: fmt.Sprintf("Usermanp dlknxapl dlkzl pzlcjvpzlij c;zljvc lzjcp;vlzj;clvjzl,nk;z lc;lz jvc;lzjc; lzj;vl c jz;clv pzocvpzcv hpzoxch v pozhv poz hcplkhz;clvkhzcl;kvhz;lkchv l;kzxch vl;kzhcv l;kzxhv lkzjcxh vlkzcxjhgv lkzxjcgh vlkzx cghvlkzxjhl kxcjhklx hlkxhzlkxhfpaokjd f;la j;dflkja;ldkj;lja;ldn;la jkn;aldnkv;aljkd;l ajkdf ;laj;lj;lajd;lnfa;lsdfj ;aldj f;lasdkj f;ladsjk f;lajds ;flaj d;flaj sd;lfj a;sld fa;ldjf ;lajsd f;lahsdj ;lasdj f;lasj ;flaj sdf;l jasd;lfj as;lfj as;ldkj s;lkj f;lakjd f;lasjk d;flkajs;lfjka s;ldj;klas-%d", id*10+i),
            }
            userChan <- user
           // time.Sleep(10 * time.Millisecond) // Simulate work
            // for j := 0; j < 10; j++ {
            //     user := User{
            //         ID:   id*10 + j,
            //         Name: fmt.Sprintf("User-%d", id*10+j),
            //     }
            //     userChan <- user
            //     time.Sleep(10 * time.Millisecond) // Simulate work
            // }
        }(i)
    }

    // Close the channel after all routines are done
    sendWg.Wait()
    close(userChan)
    
    // Wait for all workers to finish
    wg.Wait()
    log.Printf("Ended in %d milliseconds", time.Now().UnixMilli() - t)
    // Verify the stored data
    // err = db.View(func(txn *badger.Txn) error {
    //     opts := badger.DefaultIteratorOptions
    //     it := txn.NewIterator(opts)
    //     defer it.Close()

    //     for it.Rewind(); it.Valid(); it.Next() {
    //         item := it.Item()
    //         key := item.Key()
    //         err := item.Value(func(val []byte) error {
    //             fmt.Printf("Key: %s, Value: %s\n", key, val)
    //             return nil
    //         })
    //         if err != nil {
    //             return err
    //         }
    //     }
    //     return nil
    // })
    if err != nil {
        log.Printf("Error reading from BadgerDB: %v", err)
    }
}