# localmemstore

[![GoDoc](https://godoc.org/github.com/atdiar/localmemstore?status.svg)](https://godoc.org/github.com/atdiar/localmemstore)

A string-based in-memory K/V store for development purposes only.

This is an implementation of a single instance (non-distributed), in-memory, unencrypted Key/Value store.

This is not fit for use in production.

## User Interface

The datastructure is a wrapper around a map type, implementing the below:

``` go
// Store defines the interface that a session store should implement.
// It should be made safe for concurrent use by multiple goroutines.
//
// NOTE: SetExpiry sets a timeout for the validity of a session
// if t = 0, the session should expire immediately.
// if t < 0, the session does not expire.
type Store interface {
	Get(id, hkey string) (res []byte, err error)
	Put(id string, hkey string, content []byte) error
	Delete(id, hkey string) error
	SetExpiry(id string, t time.Duration) error
}
```



Key/Value pairs are specific to a given id, for instance a User id.
The persistence is set by id, meaning that Key/Value pairs belonging to the same id have the same expiration date.

## License
MIT
