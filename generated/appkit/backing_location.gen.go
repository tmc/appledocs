
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [backingLocation] class.
var backingLocationClass _backingLocationClass

func init() {
	backingLocationClass = _backingLocationClass{objc.GetClass("backingLocation")}
}

type _backingLocationClass struct {
	objc.Class
}

// An interface definition for the [backingLocation] class.
type IbackingLocation interface {
	ID() objc.ID
}

type backingLocation struct {
	id objc.ID
}

func backingLocationFrom(ptr unsafe.Pointer) backingLocation {
	return backingLocation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ backingLocation) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _backingLocationClass) Alloc() backingLocation {
	rv := objc.Send[backingLocation](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _backingLocationClass) New() backingLocation {
	rv := objc.Send[backingLocation](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbackingLocation creates and returns a new initialized instance.
func NewbackingLocation() backingLocation {
	return backingLocationClass.New()
}

// Init initializes the instance.
func (b_ backingLocation) Init() backingLocation {
	rv := objc.Send[backingLocation](b_.ID(), selInit)
	return rv
}
