// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ActionCell] class.
var actionCellClass = _ActionCellClass{objc.GetClass("NSActionCell")}

type _ActionCellClass struct {
	class objc.Class
}

// An interface definition for the [ActionCell] class.
type IActionCell interface {
	ICell
}

// An active area inside a control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell

type ActionCell struct {
	Cell
}

// ActionCellFrom constructs a [ActionCell] from an unsafe.Pointer.
//
// An active area inside a control.
func ActionCellFrom(ptr unsafe.Pointer) ActionCell {
	return ActionCell{
		Cell: CellFrom(ptr),
	}
}



