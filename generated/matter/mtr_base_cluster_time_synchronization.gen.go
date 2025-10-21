// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterTimeSynchronization] class.
var (
	MTRBaseClusterTimeSynchronizationClass     _MTRBaseClusterTimeSynchronizationClass
	MTRBaseClusterTimeSynchronizationClassOnce sync.Once
)

func getMTRBaseClusterTimeSynchronizationClass() _MTRBaseClusterTimeSynchronizationClass {
	MTRBaseClusterTimeSynchronizationClassOnce.Do(func() {
		MTRBaseClusterTimeSynchronizationClass = _MTRBaseClusterTimeSynchronizationClass{objc.GetClass("MTRBaseClusterTimeSynchronization")}
	})
	return MTRBaseClusterTimeSynchronizationClass
}

type _MTRBaseClusterTimeSynchronizationClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterTimeSynchronization] class.
type IMTRBaseClusterTimeSynchronization interface {
	IMTRGenericBaseCluster
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeDSTOffsetWithCompletion(completion unsafe.Pointer)
	ReadAttributeDSTOffsetListMaxSizeWithCompletion(completion unsafe.Pointer)
	ReadAttributeDefaultNTPWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeGranularityWithCompletion(completion unsafe.Pointer)
	ReadAttributeLocalTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeNTPServerAvailableWithCompletion(completion unsafe.Pointer)
	ReadAttributeSupportsDNSResolveWithCompletion(completion unsafe.Pointer)
	ReadAttributeTimeSourceWithCompletion(completion unsafe.Pointer)
	ReadAttributeTimeZoneWithCompletion(completion unsafe.Pointer)
	ReadAttributeTimeZoneDatabaseWithCompletion(completion unsafe.Pointer)
	ReadAttributeTimeZoneListMaxSizeWithCompletion(completion unsafe.Pointer)
	ReadAttributeTrustedTimeSourceWithCompletion(completion unsafe.Pointer)
	ReadAttributeUTCTimeWithCompletion(completion unsafe.Pointer)
	SetDSTOffsetWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SetDefaultNTPWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SetTimeZoneWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SetTrustedTimeSourceWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SetUTCTimeWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeDSTOffsetWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeDSTOffsetListMaxSizeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeDefaultNTPWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGranularityWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeLocalTimeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNTPServerAvailableWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSupportsDNSResolveWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeTimeSourceWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeTimeZoneWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeTimeZoneDatabaseWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeTimeZoneListMaxSizeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeTrustedTimeSourceWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeUTCTimeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Time Synchronization
//
// Accurate time is required for a number of reasons, including scheduling, display and validating security materials.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization
type MTRBaseClusterTimeSynchronization struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterTimeSynchronizationFrom constructs a [MTRBaseClusterTimeSynchronization] from an unsafe.Pointer.
//
// Cluster Time Synchronization
func MTRBaseClusterTimeSynchronizationFrom(ptr unsafe.Pointer) MTRBaseClusterTimeSynchronization {
	return MTRBaseClusterTimeSynchronization{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterTimeSynchronizationClass) Alloc() MTRBaseClusterTimeSynchronization {
	rv := objc.Send[MTRBaseClusterTimeSynchronization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterTimeSynchronizationClass) New() MTRBaseClusterTimeSynchronization {
	rv := objc.Send[MTRBaseClusterTimeSynchronization](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterTimeSynchronization) Init() MTRBaseClusterTimeSynchronization {
	rv := objc.Send[MTRBaseClusterTimeSynchronization](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterTimeSynchronization) Autorelease() MTRBaseClusterTimeSynchronization {
	rv := objc.Send[MTRBaseClusterTimeSynchronization](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterTimeSynchronization creates a new MTRBaseClusterTimeSynchronization instance.
func NewMTRBaseClusterTimeSynchronization() MTRBaseClusterTimeSynchronization {
	return getMTRBaseClusterTimeSynchronizationClass().New()
}


// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/init(device:endpointID:queue:)
func NewMTRBaseClusterTimeSynchronizationWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRBaseClusterTimeSynchronization {
	instance := getMTRBaseClusterTimeSynchronizationClass().Alloc()
	rv := objc.Send[MTRBaseClusterTimeSynchronization](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeDSTOffset(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeDSTOffsetWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDSTOffsetWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeDSTOffsetListMaxSize(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeDSTOffsetListMaxSizeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDSTOffsetListMaxSizeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeDefaultNTP(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeDefaultNTPWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDefaultNTPWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeGranularity(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeGranularityWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGranularityWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeLocalTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeLocalTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeLocalTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeNTPServerAvailable(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeNTPServerAvailableWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNTPServerAvailableWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeSupportsDNSResolve(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeSupportsDNSResolveWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportsDNSResolveWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeTimeSource(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeTimeSourceWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTimeSourceWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeTimeZone(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeTimeZoneWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTimeZoneWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeTimeZoneDatabase(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeTimeZoneDatabaseWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTimeZoneDatabaseWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeTimeZoneListMaxSize(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeTimeZoneListMaxSizeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTimeZoneListMaxSizeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeTrustedTimeSource(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeTrustedTimeSourceWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTrustedTimeSourceWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeUTCTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterTimeSynchronizationClass) ReadAttributeUTCTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUTCTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeDSTOffset(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeDSTOffsetWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeDSTOffsetWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeDSTOffsetListMaxSize(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeDSTOffsetListMaxSizeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeDSTOffsetListMaxSizeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeDefaultNTP(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeDefaultNTPWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeDefaultNTPWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeGranularity(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeGranularityWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGranularityWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeLocalTime(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeLocalTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeLocalTimeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeNTPServerAvailable(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeNTPServerAvailableWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNTPServerAvailableWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeSupportsDNSResolve(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeSupportsDNSResolveWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupportsDNSResolveWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeTimeSource(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeTimeSourceWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeTimeSourceWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeTimeZone(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeTimeZoneWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeTimeZoneWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeTimeZoneDatabase(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeTimeZoneDatabaseWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeTimeZoneDatabaseWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeTimeZoneListMaxSize(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeTimeZoneListMaxSizeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeTimeZoneListMaxSizeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeTrustedTimeSource(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeTrustedTimeSourceWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeTrustedTimeSourceWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/readAttributeUTCTime(completion:)
func (m_ MTRBaseClusterTimeSynchronization) ReadAttributeUTCTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeUTCTimeWithCompletion:"), completion)
}

// Command SetDSTOffset
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/setDSTOffsetWith(_:completion:)
func (m_ MTRBaseClusterTimeSynchronization) SetDSTOffsetWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDSTOffsetWithParams:completion:"), params, completion)
}

// Command SetDefaultNTP
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/setDefaultNTPWith(_:completion:)
func (m_ MTRBaseClusterTimeSynchronization) SetDefaultNTPWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultNTPWithParams:completion:"), params, completion)
}

// Command SetTimeZone
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/setTimeZoneWith(_:completion:)
func (m_ MTRBaseClusterTimeSynchronization) SetTimeZoneWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeZoneWithParams:completion:"), params, completion)
}

// Command SetTrustedTimeSource
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/setTrustedTimeSourceWith(_:completion:)
func (m_ MTRBaseClusterTimeSynchronization) SetTrustedTimeSourceWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrustedTimeSourceWithParams:completion:"), params, completion)
}

// Command SetUTCTime
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/setUTCTimeWith(_:completion:)
func (m_ MTRBaseClusterTimeSynchronization) SetUTCTimeWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUTCTimeWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeDSTOffset(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeDSTOffsetWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeDSTOffsetWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeDSTOffsetListMaxSize(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeDSTOffsetListMaxSizeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeDSTOffsetListMaxSizeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeDefaultNTP(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeDefaultNTPWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeDefaultNTPWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeGranularity(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeGranularityWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGranularityWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeLocalTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeLocalTimeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeLocalTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeNTPServerAvailable(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeNTPServerAvailableWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNTPServerAvailableWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeSupportsDNSResolve(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeSupportsDNSResolveWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupportsDNSResolveWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeTimeSource(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeTimeSourceWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeTimeSourceWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeTimeZone(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeTimeZoneWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeTimeZoneWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeTimeZoneDatabase(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeTimeZoneDatabaseWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeTimeZoneDatabaseWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeTimeZoneListMaxSize(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeTimeZoneListMaxSizeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeTimeZoneListMaxSizeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeTrustedTimeSource(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeTrustedTimeSourceWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeTrustedTimeSourceWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeUTCTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeUTCTimeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUTCTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


