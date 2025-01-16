package dom

// ----------------------------------------------------------------------------
func AddClass(elem string, class string) {
	GetElementValue(elem, "classList").Call("add", class)
}

// ----------------------------------------------------------------------------
func RemoveClass(elem string, class string) {
	classList := GetElementValue(elem, "classList")
	if classList.Call("contains", class).Bool() {
		classList.Call("remove", class)
	}
}
