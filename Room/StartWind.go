package Room

import (
	"GoTest/HttpRequest"
	"fmt"
	"time"
)

type StartWindRequestBody struct {
	FanSpeed   string  `json:"fanSpeed"`
	TargetTemp float64 `json:"targetTemp"`
}

// 发送送风请求
type StartWindResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func StartWind(room *Room) error {
	//首先，确保上一个请求已暂停（这非常重要！！）
	StopWind()
	//发送送风请求
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
		var requestState string
		//循环获取请求状态（间隔1秒），当请求状态为Doing或者Done时，停止循环
		for {
			err, requestState = GetRequestState()
			if err != nil {
				fmt.Println("获取请求状态时发生错误，结束循环：", err)
				break
			} else if requestState != "Pending" {
				fmt.Println("请求等待执行中")
				break
			}
			time.Sleep(1 * time.Second)
		}
		state0 := requestState
		//获得状态后，再执行以下代码。若此时请求状态为Done，则停止送风
		if state0 == "Done" {
			StopWind()
		} else if state0 == "Doing" {
			//若此时请求状态为Doing，则开始送风；同时监听请求状态，当请求状态为Done时，停止送风
			target := room.TargetTemperature
			var flag float64
			if room.WorkStatus == "Warm" {
				flag = 1
			} else {
				flag = -1
			}
			var degreeLevel float64
			switch room.WindSpeed {
			case "low":
				degreeLevel = 0.5 * flag
				break
			case "medium":
				degreeLevel = 1 * flag
				break
			case "high":
				degreeLevel = 1.5
				break
			}
			diff := (target - room.Temperature) * flag
			//循环获取请求状态（间隔1秒），当请求状态为Done时：stop <- true，停止送风
			for state0 == "Doing" && diff > 1 {
				room.Temperature += degreeLevel
				err, state0 = GetRequestState()
				if err != nil {
					fmt.Println("获取请求状态时发生错误，结束循环：", err)
					break
				}
				time.Sleep(1 * time.Second)
			}
		}
	} else {
		fmt.Println("送风请求失败：", response.Message)
	}
	return nil
}
