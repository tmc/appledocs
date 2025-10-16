
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isZoomable] class.
var isZoomableClass _isZoomableClass

func init() {
	isZoomableClass = _isZoomableClass{objc.GetClass("isZoomable")}
}

type _isZoomableClass struct {
	objc.Class
}

// An interface definition for the [isZoomable] class.
type IisZoomable interface {
	ID() objc.ID
}

type isZoomable struct {
	id objc.ID
}

func isZoomableFrom(ptr unsafe.Pointer) isZoomable {
	return isZoomable{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isZoomable) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isZoomableClass) Alloc() isZoomable {
	rv := objc.Send[isZoomable](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isZoomableClass) New() isZoomable {
	rv := objc.Send[isZoomable](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisZoomable creates and returns a new initialized instance.
func NewisZoomable() isZoomable {
	return isZoomableClass.New()
}

// Init initializes the instance.
func (i_ isZoomable) Init() isZoomable {
	rv := objc.Send[isZoomable](i_.ID(), selInit)
	return rv
}
