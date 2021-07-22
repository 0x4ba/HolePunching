package main

import (
	"fmt"
	"os"
)

// var (
// 	host = "116.63.143.23"
// 	port = "666"
// )

type host struct {
	ipaddr string
	port   string
}

func (h *host) ToString() string {
	return h.ipaddr + ":" + h.port
}

var (
	s         = host{ipaddr: "116.63.143.23", port: "666"}
	localhost = host{ipaddr: "", port: "666"}
)

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
}
