// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PathComponentCell] class.
var PathComponentCellClass objc.Class

func init() {
	PathComponentCellClass = objc.GetClass("NSPathComponentCell")
}

type PathComponentCell struct {
	objc.ID
}

func PathComponentCellFrom(ptr unsafe.Pointer) PathComponentCell {
	return PathComponentCell{
		ID: objc.ID(ptr),
	}
}




