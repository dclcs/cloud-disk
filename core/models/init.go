package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
)

var Engine = Init()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "ubuntu:Ubuntu_123@tcp(localhost)/cloud-disk")

	if err != nil {
		log.Printf("New Engine Failed: %v\n", err)
		return nil
	}

	return engine

}
