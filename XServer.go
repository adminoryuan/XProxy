package main

import (
	"fmt"
	"net"

	"github.com/XProxy/Http"
	core "github.com/XProxy/ProxyCore"
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

		url, err := HttpUntity.AnalyHttpReqUrl(bodys[:n])
		if err != nil {
			fmt.Println("解析出错")
			return
		}
		addr, e := untity.PaserIP(url)
		if e != nil {
			fmt.Println("协议解析出错")
			break
		}
		fmt.Println(addr)
		cli, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Println("链接失败")
		}
		xProxy := core.XProxyCore{}

		cli.Write(bodys[:n])

		xProxy.SetnetCli(cli)

		xProxy.SetProxyCli(conn)

		xProxy.Runproxy()

	}

}
