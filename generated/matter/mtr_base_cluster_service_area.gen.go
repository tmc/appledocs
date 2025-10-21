// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterServiceArea] class.
var (
	MTRBaseClusterServiceAreaClass     _MTRBaseClusterServiceAreaClass
	MTRBaseClusterServiceAreaClassOnce sync.Once
)

func getMTRBaseClusterServiceAreaClass() _MTRBaseClusterServiceAreaClass {
	MTRBaseClusterServiceAreaClassOnce.Do(func() {
		MTRBaseClusterServiceAreaClass = _MTRBaseClusterServiceAreaClass{objc.GetClass("MTRBaseClusterServiceArea")}
	})
	return MTRBaseClusterServiceAreaClass
}

type _MTRBaseClusterServiceAreaClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterServiceArea] class.
type IMTRBaseClusterServiceArea interface {
	IMTRGenericBaseCluster
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentAreaWithCompletion(completion unsafe.Pointer)
	ReadAttributeEstimatedEndTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeProgressWithCompletion(completion unsafe.Pointer)
	ReadAttributeSelectedAreasWithCompletion(completion unsafe.Pointer)
	ReadAttributeSupportedAreasWithCompletion(completion unsafe.Pointer)
	ReadAttributeSupportedMapsWithCompletion(completion unsafe.Pointer)
	SelectAreasWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SkipAreaWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentAreaWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeEstimatedEndTimeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeProgressWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSelectedAreasWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSupportedAreasWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSupportedMapsWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Service Area
//
// The Service Area cluster provides an interface for controlling the areas where a device should operate, and for querying the current area being serviced.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea
type MTRBaseClusterServiceArea struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterServiceAreaFrom constructs a [MTRBaseClusterServiceArea] from an unsafe.Pointer.
//
// Cluster Service Area
func MTRBaseClusterServiceAreaFrom(ptr unsafe.Pointer) MTRBaseClusterServiceArea {
	return MTRBaseClusterServiceArea{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterServiceAreaClass) Alloc() MTRBaseClusterServiceArea {
	rv := objc.Send[MTRBaseClusterServiceArea](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterServiceAreaClass) New() MTRBaseClusterServiceArea {
	rv := objc.Send[MTRBaseClusterServiceArea](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterServiceArea) Init() MTRBaseClusterServiceArea {
	rv := objc.Send[MTRBaseClusterServiceArea](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterServiceArea) Autorelease() MTRBaseClusterServiceArea {
	rv := objc.Send[MTRBaseClusterServiceArea](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterServiceArea creates a new MTRBaseClusterServiceArea instance.
func NewMTRBaseClusterServiceArea() MTRBaseClusterServiceArea {
	return getMTRBaseClusterServiceAreaClass().New()
}


// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/init(device:endpointID:queue:)
func NewMTRBaseClusterServiceAreaWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRBaseClusterServiceArea {
	instance := getMTRBaseClusterServiceAreaClass().Alloc()
	rv := objc.Send[MTRBaseClusterServiceArea](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterServiceAreaClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterServiceAreaClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterServiceAreaClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeCurrentArea(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterServiceAreaClass) ReadAttributeCurrentAreaWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentAreaWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeEstimatedEndTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterServiceAreaClass) ReadAttributeEstimatedEndTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEstimatedEndTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterServiceAreaClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterServiceAreaClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeProgress(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterServiceAreaClass) ReadAttributeProgressWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeProgressWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeSelectedAreas(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterServiceAreaClass) ReadAttributeSelectedAreasWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSelectedAreasWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeSupportedAreas(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterServiceAreaClass) ReadAttributeSupportedAreasWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedAreasWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeSupportedMaps(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterServiceAreaClass) ReadAttributeSupportedMapsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedMapsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterServiceArea) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterServiceArea) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterServiceArea) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeCurrentArea(completion:)
func (m_ MTRBaseClusterServiceArea) ReadAttributeCurrentAreaWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentAreaWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeEstimatedEndTime(completion:)
func (m_ MTRBaseClusterServiceArea) ReadAttributeEstimatedEndTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeEstimatedEndTimeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterServiceArea) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterServiceArea) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeProgress(completion:)
func (m_ MTRBaseClusterServiceArea) ReadAttributeProgressWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeProgressWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeSelectedAreas(completion:)
func (m_ MTRBaseClusterServiceArea) ReadAttributeSelectedAreasWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSelectedAreasWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeSupportedAreas(completion:)
func (m_ MTRBaseClusterServiceArea) ReadAttributeSupportedAreasWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupportedAreasWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/readAttributeSupportedMaps(completion:)
func (m_ MTRBaseClusterServiceArea) ReadAttributeSupportedMapsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupportedMapsWithCompletion:"), completion)
}

// Command SelectAreas
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/selectAreas(with:completion:)
func (m_ MTRBaseClusterServiceArea) SelectAreasWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("selectAreasWithParams:completion:"), params, completion)
}

// Command SkipArea
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/skip(with:completion:)
func (m_ MTRBaseClusterServiceArea) SkipAreaWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("skipAreaWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterServiceArea) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterServiceArea) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterServiceArea) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/subscribeAttributeCurrentArea(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterServiceArea) SubscribeAttributeCurrentAreaWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentAreaWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/subscribeAttributeEstimatedEndTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterServiceArea) SubscribeAttributeEstimatedEndTimeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeEstimatedEndTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterServiceArea) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterServiceArea) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/subscribeAttributeProgress(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterServiceArea) SubscribeAttributeProgressWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeProgressWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/subscribeAttributeSelectedAreas(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterServiceArea) SubscribeAttributeSelectedAreasWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSelectedAreasWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/subscribeAttributeSupportedAreas(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterServiceArea) SubscribeAttributeSupportedAreasWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupportedAreasWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/subscribeAttributeSupportedMaps(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterServiceArea) SubscribeAttributeSupportedMapsWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupportedMapsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


