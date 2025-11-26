package Task2

//package ProduceConsumeTest2
//
//import (
//	"context"
//	"fmt"
//	"sync"
//	"time"
//)
//
//type Message struct {
//	ID      int
//	Content string
//	Time    time.Time
//}
//
//func ProduceConsumeRun() {
//	// 创建带缓冲的通道，容量为10
//	messageChan := make(chan Message, 10)
//	var wg sync.WaitGroup
//
//	// 创建带取消的上下文
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	// 启动5个生产者
//	for i := 1; i <= 5; i++ {
//		wg.Add(1)
//		go producer(ctx, i, messageChan, &wg)
//	}
//
//	// 启动10个消费者
//	for i := 1; i <= 10; i++ {
//		wg.Add(1)
//		go consumer(ctx, i, messageChan, &wg)
//	}
//
//	// 等待所有生产者完成
//	fmt.Println("等待生产者完成...")
//	wg.Wait()
//
//	// 关闭通道，通知消费者不再有新的数据
//	close(messageChan)
//
//	fmt.Println("所有任务都处理完成")
//}
//
//// producer 生产者函数
//func producer(ctx context.Context, id int, ch chan<- Message, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	for i := 1; i <= 3; i++ {
//		msg := Message{
//			ID:      i,
//			Content: fmt.Sprintf("消息%d来自生产者%d", i, id),
//			Time:    time.Now(),
//		}
//
//		select {
//		case <-ctx.Done():
//			fmt.Printf("生产者 %d 被取消\n", id)
//			return
//		case ch <- msg:
//			fmt.Printf("生产者 %d 生产了: %s\n", id, msg.Content)
//		}
//
//		time.Sleep(time.Millisecond * 100)
//	}
//
//	fmt.Printf("生产者 %d 已完成生产\n", id)
//}
//
//// consumer 消费者函数
//func consumer(ctx context.Context, id int, ch <-chan Message, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	for {
//		select {
//		case <-ctx.Done():
//			fmt.Printf("消费者 %d 被取消\n", id)
//			return
//		case _, ok := <-ch:
//			if !ok {
//				fmt.Printf("消费者 %d 检测到通道关闭\n", id)
//				return
//			}
//		default:
//			select {
//			case _, ok := <-ch:
//				if !ok {
//					fmt.Printf("消费者 %d 检测到通道关闭\n", id)
//					return
//				}
//			case <-time.After(time.Millisecond * 50):
//				// 短暂等待后继续检查
//				continue
//			}
//		}
//	}
//}
