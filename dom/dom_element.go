package dom

import "syscall/js"

// ----------------------------------------------------------------------------
func AddToElement(target, elem js.Value) {
	target.Call("appendChild", elem)
}

// ----------------------------------------------------------------------------
func GetElementById(elem string) js.Value {
	return document.Call("getElementById", elem)
}

// ----------------------------------------------------------------------------
func GetElementValue(elem string, value string) js.Value {
	return GetElementById(elem).Get(value)
}

// ----------------------------------------------------------------------------
func Hide(elem string) {
	GetElementValue(elem, "style").Call("setProperty", "display", "none")
}

// ----------------------------------------------------------------------------
func Show(elem string) {
	GetElementValue(elem, "style").Call("setProperty", "display", "block")
}

// ----------------------------------------------------------------------------
func SetFocus(elem string) {
	GetElementById(elem).Call("focus")
}

// ----------------------------------------------------------------------------
func GetString(elem string, value string) string {
	return GetElementValue(elem, value).String()
}

// ----------------------------------------------------------------------------
func SetValue(elem string, key string, value string) {
	GetElementById(elem).Set(key, value)
}
