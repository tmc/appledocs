// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterActions] class.
var (
	MTRClusterActionsClass     _MTRClusterActionsClass
	MTRClusterActionsClassOnce sync.Once
)

func getMTRClusterActionsClass() _MTRClusterActionsClass {
	MTRClusterActionsClassOnce.Do(func() {
		MTRClusterActionsClass = _MTRClusterActionsClass{objc.GetClass("MTRClusterActions")}
	})
	return MTRClusterActionsClass
}

type _MTRClusterActionsClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterActions] class.
type IMTRClusterActions interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterActions
type MTRClusterActions struct {
	MTRGenericCluster
}

// MTRClusterActionsFrom constructs a [MTRClusterActions] from an unsafe.Pointer.
func MTRClusterActionsFrom(ptr unsafe.Pointer) MTRClusterActions {
	return MTRClusterActions{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterActionsClass) Alloc() MTRClusterActions {
	rv := objc.Send[MTRClusterActions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterActionsClass) New() MTRClusterActions {
	rv := objc.Send[MTRClusterActions](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterActions) Init() MTRClusterActions {
	rv := objc.Send[MTRClusterActions](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterActions) Autorelease() MTRClusterActions {
	rv := objc.Send[MTRClusterActions](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterActions creates a new MTRClusterActions instance.
func NewMTRClusterActions() MTRClusterActions {
	return getMTRClusterActionsClass().New()
}




