package main

import (
	"fmt"
	"time"
)

const (
	defaultCaching = false
	defaultTimeout = 10
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

type options struct {
	cache   bool
	timeout time.Duration
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
		fmt.Println("executed")
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.cache = cache
	})

}

func NewConnect(addr string, opts ...Option) (*Connection, error) {
	options := options{
		cache:   false,
		timeout: 0,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &Connection{
		addr:    addr,
		cache:   options.cache,
		timeout: options.timeout,
	}, nil
}


//define a function type



type testFunc func(a string)

func (f testFunc) name(n string)  {
	f(n)
}

func pName(name string)  {
   fmt.Println(name)

}


func main() {
	c, _ := NewConnect("127", WithTimeout(200))
	fmt.Println(c.timeout)

	var t1 testFunc = pName
	t1.name("buddy is bad")





}
