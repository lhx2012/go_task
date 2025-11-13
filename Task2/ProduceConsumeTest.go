package Task2

import (
	"fmt"
	"sync"
	"time"
)

func ProduceConsumeRun() {
	//创建带缓冲的通道
	message := make(chan Message, 10)

	//使用sync.WaitGroup等待所有goroutine完成
	var wg sync.WaitGroup

	//启动5个生产中
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go producer(i, message, &wg)
	}

	//等待所有生产者完成
	wg.Wait()

	//关闭通道，通知消费者不在有新的数据
	close(message)

	//var wgConsumer sync.WaitGroup
	//启动10个消费者
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go consumer(i, message, &wg)
	}

	wg.Wait()

	////等待所有生产者完成
	//wg.Wait()
	//
	////关闭通道，通知消费者不在有新的数据
	//close(message)

	time.Sleep(1 * time.Second)
	fmt.Println("所有任务都处理完成")
}

type Message struct {
	ID      int
	Content string
	Time    time.Time
}

func producer(id int, ch chan<- Message, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := Message{
		ID:      id,
		Content: "驰名商品",
		Time:    time.Now(),
	}

	//发送数据到通道
	ch <- msg
	fmt.Printf("生产者%d生产了：%s\n", msg.ID, msg.Content)
	//模拟生产产品
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("生产者%d已经完成生产\n", id)
}

func consumer(id int, ch <-chan Message, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range ch {
		fmt.Printf("消费者%d消费了%d：%s(时间：%s)\n", id, msg.ID, msg.Content, msg.Time.Format("15:04:05"))
		//模拟处理时间
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Printf("消费者%d已经完成消费\n", id)
}
