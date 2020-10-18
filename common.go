package embed_javascript_in_go_demo

import "github.com/robertkrimen/otto"

//执行一段脚本
func ExecScript(vm *otto.Otto,script string,receiver func(returnVal otto.Value)) error {
	return nil
}
