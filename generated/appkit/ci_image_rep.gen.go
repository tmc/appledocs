
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CIImageRep] class.
var CIImageRepClass _CIImageRepClass

func init() {
	CIImageRepClass = _CIImageRepClass{objc.GetClass("NSCIImageRep")}
}

type _CIImageRepClass struct {
	objc.Class
}

// An interface definition for the [CIImageRep] class.
type ICIImageRep interface {
	ID() objc.ID
}

type CIImageRep struct {
	id objc.ID
}

func CIImageRepFrom(ptr unsafe.Pointer) CIImageRep {
	return CIImageRep{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ CIImageRep) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _CIImageRepClass) Alloc() CIImageRep {
	rv := objc.Send[CIImageRep](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _CIImageRepClass) New() CIImageRep {
	rv := objc.Send[CIImageRep](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewCIImageRep creates and returns a new initialized instance.
func NewCIImageRep() CIImageRep {
	return CIImageRepClass.New()
}

// Init initializes the instance.
func (i_ CIImageRep) Init() CIImageRep {
	rv := objc.Send[CIImageRep](i_.ID(), selInit)
	return rv
}
