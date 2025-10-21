// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterWaterHeaterManagement] class.
var (
	MTRClusterWaterHeaterManagementClass     _MTRClusterWaterHeaterManagementClass
	MTRClusterWaterHeaterManagementClassOnce sync.Once
)

func getMTRClusterWaterHeaterManagementClass() _MTRClusterWaterHeaterManagementClass {
	MTRClusterWaterHeaterManagementClassOnce.Do(func() {
		MTRClusterWaterHeaterManagementClass = _MTRClusterWaterHeaterManagementClass{objc.GetClass("MTRClusterWaterHeaterManagement")}
	})
	return MTRClusterWaterHeaterManagementClass
}

type _MTRClusterWaterHeaterManagementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterWaterHeaterManagement] class.
type IMTRClusterWaterHeaterManagement interface {
	IMTRGenericCluster
	BoostWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer)
	CancelBoostWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer)
	CancelBoostWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeBoostStateWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeEstimatedHeatRequiredWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeHeatDemandWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeHeaterTypesWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeTankPercentageWithParams(params unsafe.Pointer) unsafe.Pointer
	ReadAttributeTankVolumeWithParams(params unsafe.Pointer) unsafe.Pointer
}

// Cluster Water Heater Management This cluster is used to allow clients to control the operation of a hot water heating appliance so that it can be used with energy management.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement
type MTRClusterWaterHeaterManagement struct {
	MTRGenericCluster
}

// MTRClusterWaterHeaterManagementFrom constructs a [MTRClusterWaterHeaterManagement] from an unsafe.Pointer.
//
// Cluster Water Heater Management This cluster is used to allow clients to control the operation of a hot water heating appliance so that it can be used with energy management.
func MTRClusterWaterHeaterManagementFrom(ptr unsafe.Pointer) MTRClusterWaterHeaterManagement {
	return MTRClusterWaterHeaterManagement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterWaterHeaterManagementClass) Alloc() MTRClusterWaterHeaterManagement {
	rv := objc.Send[MTRClusterWaterHeaterManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterWaterHeaterManagementClass) New() MTRClusterWaterHeaterManagement {
	rv := objc.Send[MTRClusterWaterHeaterManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterWaterHeaterManagement) Init() MTRClusterWaterHeaterManagement {
	rv := objc.Send[MTRClusterWaterHeaterManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterWaterHeaterManagement) Autorelease() MTRClusterWaterHeaterManagement {
	rv := objc.Send[MTRClusterWaterHeaterManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterWaterHeaterManagement creates a new MTRClusterWaterHeaterManagement instance.
func NewMTRClusterWaterHeaterManagement() MTRClusterWaterHeaterManagement {
	return getMTRClusterWaterHeaterManagementClass().New()
}




// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/init(device:endpointID:queue:)
func NewMTRClusterWaterHeaterManagementWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID foundation.Number, queue unsafe.Pointer) MTRClusterWaterHeaterManagement {
	instance := getMTRClusterWaterHeaterManagementClass().Alloc()
	rv := objc.Send[MTRClusterWaterHeaterManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/boost(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterWaterHeaterManagement) BoostWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("boostWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/cancelBoost(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterWaterHeaterManagement) CancelBoostWithParamsExpectedValuesExpectedValueIntervalCompletion(params unsafe.Pointer, expectedDataValueDictionaries unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelBoostWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/cancelBoost(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterWaterHeaterManagement) CancelBoostWithExpectedValuesExpectedValueIntervalCompletion(expectedValues unsafe.Pointer, expectedValueIntervalMs foundation.Number, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelBoostWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterWaterHeaterManagement) ReadAttributeAcceptedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/readAttributeAttributeList(with:)
func (m_ MTRClusterWaterHeaterManagement) ReadAttributeAttributeListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/readAttributeBoostState(with:)
func (m_ MTRClusterWaterHeaterManagement) ReadAttributeBoostStateWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeBoostStateWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/readAttributeClusterRevision(with:)
func (m_ MTRClusterWaterHeaterManagement) ReadAttributeClusterRevisionWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/readAttributeEstimatedHeatRequired(with:)
func (m_ MTRClusterWaterHeaterManagement) ReadAttributeEstimatedHeatRequiredWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeEstimatedHeatRequiredWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/readAttributeFeatureMap(with:)
func (m_ MTRClusterWaterHeaterManagement) ReadAttributeFeatureMapWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterWaterHeaterManagement) ReadAttributeGeneratedCommandListWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/readAttributeHeatDemand(with:)
func (m_ MTRClusterWaterHeaterManagement) ReadAttributeHeatDemandWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeHeatDemandWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/readAttributeHeaterTypes(with:)
func (m_ MTRClusterWaterHeaterManagement) ReadAttributeHeaterTypesWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeHeaterTypesWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/readAttributeTankPercentage(with:)
func (m_ MTRClusterWaterHeaterManagement) ReadAttributeTankPercentageWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeTankPercentageWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/readAttributeTankVolume(with:)
func (m_ MTRClusterWaterHeaterManagement) ReadAttributeTankVolumeWithParams(params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeTankVolumeWithParams:"), params)
	return rv
}


