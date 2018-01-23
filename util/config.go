package util

import (
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v1"
)

//Config 全局配置
type Config struct {
	Debug          bool   `yaml:"debug"`
	APP            string `yaml:"app"`
	Auto           bool   `yaml:"automatic"`
	Device         string `yaml:"device"`
	OCR            string `yaml:"ocr"`
	AdbAddress     string `yaml:"adb_address"`
	WdaAddress     string `yaml:"wda_address"`
	BaiduAPIKey    string `yaml:"Baidu_API_Key"`
	BaiduSecretKey string `yaml:"Baidu_Secret_Key"`

	//西瓜视频截图题目位置
	XgQx int `yaml:"xg_q_x"`
	XgQy int `yaml:"xg_q_y"`
	XgQw int `yaml:"xg_q_w"`
	XgQh int `yaml:"xg_q_h"`
	//西瓜视频截取答案位置
	XgAx int `yaml:"xg_a_x"`
	XgAy int `yaml:"xg_a_y"`
	XgAw int `yaml:"xg_a_w"`
	XgAh int `yaml:"xg_a_h"`

	//冲顶大会截图题目位置
	CdQx int `yaml:"cd_q_x"`
	CdQy int `yaml:"cd_q_y"`
	CdQw int `yaml:"cd_q_w"`
	CdQh int `yaml:"cd_q_h"`
	//冲顶大会截取答案位置
	CdAx int `yaml:"cd_a_x"`
	CdAy int `yaml:"cd_a_y"`
	CdAw int `yaml:"cd_a_w"`
	CdAh int `yaml:"cd_a_h"`

	//花椒直播截图题目位置
	HjQx int `yaml:"hj_q_x"`
	HjQy int `yaml:"hj_q_y"`
	HjQw int `yaml:"hj_q_w"`
	HjQh int `yaml:"hj_q_h"`
	//花椒直播截取答案位置
	HjAx int `yaml:"hj_a_x"`
	HjAy int `yaml:"hj_a_y"`
	HjAw int `yaml:"hj_a_w"`
	HjAh int `yaml:"hj_a_h"`

	//芝士超人截图题目位置
	ZsQx int `yaml:"zs_q_x"`
	ZsQy int `yaml:"zs_q_y"`
	ZsQw int `yaml:"zs_q_w"`
	ZsQh int `yaml:"zs_q_h"`
	//芝士超人截取答案位置
	ZsAx int `yaml:"zs_a_x"`
	ZsAy int `yaml:"zs_a_y"`
	ZsAw int `yaml:"zs_a_w"`
	ZsAh int `yaml:"zs_a_h"`
}

var cfg *Config

var cfgFilename = "./config.yml"

//SetConfigFile 设置配置文件地址
func SetConfigFile(path string) {
	cfgFilename = path
}

//GetConfig 解析配置
func GetConfig() *Config {
	if cfg != nil {
		return cfg
	}
	filename, _ := filepath.Abs(cfgFilename)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	var c *Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		panic(err)
	}
	cfg = c
	return cfg
}
