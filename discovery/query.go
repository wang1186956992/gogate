package discovery

import (
	"strconv"
)

func QueryEureka() ([]*InstanceInfo, error) {
	apps, err := euClient.GetApplications()
	if nil != err {
		return nil, err
	}

	var instances []*InstanceInfo
	for _, app := range apps.Applications {
		// 服务名
		servName := app.Name

		// 遍历每一个实例
		for _, ins := range app.Instances {
			// 跳过无效实例
			if nil == ins.Port || ins.Status != "UP" {
				continue
			}

			addr := ins.HostName + ":" + strconv.Itoa(ins.Port.Port)
			var meta map[string]string
			if nil != ins.Metadata {
				meta = ins.Metadata.Map
			}

			instances = append(
				instances,
				&InstanceInfo{
					ServiceName: servName,
					Addr: addr,
					Meta: meta,
				},
			)
		}
	}

	return instances, nil
}
