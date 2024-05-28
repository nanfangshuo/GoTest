package main

import "fmt"

type LogoutRequest struct {
	RoomNumber string `json:"roomNumber"`
}

func Logout(roomNumber string) error {
	data := LogoutRequest{
		RoomNumber: roomNumber,
	}

	err := sendPostRequest("/admin/logout", data, nil)
	if err != nil {
		fmt.Println("注销失败：", err)
		return err
	}

	fmt.Println("成功注销")
	return nil
}
