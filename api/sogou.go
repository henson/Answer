package api

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/henson/Answer/util"
	"github.com/parnurzeal/gorequest"
)

//Sogos 搜狗json解析类型
type Sogos struct {
	Code   int     `json:"code"`
	Result []Reply `json:"result"`
}

//Reply 类型
type Reply struct {
	Answers     []string     `json:"answers"`
	CDID        string       `json:"cd_id"`
	Channel     string       `json:"channel"`
	Choices     string       `json:"choices"`
	Debug       string       `json:"debug"`
	Recommend   string       `json:"recommend"`
	Result      string       `json:"result"`
	Searchinfos []Searchinfo `json:"search_infos"`
	Title       string       `json:"title"`
	UID         string       `json:"uid"`
}

//Searchinfo 类型
type Searchinfo struct {
	Summary string `json:"summary"`
	Title   string `json:"title"`
	URL     string `json:"url"`
}

//Sogou 获取搜狗的结果
func Sogou(app string) {
	c := util.GetCache()
	c.Set(util.QuestionInCache, "first init", time.Duration(60)*time.Second)
	for {
		_, body, _ := gorequest.New().Get("http://140.143.49.31/api/ans2?key="+app).
			Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_2_2 like Mac OS X) AppleWebKit/604.4.7 (KHTML, like Gecko) Mobile/15C202 Sogousearch/Ios/5.9.8").
			Set("Referer", "http://nb.sa.sogou.com/").End()
		body = strings.Replace(body, "undefined(", "", -1)
		body = strings.Replace(body, ")", "", -1)
		body = strings.Replace(body, "\"{", "{", -1)
		body = strings.Replace(body, "}\"", "}", -1)
		body = strings.Replace(body, "\\", "", -1)
		var x Sogos
		json.Unmarshal([]byte(body), &x)
		if len(x.Result) > 1 {
			QInCache, found := c.Get(util.QuestionInCache)
			if found {
				if x.Result[1].Title != QInCache.(string) {
					fmt.Println("=====================搜狗结果===================")
					fmt.Println(x.Result[1].Title)
					fmt.Println(x.Result[1].Answers)
					fmt.Println("\n搜狗推荐答案是【", x.Result[1].Result, "】")
					c.Set(util.QuestionInCache, x.Result[1].Title, time.Duration(60)*time.Second)
				}
			}
		}

		time.Sleep(time.Duration(1) * time.Second)
	}
}
