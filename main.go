package main

import (
	"flag"
	"fmt"
)

var (
	h           bool
	apiScheme   string
	apiHost     string
	apiPort     string
	apiUser     string
	apiPassword string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&apiScheme, "apiScheme", "http", "http or https (depending on server configuration, http is default)")
	flag.StringVar(&apiHost, "apiHost", "127.0.0.1", "IP address of ePM server")
	flag.StringVar(&apiPort, "apiPort", "8080", "Port of ePM server")
	flag.StringVar(&apiUser, "apiUser", "admin", "ePM user username")
	flag.StringVar(&apiPassword, "apiPassword", "admin", "ePM user password")

	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	// flag.Usage = usage
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}
}

func ExecuteEggplantAPI() {
	baseUri := fmt.Sprintf("%s://%s:%s", apiScheme, apiHost, apiPort)
	fmt.Println(baseUri)

	apiUriPrefix := "api"
	testResourceName := "test"
	testRunResourceName := "test_run"
	configurationResourceName := "execution_configuration"
	execute := "execute"
	executePollDuration := 5
}
