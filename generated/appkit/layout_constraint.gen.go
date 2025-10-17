
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutConstraint] class.
var LayoutConstraintClass _LayoutConstraintClass

func init() {
	LayoutConstraintClass = _LayoutConstraintClass{objc.GetClass("NSLayoutConstraint")}
}

type _LayoutConstraintClass struct {
	objc.Class
}

// An interface definition for the [LayoutConstraint] class.
type ILayoutConstraint interface {
	ID() objc.ID
}

type LayoutConstraint struct {
	id objc.ID
}

func LayoutConstraintFrom(ptr unsafe.Pointer) LayoutConstraint {
	return LayoutConstraint{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ LayoutConstraint) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutConstraintClass) Alloc() LayoutConstraint {
	rv := objc.Send[LayoutConstraint](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _LayoutConstraintClass) New() LayoutConstraint {
	rv := objc.Send[LayoutConstraint](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewLayoutConstraint creates and returns a new initialized instance.
func NewLayoutConstraint() LayoutConstraint {
	return LayoutConstraintClass.New()
}

// Init initializes the instance.
func (l_ LayoutConstraint) Init() LayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID(), selInit)
	return rv
}
