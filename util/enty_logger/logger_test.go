package enty_logger

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	logs()
}
func logs() {
	os.Setenv("LOG_LEVEL", "Debug")
	Init()
	Info("hello", "world", 1)
	Debug("hello", "world", 2)
	Trace("hello", "world", 3)
	Info("hello", "world", 4)
}
func threads() {
	go func() {
		gid := GetGid()
		fmt.Printf("child goruntine1 gid:%v \n", gid)
	}()
	go func() {
		gid := GetGid()
		fmt.Printf("child goruntine2 gid:%v \n", gid)
	}()
	go func() {
		gid := GetGid()
		fmt.Printf("child goruntine3 gid:%v \n", gid)
	}()
	go func() {
		gid := GetGid()
		fmt.Printf("child goruntine4 gid:%v \n", gid)
	}()
	go func() {
		gid := GetGid()
		fmt.Printf("child goruntine5 gid:%v \n", gid)
	}()
	gid := GetGid()
	fmt.Printf("main goruntine gid:%v \n", gid)
	time.Sleep(time.Second)
}
