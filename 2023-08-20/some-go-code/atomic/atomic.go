package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var totalCounter uint64

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		fmt.Println("increase by ", i, " current counter ", totalCounter)
		atomic.AddUint64(&totalCounter, uint64(i))
	}
}

//ignore for another main
//func main() {
//	var wg sync.WaitGroup
//	totalCounter = 0
//	wg.Add(2)
//	go worker(&wg)
//	go worker(&wg)
//	wg.Wait()
//
//	fmt.Println(totalCounter)
//}

type OnceDemo struct {
	onceFlag uint32
	mu       sync.Mutex
}

func (o *OnceDemo) Do(fn func()) {
	if atomic.LoadUint32(&o.onceFlag) == 1 {
		return
	}

	o.mu.Lock()
	defer o.mu.Unlock()
	if o.onceFlag == 0 {
		defer atomic.StoreUint32(&o.onceFlag, 1)
		fn()
	}
}

func main() {
	once := &OnceDemo{}

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(ind int) {
			once.Do(func() {
				fmt.Println("hello world! ", ind)
			})
			wg.Done()
		}(i)
	}

	wg.Wait()

}
