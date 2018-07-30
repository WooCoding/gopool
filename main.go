package main

import (
	"github.com/WooCoding/goreq"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	// 等待协程
	wg sync.WaitGroup
	// 三个通道：1.新代理 2.有效代理 3.删除代理
	chanNewProxy    = make(chan Proxy, 500)
	chanValidProxy  = make(chan Proxy, 500)
	chanFailedProxy = make(chan Proxy, 500)
)

func main() {
	defer db.Close()
	// 通道
	go func() {
		for {
			select {
			case px := <-chanFailedProxy:
				Del(&px)
			case px := <-chanValidProxy:
				Save(&px)
			case px := <-chanNewProxy:
				if !IsExist(&px) {
					db.Create(&px)
				}
			default:

			}
		}
	}()
	// 启动api
	go API()

	for {
		// 获取数据库代理，并验证
		pxs := All()
		for _, px := range pxs {
			wg.Add(1)
			go Validator(px)
		}
		wg.Wait()
		// 查询已经验证的代理数目
		log.Printf("alive proxies:%d", Count())

		// 抓取代理
		for _, pw := range webList {
			go Spider(pw)
		}
		<-time.After(3 * time.Minute)
	}
}

func Validator(px Proxy) {

	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			//log.Println(r)
			chanFailedProxy <- px
		}
	}()

	host := "www.taobao.com"
	URL := fmt.Sprintf("%s://%s", px.Scheme, host)

	opt := &goreq.Option{
		Header:  GetRandomHeader(),
		Proxy:   px.String(),
		Timeout: 30 * time.Second,
	}

	start := time.Now()
	res, err := goreq.Get(URL, opt)
	if err != nil {
		panic(fmt.Errorf("校验代理:%s, 返回错误:%s", px.String(), err))
	}

	if res.StatusCode == 200 {
		// 转换为毫秒
		speed := int64(time.Since(start) / time.Millisecond)
		if px.Speed == -1 {
			px.Speed = speed
		} else {
			px.Speed = (px.Speed + speed) / 2
		}
		//有效修改
		chanValidProxy <- px
	} else {
		//删除
		chanFailedProxy <- px
	}
}
