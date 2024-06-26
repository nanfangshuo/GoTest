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
	_, responseStatus := HttpRequest.SendPostRequestWithToken("/room/poll/request", requestBody, &response)
	if responseStatus == 200 {
		fmt.Println("获取请求状态成功, 请求状态为：", response.Data.RequestStatus)
		return nil, response.Data.RequestStatus
	} else {
		fmt.Println("获取请求状态失败：", response.Message, "错误码", responseStatus)

	}
	return nil, ""
}
