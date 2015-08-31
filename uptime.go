package main

import (
	"fmt"
	//"strconv"
	"syscall"
	"time"
)

func Load() (int64) {
	info := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(info)
	if err != nil {
		fmt.Printf("%s\n", err)
	} 

	return info.Uptime
}

func Data(i int64) (string) {
  	q := time.Duration(i)
	d := fmt.Sprint(q * time.Second)

	return d
}