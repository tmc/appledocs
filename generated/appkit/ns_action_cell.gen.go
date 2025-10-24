// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ActionCell] class.
var (
	ActionCellClass     _ActionCellClass
	ActionCellClassOnce sync.Once
)

func getActionCellClass() _ActionCellClass {
	ActionCellClassOnce.Do(func() {
		ActionCellClass = _ActionCellClass{objc.GetClass("NSActionCell")}
	})
	return ActionCellClass
}

type _ActionCellClass struct {
	class objc.Class
}

// An interface definition for the [ActionCell] class.
type IActionCell interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other AppKit classes.


// A parent class referenced by other AppKit classes. [Full Topic]
type ActionCell struct {
	objectivec.Object
}

// ActionCellFrom constructs a [ActionCell] from an unsafe.Pointer.
//
// A parent class referenced by other AppKit classes.
func ActionCellFrom(ptr unsafe.Pointer) ActionCell {
	return ActionCell{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ActionCellClass) Alloc() ActionCell {
	rv := objc.Send[ActionCell](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




