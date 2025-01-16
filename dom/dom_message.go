package dom

import "syscall/js"

// ----------------------------------------------------------------------------
func ShowAlert(msg string) {
	js.Global().Call("alert", msg)
}
