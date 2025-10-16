
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [resizeIncrements] class.
var resizeIncrementsClass _resizeIncrementsClass

func init() {
	resizeIncrementsClass = _resizeIncrementsClass{objc.GetClass("resizeIncrements")}
}

type _resizeIncrementsClass struct {
	objc.Class
}

// An interface definition for the [resizeIncrements] class.
type IresizeIncrements interface {
	ID() objc.ID
}

type resizeIncrements struct {
	id objc.ID
}

func resizeIncrementsFrom(ptr unsafe.Pointer) resizeIncrements {
	return resizeIncrements{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ resizeIncrements) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _resizeIncrementsClass) Alloc() resizeIncrements {
	rv := objc.Send[resizeIncrements](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _resizeIncrementsClass) New() resizeIncrements {
	rv := objc.Send[resizeIncrements](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewresizeIncrements creates and returns a new initialized instance.
func NewresizeIncrements() resizeIncrements {
	return resizeIncrementsClass.New()
}

// Init initializes the instance.
func (r_ resizeIncrements) Init() resizeIncrements {
	rv := objc.Send[resizeIncrements](r_.ID(), selInit)
	return rv
}
