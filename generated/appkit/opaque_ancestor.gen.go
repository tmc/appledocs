
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [opaqueAncestor] class.
var opaqueAncestorClass _opaqueAncestorClass

func init() {
	opaqueAncestorClass = _opaqueAncestorClass{objc.GetClass("opaqueAncestor")}
}

type _opaqueAncestorClass struct {
	objc.Class
}

// An interface definition for the [opaqueAncestor] class.
type IopaqueAncestor interface {
	ID() objc.ID
}

type opaqueAncestor struct {
	id objc.ID
}

func opaqueAncestorFrom(ptr unsafe.Pointer) opaqueAncestor {
	return opaqueAncestor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ opaqueAncestor) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _opaqueAncestorClass) Alloc() opaqueAncestor {
	rv := objc.Send[opaqueAncestor](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _opaqueAncestorClass) New() opaqueAncestor {
	rv := objc.Send[opaqueAncestor](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewopaqueAncestor creates and returns a new initialized instance.
func NewopaqueAncestor() opaqueAncestor {
	return opaqueAncestorClass.New()
}

// Init initializes the instance.
func (o_ opaqueAncestor) Init() opaqueAncestor {
	rv := objc.Send[opaqueAncestor](o_.ID(), selInit)
	return rv
}
