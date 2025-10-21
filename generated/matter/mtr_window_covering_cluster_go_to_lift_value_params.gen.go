// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRWindowCoveringClusterGoToLiftValueParams] class.
var (
	MTRWindowCoveringClusterGoToLiftValueParamsClass     _MTRWindowCoveringClusterGoToLiftValueParamsClass
	MTRWindowCoveringClusterGoToLiftValueParamsClassOnce sync.Once
)

func getMTRWindowCoveringClusterGoToLiftValueParamsClass() _MTRWindowCoveringClusterGoToLiftValueParamsClass {
	MTRWindowCoveringClusterGoToLiftValueParamsClassOnce.Do(func() {
		MTRWindowCoveringClusterGoToLiftValueParamsClass = _MTRWindowCoveringClusterGoToLiftValueParamsClass{objc.GetClass("MTRWindowCoveringClusterGoToLiftValueParams")}
	})
	return MTRWindowCoveringClusterGoToLiftValueParamsClass
}

type _MTRWindowCoveringClusterGoToLiftValueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWindowCoveringClusterGoToLiftValueParams] class.
type IMTRWindowCoveringClusterGoToLiftValueParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWindowCoveringClusterGoToLiftValueParams
type MTRWindowCoveringClusterGoToLiftValueParams struct {
	objectivec.Object
}

// MTRWindowCoveringClusterGoToLiftValueParamsFrom constructs a [MTRWindowCoveringClusterGoToLiftValueParams] from an unsafe.Pointer.
func MTRWindowCoveringClusterGoToLiftValueParamsFrom(ptr unsafe.Pointer) MTRWindowCoveringClusterGoToLiftValueParams {
	return MTRWindowCoveringClusterGoToLiftValueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWindowCoveringClusterGoToLiftValueParamsClass) Alloc() MTRWindowCoveringClusterGoToLiftValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftValueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWindowCoveringClusterGoToLiftValueParamsClass) New() MTRWindowCoveringClusterGoToLiftValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftValueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) Init() MTRWindowCoveringClusterGoToLiftValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftValueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWindowCoveringClusterGoToLiftValueParams) Autorelease() MTRWindowCoveringClusterGoToLiftValueParams {
	rv := objc.Send[MTRWindowCoveringClusterGoToLiftValueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWindowCoveringClusterGoToLiftValueParams creates a new MTRWindowCoveringClusterGoToLiftValueParams instance.
func NewMTRWindowCoveringClusterGoToLiftValueParams() MTRWindowCoveringClusterGoToLiftValueParams {
	return getMTRWindowCoveringClusterGoToLiftValueParamsClass().New()
}




