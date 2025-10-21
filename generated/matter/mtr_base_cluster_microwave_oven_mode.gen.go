// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterMicrowaveOvenMode] class.
var (
	MTRBaseClusterMicrowaveOvenModeClass     _MTRBaseClusterMicrowaveOvenModeClass
	MTRBaseClusterMicrowaveOvenModeClassOnce sync.Once
)

func getMTRBaseClusterMicrowaveOvenModeClass() _MTRBaseClusterMicrowaveOvenModeClass {
	MTRBaseClusterMicrowaveOvenModeClassOnce.Do(func() {
		MTRBaseClusterMicrowaveOvenModeClass = _MTRBaseClusterMicrowaveOvenModeClass{objc.GetClass("MTRBaseClusterMicrowaveOvenMode")}
	})
	return MTRBaseClusterMicrowaveOvenModeClass
}

type _MTRBaseClusterMicrowaveOvenModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterMicrowaveOvenMode] class.
type IMTRBaseClusterMicrowaveOvenMode interface {
	IMTRGenericBaseCluster
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentModeWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeSupportedModesWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentModeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSupportedModesWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Microwave Oven Mode
//
// Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode
type MTRBaseClusterMicrowaveOvenMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterMicrowaveOvenModeFrom constructs a [MTRBaseClusterMicrowaveOvenMode] from an unsafe.Pointer.
//
// Cluster Microwave Oven Mode
func MTRBaseClusterMicrowaveOvenModeFrom(ptr unsafe.Pointer) MTRBaseClusterMicrowaveOvenMode {
	return MTRBaseClusterMicrowaveOvenMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterMicrowaveOvenModeClass) Alloc() MTRBaseClusterMicrowaveOvenMode {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterMicrowaveOvenModeClass) New() MTRBaseClusterMicrowaveOvenMode {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterMicrowaveOvenMode) Init() MTRBaseClusterMicrowaveOvenMode {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterMicrowaveOvenMode) Autorelease() MTRBaseClusterMicrowaveOvenMode {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterMicrowaveOvenMode creates a new MTRBaseClusterMicrowaveOvenMode instance.
func NewMTRBaseClusterMicrowaveOvenMode() MTRBaseClusterMicrowaveOvenMode {
	return getMTRBaseClusterMicrowaveOvenModeClass().New()
}


// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/init(device:endpointID:queue:)
func NewMTRBaseClusterMicrowaveOvenModeWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRBaseClusterMicrowaveOvenMode {
	instance := getMTRBaseClusterMicrowaveOvenModeClass().Alloc()
	rv := objc.Send[MTRBaseClusterMicrowaveOvenMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenModeClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenModeClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenModeClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeCurrentMode(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenModeClass) ReadAttributeCurrentModeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentModeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenModeClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenModeClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeSupportedModes(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenModeClass) ReadAttributeSupportedModesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedModesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterMicrowaveOvenMode) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterMicrowaveOvenMode) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterMicrowaveOvenMode) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeCurrentMode(completion:)
func (m_ MTRBaseClusterMicrowaveOvenMode) ReadAttributeCurrentModeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentModeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterMicrowaveOvenMode) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterMicrowaveOvenMode) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/readAttributeSupportedModes(completion:)
func (m_ MTRBaseClusterMicrowaveOvenMode) ReadAttributeSupportedModesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupportedModesWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenMode) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenMode) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenMode) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/subscribeAttributeCurrentMode(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenMode) SubscribeAttributeCurrentModeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentModeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenMode) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenMode) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenMode/subscribeAttributeSupportedModes(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenMode) SubscribeAttributeSupportedModesWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupportedModesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


