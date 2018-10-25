package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)


func main() {
	db, err := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		fmt.Println("error opening DB "+err.Error())
	}

	type Total struct {
		Total int `gorm:"column:lele"`
	}

	total := Total{}
	db.Debug().Raw("select count(1) as lele from test").Scan(&total)
	fmt.Print(total.Total)
}