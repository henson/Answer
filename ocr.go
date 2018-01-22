package Answer

import (
	"github.com/henson/Answer/ocr"
)

//Ocr ocr 识别图片文字
type Ocr interface {
	GetText(imgPath string) (string, error)
}

func tesseractOCR() Ocr {
	return ocr.NewTesseract()
}

func baiduOCR() Ocr {
	return ocr.NewBaidu()
}
