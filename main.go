package main

import "fmt"

func main() {
	var roomId, mode, temperature = Login()
	fmt.Println("登录成功，当前房间号为" + roomId)
	fmt.Println("模式为：" + mode)
	fmt.Println("温度为：" + string(temperature))
}
