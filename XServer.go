package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"

	"github.com/XProxy/Http"
	"github.com/XProxy/untity"
)

type HttpXproxy struct{}

var HttpUntity Http.HttpUntity = Http.HttpUntity{}

var CachePool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 512)
	},
}

//启动一个http代理服务
func (h HttpXproxy) StartXproxy(addr string) {

	fmt.Println("Start HyProxy ")
	fmt.Printf("addr:%s\n", addr)
	S, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		cli, err := S.Accept()
		if err != nil {
			fmt.Println("链接出错")
			continue
		}
		go h.HandleReq(cli)
	}

}
func (h HttpXproxy) HandleReq(conn net.Conn) {

	res := CachePool.Get().([]byte)
	n, er := conn.Read(res[:])
	if er == io.EOF {

		return
	}
	host, err := HttpUntity.AnalyHttpReqUrl(res[:n])

	CachePool.Put(res)

	fmt.Println(host)
	if err != nil {
		fmt.Println("解析出错")
		return
	}

	addr, e := untity.PaserIP(host)

	if e != nil {
		fmt.Println("解析失败")
		fmt.Println(e.Error())
		return
	}

	cli, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("链接失败")
		return
	}
	if strings.Contains(addr, ":443") {
		fmt.Println("pk")
		conn.Write([]byte("HTTP/1.0 200 ok\r\n\r\n"))
	} else {
		cli.Write(res[:n])
	}
	fmt.Println(addr)

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
