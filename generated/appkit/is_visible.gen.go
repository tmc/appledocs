
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isVisible] class.
var isVisibleClass _isVisibleClass

func init() {
	isVisibleClass = _isVisibleClass{objc.GetClass("isVisible")}
}

type _isVisibleClass struct {
	objc.Class
}

// An interface definition for the [isVisible] class.
type IisVisible interface {
	ID() objc.ID
}

type isVisible struct {
	id objc.ID
}

func isVisibleFrom(ptr unsafe.Pointer) isVisible {
	return isVisible{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isVisible) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isVisibleClass) Alloc() isVisible {
	rv := objc.Send[isVisible](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isVisibleClass) New() isVisible {
	rv := objc.Send[isVisible](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisVisible creates and returns a new initialized instance.
func NewisVisible() isVisible {
	return isVisibleClass.New()
}

// Init initializes the instance.
func (i_ isVisible) Init() isVisible {
	rv := objc.Send[isVisible](i_.ID(), selInit)
	return rv
}
