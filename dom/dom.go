package dom

// Package dom provides a small DOM library for TinyGo, enabling DOM manipulation
// in WebAssembly TinyGo applications.

import (
	"syscall/js"
)

// ----------------------------------------------------------------------------
var document js.Value
var body js.Value
var console js.Value

// ----------------------------------------------------------------------------
func init() {
	document = js.Global().Get("document")
	body = document.Get("body")
	console = js.Global().Get("console")
}

// ----------------------------------------------------------------------------
//
// Adds the specified element directly to the document body.
func AddToBody(element js.Value) {
	AddToElement(body, element)
}

// ----------------------------------------------------------------------------
//
// Log calls JavaScript console.log in the browser.
func Log(message string) {
	console.Call("log", message)
}
