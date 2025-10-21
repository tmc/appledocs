// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRColorControlClusterMoveToColorTemperatureParams] class.
var (
	MTRColorControlClusterMoveToColorTemperatureParamsClass     _MTRColorControlClusterMoveToColorTemperatureParamsClass
	MTRColorControlClusterMoveToColorTemperatureParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToColorTemperatureParamsClass() _MTRColorControlClusterMoveToColorTemperatureParamsClass {
	MTRColorControlClusterMoveToColorTemperatureParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToColorTemperatureParamsClass = _MTRColorControlClusterMoveToColorTemperatureParamsClass{objc.GetClass("MTRColorControlClusterMoveToColorTemperatureParams")}
	})
	return MTRColorControlClusterMoveToColorTemperatureParamsClass
}

type _MTRColorControlClusterMoveToColorTemperatureParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveToColorTemperatureParams] class.
type IMTRColorControlClusterMoveToColorTemperatureParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorTemperatureParams
type MTRColorControlClusterMoveToColorTemperatureParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToColorTemperatureParamsFrom constructs a [MTRColorControlClusterMoveToColorTemperatureParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToColorTemperatureParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToColorTemperatureParams {
	return MTRColorControlClusterMoveToColorTemperatureParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToColorTemperatureParamsClass) Alloc() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveToColorTemperatureParamsClass) New() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) Init() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToColorTemperatureParams) Autorelease() MTRColorControlClusterMoveToColorTemperatureParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorTemperatureParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToColorTemperatureParams creates a new MTRColorControlClusterMoveToColorTemperatureParams instance.
func NewMTRColorControlClusterMoveToColorTemperatureParams() MTRColorControlClusterMoveToColorTemperatureParams {
	return getMTRColorControlClusterMoveToColorTemperatureParamsClass().New()
}




