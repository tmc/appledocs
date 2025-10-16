
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [fullScreenAccessoryViewMaxHeight] class.
var fullScreenAccessoryViewMaxHeightClass _fullScreenAccessoryViewMaxHeightClass

func init() {
	fullScreenAccessoryViewMaxHeightClass = _fullScreenAccessoryViewMaxHeightClass{objc.GetClass("fullScreenAccessoryViewMaxHeight")}
}

type _fullScreenAccessoryViewMaxHeightClass struct {
	objc.Class
}

// An interface definition for the [fullScreenAccessoryViewMaxHeight] class.
type IfullScreenAccessoryViewMaxHeight interface {
	ID() objc.ID
}

type fullScreenAccessoryViewMaxHeight struct {
	id objc.ID
}

func fullScreenAccessoryViewMaxHeightFrom(ptr unsafe.Pointer) fullScreenAccessoryViewMaxHeight {
	return fullScreenAccessoryViewMaxHeight{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ fullScreenAccessoryViewMaxHeight) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _fullScreenAccessoryViewMaxHeightClass) Alloc() fullScreenAccessoryViewMaxHeight {
	rv := objc.Send[fullScreenAccessoryViewMaxHeight](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _fullScreenAccessoryViewMaxHeightClass) New() fullScreenAccessoryViewMaxHeight {
	rv := objc.Send[fullScreenAccessoryViewMaxHeight](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfullScreenAccessoryViewMaxHeight creates and returns a new initialized instance.
func NewfullScreenAccessoryViewMaxHeight() fullScreenAccessoryViewMaxHeight {
	return fullScreenAccessoryViewMaxHeightClass.New()
}

// Init initializes the instance.
func (f_ fullScreenAccessoryViewMaxHeight) Init() fullScreenAccessoryViewMaxHeight {
	rv := objc.Send[fullScreenAccessoryViewMaxHeight](f_.ID(), selInit)
	return rv
}
