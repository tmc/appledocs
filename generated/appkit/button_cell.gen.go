// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ButtonCell] class.
var (
	buttonCellClass     _ButtonCellClass
	buttonCellClassOnce sync.Once
)

func getButtonCellClass() _ButtonCellClass {
	buttonCellClassOnce.Do(func() {
		buttonCellClass = _ButtonCellClass{objc.GetClass("NSButtonCell")}
	})
	return buttonCellClass
}

type _ButtonCellClass struct {
	class objc.Class
}

// An interface definition for the [ButtonCell] class.
type IButtonCell interface {
	IActionCell
}

// An object that defines the user interface of a button or other clickable region of a view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonCell

type ButtonCell struct {
	ActionCell
}

// ButtonCellFrom constructs a [ButtonCell] from an unsafe.Pointer.
//
// An object that defines the user interface of a button or other clickable region of a view.
func ButtonCellFrom(ptr unsafe.Pointer) ButtonCell {
	return ButtonCell{
		ActionCell: ActionCellFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (bc _ButtonCellClass) Alloc() ButtonCell {
	rv := objc.Send[ButtonCell](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _ButtonCellClass) New() ButtonCell {
	rv := objc.Send[ButtonCell](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ ButtonCell) Init() ButtonCell {
	rv := objc.Send[ButtonCell](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ ButtonCell) Autorelease() ButtonCell {
	rv := objc.Send[ButtonCell](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewButtonCell creates a new ButtonCell instance.
func NewButtonCell() ButtonCell {
	return getButtonCellClass().New()
}




