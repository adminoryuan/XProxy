package untity

import (
	"fmt"
	"testing"
)

func TestHttp(t *testing.T) {
	a, er := PaserIP("v.qq.com:443")
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println("ss")
	fmt.Println(a)
}
