// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StatusBarButton] class.
var StatusBarButtonClass objc.Class

func init() {
	StatusBarButtonClass = objc.GetClass("NSStatusBarButton")
}

type StatusBarButton struct {
	objc.ID
}

func StatusBarButtonFrom(ptr unsafe.Pointer) StatusBarButton {
	return StatusBarButton{
		ID: objc.ID(ptr),
	}
}




