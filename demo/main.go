package main

import (
	"fmt"
	"syscall/js"

	"github.com/ewaldhorn/dommie/dom"
)

// ----------------------------------------------------------------------------
var booCounter int
var applicationElement js.Value

// ------------------------------------------ Externally provided by JavaScript

//export bootstrap
func bootstrap()

// ----------------------------------------------------------------------------
func main() {
	setCallbacks()
	dom.Hide("loading")
	dom.Show("controls")
	dom.Show("information")
	bootstrap()

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch
}

// ----------------------------------------------------------------------------
func setCallbacks() {
	setVersionCallback()
	setApplicationContainerCallback()
	setDoSomethingCallback()
}

// ----------------------------------------------------------------------------
func setApplicationContainerCallback() {
	js.Global().Set("setApplicationContainer", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		applicationElement = dom.GetElementById(args[0].String())
		// TODO: handle case where element is not found
		return nil
	}))
}

// ----------------------------------------------------------------------------
func setDoSomethingCallback() {
	dom.AddEventListener("doSomethingButton", "click", func() {
		booCounter++
		p := dom.CreateParagraphWithText(fmt.Sprintf("Boo! (%d)", booCounter))
		dom.AddToElement(applicationElement, p)
	})
}
