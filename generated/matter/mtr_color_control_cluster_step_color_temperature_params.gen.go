// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterStepColorTemperatureParams] class.
var (
	MTRColorControlClusterStepColorTemperatureParamsClass     _MTRColorControlClusterStepColorTemperatureParamsClass
	MTRColorControlClusterStepColorTemperatureParamsClassOnce sync.Once
)

func getMTRColorControlClusterStepColorTemperatureParamsClass() _MTRColorControlClusterStepColorTemperatureParamsClass {
	MTRColorControlClusterStepColorTemperatureParamsClassOnce.Do(func() {
		MTRColorControlClusterStepColorTemperatureParamsClass = _MTRColorControlClusterStepColorTemperatureParamsClass{objc.GetClass("MTRColorControlClusterStepColorTemperatureParams")}
	})
	return MTRColorControlClusterStepColorTemperatureParamsClass
}

type _MTRColorControlClusterStepColorTemperatureParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterStepColorTemperatureParams] class.
type IMTRColorControlClusterStepColorTemperatureParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepColorTemperatureParams
type MTRColorControlClusterStepColorTemperatureParams struct {
	objectivec.Object
}

// MTRColorControlClusterStepColorTemperatureParamsFrom constructs a [MTRColorControlClusterStepColorTemperatureParams] from an unsafe.Pointer.
func MTRColorControlClusterStepColorTemperatureParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterStepColorTemperatureParams {
	return MTRColorControlClusterStepColorTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterStepColorTemperatureParamsClass) Alloc() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterStepColorTemperatureParamsClass) New() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterStepColorTemperatureParams) Init() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterStepColorTemperatureParams) Autorelease() MTRColorControlClusterStepColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterStepColorTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterStepColorTemperatureParams creates a new MTRColorControlClusterStepColorTemperatureParams instance.
func NewMTRColorControlClusterStepColorTemperatureParams() MTRColorControlClusterStepColorTemperatureParams {
	return getMTRColorControlClusterStepColorTemperatureParamsClass().New()
}




