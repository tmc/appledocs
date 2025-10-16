
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [sizeToFit] class.
var sizeToFitClass _sizeToFitClass

func init() {
	sizeToFitClass = _sizeToFitClass{objc.GetClass("sizeToFit")}
}

type _sizeToFitClass struct {
	objc.Class
}

// An interface definition for the [sizeToFit] class.
type IsizeToFit interface {
	ID() objc.ID
}

type sizeToFit struct {
	id objc.ID
}

func sizeToFitFrom(ptr unsafe.Pointer) sizeToFit {
	return sizeToFit{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ sizeToFit) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _sizeToFitClass) Alloc() sizeToFit {
	rv := objc.Send[sizeToFit](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _sizeToFitClass) New() sizeToFit {
	rv := objc.Send[sizeToFit](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewsizeToFit creates and returns a new initialized instance.
func NewsizeToFit() sizeToFit {
	return sizeToFitClass.New()
}

// Init initializes the instance.
func (s_ sizeToFit) Init() sizeToFit {
	rv := objc.Send[sizeToFit](s_.ID(), selInit)
	return rv
}
