// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StatusBar] class.
var StatusBarClass objc.Class

func init() {
	StatusBarClass = objc.GetClass("NSStatusBar")
}

type StatusBar struct {
	objc.ID
}

func StatusBarFrom(ptr unsafe.Pointer) StatusBar {
	return StatusBar{
		ID: objc.ID(ptr),
	}
}



