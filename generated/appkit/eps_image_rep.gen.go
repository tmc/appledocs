
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [EPSImageRep] class.
var EPSImageRepClass _EPSImageRepClass

func init() {
	EPSImageRepClass = _EPSImageRepClass{objc.GetClass("NSEPSImageRep")}
}

type _EPSImageRepClass struct {
	objc.Class
}

// An interface definition for the [EPSImageRep] class.
type IEPSImageRep interface {
	ID() objc.ID
}

type EPSImageRep struct {
	id objc.ID
}

func EPSImageRepFrom(ptr unsafe.Pointer) EPSImageRep {
	return EPSImageRep{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ EPSImageRep) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _EPSImageRepClass) Alloc() EPSImageRep {
	rv := objc.Send[EPSImageRep](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _EPSImageRepClass) New() EPSImageRep {
	rv := objc.Send[EPSImageRep](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewEPSImageRep creates and returns a new initialized instance.
func NewEPSImageRep() EPSImageRep {
	return EPSImageRepClass.New()
}

// Init initializes the instance.
func (e_ EPSImageRep) Init() EPSImageRep {
	rv := objc.Send[EPSImageRep](e_.ID(), selInit)
	return rv
}
