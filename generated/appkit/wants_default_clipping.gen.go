
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [wantsDefaultClipping] class.
var wantsDefaultClippingClass _wantsDefaultClippingClass

func init() {
	wantsDefaultClippingClass = _wantsDefaultClippingClass{objc.GetClass("wantsDefaultClipping")}
}

type _wantsDefaultClippingClass struct {
	objc.Class
}

// An interface definition for the [wantsDefaultClipping] class.
type IwantsDefaultClipping interface {
	ID() objc.ID
}

type wantsDefaultClipping struct {
	id objc.ID
}

func wantsDefaultClippingFrom(ptr unsafe.Pointer) wantsDefaultClipping {
	return wantsDefaultClipping{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ wantsDefaultClipping) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _wantsDefaultClippingClass) Alloc() wantsDefaultClipping {
	rv := objc.Send[wantsDefaultClipping](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _wantsDefaultClippingClass) New() wantsDefaultClipping {
	rv := objc.Send[wantsDefaultClipping](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewwantsDefaultClipping creates and returns a new initialized instance.
func NewwantsDefaultClipping() wantsDefaultClipping {
	return wantsDefaultClippingClass.New()
}

// Init initializes the instance.
func (w_ wantsDefaultClipping) Init() wantsDefaultClipping {
	rv := objc.Send[wantsDefaultClipping](w_.ID(), selInit)
	return rv
}
