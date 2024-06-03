package Room

import (
	"GoTest/HttpRequest"
	"fmt"
)

type ReportStatusRequestBody struct {
	Status      string  `json:"status"`
	Temperature float64 `json:"temperature"`
}
type ReportStatusResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// 向主控机汇报从控机的状态
func ReportStatus(status string, temperature float64) error {
	requestBody := ReportStatusRequestBody{
		Status:      status,
		Temperature: temperature,
	}
	var response ReportStatusResponse
	err, responseStatus := HttpRequest.SendPostRequestWithToken("/room/poll/room_status", requestBody, &response)
	if err != nil {
		fmt.Println("上报状态请求发送错误：", err)
		return err
	}
	if responseStatus == 200 {
		fmt.Println("上报状态请求成功")
		//TODO：上报状态成功后的操作
	} else {
		fmt.Println("上报状态请求失败：", response.Message)
	}
	return err
}
