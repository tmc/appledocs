
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [orientation] class.
var orientationClass _orientationClass

func init() {
	orientationClass = _orientationClass{objc.GetClass("orientation")}
}

type _orientationClass struct {
	objc.Class
}

// An interface definition for the [orientation] class.
type Iorientation interface {
	ID() objc.ID
}

type orientation struct {
	id objc.ID
}

func orientationFrom(ptr unsafe.Pointer) orientation {
	return orientation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ orientation) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _orientationClass) Alloc() orientation {
	rv := objc.Send[orientation](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _orientationClass) New() orientation {
	rv := objc.Send[orientation](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Neworientation creates and returns a new initialized instance.
func Neworientation() orientation {
	return orientationClass.New()
}

// Init initializes the instance.
func (o_ orientation) Init() orientation {
	rv := objc.Send[orientation](o_.ID(), selInit)
	return rv
}
