package main

import (
	"fmt"
	"syscall/js"

	"github.com/ewaldhorn/dommie/dom"
)

// ----------------------------------------------------------------------------
var isReady bool
var booCounter int
var applicationElement js.Value

// ------------------------------------------ Externally provided by JavaScript

//export bootstrap
func bootstrap()

// ----------------------------------------------------------------------------
func main() {
	setCallbacks()
	toggleElements()
	bootstrap()

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch
}

// ----------------------------------------------------------------------------
func toggleElements() {
	dom.Hide("loading")
	dom.Show("controls")
	dom.Show("information")
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

		if applicationElement.IsNull() {
			dom.ShowAlert("Unable to find container element.")
		} else {
			isReady = true
		}

		return nil
	}))
}

// ----------------------------------------------------------------------------
func setDoSomethingCallback() {
	dom.AddEventListener("doSomethingButton", "click", func() {
		if isReady {
			booCounter++
			p := dom.CreateParagraphWithText(fmt.Sprintf("Boo! (%d)", booCounter))
			dom.AddToElement(applicationElement, p)
		}
	})
}
