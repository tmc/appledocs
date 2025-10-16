
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [supportedWindowDepths] class.
var supportedWindowDepthsClass _supportedWindowDepthsClass

func init() {
	supportedWindowDepthsClass = _supportedWindowDepthsClass{objc.GetClass("supportedWindowDepths")}
}

type _supportedWindowDepthsClass struct {
	objc.Class
}

// An interface definition for the [supportedWindowDepths] class.
type IsupportedWindowDepths interface {
	ID() objc.ID
}

type supportedWindowDepths struct {
	id objc.ID
}

func supportedWindowDepthsFrom(ptr unsafe.Pointer) supportedWindowDepths {
	return supportedWindowDepths{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ supportedWindowDepths) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _supportedWindowDepthsClass) Alloc() supportedWindowDepths {
	rv := objc.Send[supportedWindowDepths](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _supportedWindowDepthsClass) New() supportedWindowDepths {
	rv := objc.Send[supportedWindowDepths](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewsupportedWindowDepths creates and returns a new initialized instance.
func NewsupportedWindowDepths() supportedWindowDepths {
	return supportedWindowDepthsClass.New()
}

// Init initializes the instance.
func (s_ supportedWindowDepths) Init() supportedWindowDepths {
	rv := objc.Send[supportedWindowDepths](s_.ID(), selInit)
	return rv
}
