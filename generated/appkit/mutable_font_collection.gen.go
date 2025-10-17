
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableFontCollection] class.
var MutableFontCollectionClass _MutableFontCollectionClass

func init() {
	MutableFontCollectionClass = _MutableFontCollectionClass{objc.GetClass("NSMutableFontCollection")}
}

type _MutableFontCollectionClass struct {
	objc.Class
}

// An interface definition for the [MutableFontCollection] class.
type IMutableFontCollection interface {
	ID() objc.ID
}

type MutableFontCollection struct {
	id objc.ID
}

func MutableFontCollectionFrom(ptr unsafe.Pointer) MutableFontCollection {
	return MutableFontCollection{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ MutableFontCollection) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _MutableFontCollectionClass) Alloc() MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _MutableFontCollectionClass) New() MutableFontCollection {
	rv := objc.Send[MutableFontCollection](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewMutableFontCollection creates and returns a new initialized instance.
func NewMutableFontCollection() MutableFontCollection {
	return MutableFontCollectionClass.New()
}

// Init initializes the instance.
func (m_ MutableFontCollection) Init() MutableFontCollection {
	rv := objc.Send[MutableFontCollection](m_.ID(), selInit)
	return rv
}
