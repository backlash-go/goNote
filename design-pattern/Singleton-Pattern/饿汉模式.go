package singleton

type singletonT struct{}

var instance = &singletonT{}

func GetsingletonT() *singletonT {
	return instance
}
