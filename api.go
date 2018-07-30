package main

import (
	"net/http"
	"fmt"
	"log"
	"strconv"
	"encoding/json"
	"strings"
)

func API()  {
	http.HandleFunc("/get", getProxy)         //设置访问的路由
	http.HandleFunc("/del", delProxy)         //设置访问的路由
	err := http.ListenAndServe(":1047", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getProxy(w http.ResponseWriter, r *http.Request)  {

	var (
		pxs    []Proxy
		pxsStr []string
		scheme  string
	)

	r.ParseForm()
	_type := r.Form["type"]
	if len(_type) == 0 || (_type[0] != "http" && _type[0] != "https") {
		scheme = "http"
	} else {
		scheme = _type[0]
	}

	limit := r.Form["limit"]
	if len(limit) == 0{
		pxs = Top(1, scheme)
	} else {
		num, _ := strconv.Atoi(limit[0])
		pxs = Top(num, scheme)
	}
	// 创建数组
	for _, px := range pxs {
		pxsStr = append(pxsStr, px.String())
	}

	if len(pxsStr) == 0 {
		fmt.Fprintf(w, "[]")
	} else {
		// 格式化json
		data, _ := json.Marshal(pxsStr)
		fmt.Fprintf(w, "%s", data)
	}
}

func delProxy(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form["proxy"]) == 1 {
		netloc := r.Form["proxy"][0]
		host := strings.Split(netloc,":")
		ip := host[0]
		port := host[1]
		px := Proxy{
			IP:ip,
			Port:port,
		}
		Del(&px)
		fmt.Fprintln(w, "del successful!")
	}
}