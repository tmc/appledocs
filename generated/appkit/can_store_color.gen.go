
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [canStoreColor] class.
var canStoreColorClass _canStoreColorClass

func init() {
	canStoreColorClass = _canStoreColorClass{objc.GetClass("canStoreColor")}
}

type _canStoreColorClass struct {
	objc.Class
}

// An interface definition for the [canStoreColor] class.
type IcanStoreColor interface {
	ID() objc.ID
}

type canStoreColor struct {
	id objc.ID
}

func canStoreColorFrom(ptr unsafe.Pointer) canStoreColor {
	return canStoreColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ canStoreColor) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _canStoreColorClass) Alloc() canStoreColor {
	rv := objc.Send[canStoreColor](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _canStoreColorClass) New() canStoreColor {
	rv := objc.Send[canStoreColor](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcanStoreColor creates and returns a new initialized instance.
func NewcanStoreColor() canStoreColor {
	return canStoreColorClass.New()
}

// Init initializes the instance.
func (c_ canStoreColor) Init() canStoreColor {
	rv := objc.Send[canStoreColor](c_.ID(), selInit)
	return rv
}
