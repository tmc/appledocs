
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsEmptySelection] class.
var allowsEmptySelectionClass _allowsEmptySelectionClass

func init() {
	allowsEmptySelectionClass = _allowsEmptySelectionClass{objc.GetClass("allowsEmptySelection")}
}

type _allowsEmptySelectionClass struct {
	objc.Class
}

// An interface definition for the [allowsEmptySelection] class.
type IallowsEmptySelection interface {
	ID() objc.ID
}

type allowsEmptySelection struct {
	id objc.ID
}

func allowsEmptySelectionFrom(ptr unsafe.Pointer) allowsEmptySelection {
	return allowsEmptySelection{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsEmptySelection) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsEmptySelectionClass) Alloc() allowsEmptySelection {
	rv := objc.Send[allowsEmptySelection](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsEmptySelectionClass) New() allowsEmptySelection {
	rv := objc.Send[allowsEmptySelection](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsEmptySelection creates and returns a new initialized instance.
func NewallowsEmptySelection() allowsEmptySelection {
	return allowsEmptySelectionClass.New()
}

// Init initializes the instance.
func (a_ allowsEmptySelection) Init() allowsEmptySelection {
	rv := objc.Send[allowsEmptySelection](a_.ID(), selInit)
	return rv
}
