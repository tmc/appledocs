// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRClusterEnergyEVSE] class.
var (
	MTRClusterEnergyEVSEClass     _MTRClusterEnergyEVSEClass
	MTRClusterEnergyEVSEClassOnce sync.Once
)

func getMTRClusterEnergyEVSEClass() _MTRClusterEnergyEVSEClass {
	MTRClusterEnergyEVSEClassOnce.Do(func() {
		MTRClusterEnergyEVSEClass = _MTRClusterEnergyEVSEClass{objc.GetClass("MTRClusterEnergyEVSE")}
	})
	return MTRClusterEnergyEVSEClass
}

type _MTRClusterEnergyEVSEClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterEnergyEVSE] class.
type IMTRClusterEnergyEVSE interface {
	IMTRGenericCluster
	// properties:
	// methods:
	ClearTargetsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterClearTargetsParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ClearTargetsWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	DisableWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterDisableParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	DisableWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	EnableChargingWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterEnableChargingParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	GetTargetsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterGetTargetsParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	GetTargetsWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeApproximateEVEfficiencyWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeChargingEnabledUntilWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeCircuitCapacityWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeFaultStateWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeMaximumChargeCurrentWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeMinimumChargeCurrentWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeNextChargeRequiredEnergyWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeNextChargeStartTimeWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeNextChargeTargetSoCWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeNextChargeTargetTimeWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeRandomizationDelayWindowWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSessionDurationWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSessionEnergyChargedWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSessionIDWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeStateWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeSupplyStateWithParams(params IMTRReadParams) foundation.IDictionary
	ReadAttributeUserMaximumChargeCurrentWithParams(params IMTRReadParams) foundation.IDictionary
	SetTargetsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterSetTargetsParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	StartDiagnosticsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterStartDiagnosticsParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	StartDiagnosticsWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeApproximateEVEfficiencyWithValueExpectedValueInterval(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */)
	WriteAttributeApproximateEVEfficiencyWithValueExpectedValueIntervalParams(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams)
	WriteAttributeRandomizationDelayWindowWithValueExpectedValueInterval(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */)
	WriteAttributeRandomizationDelayWindowWithValueExpectedValueIntervalParams(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams)
	WriteAttributeUserMaximumChargeCurrentWithValueExpectedValueInterval(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */)
	WriteAttributeUserMaximumChargeCurrentWithValueExpectedValueIntervalParams(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams)
}

// Cluster Energy EVSE Electric Vehicle Supply Equipment (EVSE) is equipment used to charge an Electric Vehicle (EV) or Plug-In Hybrid Electric Vehicle. This cluster provides an interface to the functionality of Electric Vehicle Supply Equipment (EVSE) management.


// Cluster Energy EVSE Electric Vehicle Supply Equipment (EVSE) is equipment used to charge an Electric Vehicle (EV) or Plug-In Hybrid Electric Vehicle. This cluster provides an interface to the functionality of Electric Vehicle Supply Equipment (EVSE) management.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE
type MTRClusterEnergyEVSE struct {
	MTRGenericCluster
}

// MTRClusterEnergyEVSEFrom constructs a [MTRClusterEnergyEVSE] from an unsafe.Pointer.
//
// Cluster Energy EVSE Electric Vehicle Supply Equipment (EVSE) is equipment used to charge an Electric Vehicle (EV) or Plug-In Hybrid Electric Vehicle. This cluster provides an interface to the functionality of Electric Vehicle Supply Equipment (EVSE) management.
func MTRClusterEnergyEVSEFrom(ptr unsafe.Pointer) MTRClusterEnergyEVSE {
	return MTRClusterEnergyEVSE{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterEnergyEVSEClass) Alloc() MTRClusterEnergyEVSE {
	rv := objc.Send[MTRClusterEnergyEVSE](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterEnergyEVSEClass) New() MTRClusterEnergyEVSE {
	rv := objc.Send[MTRClusterEnergyEVSE](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterEnergyEVSE) Init() MTRClusterEnergyEVSE {
	rv := objc.Send[MTRClusterEnergyEVSE](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterEnergyEVSE) Autorelease() MTRClusterEnergyEVSE {
	rv := objc.Send[MTRClusterEnergyEVSE](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterEnergyEVSE creates a new MTRClusterEnergyEVSE instance.
func NewMTRClusterEnergyEVSE() MTRClusterEnergyEVSE {
	return getMTRClusterEnergyEVSEClass().New()
}



// For all instance methods that take a completion (i.e. command invocations), the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/init(device:endpointID:queue:)
func NewMTRClusterEnergyEVSEWithDeviceEndpointIDQueue(device IMTRDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRClusterEnergyEVSE {
	instance := getMTRClusterEnergyEVSEClass().Alloc()
	rv := objc.Send[MTRClusterEnergyEVSE](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/clearTargets(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) ClearTargetsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterClearTargetsParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearTargetsWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/clearTargets(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) ClearTargetsWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearTargetsWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/disable(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) DisableWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterDisableParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("disableWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/disable(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) DisableWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("disableWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/enableCharging(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) EnableChargingWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterEnableChargingParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("enableChargingWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/getTargetsWith(_:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) GetTargetsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterGetTargetsParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getTargetsWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/getTargetsWithExpectedValues(_:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) GetTargetsWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getTargetsWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeApproximateEVEfficiency(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeApproximateEVEfficiencyWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeApproximateEVEfficiencyWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeAttributeList(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeAttributeListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeChargingEnabledUntil(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeChargingEnabledUntilWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeChargingEnabledUntilWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeCircuitCapacity(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeCircuitCapacityWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeCircuitCapacityWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeClusterRevision(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeFaultState(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeFaultStateWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFaultStateWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeFeatureMap(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeFeatureMapWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeMaximumChargeCurrent(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeMaximumChargeCurrentWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeMaximumChargeCurrentWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeMinimumChargeCurrent(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeMinimumChargeCurrentWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeMinimumChargeCurrentWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeNextChargeRequiredEnergy(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeNextChargeRequiredEnergyWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeNextChargeRequiredEnergyWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeNextChargeStartTime(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeNextChargeStartTimeWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeNextChargeStartTimeWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeNextChargeTargetSoC(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeNextChargeTargetSoCWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeNextChargeTargetSoCWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeNextChargeTargetTime(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeNextChargeTargetTimeWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeNextChargeTargetTimeWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeRandomizationDelayWindow(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeRandomizationDelayWindowWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeRandomizationDelayWindowWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeSessionDuration(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeSessionDurationWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSessionDurationWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeSessionEnergyCharged(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeSessionEnergyChargedWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSessionEnergyChargedWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeSessionID(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeSessionIDWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSessionIDWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeState(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeStateWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeStateWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeSupplyState(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeSupplyStateWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeSupplyStateWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeUserMaximumChargeCurrent(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeUserMaximumChargeCurrentWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeUserMaximumChargeCurrentWithParams:"), params)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/setTargetsWith(_:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) SetTargetsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterSetTargetsParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetsWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/startDiagnostics(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) StartDiagnosticsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterStartDiagnosticsParams, expectedDataValueDictionaries foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startDiagnosticsWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/startDiagnostics(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) StartDiagnosticsWithExpectedValuesExpectedValueIntervalCompletion(expectedValues foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startDiagnosticsWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/writeAttributeApproximateEVEfficiency(withValue:expectedValueInterval:)
func (m_ MTRClusterEnergyEVSE) WriteAttributeApproximateEVEfficiencyWithValueExpectedValueInterval(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeApproximateEVEfficiencyWithValue:expectedValueInterval:"), dataValueDictionary, expectedValueIntervalMs)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/writeAttributeApproximateEVEfficiency(withValue:expectedValueInterval:params:)
func (m_ MTRClusterEnergyEVSE) WriteAttributeApproximateEVEfficiencyWithValueExpectedValueIntervalParams(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeApproximateEVEfficiencyWithValue:expectedValueInterval:params:"), dataValueDictionary, expectedValueIntervalMs, params)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/writeAttributeRandomizationDelayWindow(withValue:expectedValueInterval:)
func (m_ MTRClusterEnergyEVSE) WriteAttributeRandomizationDelayWindowWithValueExpectedValueInterval(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeRandomizationDelayWindowWithValue:expectedValueInterval:"), dataValueDictionary, expectedValueIntervalMs)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/writeAttributeRandomizationDelayWindow(withValue:expectedValueInterval:params:)
func (m_ MTRClusterEnergyEVSE) WriteAttributeRandomizationDelayWindowWithValueExpectedValueIntervalParams(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeRandomizationDelayWindowWithValue:expectedValueInterval:params:"), dataValueDictionary, expectedValueIntervalMs, params)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/writeAttributeUserMaximumChargeCurrent(withValue:expectedValueInterval:)
func (m_ MTRClusterEnergyEVSE) WriteAttributeUserMaximumChargeCurrentWithValueExpectedValueInterval(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUserMaximumChargeCurrentWithValue:expectedValueInterval:"), dataValueDictionary, expectedValueIntervalMs)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/writeAttributeUserMaximumChargeCurrent(withValue:expectedValueInterval:params:)
func (m_ MTRClusterEnergyEVSE) WriteAttributeUserMaximumChargeCurrentWithValueExpectedValueIntervalParams(dataValueDictionary foundation.IDictionary, expectedValueIntervalMs objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUserMaximumChargeCurrentWithValue:expectedValueInterval:params:"), dataValueDictionary, expectedValueIntervalMs, params)
}


