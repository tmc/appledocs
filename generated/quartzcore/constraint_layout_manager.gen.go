// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ConstraintLayoutManager] class.
var (
	constraintLayoutManagerClass     _ConstraintLayoutManagerClass
	constraintLayoutManagerClassOnce sync.Once
)

func getConstraintLayoutManagerClass() _ConstraintLayoutManagerClass {
	constraintLayoutManagerClassOnce.Do(func() {
		constraintLayoutManagerClass = _ConstraintLayoutManagerClass{objc.GetClass("CAConstraintLayoutManager")}
	})
	return constraintLayoutManagerClass
}

type _ConstraintLayoutManagerClass struct {
	class objc.Class
}

// An interface definition for the [ConstraintLayoutManager] class.
type IConstraintLayoutManager interface {
	objectivec.IObject
}

// An object that provides a constraint-based layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintLayoutManager
type ConstraintLayoutManager struct {
	objectivec.Object
}

// ConstraintLayoutManagerFrom constructs a [ConstraintLayoutManager] from an unsafe.Pointer.
//
// An object that provides a constraint-based layout manager.
func ConstraintLayoutManagerFrom(ptr unsafe.Pointer) ConstraintLayoutManager {
	return ConstraintLayoutManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ConstraintLayoutManagerClass) Alloc() ConstraintLayoutManager {
	rv := objc.Send[ConstraintLayoutManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ConstraintLayoutManagerClass) New() ConstraintLayoutManager {
	rv := objc.Send[ConstraintLayoutManager](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ConstraintLayoutManager) Init() ConstraintLayoutManager {
	rv := objc.Send[ConstraintLayoutManager](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ConstraintLayoutManager) Autorelease() ConstraintLayoutManager {
	rv := objc.Send[ConstraintLayoutManager](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConstraintLayoutManager creates a new ConstraintLayoutManager instance.
func NewConstraintLayoutManager() ConstraintLayoutManager {
	return getConstraintLayoutManagerClass().New()
}


// Returns the shared layout manager object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintLayoutManager/layoutManager
func (cc _ConstraintLayoutManagerClass) LayoutManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layoutManager"))
	return rv
}


