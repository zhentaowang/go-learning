package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 15
	c <- 34
	c <- 65
	close(c)
	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)
	
	c <- 1
}