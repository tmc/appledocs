// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBaseClusterDeviceEnergyManagementMode] class.
var (
	MTRBaseClusterDeviceEnergyManagementModeClass     _MTRBaseClusterDeviceEnergyManagementModeClass
	MTRBaseClusterDeviceEnergyManagementModeClassOnce sync.Once
)

func getMTRBaseClusterDeviceEnergyManagementModeClass() _MTRBaseClusterDeviceEnergyManagementModeClass {
	MTRBaseClusterDeviceEnergyManagementModeClassOnce.Do(func() {
		MTRBaseClusterDeviceEnergyManagementModeClass = _MTRBaseClusterDeviceEnergyManagementModeClass{objc.GetClass("MTRBaseClusterDeviceEnergyManagementMode")}
	})
	return MTRBaseClusterDeviceEnergyManagementModeClass
}

type _MTRBaseClusterDeviceEnergyManagementModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterDeviceEnergyManagementMode] class.
type IMTRBaseClusterDeviceEnergyManagementMode interface {
	IMTRGenericBaseCluster
	ChangeToModeWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
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

// Cluster Device Energy Management Mode
//
// Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode
type MTRBaseClusterDeviceEnergyManagementMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterDeviceEnergyManagementModeFrom constructs a [MTRBaseClusterDeviceEnergyManagementMode] from an unsafe.Pointer.
//
// Cluster Device Energy Management Mode
func MTRBaseClusterDeviceEnergyManagementModeFrom(ptr unsafe.Pointer) MTRBaseClusterDeviceEnergyManagementMode {
	return MTRBaseClusterDeviceEnergyManagementMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterDeviceEnergyManagementModeClass) Alloc() MTRBaseClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagementMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterDeviceEnergyManagementModeClass) New() MTRBaseClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagementMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterDeviceEnergyManagementMode) Init() MTRBaseClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagementMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterDeviceEnergyManagementMode) Autorelease() MTRBaseClusterDeviceEnergyManagementMode {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagementMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterDeviceEnergyManagementMode creates a new MTRBaseClusterDeviceEnergyManagementMode instance.
func NewMTRBaseClusterDeviceEnergyManagementMode() MTRBaseClusterDeviceEnergyManagementMode {
	return getMTRBaseClusterDeviceEnergyManagementModeClass().New()
}




// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/init(device:endpointID:queue:)
func NewMTRBaseClusterDeviceEnergyManagementModeWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID foundation.Number, queue unsafe.Pointer) MTRBaseClusterDeviceEnergyManagementMode {
	instance := getMTRBaseClusterDeviceEnergyManagementModeClass().Alloc()
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagementMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementModeClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementModeClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementModeClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeCurrentMode(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementModeClass) ReadAttributeCurrentModeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentModeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementModeClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementModeClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeSupportedModes(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementModeClass) ReadAttributeSupportedModesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedModesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

// Command ChangeToMode
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/changeToMode(with:completion:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) ChangeToModeWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeCurrentMode(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) ReadAttributeCurrentModeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentModeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/readAttributeSupportedModes(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) ReadAttributeSupportedModesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupportedModesWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/subscribeAttributeCurrentMode(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) SubscribeAttributeCurrentModeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentModeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagementMode/subscribeAttributeSupportedModes(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagementMode) SubscribeAttributeSupportedModesWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupportedModesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


