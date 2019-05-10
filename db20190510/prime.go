package db20190510

import "fmt"

/**
  获取素数

  利用channel 懒加载

  make(chan int) 没有缓存
*/

func Counter(out chan<- int) {
	for i := 2; ; i++ {
		out <- i
	}
}

func PrimeFilter(prime int, in <-chan int, out chan<- int) {

	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}

}

func PrimeSieve(out chan<- int) {
	c := make(chan int)
	go Counter(c)
	for {
		prime := <-c
		out <- prime
		nc := make(chan int)
		go PrimeFilter(prime, c, nc)
		c = nc
	}
}

func PrimeTest() {
	primes := make(chan int)
	go PrimeSieve(primes)
	for i := 0; i < 5; i++ {
		fmt.Println(<-primes)
	}

}
