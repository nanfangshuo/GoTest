package Room

import (
	"GoTest/HttpRequest"
	"fmt"
)

//停止送风

func StopWind() error {
	var requestBody map[string]interface{}
	var response map[string]interface{}
	err, responseStatus := HttpRequest.SendPostRequestWithToken("/room/blowing/stop", requestBody, &response)
	if err != nil {
		fmt.Println("停止送风请求发送错误：", err)
		return err
	}
	if responseStatus == 200 {
		fmt.Println("发送停止送风请求成功")
		//TODO：停止送风的操作
		return nil
	} else {
		fmt.Println("停止送风请求失败：", response["message"])
	}
	return nil
}
