// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterBooleanState] class.
var (
	MTRClusterBooleanStateClass     _MTRClusterBooleanStateClass
	MTRClusterBooleanStateClassOnce sync.Once
)

func getMTRClusterBooleanStateClass() _MTRClusterBooleanStateClass {
	MTRClusterBooleanStateClassOnce.Do(func() {
		MTRClusterBooleanStateClass = _MTRClusterBooleanStateClass{objc.GetClass("MTRClusterBooleanState")}
	})
	return MTRClusterBooleanStateClass
}

type _MTRClusterBooleanStateClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterBooleanState] class.
type IMTRClusterBooleanState interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterBooleanState
type MTRClusterBooleanState struct {
	MTRGenericCluster
}

// MTRClusterBooleanStateFrom constructs a [MTRClusterBooleanState] from an unsafe.Pointer.
func MTRClusterBooleanStateFrom(ptr unsafe.Pointer) MTRClusterBooleanState {
	return MTRClusterBooleanState{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterBooleanStateClass) Alloc() MTRClusterBooleanState {
	rv := objc.Send[MTRClusterBooleanState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterBooleanStateClass) New() MTRClusterBooleanState {
	rv := objc.Send[MTRClusterBooleanState](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterBooleanState) Init() MTRClusterBooleanState {
	rv := objc.Send[MTRClusterBooleanState](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterBooleanState) Autorelease() MTRClusterBooleanState {
	rv := objc.Send[MTRClusterBooleanState](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterBooleanState creates a new MTRClusterBooleanState instance.
func NewMTRClusterBooleanState() MTRClusterBooleanState {
	return getMTRClusterBooleanStateClass().New()
}




