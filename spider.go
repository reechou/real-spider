package main

import (
	"fmt"
	"os"
	"time"
	//"path/filepath"
	//"strings"
	"regexp"
	
	"github.com/PuerkitoBio/goquery"
)

var (
	anjukeHosts = []string{
		"quzhou",
		"jiaxing",
		"beijing",
		"shanghai",
		"guangzhou",
		"shenzhen",
		"chengdu",
		"nanjing",
		"tianjin",
		"hangzhou",
		"suzhou",
		"chongqing",
		"wuhan",
		"zhengzhou",
		"changsha",
		"changzhou",
		"changchun",
		"dalian",
		"dg",
		"foshan",
		"fuzhou",
		"guiyang",
		"hefei",
		"heb",
		"haikou",
		"huizhou",
		"jinan",
		"jinhua",
		"lanzhou",
		"luoyang",
		"mianyang",
		"nanjing",
		"ningbo",
		"nanchang",
		"nanning",
		"qingdao",
		"quanzhou",
		"sjz",
		"shenyang",
		"hsanya",
		"shaoxing",
		"shantou",
		"taiyuan",
		"tanshan",
		"wuxi",
		"weihai",
		"wenzhou",
		"xa",
		"xiamen",
		"xuzhou",
		"yantai",
		"yangzhou",
		"zhuhai",
		"zhongshan",
	}
)

func SpiderAnjuke() {
	for _, v := range anjukeHosts {
		host := fmt.Sprintf("https://%s.anjuke.com/tycoon/p1/", v)
		phoneMap := make(map[string]int)
		var valid = regexp.MustCompile("[0-9]+")
		for {
			fmt.Println("start to spider:", host)
			doc, err := goquery.NewDocument(host)
			if err != nil {
				fmt.Println("goquery new document error:", err)
				return
			}
			doc.Find(".jjr-side").Each(func(i int, s *goquery.Selection) {
				val := valid.FindStringSubmatch(s.Text())
				if len(val) > 0 {
					phoneMap[val[0]] = 1
				}
			})
			nextUrl, _ := doc.Find(".aNxt").Attr("href")
			if nextUrl == "" {
				fmt.Printf("[%s] is the end, num: %d\n", v, len(phoneMap))
				break
			} else {
				host = nextUrl
			}

			time.Sleep(17 * time.Second)
		}
		f, _ := os.OpenFile("anjuke_"+v, os.O_CREATE|os.O_RDWR, 0666)
		for k, _ := range phoneMap {
			f.WriteString(k+"\n")
		}
		f.Close()
		time.Sleep(17 * time.Second)
	}
}

func SpiderBaixing(cate string) {
	host := "http://hangzhou.baixing.com"
	i := 1
	phoneMap := make(map[string]int)
	for {
		if i > 100 {
			break
		}
		originUrl := fmt.Sprintf("%s/%s/?page=%d", host, cate, i)
		doc, err := goquery.NewDocument(originUrl)
		if err != nil {
			fmt.Println("goquery new document error:", err)
			return
		}
		doc.Find(".contact-button").Each(func(i int, s *goquery.Selection) {
			val, ok := s.Attr("data-contact")
			if ok {
				phoneMap[val] = 1
			}
		})
		time.Sleep(10 * time.Second)
		
		i++
	}
	f, _ := os.OpenFile(cate, os.O_CREATE|os.O_RDWR, 0666)
	defer f.Close()
	for k, _ := range phoneMap {
		f.WriteString(k+"\n")
	}
}

func main() {
	//SpiderBaixing("kuaijifuwu")
	//SpiderBaixing("jiaoyupeixun")
	
	SpiderAnjuke()
}
