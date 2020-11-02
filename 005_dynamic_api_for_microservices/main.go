package main

import "github.com/robertkrimen/otto"

type User struct {
	ID int
	Name string
}

type ServiceUser struct {

}

func (service *ServiceUser) GetUserByID(id int) *User {
	return &User{
		ID:1,
		Name: "Misko Lee",
	}
}

type Order struct {
	ID int
	UID int
	Number string
	Subtotal float64
}
type ServiceOrder struct {

}

func (service *ServiceOrder) CreateOrder(uid int,subtotal float64) *Order {
	return &Order{
		ID : 1,
		UID: uid,
		Number: "20201111111111001111",
		Subtotal: subtotal,
	}
}

func NewVM() *otto.Otto {
	vm :=  otto.New()
	vm.Set("NewUserService",func()*ServiceUser{return &ServiceUser{}})
	vm.Set("NewOrderService",func()*ServiceOrder{return &ServiceOrder{}})
	return vm
}

func main() {

	script := `
		function main(req) {
			var userService = NewUserService();
			var orderService = NewOrderService();
			var user = userService.GetUserByID(req.UID);
			//杀熟与救济穷人
			var subTotal = 100.00;
			if(user.Name == "Misko Lee") {
				subTotal = 1000.00;
			}else if (user.Name == "Misko Chan") {
				subTotal = 1.00;
			}
			var order = orderService.CreateOrder(user.ID,subTotal);
			console.log(order.Subtotal);
		}
  `
	vm := NewVM()
	vm.Run(script)
	_,err := vm.Call("main",nil,
		//模拟HTTP数据
		map[string]interface{}{
		"UID" : 1,
	})
	if err != nil {
		panic("Can not call script:" + err.Error())
	}

}