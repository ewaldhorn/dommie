package dom

import "syscall/js"

// ----------------------------------------------------------------------------
//
// Wraps a Go function and sets it as the callback for the specified event on
// the target element
func AddEventListener(elem string, event string, fn func()) {
	GetElementById(elem).Call("addEventListener", event, js.FuncOf(wrapGoFunction(fn)))
}
