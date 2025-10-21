// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterPauseActionParams] class.
var (
	MTRActionsClusterPauseActionParamsClass     _MTRActionsClusterPauseActionParamsClass
	MTRActionsClusterPauseActionParamsClassOnce sync.Once
)

func getMTRActionsClusterPauseActionParamsClass() _MTRActionsClusterPauseActionParamsClass {
	MTRActionsClusterPauseActionParamsClassOnce.Do(func() {
		MTRActionsClusterPauseActionParamsClass = _MTRActionsClusterPauseActionParamsClass{objc.GetClass("MTRActionsClusterPauseActionParams")}
	})
	return MTRActionsClusterPauseActionParamsClass
}

type _MTRActionsClusterPauseActionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterPauseActionParams] class.
type IMTRActionsClusterPauseActionParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterPauseActionParams
type MTRActionsClusterPauseActionParams struct {
	objectivec.Object
}

// MTRActionsClusterPauseActionParamsFrom constructs a [MTRActionsClusterPauseActionParams] from an unsafe.Pointer.
func MTRActionsClusterPauseActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterPauseActionParams {
	return MTRActionsClusterPauseActionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterPauseActionParamsClass) Alloc() MTRActionsClusterPauseActionParams {
	rv := objc.Send[MTRActionsClusterPauseActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterPauseActionParamsClass) New() MTRActionsClusterPauseActionParams {
	rv := objc.Send[MTRActionsClusterPauseActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterPauseActionParams) Init() MTRActionsClusterPauseActionParams {
	rv := objc.Send[MTRActionsClusterPauseActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterPauseActionParams) Autorelease() MTRActionsClusterPauseActionParams {
	rv := objc.Send[MTRActionsClusterPauseActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterPauseActionParams creates a new MTRActionsClusterPauseActionParams instance.
func NewMTRActionsClusterPauseActionParams() MTRActionsClusterPauseActionParams {
	return getMTRActionsClusterPauseActionParamsClass().New()
}




