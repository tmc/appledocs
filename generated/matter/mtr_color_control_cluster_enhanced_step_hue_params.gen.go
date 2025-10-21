// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRColorControlClusterEnhancedStepHueParams] class.
var (
	MTRColorControlClusterEnhancedStepHueParamsClass     _MTRColorControlClusterEnhancedStepHueParamsClass
	MTRColorControlClusterEnhancedStepHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterEnhancedStepHueParamsClass() _MTRColorControlClusterEnhancedStepHueParamsClass {
	MTRColorControlClusterEnhancedStepHueParamsClassOnce.Do(func() {
		MTRColorControlClusterEnhancedStepHueParamsClass = _MTRColorControlClusterEnhancedStepHueParamsClass{objc.GetClass("MTRColorControlClusterEnhancedStepHueParams")}
	})
	return MTRColorControlClusterEnhancedStepHueParamsClass
}

type _MTRColorControlClusterEnhancedStepHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterEnhancedStepHueParams] class.
type IMTRColorControlClusterEnhancedStepHueParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedStepHueParams
type MTRColorControlClusterEnhancedStepHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterEnhancedStepHueParamsFrom constructs a [MTRColorControlClusterEnhancedStepHueParams] from an unsafe.Pointer.
func MTRColorControlClusterEnhancedStepHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterEnhancedStepHueParams {
	return MTRColorControlClusterEnhancedStepHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterEnhancedStepHueParamsClass) Alloc() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterEnhancedStepHueParamsClass) New() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterEnhancedStepHueParams) Init() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterEnhancedStepHueParams) Autorelease() MTRColorControlClusterEnhancedStepHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedStepHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterEnhancedStepHueParams creates a new MTRColorControlClusterEnhancedStepHueParams instance.
func NewMTRColorControlClusterEnhancedStepHueParams() MTRColorControlClusterEnhancedStepHueParams {
	return getMTRColorControlClusterEnhancedStepHueParamsClass().New()
}




