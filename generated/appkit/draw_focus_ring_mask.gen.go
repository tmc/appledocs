
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [drawFocusRingMask] class.
var drawFocusRingMaskClass _drawFocusRingMaskClass

func init() {
	drawFocusRingMaskClass = _drawFocusRingMaskClass{objc.GetClass("drawFocusRingMask")}
}

type _drawFocusRingMaskClass struct {
	objc.Class
}

// An interface definition for the [drawFocusRingMask] class.
type IdrawFocusRingMask interface {
	ID() objc.ID
}

type drawFocusRingMask struct {
	id objc.ID
}

func drawFocusRingMaskFrom(ptr unsafe.Pointer) drawFocusRingMask {
	return drawFocusRingMask{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ drawFocusRingMask) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _drawFocusRingMaskClass) Alloc() drawFocusRingMask {
	rv := objc.Send[drawFocusRingMask](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _drawFocusRingMaskClass) New() drawFocusRingMask {
	rv := objc.Send[drawFocusRingMask](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdrawFocusRingMask creates and returns a new initialized instance.
func NewdrawFocusRingMask() drawFocusRingMask {
	return drawFocusRingMaskClass.New()
}

// Init initializes the instance.
func (d_ drawFocusRingMask) Init() drawFocusRingMask {
	rv := objc.Send[drawFocusRingMask](d_.ID(), selInit)
	return rv
}
