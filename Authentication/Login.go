// 用于开机时的认证登录
// 调用func Login()，会一直执行登录直至登录成功。返回房间号、空调模式、缺省温度

package Authentication

import (
	"GoTest/HttpRequest"
	"GoTest/Room"
	"fmt"
)

type LoginRequestBody struct {
	RoomNumber string `json:"roomId"`
	Id         string `json:"identity"`
}

type LoginResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Mode               string  `json:"mode"`
		DefaultTemperature float64 `json:"defaultTemp"`
		refreshRate        int     `json:"refreshRate"`
		Token              string  `json:"token"`
	} `json:"data"`
}

// 用于开机时的认证登录
// 调用func Login()，会一直执行登录直至登录成功。
// 使用得到的response信息，构造并返回新创建的room结构体
func Login() *Room.Room {
	var room *Room.Room
	var roomId string
	var identity string
	for {
		fmt.Println("请输入roomId：")
		fmt.Scanln(&roomId)
		fmt.Println("请输入identity：")
		fmt.Scanln(&identity)

		var data = LoginRequestBody{
			RoomNumber: roomId,
			Id:         identity,
		}
		var loginResp LoginResponse

		err, responseStatus := HttpRequest.SendPostRequestWithoutToken("/auth/login", data, &loginResp)
		if err != nil {
			fmt.Println("发送登录请求错误：", err)
			continue
		}

		if responseStatus == 200 {
			room = Room.NewRoom(roomId, loginResp.Data.Mode, loginResp.Data.DefaultTemperature)
			HttpRequest.Token = loginResp.Data.Token
			fmt.Println("登录成功")
			break
		} else {
			fmt.Println("响应码：", responseStatus)
			fmt.Println("登录失败：", loginResp.Message)
		}
	}
	return room
}
