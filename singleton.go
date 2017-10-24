package singleton

import (
	"sync"
	"fmt"
)

type Administrator struct{
	name string
}

var instance *Administrator
var once sync.Once

func GetInstance(n string) *Administrator{
	once.Do(func(){
		instance=&Administrator{n}
	})
	return instance
}

func(admin *Administrator) PrintName(){
	fmt.Println(admin.name)
}