
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isVerticalContentSizeConstraintActive] class.
var isVerticalContentSizeConstraintActiveClass _isVerticalContentSizeConstraintActiveClass

func init() {
	isVerticalContentSizeConstraintActiveClass = _isVerticalContentSizeConstraintActiveClass{objc.GetClass("isVerticalContentSizeConstraintActive")}
}

type _isVerticalContentSizeConstraintActiveClass struct {
	objc.Class
}

// An interface definition for the [isVerticalContentSizeConstraintActive] class.
type IisVerticalContentSizeConstraintActive interface {
	ID() objc.ID
}

type isVerticalContentSizeConstraintActive struct {
	id objc.ID
}

func isVerticalContentSizeConstraintActiveFrom(ptr unsafe.Pointer) isVerticalContentSizeConstraintActive {
	return isVerticalContentSizeConstraintActive{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isVerticalContentSizeConstraintActive) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isVerticalContentSizeConstraintActiveClass) Alloc() isVerticalContentSizeConstraintActive {
	rv := objc.Send[isVerticalContentSizeConstraintActive](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isVerticalContentSizeConstraintActiveClass) New() isVerticalContentSizeConstraintActive {
	rv := objc.Send[isVerticalContentSizeConstraintActive](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisVerticalContentSizeConstraintActive creates and returns a new initialized instance.
func NewisVerticalContentSizeConstraintActive() isVerticalContentSizeConstraintActive {
	return isVerticalContentSizeConstraintActiveClass.New()
}

// Init initializes the instance.
func (i_ isVerticalContentSizeConstraintActive) Init() isVerticalContentSizeConstraintActive {
	rv := objc.Send[isVerticalContentSizeConstraintActive](i_.ID(), selInit)
	return rv
}
