
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [superview] class.
var superviewClass _superviewClass

func init() {
	superviewClass = _superviewClass{objc.GetClass("superview")}
}

type _superviewClass struct {
	objc.Class
}

// An interface definition for the [superview] class.
type Isuperview interface {
	ID() objc.ID
}

type superview struct {
	id objc.ID
}

func superviewFrom(ptr unsafe.Pointer) superview {
	return superview{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ superview) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _superviewClass) Alloc() superview {
	rv := objc.Send[superview](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _superviewClass) New() superview {
	rv := objc.Send[superview](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newsuperview creates and returns a new initialized instance.
func Newsuperview() superview {
	return superviewClass.New()
}

// Init initializes the instance.
func (s_ superview) Init() superview {
	rv := objc.Send[superview](s_.ID(), selInit)
	return rv
}
