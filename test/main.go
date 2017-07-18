package main

import (
	"github.com/wangzyg/netflowAnalyser/qqwry"
	"log"
	"fmt"
	"reflect"
	"net/url"
	//"io/ioutil"
	//"bytes"
	//
	//"golang.org/x/text/encoding/simplifiedchinese"
	//"golang.org/x/text/transform"
)


func main(){
	datFile := "./qqwry.dat"

	qqwry.IPData.FilePath = datFile

	res := qqwry.IPData.InitIPData()

	if v, ok := res.(error); ok {
		log.Panic(v)
	}

	qqWry := qqwry.NewQQwry()

	ipLocate := qqWry.Find("101.81.30.99")

	fmt.Println(reflect.TypeOf(ipLocate))
	fmt.Print(ipLocate)

	//src:="编码转换内容内容"
	////data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))
	//fmt.Println(data) //byte
	//fmt.Println(string(data))  //打印为乱码

	fmt.Println("-----")
	s := "%7B%22agentEuqTasks%22%3A%5B%7B%22agentHost%22%3A%221171_80_1%22%2C%22equId%22%3A80%2C%22equName%22%3A%22hx-zabbix-test-161%22%2C%22ipAddress%22%3A%22172.30.60.161%22%2C%22status%22%3A1%2C%22taskId%22%3A1171%7D%5D%2C%22cmHxZabbixItemBeans%22%3A%5B%7B%22hxTaskId%22%3A0%2C%22itemId%22%3A0%2C%22taskDiscoveryRuleId%22%3A0%2C%22vId%22%3A11%2C%22zabbixHostId%22%3A0%2C%22zabbixItemId%22%3A0%2C%22zbApplicationId%22%3A1%2C%22zbApplicationName%22%3A%22%22%2C%22zbDataType%22%3A0%2C%22zbDelay%22%3A60%2C%22zbDescription%22%3A%22CPU%E4%BD%BF%E7%94%A8%E6%97%B6%E9%97%B4%22%2C%22zbInterfaceType%22%3A1%2C%22zbItemId%22%3A1%2C%22zbItmeFlag%22%3A1%2C%22zbKey%22%3A%22system.cpu.util%5B%2Cuser%5D%22%2C%22zbName%22%3A%22CPU%E4%BD%BF%E7%94%A8%E6%97%B6%E9%97%B4+%5BCPU+user+util++time%5D%22%2C%22zbTempletItemId%22%3A35%2C%22zbTempletItemName%22%3A%22%22%2C%22zbType%22%3A0%2C%22zbUnits%22%3A%22%25%22%2C%22zbValueType%22%3A0%7D%2C%7B%22hxTaskId%22%3A0%2C%22itemId%22%3A0%2C%22taskDiscoveryRuleId%22%3A0%2C%22vId%22%3A13%2C%22zabbixHostId%22%3A0%2C%22zabbixItemId%22%3A0%2C%22zbApplicationId%22%3A1%2C%22zbApplicationName%22%3A%22%22%2C%22zbDataType%22%3A0%2C%22zbDelay%22%3A60%2C%22zbDescription%22%3A%22CPU%E9%9D%99%E9%BB%98%E6%97%B6%E9%97%B4%22%2C%22zbInterfaceType%22%3A1%2C%22zbItemId%22%3A3%2C%22zbItmeFlag%22%3A1%2C%22zbKey%22%3A%22system.cpu.util%5B%2Csteal%5D%22%2C%22zbName%22%3A%22CPU%E9%9D%99%E9%BB%98%E6%97%B6%E9%97%B4+%5BCPU+steal+time%5D%22%2C%22zbTempletItemId%22%3A37%2C%22zbTempletItemName%22%3A%22%22%2C%22zbType%22%3A0%2C%22zbUnits%22%3A%22%25%22%2C%22zbValueType%22%3A0%7D%2C%7B%22hxTaskId%22%3A0%2C%22itemId%22%3A0%2C%22taskDiscoveryRuleId%22%3A0%2C%22vId%22%3A15%2C%22zabbixHostId%22%3A0%2C%22zabbixItemId%22%3A0%2C%22zbApplicationId%22%3A1%2C%22zbApplicationName%22%3A%22%22%2C%22zbDataType%22%3A0%2C%22zbDelay%22%3A60%2C%22zbDescription%22%3A%22CPU%E6%AD%A3%E5%B8%B8%E4%BD%BF%E7%94%A8%E6%97%B6%E9%97%B4%22%2C%22zbInterfaceType%22%3A1%2C%22zbItemId%22%3A5%2C%22zbItmeFlag%22%3A1%2C%22zbKey%22%3A%22system.cpu.util%5B%2Cnice%5D%22%2C%22zbName%22%3A%22CPU%E6%AD%A3%E5%B8%B8%E4%BD%BF%E7%94%A8%E6%97%B6%E9%97%B4+%5BCPU+nice+time%5D%22%2C%22zbTempletItemId%22%3A39%2C%22zbTempletItemName%22%3A%22%22%2C%22zbType%22%3A0%2C%22zbUnits%22%3A%22%25%22%2C%22zbValueType%22%3A0%7D%2C%7B%22hxTaskId%22%3A0%2C%22itemId%22%3A0%2C%22taskDiscoveryRuleId%22%3A0%2C%22vId%22%3A17%2C%22zabbixHostId%22%3A0%2C%22zabbixItemId%22%3A0%2C%22zbApplicationId%22%3A1%2C%22zbApplicationName%22%3A%22%22%2C%22zbDataType%22%3A0%2C%22zbDelay%22%3A60%2C%22zbDescription%22%3A%22CPU%E4%B8%AD%E6%96%AD%E6%97%B6%E9%97%B4%22%2C%22zbInterfaceType%22%3A1%2C%22zbItemId%22%3A7%2C%22zbItmeFlag%22%3A1%2C%22zbKey%22%3A%22system.cpu.util%5B%2Cinterrupt%5D%22%2C%22zbName%22%3A%22CPU%E4%B8%AD%E6%96%AD%E6%97%B6%E9%97%B4+%5BCPU+interrupt+time%5D%22%2C%22zbTempletItemId%22%3A41%2C%22zbTempletItemName%22%3A%22%22%2C%22zbType%22%3A0%2C%22zbUnits%22%3A%22%25%22%2C%22zbValueType%22%3A0%7D%2C%7B%22hxTaskId%22%3A0%2C%22itemId%22%3A0%2C%22taskDiscoveryRuleId%22%3A0%2C%22vId%22%3A19%2C%22zabbixHostId%22%3A0%2C%22zabbixItemId%22%3A0%2C%22zbApplicationId%22%3A1%2C%22zbApplicationName%22%3A%22%22%2C%22zbDataType%22%3A0%2C%22zbDelay%22%3A60%2C%22zbDescription%22%3A%22CPU%E4%B8%8A%E4%B8%8B%E6%96%87%E5%88%87%E6%8D%A2%E9%80%9F%E5%BA%A6%22%2C%22zbInterfaceType%22%3A1%2C%22zbItemId%22%3A9%2C%22zbItmeFlag%22%3A1%2C%22zbKey%22%3A%22system.cpu.switches%22%2C%22zbName%22%3A%22CPU%E4%B8%8A%E4%B8%8B%E6%96%87%E5%88%87%E6%8D%A2%E9%80%9F%E5%BA%A6+%5BContext+switches+per+second%5D%22%2C%22zbTempletItemId%22%3A43%2C%22zbTempletItemName%22%3A%22%22%2C%22zbType%22%3A0%2C%22zbUnits%22%3A%22sps%22%2C%22zbValueType%22%3A3%7D%2C%7B%22hxTaskId%22%3A0%2C%22itemId%22%3A0%2C%22taskDiscoveryRuleId%22%3A0%2C%22vId%22%3A111%2C%22zabbixHostId%22%3A0%2C%22zabbixItemId%22%3A0%2C%22zbApplicationId%22%3A1%2C%22zbApplicationName%22%3A%22%22%2C%22zbDataType%22%3A0%2C%22zbDelay%22%3A60%2C%22zbDescription%22%3A%22CPU%E8%B4%9F%E8%BD%BD%28%E6%AF%8F%E8%8A%AF15%E5%88%86%E9%92%9F%E5%B9%B3%E5%9D%87%E5%80%BC%29%22%2C%22zbInterfaceType%22%3A1%2C%22zbItemId%22%3A11%2C%22zbItmeFlag%22%3A1%2C%22zbKey%22%3A%22system.cpu.load%5Bpercpu%2Cavg15%5D%22%2C%22zbName%22%3A%22CPU%E8%B4%9F%E8%BD%BD%28%E6%AF%8F%E8%8A%AF15%E5%88%86%E9%92%9F%E5%B9%B3%E5%9D%87%E5%80%BC%29%5BProcessor+load+%2815+min+average+per+core%29%5D%22%2C%22zbTempletItemId%22%3A45%2C%22zbTempletItemName%22%3A%22%22%2C%22zbType%22%3A0%2C%22zbUnits%22%3A%22%22%2C%22zbValueType%22%3A0%7D%2C%7B%22hxTaskId%22%3A0%2C%22itemId%22%3A0%2C%22taskDiscoveryRuleId%22%3A0%2C%22vId%22%3A113%2C%22zabbixHostId%22%3A0%2C%22zabbixItemId%22%3A0%2C%22zbApplicationId%22%3A1%2C%22zbApplicationName%22%3A%22%22%2C%22zbDataType%22%3A0%2C%22zbDelay%22%3A60%2C%22zbDescription%22%3A%22CPU%E8%B4%9F%E8%BD%BD%28%E6%AF%8F%E8%8A%AF%E6%AF%8F%E5%88%86%E9%92%9F%E5%B9%B3%E5%9D%87%E5%80%BC%29%22%2C%22zbInterfaceType%22%3A1%2C%22zbItemId%22%3A13%2C%22zbItmeFlag%22%3A1%2C%22zbKey%22%3A%22system.cpu.load%5Bpercpu%2Cavg1%5D%22%2C%22zbName%22%3A%22CPU%E8%B4%9F%E8%BD%BD%28%E6%AF%8F%E8%8A%AF%E6%AF%8F%E5%88%86%E9%92%9F%E5%B9%B3%E5%9D%87%E5%80%BC%29%5BProcessor+load+%281+min+average+per+core%29%5D%22%2C%22zbTempletItemId%22%3A47%2C%22zbTempletItemName%22%3A%22%22%2C%22zbType%22%3A0%2C%22zbUnits%22%3A%22%22%2C%22zbValueType%22%3A0%7D%2C%7B%22hxTaskId%22%3A0%2C%22itemId%22%3A0%2C%22taskDiscoveryRuleId%22%3A0%2C%22vId%22%3A12%2C%22zabbixHostId%22%3A0%2C%22zabbixItemId%22%3A0%2C%22zbApplicationId%22%3A1%2C%22zbApplicationName%22%3A%22%22%2C%22zbDataType%22%3A0%2C%22zbDelay%22%3A60%2C%22zbDescription%22%3A%22CPU%E7%B3%BB%E7%BB%9F%E6%97%B6%E9%97%B4%22%2C%22zbInterfaceType%22%3A1%2C%22zbItemId%22%3A2%2C%22zbItmeFlag%22%3A1%2C%22zbKey%22%3A%22system.cpu.util%5B%2Csystem%5D%22%2C%22zbName%22%3A%22CPU%E7%B3%BB%E7%BB%9F%E6%97%B6%E9%97%B4+%5BCPU+system+time%5D%22%2C%22zbTempletItemId%22%3A36%2C%22zbTempletItemName%22%3A%22%22%2C%22zbType%22%3A0%2C%22zbUnits%22%3A%22%25%22%2C%22zbValueType%22%3A0%7D%2C%7B%22hxTaskId%22%3A0%2C%22itemId%22%3A0%2C%22taskDiscoveryRuleId%22%3A0%2C%22vId%22%3A14%2C%22zabbixHostId%22%3A0%2C%22zabbixItemId%22%3A0%2C%22zbApplicationId%22%3A1%2C%22zbApplicationName%22%3A%22%22%2C%22zbDataType%22%3A0%2C%22zbDelay%22%3A60%2C%22zbDescription%22%3A%22CPU%E8%BD%AF%E4%BB%B6%E4%B8%AD%E6%96%AD%E6%97%B6%E9%97%B4%22%2C%22zbInterfaceType%22%3A1%2C%22zbItemId%22%3A4%2C%22zbItmeFlag%22%3A1%2C%22zbKey%22%3A%22system.cpu.util%5B%2Csoftirq%5D%22%2C%22zbName%22%3A%22CPU%E8%BD%AF%E4%BB%B6%E4%B8%AD%E6%96%AD%E6%97%B6%E9%97%B4+%5BCPU+softirq+time%5D%22%2C%22zbTempletItemId%22%3A38%2C%22zbTempletItemName%22%3A%22%22%2C%22zbType%22%3A0%2C%22zbUnits%22%3A%22%25%22%2C%22zbValueType%22%3A0%7D%2C%7B%22hxTaskId%22%3A0%2C%22itemId%22%3A0%2C%22taskDiscoveryRuleId%22%3A0%2C%22vId%22%3A16%2C%22zabbixHostId%22%3A0%2C%22zabbixItemId%22%3A0%2C%22zbApplicationId%22%3A1%2C%22zbApplicationName%22%3A%22%22%2C%22zbDataType%22%3A0%2C%22zbDelay%22%3A60%2C%22zbDescription%22%3A%22CPU+IO++%E7%AD%89%E5%BE%85%E6%97%B6%E9%97%B4%22%2C%22zbInterfaceType%22%3A1%2C%22zbItemId%22%3A6%2C%22zbItmeFlag%22%3A1%2C%22zbKey%22%3A%22system.cpu.util%5B%2Ciowait%5D%22%2C%22zbName%22%3A%22CPU+IO++%E7%AD%89%E5%BE%85%E6%97%B6%E9%97%B4+%5BCPU+iowait+time%5D%22%2C%22zbTempletItemId%22%3A40%2C%22zbTempletItemName%22%3A%22%22%2C%22zbType%22%3A0%2C%22zbUnits%22%3A%22%25%22%2C%22zbValueType%22%3A0%7D%2C%7B%22hxTaskId%22%3A0%2C%22itemId%22%3A0%2C%22taskDiscoveryRuleId%22%3A0%2C%22vId%22%3A18%2C%22zabbixHostId%22%3A0%2C%22zabbixItemId%22%3A0%2C%22zbApplicationId%22%3A1%2C%22zbApplicationName%22%3A%22%22%2C%22zbDataType%22%3A0%2C%22zbDelay%22%3A60%2C%22zbDescription%22%3A%22CPU%E7%A9%BA%E9%97%B2%E6%97%B6%E9%97%B4%22%2C%22zbInterfaceType%22%3A1%2C%22zbItemId%22%3A8%2C%22zbItmeFlag%22%3A1%2C%22zbKey%22%3A%22system.cpu.util%5B%2Cidle%5D%22%2C%22zbName%22%3A%22CPU%E7%A9%BA%E9%97%B2%E6%97%B6%E9%97%B4+%5BCPU+idle+time%5D%22%2C%22zbTempletItemId%22%3A42%2C%22zbTempletItemName%22%3A%22%22%2C%22zbType%22%3A0%2C%22zbUnits%22%3A%22%25%22%2C%22zbValueType%22%3A0%7D%2C%7B%22hxTaskId%22%3A0%2C%22itemId%22%3A0%2C%22taskDiscoveryRuleId%22%3A0%2C%22vId%22%3A110%2C%22zabbixHostId%22%3A0%2C%22zabbixItemId%22%3A0%2C%22zbApplicationId%22%3A1%2C%22zbApplicationName%22%3A%22%22%2C%22zbDataType%22%3A0%2C%22zbDelay%22%3A60%2C%22zbDescription%22%3A%22CPU%E4%B8%AD%E6%96%AD%E6%AC%A1%E6%95%B0%22%2C%22zbInterfaceType%22%3A1%2C%22zbItemId%22%3A10%2C%22zbItmeFlag%22%3A1%2C%22zbKey%22%3A%22system.cpu.intr%22%2C%22zbName%22%3A%22CPU%E4%B8%AD%E6%96%AD%E6%AC%A1%E6%95%B0+%5BInterrupts+per+second%5D%22%2C%22zbTempletItemId%22%3A44%2C%22zbTempletItemName%22%3A%22%22%2C%22zbType%22%3A0%2C%22zbUnits%22%3A%22ips%22%2C%22zbValueType%22%3A3%7D%2C%7B%22hxTaskId%22%3A0%2C%22itemId%22%3A0%2C%22taskDiscoveryRuleId%22%3A0%2C%22vId%22%3A112%2C%22zabbixHostId%22%3A0%2C%22zabbixItemId%22%3A0%2C%22zbApplicationId%22%3A1%2C%22zbApplicationName%22%3A%22%22%2C%22zbDataType%22%3A0%2C%22zbDelay%22%3A60%2C%22zbDescription%22%3A%22CPU%E8%B4%9F%E8%BD%BD%28%E6%AF%8F%E8%8A%AF5%E5%88%86%E9%92%9F%E5%B9%B3%E5%9D%87%E5%80%BC%29%22%2C%22zbInterfaceType%22%3A1%2C%22zbItemId%22%3A12%2C%22zbItmeFlag%22%3A1%2C%22zbKey%22%3A%22system.cpu.load%5Bpercpu%2Cavg5%5D%22%2C%22zbName%22%3A%22CPU%E8%B4%9F%E8%BD%BD%28%E6%AF%8F%E8%8A%AF5%E5%88%86%E9%92%9F%E5%B9%B3%E5%9D%87%E5%80%BC%29%5BProcessor+load+%285+min+average+per+core%29%5D%22%2C%22zbTempletItemId%22%3A46%2C%22zbTempletItemName%22%3A%22%22%2C%22zbType%22%3A0%2C%22zbUnits%22%3A%22%22%2C%22zbValueType%22%3A0%7D%5D%2C%22equInterfaceMap%22%3A%7B%2280%22%3A%5B%7B%22equId%22%3A%2280%22%2C%22interType%22%3A%221%22%2C%22dnsName%22%3A%22%22%2C%22isUseIp%22%3A1%2C%22portNo%22%3A%2210050%22%7D%5D%7D%7D"
	fmt.Println(url.PathUnescape(s))

}
