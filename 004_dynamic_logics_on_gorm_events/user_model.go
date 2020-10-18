package main

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model
	ClientID string //SaaS平台的客户ID
	Name string
	Age  string
}

func (u *UserModel) AfterCreated() {

}


