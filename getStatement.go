// 用于发送/send/statement/请求，获取统计报表
// 调用func getStatement(roomNumber string, statementType string)，返回Statement
package main

import (
	"fmt"
)

// 发送请求时的body结构
type StatementRequestBody struct {
	RoomNumber    string `json:"roomNumber"`
	StatementType string `json:"statementType"`
}

// 获取到的报表结构，这里要获取总能量就用TotalCost除以电费单价吧
type Statement struct {
	Code                int                  `json:"code"`    //响应内容的一部分，code在最后的报表中无需显示
	Message             string               `json:"message"` //响应内容的一部分，message在最后的报表中无需显示
	RoomNumber          string               `json:"roomNumber"`
	OnOffCount          int                  `json:"onOffCount"`
	TemperatureRequests []TemperatureRequest `json:"temperatureRequests"`
	TotalCost           int                  `json:"totalCost"`
}

// 温控请求结构
type TemperatureRequest struct {
	StartTime         string `json:"startTime"`
	StartTemperature  int    `json:"startTemperature"`
	EnergyConsumption int    `json:"energyConsumption"`
	Cost              int    `json:"cost"`
}

func getStatement(roomNumber string, statementType string) (Statement, error) {
	var emptyStatement Statement
	data := StatementRequestBody{
		RoomNumber:    roomNumber,
		StatementType: statementType,
	}
	var statement Statement
	err := sendPostRequest("/send/statement", data, &statement)
	if err != nil {
		fmt.Print("发送请求错误：")
		fmt.Println(err)
		return emptyStatement, err
	}

	//TODO： 打印报表，可以删去
	fmt.Println("成功发送统计报表，房间号：", statement.RoomNumber)
	fmt.Println("开关机次数：", statement.OnOffCount)
	fmt.Println("总费用：", statement.TotalCost)
	for _, tempReq := range statement.TemperatureRequests {
		fmt.Println("开始时间：", tempReq.StartTime)
		fmt.Println("开始温度：", tempReq.StartTemperature)
		fmt.Println("能耗：", tempReq.EnergyConsumption)
		fmt.Println("费用：", tempReq.Cost)
	}

	return statement, nil
}
