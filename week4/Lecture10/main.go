package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	times := 10
	cp := &ConcurrentPrinter{}
	cp.PrintFoo(times)
	cp.PrintBar(times)
	cp.Wait()

}

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
}

func (cp *ConcurrentPrinter) PrintFoo(times int) {
	cp.Add(1)
	go func() {
		for i := 0; i < times; i++ {
			if i%2 == 0 {
				cp.Lock()
				fmt.Println("Foo")
				cp.Unlock()

				time.Sleep(time.Millisecond * 10)

			} else {

				time.Sleep(time.Millisecond * 10)
				continue
			}

		}
		cp.Done()
	}()

}

func (cp *ConcurrentPrinter) PrintBar(times int) {
	cp.Add(1)
	go func() {
		for i := 0; i < times; i++ {
			if i%2 == 0 {
				cp.Lock()
				fmt.Println("Bar")
				cp.Unlock()

				time.Sleep(time.Millisecond * 10)

			} else {

				time.Sleep(time.Millisecond * 10)
				continue
			}

		}
	}()
	cp.Done()
}
