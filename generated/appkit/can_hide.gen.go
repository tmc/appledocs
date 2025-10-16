
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [canHide] class.
var canHideClass _canHideClass

func init() {
	canHideClass = _canHideClass{objc.GetClass("canHide")}
}

type _canHideClass struct {
	objc.Class
}

// An interface definition for the [canHide] class.
type IcanHide interface {
	ID() objc.ID
}

type canHide struct {
	id objc.ID
}

func canHideFrom(ptr unsafe.Pointer) canHide {
	return canHide{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ canHide) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _canHideClass) Alloc() canHide {
	rv := objc.Send[canHide](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _canHideClass) New() canHide {
	rv := objc.Send[canHide](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcanHide creates and returns a new initialized instance.
func NewcanHide() canHide {
	return canHideClass.New()
}

// Init initializes the instance.
func (c_ canHide) Init() canHide {
	rv := objc.Send[canHide](c_.ID(), selInit)
	return rv
}
