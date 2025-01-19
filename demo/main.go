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
	injectBodyCSS()
	createVersionElements()

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
	setApplicationContainerCallback()
	setVersionCallback()
	setToggleVersionCallback()
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
// Now that the page has loaded, inject some CSS styles to the body from the
// wasm side.
func injectBodyCSS() {
	dom.CreateNewStyleElement("body")
}

// ----------------------------------------------------------------------------
func createVersionDivAndContent() {

}

// ----------------------------------------------------------------------------
func createVersionElements() {
	createToggleVersionButton()
	createVersionDivAndContent()
}

// ----------------------------------------------------------------------------
func setToggleVersionCallback() {
	js.Global().Set("toggleDisplayVersion", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if isReady {
			toggleOn := args[0].Bool()
			if toggleOn {
				showVersion()
			} else {
				hideVersion()
			}
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
			dom.AddElementTo(applicationElement, p)
		}
	})
}
