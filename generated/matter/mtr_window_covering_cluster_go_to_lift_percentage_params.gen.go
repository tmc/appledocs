// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWindowCoveringClusterGoToLiftPercentageParams] class.
var (
	MTRWindowCoveringClusterGoToLiftPercentageParamsClass     _MTRWindowCoveringClusterGoToLiftPercentageParamsClass
	MTRWindowCoveringClusterGoToLiftPercentageParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterGoToLiftPercentageParamsClass() _MTRWindowCoveringClusterGoToLiftPercentageParamsClass {
	MTRWindowCoveringClusterGoToLiftPercentageParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterGoToLiftPercentageParamsClass = _MTRWindowCoveringClusterGoToLiftPercentageParamsClass{objc.GetClass("MTRWindowCoveringClusterGoToLiftPercentageParams")}
	})
	return MTRWindowCoveringClusterGoToLiftPercentageParamsClass
}

type _MTRWindowCoveringClusterGoToLiftPercentageParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWindowCoveringClusterGoToLiftPercentageParams] class.
type IMTRWindowCoveringClusterGoToLiftPercentageParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftPercentageParams
type MTRWindowCoveringClusterGoToLiftPercentageParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterGoToLiftPercentageParamsFrom constructs a [MTRWindowCoveringClusterGoToLiftPercentageParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterGoToLiftPercentageParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterGoToLiftPercentageParams {
	return MTRWindowCoveringClusterGoToLiftPercentageParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterGoToLiftPercentageParamsClass) Alloc() MTRWindowCoveringClusterGoToLiftPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftPercentageParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWindowCoveringClusterGoToLiftPercentageParamsClass) New() MTRWindowCoveringClusterGoToLiftPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftPercentageParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) Init() MTRWindowCoveringClusterGoToLiftPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftPercentageParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterGoToLiftPercentageParams) Autorelease() MTRWindowCoveringClusterGoToLiftPercentageParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftPercentageParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterGoToLiftPercentageParams creates a new MTRWindowCoveringClusterGoToLiftPercentageParams instance.
func NewMTRWindowCoveringClusterGoToLiftPercentageParams() MTRWindowCoveringClusterGoToLiftPercentageParams {
	return getMTRWindowCoveringClusterGoToLiftPercentageParamsClass().New()
}




