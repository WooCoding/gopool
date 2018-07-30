package main

import (
	"github.com/WooCoding/goreq"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"time"
)

func Spider(pw *ProxyWeb) {

	for _, url := range pw.URLs {

		go func(url string) {

			defer func() {
				if r := recover(); r != nil {
					//log.Println(r)
				}
			}()

			opt := &goreq.Option{
				Header:  GetRandomHeader(),
				Proxy:   GetRandomProxy(),
				Timeout: 1 * time.Minute,
			}

			res, err := goreq.Get(url, opt)
			if err != nil {
				panic(fmt.Errorf("请求:%s, 返回错误:%s", url, err))
			}
			//fmt.Println(url, res.StatusCode)
			if res.StatusCode == 200 {

				switch pw.parse {
				case "css":
					doc, err := goquery.NewDocumentFromResponse(res)
					if err != nil {
						panic(fmt.Errorf("css解析:%s, 返回错误:%s", url, err))
					}
					trs := doc.Find(pw.pattern)
					trs.Each(func(i int, sel *goquery.Selection) {
						ip := sel.Find(pw.pos.ip).Text()
						port := sel.Find(pw.pos.port).Text()
						scheme := sel.Find(pw.pos.scheme).Text()
						if px, ok := IsValid(ip, port, scheme); ok {
							// 爬到的代理放入通道
							chanNewProxy <- px
						}
					})
				case "json":
					bodyByte, err := ioutil.ReadAll(res.Body)
					if err != nil {
						panic(fmt.Errorf("json解析:%s, 返回错误:%s", url, err))
					}
					defer res.Body.Close()
					bodyStr := string(bodyByte)
					val := gjson.Get(bodyStr, pw.pattern)
					val.ForEach(func(key, value gjson.Result) bool {
						ip := value.Get(pw.pos.ip).String()
						port := value.Get(pw.pos.port).String()
						scheme := value.Get(pw.pos.scheme).String()
						if px, ok := IsValid(ip, port, scheme); ok {
							// 爬到的代理放入通道
							chanNewProxy <- px
						}
						return true
					})
				}
			}
		}(url)
	}
}
