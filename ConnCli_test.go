package main

import (
	"fmt"
	"net"
	"strings"
	"testing"
)

func TestConn(t *testing.T) {
	url := strings.ReplaceAll("http://www.240ps.com/", "/", "")
	url = strings.ReplaceAll(url, "http:", "")
	fmt.Println(url)
	c, _ := net.ResolveIPAddr("ip", url)
	fmt.Println(c.IP)
}
