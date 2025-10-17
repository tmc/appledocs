// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Nib] class.
var NibClass objc.Class

func init() {
	NibClass = objc.GetClass("NSNib")
}

type Nib struct {
	objc.ID
}

func NibFrom(ptr unsafe.Pointer) Nib {
	return Nib{
		ID: objc.ID(ptr),
	}
}



