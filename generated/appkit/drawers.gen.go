
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [drawers] class.
var drawersClass _drawersClass

func init() {
	drawersClass = _drawersClass{objc.GetClass("drawers")}
}

type _drawersClass struct {
	objc.Class
}

// An interface definition for the [drawers] class.
type Idrawers interface {
	ID() objc.ID
}

type drawers struct {
	id objc.ID
}

func drawersFrom(ptr unsafe.Pointer) drawers {
	return drawers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ drawers) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _drawersClass) Alloc() drawers {
	rv := objc.Send[drawers](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _drawersClass) New() drawers {
	rv := objc.Send[drawers](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newdrawers creates and returns a new initialized instance.
func Newdrawers() drawers {
	return drawersClass.New()
}

// Init initializes the instance.
func (d_ drawers) Init() drawers {
	rv := objc.Send[drawers](d_.ID(), selInit)
	return rv
}
