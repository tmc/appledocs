
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [drawsBackground] class.
var drawsBackgroundClass _drawsBackgroundClass

func init() {
	drawsBackgroundClass = _drawsBackgroundClass{objc.GetClass("drawsBackground")}
}

type _drawsBackgroundClass struct {
	objc.Class
}

// An interface definition for the [drawsBackground] class.
type IdrawsBackground interface {
	ID() objc.ID
}

type drawsBackground struct {
	id objc.ID
}

func drawsBackgroundFrom(ptr unsafe.Pointer) drawsBackground {
	return drawsBackground{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ drawsBackground) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _drawsBackgroundClass) Alloc() drawsBackground {
	rv := objc.Send[drawsBackground](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _drawsBackgroundClass) New() drawsBackground {
	rv := objc.Send[drawsBackground](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdrawsBackground creates and returns a new initialized instance.
func NewdrawsBackground() drawsBackground {
	return drawsBackgroundClass.New()
}

// Init initializes the instance.
func (d_ drawsBackground) Init() drawsBackground {
	rv := objc.Send[drawsBackground](d_.ID(), selInit)
	return rv
}
