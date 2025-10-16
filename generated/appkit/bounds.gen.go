
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [bounds] class.
var boundsClass _boundsClass

func init() {
	boundsClass = _boundsClass{objc.GetClass("bounds")}
}

type _boundsClass struct {
	objc.Class
}

// An interface definition for the [bounds] class.
type Ibounds interface {
	ID() objc.ID
}

type bounds struct {
	id objc.ID
}

func boundsFrom(ptr unsafe.Pointer) bounds {
	return bounds{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ bounds) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _boundsClass) Alloc() bounds {
	rv := objc.Send[bounds](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _boundsClass) New() bounds {
	rv := objc.Send[bounds](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newbounds creates and returns a new initialized instance.
func Newbounds() bounds {
	return boundsClass.New()
}

// Init initializes the instance.
func (b_ bounds) Init() bounds {
	rv := objc.Send[bounds](b_.ID(), selInit)
	return rv
}
