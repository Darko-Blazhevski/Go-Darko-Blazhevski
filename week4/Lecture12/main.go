package main

import (
	"log"
	"time"
)

func main() {
	out := generateThrottled("foo", 2, time.Second)
	for f := range out {
		log.Println(f)
	}

}

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	ch := make(chan string, bufferLimit)
	go func() {
		for {
			for i := 0; i < bufferLimit; i++ {
				ch <- data
			}

			time.Sleep(clearInterval)
		}
	}()

	return ch

}
