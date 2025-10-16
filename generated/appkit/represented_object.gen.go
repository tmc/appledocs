
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [representedObject] class.
var representedObjectClass _representedObjectClass

func init() {
	representedObjectClass = _representedObjectClass{objc.GetClass("representedObject")}
}

type _representedObjectClass struct {
	objc.Class
}

// An interface definition for the [representedObject] class.
type IrepresentedObject interface {
	ID() objc.ID
}

type representedObject struct {
	id objc.ID
}

func representedObjectFrom(ptr unsafe.Pointer) representedObject {
	return representedObject{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ representedObject) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _representedObjectClass) Alloc() representedObject {
	rv := objc.Send[representedObject](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _representedObjectClass) New() representedObject {
	rv := objc.Send[representedObject](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrepresentedObject creates and returns a new initialized instance.
func NewrepresentedObject() representedObject {
	return representedObjectClass.New()
}

// Init initializes the instance.
func (r_ representedObject) Init() representedObject {
	rv := objc.Send[representedObject](r_.ID(), selInit)
	return rv
}
