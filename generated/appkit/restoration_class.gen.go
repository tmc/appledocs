
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [restorationClass] class.
var restorationClassClass _restorationClassClass

func init() {
	restorationClassClass = _restorationClassClass{objc.GetClass("restorationClass")}
}

type _restorationClassClass struct {
	objc.Class
}

// An interface definition for the [restorationClass] class.
type IrestorationClass interface {
	ID() objc.ID
}

type restorationClass struct {
	id objc.ID
}

func restorationClassFrom(ptr unsafe.Pointer) restorationClass {
	return restorationClass{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ restorationClass) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _restorationClassClass) Alloc() restorationClass {
	rv := objc.Send[restorationClass](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _restorationClassClass) New() restorationClass {
	rv := objc.Send[restorationClass](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrestorationClass creates and returns a new initialized instance.
func NewrestorationClass() restorationClass {
	return restorationClassClass.New()
}

// Init initializes the instance.
func (r_ restorationClass) Init() restorationClass {
	rv := objc.Send[restorationClass](r_.ID(), selInit)
	return rv
}
