package main

import (
	"syscall/js"

	"github.com/ewaldhorn/dommie/dom"
)

// ----------------------------------------------------------------------------
const VERSION = "0.0.2c"
const NAME = "Dommie Demo"

// ----------------------------------------------------------------------------
func getVersionString() string {
	return NAME + " v" + VERSION
}

// ----------------------------------------------------------------------------
func setVersionCallback() {
	js.Global().Set("getVersion", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return getVersionString()
	}))
}

// ----------------------------------------------------------------------------
func createToggleVersionButton() {
	if isReady {
		controls := dom.GetElementById("controls")
		button := dom.CreateButton("Toggle Version")
		dom.AddElementTo(controls, button)
	}
}

// ----------------------------------------------------------------------------
func showVersion() {}

// ----------------------------------------------------------------------------
func hideVersion() {}
