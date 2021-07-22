package main

import (
	"fmt"
	"net"
	"os"
)

func client() {
	fmt.Println("client")
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

	conn.Write([]byte("test"))

	hostaddr, err := net.ResolveUDPAddr("udp", conn.LocalAddr().String())
	if err != nil {
		fmt.Println("resolve addr error", err)
		os.Exit(-1)
	}

	udpconn, err := net.ListenUDP("udp", hostaddr)
	if err != nil {
		fmt.Println("listening error", err)
		os.Exit(-1)
	}
	buf := make([]byte, 128)

	for {
		_, _, err = udpconn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("client recv address error", err)
		}
		recvMsg(udpconn)
	}

}

func recvMsg(conn *net.UDPConn) {
	buf := make([]byte, 512)
	for {
		_, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(buf[:])
		buf = buf[0:0]
	}
}
