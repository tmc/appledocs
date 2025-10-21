// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBaseClusterTemperatureControl] class.
var (
	MTRBaseClusterTemperatureControlClass     _MTRBaseClusterTemperatureControlClass
	MTRBaseClusterTemperatureControlClassOnce sync.Once
)

func getMTRBaseClusterTemperatureControlClass() _MTRBaseClusterTemperatureControlClass {
	MTRBaseClusterTemperatureControlClassOnce.Do(func() {
		MTRBaseClusterTemperatureControlClass = _MTRBaseClusterTemperatureControlClass{objc.GetClass("MTRBaseClusterTemperatureControl")}
	})
	return MTRBaseClusterTemperatureControlClass
}

type _MTRBaseClusterTemperatureControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterTemperatureControl] class.
type IMTRBaseClusterTemperatureControl interface {
	IMTRGenericBaseCluster
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeMaxTemperatureWithCompletion(completion unsafe.Pointer)
	ReadAttributeMinTemperatureWithCompletion(completion unsafe.Pointer)
	ReadAttributeSelectedTemperatureLevelWithCompletion(completion unsafe.Pointer)
	ReadAttributeStepWithCompletion(completion unsafe.Pointer)
	ReadAttributeSupportedTemperatureLevelsWithCompletion(completion unsafe.Pointer)
	ReadAttributeTemperatureSetpointWithCompletion(completion unsafe.Pointer)
	SetTemperatureWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SetTemperatureWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMaxTemperatureWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMinTemperatureWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSelectedTemperatureLevelWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeStepWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSupportedTemperatureLevelsWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeTemperatureSetpointWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Temperature Control
//
// Attributes and commands for configuring the temperature control, and reporting temperature.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl
type MTRBaseClusterTemperatureControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterTemperatureControlFrom constructs a [MTRBaseClusterTemperatureControl] from an unsafe.Pointer.
//
// Cluster Temperature Control
func MTRBaseClusterTemperatureControlFrom(ptr unsafe.Pointer) MTRBaseClusterTemperatureControl {
	return MTRBaseClusterTemperatureControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterTemperatureControlClass) Alloc() MTRBaseClusterTemperatureControl {
	rv := objc.Send[MTRBaseClusterTemperatureControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterTemperatureControlClass) New() MTRBaseClusterTemperatureControl {
	rv := objc.Send[MTRBaseClusterTemperatureControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterTemperatureControl) Init() MTRBaseClusterTemperatureControl {
	rv := objc.Send[MTRBaseClusterTemperatureControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterTemperatureControl) Autorelease() MTRBaseClusterTemperatureControl {
	rv := objc.Send[MTRBaseClusterTemperatureControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterTemperatureControl creates a new MTRBaseClusterTemperatureControl instance.
func NewMTRBaseClusterTemperatureControl() MTRBaseClusterTemperatureControl {
	return getMTRBaseClusterTemperatureControlClass().New()
}




// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/init(device:endpointID:queue:)
func NewMTRBaseClusterTemperatureControlWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID foundation.Number, queue unsafe.Pointer) MTRBaseClusterTemperatureControl {
	instance := getMTRBaseClusterTemperatureControlClass().Alloc()
	rv := objc.Send[MTRBaseClusterTemperatureControl](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTemperatureControlClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTemperatureControlClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTemperatureControlClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTemperatureControlClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTemperatureControlClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeMaxTemperature(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTemperatureControlClass) ReadAttributeMaxTemperatureWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxTemperatureWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeMinTemperature(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTemperatureControlClass) ReadAttributeMinTemperatureWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinTemperatureWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeSelectedTemperatureLevel(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTemperatureControlClass) ReadAttributeSelectedTemperatureLevelWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSelectedTemperatureLevelWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeStep(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTemperatureControlClass) ReadAttributeStepWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeStepWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeSupportedTemperatureLevels(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTemperatureControlClass) ReadAttributeSupportedTemperatureLevelsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedTemperatureLevelsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeTemperatureSetpoint(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTemperatureControlClass) ReadAttributeTemperatureSetpointWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTemperatureSetpointWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterTemperatureControl) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterTemperatureControl) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterTemperatureControl) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterTemperatureControl) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterTemperatureControl) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeMaxTemperature(completion:)
func (m_ MTRBaseClusterTemperatureControl) ReadAttributeMaxTemperatureWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMaxTemperatureWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeMinTemperature(completion:)
func (m_ MTRBaseClusterTemperatureControl) ReadAttributeMinTemperatureWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMinTemperatureWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeSelectedTemperatureLevel(completion:)
func (m_ MTRBaseClusterTemperatureControl) ReadAttributeSelectedTemperatureLevelWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSelectedTemperatureLevelWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeStep(completion:)
func (m_ MTRBaseClusterTemperatureControl) ReadAttributeStepWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeStepWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeSupportedTemperatureLevels(completion:)
func (m_ MTRBaseClusterTemperatureControl) ReadAttributeSupportedTemperatureLevelsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupportedTemperatureLevelsWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/readAttributeTemperatureSetpoint(completion:)
func (m_ MTRBaseClusterTemperatureControl) ReadAttributeTemperatureSetpointWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeTemperatureSetpointWithCompletion:"), completion)
}

// Command SetTemperature
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/setTemperatureWith(_:completion:)
func (m_ MTRBaseClusterTemperatureControl) SetTemperatureWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTemperatureWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/setTemperatureWithCompletion(_:)
func (m_ MTRBaseClusterTemperatureControl) SetTemperatureWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTemperatureWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTemperatureControl) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTemperatureControl) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTemperatureControl) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTemperatureControl) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTemperatureControl) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/subscribeAttributeMaxTemperature(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTemperatureControl) SubscribeAttributeMaxTemperatureWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMaxTemperatureWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/subscribeAttributeMinTemperature(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTemperatureControl) SubscribeAttributeMinTemperatureWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMinTemperatureWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/subscribeAttributeSelectedTemperatureLevel(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTemperatureControl) SubscribeAttributeSelectedTemperatureLevelWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSelectedTemperatureLevelWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/subscribeAttributeStep(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTemperatureControl) SubscribeAttributeStepWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeStepWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/subscribeAttributeSupportedTemperatureLevels(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTemperatureControl) SubscribeAttributeSupportedTemperatureLevelsWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupportedTemperatureLevelsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureControl/subscribeAttributeTemperatureSetpoint(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTemperatureControl) SubscribeAttributeTemperatureSetpointWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeTemperatureSetpointWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


