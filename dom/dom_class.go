package dom

// ----------------------------------------------------------------------------
//
// Adds a CSS class to the specified element
func AddClass(elem string, class string) {
	GetElementValue(elem, "classList").Call("add", class)
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
