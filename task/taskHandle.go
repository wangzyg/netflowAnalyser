package task

import (
	"os"
	"log"
	"bytes"
	"regexp"
	"encoding/json"
	"fmt"
)


var Tasks AllTask

/**
任务结构
 */
type AllTask struct{
	Task []Task `json:Task`
}
type Task struct {
	TaskName string `json:TaskName`
	RouterFilter RouterFilter `json:RouterFilter`
	SourceFilter SourceFilter `json:SourceFilter`
	DestFilter DestFilter `json:DestFilter`
}

type RouterFilter struct{
	RouterAddress []string `json:RouterAddress`
}

type SourceFilter struct{
	SourceAddress []string `json:SourceAddress`
}

type DestFilter struct{
	DestAddress []string `json:DestAddress`
}

func LoadConfig(taskJson string)(config AllTask, err error){
	path := taskJson
	config_file, err := os.Open(path)
	if err != nil {
		emit("Failed to open config file '%s': %s\n", path, err)
		return
	}

	fi, _ := config_file.Stat()
	//if size := fi.Size(); size > (configFileSizeLimit) {
	//	emit("config file (%q) size exceeds reasonable limit (%d) - aborting", path, size)
	//}

	if fi.Size() == 0 {
		emit("config file (%q) is empty, skipping", path)
	}

	buffer := make([]byte, fi.Size())
	_, err = config_file.Read(buffer)

	buffer, err = StripComments(buffer) //去掉注释
	if err != nil {
		emit("Failed to strip comments from json: %s\n", err)
		return
	}

	buffer = []byte(os.ExpandEnv(string(buffer))) //特殊

	err = json.Unmarshal(buffer, &config) //解析json格式数据
	if err != nil {
		emit("Failed unmarshalling json: %s\n", err)
		return
	}

	for _, v := range config.Task {
		fmt.Println("json:", v.TaskName)
		fmt.Println(v.RouterFilter.RouterAddress)
		fmt.Println(v.SourceFilter.SourceAddress)
		fmt.Println(v.DestFilter.DestAddress)

		if v.RouterFilter.RouterAddress == nil {
			fmt.Println("-------------------------------")
		}
	}


	return
}


func StripComments(data []byte) ([]byte, error) {
	data = bytes.Replace(data, []byte("\r"), []byte(""), 0) // Windows
	lines := bytes.Split(data, []byte("\n"))                //split to muli lines
	filtered := make([][]byte, 0)

	for _, line := range lines {
		match, err := regexp.Match(`^\s*#`, line)
		if err != nil {
			return nil, err
		}
		if !match {
			filtered = append(filtered, line)
		}
	}

	return bytes.Join(filtered, []byte("\n")), nil
}

func emit(msgfmt string, args ...interface{}) {
	log.Printf(msgfmt, args...)
}