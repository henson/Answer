package device

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os/exec"

	"github.com/henson/Answer/util"
)

//Android android
type Android struct{}

//NewAndroid new
func NewAndroid(cfg *util.Config) *Android {
	return new(Android)
}

//GetImage 通过adb获取截图
/* func (android *Android) GetImage() (img image.Image, err error) {
	err = exec.Command("adb", "shell", "screencap", "-p", "/sdcard/screenshot.png").Run()
	if err != nil {
		return
	}
	originImagePath := util.ImagePath + "origin.png"
	err = exec.Command("adb", "pull", "/sdcard/screenshot.png", originImagePath).Run()
	if err != nil {
		return
	}
	img, err = util.OpenPNG(originImagePath)
	return
} */

//GetImage 直接读取adb截图数据，速度更快
func (android *Android) GetImage() (img image.Image, err error) {
	cmd := exec.Command("adb", "shell", "screencap", "-p")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err = cmd.Run(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	x := bytes.Replace(out.Bytes(), []byte("\r\r\n"), []byte("\n"), -1)
	img, err = png.Decode(bytes.NewReader(x))
	return
}
