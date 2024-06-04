package GetBills

import (
	"GoTest/HttpRequest"
	"fmt"
	"time"
)

type getBillsRequestBody struct {
	Period string `json:"period"`
}

type getBillsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RoomId      string    `json:"roomId"`
		SwitchTimes int       `json:"switchTimes"`
		Requests    []Request `json:"requests"`
		TotalCost   float64   `json:"totalCost"`
	} `json:"data"`
}

type Request struct {
	StartTime        time.Time `json:"startTime"`
	EndTime          time.Time `json:"endTime"`
	StartTemperature int       `json:"startTemperature"`
	EndTemperature   int       `json:"endTemperature"`
	FanSpeed         string    `json:"fanSpeed"`
	Cost             float64   `json:"cost"`
}

func GetBills(period string) error {
	//把period填入requestBody
	requestBody := getBillsRequestBody{
		Period: period,
	}

	var response getBillsResponse
	err, rspStatus := HttpRequest.SendPostRequestWithToken("/room/report", requestBody, &response)
	if err != nil {
		fmt.Println("获取账单请求发送错误：", err)
		return err
	}
	if rspStatus == 200 {
		fmt.Println("获取账单请求成功")
		// 打印账单
		fmt.Println("以下是选中时间内的账单")
		fmt.Println("房间号：", response.Data.RoomId)
		fmt.Println("开关次数：", response.Data.SwitchTimes)
		fmt.Println("请求详情如下")
		for i, request := range response.Data.Requests {
			fmt.Println("请求", i+1)
			fmt.Println("开始时间：", request.StartTime)
			fmt.Println("结束时间：", request.EndTime)
			fmt.Println("开始温度：", request.StartTemperature)
			fmt.Println("结束温度：", request.EndTemperature)
			fmt.Println("风速：", request.FanSpeed)
			fmt.Println("费用：", request.Cost)
			fmt.Println("总费用：", response.Data.TotalCost)
		}
	} else {
		fmt.Println("获取账单请求失败：", response.Message)
	}

	return nil
}
