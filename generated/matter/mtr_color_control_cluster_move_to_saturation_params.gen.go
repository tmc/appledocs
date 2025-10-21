// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRColorControlClusterMoveToSaturationParams] class.
var (
	MTRColorControlClusterMoveToSaturationParamsClass     _MTRColorControlClusterMoveToSaturationParamsClass
	MTRColorControlClusterMoveToSaturationParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToSaturationParamsClass() _MTRColorControlClusterMoveToSaturationParamsClass {
	MTRColorControlClusterMoveToSaturationParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToSaturationParamsClass = _MTRColorControlClusterMoveToSaturationParamsClass{objc.GetClass("MTRColorControlClusterMoveToSaturationParams")}
	})
	return MTRColorControlClusterMoveToSaturationParamsClass
}

type _MTRColorControlClusterMoveToSaturationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveToSaturationParams] class.
type IMTRColorControlClusterMoveToSaturationParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToSaturationParams
type MTRColorControlClusterMoveToSaturationParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToSaturationParamsFrom constructs a [MTRColorControlClusterMoveToSaturationParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToSaturationParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToSaturationParams {
	return MTRColorControlClusterMoveToSaturationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToSaturationParamsClass) Alloc() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveToSaturationParamsClass) New() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToSaturationParams) Init() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToSaturationParams) Autorelease() MTRColorControlClusterMoveToSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveToSaturationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToSaturationParams creates a new MTRColorControlClusterMoveToSaturationParams instance.
func NewMTRColorControlClusterMoveToSaturationParams() MTRColorControlClusterMoveToSaturationParams {
	return getMTRColorControlClusterMoveToSaturationParamsClass().New()
}




