
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [minItemSize] class.
var minItemSizeClass _minItemSizeClass

func init() {
	minItemSizeClass = _minItemSizeClass{objc.GetClass("minItemSize")}
}

type _minItemSizeClass struct {
	objc.Class
}

// An interface definition for the [minItemSize] class.
type IminItemSize interface {
	ID() objc.ID
}

type minItemSize struct {
	id objc.ID
}

func minItemSizeFrom(ptr unsafe.Pointer) minItemSize {
	return minItemSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ minItemSize) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _minItemSizeClass) Alloc() minItemSize {
	rv := objc.Send[minItemSize](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _minItemSizeClass) New() minItemSize {
	rv := objc.Send[minItemSize](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewminItemSize creates and returns a new initialized instance.
func NewminItemSize() minItemSize {
	return minItemSizeClass.New()
}

// Init initializes the instance.
func (m_ minItemSize) Init() minItemSize {
	rv := objc.Send[minItemSize](m_.ID(), selInit)
	return rv
}
