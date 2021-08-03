package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//测试os中的参数，包括command传入的参数 和 os包自带的常用项

func worker(stop <-chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("exit")
			return
		default:
			fmt.Println("running...")
			time.Sleep(3)
		}

	}
}

func waitForSignal() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	fmt.Println(<-sigs)
}

func main() {
	stop := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(stop <-chan bool) {
			defer wg.Done()
			worker(stop)
		}(stop)
	}
	waitForSignal()
	close(stop)
	fmt.Println("stopping all job")
	wg.Wait()
}
