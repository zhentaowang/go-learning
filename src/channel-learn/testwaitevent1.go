package main

import "fmt"

func worker(start chan bool, index int) {
	<-start
	fmt.Println("This is Worker:", index)
}

func main() {
	start := make(chan bool)
	for i := 1; i <= 100; i++ {
		go worker(start, i)
	}
	close(start)
	select {} //deadlock we expected
}