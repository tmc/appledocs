
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsMultipleSelection] class.
var allowsMultipleSelectionClass _allowsMultipleSelectionClass

func init() {
	allowsMultipleSelectionClass = _allowsMultipleSelectionClass{objc.GetClass("allowsMultipleSelection")}
}

type _allowsMultipleSelectionClass struct {
	objc.Class
}

// An interface definition for the [allowsMultipleSelection] class.
type IallowsMultipleSelection interface {
	ID() objc.ID
}

type allowsMultipleSelection struct {
	id objc.ID
}

func allowsMultipleSelectionFrom(ptr unsafe.Pointer) allowsMultipleSelection {
	return allowsMultipleSelection{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsMultipleSelection) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsMultipleSelectionClass) Alloc() allowsMultipleSelection {
	rv := objc.Send[allowsMultipleSelection](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsMultipleSelectionClass) New() allowsMultipleSelection {
	rv := objc.Send[allowsMultipleSelection](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsMultipleSelection creates and returns a new initialized instance.
func NewallowsMultipleSelection() allowsMultipleSelection {
	return allowsMultipleSelectionClass.New()
}

// Init initializes the instance.
func (a_ allowsMultipleSelection) Init() allowsMultipleSelection {
	rv := objc.Send[allowsMultipleSelection](a_.ID(), selInit)
	return rv
}
