// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRIdentifyClusterTriggerEffectParams] class.
var (
	MTRIdentifyClusterTriggerEffectParamsClass     _MTRIdentifyClusterTriggerEffectParamsClass
	MTRIdentifyClusterTriggerEffectParamsClassOnce sync.Once
)

func getMTRIdentifyClusterTriggerEffectParamsClass() _MTRIdentifyClusterTriggerEffectParamsClass {
	MTRIdentifyClusterTriggerEffectParamsClassOnce.Do(func() {
		MTRIdentifyClusterTriggerEffectParamsClass = _MTRIdentifyClusterTriggerEffectParamsClass{objc.GetClass("MTRIdentifyClusterTriggerEffectParams")}
	})
	return MTRIdentifyClusterTriggerEffectParamsClass
}

type _MTRIdentifyClusterTriggerEffectParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRIdentifyClusterTriggerEffectParams] class.
type IMTRIdentifyClusterTriggerEffectParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRIdentifyClusterTriggerEffectParams
type MTRIdentifyClusterTriggerEffectParams struct {
	objectivec.Object
}

// MTRIdentifyClusterTriggerEffectParamsFrom constructs a [MTRIdentifyClusterTriggerEffectParams] from an unsafe.Pointer.
func MTRIdentifyClusterTriggerEffectParamsFrom(ptr unsafe.Pointer) MTRIdentifyClusterTriggerEffectParams {
	return MTRIdentifyClusterTriggerEffectParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRIdentifyClusterTriggerEffectParamsClass) Alloc() MTRIdentifyClusterTriggerEffectParams {
	rv := objc.Send[MTRIdentifyClusterTriggerEffectParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRIdentifyClusterTriggerEffectParamsClass) New() MTRIdentifyClusterTriggerEffectParams {
	rv := objc.Send[MTRIdentifyClusterTriggerEffectParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRIdentifyClusterTriggerEffectParams) Init() MTRIdentifyClusterTriggerEffectParams {
	rv := objc.Send[MTRIdentifyClusterTriggerEffectParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRIdentifyClusterTriggerEffectParams) Autorelease() MTRIdentifyClusterTriggerEffectParams {
	rv := objc.Send[MTRIdentifyClusterTriggerEffectParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRIdentifyClusterTriggerEffectParams creates a new MTRIdentifyClusterTriggerEffectParams instance.
func NewMTRIdentifyClusterTriggerEffectParams() MTRIdentifyClusterTriggerEffectParams {
	return getMTRIdentifyClusterTriggerEffectParamsClass().New()
}




