package main

import (
	"fmt"
	"net"
	"os"
)

func client() {

	addr, err := net.ResolveUDPAddr("udp", s.ToString())
	if err != nil {
		fmt.Println("resolve addr error", err, s.ToString())
		os.Exit(-1)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("dialudp error", err, s.ToString())
		os.Exit(-1)
	}

	conn.Close()

}
