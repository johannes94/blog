package main

import (
	"fmt"
	"time"
)

func sendWithDelay(nums []int, delay time.Duration) chan int {
	c := make(chan int)

	go func() {
		for _, n := range nums {
			time.Sleep(delay)
			c <- n
		}

		close(c)
	}()

	return c
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	chanA := sendWithDelay(nums, time.Duration(500)*time.Millisecond)
	chanB := sendWithDelay(nums, time.Duration(1000)*time.Millisecond)

	for chanA != nil || chanB != nil {
		select {
		case v, ok := <-chanA:
			if !ok {
				chanA = nil
				continue
			}
			fmt.Println("From a: ", v)
		case v, ok := <-chanB:
			if !ok {
				chanB = nil
				continue
			}
			fmt.Println("From b: ", v)
		}
	}
}
