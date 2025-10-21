// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterRVCOperationalState] class.
var (
	MTRBaseClusterRVCOperationalStateClass     _MTRBaseClusterRVCOperationalStateClass
	MTRBaseClusterRVCOperationalStateClassOnce sync.Once
)

func getMTRBaseClusterRVCOperationalStateClass() _MTRBaseClusterRVCOperationalStateClass {
	MTRBaseClusterRVCOperationalStateClassOnce.Do(func() {
		MTRBaseClusterRVCOperationalStateClass = _MTRBaseClusterRVCOperationalStateClass{objc.GetClass("MTRBaseClusterRVCOperationalState")}
	})
	return MTRBaseClusterRVCOperationalStateClass
}

type _MTRBaseClusterRVCOperationalStateClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterRVCOperationalState] class.
type IMTRBaseClusterRVCOperationalState interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRVCOperationalState
type MTRBaseClusterRVCOperationalState struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterRVCOperationalStateFrom constructs a [MTRBaseClusterRVCOperationalState] from an unsafe.Pointer.
func MTRBaseClusterRVCOperationalStateFrom(ptr unsafe.Pointer) MTRBaseClusterRVCOperationalState {
	return MTRBaseClusterRVCOperationalState{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterRVCOperationalStateClass) Alloc() MTRBaseClusterRVCOperationalState {
	rv := objc.Send[MTRBaseClusterRVCOperationalState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterRVCOperationalStateClass) New() MTRBaseClusterRVCOperationalState {
	rv := objc.Send[MTRBaseClusterRVCOperationalState](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterRVCOperationalState) Init() MTRBaseClusterRVCOperationalState {
	rv := objc.Send[MTRBaseClusterRVCOperationalState](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterRVCOperationalState) Autorelease() MTRBaseClusterRVCOperationalState {
	rv := objc.Send[MTRBaseClusterRVCOperationalState](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterRVCOperationalState creates a new MTRBaseClusterRVCOperationalState instance.
func NewMTRBaseClusterRVCOperationalState() MTRBaseClusterRVCOperationalState {
	return getMTRBaseClusterRVCOperationalStateClass().New()
}




