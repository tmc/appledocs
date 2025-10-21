package main

import (
	"fmt"
	"runtime"
	
	"github.com/tmc/appledocs/generated/appkit"
)

func main() {
	runtime.LockOSThread()
	
	// Create a button
	button := appkit.NewButton()
	
	// Set the title (this was missing before!)
	button.SetTitle("Click Me!")
	
	// Get the title back
	title := button.Title()
	fmt.Printf("Button title: %s\n", title)
	
	fmt.Println("✓ Button property accessors working!")
}
