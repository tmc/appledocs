// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCRunModeClusterChangeToModeResponseParams] class.
var (
	MTRRVCRunModeClusterChangeToModeResponseParamsClass     _MTRRVCRunModeClusterChangeToModeResponseParamsClass
	MTRRVCRunModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTRRVCRunModeClusterChangeToModeResponseParamsClass() _MTRRVCRunModeClusterChangeToModeResponseParamsClass {
	MTRRVCRunModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTRRVCRunModeClusterChangeToModeResponseParamsClass = _MTRRVCRunModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTRRVCRunModeClusterChangeToModeResponseParams")}
	})
	return MTRRVCRunModeClusterChangeToModeResponseParamsClass
}

type _MTRRVCRunModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCRunModeClusterChangeToModeResponseParams] class.
type IMTRRVCRunModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCRunModeClusterChangeToModeResponseParams
type MTRRVCRunModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTRRVCRunModeClusterChangeToModeResponseParamsFrom constructs a [MTRRVCRunModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTRRVCRunModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTRRVCRunModeClusterChangeToModeResponseParams {
	return MTRRVCRunModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCRunModeClusterChangeToModeResponseParamsClass) Alloc() MTRRVCRunModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRVCRunModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCRunModeClusterChangeToModeResponseParamsClass) New() MTRRVCRunModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRVCRunModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCRunModeClusterChangeToModeResponseParams) Init() MTRRVCRunModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRVCRunModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCRunModeClusterChangeToModeResponseParams) Autorelease() MTRRVCRunModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTRRVCRunModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCRunModeClusterChangeToModeResponseParams creates a new MTRRVCRunModeClusterChangeToModeResponseParams instance.
func NewMTRRVCRunModeClusterChangeToModeResponseParams() MTRRVCRunModeClusterChangeToModeResponseParams {
	return getMTRRVCRunModeClusterChangeToModeResponseParamsClass().New()
}




