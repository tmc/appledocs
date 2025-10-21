// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRColorControlClusterEnhancedMoveToHueParams] class.
var (
	MTRColorControlClusterEnhancedMoveToHueParamsClass     _MTRColorControlClusterEnhancedMoveToHueParamsClass
	MTRColorControlClusterEnhancedMoveToHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterEnhancedMoveToHueParamsClass() _MTRColorControlClusterEnhancedMoveToHueParamsClass {
	MTRColorControlClusterEnhancedMoveToHueParamsClassOnce.Do(func() {
		MTRColorControlClusterEnhancedMoveToHueParamsClass = _MTRColorControlClusterEnhancedMoveToHueParamsClass{objc.GetClass("MTRColorControlClusterEnhancedMoveToHueParams")}
	})
	return MTRColorControlClusterEnhancedMoveToHueParamsClass
}

type _MTRColorControlClusterEnhancedMoveToHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterEnhancedMoveToHueParams] class.
type IMTRColorControlClusterEnhancedMoveToHueParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterEnhancedMoveToHueParams
type MTRColorControlClusterEnhancedMoveToHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterEnhancedMoveToHueParamsFrom constructs a [MTRColorControlClusterEnhancedMoveToHueParams] from an unsafe.Pointer.
func MTRColorControlClusterEnhancedMoveToHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterEnhancedMoveToHueParams {
	return MTRColorControlClusterEnhancedMoveToHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterEnhancedMoveToHueParamsClass) Alloc() MTRColorControlClusterEnhancedMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterEnhancedMoveToHueParamsClass) New() MTRColorControlClusterEnhancedMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) Init() MTRColorControlClusterEnhancedMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterEnhancedMoveToHueParams) Autorelease() MTRColorControlClusterEnhancedMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterEnhancedMoveToHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterEnhancedMoveToHueParams creates a new MTRColorControlClusterEnhancedMoveToHueParams instance.
func NewMTRColorControlClusterEnhancedMoveToHueParams() MTRColorControlClusterEnhancedMoveToHueParams {
	return getMTRColorControlClusterEnhancedMoveToHueParamsClass().New()
}




