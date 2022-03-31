package untity

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

func PaserIP(DomainName string) (string, error) {
	DomainName = strings.ReplaceAll(DomainName, "/", "")
	DomainName = strings.ReplaceAll(DomainName, "http:", "")
	if strings.EqualFold(DomainName, "") {
		return "", errors.New("地址错误")
	}

	c, err := net.ResolveIPAddr("ip", DomainName)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:80", c.IP.String()), nil
}
