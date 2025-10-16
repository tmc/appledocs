
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isOpaque] class.
var isOpaqueClass _isOpaqueClass

func init() {
	isOpaqueClass = _isOpaqueClass{objc.GetClass("isOpaque")}
}

type _isOpaqueClass struct {
	objc.Class
}

// An interface definition for the [isOpaque] class.
type IisOpaque interface {
	ID() objc.ID
}

type isOpaque struct {
	id objc.ID
}

func isOpaqueFrom(ptr unsafe.Pointer) isOpaque {
	return isOpaque{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isOpaque) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isOpaqueClass) Alloc() isOpaque {
	rv := objc.Send[isOpaque](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isOpaqueClass) New() isOpaque {
	rv := objc.Send[isOpaque](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisOpaque creates and returns a new initialized instance.
func NewisOpaque() isOpaque {
	return isOpaqueClass.New()
}

// Init initializes the instance.
func (i_ isOpaque) Init() isOpaque {
	rv := objc.Send[isOpaque](i_.ID(), selInit)
	return rv
}
