// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRColorControlClusterEnhancedMoveHueParams] class.
var (
	MTRColorControlClusterEnhancedMoveHueParamsClass     _MTRColorControlClusterEnhancedMoveHueParamsClass
	MTRColorControlClusterEnhancedMoveHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterEnhancedMoveHueParamsClass() _MTRColorControlClusterEnhancedMoveHueParamsClass {
	MTRColorControlClusterEnhancedMoveHueParamsClassOnce.Do(func() {
		MTRColorControlClusterEnhancedMoveHueParamsClass = _MTRColorControlClusterEnhancedMoveHueParamsClass{objc.GetClass("MTRColorControlClusterEnhancedMoveHueParams")}
	})
	return MTRColorControlClusterEnhancedMoveHueParamsClass
}

type _MTRColorControlClusterEnhancedMoveHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterEnhancedMoveHueParams] class.
type IMTRColorControlClusterEnhancedMoveHueParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveHueParams
type MTRColorControlClusterEnhancedMoveHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterEnhancedMoveHueParamsFrom constructs a [MTRColorControlClusterEnhancedMoveHueParams] from an unsafe.Pointer.
func MTRColorControlClusterEnhancedMoveHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterEnhancedMoveHueParams {
	return MTRColorControlClusterEnhancedMoveHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterEnhancedMoveHueParamsClass) Alloc() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterEnhancedMoveHueParamsClass) New() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterEnhancedMoveHueParams) Init() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterEnhancedMoveHueParams) Autorelease() MTRColorControlClusterEnhancedMoveHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterEnhancedMoveHueParams creates a new MTRColorControlClusterEnhancedMoveHueParams instance.
func NewMTRColorControlClusterEnhancedMoveHueParams() MTRColorControlClusterEnhancedMoveHueParams {
	return getMTRColorControlClusterEnhancedMoveHueParamsClass().New()
}




