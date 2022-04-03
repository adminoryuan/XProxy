package main

//处理一些命令行参数
type XArgs struct {
	xargsFunc map[string]func() interface{}
}
