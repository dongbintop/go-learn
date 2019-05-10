package db20190510

import "fmt"

/**
  阶乘计算
  channel 递归
*/
func FactCalc(in <-chan int, out chan<- int) {
	var subIn, subOut chan int

	for {
		n := <-in
		if n == 0 {
			out <- 1
		} else {
			if subIn == nil {
				subIn, subOut = make(chan int), make(chan int)
				go FactCalc(subIn, subOut)
			}
			subIn <- n - 1
			r := <-subOut
			out <- n * r
		}
	}
}

func MakeFactFunc() func(int) int {
	in, out := make(chan int), make(chan int)
	go FactCalc(in, out)
	return func(x int) int {
		in <- x
		return <-out
	}
}

func Test() {
	factFunc := MakeFactFunc()
	fmt.Println(factFunc(3))
}
