package semaphore

/**
通过一个channel实现 semaphore
有一个问题是 先调用Signal 会导致阻塞

这个可以强允许只有 initVal go程在跑
*/

type Semaphore chan struct{}

func (s Semaphore) Wait() {
	s <- struct{}{}
}
func (s Semaphore) Signal() {
	<-s
}
func NewSemaphore(initVal uint) Semaphore {
	return make(Semaphore, initVal)
}
