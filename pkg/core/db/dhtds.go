package db

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"runtime"
// 	"strings"
// 	"sync"
// 	"time"

// 	badger "github.com/dgraph-io/badger/v4"
// 	ds "github.com/ipfs/go-datastore"
// 	dsq "github.com/ipfs/go-datastore/query"
// 	goprocess "github.com/jbenet/goprocess"
// )

// type DhtDatastore struct {
// 	DB *badger.DB

// 	closeLk   sync.RWMutex
// 	closed    bool
// 	closeOnce sync.Once
// 	closing   chan struct{}

// 	gcDiscardRatio float64
// 	gcSleep        time.Duration
// 	gcInterval     time.Duration

// 	syncWrites bool
// }

// func init() {
// 	DefaultOptions = Options{
// 		GcDiscardRatio: 0.2,
// 		GcInterval:     15 * time.Minute,
// 		GcSleep:        10 * time.Second,
// 		Options:        badger.LSMOnlyOptions(""),
// 	}
// 	// This is to optimize the database on close so it can be opened
// 	// read-only and efficiently queried. We don't do that and hanging on
// 	// stop isn't nice.
// 	DefaultOptions.Options.CompactL0OnClose = false

// }

// var _ ds.Datastore = (*DhtDatastore)(nil)
// var _ ds.TxnDatastore = (*DhtDatastore)(nil)
// var _ ds.TTLDatastore = (*DhtDatastore)(nil)
// var _ ds.GCDatastore = (*DhtDatastore)(nil)
// var _ ds.Batching = (*DhtDatastore)(nil)

// // NewDatastore creates a new badger datastore.
// //
// // DO NOT set the Dir and/or ValuePath fields of opt, they will be set for you.
// func NewDhtDatastore(path string, options *Options) (*DhtDatastore, error) {
// 	// Copy the options because we modify them.
// 	var opt badger.Options
// 	var gcDiscardRatio float64
// 	var gcSleep time.Duration
// 	var gcInterval time.Duration
// 	if options == nil {
// 		opt = badger.DefaultOptions(path).WithInMemory(false)
// 		gcDiscardRatio = DefaultOptions.GcDiscardRatio
// 		gcSleep = DefaultOptions.GcSleep
// 		gcInterval = DefaultOptions.GcInterval
// 	} else {
// 		opt = options.Options
// 		gcDiscardRatio = options.GcDiscardRatio
// 		gcSleep = options.GcSleep
// 		gcInterval = options.GcInterval
// 	}

// 	if gcSleep <= 0 {
// 		// If gcSleep is 0, we don't perform multiple rounds of GC per
// 		// cycle.
// 		gcSleep = gcInterval
// 	}

// 	opt.Dir = path
// 	opt.ValueDir = path
// 	opt.Logger = &badgerLog{*log}

// 	kv, err := badger.Open(opt)

// 	if err != nil {
// 		if strings.HasPrefix(err.Error(), "manifest has unsupported version:") {
// 			err = fmt.Errorf("unsupported badger version, use github.com/ipfs/badgerds-upgrade to upgrade: %s", err.Error())
// 		}
// 		return nil, err
// 	}

// 	ds := &DhtDatastore{
// 		DB:             kv,
// 		closing:        make(chan struct{}),
// 		gcDiscardRatio: gcDiscardRatio,
// 		gcSleep:        gcSleep,
// 		gcInterval:     gcInterval,
// 		syncWrites:     opt.SyncWrites,
// 	}

// 	// Start the GC process if requested.
// 	if ds.gcInterval > 0 {
// 		go ds.periodicGC()
// 	}

// 	return ds, nil
// }

// // Keep scheduling GC's AFTER `gcInterval` has passed since the previous GC
// func (d *DhtDatastore) periodicGC() {
// 	gcTimeout := time.NewTimer(d.gcInterval)
// 	defer gcTimeout.Stop()

// 	for {
// 		select {
// 		case <-gcTimeout.C:
// 			switch err := d.gcOnce(); err {
// 			case badger.ErrNoRewrite, badger.ErrRejected:
// 				// No rewrite means we've fully garbage collected.
// 				// Rejected means someone else is running a GC
// 				// or we're closing.
// 				gcTimeout.Reset(d.gcInterval)
// 			case nil:
// 				gcTimeout.Reset(d.gcSleep)
// 			case ErrClosed:
// 				return
// 			default:
// 				log.Errorf("error during a GC cycle: %s", err)
// 				// Not much we can do on a random error but log it and continue.
// 				gcTimeout.Reset(d.gcInterval)
// 			}
// 		case <-d.closing:
// 			return
// 		}
// 	}
// }

// func (d *DhtDatastore) Put(ctx context.Context, key ds.Key, value []byte) error {
// 	logger.Debugf("KEYYYYY %s value %d vvv %v", string(key.Bytes()))
// 	// return d.DB.Update(func(txn *badger.Txn) error {
// 	// 	return txn.Set(key.Bytes(), value)
// 	// })

// 	//  d.closeLk.RLock()
// 	// defer d.closeLk.RUnlock()
// 	// if d.closed {
// 	// 	return ErrClosed
// 	// }

// 	// txn := d.newImplicitTransaction(false)
// 	// defer txn.discard()

// 	// if err := txn.Set(ctx, key, value, true); err != nil {
// 	// 	return err
// 	// }

// 	// return txn.commit()
// 	return nil
// }

// func (d *DhtDatastore) Get(ctx context.Context, key ds.Key) (value []byte, err error) {

// 	err = d.DB.View(func(txn *badger.Txn) error {

//         item, err := txn.Get(key.Bytes())

//         if err != nil {
//             return err
//         }
//         err = item.Value(func(val []byte) error {
//             // Copying the value to retrievedValue
//             value = append([]byte{}, val...)
//             return nil
//         })
//         return err
//     })
// 	return value, err
// }

// func (d *DhtDatastore) Has(ctx context.Context, key ds.Key) (bool, error) {
// 	err := d.DB.View(func(txn *badger.Txn) error {
//         _, err := txn.Get(key.Bytes())
//         return err
//     })
// 	switch err {
// 	case badger.ErrKeyNotFound:
// 		return false, nil
// 	case nil:
// 		return true, nil
// 	default:
// 		return false, err
// 	}
// }

// func (d *DhtDatastore) Delete(ctx context.Context, key ds.Key) error {
// 	err := d.DB.Update(func(txn *badger.Txn) error {
//         err := txn.Delete(key.Bytes())
//         return err
//     })
// 	if err != nil {
// 		return err
// 	}
// }

// func (d *DhtDatastore) Query(ctx context.Context, q dsq.Query) (result dsq.Results, err error) {

// 	err = d.DB.View(func(txn *badger.Txn) error {
//         opt := badger.DefaultIteratorOptions
// 		opt.PrefetchValues = !q.KeysOnly
// 		prefix := ds.NewKey(q.Prefix).String()
// 		if prefix != "/" {
// 			opt.Prefix = []byte(prefix + "/")
// 		}

// 		// Handle ordering
// 		if len(q.Orders) > 0 {
// 			switch q.Orders[0].(type) {
// 			case dsq.OrderByKey, *dsq.OrderByKey:
// 			// We order by key by default.
// 			case dsq.OrderByKeyDescending, *dsq.OrderByKeyDescending:
// 				// Reverse order by key
// 				opt.Reverse = true
// 			default:
// 				// Ok, we have a weird order we can't handle. Let's
// 				// perform the _base_ query (prefix, filter, etc.), then
// 				// handle sort/offset/limit later.

// 				// Skip the stuff we can't apply.
// 				baseQuery := q
// 				baseQuery.Limit = 0
// 				baseQuery.Offset = 0
// 				baseQuery.Orders = nil

// 				// perform the base query.
// 				res, err := t.query(baseQuery)
// 				if err != nil {
// 					return nil, err
// 				}

// 				// fix the query
// 				res = dsq.ResultsReplaceQuery(res, q)

// 				// Remove the parts we've already applied.
// 				naiveQuery := q
// 				naiveQuery.Prefix = ""
// 				naiveQuery.Filters = nil

// 				// Apply the rest of the query
// 				result = dsq.NaiveQueryApply(naiveQuery, res)
// 				return err
// 			}
// 		}
// 		ch := make(chan dsq.Result, 100)
// 		result = dsq.ResultsWithChan(q, ch)
// 		it := txn.NewIterator(opt)
// 		for it.Rewind(); it.ValidForPrefix([]byte(prefix)); it.Next() {
// 			item := it.Item()
// 			key := item.Key()
// 			err := item.Value(func(val []byte) error {
// 				entry := dsq.Entry{
// 					Key:   string(key),
// 					Value: val,
// 				}
// 				ch <- dsq.Result{
// 					Entry: entry,
// 				}
// 				return nil
// 			})
// 			if err != nil {
// 				return err
// 			}
// 		}

// 	if err != nil {
// 		ch <- dsq.Result{
// 			Error: err,
// 		}
// 	}

// 		return err
//     })
// 	return result, err
// }

// func (d *DhtDatastore) Close() error {
// 	d.closeOnce.Do(func() {
// 		close(d.closing)
// 	})
// 	d.closeLk.Lock()
// 	defer d.closeLk.Unlock()
// 	if d.closed {
// 		return ErrClosed
// 	}
// 	d.closed = true
// 	return d.DB.Close()
// }
// type dhtBatch struct {
// 	ds         *DhtDatastore
// 	writeBatch *badger.WriteBatch
// }

// // Batch creats a new Batch object. This provides a way to do many writes, when
// // there may be too many to fit into a single transaction.
// // func (d *DhtDatastore) Batch(ctx context.Context) (ds.Batch, error) {
// // 	d.closeLk.RLock()
// // 	defer d.closeLk.RUnlock()
// // 	if d.closed {
// // 		return nil, ErrClosed
// // 	}

// // 	b := &batch{d, d.DB.NewWriteBatch()}
// // 	// Ensure that incomplete transaction resources are cleaned up in case
// // 	// batch is abandoned.
// // 	// runtime.SetFinalizer(b, func(b *batch) {
// // 	// 	b.cancel()
// // 	// 	log.Error("batch not committed or canceled")
// // 	// })

// // 	return b, nil
// // }

// func (d *DhtDatastore) CollectGarbage(ctx context.Context) (err error) {
// 	// The idea is to keep calling DB.RunValueLogGC() till Badger no longer has any log files
// 	// to GC(which would be indicated by an error, please refer to Badger GC docs).
// 	for err == nil {
// 		err = d.gcOnce()
// 	}

// 	if err == badger.ErrNoRewrite {
// 		err = nil
// 	}

// 	return err
// }

// func (d *DhtDatastore) gcOnce() error {
// 	d.closeLk.RLock()
// 	defer d.closeLk.RUnlock()
// 	if d.closed {
// 		return ErrClosed
// 	}
// 	log.Info("Running GC round")
// 	defer log.Info("Finished running GC round")
// 	return d.DB.RunValueLogGC(d.gcDiscardRatio)
// }

// var _ ds.Batch = (*batch)(nil)

// func (b *batch) Put(ctx context.Context, key ds.Key, value []byte) error {
// 	// b.ds.closeLk.RLock()
// 	// defer b.ds.closeLk.RUnlock()
// 	// if b.ds.closed {
// 	// 	return ErrClosed
// 	// }
// 	// return b.ds.Set(ctx, key, value, true)
// 	return nil
// }

// func (b *batch) put(key ds.Key, value []byte) error {
// 	return b.writeBatch.Set(key.Bytes(), value)
// }

// func (b *batch) Delete(ctx context.Context, key ds.Key) error {
// 	b.ds.closeLk.RLock()
// 	defer b.ds.closeLk.RUnlock()
// 	if b.ds.closed {
// 		return ErrClosed
// 	}

// 	return b.delete(key)
// }

// func (b *batch) delete(key ds.Key) error {
// 	return b.writeBatch.Delete(key.Bytes())
// }

// func (b *batch) Commit(ctx context.Context) error {
// 	b.ds.closeLk.RLock()
// 	defer b.ds.closeLk.RUnlock()
// 	if b.ds.closed {
// 		return ErrClosed
// 	}

// 	return b.commit()
// }

// func (b *batch) commit() error {
// 	err := b.writeBatch.Flush()
// 	if err != nil {
// 		// Discard incomplete transaction held by b.writeBatch
// 		b.cancel()
// 		return err
// 	}
// 	runtime.SetFinalizer(b, nil)
// 	return nil
// }

// func (b *batch) Cancel() error {
// 	b.ds.closeLk.RLock()
// 	defer b.ds.closeLk.RUnlock()
// 	if b.ds.closed {
// 		return ErrClosed
// 	}

// 	b.cancel()
// 	return nil
// }

// func (b *batch) cancel() {
// 	b.writeBatch.Cancel()
// 	runtime.SetFinalizer(b, nil)
// }

// var _ ds.DhtDatastore = (*txn)(nil)
// var _ ds.TTLDatastore = (*txn)(nil)

// func (t *txn) Put(ctx context.Context, key ds.Key, value []byte) error {
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return ErrClosed
// 	}
// 	return t.Set(ctx, key, value, true)
// }

// func (t *txn) Set(ctx context.Context, key ds.Key, value []byte, replace bool) error {
// 	if !replace {
// 		has, err := t.has(key)
// 		if has && err == nil {
// 			return errors.New("Key already exists")
// 		}
// 	}
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return ErrClosed
// 	}
// 	return t.put(key, value)
// }

// func (t *txn) put(key ds.Key, value []byte) error {
// 	return t.txn.Set(key.Bytes(), value)
// }

// func (t *txn) Sync(ctx context.Context, prefix ds.Key) error {
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return ErrClosed
// 	}

// 	return nil
// }

// func (t *txn) PutWithTTL(ctx context.Context, key ds.Key, value []byte, ttl time.Duration) error {
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return ErrClosed
// 	}
// 	return t.putWithTTL(key, value, ttl)
// }

// func (t *txn) putWithTTL(key ds.Key, value []byte, ttl time.Duration) error {
// 	return t.txn.SetEntry(badger.NewEntry(key.Bytes(), value).WithTTL(ttl))
// }

// func (t *txn) GetExpiration(ctx context.Context, key ds.Key) (time.Time, error) {
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return time.Time{}, ErrClosed
// 	}

// 	return t.getExpiration(key)
// }

// func (t *txn) getExpiration(key ds.Key) (time.Time, error) {
// 	item, err := t.txn.Get(key.Bytes())
// 	if err == badger.ErrKeyNotFound {
// 		return time.Time{}, ds.ErrNotFound
// 	} else if err != nil {
// 		return time.Time{}, err
// 	}
// 	return time.Unix(int64(item.ExpiresAt()), 0), nil
// }

// func (t *txn) SetTTL(ctx context.Context, key ds.Key, ttl time.Duration) error {
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return ErrClosed
// 	}

// 	return t.setTTL(key, ttl)
// }

// func (t *txn) setTTL(key ds.Key, ttl time.Duration) error {
// 	item, err := t.txn.Get(key.Bytes())
// 	if err != nil {
// 		return err
// 	}
// 	return item.Value(func(data []byte) error {
// 		return t.putWithTTL(key, data, ttl)
// 	})

// }

// func (t *txn) Get(ctx context.Context, key ds.Key) ([]byte, error) {
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return nil, ErrClosed
// 	}

// 	return t.get(key)
// }

// func (t *txn) get(key ds.Key) ([]byte, error) {
// 	item, err := t.txn.Get(key.Bytes())
// 	if err == badger.ErrKeyNotFound {
// 		err = ds.ErrNotFound
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	return item.ValueCopy(nil)
// }

// func (t *txn) Has(ctx context.Context, key ds.Key) (bool, error) {
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return false, ErrClosed
// 	}

// 	return t.has(key)
// }

// func (t *txn) has(key ds.Key) (bool, error) {
// 	_, err := t.txn.Get(key.Bytes())
// 	switch err {
// 	case badger.ErrKeyNotFound:
// 		return false, nil
// 	case nil:
// 		return true, nil
// 	default:
// 		return false, err
// 	}
// }

// func (t *txn) GetSize(ctx context.Context, key ds.Key) (int, error) {
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return -1, ErrClosed
// 	}

// 	return t.getSize(key)
// }

// func (t *txn) getSize(key ds.Key) (int, error) {
// 	item, err := t.txn.Get(key.Bytes())
// 	switch err {
// 	case nil:
// 		return int(item.ValueSize()), nil
// 	case badger.ErrKeyNotFound:
// 		return -1, ds.ErrNotFound
// 	default:
// 		return -1, err
// 	}
// }

// func (t *txn) Delete(ctx context.Context, key ds.Key) error {
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return ErrClosed
// 	}

// 	return t.delete(key)
// }

// func (t *txn) delete(key ds.Key) error {
// 	return t.txn.Delete(key.Bytes())
// }

// func (t *txn) Query(ctx context.Context, q dsq.Query) (dsq.Results, error) {
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return nil, ErrClosed
// 	}

// 	return t.query(q)
// }

// func (t *txn) query(q dsq.Query) (dsq.Results, error) {
// 	opt := badger.DefaultIteratorOptions
// 	opt.PrefetchValues = !q.KeysOnly
// 	prefix := ds.NewKey(q.Prefix).String()
// 	if prefix != "/" {
// 		opt.Prefix = []byte(prefix + "/")
// 	}

// 	// Handle ordering
// 	if len(q.Orders) > 0 {
// 		switch q.Orders[0].(type) {
// 		case dsq.OrderByKey, *dsq.OrderByKey:
// 		// We order by key by default.
// 		case dsq.OrderByKeyDescending, *dsq.OrderByKeyDescending:
// 			// Reverse order by key
// 			opt.Reverse = true
// 		default:
// 			// Ok, we have a weird order we can't handle. Let's
// 			// perform the _base_ query (prefix, filter, etc.), then
// 			// handle sort/offset/limit later.

// 			// Skip the stuff we can't apply.
// 			baseQuery := q
// 			baseQuery.Limit = 0
// 			baseQuery.Offset = 0
// 			baseQuery.Orders = nil

// 			// perform the base query.
// 			res, err := t.query(baseQuery)
// 			if err != nil {
// 				return nil, err
// 			}

// 			// fix the query
// 			res = dsq.ResultsReplaceQuery(res, q)

// 			// Remove the parts we've already applied.
// 			naiveQuery := q
// 			naiveQuery.Prefix = ""
// 			naiveQuery.Filters = nil

// 			// Apply the rest of the query
// 			return dsq.NaiveQueryApply(naiveQuery, res), nil
// 		}
// 	}

// 	it := t.txn.NewIterator(opt)
// 	qrb := dsq.NewResultBuilder(q)
// 	qrb.Process.Go(func(worker goprocess.Process) {
// 		t.ds.closeLk.RLock()
// 		closedEarly := false
// 		defer func() {
// 			t.ds.closeLk.RUnlock()
// 			if closedEarly {
// 				select {
// 				case qrb.Output <- dsq.Result{
// 					Error: ErrClosed,
// 				}:
// 				case <-qrb.Process.Closing():
// 				}
// 			}

// 		}()
// 		if t.ds.closed {
// 			closedEarly = true
// 			return
// 		}

// 		// this iterator is part of an implicit transaction, so when
// 		// we're done we must discard the transaction. It's safe to
// 		// discard the txn it because it contains the iterator only.
// 		if t.implicit {
// 			defer t.discard()
// 		}

// 		defer it.Close()

// 		// All iterators must be started by rewinding.
// 		it.Rewind()

// 		// skip to the offset
// 		for skipped := 0; skipped < q.Offset && it.Valid(); it.Next() {
// 			// On the happy path, we have no filters and we can go
// 			// on our way.
// 			if len(q.Filters) == 0 {
// 				skipped++
// 				continue
// 			}

// 			// On the sad path, we need to apply filters before
// 			// counting the item as "skipped" as the offset comes
// 			// _after_ the filter.
// 			item := it.Item()

// 			matches := true
// 			check := func(value []byte) error {
// 				e := dsq.Entry{
// 					Key:   string(item.Key()),
// 					Value: value,
// 					Size:  int(item.ValueSize()), // this function is basically free
// 				}

// 				// Only calculate expirations if we need them.
// 				if q.ReturnExpirations {
// 					e.Expiration = expires(item)
// 				}
// 				matches = filter(q.Filters, e)
// 				return nil
// 			}

// 			// Maybe check with the value, only if we need it.
// 			var err error
// 			if q.KeysOnly {
// 				err = check(nil)
// 			} else {
// 				err = item.Value(check)
// 			}

// 			if err != nil {
// 				select {
// 				case qrb.Output <- dsq.Result{Error: err}:
// 				case <-t.ds.closing: // datastore closing.
// 					closedEarly = true
// 					return
// 				case <-worker.Closing(): // client told us to close early
// 					return
// 				}
// 			}
// 			if !matches {
// 				skipped++
// 			}
// 		}

// 		for sent := 0; (q.Limit <= 0 || sent < q.Limit) && it.Valid(); it.Next() {
// 			item := it.Item()
// 			e := dsq.Entry{Key: string(item.Key())}

// 			// Maybe get the value
// 			var result dsq.Result
// 			if !q.KeysOnly {
// 				b, err := item.ValueCopy(nil)
// 				if err != nil {
// 					result = dsq.Result{Error: err}
// 				} else {
// 					e.Value = b
// 					e.Size = len(b)
// 					result = dsq.Result{Entry: e}
// 				}
// 			} else {
// 				e.Size = int(item.ValueSize())
// 				result = dsq.Result{Entry: e}
// 			}

// 			if q.ReturnExpirations {
// 				result.Expiration = expires(item)
// 			}

// 			// Finally, filter it (unless we're dealing with an error).
// 			if result.Error == nil && filter(q.Filters, e) {
// 				continue
// 			}

// 			select {
// 			case qrb.Output <- result:
// 				sent++
// 			case <-t.ds.closing: // datastore closing.
// 				closedEarly = true
// 				return
// 			case <-worker.Closing(): // client told us to close early
// 				return
// 			}
// 		}
// 	})

// 	go qrb.Process.CloseAfterChildren() //nolint

// 	return qrb.Results(), nil
// }

// func (t *txn) Commit(ctx context.Context) error {
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return ErrClosed
// 	}

// 	return t.commit()
// }

// func (t *txn) commit() error {
// 	return t.txn.Commit()
// }

// // Alias to commit
// func (t *txn) Close() error {
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return ErrClosed
// 	}
// 	return t.close()
// }

// func (t *txn) close() error {
// 	return t.txn.Commit()
// }

// func (t *txn) Discard(ctx context.Context) {
// 	t.ds.closeLk.RLock()
// 	defer t.ds.closeLk.RUnlock()
// 	if t.ds.closed {
// 		return
// 	}

// 	t.discard()
// }

// func (t *txn) discard() {
// 	t.txn.Discard()
// }

// // filter returns _true_ if we should filter (skip) the entry
// func filter(filters []dsq.Filter, entry dsq.Entry) bool {
// 	for _, f := range filters {
// 		if !f.Filter(entry) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func expires(item *badger.Item) time.Time {
// 	return time.Unix(int64(item.ExpiresAt()), 0)
// }
