// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterTargetNavigator] class.
var (
	MTRBaseClusterTargetNavigatorClass     _MTRBaseClusterTargetNavigatorClass
	MTRBaseClusterTargetNavigatorClassOnce sync.Once
)

func getMTRBaseClusterTargetNavigatorClass() _MTRBaseClusterTargetNavigatorClass {
	MTRBaseClusterTargetNavigatorClassOnce.Do(func() {
		MTRBaseClusterTargetNavigatorClass = _MTRBaseClusterTargetNavigatorClass{objc.GetClass("MTRBaseClusterTargetNavigator")}
	})
	return MTRBaseClusterTargetNavigatorClass
}

type _MTRBaseClusterTargetNavigatorClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterTargetNavigator] class.
type IMTRBaseClusterTargetNavigator interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTargetNavigator
type MTRBaseClusterTargetNavigator struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterTargetNavigatorFrom constructs a [MTRBaseClusterTargetNavigator] from an unsafe.Pointer.
func MTRBaseClusterTargetNavigatorFrom(ptr unsafe.Pointer) MTRBaseClusterTargetNavigator {
	return MTRBaseClusterTargetNavigator{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterTargetNavigatorClass) Alloc() MTRBaseClusterTargetNavigator {
	rv := objc.Send[MTRBaseClusterTargetNavigator](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterTargetNavigatorClass) New() MTRBaseClusterTargetNavigator {
	rv := objc.Send[MTRBaseClusterTargetNavigator](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterTargetNavigator) Init() MTRBaseClusterTargetNavigator {
	rv := objc.Send[MTRBaseClusterTargetNavigator](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterTargetNavigator) Autorelease() MTRBaseClusterTargetNavigator {
	rv := objc.Send[MTRBaseClusterTargetNavigator](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterTargetNavigator creates a new MTRBaseClusterTargetNavigator instance.
func NewMTRBaseClusterTargetNavigator() MTRBaseClusterTargetNavigator {
	return getMTRBaseClusterTargetNavigatorClass().New()
}
