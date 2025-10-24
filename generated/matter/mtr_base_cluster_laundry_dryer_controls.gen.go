// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBaseClusterLaundryDryerControls] class.
var (
	MTRBaseClusterLaundryDryerControlsClass     _MTRBaseClusterLaundryDryerControlsClass
	MTRBaseClusterLaundryDryerControlsClassOnce sync.Once
)

func getMTRBaseClusterLaundryDryerControlsClass() _MTRBaseClusterLaundryDryerControlsClass {
	MTRBaseClusterLaundryDryerControlsClassOnce.Do(func() {
		MTRBaseClusterLaundryDryerControlsClass = _MTRBaseClusterLaundryDryerControlsClass{objc.GetClass("MTRBaseClusterLaundryDryerControls")}
	})
	return MTRBaseClusterLaundryDryerControlsClass
}

type _MTRBaseClusterLaundryDryerControlsClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterLaundryDryerControls] class.
type IMTRBaseClusterLaundryDryerControls interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeSelectedDrynessLevelWithCompletion(completion unsafe.Pointer)
	ReadAttributeSupportedDrynessLevelsWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSelectedDrynessLevelWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSupportedDrynessLevelsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	WriteAttributeSelectedDrynessLevelWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeSelectedDrynessLevelWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
}

// Cluster Laundry Dryer Controls
//
// This cluster provides a way to access options associated with the operation of a laundry dryer device type.


// Cluster Laundry Dryer Controls
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls
type MTRBaseClusterLaundryDryerControls struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterLaundryDryerControlsFrom constructs a [MTRBaseClusterLaundryDryerControls] from an unsafe.Pointer.
//
// Cluster Laundry Dryer Controls
func MTRBaseClusterLaundryDryerControlsFrom(ptr unsafe.Pointer) MTRBaseClusterLaundryDryerControls {
	return MTRBaseClusterLaundryDryerControls{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterLaundryDryerControlsClass) Alloc() MTRBaseClusterLaundryDryerControls {
	rv := objc.Send[MTRBaseClusterLaundryDryerControls](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterLaundryDryerControlsClass) New() MTRBaseClusterLaundryDryerControls {
	rv := objc.Send[MTRBaseClusterLaundryDryerControls](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterLaundryDryerControls) Init() MTRBaseClusterLaundryDryerControls {
	rv := objc.Send[MTRBaseClusterLaundryDryerControls](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterLaundryDryerControls) Autorelease() MTRBaseClusterLaundryDryerControls {
	rv := objc.Send[MTRBaseClusterLaundryDryerControls](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterLaundryDryerControls creates a new MTRBaseClusterLaundryDryerControls instance.
func NewMTRBaseClusterLaundryDryerControls() MTRBaseClusterLaundryDryerControls {
	return getMTRBaseClusterLaundryDryerControlsClass().New()
}



// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/init(device:endpointID:queue:)
func NewMTRBaseClusterLaundryDryerControlsWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterLaundryDryerControls {
	instance := getMTRBaseClusterLaundryDryerControlsClass().Alloc()
	rv := objc.Send[MTRBaseClusterLaundryDryerControls](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryDryerControlsClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryDryerControlsClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryDryerControlsClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryDryerControlsClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryDryerControlsClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeSelectedDrynessLevel(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryDryerControlsClass) ReadAttributeSelectedDrynessLevelWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSelectedDrynessLevelWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeSupportedDrynessLevels(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLaundryDryerControlsClass) ReadAttributeSupportedDrynessLevelsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedDrynessLevelsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterLaundryDryerControls) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterLaundryDryerControls) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterLaundryDryerControls) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterLaundryDryerControls) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterLaundryDryerControls) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeSelectedDrynessLevel(completion:)
func (m_ MTRBaseClusterLaundryDryerControls) ReadAttributeSelectedDrynessLevelWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSelectedDrynessLevelWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/readAttributeSupportedDrynessLevels(completion:)
func (m_ MTRBaseClusterLaundryDryerControls) ReadAttributeSupportedDrynessLevelsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupportedDrynessLevelsWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryDryerControls) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryDryerControls) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryDryerControls) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryDryerControls) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryDryerControls) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/subscribeAttributeSelectedDrynessLevel(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryDryerControls) SubscribeAttributeSelectedDrynessLevelWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSelectedDrynessLevelWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/subscribeAttributeSupportedDrynessLevels(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLaundryDryerControls) SubscribeAttributeSupportedDrynessLevelsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupportedDrynessLevelsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/writeAttributeSelectedDrynessLevel(withValue:completion:)
func (m_ MTRBaseClusterLaundryDryerControls) WriteAttributeSelectedDrynessLevelWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSelectedDrynessLevelWithValue:completion:"), value, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLaundryDryerControls/writeAttributeSelectedDrynessLevel(withValue:params:completion:)
func (m_ MTRBaseClusterLaundryDryerControls) WriteAttributeSelectedDrynessLevelWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSelectedDrynessLevelWithValue:params:completion:"), value, params, completion)
}


