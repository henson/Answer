package Answer

import (
	"image"

	"github.com/henson/Answer/device"
	"github.com/henson/Answer/util"
)

//Screenshot 获取屏幕截图
type Screenshot interface {
	GetImage() (image.Image, error)
}

//NewScreenshot 根据手机系统区分
func NewScreenshot(cfg *util.Config) Screenshot {
	if cfg.Device == util.DeviceiOS {
		return device.NewIOS(cfg)
	}
	return device.NewAndroid(cfg)
}
