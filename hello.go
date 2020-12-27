package main

import "fmt"

func go1(num int, c chan int){
	c <- num
}

func main() {
	var ch = make(chan int)
	go go1(10, ch)
	go go1(5, ch)
	var a = <-ch
	var b = <-ch
	fmt.Print("a + b = ", a+b)
}
