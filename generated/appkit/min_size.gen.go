
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [minSize] class.
var minSizeClass _minSizeClass

func init() {
	minSizeClass = _minSizeClass{objc.GetClass("minSize")}
}

type _minSizeClass struct {
	objc.Class
}

// An interface definition for the [minSize] class.
type IminSize interface {
	ID() objc.ID
}

type minSize struct {
	id objc.ID
}

func minSizeFrom(ptr unsafe.Pointer) minSize {
	return minSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ minSize) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _minSizeClass) Alloc() minSize {
	rv := objc.Send[minSize](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _minSizeClass) New() minSize {
	rv := objc.Send[minSize](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewminSize creates and returns a new initialized instance.
func NewminSize() minSize {
	return minSizeClass.New()
}

// Init initializes the instance.
func (m_ minSize) Init() minSize {
	rv := objc.Send[minSize](m_.ID(), selInit)
	return rv
}
