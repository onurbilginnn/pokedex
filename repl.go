package main

import "strings"

func cleanInput(text string) []string {
	// strings.Fields trims whitespace and splits on any run of whitespace
	return strings.Fields(text)
}
