// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterBooleanState] class.
var (
	MTRBaseClusterBooleanStateClass     _MTRBaseClusterBooleanStateClass
	MTRBaseClusterBooleanStateClassOnce sync.Once
)

func getMTRBaseClusterBooleanStateClass() _MTRBaseClusterBooleanStateClass {
	MTRBaseClusterBooleanStateClassOnce.Do(func() {
		MTRBaseClusterBooleanStateClass = _MTRBaseClusterBooleanStateClass{objc.GetClass("MTRBaseClusterBooleanState")}
	})
	return MTRBaseClusterBooleanStateClass
}

type _MTRBaseClusterBooleanStateClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterBooleanState] class.
type IMTRBaseClusterBooleanState interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterBooleanState
type MTRBaseClusterBooleanState struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterBooleanStateFrom constructs a [MTRBaseClusterBooleanState] from an unsafe.Pointer.
func MTRBaseClusterBooleanStateFrom(ptr unsafe.Pointer) MTRBaseClusterBooleanState {
	return MTRBaseClusterBooleanState{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterBooleanStateClass) Alloc() MTRBaseClusterBooleanState {
	rv := objc.Send[MTRBaseClusterBooleanState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterBooleanStateClass) New() MTRBaseClusterBooleanState {
	rv := objc.Send[MTRBaseClusterBooleanState](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterBooleanState) Init() MTRBaseClusterBooleanState {
	rv := objc.Send[MTRBaseClusterBooleanState](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterBooleanState) Autorelease() MTRBaseClusterBooleanState {
	rv := objc.Send[MTRBaseClusterBooleanState](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterBooleanState creates a new MTRBaseClusterBooleanState instance.
func NewMTRBaseClusterBooleanState() MTRBaseClusterBooleanState {
	return getMTRBaseClusterBooleanStateClass().New()
}
