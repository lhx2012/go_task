package Task2

import (
	"fmt"
	"time"
)

func channelRun() {
	ch := make(chan int, 5)
	go sendNum(ch)
	go receiveNum(ch)

	timeout := time.After(2 * time.Second)

	for true {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Printf("主goroutine接受到：%d\n", v)
		case <-timeout:
			fmt.Println("timeout")
			return
		default:
			fmt.Println("没有数据，等待中......")
			time.Sleep(1 * time.Second)
		}
	}
}

func sendNum(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func receiveNum(ch <-chan int) {

	for val := range ch {
		fmt.Printf("接收到：%d\n", val)
	}
}
