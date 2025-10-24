// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBaseClusterPowerTopology] class.
var (
	MTRBaseClusterPowerTopologyClass     _MTRBaseClusterPowerTopologyClass
	MTRBaseClusterPowerTopologyClassOnce sync.Once
)

func getMTRBaseClusterPowerTopologyClass() _MTRBaseClusterPowerTopologyClass {
	MTRBaseClusterPowerTopologyClassOnce.Do(func() {
		MTRBaseClusterPowerTopologyClass = _MTRBaseClusterPowerTopologyClass{objc.GetClass("MTRBaseClusterPowerTopology")}
	})
	return MTRBaseClusterPowerTopologyClass
}

type _MTRBaseClusterPowerTopologyClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterPowerTopology] class.
type IMTRBaseClusterPowerTopology interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeActiveEndpointsWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAvailableEndpointsWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeActiveEndpointsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAvailableEndpointsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Power Topology
//
// The Power Topology Cluster provides a mechanism for expressing how power is flowing between endpoints.


// Cluster Power Topology
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology
type MTRBaseClusterPowerTopology struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterPowerTopologyFrom constructs a [MTRBaseClusterPowerTopology] from an unsafe.Pointer.
//
// Cluster Power Topology
func MTRBaseClusterPowerTopologyFrom(ptr unsafe.Pointer) MTRBaseClusterPowerTopology {
	return MTRBaseClusterPowerTopology{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterPowerTopologyClass) Alloc() MTRBaseClusterPowerTopology {
	rv := objc.Send[MTRBaseClusterPowerTopology](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterPowerTopologyClass) New() MTRBaseClusterPowerTopology {
	rv := objc.Send[MTRBaseClusterPowerTopology](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterPowerTopology) Init() MTRBaseClusterPowerTopology {
	rv := objc.Send[MTRBaseClusterPowerTopology](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterPowerTopology) Autorelease() MTRBaseClusterPowerTopology {
	rv := objc.Send[MTRBaseClusterPowerTopology](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterPowerTopology creates a new MTRBaseClusterPowerTopology instance.
func NewMTRBaseClusterPowerTopology() MTRBaseClusterPowerTopology {
	return getMTRBaseClusterPowerTopologyClass().New()
}



// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/init(device:endpointID:queue:)
func NewMTRBaseClusterPowerTopologyWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterPowerTopology {
	instance := getMTRBaseClusterPowerTopologyClass().Alloc()
	rv := objc.Send[MTRBaseClusterPowerTopology](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterPowerTopologyClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeActiveEndpoints(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterPowerTopologyClass) ReadAttributeActiveEndpointsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeActiveEndpointsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterPowerTopologyClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeAvailableEndpoints(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterPowerTopologyClass) ReadAttributeAvailableEndpointsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAvailableEndpointsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterPowerTopologyClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterPowerTopologyClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterPowerTopologyClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterPowerTopology) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeActiveEndpoints(completion:)
func (m_ MTRBaseClusterPowerTopology) ReadAttributeActiveEndpointsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeActiveEndpointsWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterPowerTopology) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeAvailableEndpoints(completion:)
func (m_ MTRBaseClusterPowerTopology) ReadAttributeAvailableEndpointsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAvailableEndpointsWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterPowerTopology) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterPowerTopology) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterPowerTopology) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterPowerTopology) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/subscribeAttributeActiveEndpoints(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterPowerTopology) SubscribeAttributeActiveEndpointsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeActiveEndpointsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterPowerTopology) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/subscribeAttributeAvailableEndpoints(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterPowerTopology) SubscribeAttributeAvailableEndpointsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAvailableEndpointsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterPowerTopology) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterPowerTopology) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerTopology/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterPowerTopology) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


