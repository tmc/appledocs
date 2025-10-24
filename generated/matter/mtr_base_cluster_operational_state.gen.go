// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterOperationalState] class.
var (
	MTRBaseClusterOperationalStateClass     _MTRBaseClusterOperationalStateClass
	MTRBaseClusterOperationalStateClassOnce sync.Once
)

func getMTRBaseClusterOperationalStateClass() _MTRBaseClusterOperationalStateClass {
	MTRBaseClusterOperationalStateClassOnce.Do(func() {
		MTRBaseClusterOperationalStateClass = _MTRBaseClusterOperationalStateClass{objc.GetClass("MTRBaseClusterOperationalState")}
	})
	return MTRBaseClusterOperationalStateClass
}

type _MTRBaseClusterOperationalStateClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterOperationalState] class.
type IMTRBaseClusterOperationalState interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOperationalState
type MTRBaseClusterOperationalState struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterOperationalStateFrom constructs a [MTRBaseClusterOperationalState] from an unsafe.Pointer.
func MTRBaseClusterOperationalStateFrom(ptr unsafe.Pointer) MTRBaseClusterOperationalState {
	return MTRBaseClusterOperationalState{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterOperationalStateClass) Alloc() MTRBaseClusterOperationalState {
	rv := objc.Send[MTRBaseClusterOperationalState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterOperationalStateClass) New() MTRBaseClusterOperationalState {
	rv := objc.Send[MTRBaseClusterOperationalState](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterOperationalState) Init() MTRBaseClusterOperationalState {
	rv := objc.Send[MTRBaseClusterOperationalState](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterOperationalState) Autorelease() MTRBaseClusterOperationalState {
	rv := objc.Send[MTRBaseClusterOperationalState](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterOperationalState creates a new MTRBaseClusterOperationalState instance.
func NewMTRBaseClusterOperationalState() MTRBaseClusterOperationalState {
	return getMTRBaseClusterOperationalStateClass().New()
}




