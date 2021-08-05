package main

import (
	"fmt"
	"os"
	"sync"
)

type host struct {
	ipaddr string
	port   string
}

func (h *host) ToString() string {
	return h.ipaddr + ":" + h.port
}

var (
	s         = host{ipaddr: "xxx.xxx.xxx.xxx", port: "666"}
	localhost = host{ipaddr: "", port: "666"}
)

var wg = sync.WaitGroup{}

func main() {
	args := os.Args[1:]

	for _, v := range args {
		switch v {
		case "-s":
			server()
		case "-c":
			client()
		}
	}
	fmt.Println("done")
	wg.Wait()
}
