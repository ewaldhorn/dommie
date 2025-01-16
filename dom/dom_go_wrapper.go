package dom

import "syscall/js"

// ----------------------------------------------------------------------------
//
// Wraps a Go function inside a JavaScript function. Handy for when you need to
// run Go code in event listeners etc.
func wrapGoFunction(fn func()) func(js.Value, []js.Value) interface{} {
	return func(_ js.Value, _ []js.Value) interface{} {
		fn()
		return nil
	}
}
