// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode] class.
var (
	MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass     _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass
	MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClassOnce sync.Once
)

func getMTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass() _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass {
	MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClassOnce.Do(func() {
		MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass = _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass{objc.GetClass("MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode")}
	})
	return MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass
}

type _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode] class.
type IMTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode interface {
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

// Cluster Refrigerator And Temperature Controlled Cabinet Mode
//
// Attributes and commands for selecting a mode from a list of supported options.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode
type MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeFrom constructs a [MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode] from an unsafe.Pointer.
//
// Cluster Refrigerator And Temperature Controlled Cabinet Mode
func MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeFrom(ptr unsafe.Pointer) MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode {
	return MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass) Alloc() MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass) New() MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) Init() MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) Autorelease() MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode {
	rv := objc.Send[MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode creates a new MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode instance.
func NewMTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode() MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode {
	return getMTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass().New()
}




// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/init(device:endpointID:queue:)
func NewMTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode {
	instance := getMTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass().Alloc()
	rv := objc.Send[MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeCurrentMode(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass) ReadAttributeCurrentModeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentModeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeSupportedModes(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetModeClass) ReadAttributeSupportedModesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedModesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

// Command ChangeToMode
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/changeToMode(with:completion:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) ChangeToModeWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("changeToModeWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeCurrentMode(completion:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeCurrentModeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentModeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/readAttributeSupportedModes(completion:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) ReadAttributeSupportedModesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupportedModesWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/subscribeAttributeCurrentMode(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) SubscribeAttributeCurrentModeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentModeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode/subscribeAttributeSupportedModes(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterRefrigeratorAndTemperatureControlledCabinetMode) SubscribeAttributeSupportedModesWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupportedModesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


