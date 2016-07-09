package main

/**
如果chan 满了， 访问会block，所以要增加default
*/
import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int64, 3)

	go func() {
		for {
			time.Sleep(2e9) // sleep two second
			<-messages
			//	fmt.Printf("Receive:  %v at %v\n", t, time.Now().Unix())
		}
	}()

	for i := 0; i < 5; i++ {
		go func() {
			for {
				t := time.Now().Unix()
				select {
				case messages <- t:
					fmt.Printf("Send:  %v, at %v \n", t, time.Now().Unix())
					//default:
					//	fmt.Println("send failed")
				}
				time.Sleep(1e9) // sleep one second
			}
		}()
	}

	time.Sleep(60e9) // sleep 60 second
}
