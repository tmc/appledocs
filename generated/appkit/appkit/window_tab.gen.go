// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WindowTab] class.
var WindowTabClass objc.Class

func init() {
	WindowTabClass = objc.GetClass("NSWindowTab")
}

type WindowTab struct {
	objc.ID
}

func WindowTabFrom(ptr unsafe.Pointer) WindowTab {
	return WindowTab{
		ID: objc.ID(ptr),
	}
}



