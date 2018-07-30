package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 表结构
type Proxy struct {
	gorm.Model
	IP     string `gorm:"not null"`
	Port   string `gorm:"not null"`
	Scheme string `gorm:"default:'http';not null"`
	Speed  int64  `gorm:"default:-1;not null"`
}

func (px *Proxy) String() string {
	return fmt.Sprintf("%s://%s:%s", px.Scheme, px.IP, px.Port)
}

var db *gorm.DB

func init() {
	var err error
	// connect to database
	db, err = gorm.Open("mysql", "root:passwd@tcp(localhost:3306)/proxy?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Proxy{})
}

// 返回所有代理
func All() []Proxy {
	var pxs []Proxy
	db.Find(&pxs)
	return pxs
}

// 排名靠前的已验证代理
func Top(num int, scheme string) []Proxy {
	var pxs []Proxy
	db.Where("speed <> ? AND scheme = ?", "-1", scheme).Order("speed").Limit(num).Find(&pxs)
	return pxs
}

// 已经验证的代理
func Count() int {
	var count int
	db.Model(&Proxy{}).Where("speed <> -1").Count(&count)
	return count
}

func Del(px *Proxy) {
	db.Delete(px)
}

func Save(px *Proxy) {
	db.Save(px)
}

func IsExist(px *Proxy) bool {
	var count int
	db.Model(&Proxy{}).Unscoped().Where(px).Count(&count)
	if count == 0 {
		return false
	} else {
		return true
	}
}



