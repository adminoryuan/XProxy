package Http

import (
	"strings"
)

type HttpUntity struct{}

//解析http协议q
//reqByte 请求的http流
func (t *HttpUntity) AnalyHttp(reqByte []byte) Requests {
	HttpAnals := strings.Split(string(reqByte), "\r")

	firstHttpRow := strings.Split(HttpAnals[0], " ")

	re := Requests{
		Url:    firstHttpRow[1],
		Method: firstHttpRow[0],
	}
	return re
}
