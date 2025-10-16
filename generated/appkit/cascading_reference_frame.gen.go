
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [cascadingReferenceFrame] class.
var cascadingReferenceFrameClass _cascadingReferenceFrameClass

func init() {
	cascadingReferenceFrameClass = _cascadingReferenceFrameClass{objc.GetClass("cascadingReferenceFrame")}
}

type _cascadingReferenceFrameClass struct {
	objc.Class
}

// An interface definition for the [cascadingReferenceFrame] class.
type IcascadingReferenceFrame interface {
	ID() objc.ID
}

type cascadingReferenceFrame struct {
	id objc.ID
}

func cascadingReferenceFrameFrom(ptr unsafe.Pointer) cascadingReferenceFrame {
	return cascadingReferenceFrame{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ cascadingReferenceFrame) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _cascadingReferenceFrameClass) Alloc() cascadingReferenceFrame {
	rv := objc.Send[cascadingReferenceFrame](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _cascadingReferenceFrameClass) New() cascadingReferenceFrame {
	rv := objc.Send[cascadingReferenceFrame](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcascadingReferenceFrame creates and returns a new initialized instance.
func NewcascadingReferenceFrame() cascadingReferenceFrame {
	return cascadingReferenceFrameClass.New()
}

// Init initializes the instance.
func (c_ cascadingReferenceFrame) Init() cascadingReferenceFrame {
	rv := objc.Send[cascadingReferenceFrame](c_.ID(), selInit)
	return rv
}
