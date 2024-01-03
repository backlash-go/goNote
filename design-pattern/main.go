package main

import (
	"fmt"
	"time"
)

const (
	defaultTimeout = 10
	defaultCaching = false
)

type options struct {
	timeout time.Duration
	caching bool
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
		fmt.Println("aaaa")
		o.timeout = t
	})
}



func main() {
	options := options{timeout: defaultTimeout, caching: defaultCaching,}
	wt := WithTimeout(100)
	wt.apply(&options)
	fmt.Println(options)



}
