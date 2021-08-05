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

	wg.Add(2)

	udpconn, err = net.ListenUDP("udp", hostaddr)
	if err != nil {
		fmt.Println("listening error", err)
		os.Exit(-1)
	}

	go func(conn *net.UDPConn, dstUDPAddr *net.UDPAddr) {

		i, err := conn.WriteToUDP([]byte("hello"), dstUDPAddr)
		fmt.Println(i, err)
		consolescanner := bufio.NewScanner(os.Stdin)

		for {
			// by default, bufio.Scanner scans newline-separated lines
			for consolescanner.Scan() {
				input := consolescanner.Bytes()
				if string(input) != "" {
					i, err := conn.WriteToUDP(input, dstUDPAddr)

					fmt.Println("ME:"+string(input[0:i]), "length:"+fmt.Sprint(len(input)), err)
				}
			}

		}
	}(udpconn, dstUDPAddr)

	go recvMsg(udpconn, hostaddr)
}

func recvMsg(conn *net.UDPConn, hostaddr *net.UDPAddr) string {

	buf := make([]byte, 512)
	for {
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		if string(buf[0:len]) != "" {
			fmt.Println(string(buf[0:len]))
		}

	}
}
