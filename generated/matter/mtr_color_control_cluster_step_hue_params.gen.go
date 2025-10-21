// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterStepHueParams] class.
var (
	MTRColorControlClusterStepHueParamsClass     _MTRColorControlClusterStepHueParamsClass
	MTRColorControlClusterStepHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterStepHueParamsClass() _MTRColorControlClusterStepHueParamsClass {
	MTRColorControlClusterStepHueParamsClassOnce.Do(func() {
		MTRColorControlClusterStepHueParamsClass = _MTRColorControlClusterStepHueParamsClass{objc.GetClass("MTRColorControlClusterStepHueParams")}
	})
	return MTRColorControlClusterStepHueParamsClass
}

type _MTRColorControlClusterStepHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterStepHueParams] class.
type IMTRColorControlClusterStepHueParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStepHueParams
type MTRColorControlClusterStepHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterStepHueParamsFrom constructs a [MTRColorControlClusterStepHueParams] from an unsafe.Pointer.
func MTRColorControlClusterStepHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterStepHueParams {
	return MTRColorControlClusterStepHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterStepHueParamsClass) Alloc() MTRColorControlClusterStepHueParams {
	rv := objc.Send[MTRColorControlClusterStepHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterStepHueParamsClass) New() MTRColorControlClusterStepHueParams {
	rv := objc.Send[MTRColorControlClusterStepHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterStepHueParams) Init() MTRColorControlClusterStepHueParams {
	rv := objc.Send[MTRColorControlClusterStepHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterStepHueParams) Autorelease() MTRColorControlClusterStepHueParams {
	rv := objc.Send[MTRColorControlClusterStepHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterStepHueParams creates a new MTRColorControlClusterStepHueParams instance.
func NewMTRColorControlClusterStepHueParams() MTRColorControlClusterStepHueParams {
	return getMTRColorControlClusterStepHueParamsClass().New()
}




