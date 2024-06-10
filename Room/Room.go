package Room

import (
	"fmt"
	"math"
)

type Room struct {
	// 房间号
	RoomId string
	// 空调工作状态，有三种，warm/cold/standby
	WorkStatus string
	// 房间当前温度，开始时应该初始化为多少呢，这里我觉得应该是服务器先写好每个房间的初始温度，然后初始化的时候从服务器获取，和开机次数和关机次数一样，可能需要继承之前的值
	Temperature float64
	// 房间目标温度
	TargetTemperature float64
	// 房间风速，低中高分别对应1，2，3，关闭时为0
	WindSpeed string
}

func NewRoom(roomId string, mode string, targetTemperature float64) *Room {
	//新建房间时检查温度是否需要发起请求
	return &Room{
		RoomId:            roomId,
		WorkStatus:        mode,
		Temperature:       20,
		TargetTemperature: targetTemperature,
		WindSpeed:         "low",
	}
}

func CheckTemperature(room *Room) {
	//检查温度是否需要发起请求
	diff := room.Temperature - room.TargetTemperature
	if math.Abs(diff) > 1 {
		//向服务器请求送风
		fmt.Printf("当前温度为%.1f，目标温度为%.1f，温度差大于1，请求送风", room.Temperature, room.TargetTemperature)
		StartWind(room)
	}
}
