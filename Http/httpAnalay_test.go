package Http

import (
	"fmt"
	"net"
	"testing"

	"github.com/XProxy/untity"
)

func TestHttprasr(t *testing.T) {

	ht, _ := untity.PaserIP("https://www.baidu.com/")
	fmt.Println(ht)
	con, e := net.Dial("tcp", "14.215.177.38:80")
	if e != nil {
		fmt.Println("链接出错")
		return
	}

	ReqHttp := "GET https://sp3.baidu.com/-0U_dTmfKgQFm2e88IuM_a/ps_fp.htm?pid=ps&fp=undefined&im=undefined&wf=undefined&br=3&qid=ca587372000038a4&bi=ED970D370B4B0059A7C508C5F0C6CA6C:FG=1&psid=ED970D370B4B0059C2CEB86AEE4E8CB8 HTTP/1.1\r\n"
	ReqHttp += "Host: sp3.baidu.com\r\n"
	ReqHttp += "User-Agent: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:98.0) Gecko/20100101 Firefox/98.0\r\n"
	ReqHttp += "Connection: keep-alive"
	ReqHttp += "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"
	ReqHttp += "\r\n\r\n"
	ReqByte := []byte(ReqHttp)

	// Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2
	// 	Accept-Encoding: gzip, deflate
	// 	Upgrade-Insecure-Requests: 1
	// 	Cache-Control: max-age=0

	// 	If-Modified-Since: Fri, 01 Apr 2022 01:26:09 GMT
	// 	If-None-Match: "1ad9e-5db8da9c9ce40"

	con.Write(ReqByte)

	untity := HttpUntity{}

	untity.ReadHttp(con)

}
