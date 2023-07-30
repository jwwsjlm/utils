package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetSubstring(str, start, end string) (string, error) {
	startIndex := strings.Index(str, start)
	if startIndex == -1 {
		return "", fmt.Errorf("找不到起始字符串 %q", start)
	}

	endIndex := strings.Index(str, end)
	if endIndex == -1 {
		return "", fmt.Errorf("找不到结束字符串 %q", end)
	}

	startIndex += len(start)
	if startIndex >= endIndex {
		return "", fmt.Errorf("起始字符串 %q 在结束字符串 %q 之后", start, end)
	}

	return str[startIndex:endIndex], nil
}

func Getip() (string, error) {
	type location struct {
		Full_ip string `json:"full_ip"`
	}

	client := http.Client{}
	response, err := client.Get("https://forge.speedtest.cn/api/location/info")
	if err != nil {
		fmt.Println("请求发送失败:", err)
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return "", err
	}

	jsonData := string(body)
	fmt.Println("get请求获取成功", jsonData)
	var loc location
	err = json.Unmarshal([]byte(jsonData), &loc)
	fmt.Println(err)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err)
		return "", err
	}

	fmt.Println("解析json成功", loc.Full_ip)
	myip := loc.Full_ip
	return myip, nil

}
