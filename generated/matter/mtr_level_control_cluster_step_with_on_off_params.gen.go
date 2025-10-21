// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRLevelControlClusterStepWithOnOffParams] class.
var (
	MTRLevelControlClusterStepWithOnOffParamsClass     _MTRLevelControlClusterStepWithOnOffParamsClass
	MTRLevelControlClusterStepWithOnOffParamsClassOnce sync.Once
)

func getMTRLevelControlClusterStepWithOnOffParamsClass() _MTRLevelControlClusterStepWithOnOffParamsClass {
	MTRLevelControlClusterStepWithOnOffParamsClassOnce.Do(func() {
		MTRLevelControlClusterStepWithOnOffParamsClass = _MTRLevelControlClusterStepWithOnOffParamsClass{objc.GetClass("MTRLevelControlClusterStepWithOnOffParams")}
	})
	return MTRLevelControlClusterStepWithOnOffParamsClass
}

type _MTRLevelControlClusterStepWithOnOffParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterStepWithOnOffParams] class.
type IMTRLevelControlClusterStepWithOnOffParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStepWithOnOffParams
type MTRLevelControlClusterStepWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterStepWithOnOffParamsFrom constructs a [MTRLevelControlClusterStepWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterStepWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterStepWithOnOffParams {
	return MTRLevelControlClusterStepWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterStepWithOnOffParamsClass) Alloc() MTRLevelControlClusterStepWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStepWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterStepWithOnOffParamsClass) New() MTRLevelControlClusterStepWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStepWithOnOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterStepWithOnOffParams) Init() MTRLevelControlClusterStepWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStepWithOnOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterStepWithOnOffParams) Autorelease() MTRLevelControlClusterStepWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStepWithOnOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterStepWithOnOffParams creates a new MTRLevelControlClusterStepWithOnOffParams instance.
func NewMTRLevelControlClusterStepWithOnOffParams() MTRLevelControlClusterStepWithOnOffParams {
	return getMTRLevelControlClusterStepWithOnOffParamsClass().New()
}




