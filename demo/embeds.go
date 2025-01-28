package main

import (
	_ "embed"
)

//go:embed styles/bodystyle.css
var bodyStyleFile string

//go:embed docs/dommie.txt
var dommieTextFile string
