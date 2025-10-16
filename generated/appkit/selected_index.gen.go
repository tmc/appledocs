
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [selectedIndex] class.
var selectedIndexClass _selectedIndexClass

func init() {
	selectedIndexClass = _selectedIndexClass{objc.GetClass("selectedIndex")}
}

type _selectedIndexClass struct {
	objc.Class
}

// An interface definition for the [selectedIndex] class.
type IselectedIndex interface {
	ID() objc.ID
}

type selectedIndex struct {
	id objc.ID
}

func selectedIndexFrom(ptr unsafe.Pointer) selectedIndex {
	return selectedIndex{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ selectedIndex) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _selectedIndexClass) Alloc() selectedIndex {
	rv := objc.Send[selectedIndex](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _selectedIndexClass) New() selectedIndex {
	rv := objc.Send[selectedIndex](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewselectedIndex creates and returns a new initialized instance.
func NewselectedIndex() selectedIndex {
	return selectedIndexClass.New()
}

// Init initializes the instance.
func (s_ selectedIndex) Init() selectedIndex {
	rv := objc.Send[selectedIndex](s_.ID(), selInit)
	return rv
}
