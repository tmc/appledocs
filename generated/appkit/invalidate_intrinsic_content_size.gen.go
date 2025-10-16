
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [invalidateIntrinsicContentSize] class.
var invalidateIntrinsicContentSizeClass _invalidateIntrinsicContentSizeClass

func init() {
	invalidateIntrinsicContentSizeClass = _invalidateIntrinsicContentSizeClass{objc.GetClass("invalidateIntrinsicContentSize")}
}

type _invalidateIntrinsicContentSizeClass struct {
	objc.Class
}

// An interface definition for the [invalidateIntrinsicContentSize] class.
type IinvalidateIntrinsicContentSize interface {
	ID() objc.ID
}

type invalidateIntrinsicContentSize struct {
	id objc.ID
}

func invalidateIntrinsicContentSizeFrom(ptr unsafe.Pointer) invalidateIntrinsicContentSize {
	return invalidateIntrinsicContentSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ invalidateIntrinsicContentSize) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _invalidateIntrinsicContentSizeClass) Alloc() invalidateIntrinsicContentSize {
	rv := objc.Send[invalidateIntrinsicContentSize](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _invalidateIntrinsicContentSizeClass) New() invalidateIntrinsicContentSize {
	rv := objc.Send[invalidateIntrinsicContentSize](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewinvalidateIntrinsicContentSize creates and returns a new initialized instance.
func NewinvalidateIntrinsicContentSize() invalidateIntrinsicContentSize {
	return invalidateIntrinsicContentSizeClass.New()
}

// Init initializes the instance.
func (i_ invalidateIntrinsicContentSize) Init() invalidateIntrinsicContentSize {
	rv := objc.Send[invalidateIntrinsicContentSize](i_.ID(), selInit)
	return rv
}
