package Behavioral_Patterns

/*


import "time"

const (
	defaultTimeout = 10
	defaultCaching = false
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

func NewConnect(addr string) (*Connection, error) {
	return &Connection{
		addr:    addr,
		cache:   defaultCaching,
		timeout: defaultTimeout,
	}, nil
}

func NewConnectWithOptions(addr string, cache bool, timeout time.Duration) (*Connection, error) {
	return &Connection{addr: addr, cache: cache, timeout: timeout,}, nil
}

 */



/*

import "time"

const (
	defaultTimeout = 10
	defaultCaching = false
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

type ConnectionOptions struct {
	Caching bool
	Timeout time.Duration
}

func NewDefaultOptions() *ConnectionOptions {
	return &ConnectionOptions{Caching: defaultCaching, Timeout: defaultTimeout}
}

func NewConnect(addr string, opts *ConnectionOptions) (*Connection, error)  {
	return &Connection{
		addr:addr,
		cache: opts.Caching,
		timeout: opts.Timeout,
	}, nil
}
 */