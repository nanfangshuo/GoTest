package WindRequest

import (
	"GoTest/HttpRequest"
	"GoTest/Room"
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

func StartWind(room *Room.Room) error {
	requestBody := StartWindRequestBody{
		FanSpeed:   room.WindSpeed,
		TargetTemp: room.TargetTemperature,
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
		//循环获取请求状态，当请求状态为Doing或者Done时，停止循环

		//若此时请求状态为Done，则停止送风

		//若此时请求状态为Doing，则开始送风；同时监听请求状态，当请求状态为Done时，停止送风
		return nil
	} else {
		fmt.Println("送风请求失败：", response.Message)
	}
	return nil
}
