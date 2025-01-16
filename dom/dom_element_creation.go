package dom

import "syscall/js"

// ----------------------------------------------------------------------------
//
// Creates the specified element type
func createElement(element string) js.Value {
	return document.Call("createElement", element)
}

// ----------------------------------------------------------------------------
//
// Creates and returns a new "div" element
func CreateDiv() js.Value {
	return createElement("div")
}

// ----------------------------------------------------------------------------
//
// Creates and returns a new "p" element
func CreateParagraph() js.Value {
	return createElement("p")
}

// ----------------------------------------------------------------------------
//
// Creates and returns a new "p" with the innerText set to the given text value
func CreateParagraphWithText(text string) js.Value {
	p := CreateParagraph()
	p.Set("innerText", text)
	return p
}
