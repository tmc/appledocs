// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBaseClusterWaterHeaterManagement] class.
var (
	MTRBaseClusterWaterHeaterManagementClass     _MTRBaseClusterWaterHeaterManagementClass
	MTRBaseClusterWaterHeaterManagementClassOnce sync.Once
)

func getMTRBaseClusterWaterHeaterManagementClass() _MTRBaseClusterWaterHeaterManagementClass {
	MTRBaseClusterWaterHeaterManagementClassOnce.Do(func() {
		MTRBaseClusterWaterHeaterManagementClass = _MTRBaseClusterWaterHeaterManagementClass{objc.GetClass("MTRBaseClusterWaterHeaterManagement")}
	})
	return MTRBaseClusterWaterHeaterManagementClass
}

type _MTRBaseClusterWaterHeaterManagementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterWaterHeaterManagement] class.
type IMTRBaseClusterWaterHeaterManagement interface {
	IMTRGenericBaseCluster
	BoostWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	CancelBoostWithCompletion(completion unsafe.Pointer)
	CancelBoostWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeBoostStateWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeEstimatedHeatRequiredWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeHeatDemandWithCompletion(completion unsafe.Pointer)
	ReadAttributeHeaterTypesWithCompletion(completion unsafe.Pointer)
	ReadAttributeTankPercentageWithCompletion(completion unsafe.Pointer)
	ReadAttributeTankVolumeWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeBoostStateWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeEstimatedHeatRequiredWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeHeatDemandWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeHeaterTypesWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeTankPercentageWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeTankVolumeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Water Heater Management
//
// This cluster is used to allow clients to control the operation of a hot water heating appliance so that it can be used with energy management.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement
type MTRBaseClusterWaterHeaterManagement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterWaterHeaterManagementFrom constructs a [MTRBaseClusterWaterHeaterManagement] from an unsafe.Pointer.
//
// Cluster Water Heater Management
func MTRBaseClusterWaterHeaterManagementFrom(ptr unsafe.Pointer) MTRBaseClusterWaterHeaterManagement {
	return MTRBaseClusterWaterHeaterManagement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterWaterHeaterManagementClass) Alloc() MTRBaseClusterWaterHeaterManagement {
	rv := objc.Send[MTRBaseClusterWaterHeaterManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterWaterHeaterManagementClass) New() MTRBaseClusterWaterHeaterManagement {
	rv := objc.Send[MTRBaseClusterWaterHeaterManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterWaterHeaterManagement) Init() MTRBaseClusterWaterHeaterManagement {
	rv := objc.Send[MTRBaseClusterWaterHeaterManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterWaterHeaterManagement) Autorelease() MTRBaseClusterWaterHeaterManagement {
	rv := objc.Send[MTRBaseClusterWaterHeaterManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterWaterHeaterManagement creates a new MTRBaseClusterWaterHeaterManagement instance.
func NewMTRBaseClusterWaterHeaterManagement() MTRBaseClusterWaterHeaterManagement {
	return getMTRBaseClusterWaterHeaterManagementClass().New()
}




// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/init(device:endpointID:queue:)
func NewMTRBaseClusterWaterHeaterManagementWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID foundation.Number, queue unsafe.Pointer) MTRBaseClusterWaterHeaterManagement {
	instance := getMTRBaseClusterWaterHeaterManagementClass().Alloc()
	rv := objc.Send[MTRBaseClusterWaterHeaterManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterManagementClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterManagementClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeBoostState(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterManagementClass) ReadAttributeBoostStateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeBoostStateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterManagementClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeEstimatedHeatRequired(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterManagementClass) ReadAttributeEstimatedHeatRequiredWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEstimatedHeatRequiredWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterManagementClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterManagementClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeHeatDemand(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterManagementClass) ReadAttributeHeatDemandWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeHeatDemandWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeHeaterTypes(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterManagementClass) ReadAttributeHeaterTypesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeHeaterTypesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeTankPercentage(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterManagementClass) ReadAttributeTankPercentageWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTankPercentageWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeTankVolume(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterManagementClass) ReadAttributeTankVolumeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTankVolumeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

// Command Boost
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/boost(with:completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) BoostWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("boostWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/cancelBoost(completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) CancelBoostWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelBoostWithCompletion:"), completion)
}

// Command CancelBoost
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/cancelBoost(with:completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) CancelBoostWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelBoostWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeBoostState(completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) ReadAttributeBoostStateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeBoostStateWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeEstimatedHeatRequired(completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) ReadAttributeEstimatedHeatRequiredWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeEstimatedHeatRequiredWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeHeatDemand(completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) ReadAttributeHeatDemandWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeHeatDemandWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeHeaterTypes(completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) ReadAttributeHeaterTypesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeHeaterTypesWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeTankPercentage(completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) ReadAttributeTankPercentageWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeTankPercentageWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/readAttributeTankVolume(completion:)
func (m_ MTRBaseClusterWaterHeaterManagement) ReadAttributeTankVolumeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeTankVolumeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterManagement) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterManagement) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/subscribeAttributeBoostState(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterManagement) SubscribeAttributeBoostStateWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeBoostStateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterManagement) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/subscribeAttributeEstimatedHeatRequired(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterManagement) SubscribeAttributeEstimatedHeatRequiredWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeEstimatedHeatRequiredWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterManagement) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterManagement) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/subscribeAttributeHeatDemand(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterManagement) SubscribeAttributeHeatDemandWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeHeatDemandWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/subscribeAttributeHeaterTypes(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterManagement) SubscribeAttributeHeaterTypesWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeHeaterTypesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/subscribeAttributeTankPercentage(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterManagement) SubscribeAttributeTankPercentageWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeTankPercentageWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterManagement/subscribeAttributeTankVolume(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterManagement) SubscribeAttributeTankVolumeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeTankVolumeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


