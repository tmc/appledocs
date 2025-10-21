// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBaseClusterICDManagement] class.
var (
	MTRBaseClusterICDManagementClass     _MTRBaseClusterICDManagementClass
	MTRBaseClusterICDManagementClassOnce sync.Once
)

func getMTRBaseClusterICDManagementClass() _MTRBaseClusterICDManagementClass {
	MTRBaseClusterICDManagementClassOnce.Do(func() {
		MTRBaseClusterICDManagementClass = _MTRBaseClusterICDManagementClass{objc.GetClass("MTRBaseClusterICDManagement")}
	})
	return MTRBaseClusterICDManagementClass
}

type _MTRBaseClusterICDManagementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterICDManagement] class.
type IMTRBaseClusterICDManagement interface {
	IMTRGenericBaseCluster
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeActiveModeDurationWithCompletion(completion unsafe.Pointer)
	ReadAttributeActiveModeThresholdWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClientsSupportedPerFabricWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeICDCounterWithCompletion(completion unsafe.Pointer)
	ReadAttributeIdleModeDurationWithCompletion(completion unsafe.Pointer)
	ReadAttributeMaximumCheckInBackOffWithCompletion(completion unsafe.Pointer)
	ReadAttributeOperatingModeWithCompletion(completion unsafe.Pointer)
	ReadAttributeRegisteredClientsWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeUserActiveModeTriggerHintWithCompletion(completion unsafe.Pointer)
	ReadAttributeUserActiveModeTriggerInstructionWithCompletion(completion unsafe.Pointer)
	RegisterClientWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	StayActiveRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeActiveModeDurationWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeActiveModeThresholdWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClientsSupportedPerFabricWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeICDCounterWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeIdleModeDurationWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMaximumCheckInBackOffWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOperatingModeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeRegisteredClientsWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeUserActiveModeTriggerHintWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeUserActiveModeTriggerInstructionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	UnregisterClientWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
}

// Cluster ICD Management
//
// Allows servers to ensure that listed clients are notified when a server is available for communication.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement
type MTRBaseClusterICDManagement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterICDManagementFrom constructs a [MTRBaseClusterICDManagement] from an unsafe.Pointer.
//
// Cluster ICD Management
func MTRBaseClusterICDManagementFrom(ptr unsafe.Pointer) MTRBaseClusterICDManagement {
	return MTRBaseClusterICDManagement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterICDManagementClass) Alloc() MTRBaseClusterICDManagement {
	rv := objc.Send[MTRBaseClusterICDManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterICDManagementClass) New() MTRBaseClusterICDManagement {
	rv := objc.Send[MTRBaseClusterICDManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterICDManagement) Init() MTRBaseClusterICDManagement {
	rv := objc.Send[MTRBaseClusterICDManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterICDManagement) Autorelease() MTRBaseClusterICDManagement {
	rv := objc.Send[MTRBaseClusterICDManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterICDManagement creates a new MTRBaseClusterICDManagement instance.
func NewMTRBaseClusterICDManagement() MTRBaseClusterICDManagement {
	return getMTRBaseClusterICDManagementClass().New()
}




// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/init(device:endpointID:queue:)
func NewMTRBaseClusterICDManagementWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID foundation.Number, queue unsafe.Pointer) MTRBaseClusterICDManagement {
	instance := getMTRBaseClusterICDManagementClass().Alloc()
	rv := objc.Send[MTRBaseClusterICDManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeActiveModeDuration(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeActiveModeDurationWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeActiveModeDurationWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeActiveModeThreshold(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeActiveModeThresholdWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeActiveModeThresholdWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeClientsSupportedPerFabric(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeClientsSupportedPerFabricWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClientsSupportedPerFabricWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeICDCounter(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeICDCounterWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeICDCounterWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeIdleModeDuration(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeIdleModeDurationWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeIdleModeDurationWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeMaximumCheckInBackOff(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeMaximumCheckInBackOffWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaximumCheckInBackOffWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeOperatingMode(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeOperatingModeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOperatingModeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeRegisteredClients(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeRegisteredClientsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeRegisteredClientsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeUserActiveModeTriggerHint(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeUserActiveModeTriggerHintWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUserActiveModeTriggerHintWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeUserActiveModeTriggerInstruction(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterICDManagementClass) ReadAttributeUserActiveModeTriggerInstructionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUserActiveModeTriggerInstructionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeActiveModeDuration(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeActiveModeDurationWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeActiveModeDurationWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeActiveModeThreshold(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeActiveModeThresholdWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeActiveModeThresholdWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeClientsSupportedPerFabric(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeClientsSupportedPerFabricWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClientsSupportedPerFabricWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeICDCounter(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeICDCounterWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeICDCounterWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeIdleModeDuration(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeIdleModeDurationWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeIdleModeDurationWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeMaximumCheckInBackOff(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeMaximumCheckInBackOffWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMaximumCheckInBackOffWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeOperatingMode(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeOperatingModeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOperatingModeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeRegisteredClients(with:completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeRegisteredClientsWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeRegisteredClientsWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeUserActiveModeTriggerHint(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeUserActiveModeTriggerHintWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeUserActiveModeTriggerHintWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/readAttributeUserActiveModeTriggerInstruction(completion:)
func (m_ MTRBaseClusterICDManagement) ReadAttributeUserActiveModeTriggerInstructionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeUserActiveModeTriggerInstructionWithCompletion:"), completion)
}

// Command RegisterClient
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/registerClient(with:completion:)
func (m_ MTRBaseClusterICDManagement) RegisterClientWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("registerClientWithParams:completion:"), params, completion)
}

// Command StayActiveRequest
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/stayActiveRequest(with:completion:)
func (m_ MTRBaseClusterICDManagement) StayActiveRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stayActiveRequestWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeActiveModeDuration(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeActiveModeDurationWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeActiveModeDurationWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeActiveModeThreshold(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeActiveModeThresholdWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeActiveModeThresholdWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeClientsSupportedPerFabric(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeClientsSupportedPerFabricWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClientsSupportedPerFabricWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeICDCounter(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeICDCounterWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeICDCounterWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeIdleModeDuration(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeIdleModeDurationWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeIdleModeDurationWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeMaximumCheckInBackOff(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeMaximumCheckInBackOffWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMaximumCheckInBackOffWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeOperatingMode(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeOperatingModeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOperatingModeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeRegisteredClients(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeRegisteredClientsWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeRegisteredClientsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeUserActiveModeTriggerHint(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeUserActiveModeTriggerHintWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUserActiveModeTriggerHintWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/subscribeAttributeUserActiveModeTriggerInstruction(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterICDManagement) SubscribeAttributeUserActiveModeTriggerInstructionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUserActiveModeTriggerInstructionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

// Command UnregisterClient
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/unregisterClient(with:completion:)
func (m_ MTRBaseClusterICDManagement) UnregisterClientWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("unregisterClientWithParams:completion:"), params, completion)
}


