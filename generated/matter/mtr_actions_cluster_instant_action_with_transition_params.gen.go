// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterInstantActionWithTransitionParams] class.
var (
	MTRActionsClusterInstantActionWithTransitionParamsClass     _MTRActionsClusterInstantActionWithTransitionParamsClass
	MTRActionsClusterInstantActionWithTransitionParamsClassOnce sync.Once
)

func getMTRActionsClusterInstantActionWithTransitionParamsClass() _MTRActionsClusterInstantActionWithTransitionParamsClass {
	MTRActionsClusterInstantActionWithTransitionParamsClassOnce.Do(func() {
		MTRActionsClusterInstantActionWithTransitionParamsClass = _MTRActionsClusterInstantActionWithTransitionParamsClass{objc.GetClass("MTRActionsClusterInstantActionWithTransitionParams")}
	})
	return MTRActionsClusterInstantActionWithTransitionParamsClass
}

type _MTRActionsClusterInstantActionWithTransitionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterInstantActionWithTransitionParams] class.
type IMTRActionsClusterInstantActionWithTransitionParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterInstantActionWithTransitionParams
type MTRActionsClusterInstantActionWithTransitionParams struct {
	objectivec.Object
}

// MTRActionsClusterInstantActionWithTransitionParamsFrom constructs a [MTRActionsClusterInstantActionWithTransitionParams] from an unsafe.Pointer.
func MTRActionsClusterInstantActionWithTransitionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterInstantActionWithTransitionParams {
	return MTRActionsClusterInstantActionWithTransitionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterInstantActionWithTransitionParamsClass) Alloc() MTRActionsClusterInstantActionWithTransitionParams {
	rv := objc.Send[MTRActionsClusterInstantActionWithTransitionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterInstantActionWithTransitionParamsClass) New() MTRActionsClusterInstantActionWithTransitionParams {
	rv := objc.Send[MTRActionsClusterInstantActionWithTransitionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterInstantActionWithTransitionParams) Init() MTRActionsClusterInstantActionWithTransitionParams {
	rv := objc.Send[MTRActionsClusterInstantActionWithTransitionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterInstantActionWithTransitionParams) Autorelease() MTRActionsClusterInstantActionWithTransitionParams {
	rv := objc.Send[MTRActionsClusterInstantActionWithTransitionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterInstantActionWithTransitionParams creates a new MTRActionsClusterInstantActionWithTransitionParams instance.
func NewMTRActionsClusterInstantActionWithTransitionParams() MTRActionsClusterInstantActionWithTransitionParams {
	return getMTRActionsClusterInstantActionWithTransitionParamsClass().New()
}




