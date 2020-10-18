package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
)


func NewVM() *otto.Otto {
	vm :=  otto.New()
	return vm
}

func CallJavascriptFunction(vm *otto.Otto) string {
	script := `function hello(name){
		return "Hello " + name + ", Are you remember me?"
	}`
	//将代码编译到vm中，相当于在html中使用<script>标签
	if _,err := vm.Run(script); err != nil {
		panic("Can not run script")
	}
	//调用自定义的js函数，并且得到返回值
	val,err := vm.Call("hello",nil,"Misko Lee")
	if err != nil {
		panic("Can not run script")
	}
	return val.String()
}

func CallGolangFunction(vm *otto.Otto) string {

	 helloFunction := func(name string) string {
		 return "Hello " + name + ", Are you remember me?"
	 }
	 //将go函数加入到vm中
	 _ = vm.Set("hello",helloFunction)

	 //调用刚才加入的go函数
	val,err := vm.Call("hello",nil,"Misko Lee")
	if err != nil {
		panic("Can not run script")
	}
	return val.String()
}

func main() {
	vm := NewVM()
	{
		output := CallJavascriptFunction(vm)
		fmt.Printf("CallJavascriptFunction:%s\n", output)
	}
	{
		output := CallGolangFunction(vm)
		fmt.Printf("CallGolangFunction:%s\n", output)
	}

}