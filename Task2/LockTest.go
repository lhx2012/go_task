package Task2

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (s *SafeCounter) Increment() {
	defer s.mu.Unlock()
	s.mu.Lock()
	s.count++
}
func (s *SafeCounter) GetCount() int {
	defer s.mu.Unlock()
	s.mu.Lock()
	return s.count
}

func LockTestRun() {
	counter := SafeCounter{
		mu:    sync.Mutex{},
		count: 0,
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()

	//time.Sleep(time.Second)
	fmt.Printf("自增完最后值：%d\n", counter.GetCount())
}
