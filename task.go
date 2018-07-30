package main

import "fmt"

// 代理网站结构
type ProxyWeb struct {
	URLs    []string
	parse   string
	pattern string
	pos     Position
}

type Position struct {
	ip     string
	port   string
	scheme string
}

// 代理网站列表
var webList = []*ProxyWeb{
	{
		URLs:    getURLs("http://www.kuaidaili.com/free/inha/%d/", 1, 2),
		parse:   "css",
		pattern: "tbody > tr",
		pos:     Position{"td:nth-child(1)", "td:nth-child(2)", "td:nth-child(4)"},
	},
	{
		URLs:    getURLs("http://www.ip181.com/daili/%d.html", 1, 2),
		parse:   "css",
		pattern: "table > tbody > tr ~ tr",
		pos:     Position{"td:nth-child(1)", "td:nth-child(2)", "td:nth-child(4)"},
	},
	{
		URLs:    getURLs("http://www.66ip.cn/areaindex_%d/1.html", 1, 34),
		parse:   "css",
		pattern: "#footer tbody > tr ~ tr",
		pos:     Position{"td:nth-child(1)", "td:nth-child(2)", ""},
	},
	{
		URLs:    getURLs("http://www.xicidaili.com/nn/%d", 1, 2),
		parse:   "css",
		pattern: "#ip_list > tbody > tr ~ tr",
		pos:     Position{"td:nth-child(2)", "td:nth-child(3)", "td:nth-child(6)"},
	},
	{
		URLs: []string{
			"http://www.data5u.com/free/gngn/index.shtml",
			"http://www.data5u.com/free/gwgn/index.shtml",
		},
		parse:   "css",
		pattern: "ul > li:nth-child(2) > ul.l2",
		pos:     Position{"span:nth-child(1) > li", "span:nth-child(2) > li", "span:nth-child(4) > li"},
	},
	{
		URLs:    getURLs("http://www.ip3366.net/?stype=1&page=%d", 1, 2),
		parse:   "css",
		pattern: "#list > table > tbody > tr",
		pos:     Position{"td:nth-child(1)", "td:nth-child(2)", "td:nth-child(4)"},
	},
	{
		URLs:    getURLs("http://www.nianshao.me/?page=%d", 1, 2),
		parse:   "css",
		pattern: "#main > div > div > table > tbody > tr",
		pos:     Position{"td:nth-child(1)", "td:nth-child(2)", "td:nth-child(5)"},
	},
	{
		URLs:    getURLs("http://www.nianshao.me/?stype=2&page=%d", 1, 2),
		parse:   "css",
		pattern: "#main > div > div > table > tbody > tr",
		pos:     Position{"td:nth-child(1)", "td:nth-child(2)", "td:nth-child(5)"},
	},
	{
		URLs:    []string{
			"http://www.mogumiao.com/proxy/free/listFreeIp",
			"http://www.mogumiao.com/proxy/api/freeIp?count=15",
		},
		parse:   "json",
		pattern: "msg",
		pos:     Position{"ip", "port", ""},
	},
	{
		URLs:    []string{"http://www.xdaili.cn/ipagent/freeip/getFreeIps?page=1&rows=10"},
		parse:   "json",
		pattern: "RESULT.rows",
		pos:     Position{"ip", "port", ""},
	},
	{
		URLs:    getURLs("http://www.kxdaili.com/dailiip/1/%d.html#ip", 1, 2),
		parse:   "css",
		pattern: "#nav_btn01 > div:nth-child(11) > table > tbody > tr",
		pos:     Position{"td:nth-child(1)", "td:nth-child(2)", "td:nth-child(4)"},
	},
	{
		URLs:    getURLs("http://www.httpsdaili.com/?stype=1&page=%d", 1, 2),
		parse:   "css",
		pattern: "#list > table > tbody > tr",
		pos:     Position{"td:nth-child(1)", "td:nth-child(2)", "td:nth-child(4)"},
	},
	{
		URLs:    []string{"http://www.iphai.com/free/ng"},
		parse:   "css",
		pattern: "body > div.container.main-container > div.table-responsive.module > table > tbody > tr ~ tr",
		pos:     Position{"td:nth-child(1)", "td:nth-child(2)", "td:nth-child(4)"},
	},
}

// 页码
func getURLs(URL string, start, end int) []string {
	var URLs []string
	for i := start; i <= end; i++ {
		URLs = append(URLs, fmt.Sprintf(URL, i))
	}
	return URLs
}
