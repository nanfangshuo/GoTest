// 用于开机时的认证登录
// 调用func Login()，会一直执行登录直至登录成功。返回房间号、空调模式、缺省温度

package main

import (
	"fmt"
)

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Mode        string `json:"mode"`
		Temperature int    `json:"temperature"`
	} `json:"data"`
}

func Login() (string, string, int) {
	var roomId string
	var password string
	for true {
		fmt.Println("请输入房间号：")
		fmt.Scanln(&roomId)
		fmt.Println("请输入房间密码：")
		fmt.Scanln(&password)
		data := map[string]string{
			"roomNumber": roomId,
			"id":         password,
		}
		var loginResp LoginResponse
		err := sendPostRequest("/admin/login", data, &loginResp)
		if err != nil {
			fmt.Println("登录失败，请重新输入")
		} else {
			fmt.Println("登录成功")
			return roomId, loginResp.Data.Mode, loginResp.Data.Temperature
		}
	}
	return "", "", 0
}
