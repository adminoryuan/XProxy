package untity

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

//解析地址
func PaserIP(DomainName string) (string, error) {
	fmt.Println(DomainName)
	port := 80
	//

	if strings.Contains(DomainName, ":443") {
		port = 443
		DomainName = strings.ReplaceAll(DomainName, ":443", "")
	} else {
		h := strings.Split(DomainName, "//")
		if len(h) < 2 {
			return "", errors.New("error")
		}
		h = strings.Split(h[1], "/")

		DomainName = h[0]

		if strings.EqualFold(DomainName, "") {

			return "", errors.New("地址错误")
		}
	}

	c, err := net.ResolveIPAddr("ip", DomainName)
	if err != nil {

		return "", err
	}
	return fmt.Sprintf("%s:%d", c.IP.String(), port), nil
}
