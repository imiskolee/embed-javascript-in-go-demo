package main

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

type UserModel struct {
	gorm.Model
	ClientID string //SaaS平台的客户ID
	Name string
	Age  string
}

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) AfterCreate() error {
	log.Print("Starting call AfterCreate...")
	script := GetCustomEvent(u.TableName(),"AfterCreate",u.ClientID)
	if script == nil {
		//没有任何自定义脚本，直接忽略处理
		return nil
	}

	vm := NewVM(&Runtime{
		ClientID: u.ClientID,
	})
	_,_ = vm.Run(script.Content)
	_,err := vm.Call("main",nil,u)
	if err != nil {
		log.Print(err)
		return errors.New("Can not run script for clientID =" + u.ClientID)
	}
	return nil
}


