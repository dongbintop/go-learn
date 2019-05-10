package main

import "go-learn/db20190510/proxy"

func main() {
	proxy.Test()
}

type Semaphore chan struct{}

func (s Semaphore) Wait() {
	s <- struct{}{}
}
func (s Semaphore) Signal() {
	<-s
}
func NewSemaphore(value uint) Semaphore {
	return make(Semaphore, value)
}
