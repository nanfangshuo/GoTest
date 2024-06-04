package Room

import (
	"GoTest/HttpRequest"
	"fmt"
)

type getRequestStateRequestBody struct {
	FanSpeed   string  `json:"fanSpeed"`
	TargetTemp float64 `json:"targetTemp"`
}

type getRequestStateResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RequestStatus string `json:"requestStatus"`
	} `json:"data"`
}

func GetRequestState() (error, string) {
	var requestBody getRequestStateRequestBody
	var response getRequestStateResponse
	err, responseStatus := HttpRequest.SendPostRequestWithToken("/room/poll/request", requestBody, &response)
	if err != nil {
		fmt.Println("获取请求状态错误：", err)
		return err, ""
	}
	if responseStatus == 200 {
		fmt.Println("获取请求状态成功")

		return nil, response.Data.RequestStatus
		//TODO:对于两种状态的处理（Doing Pending Done）

	} else {
		fmt.Println("获取请求状态失败：", response.Message)
	}
	return nil, ""
}
