
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [maxMagnification] class.
var maxMagnificationClass _maxMagnificationClass

func init() {
	maxMagnificationClass = _maxMagnificationClass{objc.GetClass("maxMagnification")}
}

type _maxMagnificationClass struct {
	objc.Class
}

// An interface definition for the [maxMagnification] class.
type ImaxMagnification interface {
	ID() objc.ID
}

type maxMagnification struct {
	id objc.ID
}

func maxMagnificationFrom(ptr unsafe.Pointer) maxMagnification {
	return maxMagnification{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ maxMagnification) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _maxMagnificationClass) Alloc() maxMagnification {
	rv := objc.Send[maxMagnification](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _maxMagnificationClass) New() maxMagnification {
	rv := objc.Send[maxMagnification](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmaxMagnification creates and returns a new initialized instance.
func NewmaxMagnification() maxMagnification {
	return maxMagnificationClass.New()
}

// Init initializes the instance.
func (m_ maxMagnification) Init() maxMagnification {
	rv := objc.Send[maxMagnification](m_.ID(), selInit)
	return rv
}
