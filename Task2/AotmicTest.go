package Task2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SafeCounter2 struct {
	count int32
}

func (s *SafeCounter2) AutoIncrement() {
	atomic.AddInt32(&(s.count), 1)
}

func (s *SafeCounter2) GetCount() int32 {
	return s.count
}

func AtomicRun() {
	counter := SafeCounter2{count: 0}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				counter.AutoIncrement()
			}
		}()
	}

	wg.Wait() // 等待所有协程完成

	// 使用原子加载安全地读取最终值
	fmt.Printf("count direct is %d\n", counter.count)
	fmt.Printf("最终计数器数值: %d\n", atomic.LoadInt32(&counter.count))

}
