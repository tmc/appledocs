// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterGroups] class.
var (
	MTRClusterGroupsClass     _MTRClusterGroupsClass
	MTRClusterGroupsClassOnce sync.Once
)

func getMTRClusterGroupsClass() _MTRClusterGroupsClass {
	MTRClusterGroupsClassOnce.Do(func() {
		MTRClusterGroupsClass = _MTRClusterGroupsClass{objc.GetClass("MTRClusterGroups")}
	})
	return MTRClusterGroupsClass
}

type _MTRClusterGroupsClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterGroups] class.
type IMTRClusterGroups interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterGroups
type MTRClusterGroups struct {
	MTRGenericCluster
}

// MTRClusterGroupsFrom constructs a [MTRClusterGroups] from an unsafe.Pointer.
func MTRClusterGroupsFrom(ptr unsafe.Pointer) MTRClusterGroups {
	return MTRClusterGroups{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterGroupsClass) Alloc() MTRClusterGroups {
	rv := objc.Send[MTRClusterGroups](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterGroupsClass) New() MTRClusterGroups {
	rv := objc.Send[MTRClusterGroups](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterGroups) Init() MTRClusterGroups {
	rv := objc.Send[MTRClusterGroups](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterGroups) Autorelease() MTRClusterGroups {
	rv := objc.Send[MTRClusterGroups](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterGroups creates a new MTRClusterGroups instance.
func NewMTRClusterGroups() MTRClusterGroups {
	return getMTRClusterGroupsClass().New()
}
