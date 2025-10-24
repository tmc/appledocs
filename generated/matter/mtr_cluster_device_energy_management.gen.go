// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterDeviceEnergyManagement] class.
var (
	MTRClusterDeviceEnergyManagementClass     _MTRClusterDeviceEnergyManagementClass
	MTRClusterDeviceEnergyManagementClassOnce sync.Once
)

func getMTRClusterDeviceEnergyManagementClass() _MTRClusterDeviceEnergyManagementClass {
	MTRClusterDeviceEnergyManagementClassOnce.Do(func() {
		MTRClusterDeviceEnergyManagementClass = _MTRClusterDeviceEnergyManagementClass{objc.GetClass("MTRClusterDeviceEnergyManagement")}
	})
	return MTRClusterDeviceEnergyManagementClass
}

type _MTRClusterDeviceEnergyManagementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterDeviceEnergyManagement] class.
type IMTRClusterDeviceEnergyManagement interface {
	IMTRGenericCluster
	// properties:
	// methods:
	CancelPowerAdjustRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	CancelPowerAdjustRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	CancelRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterCancelRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	CancelRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ModifyForecastRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterModifyForecastRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	PauseRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterPauseRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	PowerAdjustRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterPowerAdjustRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ReadAttributeAbsMaxPowerWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAbsMinPowerWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeESACanGenerateWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeESAStateWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeESATypeWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeForecastWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeOptOutStateWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributePowerAdjustmentCapabilityWithParams(params IMTRReadParams) foundation.IDictionary
	RequestConstraintBasedForecastWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ResumeRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterResumeRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ResumeRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	StartTimeAdjustRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
}

// Cluster Device Energy Management This cluster allows a client to manage the power draw of a device. An example of such a client could be an Energy Management System (EMS) which controls an Energy Smart Appliance (ESA).


// Cluster Device Energy Management This cluster allows a client to manage the power draw of a device. An example of such a client could be an Energy Management System (EMS) which controls an Energy Smart Appliance (ESA).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement
type MTRClusterDeviceEnergyManagement struct {
	MTRGenericCluster
}

// MTRClusterDeviceEnergyManagementFrom constructs a [MTRClusterDeviceEnergyManagement] from an unsafe.Pointer.
//
// Cluster Device Energy Management This cluster allows a client to manage the power draw of a device. An example of such a client could be an Energy Management System (EMS) which controls an Energy Smart Appliance (ESA).
func MTRClusterDeviceEnergyManagementFrom(ptr unsafe.Pointer) MTRClusterDeviceEnergyManagement {
	return MTRClusterDeviceEnergyManagement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterDeviceEnergyManagementClass) Alloc() MTRClusterDeviceEnergyManagement {
	rv := objc.Send[MTRClusterDeviceEnergyManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterDeviceEnergyManagementClass) New() MTRClusterDeviceEnergyManagement {
	rv := objc.Send[MTRClusterDeviceEnergyManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterDeviceEnergyManagement) Init() MTRClusterDeviceEnergyManagement {
	rv := objc.Send[MTRClusterDeviceEnergyManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterDeviceEnergyManagement) Autorelease() MTRClusterDeviceEnergyManagement {
	rv := objc.Send[MTRClusterDeviceEnergyManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterDeviceEnergyManagement creates a new MTRClusterDeviceEnergyManagement instance.
func NewMTRClusterDeviceEnergyManagement() MTRClusterDeviceEnergyManagement {
	return getMTRClusterDeviceEnergyManagementClass().New()
}



// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/init(device:endpointID:queue:)
func NewMTRClusterDeviceEnergyManagementWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterDeviceEnergyManagement {
	instance := getMTRClusterDeviceEnergyManagementClass().Alloc()
	rv := objc.Send[MTRClusterDeviceEnergyManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/cancelPowerAdjustRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDeviceEnergyManagement) CancelPowerAdjustRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterCancelPowerAdjustRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelPowerAdjustRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/cancelPowerAdjustRequest(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDeviceEnergyManagement) CancelPowerAdjustRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelPowerAdjustRequestWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/cancelRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDeviceEnergyManagement) CancelRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterCancelRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/cancelRequest(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDeviceEnergyManagement) CancelRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelRequestWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/modifyForecastRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDeviceEnergyManagement) ModifyForecastRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterModifyForecastRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("modifyForecastRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/pauseRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDeviceEnergyManagement) PauseRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterPauseRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("pauseRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/powerAdjustRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDeviceEnergyManagement) PowerAdjustRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterPowerAdjustRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("powerAdjustRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/readAttributeAbsMaxPower(with:)
func (m_ MTRClusterDeviceEnergyManagement) ReadAttributeAbsMaxPowerWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAbsMaxPowerWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/readAttributeAbsMinPower(with:)
func (m_ MTRClusterDeviceEnergyManagement) ReadAttributeAbsMinPowerWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAbsMinPowerWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterDeviceEnergyManagement) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/readAttributeAttributeList(with:)
func (m_ MTRClusterDeviceEnergyManagement) ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/readAttributeClusterRevision(with:)
func (m_ MTRClusterDeviceEnergyManagement) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/readAttributeESACanGenerate(with:)
func (m_ MTRClusterDeviceEnergyManagement) ReadAttributeESACanGenerateWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeESACanGenerateWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/readAttributeESAState(with:)
func (m_ MTRClusterDeviceEnergyManagement) ReadAttributeESAStateWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeESAStateWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/readAttributeESAType(with:)
func (m_ MTRClusterDeviceEnergyManagement) ReadAttributeESATypeWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeESATypeWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/readAttributeFeatureMap(with:)
func (m_ MTRClusterDeviceEnergyManagement) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/readAttributeForecast(with:)
func (m_ MTRClusterDeviceEnergyManagement) ReadAttributeForecastWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeForecastWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterDeviceEnergyManagement) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/readAttributeOptOutState(with:)
func (m_ MTRClusterDeviceEnergyManagement) ReadAttributeOptOutStateWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeOptOutStateWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/readAttributePowerAdjustmentCapability(with:)
func (m_ MTRClusterDeviceEnergyManagement) ReadAttributePowerAdjustmentCapabilityWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributePowerAdjustmentCapabilityWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/requestConstraintBasedForecast(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDeviceEnergyManagement) RequestConstraintBasedForecastWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterRequestConstraintBasedForecastParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("requestConstraintBasedForecastWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/resumeRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDeviceEnergyManagement) ResumeRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterResumeRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("resumeRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/resumeRequest(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDeviceEnergyManagement) ResumeRequestWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("resumeRequestWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterDeviceEnergyManagement/startTimeAdjustRequest(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterDeviceEnergyManagement) StartTimeAdjustRequestWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTRDeviceEnergyManagementClusterStartTimeAdjustRequestParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startTimeAdjustRequestWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


