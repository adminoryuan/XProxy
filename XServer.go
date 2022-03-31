package main

import (
	"fmt"
	"net"

	"github.com/XProxy/Http"
	"github.com/XProxy/untity"
)

type HttpXproxy struct{}

var HttpUntity Http.HttpUntity = Http.HttpUntity{}

//启动一个http代理服务
func (h HttpXproxy) StartXproxy(addr string) {

	S, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		cli, err := S.Accept()
		if err != nil {
			fmt.Println("链接出错")
		}
		go h.HandleReq(cli)
	}

}
func (h HttpXproxy) HandleReq(conn net.Conn) {
	for {

		bodys := make([]byte, 5012)

		n, _ := conn.Read(bodys)

		req := HttpUntity.AnalyHttp(bodys[:n])

		R := RwHandel{}

		addr, e := untity.PaserIP(req.Url)
		if e != nil {

			break
		}
		fmt.Println(addr)
		cli, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Println("链接失败")
		}
		R.ReadChanle = make(chan []byte)
		R.WriteChanle = make(chan []byte)
		go R.OnRead(cli)

		go R.OnWrite(cli)

		R.WriteChanle <- bodys[:n]

		conn.Write(<-R.ReadChanle)

	}

}
