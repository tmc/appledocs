// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterThreadBorderRouterManagement] class.
var (
	MTRBaseClusterThreadBorderRouterManagementClass     _MTRBaseClusterThreadBorderRouterManagementClass
	MTRBaseClusterThreadBorderRouterManagementClassOnce sync.Once
)

func getMTRBaseClusterThreadBorderRouterManagementClass() _MTRBaseClusterThreadBorderRouterManagementClass {
	MTRBaseClusterThreadBorderRouterManagementClassOnce.Do(func() {
		MTRBaseClusterThreadBorderRouterManagementClass = _MTRBaseClusterThreadBorderRouterManagementClass{objc.GetClass("MTRBaseClusterThreadBorderRouterManagement")}
	})
	return MTRBaseClusterThreadBorderRouterManagementClass
}

type _MTRBaseClusterThreadBorderRouterManagementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterThreadBorderRouterManagement] class.
type IMTRBaseClusterThreadBorderRouterManagement interface {
	IMTRGenericBaseCluster
	GetActiveDatasetRequestWithCompletion(completion unsafe.Pointer)
	GetActiveDatasetRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	GetPendingDatasetRequestWithCompletion(completion unsafe.Pointer)
	GetPendingDatasetRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeActiveDatasetTimestampWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeBorderAgentIDWithCompletion(completion unsafe.Pointer)
	ReadAttributeBorderRouterNameWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeInterfaceEnabledWithCompletion(completion unsafe.Pointer)
	ReadAttributePendingDatasetTimestampWithCompletion(completion unsafe.Pointer)
	ReadAttributeThreadVersionWithCompletion(completion unsafe.Pointer)
	SetActiveDatasetRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SetPendingDatasetRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeActiveDatasetTimestampWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeBorderAgentIDWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeBorderRouterNameWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeInterfaceEnabledWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePendingDatasetTimestampWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeThreadVersionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Thread Border Router Management
//
// Manage the Thread network of Thread Border Router
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement
type MTRBaseClusterThreadBorderRouterManagement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterThreadBorderRouterManagementFrom constructs a [MTRBaseClusterThreadBorderRouterManagement] from an unsafe.Pointer.
//
// Cluster Thread Border Router Management
func MTRBaseClusterThreadBorderRouterManagementFrom(ptr unsafe.Pointer) MTRBaseClusterThreadBorderRouterManagement {
	return MTRBaseClusterThreadBorderRouterManagement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) Alloc() MTRBaseClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRBaseClusterThreadBorderRouterManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) New() MTRBaseClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRBaseClusterThreadBorderRouterManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterThreadBorderRouterManagement) Init() MTRBaseClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRBaseClusterThreadBorderRouterManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterThreadBorderRouterManagement) Autorelease() MTRBaseClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRBaseClusterThreadBorderRouterManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterThreadBorderRouterManagement creates a new MTRBaseClusterThreadBorderRouterManagement instance.
func NewMTRBaseClusterThreadBorderRouterManagement() MTRBaseClusterThreadBorderRouterManagement {
	return getMTRBaseClusterThreadBorderRouterManagementClass().New()
}


// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/init(device:endpointID:queue:)
func NewMTRBaseClusterThreadBorderRouterManagementWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRBaseClusterThreadBorderRouterManagement {
	instance := getMTRBaseClusterThreadBorderRouterManagementClass().Alloc()
	rv := objc.Send[MTRBaseClusterThreadBorderRouterManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeActiveDatasetTimestamp(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) ReadAttributeActiveDatasetTimestampWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeActiveDatasetTimestampWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeBorderAgentID(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) ReadAttributeBorderAgentIDWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeBorderAgentIDWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeBorderRouterName(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) ReadAttributeBorderRouterNameWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeBorderRouterNameWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeInterfaceEnabled(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) ReadAttributeInterfaceEnabledWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeInterfaceEnabledWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributePendingDatasetTimestamp(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) ReadAttributePendingDatasetTimestampWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePendingDatasetTimestampWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeThreadVersion(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) ReadAttributeThreadVersionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeThreadVersionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/getActiveDatasetRequest(completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) GetActiveDatasetRequestWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getActiveDatasetRequestWithCompletion:"), completion)
}

// Command GetActiveDatasetRequest
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/getActiveDatasetRequest(with:completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) GetActiveDatasetRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getActiveDatasetRequestWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/getPendingDatasetRequest(completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) GetPendingDatasetRequestWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getPendingDatasetRequestWithCompletion:"), completion)
}

// Command GetPendingDatasetRequest
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/getPendingDatasetRequest(with:completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) GetPendingDatasetRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getPendingDatasetRequestWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeActiveDatasetTimestamp(completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) ReadAttributeActiveDatasetTimestampWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeActiveDatasetTimestampWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeBorderAgentID(completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) ReadAttributeBorderAgentIDWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeBorderAgentIDWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeBorderRouterName(completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) ReadAttributeBorderRouterNameWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeBorderRouterNameWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeInterfaceEnabled(completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) ReadAttributeInterfaceEnabledWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeInterfaceEnabledWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributePendingDatasetTimestamp(completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) ReadAttributePendingDatasetTimestampWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePendingDatasetTimestampWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/readAttributeThreadVersion(completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) ReadAttributeThreadVersionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeThreadVersionWithCompletion:"), completion)
}

// Command SetActiveDatasetRequest
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/setActiveDatasetRequestWith(_:completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SetActiveDatasetRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveDatasetRequestWithParams:completion:"), params, completion)
}

// Command SetPendingDatasetRequest
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/setPendingDatasetRequestWith(_:completion:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SetPendingDatasetRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPendingDatasetRequestWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/subscribeAttributeActiveDatasetTimestamp(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SubscribeAttributeActiveDatasetTimestampWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeActiveDatasetTimestampWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/subscribeAttributeBorderAgentID(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SubscribeAttributeBorderAgentIDWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeBorderAgentIDWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/subscribeAttributeBorderRouterName(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SubscribeAttributeBorderRouterNameWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeBorderRouterNameWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/subscribeAttributeInterfaceEnabled(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SubscribeAttributeInterfaceEnabledWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeInterfaceEnabledWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/subscribeAttributePendingDatasetTimestamp(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SubscribeAttributePendingDatasetTimestampWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePendingDatasetTimestampWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/subscribeAttributeThreadVersion(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SubscribeAttributeThreadVersionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeThreadVersionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


