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

func NewConnect(addr string) (*Connect, error)  {
	return &Connect{
		addr:    addr,
		cache:   defaultCaching,
		timeout: defaultTimeout,
	}, nil
}





func NewConnectWithOption(addr string, cache bool, timeout time.Duration) (*Connect, error)  {
    return &Connect{
		addr:    addr,
		cache:   cache,
		timeout: timeout,
	}, nil
}


func main() {

}
