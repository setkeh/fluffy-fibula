package main

import (
	//"fmt"
	"github.com/xconstruct/go-pushbullet"
	//"strconv"
	//"syscall"
)

func main() {

	pb := pushbullet.New("VMpvLPbc593WK3Hj9vshQxn7CIt5szCW")
	devs, err := pb.Devices()
	if err != nil {
		panic(err)
	}

	day := Data(Load())
	err = pb.PushNote(devs[0].Iden, "Uptime:", day)
	if err != nil {
		panic(err)
	}
}
