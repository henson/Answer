package Answer

import (
	"flag"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/henson/Answer/api"
	"github.com/henson/Answer/util"
	"github.com/ngaut/log"
	termbox "github.com/nsf/termbox-go"
)

var cfgFilename = flag.String("config", "config.yml", "配置文件路径")
var attentionWord = []string{"是错", "错误", "没有", "不是", "不能", "不对", "不属于", "不可以", "不正确", "不提供", "不包含", "不包括", "不存在", "不经过", "不可能", "不匹配"}

func init() {
	flag.Parse()
}

//Run 启动
func Run() {
	util.SetConfigFile(*cfgFilename)
	cfg := util.GetConfig()
	util.MkDirIfNotExist(util.ImagePath)

	termbox.Init()
	defer termbox.Close()

	if !cfg.Debug {
		log.SetLevel(log.LOG_LEVEL_INFO)
	}
	fmt.Println("APP：", cfg.APP, "\t调试：", cfg.Debug)
	fmt.Println("设备：", cfg.Device, "\t配置文件：", *cfgFilename)

	//区分ios 或android 获取图像
	screenshot := NewScreenshot(cfg)

	if cfg.Device == "android" {
		//连接android
		cmd := exec.Command("adb", "connect", cfg.AdbAddress)
		out, err := cmd.Output()
		if err != nil {
			log.Errorf(err.Error())
		}
		fmt.Println(string(out))
	}
	fmt.Println("\n请按空格键开始搜索答案...")

Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeySpace:
				answerQuestion(screenshot, cfg)
				fmt.Println("\n\n请按空格键开始搜索答案...")
			default:
				break Loop
			}
		}
	}

}

func answerQuestion(sc Screenshot, cfg *util.Config) {
	start := time.Now()
	var wig sync.WaitGroup
	var questionText string
	var answerArr []string
	var flag bool
	imgChan1 := make(chan string, 1)
	imgChan2 := make(chan string, 1)
	qchan1 := make(chan string, 1)
	qchan2 := make(chan string, 1)
	ocr := NewOcr(cfg)
	go func() {
		go func() {
			api.Sogou(cfg.APP)
		}()
		knowledge(<-qchan1)
	}()
	wig.Add(3)
	go func() {
		defer wig.Done()
		//qText, err := tesseractOCR().GetText(util.QuestionImage)
		qText, err := ocr.GetText(<-imgChan1)
		if err != nil {
			log.Errorf("识别题目失败，%v", err.Error())
			return
		}
		questionText = processQuestion(qText)
		qchan1 <- questionText
		qchan2 <- questionText
	}()
	go func() {
		defer wig.Done()
		//answerText, err := baiduOCR().GetText(util.AnswerImage)
		answerText, err := ocr.GetText(<-imgChan2)
		if err != nil {
			log.Errorf("识别答案失败，%v", err.Error())
			return
		}
		answerArr = processAnswer(cfg, answerText)
	}()
	go func() {
		defer wig.Done()
		k := <-qchan2
		for _, v := range attentionWord {
			if strings.Contains(k, v) {
				//fmt.Println("请注意题干：", util.SubString(k, util.UnicodeIndex(k, v), 3)+"...")
				fmt.Printf("请注意题干：%s...\n", v)
				flag = true
				break
			}
		}
	}()
	go func() {
		if !cfg.Debug {
			png, err := sc.GetImage()
			if err != nil {
				log.Errorf("获取截图失败，%v", err.Error())
				return
			}
			err = saveImage(png, cfg, imgChan1, imgChan2)
			if err != nil {
				log.Errorf("保存图片失败，%v", err.Error())
				return
			}
		} else {
			imgChan1 <- util.QuestionImage
			imgChan2 <- util.AnswerImage
		}
	}()
	wig.Wait()

	var input []*Answers
	var finalAnswer string
	//搜索答案并显示
	result := GetSearchResult(questionText, answerArr)
	for engine, answerResult := range result {
		fmt.Printf("=====================%s搜索===================\n", engine)
		fmt.Println(questionText + "\n")
		for key, val := range answerResult {
			fmt.Printf("%s：\t结果总数%d，\t答案出现频率：%d\n", answerArr[key], val.Sum, val.Freq)
			input = append(input, &Answers{answerArr[key], val.Sum, val.Freq})
		}
	}
	finalAnswer = processLogic(input, flag)
	fmt.Println("\n最终答案可能是【", finalAnswer, "】")
	if cfg.Debug {
		go func() {
			//x, y := PressArea(findSliceIndex(answerArr, finalAnswer), cfg)
			//Press(x, y)
			PressEcho(cfg)
		}()
	}
	fmt.Println("================================================")
	fmt.Printf("总耗时：%v", time.Now().Sub(start))
}

func findSliceIndex(s []string, k string) int {
	for i, v := range s {
		if v == k {
			return i
		}
	}
	return 0
}

func processQuestion(text string) string {
	text = strings.TrimSpace(strings.Replace(text, " ", "", -1))
	text = strings.Replace(text, "\r\n", "", -1)
	text = strings.Replace(text, "-", ".", -1)
	//去除编号
	re, _ := regexp.Compile("\\d{1,2}\\.")
	text = re.ReplaceAllString(text, "")
	return text
}

func processAnswer(cfg *util.Config, text string) []string {
	if cfg.OCR == "tesseract" {
		text = strings.TrimSpace(strings.Replace(text, " ", "", -1))
		text = strings.Replace(text, "\r", "", -1)
	}
	arr := strings.Split(text, "\n")
	//去除空白
	textArr := []string{}
	for _, val := range arr {
		if strings.TrimSpace(val) == "" {
			continue
		}
		textArr = append(textArr, val)
	}
	return textArr
}
