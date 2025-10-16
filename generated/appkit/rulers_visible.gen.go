
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [rulersVisible] class.
var rulersVisibleClass _rulersVisibleClass

func init() {
	rulersVisibleClass = _rulersVisibleClass{objc.GetClass("rulersVisible")}
}

type _rulersVisibleClass struct {
	objc.Class
}

// An interface definition for the [rulersVisible] class.
type IrulersVisible interface {
	ID() objc.ID
}

type rulersVisible struct {
	id objc.ID
}

func rulersVisibleFrom(ptr unsafe.Pointer) rulersVisible {
	return rulersVisible{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ rulersVisible) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _rulersVisibleClass) Alloc() rulersVisible {
	rv := objc.Send[rulersVisible](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _rulersVisibleClass) New() rulersVisible {
	rv := objc.Send[rulersVisible](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrulersVisible creates and returns a new initialized instance.
func NewrulersVisible() rulersVisible {
	return rulersVisibleClass.New()
}

// Init initializes the instance.
func (r_ rulersVisible) Init() rulersVisible {
	rv := objc.Send[rulersVisible](r_.ID(), selInit)
	return rv
}
