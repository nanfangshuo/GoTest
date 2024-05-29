package main

import "fmt"

type EnergyCostResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		EnergyConsumed int `json:"energyConsumed"`
		AmountDue      int `json:"amountDue"`
	} `json:"data"`
}

func getEnergyAndCost(roomNumber string) (int, int) {
	data := map[string]string{
		"roomNumber": roomNumber,
	}

	var response EnergyCostResponse
	err := sendPostRequest("poll/billing", data, &response)
	if err != nil {
		fmt.Println("Error getting energy and cost: ", err)
		return 0, 0
	}

	return response.Data.EnergyConsumed, response.Data.AmountDue
}

type RequestStateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		status string `json:"status"`
	} `json:"data"`
}

func getRequestState(roomNumber string) {
	data := map[string]string{
		"roomNumber": roomNumber,
	}

	var response RequestStateResponse
	err := sendPostRequest("poll/requestState", data, &response)
	if err != nil {
		fmt.Println("Error getting request state: ", err)
		return
	}

	fmt.Println("Request state: ", response.Data.status)
}
