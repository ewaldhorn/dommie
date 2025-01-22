package dom

import "strings"

// ----------------------------------------------------------------------------
//
// Creates the specified element type
func CreateElement(element string) HTMLElement {
	return document.Call("createElement", element)
}

// ----------------------------------------------------------------------------
//
// Creates and returns a new "div" element
func CreateDiv() HTMLElement {
	return CreateElement("div")
}

// ----------------------------------------------------------------------------
//
// Wraps an existing element in a new div of the type specified
func WrapElementWithNewDiv(element HTMLElement, divClasses ...string) HTMLElement {
	div := CreateDiv()
	div.Set("className", strings.Join(divClasses, " "))
	AddElementTo(div, element)

	return div
}

// ----------------------------------------------------------------------------
//
// Creates and returns a new "p" element
func CreateParagraph() HTMLElement {
	return CreateElement("p")
}

// ----------------------------------------------------------------------------
//
// Creates and returns a new "p" with the innerText set to the given text value
func CreateParagraphWithText(text string) HTMLElement {
	p := CreateParagraph()
	p.Set("innerText", text)
	return p
}

// ----------------------------------------------------------------------------
//
// Creates a "button" element
func CreateButton(text string) HTMLElement {
	button := CreateElement("button")
	button.Set("type", "button")
	button.Set("innerText", text)
	return button
}

// ----------------------------------------------------------------------------
func CreateImg(src string) HTMLElement {
	img := CreateElement("img")
	img.Set("src", src)
	return img
}
