package main

import (
	"fmt"
	"io"
	"strings"

	_ "github.com/axgle/mahonia"
)

type RwHandel struct {
	ReadChanle  chan []byte
	WriteChanle chan []byte
}

func (rwhandel RwHandel) OnRead(R io.Reader) {
	for {
		Body := make([]byte, 5012)
		n, _ := R.Read(Body)
		fmt.Println("读取到数据")
		//	dec := mahonia.NewDecoder("GBK")
		r := strings.ReplaceAll(string(Body)[:n], "http://www.240ps.com/jc/xinshouzixue.asp", "https://www.baidu.com")

		fmt.Println(r)
		rwhandel.ReadChanle <- []byte(r)
	}
}
func (RwHandel RwHandel) OnWrite(W io.Writer) {
	for {
		select {
		case b := <-RwHandel.WriteChanle:
			fmt.Printf(string(b))
			W.Write(b)
		default:

		}

	}
}
