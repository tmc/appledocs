// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterDisableActionWithDurationParams] class.
var (
	MTRActionsClusterDisableActionWithDurationParamsClass     _MTRActionsClusterDisableActionWithDurationParamsClass
	MTRActionsClusterDisableActionWithDurationParamsClassOnce sync.Once
)

func getMTRActionsClusterDisableActionWithDurationParamsClass() _MTRActionsClusterDisableActionWithDurationParamsClass {
	MTRActionsClusterDisableActionWithDurationParamsClassOnce.Do(func() {
		MTRActionsClusterDisableActionWithDurationParamsClass = _MTRActionsClusterDisableActionWithDurationParamsClass{objc.GetClass("MTRActionsClusterDisableActionWithDurationParams")}
	})
	return MTRActionsClusterDisableActionWithDurationParamsClass
}

type _MTRActionsClusterDisableActionWithDurationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterDisableActionWithDurationParams] class.
type IMTRActionsClusterDisableActionWithDurationParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionWithDurationParams
type MTRActionsClusterDisableActionWithDurationParams struct {
	objectivec.Object
}

// MTRActionsClusterDisableActionWithDurationParamsFrom constructs a [MTRActionsClusterDisableActionWithDurationParams] from an unsafe.Pointer.
func MTRActionsClusterDisableActionWithDurationParamsFrom(ptr unsafe.Pointer) MTRActionsClusterDisableActionWithDurationParams {
	return MTRActionsClusterDisableActionWithDurationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterDisableActionWithDurationParamsClass) Alloc() MTRActionsClusterDisableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterDisableActionWithDurationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterDisableActionWithDurationParamsClass) New() MTRActionsClusterDisableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterDisableActionWithDurationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterDisableActionWithDurationParams) Init() MTRActionsClusterDisableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterDisableActionWithDurationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterDisableActionWithDurationParams) Autorelease() MTRActionsClusterDisableActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterDisableActionWithDurationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterDisableActionWithDurationParams creates a new MTRActionsClusterDisableActionWithDurationParams instance.
func NewMTRActionsClusterDisableActionWithDurationParams() MTRActionsClusterDisableActionWithDurationParams {
	return getMTRActionsClusterDisableActionWithDurationParamsClass().New()
}




