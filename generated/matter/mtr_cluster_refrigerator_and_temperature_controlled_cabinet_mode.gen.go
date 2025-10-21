// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterRefrigeratorAndTemperatureControlledCabinetMode] class.
var (
	MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass     _MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass
	MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClassOnce sync.Once
)

func getMTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass() _MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass {
	MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClassOnce.Do(func() {
		MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass = _MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass{objc.GetClass("MTRClusterRefrigeratorAndTemperatureControlledCabinetMode")}
	})
	return MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass
}

type _MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterRefrigeratorAndTemperatureControlledCabinetMode] class.
type IMTRClusterRefrigeratorAndTemperatureControlledCabinetMode interface {
	IMTRGenericCluster
	ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeCurrentModeWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeSupportedModesWithParams(params IMTRReadParams) unsafe.Pointer
}

// Cluster Refrigerator And Temperature Controlled Cabinet Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAndTemperatureControlledCabinetMode
type MTRClusterRefrigeratorAndTemperatureControlledCabinetMode struct {
	MTRGenericCluster
}

// MTRClusterRefrigeratorAndTemperatureControlledCabinetModeFrom constructs a [MTRClusterRefrigeratorAndTemperatureControlledCabinetMode] from an unsafe.Pointer.
//
// Cluster Refrigerator And Temperature Controlled Cabinet Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterRefrigeratorAndTemperatureControlledCabinetModeFrom(ptr unsafe.Pointer) MTRClusterRefrigeratorAndTemperatureControlledCabinetMode {
	return MTRClusterRefrigeratorAndTemperatureControlledCabinetMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass) Alloc() MTRClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRClusterRefrigeratorAndTemperatureControlledCabinetMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass) New() MTRClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRClusterRefrigeratorAndTemperatureControlledCabinetMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterRefrigeratorAndTemperatureControlledCabinetMode) Init() MTRClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRClusterRefrigeratorAndTemperatureControlledCabinetMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterRefrigeratorAndTemperatureControlledCabinetMode) Autorelease() MTRClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRClusterRefrigeratorAndTemperatureControlledCabinetMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterRefrigeratorAndTemperatureControlledCabinetMode creates a new MTRClusterRefrigeratorAndTemperatureControlledCabinetMode instance.
func NewMTRClusterRefrigeratorAndTemperatureControlledCabinetMode() MTRClusterRefrigeratorAndTemperatureControlledCabinetMode {
	return getMTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass().New()
}




// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAndTemperatureControlledCabinetMode/init(device:endpointID:queue:)
func NewMTRClusterRefrigeratorAndTemperatureControlledCabinetModeWithDeviceEndpointIDQueue(device IMTRDevice, endpointID foundation.INumber, queue unsafe.Pointer) MTRClusterRefrigeratorAndTemperatureControlledCabinetMode {
	instance := getMTRClusterRefrigeratorAndTemperatureControlledCabinetModeClass().Alloc()
	rv := objc.Send[MTRClusterRefrigeratorAndTemperatureControlledCabinetMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAndTemperatureControlledCabinetMode/changeToMode(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterRefrigeratorAndTemperatureControlledCabinetMode) ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterChangeToModeParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeAttributeList(with:)
func (m_ MTRClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeClusterRevision(with:)
func (m_ MTRClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeCurrentMode(with:)
func (m_ MTRClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeCurrentModeWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeCurrentModeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeFeatureMap(with:)
func (m_ MTRClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeSupportedModes(with:)
func (m_ MTRClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeSupportedModesWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSupportedModesWithParams:"), params)
	return rv
}


