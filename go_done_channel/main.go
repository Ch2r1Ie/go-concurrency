package godonechannel

import (
	"fmt"
	"time"
)

func multiples(i int) (chan int, chan struct{}) {

	current := 0

	out := make(chan int)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case out <- current * i:
				current++
			case <-done:
				fmt.Println("done")
				return
			}
		}
	}()

	return out, done
}

func Init() {

	twosChan, done := multiples(2)
	for i := range twosChan {
		if i > 20 {
			break
		}
		fmt.Println(i)
	}
	close(done)

	time.Sleep(1 * time.Second)
}
