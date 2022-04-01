package Http

import (
	"errors"
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
