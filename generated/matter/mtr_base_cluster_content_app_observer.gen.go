// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterContentAppObserver] class.
var (
	MTRBaseClusterContentAppObserverClass     _MTRBaseClusterContentAppObserverClass
	MTRBaseClusterContentAppObserverClassOnce sync.Once
)

func getMTRBaseClusterContentAppObserverClass() _MTRBaseClusterContentAppObserverClass {
	MTRBaseClusterContentAppObserverClassOnce.Do(func() {
		MTRBaseClusterContentAppObserverClass = _MTRBaseClusterContentAppObserverClass{objc.GetClass("MTRBaseClusterContentAppObserver")}
	})
	return MTRBaseClusterContentAppObserverClass
}

type _MTRBaseClusterContentAppObserverClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterContentAppObserver] class.
type IMTRBaseClusterContentAppObserver interface {
	IMTRGenericBaseCluster
	ContentAppMessageWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Content App Observer
//
// This cluster provides an interface for sending targeted commands to an Observer of a Content App on a Video Player device such as a Streaming Media Player, Smart TV or Smart Screen. The cluster server for Content App Observer is implemented by an endpoint that communicates with a Content App, such as a Casting Video Client. The cluster client for Content App Observer is implemented by a Content App endpoint. A Content App is informed of the NodeId of an Observer when a binding is set on the Content App. The Content App can then send the ContentAppMessage to the Observer (server cluster), and the Observer responds with a ContentAppMessageResponse.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver
type MTRBaseClusterContentAppObserver struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterContentAppObserverFrom constructs a [MTRBaseClusterContentAppObserver] from an unsafe.Pointer.
//
// Cluster Content App Observer
func MTRBaseClusterContentAppObserverFrom(ptr unsafe.Pointer) MTRBaseClusterContentAppObserver {
	return MTRBaseClusterContentAppObserver{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterContentAppObserverClass) Alloc() MTRBaseClusterContentAppObserver {
	rv := objc.Send[MTRBaseClusterContentAppObserver](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterContentAppObserverClass) New() MTRBaseClusterContentAppObserver {
	rv := objc.Send[MTRBaseClusterContentAppObserver](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterContentAppObserver) Init() MTRBaseClusterContentAppObserver {
	rv := objc.Send[MTRBaseClusterContentAppObserver](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterContentAppObserver) Autorelease() MTRBaseClusterContentAppObserver {
	rv := objc.Send[MTRBaseClusterContentAppObserver](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterContentAppObserver creates a new MTRBaseClusterContentAppObserver instance.
func NewMTRBaseClusterContentAppObserver() MTRBaseClusterContentAppObserver {
	return getMTRBaseClusterContentAppObserverClass().New()
}


// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/init(device:endpointID:queue:)
func NewMTRBaseClusterContentAppObserverWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRBaseClusterContentAppObserver {
	instance := getMTRBaseClusterContentAppObserverClass().Alloc()
	rv := objc.Send[MTRBaseClusterContentAppObserver](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterContentAppObserverClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterContentAppObserverClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterContentAppObserverClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterContentAppObserverClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterContentAppObserverClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

// Command ContentAppMessage
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/contentAppMessage(with:completion:)
func (m_ MTRBaseClusterContentAppObserver) ContentAppMessageWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("contentAppMessageWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterContentAppObserver) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterContentAppObserver) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterContentAppObserver) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterContentAppObserver) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterContentAppObserver) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterContentAppObserver) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterContentAppObserver) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterContentAppObserver) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterContentAppObserver) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentAppObserver/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterContentAppObserver) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


