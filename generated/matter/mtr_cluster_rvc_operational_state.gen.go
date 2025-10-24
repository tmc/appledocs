// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterRVCOperationalState] class.
var (
	MTRClusterRVCOperationalStateClass     _MTRClusterRVCOperationalStateClass
	MTRClusterRVCOperationalStateClassOnce sync.Once
)

func getMTRClusterRVCOperationalStateClass() _MTRClusterRVCOperationalStateClass {
	MTRClusterRVCOperationalStateClassOnce.Do(func() {
		MTRClusterRVCOperationalStateClass = _MTRClusterRVCOperationalStateClass{objc.GetClass("MTRClusterRVCOperationalState")}
	})
	return MTRClusterRVCOperationalStateClass
}

type _MTRClusterRVCOperationalStateClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterRVCOperationalState] class.
type IMTRClusterRVCOperationalState interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRVCOperationalState
type MTRClusterRVCOperationalState struct {
	MTRGenericCluster
}

// MTRClusterRVCOperationalStateFrom constructs a [MTRClusterRVCOperationalState] from an unsafe.Pointer.
func MTRClusterRVCOperationalStateFrom(ptr unsafe.Pointer) MTRClusterRVCOperationalState {
	return MTRClusterRVCOperationalState{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterRVCOperationalStateClass) Alloc() MTRClusterRVCOperationalState {
	rv := objc.Send[MTRClusterRVCOperationalState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterRVCOperationalStateClass) New() MTRClusterRVCOperationalState {
	rv := objc.Send[MTRClusterRVCOperationalState](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterRVCOperationalState) Init() MTRClusterRVCOperationalState {
	rv := objc.Send[MTRClusterRVCOperationalState](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterRVCOperationalState) Autorelease() MTRClusterRVCOperationalState {
	rv := objc.Send[MTRClusterRVCOperationalState](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterRVCOperationalState creates a new MTRClusterRVCOperationalState instance.
func NewMTRClusterRVCOperationalState() MTRClusterRVCOperationalState {
	return getMTRClusterRVCOperationalStateClass().New()
}




