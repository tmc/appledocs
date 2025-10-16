
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [fullScreenAccessoryViewMinHeight] class.
var fullScreenAccessoryViewMinHeightClass _fullScreenAccessoryViewMinHeightClass

func init() {
	fullScreenAccessoryViewMinHeightClass = _fullScreenAccessoryViewMinHeightClass{objc.GetClass("fullScreenAccessoryViewMinHeight")}
}

type _fullScreenAccessoryViewMinHeightClass struct {
	objc.Class
}

// An interface definition for the [fullScreenAccessoryViewMinHeight] class.
type IfullScreenAccessoryViewMinHeight interface {
	ID() objc.ID
}

type fullScreenAccessoryViewMinHeight struct {
	id objc.ID
}

func fullScreenAccessoryViewMinHeightFrom(ptr unsafe.Pointer) fullScreenAccessoryViewMinHeight {
	return fullScreenAccessoryViewMinHeight{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ fullScreenAccessoryViewMinHeight) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _fullScreenAccessoryViewMinHeightClass) Alloc() fullScreenAccessoryViewMinHeight {
	rv := objc.Send[fullScreenAccessoryViewMinHeight](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _fullScreenAccessoryViewMinHeightClass) New() fullScreenAccessoryViewMinHeight {
	rv := objc.Send[fullScreenAccessoryViewMinHeight](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfullScreenAccessoryViewMinHeight creates and returns a new initialized instance.
func NewfullScreenAccessoryViewMinHeight() fullScreenAccessoryViewMinHeight {
	return fullScreenAccessoryViewMinHeightClass.New()
}

// Init initializes the instance.
func (f_ fullScreenAccessoryViewMinHeight) Init() fullScreenAccessoryViewMinHeight {
	rv := objc.Send[fullScreenAccessoryViewMinHeight](f_.ID(), selInit)
	return rv
}
