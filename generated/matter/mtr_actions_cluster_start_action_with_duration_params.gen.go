// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRActionsClusterStartActionWithDurationParams] class.
var (
	MTRActionsClusterStartActionWithDurationParamsClass     _MTRActionsClusterStartActionWithDurationParamsClass
	MTRActionsClusterStartActionWithDurationParamsClassOnce sync.Once
)

func getMTRActionsClusterStartActionWithDurationParamsClass() _MTRActionsClusterStartActionWithDurationParamsClass {
	MTRActionsClusterStartActionWithDurationParamsClassOnce.Do(func() {
		MTRActionsClusterStartActionWithDurationParamsClass = _MTRActionsClusterStartActionWithDurationParamsClass{objc.GetClass("MTRActionsClusterStartActionWithDurationParams")}
	})
	return MTRActionsClusterStartActionWithDurationParamsClass
}

type _MTRActionsClusterStartActionWithDurationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterStartActionWithDurationParams] class.
type IMTRActionsClusterStartActionWithDurationParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStartActionWithDurationParams
type MTRActionsClusterStartActionWithDurationParams struct {
	objectivec.Object
}

// MTRActionsClusterStartActionWithDurationParamsFrom constructs a [MTRActionsClusterStartActionWithDurationParams] from an unsafe.Pointer.
func MTRActionsClusterStartActionWithDurationParamsFrom(ptr unsafe.Pointer) MTRActionsClusterStartActionWithDurationParams {
	return MTRActionsClusterStartActionWithDurationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterStartActionWithDurationParamsClass) Alloc() MTRActionsClusterStartActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterStartActionWithDurationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterStartActionWithDurationParamsClass) New() MTRActionsClusterStartActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterStartActionWithDurationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterStartActionWithDurationParams) Init() MTRActionsClusterStartActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterStartActionWithDurationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterStartActionWithDurationParams) Autorelease() MTRActionsClusterStartActionWithDurationParams {
	rv := objc.Send[MTRActionsClusterStartActionWithDurationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterStartActionWithDurationParams creates a new MTRActionsClusterStartActionWithDurationParams instance.
func NewMTRActionsClusterStartActionWithDurationParams() MTRActionsClusterStartActionWithDurationParams {
	return getMTRActionsClusterStartActionWithDurationParamsClass().New()
}




