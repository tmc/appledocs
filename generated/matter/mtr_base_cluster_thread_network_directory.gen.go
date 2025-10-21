// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterThreadNetworkDirectory] class.
var (
	MTRBaseClusterThreadNetworkDirectoryClass     _MTRBaseClusterThreadNetworkDirectoryClass
	MTRBaseClusterThreadNetworkDirectoryClassOnce sync.Once
)

func getMTRBaseClusterThreadNetworkDirectoryClass() _MTRBaseClusterThreadNetworkDirectoryClass {
	MTRBaseClusterThreadNetworkDirectoryClassOnce.Do(func() {
		MTRBaseClusterThreadNetworkDirectoryClass = _MTRBaseClusterThreadNetworkDirectoryClass{objc.GetClass("MTRBaseClusterThreadNetworkDirectory")}
	})
	return MTRBaseClusterThreadNetworkDirectoryClass
}

type _MTRBaseClusterThreadNetworkDirectoryClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterThreadNetworkDirectory] class.
type IMTRBaseClusterThreadNetworkDirectory interface {
	IMTRGenericBaseCluster
	AddNetworkWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	GetOperationalDatasetWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributePreferredExtendedPanIDWithCompletion(completion unsafe.Pointer)
	ReadAttributeThreadNetworkTableSizeWithCompletion(completion unsafe.Pointer)
	ReadAttributeThreadNetworksWithCompletion(completion unsafe.Pointer)
	RemoveNetworkWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePreferredExtendedPanIDWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeThreadNetworkTableSizeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeThreadNetworksWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	WriteAttributePreferredExtendedPanIDWithValueCompletion(value unsafe.Pointer, completion unsafe.Pointer)
	WriteAttributePreferredExtendedPanIDWithValueParamsCompletion(value unsafe.Pointer, params unsafe.Pointer, completion unsafe.Pointer)
}

// Cluster Thread Network Directory
//
// Manages the names and credentials of Thread networks visible to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory
type MTRBaseClusterThreadNetworkDirectory struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterThreadNetworkDirectoryFrom constructs a [MTRBaseClusterThreadNetworkDirectory] from an unsafe.Pointer.
//
// Cluster Thread Network Directory
func MTRBaseClusterThreadNetworkDirectoryFrom(ptr unsafe.Pointer) MTRBaseClusterThreadNetworkDirectory {
	return MTRBaseClusterThreadNetworkDirectory{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterThreadNetworkDirectoryClass) Alloc() MTRBaseClusterThreadNetworkDirectory {
	rv := objc.Send[MTRBaseClusterThreadNetworkDirectory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterThreadNetworkDirectoryClass) New() MTRBaseClusterThreadNetworkDirectory {
	rv := objc.Send[MTRBaseClusterThreadNetworkDirectory](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterThreadNetworkDirectory) Init() MTRBaseClusterThreadNetworkDirectory {
	rv := objc.Send[MTRBaseClusterThreadNetworkDirectory](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterThreadNetworkDirectory) Autorelease() MTRBaseClusterThreadNetworkDirectory {
	rv := objc.Send[MTRBaseClusterThreadNetworkDirectory](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterThreadNetworkDirectory creates a new MTRBaseClusterThreadNetworkDirectory instance.
func NewMTRBaseClusterThreadNetworkDirectory() MTRBaseClusterThreadNetworkDirectory {
	return getMTRBaseClusterThreadNetworkDirectoryClass().New()
}


// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/init(device:endpointID:queue:)
func NewMTRBaseClusterThreadNetworkDirectoryWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRBaseClusterThreadNetworkDirectory {
	instance := getMTRBaseClusterThreadNetworkDirectoryClass().Alloc()
	rv := objc.Send[MTRBaseClusterThreadNetworkDirectory](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadNetworkDirectoryClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadNetworkDirectoryClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadNetworkDirectoryClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadNetworkDirectoryClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadNetworkDirectoryClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributePreferredExtendedPanID(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadNetworkDirectoryClass) ReadAttributePreferredExtendedPanIDWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePreferredExtendedPanIDWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeThreadNetworkTableSize(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadNetworkDirectoryClass) ReadAttributeThreadNetworkTableSizeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeThreadNetworkTableSizeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeThreadNetworks(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadNetworkDirectoryClass) ReadAttributeThreadNetworksWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeThreadNetworksWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

// Command AddNetwork
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/addNetwork(with:completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) AddNetworkWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addNetworkWithParams:completion:"), params, completion)
}

// Command GetOperationalDataset
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/getOperationalDataset(with:completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) GetOperationalDatasetWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getOperationalDatasetWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributePreferredExtendedPanID(completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) ReadAttributePreferredExtendedPanIDWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePreferredExtendedPanIDWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeThreadNetworkTableSize(completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) ReadAttributeThreadNetworkTableSizeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeThreadNetworkTableSizeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/readAttributeThreadNetworks(completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) ReadAttributeThreadNetworksWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeThreadNetworksWithCompletion:"), completion)
}

// Command RemoveNetwork
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/removeNetwork(with:completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) RemoveNetworkWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeNetworkWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadNetworkDirectory) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadNetworkDirectory) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadNetworkDirectory) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadNetworkDirectory) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadNetworkDirectory) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/subscribeAttributePreferredExtendedPanID(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadNetworkDirectory) SubscribeAttributePreferredExtendedPanIDWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePreferredExtendedPanIDWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/subscribeAttributeThreadNetworkTableSize(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadNetworkDirectory) SubscribeAttributeThreadNetworkTableSizeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeThreadNetworkTableSizeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/subscribeAttributeThreadNetworks(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadNetworkDirectory) SubscribeAttributeThreadNetworksWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeThreadNetworksWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/writeAttributePreferredExtendedPanID(withValue:completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) WriteAttributePreferredExtendedPanIDWithValueCompletion(value unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributePreferredExtendedPanIDWithValue:completion:"), value, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadNetworkDirectory/writeAttributePreferredExtendedPanID(withValue:params:completion:)
func (m_ MTRBaseClusterThreadNetworkDirectory) WriteAttributePreferredExtendedPanIDWithValueParamsCompletion(value unsafe.Pointer, params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributePreferredExtendedPanIDWithValue:params:completion:"), value, params, completion)
}


