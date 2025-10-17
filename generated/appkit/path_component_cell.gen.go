// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PathComponentCell] class.
var pathComponentCellClass = _PathComponentCellClass{objc.GetClass("NSPathComponentCell")}

type _PathComponentCellClass struct {
	class objc.Class
}

// An interface definition for the [PathComponentCell] class.
type IPathComponentCell interface {
	ITextFieldCell
}

// A component of a path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathComponentCell

type PathComponentCell struct {
	TextFieldCell
}

// PathComponentCellFrom constructs a [PathComponentCell] from an unsafe.Pointer.
//
// A component of a path.
func PathComponentCellFrom(ptr unsafe.Pointer) PathComponentCell {
	return PathComponentCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}



