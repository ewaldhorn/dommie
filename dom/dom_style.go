package dom

import "strings"

// ----------------------------------------------------------------------------
//
// Checks whether the Dommie stylesheet exists yet or not. If not, it is created
// and added to the Head element.
func isDommieStylesheetAvailable() bool {
	if head.IsNull() {
		return false
	}

	if dommieStyleSheet.IsNull() {
		// not available yet, create it
		dommieStyleSheet = CreateElement("style")
		dommieStyleSheet.Set("type", "text/css")

		AddElementTo(head, dommieStyleSheet)
	}

	return true
}

// ----------------------------------------------------------------------------
//
// Adds a CSS class to the specified element
func AddClass(elem string, class string) {
	GetElementValue(elem, "classList").Call("add", class)
}

// ----------------------------------------------------------------------------
//
// Replaces all the classes for the element with the provided list
func ReplaceClasses(element HTMLElement, classes ...string) HTMLElement {
	element.Set("className", strings.Join(classes, " "))

	return element
}

// ----------------------------------------------------------------------------
//
// Removes the specified CSS class from the target element
func RemoveClass(elem string, class string) {
	classList := GetElementValue(elem, "classList")
	if classList.Call("contains", class).Bool() {
		classList.Call("remove", class)
	}
}

// ----------------------------------------------------------------------------
//
// Creates a new style element for the page
func AddNewStyleElement(style string) {
	if !head.IsNull() {
		tmpStyleSheet := CreateElement("style")
		tmpStyleSheet.Set("type", "text/css")
		tmpStyleSheet.Set("innerHTML", style)

		AddElementTo(head, tmpStyleSheet)
	}
}
