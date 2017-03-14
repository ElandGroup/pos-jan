package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		//1.flag.String(),Bool(),Int()
		httpAddr = flag.String("http.addr", ":5000", "http listen address")
		name     string
		//3.bind customer type
	)
	//2.bind flag into a varable
	flag.StringVar(&name, "xiao", "123", "help message for flagname")
	flag.Parse()
	fmt.Println(*httpAddr)
	fmt.Println(name)

	//go run ./flag/cmd/main.go -'http.addr' :5001

}
