package dom

import "strings"

// ----------------------------------------------------------------------------
//
// Log calls JavaScript console.log in the browser.
func Log(message ...string) {
	console.Call("log", strings.Join(message, " "))
}
