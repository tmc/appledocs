// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBaseClusterEnergyEVSE] class.
var (
	MTRBaseClusterEnergyEVSEClass     _MTRBaseClusterEnergyEVSEClass
	MTRBaseClusterEnergyEVSEClassOnce sync.Once
)

func getMTRBaseClusterEnergyEVSEClass() _MTRBaseClusterEnergyEVSEClass {
	MTRBaseClusterEnergyEVSEClassOnce.Do(func() {
		MTRBaseClusterEnergyEVSEClass = _MTRBaseClusterEnergyEVSEClass{objc.GetClass("MTRBaseClusterEnergyEVSE")}
	})
	return MTRBaseClusterEnergyEVSEClass
}

type _MTRBaseClusterEnergyEVSEClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterEnergyEVSE] class.
type IMTRBaseClusterEnergyEVSE interface {
	IMTRGenericBaseCluster
	ClearTargetsWithCompletion(completion unsafe.Pointer)
	ClearTargetsWithParamsCompletion(params IMTREnergyEVSEClusterClearTargetsParams, completion unsafe.Pointer)
	DisableWithCompletion(completion unsafe.Pointer)
	DisableWithParamsCompletion(params IMTREnergyEVSEClusterDisableParams, completion unsafe.Pointer)
	EnableChargingWithParamsCompletion(params IMTREnergyEVSEClusterEnableChargingParams, completion unsafe.Pointer)
	GetTargetsWithParamsCompletion(params IMTREnergyEVSEClusterGetTargetsParams, completion unsafe.Pointer)
	GetTargetsWithCompletion(completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeApproximateEVEfficiencyWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeChargingEnabledUntilWithCompletion(completion unsafe.Pointer)
	ReadAttributeCircuitCapacityWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFaultStateWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeMaximumChargeCurrentWithCompletion(completion unsafe.Pointer)
	ReadAttributeMinimumChargeCurrentWithCompletion(completion unsafe.Pointer)
	ReadAttributeNextChargeRequiredEnergyWithCompletion(completion unsafe.Pointer)
	ReadAttributeNextChargeStartTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeNextChargeTargetSoCWithCompletion(completion unsafe.Pointer)
	ReadAttributeNextChargeTargetTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeRandomizationDelayWindowWithCompletion(completion unsafe.Pointer)
	ReadAttributeSessionDurationWithCompletion(completion unsafe.Pointer)
	ReadAttributeSessionEnergyChargedWithCompletion(completion unsafe.Pointer)
	ReadAttributeSessionIDWithCompletion(completion unsafe.Pointer)
	ReadAttributeStateWithCompletion(completion unsafe.Pointer)
	ReadAttributeSupplyStateWithCompletion(completion unsafe.Pointer)
	ReadAttributeUserMaximumChargeCurrentWithCompletion(completion unsafe.Pointer)
	SetTargetsWithParamsCompletion(params IMTREnergyEVSEClusterSetTargetsParams, completion unsafe.Pointer)
	StartDiagnosticsWithCompletion(completion unsafe.Pointer)
	StartDiagnosticsWithParamsCompletion(params IMTREnergyEVSEClusterStartDiagnosticsParams, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeApproximateEVEfficiencyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeChargingEnabledUntilWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCircuitCapacityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFaultStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMaximumChargeCurrentWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMinimumChargeCurrentWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNextChargeRequiredEnergyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNextChargeStartTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNextChargeTargetSoCWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNextChargeTargetTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeRandomizationDelayWindowWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSessionDurationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSessionEnergyChargedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSessionIDWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSupplyStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeUserMaximumChargeCurrentWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	WriteAttributeApproximateEVEfficiencyWithValueCompletion(value foundation.INumber, completion unsafe.Pointer)
	WriteAttributeApproximateEVEfficiencyWithValueParamsCompletion(value foundation.INumber, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeRandomizationDelayWindowWithValueCompletion(value foundation.INumber, completion unsafe.Pointer)
	WriteAttributeRandomizationDelayWindowWithValueParamsCompletion(value foundation.INumber, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeUserMaximumChargeCurrentWithValueCompletion(value foundation.INumber, completion unsafe.Pointer)
	WriteAttributeUserMaximumChargeCurrentWithValueParamsCompletion(value foundation.INumber, params IMTRWriteParams, completion unsafe.Pointer)
}

// Cluster Energy EVSE
//
// Electric Vehicle Supply Equipment (EVSE) is equipment used to charge an Electric Vehicle (EV) or Plug-In Hybrid Electric Vehicle. This cluster provides an interface to the functionality of Electric Vehicle Supply Equipment (EVSE) management.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE
type MTRBaseClusterEnergyEVSE struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterEnergyEVSEFrom constructs a [MTRBaseClusterEnergyEVSE] from an unsafe.Pointer.
//
// Cluster Energy EVSE
func MTRBaseClusterEnergyEVSEFrom(ptr unsafe.Pointer) MTRBaseClusterEnergyEVSE {
	return MTRBaseClusterEnergyEVSE{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterEnergyEVSEClass) Alloc() MTRBaseClusterEnergyEVSE {
	rv := objc.Send[MTRBaseClusterEnergyEVSE](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterEnergyEVSEClass) New() MTRBaseClusterEnergyEVSE {
	rv := objc.Send[MTRBaseClusterEnergyEVSE](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterEnergyEVSE) Init() MTRBaseClusterEnergyEVSE {
	rv := objc.Send[MTRBaseClusterEnergyEVSE](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterEnergyEVSE) Autorelease() MTRBaseClusterEnergyEVSE {
	rv := objc.Send[MTRBaseClusterEnergyEVSE](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterEnergyEVSE creates a new MTRBaseClusterEnergyEVSE instance.
func NewMTRBaseClusterEnergyEVSE() MTRBaseClusterEnergyEVSE {
	return getMTRBaseClusterEnergyEVSEClass().New()
}




// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/init(device:endpointID:queue:)
func NewMTRBaseClusterEnergyEVSEWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID foundation.INumber, queue unsafe.Pointer) MTRBaseClusterEnergyEVSE {
	instance := getMTRBaseClusterEnergyEVSEClass().Alloc()
	rv := objc.Send[MTRBaseClusterEnergyEVSE](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeApproximateEVEfficiency(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeApproximateEVEfficiencyWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeApproximateEVEfficiencyWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeChargingEnabledUntil(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeChargingEnabledUntilWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeChargingEnabledUntilWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeCircuitCapacity(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeCircuitCapacityWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCircuitCapacityWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeFaultState(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeFaultStateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFaultStateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeMaximumChargeCurrent(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeMaximumChargeCurrentWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaximumChargeCurrentWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeMinimumChargeCurrent(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeMinimumChargeCurrentWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinimumChargeCurrentWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeNextChargeRequiredEnergy(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeNextChargeRequiredEnergyWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNextChargeRequiredEnergyWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeNextChargeStartTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeNextChargeStartTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNextChargeStartTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeNextChargeTargetSoC(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeNextChargeTargetSoCWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNextChargeTargetSoCWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeNextChargeTargetTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeNextChargeTargetTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNextChargeTargetTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeRandomizationDelayWindow(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeRandomizationDelayWindowWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeRandomizationDelayWindowWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeSessionDuration(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeSessionDurationWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSessionDurationWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeSessionEnergyCharged(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeSessionEnergyChargedWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSessionEnergyChargedWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeSessionID(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeSessionIDWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSessionIDWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeState(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeStateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeStateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeSupplyState(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeSupplyStateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupplyStateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeUserMaximumChargeCurrent(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterEnergyEVSEClass) ReadAttributeUserMaximumChargeCurrentWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUserMaximumChargeCurrentWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/clearTargets(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ClearTargetsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearTargetsWithCompletion:"), completion)
}

// Command ClearTargets
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/clearTargets(with:completion:)
func (m_ MTRBaseClusterEnergyEVSE) ClearTargetsWithParamsCompletion(params IMTREnergyEVSEClusterClearTargetsParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("clearTargetsWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/disable(completion:)
func (m_ MTRBaseClusterEnergyEVSE) DisableWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("disableWithCompletion:"), completion)
}

// Command Disable
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/disable(with:completion:)
func (m_ MTRBaseClusterEnergyEVSE) DisableWithParamsCompletion(params IMTREnergyEVSEClusterDisableParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("disableWithParams:completion:"), params, completion)
}

// Command EnableCharging
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/enableCharging(with:completion:)
func (m_ MTRBaseClusterEnergyEVSE) EnableChargingWithParamsCompletion(params IMTREnergyEVSEClusterEnableChargingParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("enableChargingWithParams:completion:"), params, completion)
}

// Command GetTargets
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/getTargetsWith(_:completion:)
func (m_ MTRBaseClusterEnergyEVSE) GetTargetsWithParamsCompletion(params IMTREnergyEVSEClusterGetTargetsParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getTargetsWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/getTargetsWithCompletion(_:)
func (m_ MTRBaseClusterEnergyEVSE) GetTargetsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getTargetsWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeApproximateEVEfficiency(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeApproximateEVEfficiencyWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeApproximateEVEfficiencyWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeChargingEnabledUntil(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeChargingEnabledUntilWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeChargingEnabledUntilWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeCircuitCapacity(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeCircuitCapacityWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCircuitCapacityWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeFaultState(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeFaultStateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFaultStateWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeMaximumChargeCurrent(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeMaximumChargeCurrentWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMaximumChargeCurrentWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeMinimumChargeCurrent(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeMinimumChargeCurrentWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMinimumChargeCurrentWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeNextChargeRequiredEnergy(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeNextChargeRequiredEnergyWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNextChargeRequiredEnergyWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeNextChargeStartTime(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeNextChargeStartTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNextChargeStartTimeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeNextChargeTargetSoC(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeNextChargeTargetSoCWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNextChargeTargetSoCWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeNextChargeTargetTime(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeNextChargeTargetTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNextChargeTargetTimeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeRandomizationDelayWindow(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeRandomizationDelayWindowWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeRandomizationDelayWindowWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeSessionDuration(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeSessionDurationWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSessionDurationWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeSessionEnergyCharged(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeSessionEnergyChargedWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSessionEnergyChargedWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeSessionID(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeSessionIDWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSessionIDWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeState(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeStateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeStateWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeSupplyState(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeSupplyStateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupplyStateWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/readAttributeUserMaximumChargeCurrent(completion:)
func (m_ MTRBaseClusterEnergyEVSE) ReadAttributeUserMaximumChargeCurrentWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeUserMaximumChargeCurrentWithCompletion:"), completion)
}

// Command SetTargets
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/setTargetsWith(_:completion:)
func (m_ MTRBaseClusterEnergyEVSE) SetTargetsWithParamsCompletion(params IMTREnergyEVSEClusterSetTargetsParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetsWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/startDiagnostics(completion:)
func (m_ MTRBaseClusterEnergyEVSE) StartDiagnosticsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startDiagnosticsWithCompletion:"), completion)
}

// Command StartDiagnostics
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/startDiagnostics(with:completion:)
func (m_ MTRBaseClusterEnergyEVSE) StartDiagnosticsWithParamsCompletion(params IMTREnergyEVSEClusterStartDiagnosticsParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startDiagnosticsWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeApproximateEVEfficiency(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeApproximateEVEfficiencyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeApproximateEVEfficiencyWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeChargingEnabledUntil(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeChargingEnabledUntilWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeChargingEnabledUntilWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeCircuitCapacity(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeCircuitCapacityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCircuitCapacityWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeFaultState(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeFaultStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFaultStateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeMaximumChargeCurrent(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeMaximumChargeCurrentWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMaximumChargeCurrentWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeMinimumChargeCurrent(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeMinimumChargeCurrentWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMinimumChargeCurrentWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeNextChargeRequiredEnergy(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeNextChargeRequiredEnergyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNextChargeRequiredEnergyWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeNextChargeStartTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeNextChargeStartTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNextChargeStartTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeNextChargeTargetSoC(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeNextChargeTargetSoCWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNextChargeTargetSoCWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeNextChargeTargetTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeNextChargeTargetTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNextChargeTargetTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeRandomizationDelayWindow(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeRandomizationDelayWindowWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeRandomizationDelayWindowWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeSessionDuration(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeSessionDurationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSessionDurationWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeSessionEnergyCharged(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeSessionEnergyChargedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSessionEnergyChargedWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeSessionID(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeSessionIDWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSessionIDWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeState(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeStateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeSupplyState(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeSupplyStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupplyStateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/subscribeAttributeUserMaximumChargeCurrent(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterEnergyEVSE) SubscribeAttributeUserMaximumChargeCurrentWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUserMaximumChargeCurrentWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/writeAttributeApproximateEVEfficiency(withValue:completion:)
func (m_ MTRBaseClusterEnergyEVSE) WriteAttributeApproximateEVEfficiencyWithValueCompletion(value foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeApproximateEVEfficiencyWithValue:completion:"), value, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/writeAttributeApproximateEVEfficiency(withValue:params:completion:)
func (m_ MTRBaseClusterEnergyEVSE) WriteAttributeApproximateEVEfficiencyWithValueParamsCompletion(value foundation.INumber, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeApproximateEVEfficiencyWithValue:params:completion:"), value, params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/writeAttributeRandomizationDelayWindow(withValue:completion:)
func (m_ MTRBaseClusterEnergyEVSE) WriteAttributeRandomizationDelayWindowWithValueCompletion(value foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeRandomizationDelayWindowWithValue:completion:"), value, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/writeAttributeRandomizationDelayWindow(withValue:params:completion:)
func (m_ MTRBaseClusterEnergyEVSE) WriteAttributeRandomizationDelayWindowWithValueParamsCompletion(value foundation.INumber, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeRandomizationDelayWindowWithValue:params:completion:"), value, params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/writeAttributeUserMaximumChargeCurrent(withValue:completion:)
func (m_ MTRBaseClusterEnergyEVSE) WriteAttributeUserMaximumChargeCurrentWithValueCompletion(value foundation.INumber, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUserMaximumChargeCurrentWithValue:completion:"), value, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterEnergyEVSE/writeAttributeUserMaximumChargeCurrent(withValue:params:completion:)
func (m_ MTRBaseClusterEnergyEVSE) WriteAttributeUserMaximumChargeCurrentWithValueParamsCompletion(value foundation.INumber, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeUserMaximumChargeCurrentWithValue:params:completion:"), value, params, completion)
}


