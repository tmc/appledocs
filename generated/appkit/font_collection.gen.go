
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FontCollection] class.
var FontCollectionClass _FontCollectionClass

func init() {
	FontCollectionClass = _FontCollectionClass{objc.GetClass("NSFontCollection")}
}

type _FontCollectionClass struct {
	objc.Class
}

// An interface definition for the [FontCollection] class.
type IFontCollection interface {
	ID() objc.ID
}

type FontCollection struct {
	id objc.ID
}

func FontCollectionFrom(ptr unsafe.Pointer) FontCollection {
	return FontCollection{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ FontCollection) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _FontCollectionClass) Alloc() FontCollection {
	rv := objc.Send[FontCollection](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _FontCollectionClass) New() FontCollection {
	rv := objc.Send[FontCollection](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewFontCollection creates and returns a new initialized instance.
func NewFontCollection() FontCollection {
	return FontCollectionClass.New()
}

// Init initializes the instance.
func (f_ FontCollection) Init() FontCollection {
	rv := objc.Send[FontCollection](f_.ID(), selInit)
	return rv
}
