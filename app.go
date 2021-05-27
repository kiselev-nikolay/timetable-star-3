package main

import "syscall/js"

func main() {
	js.Global().Set("render", Render())
	select {}
}

func Render() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		data := args[0].String()
		data = "!!" + data + "!!"
		return data
	})
}
