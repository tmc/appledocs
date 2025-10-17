// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FormCell] class.
var FormCellClass objc.Class

func init() {
	FormCellClass = objc.GetClass("NSFormCell")
}

type FormCell struct {
	objc.ID
}

func FormCellFrom(ptr unsafe.Pointer) FormCell {
	return FormCell{
		ID: objc.ID(ptr),
	}
}



