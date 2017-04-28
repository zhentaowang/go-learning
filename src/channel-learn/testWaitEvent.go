package main

import "fmt"

func main() {
	fmt.Println("Begin doing something!")
	c := make(chan bool)
	go func() {
		fmt.Println("Doing something...")
		c<-true
	}()
	fmt.Println("Done!")
	<-c
	fmt.Println("Last")
}