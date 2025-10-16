
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [edgeInsets] class.
var edgeInsetsClass _edgeInsetsClass

func init() {
	edgeInsetsClass = _edgeInsetsClass{objc.GetClass("edgeInsets")}
}

type _edgeInsetsClass struct {
	objc.Class
}

// An interface definition for the [edgeInsets] class.
type IedgeInsets interface {
	ID() objc.ID
}

type edgeInsets struct {
	id objc.ID
}

func edgeInsetsFrom(ptr unsafe.Pointer) edgeInsets {
	return edgeInsets{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ edgeInsets) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _edgeInsetsClass) Alloc() edgeInsets {
	rv := objc.Send[edgeInsets](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _edgeInsetsClass) New() edgeInsets {
	rv := objc.Send[edgeInsets](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewedgeInsets creates and returns a new initialized instance.
func NewedgeInsets() edgeInsets {
	return edgeInsetsClass.New()
}

// Init initializes the instance.
func (e_ edgeInsets) Init() edgeInsets {
	rv := objc.Send[edgeInsets](e_.ID(), selInit)
	return rv
}
