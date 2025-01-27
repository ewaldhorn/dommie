package dom

// Package dom provides a small DOM library for TinyGo, enabling DOM manipulation
// in WebAssembly TinyGo applications.

import (
	"syscall/js"
)

// ----------------------------------------------------------------------------
// Globals used by Dommie to speed things up a bit by not having to get
// references over and over.
var document HTMLElement
var body HTMLElement
var head HTMLElement
var console js.Value
var dommieStyleSheet js.Value

// ----------------------------------------------------------------------------
// Set up references to some basics Dommie uses a lot.
func init() {
	document = js.Global().Get("document")
	head = document.Get("head")
	body = document.Get("body")
	console = js.Global().Get("console")
}

// ----------------------------------------------------------------------------
//
// Adds the specified element directly to the document body.
func AddToBody(element HTMLElement) {
	AddElementTo(body, element)
}
