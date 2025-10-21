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
	ClearTargetsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterClearTargetsParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	ClearTargetsWithExpectedValuesExpectedValueIntervalCompletion(expectedValues []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	DisableWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterDisableParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	DisableWithExpectedValuesExpectedValueIntervalCompletion(expectedValues []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	EnableChargingWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterEnableChargingParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	GetTargetsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterGetTargetsParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	GetTargetsWithExpectedValuesExpectedValueIntervalCompletion(expectedValues []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeApproximateEVEfficiencyWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeChargingEnabledUntilWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeCircuitCapacityWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeFaultStateWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeMaximumChargeCurrentWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeMinimumChargeCurrentWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeNextChargeRequiredEnergyWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeNextChargeStartTimeWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeNextChargeTargetSoCWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeNextChargeTargetTimeWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeRandomizationDelayWindowWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeSessionDurationWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeSessionEnergyChargedWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeSessionIDWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeStateWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeSupplyStateWithParams(params IMTRReadParams) unsafe.Pointer
	ReadAttributeUserMaximumChargeCurrentWithParams(params IMTRReadParams) unsafe.Pointer
	SetTargetsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterSetTargetsParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	StartDiagnosticsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterStartDiagnosticsParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	StartDiagnosticsWithExpectedValuesExpectedValueIntervalCompletion(expectedValues []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer)
	WriteAttributeApproximateEVEfficiencyWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber)
	WriteAttributeApproximateEVEfficiencyWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber, params IMTRWriteParams)
	WriteAttributeRandomizationDelayWindowWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber)
	WriteAttributeRandomizationDelayWindowWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber, params IMTRWriteParams)
	WriteAttributeUserMaximumChargeCurrentWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber)
	WriteAttributeUserMaximumChargeCurrentWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber, params IMTRWriteParams)
}

// Cluster Energy EVSE Electric Vehicle Supply Equipment (EVSE) is equipment used to charge an Electric Vehicle (EV) or Plug-In Hybrid Electric Vehicle. This cluster provides an interface to the functionality of Electric Vehicle Supply Equipment (EVSE) management.
//
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/init(device:endpointID:queue:)
func NewMTRClusterEnergyEVSEWithDeviceEndpointIDQueue(device IMTRDevice, endpointID foundation.INumber, queue unsafe.Pointer) MTRClusterEnergyEVSE {
	instance := getMTRClusterEnergyEVSEClass().Alloc()
	rv := objc.Send[MTRClusterEnergyEVSE](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/clearTargets(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) ClearTargetsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterClearTargetsParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearTargetsWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/clearTargets(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) ClearTargetsWithExpectedValuesExpectedValueIntervalCompletion(expectedValues []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearTargetsWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/disable(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) DisableWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterDisableParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("disableWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/disable(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) DisableWithExpectedValuesExpectedValueIntervalCompletion(expectedValues []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("disableWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/enableCharging(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) EnableChargingWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterEnableChargingParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("enableChargingWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/getTargetsWith(_:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) GetTargetsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterGetTargetsParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getTargetsWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/getTargetsWithExpectedValues(_:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) GetTargetsWithExpectedValuesExpectedValueIntervalCompletion(expectedValues []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getTargetsWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeAcceptedCommandList(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeAcceptedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeApproximateEVEfficiency(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeApproximateEVEfficiencyWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeApproximateEVEfficiencyWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeAttributeList(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeAttributeListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeAttributeListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeChargingEnabledUntil(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeChargingEnabledUntilWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeChargingEnabledUntilWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeCircuitCapacity(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeCircuitCapacityWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeCircuitCapacityWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeClusterRevision(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeClusterRevisionWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeClusterRevisionWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeFaultState(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeFaultStateWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFaultStateWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeFeatureMap(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeFeatureMapWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeFeatureMapWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeMaximumChargeCurrent(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeMaximumChargeCurrentWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeMaximumChargeCurrentWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeMinimumChargeCurrent(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeMinimumChargeCurrentWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeMinimumChargeCurrentWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeNextChargeRequiredEnergy(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeNextChargeRequiredEnergyWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeNextChargeRequiredEnergyWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeNextChargeStartTime(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeNextChargeStartTimeWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeNextChargeStartTimeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeNextChargeTargetSoC(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeNextChargeTargetSoCWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeNextChargeTargetSoCWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeNextChargeTargetTime(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeNextChargeTargetTimeWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeNextChargeTargetTimeWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeRandomizationDelayWindow(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeRandomizationDelayWindowWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeRandomizationDelayWindowWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeSessionDuration(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeSessionDurationWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSessionDurationWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeSessionEnergyCharged(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeSessionEnergyChargedWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSessionEnergyChargedWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeSessionID(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeSessionIDWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSessionIDWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeState(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeStateWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeStateWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeSupplyState(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeSupplyStateWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeSupplyStateWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/readAttributeUserMaximumChargeCurrent(with:)
func (m_ MTRClusterEnergyEVSE) ReadAttributeUserMaximumChargeCurrentWithParams(params IMTRReadParams) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readAttributeUserMaximumChargeCurrentWithParams:"), params)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/setTargetsWith(_:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) SetTargetsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterSetTargetsParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetsWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/startDiagnostics(with:expectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) StartDiagnosticsWithParamsExpectedValuesExpectedValueIntervalCompletion(params IMTREnergyEVSEClusterStartDiagnosticsParams, expectedDataValueDictionaries []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startDiagnosticsWithParams:expectedValues:expectedValueInterval:completion:"), params, expectedDataValueDictionaries, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/startDiagnostics(withExpectedValues:expectedValueInterval:completion:)
func (m_ MTRClusterEnergyEVSE) StartDiagnosticsWithExpectedValuesExpectedValueIntervalCompletion(expectedValues []foundation.IDictionary, expectedValueIntervalMs foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startDiagnosticsWithExpectedValues:expectedValueInterval:completion:"), expectedValues, expectedValueIntervalMs, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/writeAttributeApproximateEVEfficiency(withValue:expectedValueInterval:)
func (m_ MTRClusterEnergyEVSE) WriteAttributeApproximateEVEfficiencyWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeApproximateEVEfficiencyWithValue:expectedValueInterval:"), dataValueDictionary, expectedValueIntervalMs)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/writeAttributeApproximateEVEfficiency(withValue:expectedValueInterval:params:)
func (m_ MTRClusterEnergyEVSE) WriteAttributeApproximateEVEfficiencyWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber, params IMTRWriteParams) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeApproximateEVEfficiencyWithValue:expectedValueInterval:params:"), dataValueDictionary, expectedValueIntervalMs, params)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/writeAttributeRandomizationDelayWindow(withValue:expectedValueInterval:)
func (m_ MTRClusterEnergyEVSE) WriteAttributeRandomizationDelayWindowWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeRandomizationDelayWindowWithValue:expectedValueInterval:"), dataValueDictionary, expectedValueIntervalMs)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/writeAttributeRandomizationDelayWindow(withValue:expectedValueInterval:params:)
func (m_ MTRClusterEnergyEVSE) WriteAttributeRandomizationDelayWindowWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber, params IMTRWriteParams) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeRandomizationDelayWindowWithValue:expectedValueInterval:params:"), dataValueDictionary, expectedValueIntervalMs, params)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/writeAttributeUserMaximumChargeCurrent(withValue:expectedValueInterval:)
func (m_ MTRClusterEnergyEVSE) WriteAttributeUserMaximumChargeCurrentWithValueExpectedValueInterval(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUserMaximumChargeCurrentWithValue:expectedValueInterval:"), dataValueDictionary, expectedValueIntervalMs)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterEnergyEVSE/writeAttributeUserMaximumChargeCurrent(withValue:expectedValueInterval:params:)
func (m_ MTRClusterEnergyEVSE) WriteAttributeUserMaximumChargeCurrentWithValueExpectedValueIntervalParams(dataValueDictionary unsafe.Pointer, expectedValueIntervalMs foundation.INumber, params IMTRWriteParams) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUserMaximumChargeCurrentWithValue:expectedValueInterval:params:"), dataValueDictionary, expectedValueIntervalMs, params)
}


