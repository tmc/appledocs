// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterPauseActionWithDurationParams] class.
var (
	MTRActionsClusterPauseActionWithDurationParamsClass     _MTRActionsClusterPauseActionWithDurationParamsClass
	MTRActionsClusterPauseActionWithDurationParamsClassOnce sync.Once
)

func getMTRActionsClusterPauseActionWithDurationParamsClass() _MTRActionsClusterPauseActionWithDurationParamsClass {
	MTRActionsClusterPauseActionWithDurationParamsClassOnce.Do(func() {
		MTRActionsClusterPauseActionWithDurationParamsClass = _MTRActionsClusterPauseActionWithDurationParamsClass{objc.GetClass("MTRActionsClusterPauseActionWithDurationParams")}
	})
	return MTRActionsClusterPauseActionWithDurationParamsClass
}

type _MTRActionsClusterPauseActionWithDurationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterPauseActionWithDurationParams] class.
type IMTRActionsClusterPauseActionWithDurationParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionWithDurationParams
type MTRActionsClusterPauseActionWithDurationParams struct {
	objectivec.Object
}

// MTRActionsClusterPauseActionWithDurationParamsFrom constructs a [MTRActionsClusterPauseActionWithDurationParams] from an unsafe.Pointer.
func MTRActionsClusterPauseActionWithDurationParamsFrom(ptr unsafe.Pointer) MTRActionsClusterPauseActionWithDurationParams {
	return MTRActionsClusterPauseActionWithDurationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterPauseActionWithDurationParamsClass) Alloc() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterPauseActionWithDurationParamsClass) New() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterPauseActionWithDurationParams) Init() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterPauseActionWithDurationParams) Autorelease() MTRActionsClusterPauseActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterPauseActionWithDurationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterPauseActionWithDurationParams creates a new MTRActionsClusterPauseActionWithDurationParams instance.
func NewMTRActionsClusterPauseActionWithDurationParams() MTRActionsClusterPauseActionWithDurationParams {
	return getMTRActionsClusterPauseActionWithDurationParamsClass().New()
}




