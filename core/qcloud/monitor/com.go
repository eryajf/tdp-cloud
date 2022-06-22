package monitor

import (
	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud"

	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
)

// 创建客户端

func NewClient(ud *midware.Userdata) (*monitor.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(ud)

	if ud.Region != "" {
		cpf.HttpProfile.Endpoint = "monitor." + ud.Region + ".tencentcloudapi.com"
	}

	return monitor.NewClient(credential, ud.Region, cpf)

}