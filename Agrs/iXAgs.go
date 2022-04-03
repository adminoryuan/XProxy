package agrs

import (
	"errors"
	"runtime"
	"strconv"
)

type iXargs interface {
	Cmmd()
}
type AgrsAgent struct {
	Ip string
}
type xargsImp struct {
	agent   AgrsAgent
	cmdFunc map[string]func(val string)
}

func NewXargs() iXargs {
	imp := xargsImp{}

	imp.cmdFunc = make(map[string]func(string))

	imp.cmdFunc["-x"] = func(val string) {
		runNumber, er := strconv.Atoi(val)
		if er != nil {
			panic(errors.New("输入错误!!!"))
		}
		runtime.GOMAXPROCS(runNumber)
	}
	imp.cmdFunc["-a"] = func(val string) {

	}
	return imp
}
func (a xargsImp) Cmmd() {

}
