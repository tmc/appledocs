// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRActionsClusterStopActionParams] class.
var (
	MTRActionsClusterStopActionParamsClass     _MTRActionsClusterStopActionParamsClass
	MTRActionsClusterStopActionParamsClassOnce sync.Once
)

func getMTRActionsClusterStopActionParamsClass() _MTRActionsClusterStopActionParamsClass {
	MTRActionsClusterStopActionParamsClassOnce.Do(func() {
		MTRActionsClusterStopActionParamsClass = _MTRActionsClusterStopActionParamsClass{objc.GetClass("MTRActionsClusterStopActionParams")}
	})
	return MTRActionsClusterStopActionParamsClass
}

type _MTRActionsClusterStopActionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRActionsClusterStopActionParams] class.
type IMTRActionsClusterStopActionParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStopActionParams
type MTRActionsClusterStopActionParams struct {
	objectivec.Object
}

// MTRActionsClusterStopActionParamsFrom constructs a [MTRActionsClusterStopActionParams] from an unsafe.Pointer.
func MTRActionsClusterStopActionParamsFrom(ptr unsafe.Pointer) MTRActionsClusterStopActionParams {
	return MTRActionsClusterStopActionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterStopActionParamsClass) Alloc() MTRActionsClusterStopActionParams {
	rv := objc.Send[MTRActionsClusterStopActionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRActionsClusterStopActionParamsClass) New() MTRActionsClusterStopActionParams {
	rv := objc.Send[MTRActionsClusterStopActionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterStopActionParams) Init() MTRActionsClusterStopActionParams {
	rv := objc.Send[MTRActionsClusterStopActionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterStopActionParams) Autorelease() MTRActionsClusterStopActionParams {
	rv := objc.Send[MTRActionsClusterStopActionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterStopActionParams creates a new MTRActionsClusterStopActionParams instance.
func NewMTRActionsClusterStopActionParams() MTRActionsClusterStopActionParams {
	return getMTRActionsClusterStopActionParamsClass().New()
}




