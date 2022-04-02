package main

import "runtime"

func main() {

	runtime.GOMAXPROCS(10)
	r := HttpXproxy{}
	r.StartXproxy("0.0.0.0:9092")
}
