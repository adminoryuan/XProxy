# XProxy 
 - # 介绍
  - 使用golang 实现的一个http https 代理，经过测试稳定运行
 - # 实现原理
 - 对于http 请求 解析host链接80 端口
 - 对于https 请求解析 host 链接443 端口
 - 将客户端 request 转发给目标host
 - 将目标host的response转发给客户端
 - # 如何使用
  - 下载源码到本地
  - go build .
  - ./XProxy &
  - 配置系统代理
- # 核心代码
 ```golang
 package ProxyCore

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type XProxyCore struct {
	proxyCli net.Conn //被代理端

	closeChanle chan struct{}

	pool sync.Pool //复用池

	Serverip string //服务器ip地址

	netCli net.Conn //代理端与目标端建立的请求

}

func (x *XProxyCore) SetProxyCli(con net.Conn) {
	x.proxyCli = con

	//fmt.Println(x.proxyCli.RemoteAddr())
}
func (x *XProxyCore) SetnetCli(con net.Conn) {
	x.netCli = con
	//	fmt.Println(x.netCli.RemoteAddr())
}

//监听被代理端的请求
func (x XProxyCore) proxyRead() {
	for {

		bodys := x.pool.Get().([]byte)
		n, err := x.proxyCli.Read(bodys)
		if err == io.EOF {
			//fmt.Println("被代理端链接被关闭")
			x.proxyCli.Close()
			x.closeChanle <- struct{}{}
			x.pool.Put(bodys)
			break
		}
		x.netCli.Write(bodys[:n])
		x.pool.Put(bodys)
	}
}

func (x XProxyCore) Runproxy() {
	fmt.Println("开启了一个代理")
	x.pool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}
	go x.proxyRead()
	go x.cliRead()
}

//监听目标服务器的响应
func (x XProxyCore) cliRead() {
	for {
		select {
		case <-x.closeChanle:
			x.netCli.Close()
			fmt.Println("关闭与服务器链接")
			return
		default:
			bodys := x.pool.Get().([]byte)
			n, err := x.netCli.Read(bodys)
			if err == io.EOF {
				fmt.Println("服务器链接被关闭")
				x.netCli.Close()
				x.pool.Put(bodys)
				return
			}
			//fmt.Println(string(bodys[:n]))
			x.proxyCli.Write(bodys[:n])
			x.pool.Put(bodys)
		}
	}
}

 ```
  
