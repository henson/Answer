package Answer

import (
	"math/rand"
	"os/exec"
	"strconv"
	"time"

	"github.com/henson/Answer/device"
	"github.com/henson/Answer/util"
	"github.com/ngaut/log"
)

//PressArea 随机区域
func PressArea(choice int, cfg *util.Config) (x, y string) {
	var P, H, K int
	switch cfg.APP {
	case "xigua":
		P, H = cfg.XgAy, cfg.XgAh
	case "cddh":
		P, H = cfg.CdAy, cfg.CdAh
	case "huajiao":
		P, H = cfg.HjAy, cfg.HjAh
	case "zscr":
		P, H = cfg.ZsAy, cfg.ZsAh
	}
	K = 20
	y = strconv.Itoa(P + K + (((H-4*K)/3)+K)*choice + 40)
	x = strconv.Itoa(RandInt(40, 800))
	return
}

//PressEcho 生成选项区域图
func PressEcho(cfg *util.Config) {
	png, _ := device.NewAndroid(cfg).GetImage()
	img0, _ := util.CutImage(png, 30, 420, 840, 90)
	img1, _ := util.CutImage(png, 30, 520, 840, 90)
	img2, _ := util.CutImage(png, 30, 625, 840, 90)
	util.SavePNG("./images/0.png", img0)
	util.SavePNG("./images/1.png", img1)
	util.SavePNG("./images/2.png", img2)
}

//Press 点击屏幕
func Press(x, y string) {
	err := exec.Command("adb", "shell", "input", "tap", x, y).Run()
	if err != nil {
		log.Errorf(err.Error())
	}
}

//RandInt 随机数
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
