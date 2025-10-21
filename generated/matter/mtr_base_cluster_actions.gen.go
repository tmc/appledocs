// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterActions] class.
var (
	MTRBaseClusterActionsClass     _MTRBaseClusterActionsClass
	MTRBaseClusterActionsClassOnce sync.Once
)

func getMTRBaseClusterActionsClass() _MTRBaseClusterActionsClass {
	MTRBaseClusterActionsClassOnce.Do(func() {
		MTRBaseClusterActionsClass = _MTRBaseClusterActionsClass{objc.GetClass("MTRBaseClusterActions")}
	})
	return MTRBaseClusterActionsClass
}

type _MTRBaseClusterActionsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterActions] class.
type IMTRBaseClusterActions interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterActions
type MTRBaseClusterActions struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterActionsFrom constructs a [MTRBaseClusterActions] from an unsafe.Pointer.
func MTRBaseClusterActionsFrom(ptr unsafe.Pointer) MTRBaseClusterActions {
	return MTRBaseClusterActions{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterActionsClass) Alloc() MTRBaseClusterActions {
	rv := objc.Send[MTRBaseClusterActions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterActionsClass) New() MTRBaseClusterActions {
	rv := objc.Send[MTRBaseClusterActions](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterActions) Init() MTRBaseClusterActions {
	rv := objc.Send[MTRBaseClusterActions](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterActions) Autorelease() MTRBaseClusterActions {
	rv := objc.Send[MTRBaseClusterActions](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterActions creates a new MTRBaseClusterActions instance.
func NewMTRBaseClusterActions() MTRBaseClusterActions {
	return getMTRBaseClusterActionsClass().New()
}




