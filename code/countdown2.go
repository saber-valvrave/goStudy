package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println("commencing countdown")
	abort := make(chan struct{})
	tick := time.Tick(1 * time.Second)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	countdown := 10

	for {
		select {
		case <-tick:
			fmt.Println(countdown)
			countdown--
		case <-abort:
			fmt.Println("abort")
			return
		}
		if countdown == 0 {
			break
		}
	}

	fmt.Println("over")
}
