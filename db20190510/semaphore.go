package db20190510

import (
	"fmt"
	"sync"
)

type Semaphore struct {
	inc chan struct{}
	dec chan struct{}
}

func (sem *Semaphore) Wait() {
	sem.dec <- struct{}{}
}

func (sem *Semaphore) Signal() {
	sem.inc <- struct{}{}
}

func MakeSemaphore(value int) *Semaphore {
	sem := Semaphore{
		inc: make(chan struct{}),
		dec: make(chan struct{}),
	}

	go func(s int) {
		for {
			if s > 0 {
				select {
				case <-sem.inc:
					s = s + 1
				case <-sem.dec:
					s = s - 1
				}
			} else {
				<-sem.inc
				s = s + 1
			}
		}

	}(value)

	return &sem
}

func SemaphoreTest() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	semaphore := MakeSemaphore(5)
	for i := 0; i < 100; i++ {
		go func(value int) {
			semaphore.Wait()
			fmt.Println(value)
			semaphore.Signal()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
