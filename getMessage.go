//获取总单

package main

import "fmt"

type MessageRequestBody struct {
	RoomNumber *string `json:"roomNumber,omitempty"`
}

type Message struct {
	RoomNumber        string `json:"roomNumber"`
	EnergyConsumption int    `json:"energyConsumption"`
	Cost              int    `json:"cost"`
}

type Response struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    Message `json:"data"`
}

func getMessage(roomNumber string) {
	data := MessageRequestBody{
		RoomNumber: &roomNumber,
	}

	var response Response
	err := sendPostRequest("/send/message", data, &response)
	if err != nil {
		fmt.Println("发送请求错误：", err)
		return
	}
	//TODO： 打印报表，可以删去
	fmt.Println("成功发送信息，房间号：", response.Data.RoomNumber)
	fmt.Println("能耗：", response.Data.EnergyConsumption)
	fmt.Println("费用：", response.Data.Cost)
}
