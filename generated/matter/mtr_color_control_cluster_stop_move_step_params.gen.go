// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRColorControlClusterStopMoveStepParams] class.
var (
	MTRColorControlClusterStopMoveStepParamsClass     _MTRColorControlClusterStopMoveStepParamsClass
	MTRColorControlClusterStopMoveStepParamsClassOnce sync.Once
)

func getMTRColorControlClusterStopMoveStepParamsClass() _MTRColorControlClusterStopMoveStepParamsClass {
	MTRColorControlClusterStopMoveStepParamsClassOnce.Do(func() {
		MTRColorControlClusterStopMoveStepParamsClass = _MTRColorControlClusterStopMoveStepParamsClass{objc.GetClass("MTRColorControlClusterStopMoveStepParams")}
	})
	return MTRColorControlClusterStopMoveStepParamsClass
}

type _MTRColorControlClusterStopMoveStepParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterStopMoveStepParams] class.
type IMTRColorControlClusterStopMoveStepParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStopMoveStepParams
type MTRColorControlClusterStopMoveStepParams struct {
	objectivec.Object
}

// MTRColorControlClusterStopMoveStepParamsFrom constructs a [MTRColorControlClusterStopMoveStepParams] from an unsafe.Pointer.
func MTRColorControlClusterStopMoveStepParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterStopMoveStepParams {
	return MTRColorControlClusterStopMoveStepParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterStopMoveStepParamsClass) Alloc() MTRColorControlClusterStopMoveStepParams {
	rv := objc.Send[MTRColorControlClusterStopMoveStepParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterStopMoveStepParamsClass) New() MTRColorControlClusterStopMoveStepParams {
	rv := objc.Send[MTRColorControlClusterStopMoveStepParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterStopMoveStepParams) Init() MTRColorControlClusterStopMoveStepParams {
	rv := objc.Send[MTRColorControlClusterStopMoveStepParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterStopMoveStepParams) Autorelease() MTRColorControlClusterStopMoveStepParams {
	rv := objc.Send[MTRColorControlClusterStopMoveStepParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterStopMoveStepParams creates a new MTRColorControlClusterStopMoveStepParams instance.
func NewMTRColorControlClusterStopMoveStepParams() MTRColorControlClusterStopMoveStepParams {
	return getMTRColorControlClusterStopMoveStepParamsClass().New()
}




