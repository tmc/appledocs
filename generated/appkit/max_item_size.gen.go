
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [maxItemSize] class.
var maxItemSizeClass _maxItemSizeClass

func init() {
	maxItemSizeClass = _maxItemSizeClass{objc.GetClass("maxItemSize")}
}

type _maxItemSizeClass struct {
	objc.Class
}

// An interface definition for the [maxItemSize] class.
type ImaxItemSize interface {
	ID() objc.ID
}

type maxItemSize struct {
	id objc.ID
}

func maxItemSizeFrom(ptr unsafe.Pointer) maxItemSize {
	return maxItemSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ maxItemSize) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _maxItemSizeClass) Alloc() maxItemSize {
	rv := objc.Send[maxItemSize](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _maxItemSizeClass) New() maxItemSize {
	rv := objc.Send[maxItemSize](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmaxItemSize creates and returns a new initialized instance.
func NewmaxItemSize() maxItemSize {
	return maxItemSizeClass.New()
}

// Init initializes the instance.
func (m_ maxItemSize) Init() maxItemSize {
	rv := objc.Send[maxItemSize](m_.ID(), selInit)
	return rv
}
