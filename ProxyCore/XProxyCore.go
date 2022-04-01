package ProxyCore

import (
	"fmt"
	"io"
	"net"
)

type XProxyCore struct {
	proxyCli net.Conn //被代理端

	closeChanle chan struct{}

	Serverip string //服务器ip地址

	netCli net.Conn //代理端与目标端建立的请求

}

func (x *XProxyCore) SetProxyCli(con net.Conn) {
	x.proxyCli = con

	fmt.Println(x.proxyCli.RemoteAddr())
}
func (x *XProxyCore) SetnetCli(con net.Conn) {
	x.netCli = con
	fmt.Println(x.netCli.RemoteAddr())
}

//监听被代理端的请求
func (x XProxyCore) proxyRead() {
	for {
		body := make([]byte, 1024)
		n, err := x.proxyCli.Read(body)
		if err == io.EOF {
			fmt.Println("被代理端链接被关闭")
			x.proxyCli.Close()
			x.closeChanle <- struct{}{}
			break
		}
		x.netCli.Write(body[:n])

	}
}

func (x XProxyCore) Runproxy() {
	fmt.Println("开启了一个代理")

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
			bodys := make([]byte, 5012)
			n, err := x.netCli.Read(bodys)
			if err == io.EOF {
				fmt.Println("服务器链接被关闭")
				x.netCli.Close()
				return
			}
			fmt.Println(string(bodys[:n]))
			x.proxyCli.Write(bodys[:n])

		}
	}
}
