package main

import "github.com/robertkrimen/otto"

type Runtime struct {
	ClientID string
	Env string
}

func NewVM(runtime *Runtime) *otto.Otto {
	vm :=  otto.New()
	vm.Set("Runtime",runtime) //将当前运行时的信息嵌入到vm中
	return vm
}


