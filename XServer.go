package main

import (
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/XProxy/Http"
	"github.com/XProxy/untity"
	workpool "github.com/XProxy/workPool"
)

type HttpXproxy struct{}

var Pool workpool.Pool = *workpool.NewPool(20)
var HttpUntity Http.HttpUntity = Http.HttpUntity{}

//启动一个http代理服务
func (h HttpXproxy) StartXproxy(addr string) {
	go Pool.Run()
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

	res := make([]byte, 1024)
	n, _ := conn.Read(res[:])

	host, err := HttpUntity.AnalyHttpReqUrl(res[:n])
	fmt.Println(host)
	if err != nil {
		fmt.Println("解析出错")
		return
	}

	addr, e := untity.PaserIP(host)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	cli, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("链接失败")
		return
	}
	if strings.Contains(addr, ":443") {
		conn.Write([]byte("HTTP/1.0 200 ok\r\n\r\n"))
	} else {
		conn.Write(res)
	}
	fmt.Println(addr)

	// t1 := workpool.NewTask(func() error {
	// 	io.Copy(cli, conn)
	// 	return nil
	// })
	// t2 := workpool.NewTask(func() error {
	// 	io.Copy(conn, cli)

	// 	return nil
	// })
	go io.Copy(cli, conn)
	go io.Copy(conn, cli)
	// Pool.EntryChannel <- t1
	// Pool.EntryChannel <- t2
	//cli.Write(res)

	// if res.IsConnection {
	// 	xProxy := core.XProxyCore{Serverip: addr}
	// 	cli.Write(res.Body)
	// 	xProxy.SetnetCli(cli)

	// 	xProxy.SetProxyCli(conn)

	// 	xProxy.Runproxy()
	// } else {

	// 	cli.Write(res.Body)

	// }

}
