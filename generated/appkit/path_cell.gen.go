// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PathCell] class.
var pathCellClass = _PathCellClass{objc.GetClass("NSPathCell")}

type _PathCellClass struct {
	class objc.Class
}

// The user interface of a path control object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathCell

type PathCell struct {
	ActionCell
}

// PathCellFrom constructs a [PathCell] from an unsafe.Pointer.
//
// The user interface of a path control object.
func PathCellFrom(ptr unsafe.Pointer) PathCell {
	return PathCell{
		ActionCell: ActionCellFrom(ptr),
	}
}



