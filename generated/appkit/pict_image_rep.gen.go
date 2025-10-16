
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PICTImageRep] class.
var PICTImageRepClass _PICTImageRepClass

func init() {
	PICTImageRepClass = _PICTImageRepClass{objc.GetClass("NSPICTImageRep")}
}

type _PICTImageRepClass struct {
	objc.Class
}

// An interface definition for the [PICTImageRep] class.
type IPICTImageRep interface {
	ID() objc.ID
}

type PICTImageRep struct {
	id objc.ID
}

func PICTImageRepFrom(ptr unsafe.Pointer) PICTImageRep {
	return PICTImageRep{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PICTImageRep) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PICTImageRepClass) Alloc() PICTImageRep {
	rv := objc.Send[PICTImageRep](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PICTImageRepClass) New() PICTImageRep {
	rv := objc.Send[PICTImageRep](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPICTImageRep creates and returns a new initialized instance.
func NewPICTImageRep() PICTImageRep {
	return PICTImageRepClass.New()
}

// Init initializes the instance.
func (p_ PICTImageRep) Init() PICTImageRep {
	rv := objc.Send[PICTImageRep](p_.ID(), selInit)
	return rv
}
