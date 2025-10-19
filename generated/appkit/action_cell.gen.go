// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ActionCell] class.
var (
	actionCellClass     _ActionCellClass
	actionCellClassOnce sync.Once
)

func getActionCellClass() _ActionCellClass {
	actionCellClassOnce.Do(func() {
		actionCellClass = _ActionCellClass{objc.GetClass("NSActionCell")}
	})
	return actionCellClass
}

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
// Alloc allocates a new instance without initialization.
func (ac _ActionCellClass) Alloc() ActionCell {
	rv := objc.Send[ActionCell](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _ActionCellClass) New() ActionCell {
	rv := objc.Send[ActionCell](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ActionCell) Init() ActionCell {
	rv := objc.Send[ActionCell](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ActionCell) Autorelease() ActionCell {
	rv := objc.Send[ActionCell](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewActionCell creates a new ActionCell instance.
func NewActionCell() ActionCell {
	return getActionCellClass().New()
}




