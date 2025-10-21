// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBaseClusterDishwasherAlarm] class.
var (
	MTRBaseClusterDishwasherAlarmClass     _MTRBaseClusterDishwasherAlarmClass
	MTRBaseClusterDishwasherAlarmClassOnce sync.Once
)

func getMTRBaseClusterDishwasherAlarmClass() _MTRBaseClusterDishwasherAlarmClass {
	MTRBaseClusterDishwasherAlarmClassOnce.Do(func() {
		MTRBaseClusterDishwasherAlarmClass = _MTRBaseClusterDishwasherAlarmClass{objc.GetClass("MTRBaseClusterDishwasherAlarm")}
	})
	return MTRBaseClusterDishwasherAlarmClass
}

type _MTRBaseClusterDishwasherAlarmClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterDishwasherAlarm] class.
type IMTRBaseClusterDishwasherAlarm interface {
	IMTRGenericBaseCluster
	ModifyEnabledAlarmsWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeLatchWithCompletion(completion unsafe.Pointer)
	ReadAttributeMaskWithCompletion(completion unsafe.Pointer)
	ReadAttributeStateWithCompletion(completion unsafe.Pointer)
	ReadAttributeSupportedWithCompletion(completion unsafe.Pointer)
	ResetWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeLatchWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMaskWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeStateWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSupportedWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Dishwasher Alarm
//
// Attributes and commands for configuring the Dishwasher alarm.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm
type MTRBaseClusterDishwasherAlarm struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterDishwasherAlarmFrom constructs a [MTRBaseClusterDishwasherAlarm] from an unsafe.Pointer.
//
// Cluster Dishwasher Alarm
func MTRBaseClusterDishwasherAlarmFrom(ptr unsafe.Pointer) MTRBaseClusterDishwasherAlarm {
	return MTRBaseClusterDishwasherAlarm{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterDishwasherAlarmClass) Alloc() MTRBaseClusterDishwasherAlarm {
	rv := objc.Send[MTRBaseClusterDishwasherAlarm](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterDishwasherAlarmClass) New() MTRBaseClusterDishwasherAlarm {
	rv := objc.Send[MTRBaseClusterDishwasherAlarm](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterDishwasherAlarm) Init() MTRBaseClusterDishwasherAlarm {
	rv := objc.Send[MTRBaseClusterDishwasherAlarm](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterDishwasherAlarm) Autorelease() MTRBaseClusterDishwasherAlarm {
	rv := objc.Send[MTRBaseClusterDishwasherAlarm](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterDishwasherAlarm creates a new MTRBaseClusterDishwasherAlarm instance.
func NewMTRBaseClusterDishwasherAlarm() MTRBaseClusterDishwasherAlarm {
	return getMTRBaseClusterDishwasherAlarmClass().New()
}




// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/init(device:endpointID:queue:)
func NewMTRBaseClusterDishwasherAlarmWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID foundation.Number, queue unsafe.Pointer) MTRBaseClusterDishwasherAlarm {
	instance := getMTRBaseClusterDishwasherAlarmClass().Alloc()
	rv := objc.Send[MTRBaseClusterDishwasherAlarm](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDishwasherAlarmClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDishwasherAlarmClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDishwasherAlarmClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDishwasherAlarmClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDishwasherAlarmClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeLatch(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDishwasherAlarmClass) ReadAttributeLatchWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLatchWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeMask(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDishwasherAlarmClass) ReadAttributeMaskWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaskWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeState(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDishwasherAlarmClass) ReadAttributeStateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeStateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeSupported(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDishwasherAlarmClass) ReadAttributeSupportedWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

// Command ModifyEnabledAlarms
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/modifyEnabledAlarms(with:completion:)
func (m_ MTRBaseClusterDishwasherAlarm) ModifyEnabledAlarmsWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("modifyEnabledAlarmsWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterDishwasherAlarm) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterDishwasherAlarm) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterDishwasherAlarm) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterDishwasherAlarm) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterDishwasherAlarm) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeLatch(completion:)
func (m_ MTRBaseClusterDishwasherAlarm) ReadAttributeLatchWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeLatchWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeMask(completion:)
func (m_ MTRBaseClusterDishwasherAlarm) ReadAttributeMaskWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMaskWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeState(completion:)
func (m_ MTRBaseClusterDishwasherAlarm) ReadAttributeStateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeStateWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/readAttributeSupported(completion:)
func (m_ MTRBaseClusterDishwasherAlarm) ReadAttributeSupportedWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupportedWithCompletion:"), completion)
}

// Command Reset
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/reset(with:completion:)
func (m_ MTRBaseClusterDishwasherAlarm) ResetWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("resetWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDishwasherAlarm) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDishwasherAlarm) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDishwasherAlarm) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDishwasherAlarm) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDishwasherAlarm) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/subscribeAttributeLatch(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDishwasherAlarm) SubscribeAttributeLatchWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeLatchWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/subscribeAttributeMask(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDishwasherAlarm) SubscribeAttributeMaskWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMaskWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/subscribeAttributeState(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDishwasherAlarm) SubscribeAttributeStateWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeStateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDishwasherAlarm/subscribeAttributeSupported(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDishwasherAlarm) SubscribeAttributeSupportedWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupportedWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


