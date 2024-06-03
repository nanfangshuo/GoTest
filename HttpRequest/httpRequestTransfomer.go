// Description: 用于发送HTTP请求的工具函数
package HttpRequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseURL = "http://82.156.104.168:80"

var Token string

// SendPostRequestWithReq 用于向指定接口发送带有token的指定消息
// 三个参数：请求接口（string类型），请求体requestBody，响应体responseBody
func SendPostRequestWithToken(api string, requestBody interface{}, responseBody interface{}) (error, int) {
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("将请求体转换为JSON错误：%v", err), -1
	}

	req, err := http.NewRequest("POST", BaseURL+api, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("构造请求错误：%v", err), -1
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求错误：%v", err), -1
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("请求失败，状态码：%d", resp.StatusCode), -1
	}

	if err := json.NewDecoder(resp.Body).Decode(responseBody); err != nil {
		return fmt.Errorf("解析响应错误：%v", err), -1
	}

	return nil, resp.StatusCode
}

func SendPostRequestWithoutToken(api string, requestBody interface{}, responseBody interface{}) (error, int) {
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("将请求体转换为JSON错误：%v", err), -1
	}

	req, err := http.NewRequest("POST", BaseURL+api, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("构造请求错误：%v", err), -1
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求错误：%v", err), -1
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("请求失败，状态码：%d", resp.StatusCode), -1
	}

	if err := json.NewDecoder(resp.Body).Decode(responseBody); err != nil {
		return fmt.Errorf("解析响应错误：%v", err), -1
	}

	return nil, resp.StatusCode
}
