package main

import (
	"fmt"
	"math/rand"
	"time"
	// "os"
	// "os/signal"
	// "syscall"
)

func worker(id int, done chan bool) {
	// sig := <-sigs
	// fmt.Println()
	// fmt.Println(sig)
	for {
		<-done
		fmt.Println(id)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(20)))
	}
}

func main() {

	// sigs := make(chan os.Signal, 1)

	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

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
