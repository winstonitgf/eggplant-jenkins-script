package main

import (
	"bytes"
	"eggplant-jenkins/models"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
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
	apiToken    string
	client      = &http.Client{}
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&apiScheme, "apiScheme", "http", "http or https (depending on server configuration, http is default)")
	flag.StringVar(&apiHost, "apiHost", "127.0.0.1", "IP address of ePM server")
	flag.StringVar(&apiPort, "apiPort", "8080", "Port of ePM server")
	flag.StringVar(&apiUser, "apiUser", "admin", "ePM user username")
	flag.StringVar(&apiPassword, "apiPassword", "admin", "ePM user password")
	flag.StringVar(&apiToken, "apiToken", "898b12d0ce0cd78d9675289525fa1a05220ba5e9e7cc98177a14853e08c25714", "ePM user token")
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	} else {
		Start()
	}
}

func Start() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ok := true
	var testResult []int
	testCaseList := GetEggplantTestCaseList()
	for _, testcase := range testCaseList.Tests {

		// 如果是 ACTIVE 狀態就跑測試
		if testcase.IsActive {

			// 執行腳本
			execute := ExecuteEggplantTestCase(testcase.Id)

			for {

				// 取得測試結果
				testExecuteResult := GetEggplantExecuteResult(execute.Id)
				if testExecuteResult.StatusCode != 100 {
					testResult = append(testResult, testExecuteResult.StatusCode)

					// 狀態不是200, 代表測試不過
					if ok && testExecuteResult.StatusCode != 200 {
						ok = false
					}
					break
				}
			}
		}
	}

	if !ok {
		panic(errors.New("測試不通過"))
	}
}

func GetEggplantTestCaseList() (testCase models.TestModel) {

	// 組合 Domain
	baseUri := fmt.Sprintf("%s://%s@%s:%s", apiScheme, apiToken, apiHost, apiPort)
	api := fmt.Sprintf("%s/api/test", baseUri)

	body := HttpDo("GET", api, nil)

	if err := json.Unmarshal(body, &testCase); err != nil {
		panic(err)
	}

	return
}

func ExecuteEggplantTestCase(id string) (testExecuteResult models.TestExecuteResultModel) {

	// 組合 Domain
	baseUri := fmt.Sprintf("%s://%s@%s:%s", apiScheme, apiToken, apiHost, apiPort)
	api := fmt.Sprintf("%s/api/test/%s/execute", baseUri, id)

	body := HttpDo("GET", api, nil)

	if err := json.Unmarshal(body, &testExecuteResult); err != nil {
		panic(err)
	}

	return
}

func GetEggplantExecuteResult(id string) (testExecuteResult models.TestExecuteResultModel) {

	// 組合 Domain
	baseUri := fmt.Sprintf("%s://%s@%s:%s", apiScheme, apiToken, apiHost, apiPort)
	api := fmt.Sprintf("%s/api/test_run/%s", baseUri, id)

	body := HttpDo("GET", api, nil)

	if err := json.Unmarshal(body, &testExecuteResult); err != nil {
		panic(err)
	}

	return
}

func LoginEggplant() {

	// login API URL
	loginApi := fmt.Sprintf("%s://%s:%s/rest2/auth_user", apiScheme, apiHost, apiPort)
	// POST BODY
	// fmt.Println(fmt.Sprintf(`{"user_name":"%s","password":"%s"}`, apiUser, apiPassword))
	jsonStr := []byte(fmt.Sprintf(`{ "user_name": "%s", "password": "%s" }`, apiUser, apiPassword))
	fmt.Println(fmt.Sprintf(`{ "user_name": %s, "password": %s }`, apiUser, apiPassword))
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
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(loginApi, ", 登入：", resp.Status)
	resp.Body.Close()
}

func HttpDo(method, url string, payload io.Reader) (body []byte) {

	// 建立 Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	// 開始請求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// 讀取結果
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("url：", url)
	fmt.Println("status: ", resp.Status)

	return
}
