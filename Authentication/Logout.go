package Authentication

import (
	"GoTest/SendRequest"
	"fmt"
	"net/http"
)

type LogoutResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func Logout(token string) error {
	req, err := http.NewRequest("POST", "/admin/logout", nil)
	if err != nil {
		fmt.Println("创建请求失败：", err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	err = SendRequest.SendPostRequestWithReq(req, nil, nil)
	if err != nil {
		fmt.Println("注销失败：", err)
		return err
	}

	fmt.Println("成功注销")
	return nil
}
