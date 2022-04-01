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

		res, _ := HttpUntity.ReadHttp(conn)

		url, err := HttpUntity.AnalyHttpReqUrl([]byte(res.Body))
		if err != nil {
			fmt.Println("解析出错")
			return
		}
		addr, e := untity.PaserIP(url)
		if e != nil {
			//	fmt.Println("域名解析出错")
			break
		}
		cli, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Println("链接失败")
		}

		if res.IsConnection {
			xProxy := core.XProxyCore{Serverip: addr}
			cli.Write(res.Body)
			xProxy.SetnetCli(cli)

			xProxy.SetProxyCli(conn)

			xProxy.Runproxy()
		} else {

			cli.Write(res.Body)

		}

	}

}
