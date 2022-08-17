package store

import "sync"

type Store struct {
	pairs map[string][]byte
	mutex sync.RWMutex
}

func NewStore(store ...*Store) *Store {
	if len(store) > 0 {
		s := store[0]
		s.mutex.RLock()
		defer s.mutex.RUnlock()

		pairs := make(map[string][]byte, len(s.pairs))
		for key, value := range s.pairs {
			pairs[key] = copyBytes(value)
		}

		return &Store{pairs: pairs}
	}

	return &Store{pairs: make(map[string][]byte)}
}

func (store *Store) Get(key string) ([]byte, bool) {
	tx := store.Begin()
	value, found, _ := tx.Get(key)
	_ = tx.Commit()
	return value, found
}

func (store *Store) Set(key string, value []byte) {
	tx := store.Begin()
	_ = tx.Set(key, value)
	_ = tx.Commit()
}

func (store *Store) Delete(key string) {
	tx := store.Begin()
	_ = tx.Delete(key)
	_ = tx.Commit()
}

func (store *Store) Begin() *Tx {
	store.mutex.Lock()
	return newTx(store)
}
