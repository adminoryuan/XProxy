package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"

	"github.com/XProxy/Http"
	"github.com/XProxy/ProxyCore"
	"github.com/XProxy/untity"
)

type HttpXproxy struct {
	cachePool sync.Pool
}

var HttpUntity Http.HttpUntity = Http.HttpUntity{}

//启动一个http代理服务
func (h *HttpXproxy) StartXproxy(addr string) {

	h.cachePool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 512)
		},
	}
	S, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Start HyProxy ")
	fmt.Printf("addr:%s\n", addr)

	for {
		cli, err := S.Accept()

		if err != nil {
			fmt.Println("链接出错")
			continue
		}
		go h.HandleReq(cli)
	}

}
func (h *HttpXproxy) HandleReq(conn net.Conn) {
	fmt.Printf("用户主机%s使用了代理", conn.RemoteAddr().String())

	res := h.cachePool.Get().([]byte)
	n, er := conn.Read(res[:])
	if er == io.EOF {
		h.cachePool.Put(res)
		return
	}
	host, err := HttpUntity.AnalyHttpReqUrl(res[:n])

	h.cachePool.Put(res)

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

		conn.Write([]byte("HTTP/1.0 200 ok\r\n\r\n"))
	} else {
		cli.Write(res[:n])
	}

	// 可以使用这种方式
	// go io.Copy(cli, conn)
	// go io.Copy(conn, cli)

	xProxy := ProxyCore.XProxyCore{Serverip: addr}

	xProxy.SetnetCli(cli)

	xProxy.SetProxyCli(conn)

	go xProxy.Runproxy()

}
