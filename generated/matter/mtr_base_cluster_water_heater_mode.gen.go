// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBaseClusterWaterHeaterMode] class.
var (
	MTRBaseClusterWaterHeaterModeClass     _MTRBaseClusterWaterHeaterModeClass
	MTRBaseClusterWaterHeaterModeClassOnce sync.Once
)

func getMTRBaseClusterWaterHeaterModeClass() _MTRBaseClusterWaterHeaterModeClass {
	MTRBaseClusterWaterHeaterModeClassOnce.Do(func() {
		MTRBaseClusterWaterHeaterModeClass = _MTRBaseClusterWaterHeaterModeClass{objc.GetClass("MTRBaseClusterWaterHeaterMode")}
	})
	return MTRBaseClusterWaterHeaterModeClass
}

type _MTRBaseClusterWaterHeaterModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterWaterHeaterMode] class.
type IMTRBaseClusterWaterHeaterMode interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
	ChangeToModeWithParamsCompletion(params IMTRWaterHeaterModeClusterChangeToModeParams, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentModeWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeSupportedModesWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSupportedModesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Water Heater Mode
//
// Attributes and commands for selecting a mode from a list of supported options.


// Cluster Water Heater Mode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode
type MTRBaseClusterWaterHeaterMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterWaterHeaterModeFrom constructs a [MTRBaseClusterWaterHeaterMode] from an unsafe.Pointer.
//
// Cluster Water Heater Mode
func MTRBaseClusterWaterHeaterModeFrom(ptr unsafe.Pointer) MTRBaseClusterWaterHeaterMode {
	return MTRBaseClusterWaterHeaterMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterWaterHeaterModeClass) Alloc() MTRBaseClusterWaterHeaterMode {
	rv := objc.Send[MTRBaseClusterWaterHeaterMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterWaterHeaterModeClass) New() MTRBaseClusterWaterHeaterMode {
	rv := objc.Send[MTRBaseClusterWaterHeaterMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterWaterHeaterMode) Init() MTRBaseClusterWaterHeaterMode {
	rv := objc.Send[MTRBaseClusterWaterHeaterMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterWaterHeaterMode) Autorelease() MTRBaseClusterWaterHeaterMode {
	rv := objc.Send[MTRBaseClusterWaterHeaterMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterWaterHeaterMode creates a new MTRBaseClusterWaterHeaterMode instance.
func NewMTRBaseClusterWaterHeaterMode() MTRBaseClusterWaterHeaterMode {
	return getMTRBaseClusterWaterHeaterModeClass().New()
}



// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/init(device:endpointID:queue:)
func NewMTRBaseClusterWaterHeaterModeWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterWaterHeaterMode {
	instance := getMTRBaseClusterWaterHeaterModeClass().Alloc()
	rv := objc.Send[MTRBaseClusterWaterHeaterMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterModeClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterModeClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterModeClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeCurrentMode(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterModeClass) ReadAttributeCurrentModeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentModeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterModeClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterModeClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeSupportedModes(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWaterHeaterModeClass) ReadAttributeSupportedModesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedModesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// Command ChangeToMode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/changeToMode(with:completion:)
func (m_ MTRBaseClusterWaterHeaterMode) ChangeToModeWithParamsCompletion(params IMTRWaterHeaterModeClusterChangeToModeParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:completion:"), params, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterWaterHeaterMode) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterWaterHeaterMode) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterWaterHeaterMode) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeCurrentMode(completion:)
func (m_ MTRBaseClusterWaterHeaterMode) ReadAttributeCurrentModeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentModeWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterWaterHeaterMode) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterWaterHeaterMode) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/readAttributeSupportedModes(completion:)
func (m_ MTRBaseClusterWaterHeaterMode) ReadAttributeSupportedModesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupportedModesWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterMode) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterMode) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterMode) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/subscribeAttributeCurrentMode(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterMode) SubscribeAttributeCurrentModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentModeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterMode) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterMode) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWaterHeaterMode/subscribeAttributeSupportedModes(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWaterHeaterMode) SubscribeAttributeSupportedModesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupportedModesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


