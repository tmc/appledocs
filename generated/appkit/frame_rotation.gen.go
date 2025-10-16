
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [frameRotation] class.
var frameRotationClass _frameRotationClass

func init() {
	frameRotationClass = _frameRotationClass{objc.GetClass("frameRotation")}
}

type _frameRotationClass struct {
	objc.Class
}

// An interface definition for the [frameRotation] class.
type IframeRotation interface {
	ID() objc.ID
}

type frameRotation struct {
	id objc.ID
}

func frameRotationFrom(ptr unsafe.Pointer) frameRotation {
	return frameRotation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ frameRotation) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _frameRotationClass) Alloc() frameRotation {
	rv := objc.Send[frameRotation](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _frameRotationClass) New() frameRotation {
	rv := objc.Send[frameRotation](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewframeRotation creates and returns a new initialized instance.
func NewframeRotation() frameRotation {
	return frameRotationClass.New()
}

// Init initializes the instance.
func (f_ frameRotation) Init() frameRotation {
	rv := objc.Send[frameRotation](f_.ID(), selInit)
	return rv
}
