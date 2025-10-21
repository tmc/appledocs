// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterOvenCavityOperationalState] class.
var (
	MTRBaseClusterOvenCavityOperationalStateClass     _MTRBaseClusterOvenCavityOperationalStateClass
	MTRBaseClusterOvenCavityOperationalStateClassOnce sync.Once
)

func getMTRBaseClusterOvenCavityOperationalStateClass() _MTRBaseClusterOvenCavityOperationalStateClass {
	MTRBaseClusterOvenCavityOperationalStateClassOnce.Do(func() {
		MTRBaseClusterOvenCavityOperationalStateClass = _MTRBaseClusterOvenCavityOperationalStateClass{objc.GetClass("MTRBaseClusterOvenCavityOperationalState")}
	})
	return MTRBaseClusterOvenCavityOperationalStateClass
}

type _MTRBaseClusterOvenCavityOperationalStateClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterOvenCavityOperationalState] class.
type IMTRBaseClusterOvenCavityOperationalState interface {
	IMTRGenericBaseCluster
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeCountdownTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentPhaseWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeOperationalErrorWithCompletion(completion unsafe.Pointer)
	ReadAttributeOperationalStateWithCompletion(completion unsafe.Pointer)
	ReadAttributeOperationalStateListWithCompletion(completion unsafe.Pointer)
	ReadAttributePhaseListWithCompletion(completion unsafe.Pointer)
	StartWithCompletion(completion unsafe.Pointer)
	StartWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	StopWithCompletion(completion unsafe.Pointer)
	StopWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCountdownTimeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentPhaseWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOperationalErrorWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOperationalStateWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOperationalStateListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePhaseListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Oven Cavity Operational State
//
// This cluster supports remotely monitoring and, where supported, changing the operational state of an Oven.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState
type MTRBaseClusterOvenCavityOperationalState struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterOvenCavityOperationalStateFrom constructs a [MTRBaseClusterOvenCavityOperationalState] from an unsafe.Pointer.
//
// Cluster Oven Cavity Operational State
func MTRBaseClusterOvenCavityOperationalStateFrom(ptr unsafe.Pointer) MTRBaseClusterOvenCavityOperationalState {
	return MTRBaseClusterOvenCavityOperationalState{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) Alloc() MTRBaseClusterOvenCavityOperationalState {
	rv := objc.Send[MTRBaseClusterOvenCavityOperationalState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) New() MTRBaseClusterOvenCavityOperationalState {
	rv := objc.Send[MTRBaseClusterOvenCavityOperationalState](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterOvenCavityOperationalState) Init() MTRBaseClusterOvenCavityOperationalState {
	rv := objc.Send[MTRBaseClusterOvenCavityOperationalState](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterOvenCavityOperationalState) Autorelease() MTRBaseClusterOvenCavityOperationalState {
	rv := objc.Send[MTRBaseClusterOvenCavityOperationalState](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterOvenCavityOperationalState creates a new MTRBaseClusterOvenCavityOperationalState instance.
func NewMTRBaseClusterOvenCavityOperationalState() MTRBaseClusterOvenCavityOperationalState {
	return getMTRBaseClusterOvenCavityOperationalStateClass().New()
}




// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/init(device:endpointID:queue:)
func NewMTRBaseClusterOvenCavityOperationalStateWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRBaseClusterOvenCavityOperationalState {
	instance := getMTRBaseClusterOvenCavityOperationalStateClass().Alloc()
	rv := objc.Send[MTRBaseClusterOvenCavityOperationalState](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeCountdownTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) ReadAttributeCountdownTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCountdownTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeCurrentPhase(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) ReadAttributeCurrentPhaseWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentPhaseWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeOperationalError(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) ReadAttributeOperationalErrorWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOperationalErrorWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeOperationalState(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) ReadAttributeOperationalStateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOperationalStateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeOperationalStateList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) ReadAttributeOperationalStateListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOperationalStateListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributePhaseList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOvenCavityOperationalStateClass) ReadAttributePhaseListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePhaseListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeCountdownTime(completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) ReadAttributeCountdownTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCountdownTimeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeCurrentPhase(completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) ReadAttributeCurrentPhaseWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentPhaseWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeOperationalError(completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) ReadAttributeOperationalErrorWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOperationalErrorWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeOperationalState(completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) ReadAttributeOperationalStateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOperationalStateWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributeOperationalStateList(completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) ReadAttributeOperationalStateListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOperationalStateListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/readAttributePhaseList(completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) ReadAttributePhaseListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePhaseListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/start(completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) StartWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startWithCompletion:"), completion)
}

// Command Start
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/start(with:completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) StartWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/stop(completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) StopWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopWithCompletion:"), completion)
}

// Command Stop
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/stop(with:completion:)
func (m_ MTRBaseClusterOvenCavityOperationalState) StopWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOvenCavityOperationalState) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOvenCavityOperationalState) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOvenCavityOperationalState) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/subscribeAttributeCountdownTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOvenCavityOperationalState) SubscribeAttributeCountdownTimeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCountdownTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/subscribeAttributeCurrentPhase(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOvenCavityOperationalState) SubscribeAttributeCurrentPhaseWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentPhaseWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOvenCavityOperationalState) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOvenCavityOperationalState) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/subscribeAttributeOperationalError(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOvenCavityOperationalState) SubscribeAttributeOperationalErrorWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOperationalErrorWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/subscribeAttributeOperationalState(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOvenCavityOperationalState) SubscribeAttributeOperationalStateWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOperationalStateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/subscribeAttributeOperationalStateList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOvenCavityOperationalState) SubscribeAttributeOperationalStateListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOperationalStateListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOvenCavityOperationalState/subscribeAttributePhaseList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOvenCavityOperationalState) SubscribeAttributePhaseListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePhaseListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


