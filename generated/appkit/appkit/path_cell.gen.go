// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PathCell] class.
var PathCellClass objc.Class

func init() {
	PathCellClass = objc.GetClass("NSPathCell")
}

type PathCell struct {
	objc.ID
}

func PathCellFrom(ptr unsafe.Pointer) PathCell {
	return PathCell{
		ID: objc.ID(ptr),
	}
}




