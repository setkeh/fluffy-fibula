package main

import (
	"fmt"
	"strconv"
	"syscall"
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

  	iu := i

	d := iu / (24 * 60 * 60)
	ds := strconv.FormatInt(d, 10)
	h := iu/(60*60) - 24*d
	hs := strconv.FormatInt(h, 10)
	m := iu/60 - 24*60*d - 60*h
	ms := strconv.FormatInt(m, 10)
	s := iu - 24*60*60*d - 60*60*h - 60*m
	ss := strconv.FormatInt(s, 10)

	return ds  + " Days "  + hs + " Hours " + ms + " Minutes " + ss + " Seconds"
  }