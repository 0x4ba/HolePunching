package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

var serverWg = sync.WaitGroup{}

func server() {
	fmt.Println("server")
	addr, err := net.ResolveUDPAddr("udp", localhost.ToString())
	if err != nil {
		fmt.Println("resolve addr error", err)
		os.Exit(-1)
	}

	//var clients = map[string]string
	// var client1 chan string
	// var client2 chan string
	client1 := make(chan string, 1)
	client2 := make(chan string, 1)

	udpconn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("listening error", err)
		os.Exit(-1)
	}

	for {
		serverWg.Add(2)
		buf := make([]byte, 64)
		_, remoteaddr1, err := udpconn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("read error", err)
		}
		client2 <- remoteaddr1.String()
		fmt.Println(string(buf))
		fmt.Println(remoteaddr1)
		go SendRealAddr(client1, udpconn, remoteaddr1.String())

		_, remoteaddr2, err := udpconn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("read error", err)
		}
		client1 <- remoteaddr2.String()
		fmt.Println(string(buf))
		fmt.Println(remoteaddr2)

		go SendRealAddr(client2, udpconn, remoteaddr2.String())
		serverWg.Wait()
	}

}

func SendRealAddr(client <-chan string, conn *net.UDPConn, Raddr string) {
	//defer conn.Close()
	cli := <-client

	//fmt.Println(cli1, cli2)
	RUDPAddr, err := net.ResolveUDPAddr("udp", Raddr)
	if err != nil {
		fmt.Println("resolve addr error", err, Raddr)
		os.Exit(-1)
	}
	conn.WriteToUDP([]byte(cli), RUDPAddr)
	serverWg.Done()
}
