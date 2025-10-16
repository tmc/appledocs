
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [selectedRange] class.
var selectedRangeClass _selectedRangeClass

func init() {
	selectedRangeClass = _selectedRangeClass{objc.GetClass("selectedRange")}
}

type _selectedRangeClass struct {
	objc.Class
}

// An interface definition for the [selectedRange] class.
type IselectedRange interface {
	ID() objc.ID
}

type selectedRange struct {
	id objc.ID
}

func selectedRangeFrom(ptr unsafe.Pointer) selectedRange {
	return selectedRange{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ selectedRange) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _selectedRangeClass) Alloc() selectedRange {
	rv := objc.Send[selectedRange](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _selectedRangeClass) New() selectedRange {
	rv := objc.Send[selectedRange](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewselectedRange creates and returns a new initialized instance.
func NewselectedRange() selectedRange {
	return selectedRangeClass.New()
}

// Init initializes the instance.
func (s_ selectedRange) Init() selectedRange {
	rv := objc.Send[selectedRange](s_.ID(), selInit)
	return rv
}
