// main.go - Minimal Swift UI launcher

package main

import (
	"log"

	"github.com/ebitengine/purego"
)

func main() {
	// Load Swift UI library
	lib, err := purego.Dlopen("./libCoreGraphicsUI.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		log.Fatal(err)
	}

	// Register and run UI
	var uiRun func()
	purego.RegisterLibFunc(&uiRun, lib, "ui_run")

	uiRun() // Blocks until window closes
}
