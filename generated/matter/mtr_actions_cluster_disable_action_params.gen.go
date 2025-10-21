// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterDisableActionParams] class.
var (
	MTRActionsClusterDisableActionParamsClass     _MTRActionsClusterDisableActionParamsClass
	MTRActionsClusterDisableActionParamsClassOnce sync.Once
)

func getMTRActionsClusterDisableActionParamsClass() _MTRActionsClusterDisableActionParamsClass {
	MTRActionsClusterDisableActionParamsClassOnce.Do(func() {
		MTRActionsClusterDisableActionParamsClass = _MTRActionsClusterDisableActionParamsClass{objc.GetClass("MTRActionsClusterDisableActionParams")}
	})
	return MTRActionsClusterDisableActionParamsClass
}

type _MTRActionsClusterDisableActionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterDisableActionParams] class.
type IMTRActionsClusterDisableActionParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterDisableActionParams
type MTRActionsClusterDisableActionParams struct {
	objectivec.Object
}

// MTRActionsClusterDisableActionParamsFrom constructs a [MTRActionsClusterDisableActionParams] from an unsafe.Pointer.
func MTRActionsClusterDisableActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterDisableActionParams {
	return MTRActionsClusterDisableActionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterDisableActionParamsClass) Alloc() MTRActionsClusterDisableActionParams {
	rv := objc.Send[MTRActionsClusterDisableActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterDisableActionParamsClass) New() MTRActionsClusterDisableActionParams {
	rv := objc.Send[MTRActionsClusterDisableActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterDisableActionParams) Init() MTRActionsClusterDisableActionParams {
	rv := objc.Send[MTRActionsClusterDisableActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterDisableActionParams) Autorelease() MTRActionsClusterDisableActionParams {
	rv := objc.Send[MTRActionsClusterDisableActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterDisableActionParams creates a new MTRActionsClusterDisableActionParams instance.
func NewMTRActionsClusterDisableActionParams() MTRActionsClusterDisableActionParams {
	return getMTRActionsClusterDisableActionParamsClass().New()
}




