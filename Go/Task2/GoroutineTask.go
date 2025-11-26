package Task2

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	id        string
	timeStamp int64
	handle    func() error
}

func (task *Task) ID() string {
	return task.id
}

func (task *Task) Execute() error {
	if task.handle != nil {
		return task.handle()
	}
	return errors.New("handle is nil")
}

type Result struct {
	ID        string
	Error     error
	TimeStamp time.Duration
}

type Scheduler struct {
	tasks   []Task
	results chan Result
	wg      sync.WaitGroup
}

func NewTaskScheduler() *Scheduler {
	return &Scheduler{
		tasks:   make([]Task, 0),
		results: make(chan Result, 10),
	}
}

func (ts *Scheduler) AddTask(id string, handle func() error) {
	ts.tasks = append(ts.tasks, Task{
		id:     id,
		handle: handle,
	})
}

func (ts *Scheduler) Execute() []Result {
	for _, task := range ts.tasks {
		ts.wg.Add(1)
		go ts.RunTask(task)
	}

	go func() {
		ts.wg.Wait()
		close(ts.results)
	}()

	var results []Result
	for rs := range ts.results {
		results = append(results, rs)
	}

	return results
}

func (ts *Scheduler) RunTask(task Task) {
	defer ts.wg.Done()

	start := time.Now()
	err := task.handle()
	duration := time.Since(start)

	//发送结果到通道
	ts.results <- Result{
		ID:        task.id,
		Error:     err,
		TimeStamp: duration,
	}
}

func goroutineRun() {
	printNum()
	time.Sleep(time.Second)

	//创建调度器
	taskScheduler := NewTaskScheduler()
	//添加任务
	taskScheduler.AddTask("First Step", FirstStep)
	taskScheduler.AddTask("Second Step", SecondStep)
	taskScheduler.AddTask("Third Step", ThirdStep)

	//开始执行任务并获得结果
	fmt.Printf("开始执行任务......\n")
	results := taskScheduler.Execute()

	// 打印执行结果
	fmt.Printf("显示任务执行结果:\n")
	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("Task :%s,Duration:%v,Error:%v\n", result.ID, result.TimeStamp, result.Error)
		} else {
			fmt.Printf("Task :%s,Duration:%v,Success\n", result.ID, result.TimeStamp)
		}

	}
}

func FirstStep() error {

	time.Sleep(2 * time.Second)
	fmt.Println("First Step completed")
	return nil
}

func SecondStep() error {
	time.Sleep(4 * time.Second)
	fmt.Println("Second Step completed")
	return nil
}

func ThirdStep() error {
	time.Sleep(3 * time.Second)
	fmt.Println("Third Step completed")
	return nil
}

func printNum() {
	go func() {
		for i := 0; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Printf("打印出偶数：%d\n", i)
			} else {
				continue
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				fmt.Printf("打印出奇数：%d\n", i)
			} else {
				continue
			}
		}
	}()
}
