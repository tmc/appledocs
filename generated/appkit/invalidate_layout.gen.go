
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [invalidateLayout] class.
var invalidateLayoutClass _invalidateLayoutClass

func init() {
	invalidateLayoutClass = _invalidateLayoutClass{objc.GetClass("invalidateLayout")}
}

type _invalidateLayoutClass struct {
	objc.Class
}

// An interface definition for the [invalidateLayout] class.
type IinvalidateLayout interface {
	ID() objc.ID
}

type invalidateLayout struct {
	id objc.ID
}

func invalidateLayoutFrom(ptr unsafe.Pointer) invalidateLayout {
	return invalidateLayout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ invalidateLayout) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _invalidateLayoutClass) Alloc() invalidateLayout {
	rv := objc.Send[invalidateLayout](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _invalidateLayoutClass) New() invalidateLayout {
	rv := objc.Send[invalidateLayout](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewinvalidateLayout creates and returns a new initialized instance.
func NewinvalidateLayout() invalidateLayout {
	return invalidateLayoutClass.New()
}

// Init initializes the instance.
func (i_ invalidateLayout) Init() invalidateLayout {
	rv := objc.Send[invalidateLayout](i_.ID(), selInit)
	return rv
}
