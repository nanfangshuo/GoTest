// 用于开机时的认证登录
// 调用func Login()，会一直执行登录直至登录成功。返回房间号、空调模式、缺省温度

package Authentication

import (
	"GoTest/Room"
	"GoTest/SendRequest"
	"fmt"
	"net/http"
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
		defaultTemperature float64 `json:"defaultTemp"`
		windSpeed          int     `json:"refreshRate"`
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
	for true {
		fmt.Println("请输入roomId：")
		fmt.Scanln(&roomId)
		fmt.Println("请输入identity：")
		fmt.Scanln(&identity)

		var data = LoginRequestBody{
			RoomNumber: roomId,
			Id:         identity,
		}
		var loginResp LoginResponse

		req, err := http.NewRequest("POST", SendRequest.BaseURL+"/auth/login", nil)
		if err != nil {
			fmt.Println("构造请求错误：", err)
			continue
		}
		err = SendRequest.SendPostRequestWithReq(req, data, &loginResp)
		if err != nil {
			fmt.Println("登录错误：", err)
			continue
		}

		if loginResp.Code == 200 {
			room = Room.NewRoom(roomId, loginResp.Data.Token, loginResp.Data.Mode, loginResp.Data.defaultTemperature, loginResp.Data.windSpeed)
			break
		} else {
			fmt.Println("登录失败：", loginResp.Message)
		}
	}
	return room
}
