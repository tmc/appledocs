// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TokenFieldCell] class.
var TokenFieldCellClass objc.Class

func init() {
	TokenFieldCellClass = objc.GetClass("NSTokenFieldCell")
}

type TokenFieldCell struct {
	objc.ID
}

func TokenFieldCellFrom(ptr unsafe.Pointer) TokenFieldCell {
	return TokenFieldCell{
		ID: objc.ID(ptr),
	}
}



