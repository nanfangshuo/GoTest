package main

import (
	"GoTest/Authentication"
	"GoTest/Room"
	"fmt"
)

var RefreshSpeed = 1

func main() {

	//登录，获得房间号、空调模式、缺省温度、房间当前温度
	//使用登录后的信息，创建房间对象，创建时，需要检查目标温度和当前温度的插值，小于一将模式设为standby风速0，大于一将模式设为warm/cold，发送温控请求
	var room *Room.Room
	room = Authentication.Login()
	fmt.Println("房间号：", room.RoomId)
	fmt.Println("空调模式：", room.WorkStatus)
	fmt.Println("缺省温度：", room.TargetTemperature)
	fmt.Println("当前温度：", room.Temperature)
}
