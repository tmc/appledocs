// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRColorControlClusterMoveSaturationParams] class.
var (
	MTRColorControlClusterMoveSaturationParamsClass     _MTRColorControlClusterMoveSaturationParamsClass
	MTRColorControlClusterMoveSaturationParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveSaturationParamsClass() _MTRColorControlClusterMoveSaturationParamsClass {
	MTRColorControlClusterMoveSaturationParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveSaturationParamsClass = _MTRColorControlClusterMoveSaturationParamsClass{objc.GetClass("MTRColorControlClusterMoveSaturationParams")}
	})
	return MTRColorControlClusterMoveSaturationParamsClass
}

type _MTRColorControlClusterMoveSaturationParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveSaturationParams] class.
type IMTRColorControlClusterMoveSaturationParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveSaturationParams
type MTRColorControlClusterMoveSaturationParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveSaturationParamsFrom constructs a [MTRColorControlClusterMoveSaturationParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveSaturationParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveSaturationParams {
	return MTRColorControlClusterMoveSaturationParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveSaturationParamsClass) Alloc() MTRColorControlClusterMoveSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveSaturationParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveSaturationParamsClass) New() MTRColorControlClusterMoveSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveSaturationParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveSaturationParams) Init() MTRColorControlClusterMoveSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveSaturationParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveSaturationParams) Autorelease() MTRColorControlClusterMoveSaturationParams {
	rv := objc.Send[MTRColorControlClusterMoveSaturationParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveSaturationParams creates a new MTRColorControlClusterMoveSaturationParams instance.
func NewMTRColorControlClusterMoveSaturationParams() MTRColorControlClusterMoveSaturationParams {
	return getMTRColorControlClusterMoveSaturationParamsClass().New()
}




