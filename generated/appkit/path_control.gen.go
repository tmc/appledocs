// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PathControl] class.
var PathControlClass objc.Class

func init() {
	PathControlClass = objc.GetClass("NSPathControl")
}

type PathControl struct {
	objc.ID
}

func PathControlFrom(ptr unsafe.Pointer) PathControl {
	return PathControl{
		ID: objc.ID(ptr),
	}
}



