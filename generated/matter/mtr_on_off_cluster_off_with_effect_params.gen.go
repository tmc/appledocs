// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROnOffClusterOffWithEffectParams] class.
var (
	MTROnOffClusterOffWithEffectParamsClass     _MTROnOffClusterOffWithEffectParamsClass
	MTROnOffClusterOffWithEffectParamsClassOnce sync.Once
)

func getMTROnOffClusterOffWithEffectParamsClass() _MTROnOffClusterOffWithEffectParamsClass {
	MTROnOffClusterOffWithEffectParamsClassOnce.Do(func() {
		MTROnOffClusterOffWithEffectParamsClass = _MTROnOffClusterOffWithEffectParamsClass{objc.GetClass("MTROnOffClusterOffWithEffectParams")}
	})
	return MTROnOffClusterOffWithEffectParamsClass
}

type _MTROnOffClusterOffWithEffectParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROnOffClusterOffWithEffectParams] class.
type IMTROnOffClusterOffWithEffectParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOffWithEffectParams
type MTROnOffClusterOffWithEffectParams struct {
	objectivec.Object
}

// MTROnOffClusterOffWithEffectParamsFrom constructs a [MTROnOffClusterOffWithEffectParams] from an unsafe.Pointer.
func MTROnOffClusterOffWithEffectParamsFrom(ptr unsafe.Pointer) MTROnOffClusterOffWithEffectParams {
	return MTROnOffClusterOffWithEffectParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROnOffClusterOffWithEffectParamsClass) Alloc() MTROnOffClusterOffWithEffectParams {
	rv := objc.Send[MTROnOffClusterOffWithEffectParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROnOffClusterOffWithEffectParamsClass) New() MTROnOffClusterOffWithEffectParams {
	rv := objc.Send[MTROnOffClusterOffWithEffectParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnOffClusterOffWithEffectParams) Init() MTROnOffClusterOffWithEffectParams {
	rv := objc.Send[MTROnOffClusterOffWithEffectParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnOffClusterOffWithEffectParams) Autorelease() MTROnOffClusterOffWithEffectParams {
	rv := objc.Send[MTROnOffClusterOffWithEffectParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnOffClusterOffWithEffectParams creates a new MTROnOffClusterOffWithEffectParams instance.
func NewMTROnOffClusterOffWithEffectParams() MTROnOffClusterOffWithEffectParams {
	return getMTROnOffClusterOffWithEffectParamsClass().New()
}




