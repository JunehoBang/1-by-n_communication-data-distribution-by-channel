package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, done chan bool) {

	for {
		<-done
		fmt.Println(id)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(20)))
	}
}

func main() {

	done := make(chan bool)

	for i := 0; i < 20; i++ {
		go worker(i, done)
	}
	time.Sleep(time.Millisecond)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		done <- true
	}

}
