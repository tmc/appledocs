// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterOperationalState] class.
var (
	MTRClusterOperationalStateClass     _MTRClusterOperationalStateClass
	MTRClusterOperationalStateClassOnce sync.Once
)

func getMTRClusterOperationalStateClass() _MTRClusterOperationalStateClass {
	MTRClusterOperationalStateClassOnce.Do(func() {
		MTRClusterOperationalStateClass = _MTRClusterOperationalStateClass{objc.GetClass("MTRClusterOperationalState")}
	})
	return MTRClusterOperationalStateClass
}

type _MTRClusterOperationalStateClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterOperationalState] class.
type IMTRClusterOperationalState interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOperationalState
type MTRClusterOperationalState struct {
	MTRGenericCluster
}

// MTRClusterOperationalStateFrom constructs a [MTRClusterOperationalState] from an unsafe.Pointer.
func MTRClusterOperationalStateFrom(ptr unsafe.Pointer) MTRClusterOperationalState {
	return MTRClusterOperationalState{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterOperationalStateClass) Alloc() MTRClusterOperationalState {
	rv := objc.Send[MTRClusterOperationalState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterOperationalStateClass) New() MTRClusterOperationalState {
	rv := objc.Send[MTRClusterOperationalState](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterOperationalState) Init() MTRClusterOperationalState {
	rv := objc.Send[MTRClusterOperationalState](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterOperationalState) Autorelease() MTRClusterOperationalState {
	rv := objc.Send[MTRClusterOperationalState](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterOperationalState creates a new MTRClusterOperationalState instance.
func NewMTRClusterOperationalState() MTRClusterOperationalState {
	return getMTRClusterOperationalStateClass().New()
}




