
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [selectedSegmentBezelColor] class.
var selectedSegmentBezelColorClass _selectedSegmentBezelColorClass

func init() {
	selectedSegmentBezelColorClass = _selectedSegmentBezelColorClass{objc.GetClass("selectedSegmentBezelColor")}
}

type _selectedSegmentBezelColorClass struct {
	objc.Class
}

// An interface definition for the [selectedSegmentBezelColor] class.
type IselectedSegmentBezelColor interface {
	ID() objc.ID
}

type selectedSegmentBezelColor struct {
	id objc.ID
}

func selectedSegmentBezelColorFrom(ptr unsafe.Pointer) selectedSegmentBezelColor {
	return selectedSegmentBezelColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ selectedSegmentBezelColor) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _selectedSegmentBezelColorClass) Alloc() selectedSegmentBezelColor {
	rv := objc.Send[selectedSegmentBezelColor](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _selectedSegmentBezelColorClass) New() selectedSegmentBezelColor {
	rv := objc.Send[selectedSegmentBezelColor](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewselectedSegmentBezelColor creates and returns a new initialized instance.
func NewselectedSegmentBezelColor() selectedSegmentBezelColor {
	return selectedSegmentBezelColorClass.New()
}

// Init initializes the instance.
func (s_ selectedSegmentBezelColor) Init() selectedSegmentBezelColor {
	rv := objc.Send[selectedSegmentBezelColor](s_.ID(), selInit)
	return rv
}
