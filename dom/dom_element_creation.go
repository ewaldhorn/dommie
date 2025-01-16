package dom

import "syscall/js"

// ----------------------------------------------------------------------------
func createElement(element string) js.Value {
	return document.Call("createElement", element)
}

// ----------------------------------------------------------------------------
func CreateDiv() js.Value {
	return createElement("div")
}

// ----------------------------------------------------------------------------
func CreateParagraph() js.Value {
	return createElement("p")
}

// ----------------------------------------------------------------------------
func CreateParagraphWithText(text string) js.Value {
	p := CreateParagraph()
	p.Set("innerText", text)
	return p
}
