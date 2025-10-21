// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRColorControlClusterMoveColorTemperatureParams] class.
var (
	MTRColorControlClusterMoveColorTemperatureParamsClass     _MTRColorControlClusterMoveColorTemperatureParamsClass
	MTRColorControlClusterMoveColorTemperatureParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveColorTemperatureParamsClass() _MTRColorControlClusterMoveColorTemperatureParamsClass {
	MTRColorControlClusterMoveColorTemperatureParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveColorTemperatureParamsClass = _MTRColorControlClusterMoveColorTemperatureParamsClass{objc.GetClass("MTRColorControlClusterMoveColorTemperatureParams")}
	})
	return MTRColorControlClusterMoveColorTemperatureParamsClass
}

type _MTRColorControlClusterMoveColorTemperatureParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveColorTemperatureParams] class.
type IMTRColorControlClusterMoveColorTemperatureParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorTemperatureParams
type MTRColorControlClusterMoveColorTemperatureParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveColorTemperatureParamsFrom constructs a [MTRColorControlClusterMoveColorTemperatureParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveColorTemperatureParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveColorTemperatureParams {
	return MTRColorControlClusterMoveColorTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveColorTemperatureParamsClass) Alloc() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveColorTemperatureParamsClass) New() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveColorTemperatureParams) Init() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveColorTemperatureParams) Autorelease() MTRColorControlClusterMoveColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveColorTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveColorTemperatureParams creates a new MTRColorControlClusterMoveColorTemperatureParams instance.
func NewMTRColorControlClusterMoveColorTemperatureParams() MTRColorControlClusterMoveColorTemperatureParams {
	return getMTRColorControlClusterMoveColorTemperatureParamsClass().New()
}




