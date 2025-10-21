// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBaseClusterWiFiNetworkManagement] class.
var (
	MTRBaseClusterWiFiNetworkManagementClass     _MTRBaseClusterWiFiNetworkManagementClass
	MTRBaseClusterWiFiNetworkManagementClassOnce sync.Once
)

func getMTRBaseClusterWiFiNetworkManagementClass() _MTRBaseClusterWiFiNetworkManagementClass {
	MTRBaseClusterWiFiNetworkManagementClassOnce.Do(func() {
		MTRBaseClusterWiFiNetworkManagementClass = _MTRBaseClusterWiFiNetworkManagementClass{objc.GetClass("MTRBaseClusterWiFiNetworkManagement")}
	})
	return MTRBaseClusterWiFiNetworkManagementClass
}

type _MTRBaseClusterWiFiNetworkManagementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterWiFiNetworkManagement] class.
type IMTRBaseClusterWiFiNetworkManagement interface {
	IMTRGenericBaseCluster
	NetworkPassphraseRequestWithCompletion(completion unsafe.Pointer)
	NetworkPassphraseRequestWithParamsCompletion(params IMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributePassphraseSurrogateWithCompletion(completion unsafe.Pointer)
	ReadAttributeSSIDWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePassphraseSurrogateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSSIDWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Wi-Fi Network Management
//
// Functionality to retrieve operational information about a managed Wi-Fi network.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement
type MTRBaseClusterWiFiNetworkManagement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterWiFiNetworkManagementFrom constructs a [MTRBaseClusterWiFiNetworkManagement] from an unsafe.Pointer.
//
// Cluster Wi-Fi Network Management
func MTRBaseClusterWiFiNetworkManagementFrom(ptr unsafe.Pointer) MTRBaseClusterWiFiNetworkManagement {
	return MTRBaseClusterWiFiNetworkManagement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterWiFiNetworkManagementClass) Alloc() MTRBaseClusterWiFiNetworkManagement {
	rv := objc.Send[MTRBaseClusterWiFiNetworkManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterWiFiNetworkManagementClass) New() MTRBaseClusterWiFiNetworkManagement {
	rv := objc.Send[MTRBaseClusterWiFiNetworkManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterWiFiNetworkManagement) Init() MTRBaseClusterWiFiNetworkManagement {
	rv := objc.Send[MTRBaseClusterWiFiNetworkManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterWiFiNetworkManagement) Autorelease() MTRBaseClusterWiFiNetworkManagement {
	rv := objc.Send[MTRBaseClusterWiFiNetworkManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterWiFiNetworkManagement creates a new MTRBaseClusterWiFiNetworkManagement instance.
func NewMTRBaseClusterWiFiNetworkManagement() MTRBaseClusterWiFiNetworkManagement {
	return getMTRBaseClusterWiFiNetworkManagementClass().New()
}




// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/init(device:endpointID:queue:)
func NewMTRBaseClusterWiFiNetworkManagementWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID foundation.INumber, queue unsafe.Pointer) MTRBaseClusterWiFiNetworkManagement {
	instance := getMTRBaseClusterWiFiNetworkManagementClass().Alloc()
	rv := objc.Send[MTRBaseClusterWiFiNetworkManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWiFiNetworkManagementClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWiFiNetworkManagementClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWiFiNetworkManagementClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWiFiNetworkManagementClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWiFiNetworkManagementClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributePassphraseSurrogate(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWiFiNetworkManagementClass) ReadAttributePassphraseSurrogateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePassphraseSurrogateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributeSSID(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWiFiNetworkManagementClass) ReadAttributeSSIDWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint foundation.INumber, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSSIDWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/networkPassphraseRequest(completion:)
func (m_ MTRBaseClusterWiFiNetworkManagement) NetworkPassphraseRequestWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("networkPassphraseRequestWithCompletion:"), completion)
}

// Command NetworkPassphraseRequest
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/networkPassphraseRequest(with:completion:)
func (m_ MTRBaseClusterWiFiNetworkManagement) NetworkPassphraseRequestWithParamsCompletion(params IMTRWiFiNetworkManagementClusterNetworkPassphraseRequestParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("networkPassphraseRequestWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterWiFiNetworkManagement) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterWiFiNetworkManagement) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterWiFiNetworkManagement) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterWiFiNetworkManagement) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterWiFiNetworkManagement) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributePassphraseSurrogate(completion:)
func (m_ MTRBaseClusterWiFiNetworkManagement) ReadAttributePassphraseSurrogateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePassphraseSurrogateWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/readAttributeSSID(completion:)
func (m_ MTRBaseClusterWiFiNetworkManagement) ReadAttributeSSIDWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSSIDWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWiFiNetworkManagement) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWiFiNetworkManagement) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWiFiNetworkManagement) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWiFiNetworkManagement) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWiFiNetworkManagement) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/subscribeAttributePassphraseSurrogate(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWiFiNetworkManagement) SubscribeAttributePassphraseSurrogateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePassphraseSurrogateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWiFiNetworkManagement/subscribeAttributeSSID(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWiFiNetworkManagement) SubscribeAttributeSSIDWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSSIDWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


