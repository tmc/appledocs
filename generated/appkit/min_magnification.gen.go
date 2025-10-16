
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [minMagnification] class.
var minMagnificationClass _minMagnificationClass

func init() {
	minMagnificationClass = _minMagnificationClass{objc.GetClass("minMagnification")}
}

type _minMagnificationClass struct {
	objc.Class
}

// An interface definition for the [minMagnification] class.
type IminMagnification interface {
	ID() objc.ID
}

type minMagnification struct {
	id objc.ID
}

func minMagnificationFrom(ptr unsafe.Pointer) minMagnification {
	return minMagnification{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ minMagnification) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _minMagnificationClass) Alloc() minMagnification {
	rv := objc.Send[minMagnification](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _minMagnificationClass) New() minMagnification {
	rv := objc.Send[minMagnification](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewminMagnification creates and returns a new initialized instance.
func NewminMagnification() minMagnification {
	return minMagnificationClass.New()
}

// Init initializes the instance.
func (m_ minMagnification) Init() minMagnification {
	rv := objc.Send[minMagnification](m_.ID(), selInit)
	return rv
}
