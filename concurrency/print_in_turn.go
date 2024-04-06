package concurrency

import (
	"fmt"
	"time"
)

// print n mission's number in turn
func _print_in_turn(n int) {
	chs := make([]chan struct{}, n)
	for i := range chs {
		chs[i] = make(chan struct{})
	}
	for i := range n {
		go func() {
			i := i
			for {
				select {
				case <-chs[(i+n-1)%n]:
					fmt.Println(i)
					time.Sleep(1 * time.Second)
					chs[i] <- struct{}{}
				}
			}
		}()
	}

	chs[n-1] <- struct{}{}
	select {}
}
