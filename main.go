package main

import "runtime"

func main() {

	runtime.GOMAXPROCS(10)
	r := HttpXproxy{}
	r.StartXproxy("127.0.0.1:9092")
}
