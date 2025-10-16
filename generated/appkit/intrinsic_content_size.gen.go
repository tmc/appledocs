
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [intrinsicContentSize] class.
var intrinsicContentSizeClass _intrinsicContentSizeClass

func init() {
	intrinsicContentSizeClass = _intrinsicContentSizeClass{objc.GetClass("intrinsicContentSize")}
}

type _intrinsicContentSizeClass struct {
	objc.Class
}

// An interface definition for the [intrinsicContentSize] class.
type IintrinsicContentSize interface {
	ID() objc.ID
}

type intrinsicContentSize struct {
	id objc.ID
}

func intrinsicContentSizeFrom(ptr unsafe.Pointer) intrinsicContentSize {
	return intrinsicContentSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ intrinsicContentSize) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _intrinsicContentSizeClass) Alloc() intrinsicContentSize {
	rv := objc.Send[intrinsicContentSize](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _intrinsicContentSizeClass) New() intrinsicContentSize {
	rv := objc.Send[intrinsicContentSize](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewintrinsicContentSize creates and returns a new initialized instance.
func NewintrinsicContentSize() intrinsicContentSize {
	return intrinsicContentSizeClass.New()
}

// Init initializes the instance.
func (i_ intrinsicContentSize) Init() intrinsicContentSize {
	rv := objc.Send[intrinsicContentSize](i_.ID(), selInit)
	return rv
}
