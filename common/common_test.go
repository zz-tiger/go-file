package common

import (
	"fmt"
	"go-file/utils"
	"strconv"
	"testing"
	"time"
)

func TestHello(t *testing.T) {

	utils.AddName("ken")

}

func TestAddName(t *testing.T) {
	AddName("Tiger")
}

func TestCore(t *testing.T) {

	core()

	var age = -1
	for i := 0; i < 5; i++ {
		fmt.Println(age << i)
	}
}

func TestSlectKey(t *testing.T) {

	var intChan chan int = make(chan int, 10)
	var strChan chan string = make(chan string, 10)

	go addInt(intChan)
	go addStr(strChan)
	index := 0
	for {

		select {

		case v := <-strChan:
			fmt.Println("读取到的str值: ", v)
		case v := <-intChan:
			fmt.Println("读取到的int值: ", v)
		default:
			fmt.Println("Game is over")
			fmt.Println("cycle counts : ", index)
			return
		}
		index++
	}

	//reflect.Kind()
}

func addInt(c chan int) {
	time.Sleep(time.Second * 100)
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func addStr(str chan string) {
	time.Sleep(time.Second * 100)
	for i := 0; i < 10; i++ {
		str <- "str" + strconv.Itoa(i)
	}
}
