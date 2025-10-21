// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRActionsClusterEnableActionWithDurationParams] class.
var (
	MTRActionsClusterEnableActionWithDurationParamsClass     _MTRActionsClusterEnableActionWithDurationParamsClass
	MTRActionsClusterEnableActionWithDurationParamsClassOnce sync.Once
)

func getMTRActionsClusterEnableActionWithDurationParamsClass() _MTRActionsClusterEnableActionWithDurationParamsClass {
	MTRActionsClusterEnableActionWithDurationParamsClassOnce.Do(func() {
		MTRActionsClusterEnableActionWithDurationParamsClass = _MTRActionsClusterEnableActionWithDurationParamsClass{objc.GetClass("MTRActionsClusterEnableActionWithDurationParams")}
	})
	return MTRActionsClusterEnableActionWithDurationParamsClass
}

type _MTRActionsClusterEnableActionWithDurationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterEnableActionWithDurationParams] class.
type IMTRActionsClusterEnableActionWithDurationParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterEnableActionWithDurationParams
type MTRActionsClusterEnableActionWithDurationParams struct {
	objectivec.Object
}

// MTRActionsClusterEnableActionWithDurationParamsFrom constructs a [MTRActionsClusterEnableActionWithDurationParams] from an unsafe.Pointer.
func MTRActionsClusterEnableActionWithDurationParamsFrom(ptr unsafe.Pointer) MTRActionsClusterEnableActionWithDurationParams {
	return MTRActionsClusterEnableActionWithDurationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterEnableActionWithDurationParamsClass) Alloc() MTRActionsClusterEnableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterEnableActionWithDurationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterEnableActionWithDurationParamsClass) New() MTRActionsClusterEnableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterEnableActionWithDurationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterEnableActionWithDurationParams) Init() MTRActionsClusterEnableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterEnableActionWithDurationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterEnableActionWithDurationParams) Autorelease() MTRActionsClusterEnableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterEnableActionWithDurationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterEnableActionWithDurationParams creates a new MTRActionsClusterEnableActionWithDurationParams instance.
func NewMTRActionsClusterEnableActionWithDurationParams() MTRActionsClusterEnableActionWithDurationParams {
	return getMTRActionsClusterEnableActionWithDurationParamsClass().New()
}




