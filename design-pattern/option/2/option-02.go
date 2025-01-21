package main

import "time"

const (
	defaultTimeout = 10
	defaultCaching = false
)

type Connect struct {
	addr string
	cache bool
	timeout time.Duration
    
}

type ConnectOptions struct {
	cache bool
	timeout time.Duration
}

func NewDefaultOptions()  *ConnectOptions {
	return &ConnectOptions{
		cache:   false,
		timeout: 0,
	}
	
}

func NewConnect(addr string, opts *ConnectOptions) (*Connect, error)  {
	return &Connect{
		addr:    addr,
		cache:   opts.cache,
		timeout: opts.timeout,
	}, nil
}








func main() {

}
