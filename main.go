package main

import (
	"GoTest/Authentication"
	"GoTest/GetBills"
	"GoTest/Room"
	"GoTest/WindRequest"
	"fmt"
	"time"
)

var RefreshSpeed = 1

func main() {
	// 创建一个新的房间
	var room = Authentication.Login()
	fmt.Println(room)

	// 开启一个线程，每10/RefreshSpeed秒向主控机汇报从控机的状态
	quit := make(chan struct{})
	go func() {
		ticker := time.NewTicker(10 * time.Second / time.Duration(RefreshSpeed))
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				// 这里调用ReportStatus函数，你需要提供适当的参数
				err := Room.ReportStatus(room.WorkStatus, room.Temperature)
				if err != nil {
					// 处理错误
					fmt.Println("ReportStatus error:", err)
				}
			case <-quit:
				return
			}
		}
	}()
	for {
		fmt.Println("输入0注销")
		fmt.Println("输入1查看房间当前状态")
		fmt.Println("输入2获取报表")
		fmt.Println("输入3更改空调温度")
		fmt.Println("输入4更改空调风速")
		var x int
		fmt.Print("请输入需要的功能：")
		fmt.Scanln(&x)
		switch x {
		case 0:
			//注销
			goto end
		case 1:
			//查看房间当前状态
			fmt.Print("当前温度：", room.Temperature, "°C\n")
			fmt.Print("目标温度：", room.Temperature, "°C\n")
			fmt.Print("风速：", room.WindSpeed, "\n")
			fmt.Print("空调工作状态：", room.WorkStatus, "\n")
			break
		case 2:
			//获取并打印报表
			var period string
			fmt.Print("请输入查询周期（daily/weekly/monthly）：")
			fmt.Scanln(&period)
			GetBills.GetBills(period)
			break
		case 3:
			//更改空调温度
			var temp float64
			fmt.Print("请输入新的温度：")
			fmt.Scanln(&temp)
			if temp < room.Temperature && room.WorkStatus == "cold" {
				room.TargetTemperature = temp
				WindRequest.StartWind(room)
			} else if temp > room.Temperature && room.WorkStatus == "warm" {
				room.TargetTemperature = temp
				WindRequest.StartWind(room)
			} else {
				fmt.Println("该温度和当前空调工作模式矛盾，设置温度失败！")
			}
			break
		case 4:
			//更改空调风速
			var speed string
			fmt.Print("请输入新的风速（low/medium/high）：")
			fmt.Scanln(&speed)
			room.WindSpeed = speed
			WindRequest.StartWind(room)
			break
		}
	}
end: //退出循环的标记
	//注销
	err := Authentication.Logout()
	if err != nil {
		fmt.Println("Logout error:", err)
	} else {
		fmt.Println("Logout success")
	}
	close(quit) // 注销后关闭quit通道，使得goroutine停止运行
}
