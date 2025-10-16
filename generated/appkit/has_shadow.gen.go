
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [hasShadow] class.
var hasShadowClass _hasShadowClass

func init() {
	hasShadowClass = _hasShadowClass{objc.GetClass("hasShadow")}
}

type _hasShadowClass struct {
	objc.Class
}

// An interface definition for the [hasShadow] class.
type IhasShadow interface {
	ID() objc.ID
}

type hasShadow struct {
	id objc.ID
}

func hasShadowFrom(ptr unsafe.Pointer) hasShadow {
	return hasShadow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ hasShadow) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _hasShadowClass) Alloc() hasShadow {
	rv := objc.Send[hasShadow](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _hasShadowClass) New() hasShadow {
	rv := objc.Send[hasShadow](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhasShadow creates and returns a new initialized instance.
func NewhasShadow() hasShadow {
	return hasShadowClass.New()
}

// Init initializes the instance.
func (h_ hasShadow) Init() hasShadow {
	rv := objc.Send[hasShadow](h_.ID(), selInit)
	return rv
}
