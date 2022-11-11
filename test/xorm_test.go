package test

import (
	"bytes"
	"cloud-disk/models"
	"encoding/json"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "ubuntu:Ubuntu_123@tcp(localhost)/cloud-disk")

	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(data)
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(dst.String())
}
