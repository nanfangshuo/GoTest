// Description: 用于发送HTTP请求的工具函数
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseURL = "http://localhost:8080"

func sendPostRequest(api string, requestBody interface{}, responseBody interface{}) error {
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("将请求体转换为JSON错误：%v", err)
	}
	var url = BaseURL + api
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("创建请求错误：%v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求错误：%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("请求失败，状态码：%d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(responseBody); err != nil {
		return fmt.Errorf("解析响应错误：%v", err)
	}

	return nil
}
