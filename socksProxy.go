package main

import (
	"fmt"
	"net"
)

type UdpServer struct{}

func (u UdpServer) RunUpdProxy() {
	TcpServer, err1 := net.Listen("tcp", ":9092")

	if err1 != nil {
		fmt.Println("udp 失败")
	}

	func() {

		for {
			cli, _ := TcpServer.Accept()

			fmt.Println(cli.RemoteAddr().String())

			read := make([]byte, 1024)
			cli.Read(read)
			fmt.Println(string(read))
		}
	}()
}
