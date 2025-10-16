
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [invalidateShadow] class.
var invalidateShadowClass _invalidateShadowClass

func init() {
	invalidateShadowClass = _invalidateShadowClass{objc.GetClass("invalidateShadow")}
}

type _invalidateShadowClass struct {
	objc.Class
}

// An interface definition for the [invalidateShadow] class.
type IinvalidateShadow interface {
	ID() objc.ID
}

type invalidateShadow struct {
	id objc.ID
}

func invalidateShadowFrom(ptr unsafe.Pointer) invalidateShadow {
	return invalidateShadow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ invalidateShadow) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _invalidateShadowClass) Alloc() invalidateShadow {
	rv := objc.Send[invalidateShadow](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _invalidateShadowClass) New() invalidateShadow {
	rv := objc.Send[invalidateShadow](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewinvalidateShadow creates and returns a new initialized instance.
func NewinvalidateShadow() invalidateShadow {
	return invalidateShadowClass.New()
}

// Init initializes the instance.
func (i_ invalidateShadow) Init() invalidateShadow {
	rv := objc.Send[invalidateShadow](i_.ID(), selInit)
	return rv
}
