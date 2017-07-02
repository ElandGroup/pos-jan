package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_Chan1(t *testing.T) {
	channel := make(chan string)
	go func() {
		channel <- "hello"
		fmt.Println("2.hello")
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("1.hello")
	msg := <-channel
	fmt.Println("2.", msg)
}

func f1(c chan<- int) {
	c <- 0
	//<-c
}

var c chan string

func ready(w string, s int) {
	time.Sleep(time.Duration(s) * time.Second)
	fmt.Println(w, "has ready!")
	c <- w
}

func Test_Chan2(t *testing.T) {
	c = make(chan string)
	go ready("coffe", 2)
	go ready("tea", 1)
	fmt.Println("1.is readying")
	fmt.Println(<-c)
	fmt.Println(<-c)
	time.Sleep(5 * time.Second)
}
