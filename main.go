package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	pid := os.Getpid()
	s := strconv.Itoa(pid)
	fmt.Printf("pid: %d\n", os.Getpid())
	d1 := []byte("#!/bin/sh\nkill -SIGUSR1 " + s)
	err := ioutil.WriteFile("stopping.sh", d1, 0644)
	if err != nil {
		fmt.Printf("err :  %v s", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGUSR1)
	stopping := false
	// Block until a signal is received.
	go func() {
		s := <-c
		fmt.Println("Got signal:", s)
		stopping = true
		//os.Exit(0)
	}()
	t := time.Tick(1 * time.Minute)
	for now := range t {
		fmt.Printf("%v %s\n", now, "test")
		if stopping {
			return
		}
	}
}
