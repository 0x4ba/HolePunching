package main

import (
	"bufio"
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
	conn.Close()
	udpconn, err := net.ListenUDP("udp", hostaddr)
	if err != nil {
		fmt.Println("listening error", err)
		os.Exit(-1)
	}
	buf := make([]byte, 128)
	var dstAddr string

	for {
		_, _, err = udpconn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("client recv address error", err)
		}
		dstAddr = recvMsg(udpconn)
		if dstAddr != "" {
			break
		}
	}

	dstUDPAddr, err := net.ResolveUDPAddr("udp", dstAddr)
	if err != nil {
		fmt.Println("resolve addr error", err, dstAddr)
		os.Exit(-1)
	}
	// conn, err = net.DialUDP("udp", nil, dstUDPAddr)
	// if err != nil {
	// 	fmt.Println("dialing error", err, dstAddr)
	// }
	go func(conn *net.UDPConn, dstUDPAddr *net.UDPAddr) {
		reader := bufio.NewReader(os.Stdin)
		for {
			buf, _, _ := reader.ReadLine()
			conn.WriteToUDP(buf, dstUDPAddr)
		}
	}(udpconn, dstUDPAddr)

	go func(conn *net.UDPConn) {
		for {
			recvMsg(udpconn)
		}
	}(udpconn)
}

func recvMsg(conn *net.UDPConn) string {
	buf := make([]byte, 512)
	_, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf[:])
	return string(buf)
}
