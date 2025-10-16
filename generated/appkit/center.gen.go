
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [center] class.
var centerClass _centerClass

func init() {
	centerClass = _centerClass{objc.GetClass("center")}
}

type _centerClass struct {
	objc.Class
}

// An interface definition for the [center] class.
type Icenter interface {
	ID() objc.ID
}

type center struct {
	id objc.ID
}

func centerFrom(ptr unsafe.Pointer) center {
	return center{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ center) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _centerClass) Alloc() center {
	rv := objc.Send[center](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _centerClass) New() center {
	rv := objc.Send[center](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newcenter creates and returns a new initialized instance.
func Newcenter() center {
	return centerClass.New()
}

// Init initializes the instance.
func (c_ center) Init() center {
	rv := objc.Send[center](c_.ID(), selInit)
	return rv
}
