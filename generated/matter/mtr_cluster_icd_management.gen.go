// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterICDManagement] class.
var (
	MTRClusterICDManagementClass     _MTRClusterICDManagementClass
	MTRClusterICDManagementClassOnce sync.Once
)

func getMTRClusterICDManagementClass() _MTRClusterICDManagementClass {
	MTRClusterICDManagementClassOnce.Do(func() {
		MTRClusterICDManagementClass = _MTRClusterICDManagementClass{objc.GetClass("MTRClusterICDManagement")}
	})
	return MTRClusterICDManagementClass
}

type _MTRClusterICDManagementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterICDManagement] class.
type IMTRClusterICDManagement interface {
	IMTRGenericCluster
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeActiveModeDurationWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeActiveModeThresholdWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeClientsSupportedPerFabricWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeICDCounterWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeIdleModeDurationWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeMaximumCheckInBackOffWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeOperatingModeWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeRegisteredClientsWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeUserActiveModeTriggerHintWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeUserActiveModeTriggerInstructionWithParams(params IMTRReadParams) unsafe.Pointer
	RegisterClientWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRICDManagementClusterRegisterClientParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	StayActiveRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRICDManagementClusterStayActiveRequestParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	UnregisterClientWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRICDManagementClusterUnregisterClientParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
}

// Cluster ICD Management Allows servers to ensure that listed clients are notified when a server is available for communication.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement
type MTRClusterICDManagement struct {
	MTRGenericCluster
}

// MTRClusterICDManagementFrom constructs a [MTRClusterICDManagement] from an unsafe.Pointer.
//
// Cluster ICD Management Allows servers to ensure that listed clients are notified when a server is available for communication.
func MTRClusterICDManagementFrom(ptr unsafe.Pointer) MTRClusterICDManagement {
	return MTRClusterICDManagement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterICDManagementClass) Alloc() MTRClusterICDManagement {
	rv := objc.Send[MTRClusterICDManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterICDManagementClass) New() MTRClusterICDManagement {
	rv := objc.Send[MTRClusterICDManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterICDManagement) Init() MTRClusterICDManagement {
	rv := objc.Send[MTRClusterICDManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterICDManagement) Autorelease() MTRClusterICDManagement {
	rv := objc.Send[MTRClusterICDManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterICDManagement creates a new MTRClusterICDManagement instance.
func NewMTRClusterICDManagement() MTRClusterICDManagement {
	return getMTRClusterICDManagementClass().New()
}




// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/init(device:endpointID:queue:)
func NewMTRClusterICDManagementWithDeviceEndpointIDQueue(device IMTRDevice, endpointID foundation.INumber, queue unsafe.Pointer) MTRClusterICDManagement {
	instance := getMTRClusterICDManagementClass().Alloc()
	rv := objc.Send[MTRClusterICDManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterICDManagement) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeActiveModeDuration(with:)
func (m_ MTRClusterICDManagement) ReadAttributeActiveModeDurationWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeActiveModeDurationWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeActiveModeThreshold(with:)
func (m_ MTRClusterICDManagement) ReadAttributeActiveModeThresholdWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeActiveModeThresholdWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeAttributeList(with:)
func (m_ MTRClusterICDManagement) ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeClientsSupportedPerFabric(with:)
func (m_ MTRClusterICDManagement) ReadAttributeClientsSupportedPerFabricWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClientsSupportedPerFabricWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeClusterRevision(with:)
func (m_ MTRClusterICDManagement) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeFeatureMap(with:)
func (m_ MTRClusterICDManagement) ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterICDManagement) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeICDCounter(with:)
func (m_ MTRClusterICDManagement) ReadAttributeICDCounterWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeICDCounterWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeIdleModeDuration(with:)
func (m_ MTRClusterICDManagement) ReadAttributeIdleModeDurationWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeIdleModeDurationWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeMaximumCheckInBackOff(with:)
func (m_ MTRClusterICDManagement) ReadAttributeMaximumCheckInBackOffWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeMaximumCheckInBackOffWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeOperatingMode(with:)
func (m_ MTRClusterICDManagement) ReadAttributeOperatingModeWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeOperatingModeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeRegisteredClients(with:)
func (m_ MTRClusterICDManagement) ReadAttributeRegisteredClientsWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeRegisteredClientsWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeUserActiveModeTriggerHint(with:)
func (m_ MTRClusterICDManagement) ReadAttributeUserActiveModeTriggerHintWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeUserActiveModeTriggerHintWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeUserActiveModeTriggerInstruction(with:)
func (m_ MTRClusterICDManagement) ReadAttributeUserActiveModeTriggerInstructionWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeUserActiveModeTriggerInstructionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/registerClient(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterICDManagement) RegisterClientWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRICDManagementClusterRegisterClientParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("registerClientWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/stayActiveRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterICDManagement) StayActiveRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRICDManagementClusterStayActiveRequestParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stayActiveRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/unregisterClient(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterICDManagement) UnregisterClientWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRICDManagementClusterUnregisterClientParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("unregisterClientWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


