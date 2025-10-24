// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBaseClusterLaundryWasherControls] class.
var (
	MTRBaseClusterLaundryWasherControlsClass     _MTRBaseClusterLaundryWasherControlsClass
	MTRBaseClusterLaundryWasherControlsClassOnce sync.Once
)

func getMTRBaseClusterLaundryWasherControlsClass() _MTRBaseClusterLaundryWasherControlsClass {
	MTRBaseClusterLaundryWasherControlsClassOnce.Do(func() {
		MTRBaseClusterLaundryWasherControlsClass = _MTRBaseClusterLaundryWasherControlsClass{objc.GetClass("MTRBaseClusterLaundryWasherControls")}
	})
	return MTRBaseClusterLaundryWasherControlsClass
}

type _MTRBaseClusterLaundryWasherControlsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterLaundryWasherControls] class.
type IMTRBaseClusterLaundryWasherControls interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfRinsesWithCompletion(completion unsafe.Pointer)
	ReadAttributeSpinSpeedCurrentWithCompletion(completion unsafe.Pointer)
	ReadAttributeSpinSpeedsWithCompletion(completion unsafe.Pointer)
	ReadAttributeSupportedRinsesWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfRinsesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSpinSpeedCurrentWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSpinSpeedsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSupportedRinsesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	WriteAttributeNumberOfRinsesWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeNumberOfRinsesWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeSpinSpeedCurrentWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeSpinSpeedCurrentWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
}

// Cluster Laundry Washer Controls
//
// This cluster supports remotely monitoring and controlling the different types of functionality available to a washing device, such as a washing machine.


// Cluster Laundry Washer Controls
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls
type MTRBaseClusterLaundryWasherControls struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterLaundryWasherControlsFrom constructs a [MTRBaseClusterLaundryWasherControls] from an unsafe.Pointer.
//
// Cluster Laundry Washer Controls
func MTRBaseClusterLaundryWasherControlsFrom(ptr unsafe.Pointer) MTRBaseClusterLaundryWasherControls {
	return MTRBaseClusterLaundryWasherControls{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterLaundryWasherControlsClass) Alloc() MTRBaseClusterLaundryWasherControls {
	rv := objc.Send[MTRBaseClusterLaundryWasherControls](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterLaundryWasherControlsClass) New() MTRBaseClusterLaundryWasherControls {
	rv := objc.Send[MTRBaseClusterLaundryWasherControls](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterLaundryWasherControls) Init() MTRBaseClusterLaundryWasherControls {
	rv := objc.Send[MTRBaseClusterLaundryWasherControls](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterLaundryWasherControls) Autorelease() MTRBaseClusterLaundryWasherControls {
	rv := objc.Send[MTRBaseClusterLaundryWasherControls](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterLaundryWasherControls creates a new MTRBaseClusterLaundryWasherControls instance.
func NewMTRBaseClusterLaundryWasherControls() MTRBaseClusterLaundryWasherControls {
	return getMTRBaseClusterLaundryWasherControlsClass().New()
}



// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/init(device:endpointID:queue:)
func NewMTRBaseClusterLaundryWasherControlsWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterLaundryWasherControls {
	instance := getMTRBaseClusterLaundryWasherControlsClass().Alloc()
	rv := objc.Send[MTRBaseClusterLaundryWasherControls](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryWasherControlsClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryWasherControlsClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryWasherControlsClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryWasherControlsClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryWasherControlsClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeNumberOfRinses(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryWasherControlsClass) ReadAttributeNumberOfRinsesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfRinsesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeSpinSpeedCurrent(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryWasherControlsClass) ReadAttributeSpinSpeedCurrentWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSpinSpeedCurrentWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeSpinSpeeds(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryWasherControlsClass) ReadAttributeSpinSpeedsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSpinSpeedsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeSupportedRinses(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryWasherControlsClass) ReadAttributeSupportedRinsesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedRinsesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterLaundryWasherControls) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterLaundryWasherControls) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterLaundryWasherControls) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterLaundryWasherControls) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterLaundryWasherControls) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeNumberOfRinses(completion:)
func (m_ MTRBaseClusterLaundryWasherControls) ReadAttributeNumberOfRinsesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfRinsesWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeSpinSpeedCurrent(completion:)
func (m_ MTRBaseClusterLaundryWasherControls) ReadAttributeSpinSpeedCurrentWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSpinSpeedCurrentWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeSpinSpeeds(completion:)
func (m_ MTRBaseClusterLaundryWasherControls) ReadAttributeSpinSpeedsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSpinSpeedsWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/readAttributeSupportedRinses(completion:)
func (m_ MTRBaseClusterLaundryWasherControls) ReadAttributeSupportedRinsesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupportedRinsesWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryWasherControls) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryWasherControls) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryWasherControls) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryWasherControls) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryWasherControls) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/subscribeAttributeNumberOfRinses(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryWasherControls) SubscribeAttributeNumberOfRinsesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfRinsesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/subscribeAttributeSpinSpeedCurrent(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryWasherControls) SubscribeAttributeSpinSpeedCurrentWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSpinSpeedCurrentWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/subscribeAttributeSpinSpeeds(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryWasherControls) SubscribeAttributeSpinSpeedsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSpinSpeedsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/subscribeAttributeSupportedRinses(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryWasherControls) SubscribeAttributeSupportedRinsesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupportedRinsesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/writeAttributeNumberOfRinses(withValue:completion:)
func (m_ MTRBaseClusterLaundryWasherControls) WriteAttributeNumberOfRinsesWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeNumberOfRinsesWithValue:completion:"), value, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/writeAttributeNumberOfRinses(withValue:params:completion:)
func (m_ MTRBaseClusterLaundryWasherControls) WriteAttributeNumberOfRinsesWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeNumberOfRinsesWithValue:params:completion:"), value, params, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/writeAttributeSpinSpeedCurrent(withValue:completion:)
func (m_ MTRBaseClusterLaundryWasherControls) WriteAttributeSpinSpeedCurrentWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSpinSpeedCurrentWithValue:completion:"), value, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryWasherControls/writeAttributeSpinSpeedCurrent(withValue:params:completion:)
func (m_ MTRBaseClusterLaundryWasherControls) WriteAttributeSpinSpeedCurrentWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSpinSpeedCurrentWithValue:params:completion:"), value, params, completion)
}


