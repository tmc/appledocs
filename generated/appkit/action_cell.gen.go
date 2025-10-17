
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ActionCell] class.
var ActionCellClass _ActionCellClass

func init() {
	ActionCellClass = _ActionCellClass{objc.GetClass("NSActionCell")}
}

type _ActionCellClass struct {
	objc.Class
}

// An interface definition for the [ActionCell] class.
type IActionCell interface {
	ID() objc.ID
}

type ActionCell struct {
	id objc.ID
}

func ActionCellFrom(ptr unsafe.Pointer) ActionCell {
	return ActionCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ ActionCell) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _ActionCellClass) Alloc() ActionCell {
	rv := objc.Send[ActionCell](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _ActionCellClass) New() ActionCell {
	rv := objc.Send[ActionCell](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewActionCell creates and returns a new initialized instance.
func NewActionCell() ActionCell {
	return ActionCellClass.New()
}

// Init initializes the instance.
func (a_ ActionCell) Init() ActionCell {
	rv := objc.Send[ActionCell](a_.ID(), selInit)
	return rv
}
// Returns the receiver’s action-message selector. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSActionCell/action
func (a_ ActionCell) Action() objc.SEL {
	rv := objc.Send[objc.SEL](a_.ID(), objc.RegisterName("action"))
	return rv
}
// SetAction sets the value of the action property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSActionCell/action
func (a_ ActionCell) SetAction(value objc.SEL) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("setAction:"), value)
}
// Returns the receiver’s tag. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSActionCell/tag
func (a_ ActionCell) Tag() int {
	rv := objc.Send[int](a_.ID(), objc.RegisterName("tag"))
	return rv
}
// SetTag sets the value of the tag property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSActionCell/tag
func (a_ ActionCell) SetTag(value int) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("setTag:"), value)
}
// Returns the receiver’s target object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSActionCell/target
func (a_ ActionCell) Target() objc.ID {
	rv := objc.Send[objc.ID](a_.ID(), objc.RegisterName("target"))
	return rv
}
// SetTarget sets the value of the target property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSActionCell/target
func (a_ ActionCell) SetTarget(value objc.ID) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("setTarget:"), value)
}
