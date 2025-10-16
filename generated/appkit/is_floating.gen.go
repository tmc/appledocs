
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isFloating] class.
var isFloatingClass _isFloatingClass

func init() {
	isFloatingClass = _isFloatingClass{objc.GetClass("isFloating")}
}

type _isFloatingClass struct {
	objc.Class
}

// An interface definition for the [isFloating] class.
type IisFloating interface {
	ID() objc.ID
}

type isFloating struct {
	id objc.ID
}

func isFloatingFrom(ptr unsafe.Pointer) isFloating {
	return isFloating{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isFloating) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isFloatingClass) Alloc() isFloating {
	rv := objc.Send[isFloating](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isFloatingClass) New() isFloating {
	rv := objc.Send[isFloating](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisFloating creates and returns a new initialized instance.
func NewisFloating() isFloating {
	return isFloatingClass.New()
}

// Init initializes the instance.
func (i_ isFloating) Init() isFloating {
	rv := objc.Send[isFloating](i_.ID(), selInit)
	return rv
}
