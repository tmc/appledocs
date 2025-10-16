
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [range_] class.
var range_Class _range_Class

func init() {
	range_Class = _range_Class{objc.GetClass("range")}
}

type _range_Class struct {
	objc.Class
}

// An interface definition for the [range_] class.
type Irange_ interface {
	ID() objc.ID
}

type range_ struct {
	id objc.ID
}

func range_From(ptr unsafe.Pointer) range_ {
	return range_{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ range_) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _range_Class) Alloc() range_ {
	rv := objc.Send[range_](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _range_Class) New() range_ {
	rv := objc.Send[range_](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newrange_ creates and returns a new initialized instance.
func Newrange_() range_ {
	return range_Class.New()
}

// Init initializes the instance.
func (r_ range_) Init() range_ {
	rv := objc.Send[range_](r_.ID(), selInit)
	return rv
}
