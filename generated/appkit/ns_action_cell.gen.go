// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	ICell
}

// An active area inside a control.
//
// An does three things: it displays text or an icon; it provides the target object and action method used by its object; and it handles mouse (cursor) tracking by properly highlighting its area and sending action messages to its target based on cursor movement. The of an is the view in which the receiver was last drawn.
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


// Returns the receiver’s action-message selector.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell/action
func (a_ ActionCell) Action() objc.SEL {
	rv := objc.Send[objc.SEL](a_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// Returns the receiver’s action-message selector.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell/action
func (a_ ActionCell) SetAction(value objc.SEL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAction:"), value)
}

// Returns the receiver’s tag.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell/tag
func (a_ ActionCell) Tag() int {
	rv := objc.Send[int](a_.ID, objc.Sel("tag"))
	return rv
}


// SetTag sets the value of the tag property.
// Returns the receiver’s tag.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell/tag
func (a_ ActionCell) SetTag(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTag:"), value)
}

// Returns the receiver’s target object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell/target
func (a_ ActionCell) Target() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// Returns the receiver’s target object.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell/target
func (a_ ActionCell) SetTarget(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTarget:"), value)
}



