package Room

import (
	"GoTest/HttpRequest"
	"fmt"
	"sync"
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
		//循环获取请求状态（间隔1秒），当请求状态为Doing或者Done时，停止循环
		ticker := time.NewTicker(1 * time.Second)
		var wg sync.WaitGroup
		wg.Add(1)
		result := make(chan string)
		go func() {
			defer wg.Done()
			for range ticker.C {
				err, requestState := GetRequestState()
				if err != nil {
					fmt.Println("获取请求状态时发生错误，结束循环：", err)
					ticker.Stop()
					result <- "Error"
					break
				} else if requestState == "Pending" {
					fmt.Println("请求等待执行中")
				} else {
					break
				}
			}
		}()
		wg.Wait()
		state0 := <-result
		close(result)
		//执行完毕上述协程后，再执行以下代码。若此时请求状态为Done，则停止送风
		if state0 == "Done" {
			StopWind()
		} else if state0 == "Doing" {
			//若此时请求状态为Doing，则开始送风；同时监听请求状态，当请求状态为Done时，停止送风
			var stopChangingTemperature = make(chan bool)
			go room.WorkingTemperatureChange(stopChangingTemperature)
			//循环获取请求状态（间隔1秒），当请求状态为Done时：stop <- true，停止送风
			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					_, requestState := GetRequestState()
					if requestState == "Done" {
						stopChangingTemperature <- true
						return nil
					}
				case <-stopChangingTemperature:
					return nil
				}
			}
			return nil
		}
	} else {
		fmt.Println("送风请求失败：", response.Message)
	}
	return nil
}
