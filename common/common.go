package common

import (
	"fmt"
	"sync/atomic"
)

func AddName(name string) {

	fmt.Println(name, " is adding...")

}

func core() {

	//runtime.GOMAXPROCS(8)

	var add int32 = 1
	add1 := atomic.AddInt32(&add, 10)

	fmt.Println(add1)
}
