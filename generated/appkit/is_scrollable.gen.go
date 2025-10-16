
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isScrollable] class.
var isScrollableClass _isScrollableClass

func init() {
	isScrollableClass = _isScrollableClass{objc.GetClass("isScrollable")}
}

type _isScrollableClass struct {
	objc.Class
}

// An interface definition for the [isScrollable] class.
type IisScrollable interface {
	ID() objc.ID
}

type isScrollable struct {
	id objc.ID
}

func isScrollableFrom(ptr unsafe.Pointer) isScrollable {
	return isScrollable{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isScrollable) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isScrollableClass) Alloc() isScrollable {
	rv := objc.Send[isScrollable](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isScrollableClass) New() isScrollable {
	rv := objc.Send[isScrollable](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisScrollable creates and returns a new initialized instance.
func NewisScrollable() isScrollable {
	return isScrollableClass.New()
}

// Init initializes the instance.
func (i_ isScrollable) Init() isScrollable {
	rv := objc.Send[isScrollable](i_.ID(), selInit)
	return rv
}
