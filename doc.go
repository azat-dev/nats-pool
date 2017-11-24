// Package implements a connection pool for nats.io which is
// thread-safe.
//
// Basic usage
//
// The basic use-case is to create a pool and then pass that pool amongst
// multiple go-routines, each of which can use it safely. To retrieve a
// connection you use Get, and to return the connection to the pool when you're
// done with it you use Put.
//
//	p, err := pool.New("nats://localhost:4222", 10)
//	if err != nil {
//		// handle error
//	}
//
//	// In another go-routine
//
//	conn, err := p.Get()
//	if err != nil {
//		// handle error
//	}
//
//	conn.Publish(....)
//
//	p.Put(conn)
//
// Shortcuts
//
// If you're doing multiple operations you may find it useful to defer the Put
// right after retrieving a connection, so that you don't have to always
// remember to do so
//
//	conn, err := p.Get()
//	if err != nil {
//		// handle error
//	}
//	defer p.Put(conn)
//
//	conn.Publish(....)
//
//
//
// Custom connections
//
// Sometimes it's necessary to run some code on each connection in a pool upon
// its creation, This can be done with NewCustom, like so
//
//	df := func(addr string) (*nats.Conn, error) {
//		client, err := nats.Connect(addr)
//		if err != nil {
//			return nil, err
//		}
//		client.Publish(....)
//
//		return client, nil
//	}
//	p, err := pool.NewCustom("nats://localhost:4222", 10, df)
package nats_pool
