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

	length, _, err := udpconn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("client recv address error", err)
	}

	dstAddr := string(buf)[0:length]
	fmt.Println(dstAddr)
	dstUDPAddr, err := net.ResolveUDPAddr("udp", dstAddr)
	if err != nil {
		fmt.Println("resolve addr error", err, dstAddr)
		os.Exit(-1)
	}
	udpconn.Close()
	// conn, err = net.DialUDP("udp", nil, dstUDPAddr)
	// if err != nil {
	// 	fmt.Println("dialing error", err, dstAddr)
	// }
	wg.Add(2)
	var done = make(chan struct{})
	//reader := bufio.NewScanner(os.Stdin)
	go func(dstUDPAddr *net.UDPAddr, done chan struct{}) {

		udpconn, err := net.DialUDP("udp", nil, dstUDPAddr)
		if err != nil {
			fmt.Println("listening error", err)
			os.Exit(-1)
		}

		conn.WriteToUDP([]byte("hello"), dstUDPAddr)
		done <- struct{}{}
		consolescanner := bufio.NewScanner(os.Stdin)

		for {
			// by default, bufio.Scanner scans newline-separated lines
			for consolescanner.Scan() {
				input := consolescanner.Bytes()
				udpconn.WriteToUDP(input, dstUDPAddr)
				fmt.Println("ME:" + string(input))
			}

		}
	}(dstUDPAddr, done)

	go recvMsg(hostaddr, done)
}

func recvMsg(hostaddr *net.UDPAddr, done chan struct{}) string {
	udpconn, err := net.ListenUDP("udp", hostaddr)
	if err != nil {
		fmt.Println("listening error", err)
		os.Exit(-1)
	}
	<-done
	buf := make([]byte, 512)
	for {
		_, _, err := udpconn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(buf))
		buf = buf[0:0]
	}
}
