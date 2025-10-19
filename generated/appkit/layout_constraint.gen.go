// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LayoutConstraint] class.
var (
	layoutConstraintClass     _LayoutConstraintClass
	layoutConstraintClassOnce sync.Once
)

func getLayoutConstraintClass() _LayoutConstraintClass {
	layoutConstraintClassOnce.Do(func() {
		layoutConstraintClass = _LayoutConstraintClass{objc.GetClass("NSLayoutConstraint")}
	})
	return layoutConstraintClass
}

type _LayoutConstraintClass struct {
	class objc.Class
}

// An interface definition for the [LayoutConstraint] class.
type ILayoutConstraint interface {
	objectivec.IObject
}

// The relationship between two user interface objects that must be satisfied by the constraint-based layout system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint
type LayoutConstraint struct {
	objectivec.Object
}

// LayoutConstraintFrom constructs a [LayoutConstraint] from an unsafe.Pointer.
//
// The relationship between two user interface objects that must be satisfied by the constraint-based layout system.
func LayoutConstraintFrom(ptr unsafe.Pointer) LayoutConstraint {
	return LayoutConstraint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutConstraintClass) Alloc() LayoutConstraint {
	rv := objc.Send[LayoutConstraint](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LayoutConstraintClass) New() LayoutConstraint {
	rv := objc.Send[LayoutConstraint](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LayoutConstraint) Init() LayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LayoutConstraint) Autorelease() LayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayoutConstraint creates a new LayoutConstraint instance.
func NewLayoutConstraint() LayoutConstraint {
	return getLayoutConstraintClass().New()
}




