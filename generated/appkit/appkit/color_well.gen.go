// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorWell] class.
var ColorWellClass objc.Class

func init() {
	ColorWellClass = objc.GetClass("NSColorWell")
}

type ColorWell struct {
	objc.ID
}

func ColorWellFrom(ptr unsafe.Pointer) ColorWell {
	return ColorWell{
		ID: objc.ID(ptr),
	}
}




