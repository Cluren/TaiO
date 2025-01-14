package service

import (
	"fmt"
	"github.com/binganao/Taio/pkg/logger"
	"github.com/binganao/Taio/service/finger"
	"github.com/binganao/Taio/service/mas"
	"github.com/binganao/Taio/service/nma"
	"github.com/binganao/Taio/service/save"
)

func Probe(ip string) {
	var web [][]string

	logger.Info("开始对目标 " + ip + " 进行端口探测")
	ports := mas.MasScan(ip)
	logger.Info("目标 " + ip + " 开启了以下端口: ")

	var dP string
	for i, p := range ports {
		if i == 0 {
			dP += p
		} else {
			dP += "," + p
		}
	}
	fmt.Println(dP)

	logger.Info("开始对目标 " + ip + " 进行服务探测")
	services := nma.NmapScan(ip, ports)
	logger.Info("目标 " + ip + " 开启了以下服务: ")

	var dS string
	for i, p := range services {
		if i == 0 {
			dS += p
		} else {
			dS += "," + p
		}
	}
	fmt.Println(dS)

	logger.Info("开始对目标 " + ip + " 进行指纹探测")
	for _, s := range services {
		if s != "" {
			wt := finger.FingerScan(ip, s)
			if wt != nil {
				web = append(web, wt)
			}
		}
	}

	var dW string
	for i, wa := range web {
		if i == 0 {
			for j, a := range wa {
				if j == 0 {
					dW += a
				} else {
					dW += "," + a
				}
			}
		} else {
			dW += ";"
			for j, a := range wa {
				if j == 0 {
					dW += a
				} else {
					dW += "," + a
				}
			}
		}
	}
	logger.Info("目标 " + ip + " 的指纹识别结果: ")
	fmt.Println(dW)

	save.Save(ip, dP, dS, dW)
}
