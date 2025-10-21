// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBaseClusterMicrowaveOvenControl] class.
var (
	MTRBaseClusterMicrowaveOvenControlClass     _MTRBaseClusterMicrowaveOvenControlClass
	MTRBaseClusterMicrowaveOvenControlClassOnce sync.Once
)

func getMTRBaseClusterMicrowaveOvenControlClass() _MTRBaseClusterMicrowaveOvenControlClass {
	MTRBaseClusterMicrowaveOvenControlClassOnce.Do(func() {
		MTRBaseClusterMicrowaveOvenControlClass = _MTRBaseClusterMicrowaveOvenControlClass{objc.GetClass("MTRBaseClusterMicrowaveOvenControl")}
	})
	return MTRBaseClusterMicrowaveOvenControlClass
}

type _MTRBaseClusterMicrowaveOvenControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterMicrowaveOvenControl] class.
type IMTRBaseClusterMicrowaveOvenControl interface {
	IMTRGenericBaseCluster
	AddMoreTimeWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeCookTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeMaxCookTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeMaxPowerWithCompletion(completion unsafe.Pointer)
	ReadAttributeMinPowerWithCompletion(completion unsafe.Pointer)
	ReadAttributePowerSettingWithCompletion(completion unsafe.Pointer)
	ReadAttributePowerStepWithCompletion(completion unsafe.Pointer)
	ReadAttributeWattRatingWithCompletion(completion unsafe.Pointer)
	SetCookingParametersWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SetCookingParametersWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCookTimeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMaxCookTimeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMaxPowerWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMinPowerWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePowerSettingWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePowerStepWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeWattRatingWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Microwave Oven Control
//
// Attributes and commands for configuring the microwave oven control, and reporting cooking stats.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl
type MTRBaseClusterMicrowaveOvenControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterMicrowaveOvenControlFrom constructs a [MTRBaseClusterMicrowaveOvenControl] from an unsafe.Pointer.
//
// Cluster Microwave Oven Control
func MTRBaseClusterMicrowaveOvenControlFrom(ptr unsafe.Pointer) MTRBaseClusterMicrowaveOvenControl {
	return MTRBaseClusterMicrowaveOvenControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterMicrowaveOvenControlClass) Alloc() MTRBaseClusterMicrowaveOvenControl {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterMicrowaveOvenControlClass) New() MTRBaseClusterMicrowaveOvenControl {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterMicrowaveOvenControl) Init() MTRBaseClusterMicrowaveOvenControl {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterMicrowaveOvenControl) Autorelease() MTRBaseClusterMicrowaveOvenControl {
	rv := objc.Send[MTRBaseClusterMicrowaveOvenControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterMicrowaveOvenControl creates a new MTRBaseClusterMicrowaveOvenControl instance.
func NewMTRBaseClusterMicrowaveOvenControl() MTRBaseClusterMicrowaveOvenControl {
	return getMTRBaseClusterMicrowaveOvenControlClass().New()
}




// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/init(device:endpointID:queue:)
func NewMTRBaseClusterMicrowaveOvenControlWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID foundation.Number, queue unsafe.Pointer) MTRBaseClusterMicrowaveOvenControl {
	instance := getMTRBaseClusterMicrowaveOvenControlClass().Alloc()
	rv := objc.Send[MTRBaseClusterMicrowaveOvenControl](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenControlClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenControlClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenControlClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeCookTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenControlClass) ReadAttributeCookTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCookTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenControlClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenControlClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeMaxCookTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenControlClass) ReadAttributeMaxCookTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxCookTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeMaxPower(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenControlClass) ReadAttributeMaxPowerWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxPowerWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeMinPower(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenControlClass) ReadAttributeMinPowerWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinPowerWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributePowerSetting(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenControlClass) ReadAttributePowerSettingWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePowerSettingWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributePowerStep(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenControlClass) ReadAttributePowerStepWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePowerStepWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeWattRating(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterMicrowaveOvenControlClass) ReadAttributeWattRatingWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint foundation.Number, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeWattRatingWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

// Command AddMoreTime
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/addMoreTime(with:completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) AddMoreTimeWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addMoreTimeWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeCookTime(completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) ReadAttributeCookTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCookTimeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeMaxCookTime(completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) ReadAttributeMaxCookTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMaxCookTimeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeMaxPower(completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) ReadAttributeMaxPowerWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMaxPowerWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeMinPower(completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) ReadAttributeMinPowerWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMinPowerWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributePowerSetting(completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) ReadAttributePowerSettingWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePowerSettingWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributePowerStep(completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) ReadAttributePowerStepWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePowerStepWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/readAttributeWattRating(completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) ReadAttributeWattRatingWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeWattRatingWithCompletion:"), completion)
}

// Command SetCookingParameters
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/setCookingParametersWith(_:completion:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SetCookingParametersWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCookingParametersWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/setCookingParametersWithCompletion(_:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SetCookingParametersWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCookingParametersWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/subscribeAttributeCookTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SubscribeAttributeCookTimeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCookTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/subscribeAttributeMaxCookTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SubscribeAttributeMaxCookTimeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMaxCookTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/subscribeAttributeMaxPower(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SubscribeAttributeMaxPowerWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMaxPowerWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/subscribeAttributeMinPower(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SubscribeAttributeMinPowerWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMinPowerWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/subscribeAttributePowerSetting(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SubscribeAttributePowerSettingWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePowerSettingWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/subscribeAttributePowerStep(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SubscribeAttributePowerStepWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePowerStepWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMicrowaveOvenControl/subscribeAttributeWattRating(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterMicrowaveOvenControl) SubscribeAttributeWattRatingWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeWattRatingWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


