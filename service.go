package main

import "github.com/go-resty/resty/v2"

const (
	baseUrl = "https://lims.fudan.edu.cn/api"
	listUrl = baseUrl + "/limsproduct/fdulims/apiForMiniProgram/listOpenableLab"
	openUrl = baseUrl + "/limsproduct/fdulims/wxAPI/openDoorPython"
)

var client = resty.New().
	SetHeaders(
		map[string]string{
			"Content-Type": "application/json",
			"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF WindowsWechat(0x63090b11)XWEB/9185",
		},
	)
