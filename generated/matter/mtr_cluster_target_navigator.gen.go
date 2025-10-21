// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterTargetNavigator] class.
var (
	MTRClusterTargetNavigatorClass     _MTRClusterTargetNavigatorClass
	MTRClusterTargetNavigatorClassOnce sync.Once
)

func getMTRClusterTargetNavigatorClass() _MTRClusterTargetNavigatorClass {
	MTRClusterTargetNavigatorClassOnce.Do(func() {
		MTRClusterTargetNavigatorClass = _MTRClusterTargetNavigatorClass{objc.GetClass("MTRClusterTargetNavigator")}
	})
	return MTRClusterTargetNavigatorClass
}

type _MTRClusterTargetNavigatorClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterTargetNavigator] class.
type IMTRClusterTargetNavigator interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTargetNavigator
type MTRClusterTargetNavigator struct {
	MTRGenericCluster
}

// MTRClusterTargetNavigatorFrom constructs a [MTRClusterTargetNavigator] from an unsafe.Pointer.
func MTRClusterTargetNavigatorFrom(ptr unsafe.Pointer) MTRClusterTargetNavigator {
	return MTRClusterTargetNavigator{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterTargetNavigatorClass) Alloc() MTRClusterTargetNavigator {
	rv := objc.Send[MTRClusterTargetNavigator](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterTargetNavigatorClass) New() MTRClusterTargetNavigator {
	rv := objc.Send[MTRClusterTargetNavigator](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterTargetNavigator) Init() MTRClusterTargetNavigator {
	rv := objc.Send[MTRClusterTargetNavigator](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterTargetNavigator) Autorelease() MTRClusterTargetNavigator {
	rv := objc.Send[MTRClusterTargetNavigator](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterTargetNavigator creates a new MTRClusterTargetNavigator instance.
func NewMTRClusterTargetNavigator() MTRClusterTargetNavigator {
	return getMTRClusterTargetNavigatorClass().New()
}




