// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TokenFieldCell] class.
var tokenFieldCellClass = _TokenFieldCellClass{objc.GetClass("NSTokenFieldCell")}

type _TokenFieldCellClass struct {
	class objc.Class
}

// A text field cell subclass that enables tokenized editing of an array of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell

type TokenFieldCell struct {
	TextFieldCell
}

// TokenFieldCellFrom constructs a [TokenFieldCell] from an unsafe.Pointer.
//
// A text field cell subclass that enables tokenized editing of an array of objects.
func TokenFieldCellFrom(ptr unsafe.Pointer) TokenFieldCell {
	return TokenFieldCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}



