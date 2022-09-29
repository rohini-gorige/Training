package logger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

type logApi interface {
	Close() error
	Rotate() error
	Write(p []byte) (n int, err error)
}
type Configlogs struct {
	Configlogs []*Config `json: configlogs`
}

type Config struct {
	Filename   string `json: "fileName"`
	MaxSize    int    `json: "maxSize"`
	MaxAge     int    `json: "maxAge"`
	MaxBackups int    `json: "maxBackups"`
	LocalTime  bool   `json: "localTime"`
	Compress   bool   `json: "compress"`
}

var l logApi

func NewLogService(path string, name string) (logApi, error) {

	os.MkdirAll(path, os.ModePerm)

	file, err := os.OpenFile(path+name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Println("error is comming form line 40")
		return nil, err

	}
	var maxAge int
	var maxBackups int
	var maxSize int
	var compress bool
	//var filename string
	var localTime bool
	fmt.Println("enter the number of variavbles and configuration ")
	fmt.Println("enter the max age in integer")
	fmt.Scanln(&maxAge)
	fmt.Println("enter the maxBackups in integer ")
	fmt.Scanln(&maxBackups)
	fmt.Println("enter the maxSixe integer")
	fmt.Scanln(&maxSize)
	fmt.Println("enter the Comperssion in boolean ")
	fmt.Scanln(&compress)
	fmt.Println("want to file format in local time  in boolean")
	//fmt.Scanln(&filename)
	fmt.Scanln(&localTime)
	data := Config{
		//Filename:   filename,
		MaxAge:     maxAge,
		MaxBackups: maxBackups,
		MaxSize:    maxSize,
		Compress:   compress,
		LocalTime:  localTime,
	}
	filenew, _ := json.MarshalIndent(data, "", "")
	_ = ioutil.WriteFile("C:/Users/SRS/workspace/telemetry-aop/logger/config.json", filenew, 0644)
	// confiPath, err3 := filepath.Abs("./config.json")
	// if err3 != nil {
	// 	log.Println("error at line 45", err3)
	// }
	// pathCon := confiPath
	// log.Println(pathCon)

	filel, err2 := ioutil.ReadFile("C:/Users/SRS/workspace/telemetry-aop/logger/config.json")
	if err2 != nil {
		log.Println("it is not reading the file ")
		return nil, err2
	}
	config := Configlogs{}

	err2 = json.Unmarshal(filel, &config)
	if err2 != nil {
		fmt.Println("error json unmarshalling ")
		fmt.Println(err2.Error())
	}
	for i := 0; i < len(config.Configlogs); i++ {
		l = &lumberjack.Logger{
			Filename:   file.Name(),                     //file.Name(),
			MaxSize:    config.Configlogs[i].MaxSize,    //data.Configlog[i].MaxSize,
			MaxAge:     config.Configlogs[i].MaxAge,     //  0,
			MaxBackups: config.Configlogs[i].MaxBackups, //5,
			LocalTime:  config.Configlogs[i].LocalTime,  // true,
			Compress:   config.Configlogs[i].Compress,   // true,
		}
	}
	log.Println(config)
	log.SetOutput(l)

	l.Rotate()

	return l, nil

}

func jsonMarshell() {

}
