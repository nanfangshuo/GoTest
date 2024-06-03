// Description: 用于发送HTTP请求的工具函数
package SendRequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BaseURL = "http://82.156.104.168:80"

// SendPostRequestWithReq 用于发送已构造HTTP请求req的工具函数
// 该函数接受一个http.Request类型的参数req，一个请求体requestBody，一个响应体responseBody
func SendPostRequestWithReq(req *http.Request, requestBody interface{}, responseBody interface{}) error {
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("将请求体转换为JSON错误：%v", err)
	}

	req.Body = io.NopCloser(bytes.NewBuffer(jsonData))
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
