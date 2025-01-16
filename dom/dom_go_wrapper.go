package dom

import "syscall/js"

// ----------------------------------------------------------------------------
func wrapGoFunction(fn func()) func(js.Value, []js.Value) interface{} {
	return func(_ js.Value, _ []js.Value) interface{} {
		fn()
		return nil
	}
}
