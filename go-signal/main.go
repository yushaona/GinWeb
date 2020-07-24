package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var sig os.Signal

var wg sync.WaitGroup
var c chan os.Signal

func main() {
	c = make(chan os.Signal)
	hookableSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
	}
	wg = sync.WaitGroup{}
	
	signal.Notify(c, hookableSignals...)
	wg.Add(1)
	go listenSignal()
	wg.Wait()
}

func listenSignal() {

	Pid := syscall.Getpid()
	for {
		sig := <-c

		switch sig {
		case syscall.SIGHUP:
			log.Println(Pid, "Recevied SIGHUP. forking...")
		case syscall.SIGINT:
			log.Println(Pid, "Recevied SIGINT. ")
		case syscall.SIGTERM:
			log.Println(Pid, "Recevied SIGTERM. ")
		default:
			log.Println("Received %v: nothing i care about...\n", sig)
		}
	}
	wg.Done()
}
