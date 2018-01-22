package Answer

import (
	"github.com/henson/Answer/ocr"
	"github.com/henson/Answer/util"
)

//Ocr ocr 识别图片文字
type Ocr interface {
	GetText(imgPath string) (string, error)
}

func tesseractOCR() Ocr {
	return ocr.NewTesseract()
}

func baiduOCR(cfg *util.Config) Ocr {
	return ocr.NewBaidu(cfg)
}
