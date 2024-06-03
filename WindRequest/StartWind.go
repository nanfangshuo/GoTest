package WindRequest

import (
	"GoTest/HttpRequest"
	"fmt"
)

//发送送风请求

type StartWindRequestBody struct {
	FanSpeed   string  `json:"fanSpeed"`
	TargetTemp float64 `json:"targetTemp"`
}

type StartWindResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func StartWind(targetTemerature float64, windSpeed string) error {
	requestBody := StartWindRequestBody{
		FanSpeed:   windSpeed,
		TargetTemp: targetTemerature,
	}
	var response StartWindResponse
	err, responseStatus := HttpRequest.SendPostRequestWithToken("/room/blowing/start", requestBody, &response)
	if err != nil {
		fmt.Println("送风请求发送错误：", err)
		return err
	}
	if responseStatus == 200 {
		fmt.Println("送风请求成功")
		//TODO：发送请求成功后的操作
		return nil
	} else {
		fmt.Println("送风请求失败：", response.Message)
	}
	return nil
}
