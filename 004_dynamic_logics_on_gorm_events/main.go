package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/robertkrimen/otto"
)

type Runtime struct {
	ClientID string
	Env string
}

func NewVM(runtime *Runtime) *otto.Otto {
	vm :=  otto.New()
	vm.Set("Runtime",runtime) //将当前运行时的信息嵌入到vm中
	return vm
}


func main() {
	db, err := gorm.Open("sqlite3","./004.db")
	 if err != nil {
	 	panic("Can not open SQLite for reason:" + err.Error())
	 }
	 db.AutoMigrate(&UserModel{})

	 //Create User for client A
	 {
	 	u := &UserModel{
	 		ClientID: "A",
	 		Name: "Misko Lee",
		}
		db.Create(u)
	 }
	//Create User for client B
	{
		u := &UserModel{
			ClientID: "B",
			Name: "Misko Chan",
		}
		db.Create(u)
	}

	//No custom script found
	{
		u := &UserModel{
			ClientID: "C",
			Name: "Linus",
		}
		db.Create(u)
	}
}



