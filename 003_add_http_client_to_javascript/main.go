package main

import "github.com/robertkrimen/otto"

func NewVM() *otto.Otto {
	vm :=  otto.New()
	vm.Set("NewHttpClient",NewHttpClient)
	return vm
}

func main() {
	vm := NewVM()
	script := `
	var httpClient = NewHttpClient();
	var respBody = httpClient.Get("https://www.baidu.com");
	console.log(respBody);
`
	if _,err := vm.Run(script); err != nil {
		panic("Can not run script")
	}
}