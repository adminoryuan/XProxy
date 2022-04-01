package Http

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type HttpUntity struct{}

//解析http协议q
//reqByte 请求的http流
func (t *HttpUntity) AnalyHttpReqUrl(reqByte []byte) (string, error) {
	HttpAnals := strings.Split(string(reqByte), "\r")

	firstHttpRow := strings.Split(HttpAnals[0], " ")

	if len(firstHttpRow) < 2 {
		return "", errors.New("协议解析出错")
	}

	return firstHttpRow[1], nil
}

//读取一个完整的Http报文
func (t *HttpUntity) ReadHttp(reder io.Reader) ([]byte, error) {

	Res := make([]byte, 128)

	httpHeaders := make([]byte, 256)
	n, _ := reder.Read(httpHeaders)
	Res = append(Res, httpHeaders[:n]...)
	httpRow := strings.Split(string(httpHeaders), "\r\n")
	httplenth := 0
	for _, row := range httpRow {
		if strings.Contains(row, "Content-Length") {

			kv := strings.Split(row, ": ")

			val, err := strconv.Atoi(kv[1])
			if err != nil {
				fmt.Println("出错了")
				return []byte(""), nil
			}
			httplenth = val
			break
		}
	}

	body := strings.Split(string(httpHeaders[:n]), "\r\n\r\n")[1]

	httplenth -= (len(body) / 8)

	for httplenth > 0 {
		n, _ := reder.Read(httpHeaders)

		Res = append(Res, httpHeaders[:n]...)
		httplenth -= 256
	}
	fmt.Println(string(Res))
	return Res, nil
}
