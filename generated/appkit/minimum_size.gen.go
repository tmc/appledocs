
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [minimumSize] class.
var minimumSizeClass _minimumSizeClass

func init() {
	minimumSizeClass = _minimumSizeClass{objc.GetClass("minimumSize")}
}

type _minimumSizeClass struct {
	objc.Class
}

// An interface definition for the [minimumSize] class.
type IminimumSize interface {
	ID() objc.ID
}

type minimumSize struct {
	id objc.ID
}

func minimumSizeFrom(ptr unsafe.Pointer) minimumSize {
	return minimumSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ minimumSize) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _minimumSizeClass) Alloc() minimumSize {
	rv := objc.Send[minimumSize](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _minimumSizeClass) New() minimumSize {
	rv := objc.Send[minimumSize](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewminimumSize creates and returns a new initialized instance.
func NewminimumSize() minimumSize {
	return minimumSizeClass.New()
}

// Init initializes the instance.
func (m_ minimumSize) Init() minimumSize {
	rv := objc.Send[minimumSize](m_.ID(), selInit)
	return rv
}
