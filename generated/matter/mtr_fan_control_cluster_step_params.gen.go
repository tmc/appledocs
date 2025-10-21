// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRFanControlClusterStepParams] class.
var (
	MTRFanControlClusterStepParamsClass     _MTRFanControlClusterStepParamsClass
	MTRFanControlClusterStepParamsClassOnce sync.Once
)

func getMTRFanControlClusterStepParamsClass() _MTRFanControlClusterStepParamsClass {
	MTRFanControlClusterStepParamsClassOnce.Do(func() {
		MTRFanControlClusterStepParamsClass = _MTRFanControlClusterStepParamsClass{objc.GetClass("MTRFanControlClusterStepParams")}
	})
	return MTRFanControlClusterStepParamsClass
}

type _MTRFanControlClusterStepParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRFanControlClusterStepParams] class.
type IMTRFanControlClusterStepParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFanControlClusterStepParams
type MTRFanControlClusterStepParams struct {
	objectivec.Object
}

// MTRFanControlClusterStepParamsFrom constructs a [MTRFanControlClusterStepParams] from an unsafe.Pointer.
func MTRFanControlClusterStepParamsFrom(ptr unsafe.Pointer) MTRFanControlClusterStepParams {
	return MTRFanControlClusterStepParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRFanControlClusterStepParamsClass) Alloc() MTRFanControlClusterStepParams {
	rv := objc.Send[MTRFanControlClusterStepParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRFanControlClusterStepParamsClass) New() MTRFanControlClusterStepParams {
	rv := objc.Send[MTRFanControlClusterStepParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRFanControlClusterStepParams) Init() MTRFanControlClusterStepParams {
	rv := objc.Send[MTRFanControlClusterStepParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRFanControlClusterStepParams) Autorelease() MTRFanControlClusterStepParams {
	rv := objc.Send[MTRFanControlClusterStepParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRFanControlClusterStepParams creates a new MTRFanControlClusterStepParams instance.
func NewMTRFanControlClusterStepParams() MTRFanControlClusterStepParams {
	return getMTRFanControlClusterStepParamsClass().New()
}




