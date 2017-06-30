package main

import "fmt"

const n = 10

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}
}
