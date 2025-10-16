
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [maxSize] class.
var maxSizeClass _maxSizeClass

func init() {
	maxSizeClass = _maxSizeClass{objc.GetClass("maxSize")}
}

type _maxSizeClass struct {
	objc.Class
}

// An interface definition for the [maxSize] class.
type ImaxSize interface {
	ID() objc.ID
}

type maxSize struct {
	id objc.ID
}

func maxSizeFrom(ptr unsafe.Pointer) maxSize {
	return maxSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ maxSize) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _maxSizeClass) Alloc() maxSize {
	rv := objc.Send[maxSize](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _maxSizeClass) New() maxSize {
	rv := objc.Send[maxSize](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmaxSize creates and returns a new initialized instance.
func NewmaxSize() maxSize {
	return maxSizeClass.New()
}

// Init initializes the instance.
func (m_ maxSize) Init() maxSize {
	rv := objc.Send[maxSize](m_.ID(), selInit)
	return rv
}
