package dom

import "syscall/js"

// ----------------------------------------------------------------------------
func AddEventListener(elem string, event string, fn func()) {
	GetElementById(elem).Call("addEventListener", event, js.FuncOf(wrapGoFunction(fn)))
}
