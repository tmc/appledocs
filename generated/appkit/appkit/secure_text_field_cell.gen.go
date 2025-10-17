// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SecureTextFieldCell] class.
var SecureTextFieldCellClass objc.Class

func init() {
	SecureTextFieldCellClass = objc.GetClass("NSSecureTextFieldCell")
}

type SecureTextFieldCell struct {
	objc.ID
}

func SecureTextFieldCellFrom(ptr unsafe.Pointer) SecureTextFieldCell {
	return SecureTextFieldCell{
		ID: objc.ID(ptr),
	}
}




