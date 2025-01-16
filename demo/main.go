package main

import (
	"github.com/ewaldhorn/dommie/dom"
)

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
	setRefreshCallback()
}

// ----------------------------------------------------------------------------
func setRefreshCallback() {
	dom.AddEventListener("refreshButton", "click", func() {
	})
}
