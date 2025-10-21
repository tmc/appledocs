// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams] class.
var (
	MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass     _MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass
	MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClassOnce sync.Once
)

func getMTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass() _MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass {
	MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClassOnce.Do(func() {
		MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass = _MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass{objc.GetClass("MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams")}
	})
	return MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass
}

type _MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams] class.
type IMTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams
type MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams struct {
	objectivec.Object
}

// MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsFrom constructs a [MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams] from an unsafe.Pointer.
func MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsFrom(ptr unsafe.Pointer) MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams {
	return MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass) Alloc() MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass) New() MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) Init() MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams) Autorelease() MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams creates a new MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams instance.
func NewMTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams() MTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParams {
	return getMTRElectricalMeasurementClusterGetMeasurementProfileResponseCommandParamsClass().New()
}




