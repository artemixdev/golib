package store

import "testing"

func TestStore(t *testing.T) {
	store := NewStore()
	store.Set("test", []byte("123"))

	tx := store.Begin()

	if err := tx.Set("hello", []byte("world")); err != nil {
		t.Fatalf("tx set 'hello' returned error: %s", err)
	}

	if err := tx.Set("abc", []byte("def")); err != nil {
		t.Fatalf("tx set 'abc' returned error: %s", err)
	}

	value, found, err := tx.Get("test")
	if err != nil {
		t.Fatalf("tx get 'test' returned error: %s", err)
	}
	if !found {
		t.Fatalf("tx get 'test' returned not found")
	}
	if data := string(value); data != "123" {
		t.Fatalf("tx get 'test' returned wrong data: %s", err)
	}

	if err := tx.Commit(); err != nil {
		t.Fatalf("tx commit returned error: %s", err)
	}

	if _, _, err := tx.Get("key"); err == nil {
		t.Fatalf("tx get 'key' after commit not returned error")
	}

	value, found = store.Get("hello")
	if !found {
		t.Fatalf("store get 'hello' returned not found")
	}
	if data := string(value); data != "world" {
		t.Fatalf("store get 'hello' returned wrong data: %s", data)
	}

	value, found = store.Get("abc")
	if !found {
		t.Fatalf("store get 'abc' returned not found")
	}
	if data := string(value); data != "def" {
		t.Fatalf("store get 'abc' returned wrong data: %s", data)
	}
}
