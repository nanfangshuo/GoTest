package Room

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
	x := 28.0
	return &Room{
		RoomId:            roomId,
		WorkStatus:        mode,
		Temperature:       x, // TODO:可能要改为从服务器获取当前温度，这里先写死
		TargetTemperature: targetTemperature,
		WindSpeed:         "low",
	}
}

func CheckTemperature(room *Room) {
	//检查温度是否需要发起请求
	diff := room.Temperature - room.TargetTemperature
	if (room.WorkStatus == "Warm" && diff < -1) || (room.WorkStatus == "Cool" && diff > 1) {
		//向服务器请求送风
		go func() {
			StartWind(room)
		}()
	}
}
