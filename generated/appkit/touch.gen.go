
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Touch] class.
var TouchClass _TouchClass

func init() {
	TouchClass = _TouchClass{objc.GetClass("NSTouch")}
}

type _TouchClass struct {
	objc.Class
}

// An interface definition for the [Touch] class.
type ITouch interface {
	ID() objc.ID
	PreviousLocationInView(view unsafe.Pointer) unsafe.Pointer
}

type Touch struct {
	id objc.ID
}

func TouchFrom(ptr unsafe.Pointer) Touch {
	return Touch{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ Touch) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TouchClass) Alloc() Touch {
	rv := objc.Send[Touch](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TouchClass) New() Touch {
	rv := objc.Send[Touch](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTouch creates and returns a new initialized instance.
func NewTouch() Touch {
	return TouchClass.New()
}

// Init initializes the instance.
func (t_ Touch) Init() Touch {
	rv := objc.Send[Touch](t_.ID(), selInit)
	return rv
}
// Indicates the previous location of the touch in the view’s coordinates. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouch/previousLocation(in:)
func (t_ Touch) PreviousLocationInView(view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("previousLocationInView:"), view)
	return rv
}
