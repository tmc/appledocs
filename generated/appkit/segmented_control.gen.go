
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SegmentedControl] class.
var SegmentedControlClass _SegmentedControlClass

func init() {
	SegmentedControlClass = _SegmentedControlClass{objc.GetClass("NSSegmentedControl")}
}

type _SegmentedControlClass struct {
	objc.Class
}

// An interface definition for the [SegmentedControl] class.
type ISegmentedControl interface {
	ID() objc.ID
}

type SegmentedControl struct {
	id objc.ID
}

func SegmentedControlFrom(ptr unsafe.Pointer) SegmentedControl {
	return SegmentedControl{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SegmentedControl) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SegmentedControlClass) Alloc() SegmentedControl {
	rv := objc.Send[SegmentedControl](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SegmentedControlClass) New() SegmentedControl {
	rv := objc.Send[SegmentedControl](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSegmentedControl creates and returns a new initialized instance.
func NewSegmentedControl() SegmentedControl {
	return SegmentedControlClass.New()
}

// Init initializes the instance.
func (s_ SegmentedControl) Init() SegmentedControl {
	rv := objc.Send[SegmentedControl](s_.ID(), selInit)
	return rv
}
// The color of the selected segment’s bezel, in appearances that support it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSegmentedControl/selectedSegmentBezelColor
func (s_ SegmentedControl) SelectedSegmentBezelColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("selectedSegmentBezelColor"))
	return rv
}
// SetSelectedSegmentBezelColor sets the value of the selectedSegmentBezelColor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSegmentedControl/selectedSegmentBezelColor
func (s_ SegmentedControl) SetSelectedSegmentBezelColor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setSelectedSegmentBezelColor:"), value)
}
