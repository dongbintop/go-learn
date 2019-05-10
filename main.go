package main

func main() {
	semaphore := NewSemaphore(5)
	semaphore.Signal()
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
