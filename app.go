package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type Response struct {
	Rs       int    `json:"rs"`
	Code     int    `json:"code"`
	Address  string `json:"address"`
	Ip       string `json:"ip"`
	IsDomain int    `json:"isDomain"`
}

func (a *App) Greet(ip string) Response {
	url := "https://www.ip.cn/api/index?ip=" + ip + "&type=1"
	resp, err := http.Get(url)
	fmt.Println(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return Response{}
	}
	defer resp.Body.Close()

	// 检查HTTP响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to get a successful response:", resp.StatusCode)
		return Response{}
	}

	// 读取响应体
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return Response{}
	}

	// 解析JSON数据到Response结构体
	var response Response
	if err := json.Unmarshal(jsonData, &response); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return Response{}
	}
	fmt.Println(response)
	return response
}
