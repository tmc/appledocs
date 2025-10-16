
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [sizeMode] class.
var sizeModeClass _sizeModeClass

func init() {
	sizeModeClass = _sizeModeClass{objc.GetClass("sizeMode")}
}

type _sizeModeClass struct {
	objc.Class
}

// An interface definition for the [sizeMode] class.
type IsizeMode interface {
	ID() objc.ID
}

type sizeMode struct {
	id objc.ID
}

func sizeModeFrom(ptr unsafe.Pointer) sizeMode {
	return sizeMode{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ sizeMode) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _sizeModeClass) Alloc() sizeMode {
	rv := objc.Send[sizeMode](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _sizeModeClass) New() sizeMode {
	rv := objc.Send[sizeMode](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewsizeMode creates and returns a new initialized instance.
func NewsizeMode() sizeMode {
	return sizeModeClass.New()
}

// Init initializes the instance.
func (s_ sizeMode) Init() sizeMode {
	rv := objc.Send[sizeMode](s_.ID(), selInit)
	return rv
}
