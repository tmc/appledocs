// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterEnergyEVSEMode] class.
var (
	MTRClusterEnergyEVSEModeClass     _MTRClusterEnergyEVSEModeClass
	MTRClusterEnergyEVSEModeClassOnce sync.Once
)

func getMTRClusterEnergyEVSEModeClass() _MTRClusterEnergyEVSEModeClass {
	MTRClusterEnergyEVSEModeClassOnce.Do(func() {
		MTRClusterEnergyEVSEModeClass = _MTRClusterEnergyEVSEModeClass{objc.GetClass("MTRClusterEnergyEVSEMode")}
	})
	return MTRClusterEnergyEVSEModeClass
}

type _MTRClusterEnergyEVSEModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterEnergyEVSEMode] class.
type IMTRClusterEnergyEVSEMode interface {
	IMTRGenericCluster
	ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEModeClusterChangeToModeParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeCurrentModeWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeSupportedModesWithParams(params IMTRReadParams) unsafe.Pointer
}

// Cluster Energy EVSE Mode Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode
type MTRClusterEnergyEVSEMode struct {
	MTRGenericCluster
}

// MTRClusterEnergyEVSEModeFrom constructs a [MTRClusterEnergyEVSEMode] from an unsafe.Pointer.
//
// Cluster Energy EVSE Mode Attributes and commands for selecting a mode from a list of supported options.
func MTRClusterEnergyEVSEModeFrom(ptr unsafe.Pointer) MTRClusterEnergyEVSEMode {
	return MTRClusterEnergyEVSEMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterEnergyEVSEModeClass) Alloc() MTRClusterEnergyEVSEMode {
	rv := objc.Send[MTRClusterEnergyEVSEMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterEnergyEVSEModeClass) New() MTRClusterEnergyEVSEMode {
	rv := objc.Send[MTRClusterEnergyEVSEMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterEnergyEVSEMode) Init() MTRClusterEnergyEVSEMode {
	rv := objc.Send[MTRClusterEnergyEVSEMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterEnergyEVSEMode) Autorelease() MTRClusterEnergyEVSEMode {
	rv := objc.Send[MTRClusterEnergyEVSEMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterEnergyEVSEMode creates a new MTRClusterEnergyEVSEMode instance.
func NewMTRClusterEnergyEVSEMode() MTRClusterEnergyEVSEMode {
	return getMTRClusterEnergyEVSEModeClass().New()
}




// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode/init(device:endpointID:queue:)
func NewMTRClusterEnergyEVSEModeWithDeviceEndpointIDQueue(device IMTRDevice, endpointID foundation.INumber, queue unsafe.Pointer) MTRClusterEnergyEVSEMode {
	instance := getMTRClusterEnergyEVSEModeClass().Alloc()
	rv := objc.Send[MTRClusterEnergyEVSEMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode/changeToMode(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSEMode) ChangeToModeWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEModeClusterChangeToModeParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterEnergyEVSEMode) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode/readAttributeAttributeList(with:)
func (m_ MTRClusterEnergyEVSEMode) ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode/readAttributeClusterRevision(with:)
func (m_ MTRClusterEnergyEVSEMode) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode/readAttributeCurrentMode(with:)
func (m_ MTRClusterEnergyEVSEMode) ReadAttributeCurrentModeWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeCurrentModeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode/readAttributeFeatureMap(with:)
func (m_ MTRClusterEnergyEVSEMode) ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterEnergyEVSEMode) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSEMode/readAttributeSupportedModes(with:)
func (m_ MTRClusterEnergyEVSEMode) ReadAttributeSupportedModesWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSupportedModesWithParams:"), params)
	return rv
}


