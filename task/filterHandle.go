package task

import (
	"fmt"
	"github.com/wangzyg/netflowAnalyser/netflow5"
)

func DoRouterFilter(p *netflow5.Packet){

	for _, task := range Tasks.Task{
		taskName := task.TaskName
		fmt.Println("taskName:", taskName)

		for _, netflowRecord := range p.Records {

			//过滤路由器ip
			var passRouterFilter bool = false
			for _, routerAddress := range task.RouterFilter.RouterAddress {
				if routerAddress == netflowRecord.RouterAddr.IP.String() {
					passRouterFilter = true
				}
			}
			fmt.Println("RouterFilter check:", passRouterFilter)
			if task.RouterFilter.RouterAddress != nil && !passRouterFilter {
				continue
			}

			//过滤源ip
			var passSourceFilter bool = false
			for _, sourceAddress := range task.SourceFilter.SourceAddress {
				if sourceAddress == netflowRecord.SrcAddr.String() {
					passSourceFilter = true
				}
			}
			fmt.Println("SourceFilter check:", passSourceFilter)
			if task.SourceFilter.SourceAddress != nil && !passSourceFilter {
				continue
			}

			//过滤源ip
			var passDestFilter bool = false
			for _, destAddress := range task.DestFilter.DestAddress {
				if destAddress == netflowRecord.DstAddr.String() {
					passDestFilter = true
				}
			}
			fmt.Println("DestFilter check:", passDestFilter)
			if task.DestFilter.DestAddress != nil && !passDestFilter {
				continue
			}

		}


	}
}
