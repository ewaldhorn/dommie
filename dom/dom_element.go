package dom

import "syscall/js"

// ----------------------------------------------------------------------------
type HTMLElement = js.Value

// ----------------------------------------------------------------------------
//
// Adds the specified element to the target element
func AddElementTo(target, elem HTMLElement) {
	target.Call("appendChild", elem)
}

// ----------------------------------------------------------------------------
//
// Attempt to find an element with the given Id
func GetElementById(elem string) HTMLElement {
	return document.Call("getElementById", elem)
}

// ----------------------------------------------------------------------------
//
// Gets the specified value from the target element
func GetElementValue(elem string, value string) HTMLElement {
	return GetElementById(elem).Get(value)
}

// ----------------------------------------------------------------------------
//
// Toggles display to "none" on the specified element
func Hide(elem string) {
	GetElementValue(elem, "style").Call("setProperty", "display", "none")
}

// ----------------------------------------------------------------------------
//
// Toggles display to "block" on the specified element
func Show(elem string) {
	GetElementValue(elem, "style").Call("setProperty", "display", "block")
}

// ----------------------------------------------------------------------------
//
// Calls "focus()" on the target element
func SetFocus(elem string) {
	GetElementById(elem).Call("focus")
}

// ----------------------------------------------------------------------------
//
// Gets the specified value from the target element as a string
func GetString(elem string, value string) string {
	return GetElementValue(elem, value).String()
}

// ----------------------------------------------------------------------------
//
// Sets the specified key to the given value on the target element
func SetValue(elem string, key string, value string) {
	GetElementById(elem).Set(key, value)
}
