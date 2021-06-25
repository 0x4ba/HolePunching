package main

import (
	"fmt"
	"net"
	"os"
)

func server() {
	fmt.Println("server")
	addr, err := net.ResolveUDPAddr("udp", localhost.ToString())
	if err != nil {
		fmt.Println("resolve addr error", err)
		os.Exit(-1)
	}

	//var clients = map[string]string
	clients := make(chan string, 2)

	for {

		udpconn, err := net.ListenUDP("udp", addr)
		if err != nil {
			fmt.Println("listening error", err)
			os.Exit(-1)
		}
		buf := make([]byte, 64)

		_, remoteaddr, err := udpconn.ReadFromUDP(buf)

		fmt.Println(remoteaddr)
		select {
		case clients <- remoteaddr.String():
			fmt.Println("get client address success")
		default:
			SendRealAddr(clients)
		}

		udpconn.Close()

	}

}

func SendRealAddr(clients <-chan string) {
	cli1 := <-clients
	cli2 := <-clients

	addr, err := net.ResolveUDPAddr("udp", cli1)
	if err != nil {
		fmt.Println("resolve addr error", err, cli1)
		os.Exit(-1)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("dialudp error", err, cli1)
		os.Exit(-1)
	}

	SendMsgHandler(conn, cli2)

	addr, err = net.ResolveUDPAddr("udp", cli2)
	if err != nil {
		fmt.Println("resolve addr error", err, cli2)
		os.Exit(-1)
	}

	conn, err = net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("dialudp error", err, cli2)
		os.Exit(-1)
	}

	SendMsgHandler(conn, cli1)
}

func SendMsgHandler(conn *net.UDPConn, msg string) bool {
	defer conn.Close()
	_, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("Send msg error", err)
		return false
	}

	return true
}
