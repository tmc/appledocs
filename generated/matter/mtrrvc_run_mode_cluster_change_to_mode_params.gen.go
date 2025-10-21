// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCRunModeClusterChangeToModeParams] class.
var (
	MTRRVCRunModeClusterChangeToModeParamsClass     _MTRRVCRunModeClusterChangeToModeParamsClass
	MTRRVCRunModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRRVCRunModeClusterChangeToModeParamsClass() _MTRRVCRunModeClusterChangeToModeParamsClass {
	MTRRVCRunModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTRRVCRunModeClusterChangeToModeParamsClass = _MTRRVCRunModeClusterChangeToModeParamsClass{objc.GetClass("MTRRVCRunModeClusterChangeToModeParams")}
	})
	return MTRRVCRunModeClusterChangeToModeParamsClass
}

type _MTRRVCRunModeClusterChangeToModeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCRunModeClusterChangeToModeParams] class.
type IMTRRVCRunModeClusterChangeToModeParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCRunModeClusterChangeToModeParams
type MTRRVCRunModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRRVCRunModeClusterChangeToModeParamsFrom constructs a [MTRRVCRunModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTRRVCRunModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRRVCRunModeClusterChangeToModeParams {
	return MTRRVCRunModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCRunModeClusterChangeToModeParamsClass) Alloc() MTRRVCRunModeClusterChangeToModeParams {
	rv := objc.Send[MTRRVCRunModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCRunModeClusterChangeToModeParamsClass) New() MTRRVCRunModeClusterChangeToModeParams {
	rv := objc.Send[MTRRVCRunModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCRunModeClusterChangeToModeParams) Init() MTRRVCRunModeClusterChangeToModeParams {
	rv := objc.Send[MTRRVCRunModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCRunModeClusterChangeToModeParams) Autorelease() MTRRVCRunModeClusterChangeToModeParams {
	rv := objc.Send[MTRRVCRunModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCRunModeClusterChangeToModeParams creates a new MTRRVCRunModeClusterChangeToModeParams instance.
func NewMTRRVCRunModeClusterChangeToModeParams() MTRRVCRunModeClusterChangeToModeParams {
	return getMTRRVCRunModeClusterChangeToModeParamsClass().New()
}




