package singleton

import "sync"

type singleton struct {

}



var ins * singleton

var once sync.Once

func (r singleton) name()  {

}

func GetInstance() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}











