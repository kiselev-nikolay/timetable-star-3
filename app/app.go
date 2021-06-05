package main

import (
	_ "embed"
	"syscall/js"

	"github.com/kiselev-nikolay/timetable-star-3/pkg/htmlgen"
)

//go:embed index.html
var index string

var state *htmlgen.State

func main() {
	state = htmlgen.New(index)
	welcome := &htmlgen.Notification{
		Key:  "test1",
		Text: "Hey, Hey",
	}
	state.Data["Welcome"] = welcome
	state.AddCallback("test1", func(action string) {
		runes := []rune(welcome.Text)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		welcome.Text = string(runes)
	})
	js.Global().Set("render", Render())
	js.Global().Set("update", Update())
	select {}
}

func Render() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return state.Render()
	})
}

func Update() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var key string = args[0].String()
		state.Update(key)
		js.Global().Call("apply")
		return nil
	})
}
