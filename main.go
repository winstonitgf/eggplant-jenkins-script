package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

var (
	h           bool
	apiScheme   string
	apiHost     string
	apiPort     string
	apiUser     string
	apiPassword string
	client      = &http.Client{}
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

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	LoginEggplant()

	ExecuteEggplantAPI()
}

func ExecuteEggplantAPI() {

	// 組合 Domain
	baseUri := fmt.Sprintf("%s://%s:%s", apiScheme, apiHost, apiPort)
	testListAPI := fmt.Sprintf("%s/api/test", baseUri)
	testExecuteAPI := fmt.Sprintf("%s/api/test/%s/execute", baseUri)
	// fmt.Println(baseUri)

	// 
	getTestList := 

	// apiUriPrefix := "api"
	// testResourceName := "test"
	// testRunResourceName := "test_run"
	// configurationResourceName := "execution_configuration"
	// execute := "execute"
	// executePollDuration := 5

}

func LoginEggplant() {

	// login API URL
	loginApi := fmt.Sprintf("%s://%s:%s/rest2/auth_user", apiScheme, apiHost, apiPort)

	// POST BODY
	// fmt.Println(fmt.Sprintf(`{"user_name":"%s","password":"%s"}`, apiUser, apiPassword))
	jsonStr := []byte(fmt.Sprintf(`{ "user_name": %s, "password": %s }`, apiUser, apiPassword))

	// 建立 Cookie 放到 Client 中，之後的請求會自動加入 Cookie
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client.Jar = jar

	// 建立 Request
	req, err := http.NewRequest("POST", loginApi, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}

	// 置入 Header
	req.Header.Set("Content-Type", `application/json`)

	// 開始請求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// 讀取結果
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
