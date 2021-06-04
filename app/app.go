package main

import (
	"syscall/js"

	"github.com/kiselev-nikolay/timetable-star-3/app/htmlgen"
)

func main() {
	js.Global().Set("render", Render())
	select {}
}

func Render() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		s := htmlgen.New()
		return s.Render()
	})
}
