
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isMovable] class.
var isMovableClass _isMovableClass

func init() {
	isMovableClass = _isMovableClass{objc.GetClass("isMovable")}
}

type _isMovableClass struct {
	objc.Class
}

// An interface definition for the [isMovable] class.
type IisMovable interface {
	ID() objc.ID
}

type isMovable struct {
	id objc.ID
}

func isMovableFrom(ptr unsafe.Pointer) isMovable {
	return isMovable{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isMovable) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isMovableClass) Alloc() isMovable {
	rv := objc.Send[isMovable](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isMovableClass) New() isMovable {
	rv := objc.Send[isMovable](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisMovable creates and returns a new initialized instance.
func NewisMovable() isMovable {
	return isMovableClass.New()
}

// Init initializes the instance.
func (i_ isMovable) Init() isMovable {
	rv := objc.Send[isMovable](i_.ID(), selInit)
	return rv
}
