// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterGroups] class.
var (
	MTRBaseClusterGroupsClass     _MTRBaseClusterGroupsClass
	MTRBaseClusterGroupsClassOnce sync.Once
)

func getMTRBaseClusterGroupsClass() _MTRBaseClusterGroupsClass {
	MTRBaseClusterGroupsClassOnce.Do(func() {
		MTRBaseClusterGroupsClass = _MTRBaseClusterGroupsClass{objc.GetClass("MTRBaseClusterGroups")}
	})
	return MTRBaseClusterGroupsClass
}

type _MTRBaseClusterGroupsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterGroups] class.
type IMTRBaseClusterGroups interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterGroups
type MTRBaseClusterGroups struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterGroupsFrom constructs a [MTRBaseClusterGroups] from an unsafe.Pointer.
func MTRBaseClusterGroupsFrom(ptr unsafe.Pointer) MTRBaseClusterGroups {
	return MTRBaseClusterGroups{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterGroupsClass) Alloc() MTRBaseClusterGroups {
	rv := objc.Send[MTRBaseClusterGroups](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterGroupsClass) New() MTRBaseClusterGroups {
	rv := objc.Send[MTRBaseClusterGroups](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterGroups) Init() MTRBaseClusterGroups {
	rv := objc.Send[MTRBaseClusterGroups](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterGroups) Autorelease() MTRBaseClusterGroups {
	rv := objc.Send[MTRBaseClusterGroups](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterGroups creates a new MTRBaseClusterGroups instance.
func NewMTRBaseClusterGroups() MTRBaseClusterGroups {
	return getMTRBaseClusterGroupsClass().New()
}




