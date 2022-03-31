package main

func main() {

	r := HttpXproxy{}
	r.StartXproxy("127.0.0.1:9091")
}
