package wasm

import "syscall/js"

// Append ...
func Append(parent, child js.Value) {
	parent.Call("appendChild", child)
}

// Create ...
func Create(tag string) js.Value {
	return js.Global().Get("document").Call("createElement", tag)
}

// GetById ...
func GetById(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

// Register ...
func Register(name string, fn func(args []js.Value)) {
	js.Global().Set(name, js.NewCallback(fn))
}