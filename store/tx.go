package store

import (
	"errors"
	"sync"
)

type Tx struct {
	store      *Store
	deletions  map[string]struct{}
	insertions map[string][]byte
	mutex      sync.RWMutex
}

func newTx(store *Store) *Tx {
	return &Tx{
		store:      store,
		deletions:  make(map[string]struct{}),
		insertions: make(map[string][]byte),
	}
}

func (tx *Tx) Get(key string) ([]byte, bool, error) {
	tx.mutex.RLock()
	defer tx.mutex.RUnlock()

	if tx.store == nil {
		return nil, false, ErrTxClosed
	}

	if _, found := tx.deletions[key]; found {
		return nil, false, nil
	}

	value, found := tx.insertions[key]
	if !found {
		value, found = tx.store.pairs[key]
	}
	if !found {
		return nil, false, nil
	}

	return copyBytes(value), true, nil
}

func (tx *Tx) Set(key string, value []byte) error {
	tx.mutex.Lock()
	defer tx.mutex.Unlock()

	if tx.store == nil {
		return ErrTxClosed
	}

	delete(tx.deletions, key)
	tx.insertions[key] = copyBytes(value)

	return nil
}

func (tx *Tx) Delete(key string) error {
	tx.mutex.Lock()
	defer tx.mutex.Unlock()

	if tx.store == nil {
		return ErrTxClosed
	}

	tx.deletions[key] = struct{}{}
	delete(tx.insertions, key)

	return nil
}

func (tx *Tx) Commit() error {
	tx.mutex.Lock()
	defer tx.mutex.Unlock()

	if tx.store == nil {
		return ErrTxClosed
	}

	for key := range tx.deletions {
		delete(tx.store.pairs, key)
	}

	for key, value := range tx.insertions {
		tx.store.pairs[key] = value
	}

	tx.store.mutex.Unlock()
	tx.store = nil
	tx.deletions = nil
	tx.insertions = nil

	return nil
}

func (tx *Tx) Rollback() error {
	tx.mutex.Lock()
	defer tx.mutex.Unlock()

	if tx.store == nil {
		return ErrTxClosed
	}

	tx.store.mutex.Unlock()
	tx.store = nil
	tx.deletions = nil
	tx.insertions = nil

	return nil
}

var ErrTxClosed = errors.New("tx closed")
