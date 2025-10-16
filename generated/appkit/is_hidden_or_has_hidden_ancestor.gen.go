
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isHiddenOrHasHiddenAncestor] class.
var isHiddenOrHasHiddenAncestorClass _isHiddenOrHasHiddenAncestorClass

func init() {
	isHiddenOrHasHiddenAncestorClass = _isHiddenOrHasHiddenAncestorClass{objc.GetClass("isHiddenOrHasHiddenAncestor")}
}

type _isHiddenOrHasHiddenAncestorClass struct {
	objc.Class
}

// An interface definition for the [isHiddenOrHasHiddenAncestor] class.
type IisHiddenOrHasHiddenAncestor interface {
	ID() objc.ID
}

type isHiddenOrHasHiddenAncestor struct {
	id objc.ID
}

func isHiddenOrHasHiddenAncestorFrom(ptr unsafe.Pointer) isHiddenOrHasHiddenAncestor {
	return isHiddenOrHasHiddenAncestor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isHiddenOrHasHiddenAncestor) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isHiddenOrHasHiddenAncestorClass) Alloc() isHiddenOrHasHiddenAncestor {
	rv := objc.Send[isHiddenOrHasHiddenAncestor](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isHiddenOrHasHiddenAncestorClass) New() isHiddenOrHasHiddenAncestor {
	rv := objc.Send[isHiddenOrHasHiddenAncestor](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisHiddenOrHasHiddenAncestor creates and returns a new initialized instance.
func NewisHiddenOrHasHiddenAncestor() isHiddenOrHasHiddenAncestor {
	return isHiddenOrHasHiddenAncestorClass.New()
}

// Init initializes the instance.
func (i_ isHiddenOrHasHiddenAncestor) Init() isHiddenOrHasHiddenAncestor {
	rv := objc.Send[isHiddenOrHasHiddenAncestor](i_.ID(), selInit)
	return rv
}
