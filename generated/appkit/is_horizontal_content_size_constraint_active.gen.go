
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isHorizontalContentSizeConstraintActive] class.
var isHorizontalContentSizeConstraintActiveClass _isHorizontalContentSizeConstraintActiveClass

func init() {
	isHorizontalContentSizeConstraintActiveClass = _isHorizontalContentSizeConstraintActiveClass{objc.GetClass("isHorizontalContentSizeConstraintActive")}
}

type _isHorizontalContentSizeConstraintActiveClass struct {
	objc.Class
}

// An interface definition for the [isHorizontalContentSizeConstraintActive] class.
type IisHorizontalContentSizeConstraintActive interface {
	ID() objc.ID
}

type isHorizontalContentSizeConstraintActive struct {
	id objc.ID
}

func isHorizontalContentSizeConstraintActiveFrom(ptr unsafe.Pointer) isHorizontalContentSizeConstraintActive {
	return isHorizontalContentSizeConstraintActive{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isHorizontalContentSizeConstraintActive) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isHorizontalContentSizeConstraintActiveClass) Alloc() isHorizontalContentSizeConstraintActive {
	rv := objc.Send[isHorizontalContentSizeConstraintActive](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isHorizontalContentSizeConstraintActiveClass) New() isHorizontalContentSizeConstraintActive {
	rv := objc.Send[isHorizontalContentSizeConstraintActive](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisHorizontalContentSizeConstraintActive creates and returns a new initialized instance.
func NewisHorizontalContentSizeConstraintActive() isHorizontalContentSizeConstraintActive {
	return isHorizontalContentSizeConstraintActiveClass.New()
}

// Init initializes the instance.
func (i_ isHorizontalContentSizeConstraintActive) Init() isHorizontalContentSizeConstraintActive {
	rv := objc.Send[isHorizontalContentSizeConstraintActive](i_.ID(), selInit)
	return rv
}
