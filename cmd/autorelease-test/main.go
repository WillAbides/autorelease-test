package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var buildVersion = "unknown"

//go:embed version.txt
var versionTxt string

func main() {
	fmt.Println("hello from version", buildVersion)
	fmt.Println("version.txt:", strings.TrimSpace(versionTxt))
}
