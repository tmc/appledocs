// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Appearance] class.
var AppearanceClass objc.Class

func init() {
	AppearanceClass = objc.GetClass("NSAppearance")
}

type Appearance struct {
	objc.ID
}

func AppearanceFrom(ptr unsafe.Pointer) Appearance {
	return Appearance{
		ID: objc.ID(ptr),
	}
}




