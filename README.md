# golib
The library contains a set of solutions I found useful.

## collection
Common functions and primitives popular when working with collections.

## lang
Missing utils in the language itself or in its core libraries.

## limiting
Primitives to control rate limit.

## messaging
This package is for working with messages going through channels.
Maybe it slightly reminds Rx* libraries.

## middleware
Middleware for most common web routers.

## store
In-memory key-value store with transactions.

Example:
```go
store := NewStore()
store.Set("key0", []byte("value0"))

tx := store.Begin()
err := tx.Set("key1", []byte("value1"))
err = tx.Set("key2", []byte("value2"))
err = tx.Commit()

data, found := store.Get("key1")
```
