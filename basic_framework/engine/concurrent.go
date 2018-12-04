package engine

import (
	"log"
)

// 并发引擎里有一个调度，一个worker数量
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

// 调度器里有提交方法,
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

//
func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParserResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	// 根据Worker数量创建worker
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	// 遍历种子数量, 然后
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out

		// 打印item
		for _, item := range result.Items {
			log.Printf("Got item :%v ", item)
		}

		// 把results里的requests送给scheduler
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)

		}
	}

}

func createWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}

	}()

}
