package main

import "strings"

func IsValidName(name string) bool {
	return strings.TrimSpace(name) != ""
}
