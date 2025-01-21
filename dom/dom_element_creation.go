package dom

import "syscall/js"

// ----------------------------------------------------------------------------
//
// Creates the specified element type
func CreateElement(element string) js.Value {
	return document.Call("createElement", element)
}

// ----------------------------------------------------------------------------
//
// Creates and returns a new "div" element
func CreateDiv() js.Value {
	return CreateElement("div")
}

// ----------------------------------------------------------------------------
//
// Creates and returns a new "p" element
func CreateParagraph() js.Value {
	return CreateElement("p")
}

// ----------------------------------------------------------------------------
//
// Creates and returns a new "p" with the innerText set to the given text value
func CreateParagraphWithText(text string) js.Value {
	p := CreateParagraph()
	p.Set("innerText", text)
	return p
}

// ----------------------------------------------------------------------------
//
// Creates a "button" element
func CreateButton(text string) js.Value {
	button := CreateElement("button")
	button.Set("type", "button")
	button.Set("innerText", text)
	return button
}

// ----------------------------------------------------------------------------
func CreateImg(src string) js.Value {
	img := CreateElement("img")
	img.Set("src", src)
	return img
}
