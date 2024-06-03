package Authentication

import (
	"GoTest/SendRequest"
	"fmt"
)

type LogoutResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func Logout(token string) error {
	var response LogoutResponse
	err := SendRequest.SendPostRequestWithToken("/admin/logout", nil, response)
	if err != nil {
		fmt.Println("注销请求发送错误：", err)
		return err
	}
	if response.Code == 200 {
		fmt.Println("成功注销")
		return nil
	} else {
		fmt.Println("注销失败：", response.Message)
	}
	return nil
}
