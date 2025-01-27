package main

import "math/rand"

var colours = []string{"red", " blue", " orange"}
var sizes = []string{"large","larger","xlarge"}

// ----------------------------------------------------------------------------
func randomColour() string {
	return colours[rand.Intn(len(colours))]
}

// ----------------------------------------------------------------------------
func randomSize()string{
	return sizes[rand.Intn(len(sizes))]
}
