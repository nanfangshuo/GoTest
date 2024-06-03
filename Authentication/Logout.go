package Authentication

import (
	"GoTest/HttpRequest"
	"fmt"
)

type LogoutResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func Logout() error {
	var response LogoutResponse
	err, responseStatus := HttpRequest.SendPostRequestWithToken("/admin/logout", nil, response)
	if err != nil {
		fmt.Println("注销请求发送错误：", err)
		return err
	}
	if responseStatus == 200 {
		fmt.Println("成功注销")
		return nil
	} else {
		fmt.Println("注销失败：", response.Message)
	}
	return nil
}
