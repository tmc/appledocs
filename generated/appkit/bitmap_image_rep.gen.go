
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [BitmapImageRep] class.
var BitmapImageRepClass _BitmapImageRepClass

func init() {
	BitmapImageRepClass = _BitmapImageRepClass{objc.GetClass("NSBitmapImageRep")}
}

type _BitmapImageRepClass struct {
	objc.Class
}

// An interface definition for the [BitmapImageRep] class.
type IBitmapImageRep interface {
	ID() objc.ID
}

type BitmapImageRep struct {
	id objc.ID
}

func BitmapImageRepFrom(ptr unsafe.Pointer) BitmapImageRep {
	return BitmapImageRep{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ BitmapImageRep) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _BitmapImageRepClass) Alloc() BitmapImageRep {
	rv := objc.Send[BitmapImageRep](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _BitmapImageRepClass) New() BitmapImageRep {
	rv := objc.Send[BitmapImageRep](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewBitmapImageRep creates and returns a new initialized instance.
func NewBitmapImageRep() BitmapImageRep {
	return BitmapImageRepClass.New()
}

// Init initializes the instance.
func (b_ BitmapImageRep) Init() BitmapImageRep {
	rv := objc.Send[BitmapImageRep](b_.ID(), selInit)
	return rv
}
