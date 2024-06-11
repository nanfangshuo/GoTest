package main

import (
	"GoTest/Authentication"
	"GoTest/GetBills"
	"GoTest/Room"
	"fmt"
	"time"
)

var RefreshSpeed = 1

func main() {
	// 创建一个新的房间
	var room = Authentication.Login()
	fmt.Println(room)

	// 开启一个线程，每4/RefreshSpeed秒向主控机汇报从控机的状态，每秒检查温度是否需要发起请求
	quit := make(chan struct{})
	go reportStatusPeriodically(room, quit)
	go checkTemperaturePeriodically(room, quit)
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
			fmt.Print("目标温度：", room.TargetTemperature, "°C\n")
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
			if (room.WorkStatus == "Cool" && temp >= 18 && temp <= 25) || (room.WorkStatus == "Warm" && temp >= 25 && temp <= 30) {
				room.TargetTemperature = temp
			} else {
				fmt.Println("该温度和当前空调工作模式矛盾，设置温度失败！")
			}
			break
		case 4:
			//更改空调风速度
			var speed string
			fmt.Print("请输入新的风速（low/medium/high）：")
			fmt.Scanln(&speed)
			room.WindSpeed = speed
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
	close(quit) // 注销后关闭quit通道，使得goroutine停止运
}

func reportStatusPeriodically(room *Room.Room, quit chan struct{}) {
	ticker1 := time.NewTicker(3 * time.Second / time.Duration(RefreshSpeed))
	defer ticker1.Stop()
	for {
		select {
		case <-ticker1.C:
			err, workStatus, refreshSpeed := Room.ReportStatus(room.WorkStatus, room.Temperature)
			if err != nil {
				fmt.Println("ReportStatus error:", err)
			} else {
				if workStatus != room.WorkStatus {
					Room.StopWind()
					room.WorkStatus = workStatus
					if room.WorkStatus == "Cool" {
						room.TargetTemperature = 22
					} else {
						room.TargetTemperature = 28
					}
				}
				if refreshSpeed != RefreshSpeed {
					RefreshSpeed = refreshSpeed
					ticker1.Stop()
					ticker1 = time.NewTicker(3 * time.Second / time.Duration(RefreshSpeed))
				}
			}
		case <-quit:
			return
		}
	}
}

func checkTemperaturePeriodically(room *Room.Room, quit chan struct{}) {
	ticker2 := time.NewTicker(1 * time.Second)
	defer ticker2.Stop()
	for {
		select {
		case <-ticker2.C:
			Room.CheckTemperature(room)
			if room.Temperature > 20.05 {
				room.Temperature -= 0.05
			} else if room.Temperature < 19.95 {
				room.Temperature += 0.05
			} else {
				room.Temperature = 20
			}
		case <-quit:
			return
		}
	}
}
