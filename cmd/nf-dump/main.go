/*
Command nf-dump decodes NetFlow packets from UDP datagrams.

Usage:
		nf-dump [flags]

Flags:
		-addr string 	Listen address (default ":2055")
*/
package main

import (
	"bytes"
	"flag"
	"log"
	"net"
	"time"
	"runtime"

	"github.com/tehmaze/netflow"
	"github.com/tehmaze/netflow/ipfix"
	"github.com/tehmaze/netflow/netflow1"
	"github.com/tehmaze/netflow/netflow5"
	"github.com/tehmaze/netflow/netflow6"
	"github.com/tehmaze/netflow/netflow7"
	"github.com/tehmaze/netflow/netflow9"
	"github.com/tehmaze/netflow/session"
	"github.com/tehmaze/netflow/packetCached"
	"github.com/tehmaze/netflow/qqwry"
	"github.com/larspensjo/config"
	"strconv"
	"github.com/tehmaze/netflow/task"
)

// Safe default
var readSize = 2 << 16
var decoders map[string]*netflow.Decoder
var analyser_thread_num int = 1

var (
	configFile = flag.String("configfile", "config.ini", "General configuration file")

	//topic list
	collect_params = make(map[string]string)
)

func main() {
	chan_collect := make(chan bool)

	Init()

	go collect_loop(chan_collect)

	for i:=0; i<analyser_thread_num; i++ {
		go analyser_loop()
	}

	<- chan_collect
}

/**
系统初始化
 */
func Init(){
	runtime.GOMAXPROCS(runtime.NumCPU())

	loadConfigIni()

	loadTask()

	//初始化缓存队列
	packetCached.QueueInit()
	//初始化路由列表
	decoders = make(map[string]*netflow.Decoder)

	//初始化qqwry库
		datFile := "./qqwry.dat"
		qqwry.IPData.FilePath = datFile
		res := qqwry.IPData.InitIPData()
		if v, ok := res.(error); ok {
			log.Panic(v)
	}

}


func loadTask(){
	task.LoadConfig()
}

/**
加载配置文件
 */
func loadConfigIni(){
	//set config file std
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	//set config file std End

	//Initialized topic from the configuration
	if cfg.HasSection("collect") {
		section, err := cfg.SectionOptions("collect")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("collect", v)
				if err == nil {
					collect_params[v] = options
				}
			}
		}
	}
	//Initialized topic from the configuration END
	analyser_thread_num, _ = strconv.Atoi(collect_params["analyser_thread_num"])
	packetCached.Max_queue_length, _ = strconv.Atoi(collect_params["max_queue_length"])

}

/**
采集线程
 */
func collect_loop(chan_collect chan bool){
	listen := flag.String("addr", "222.73.216.10:9900", "Listen address")
	flag.Parse()

	var addr *net.UDPAddr
	var err error
	if addr, err = net.ResolveUDPAddr("udp", *listen); err != nil {
		log.Fatal(err)
	}

	var server *net.UDPConn
	if server, err = net.ListenUDP("udp", addr); err != nil {
		log.Fatal(err)
	}

	if err = server.SetReadBuffer(readSize); err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 8192)
		var remote *net.UDPAddr
		var octets int
		if octets, remote, err = server.ReadFromUDP(buf); err != nil {
			log.Printf("error reading from %s: %v\n", remote, err)
			continue
		}

		//log.Printf("received %d bytes from %s\n", octets, remote)

		//fmt.Println("received data address:buf:", &buf[0], " octets:", &octets, " remote:", &remote)
		packetCached.Queue.PushQueue(buf, octets, remote)


	}

	chan_collect <- true
}

/**
解析netflow包线程
 */
func analyser_loop(){
	//每个线程分配一个ip地址查询器
	qqWry := qqwry.NewQQwry()

	for{

		//data, err := packetCached.GetQueueData()
		data, err := packetCached.Queue.GetQueueData()
		if err{
			//fmt.Println("pop data_queue size:", data_queue.Size())
			//var buf []byte = ([]byte)data
			processPacket(data.(packetCached.UDPPacket), qqWry)
		} else{
			time.Sleep(500)
		}
	}
}

/**
处理包
 */
func processPacket(p packetCached.UDPPacket, qqWry qqwry.QQwry){
	d, found := decoders[p.RouterAddr.String()]
	if !found {
		s := session.New()
		d = netflow.NewDecoder(s)
		decoders[p.RouterAddr.String()] = d
	}

	go makeNetFlow(d, p.Buf, p.Octets, qqWry)
}

/**
构造netflow对象
 */
func makeNetFlow(d *netflow.Decoder, buf []byte, octets int, qqWry qqwry.QQwry){

	m, err := d.Read(bytes.NewBuffer(buf[:octets]))
	if err != nil {
		log.Println("decoder error:", err)
		return
	}

	switch p := m.(type) {
	case *netflow1.Packet:
		netflow1.Dump(p)

	case *netflow5.Packet:
		nf_packet := m.(*netflow5.Packet)
		for _, record := range nf_packet.Records {
			record.IpAddrArea = qqWry.Find(record.SrcAddr.String())
		}
		netflow5.Dump(p)

	case *netflow6.Packet:
		netflow6.Dump(p)

	case *netflow7.Packet:
		netflow7.Dump(p)

	case *netflow9.Packet:
		netflow9.Dump(p)

	case *ipfix.Message:
		ipfix.Dump(p)
	}

}