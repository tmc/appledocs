// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterDeviceEnergyManagement] class.
var (
	MTRBaseClusterDeviceEnergyManagementClass     _MTRBaseClusterDeviceEnergyManagementClass
	MTRBaseClusterDeviceEnergyManagementClassOnce sync.Once
)

func getMTRBaseClusterDeviceEnergyManagementClass() _MTRBaseClusterDeviceEnergyManagementClass {
	MTRBaseClusterDeviceEnergyManagementClassOnce.Do(func() {
		MTRBaseClusterDeviceEnergyManagementClass = _MTRBaseClusterDeviceEnergyManagementClass{objc.GetClass("MTRBaseClusterDeviceEnergyManagement")}
	})
	return MTRBaseClusterDeviceEnergyManagementClass
}

type _MTRBaseClusterDeviceEnergyManagementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterDeviceEnergyManagement] class.
type IMTRBaseClusterDeviceEnergyManagement interface {
	IMTRGenericBaseCluster
	CancelPowerAdjustRequestWithCompletion(completion unsafe.Pointer)
	CancelPowerAdjustRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	CancelRequestWithCompletion(completion unsafe.Pointer)
	CancelRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	ModifyForecastRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	PauseRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	PowerAdjustRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeAbsMaxPowerWithCompletion(completion unsafe.Pointer)
	ReadAttributeAbsMinPowerWithCompletion(completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeESACanGenerateWithCompletion(completion unsafe.Pointer)
	ReadAttributeESAStateWithCompletion(completion unsafe.Pointer)
	ReadAttributeESATypeWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeForecastWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeOptOutStateWithCompletion(completion unsafe.Pointer)
	ReadAttributePowerAdjustmentCapabilityWithCompletion(completion unsafe.Pointer)
	RequestConstraintBasedForecastWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	ResumeRequestWithCompletion(completion unsafe.Pointer)
	ResumeRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	StartTimeAdjustRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer)
	SubscribeAttributeAbsMaxPowerWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAbsMinPowerWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeESACanGenerateWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeESAStateWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeESATypeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeForecastWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOptOutStateWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePowerAdjustmentCapabilityWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
}

// Cluster Device Energy Management
//
// This cluster allows a client to manage the power draw of a device. An example of such a client could be an Energy Management System (EMS) which controls an Energy Smart Appliance (ESA).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement
type MTRBaseClusterDeviceEnergyManagement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterDeviceEnergyManagementFrom constructs a [MTRBaseClusterDeviceEnergyManagement] from an unsafe.Pointer.
//
// Cluster Device Energy Management
func MTRBaseClusterDeviceEnergyManagementFrom(ptr unsafe.Pointer) MTRBaseClusterDeviceEnergyManagement {
	return MTRBaseClusterDeviceEnergyManagement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterDeviceEnergyManagementClass) Alloc() MTRBaseClusterDeviceEnergyManagement {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterDeviceEnergyManagementClass) New() MTRBaseClusterDeviceEnergyManagement {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterDeviceEnergyManagement) Init() MTRBaseClusterDeviceEnergyManagement {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterDeviceEnergyManagement) Autorelease() MTRBaseClusterDeviceEnergyManagement {
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterDeviceEnergyManagement creates a new MTRBaseClusterDeviceEnergyManagement instance.
func NewMTRBaseClusterDeviceEnergyManagement() MTRBaseClusterDeviceEnergyManagement {
	return getMTRBaseClusterDeviceEnergyManagementClass().New()
}




// For all instance methods (reads, writes, commands) that take a completion, the completion will be called on the provided queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/init(device:endpointID:queue:)
func NewMTRBaseClusterDeviceEnergyManagementWithDeviceEndpointIDQueue(device unsafe.Pointer, endpointID unsafe.Pointer, queue unsafe.Pointer) MTRBaseClusterDeviceEnergyManagement {
	instance := getMTRBaseClusterDeviceEnergyManagementClass().Alloc()
	rv := objc.Send[MTRBaseClusterDeviceEnergyManagement](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeAbsMaxPower(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementClass) ReadAttributeAbsMaxPowerWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAbsMaxPowerWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeAbsMinPower(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementClass) ReadAttributeAbsMinPowerWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAbsMinPowerWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeESACanGenerate(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementClass) ReadAttributeESACanGenerateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeESACanGenerateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeESAState(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementClass) ReadAttributeESAStateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeESAStateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeESAType(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementClass) ReadAttributeESATypeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeESATypeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeForecast(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementClass) ReadAttributeForecastWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeForecastWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeOptOutState(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementClass) ReadAttributeOptOutStateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOptOutStateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributePowerAdjustmentCapability(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterDeviceEnergyManagementClass) ReadAttributePowerAdjustmentCapabilityWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer unsafe.Pointer, endpoint unsafe.Pointer, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePowerAdjustmentCapabilityWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/cancelPowerAdjustRequest(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) CancelPowerAdjustRequestWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelPowerAdjustRequestWithCompletion:"), completion)
}

// Command CancelPowerAdjustRequest
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/cancelPowerAdjustRequest(with:completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) CancelPowerAdjustRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelPowerAdjustRequestWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/cancelRequest(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) CancelRequestWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelRequestWithCompletion:"), completion)
}

// Command CancelRequest
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/cancelRequest(with:completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) CancelRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("cancelRequestWithParams:completion:"), params, completion)
}

// Command ModifyForecastRequest
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/modifyForecastRequest(with:completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ModifyForecastRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("modifyForecastRequestWithParams:completion:"), params, completion)
}

// Command PauseRequest
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/pauseRequest(with:completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) PauseRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("pauseRequestWithParams:completion:"), params, completion)
}

// Command PowerAdjustRequest
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/powerAdjustRequest(with:completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) PowerAdjustRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("powerAdjustRequestWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeAbsMaxPower(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ReadAttributeAbsMaxPowerWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAbsMaxPowerWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeAbsMinPower(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ReadAttributeAbsMinPowerWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAbsMinPowerWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeESACanGenerate(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ReadAttributeESACanGenerateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeESACanGenerateWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeESAState(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ReadAttributeESAStateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeESAStateWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeESAType(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ReadAttributeESATypeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeESATypeWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeForecast(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ReadAttributeForecastWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeForecastWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributeOptOutState(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ReadAttributeOptOutStateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOptOutStateWithCompletion:"), completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/readAttributePowerAdjustmentCapability(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ReadAttributePowerAdjustmentCapabilityWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePowerAdjustmentCapabilityWithCompletion:"), completion)
}

// Command RequestConstraintBasedForecast
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/requestConstraintBasedForecast(with:completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) RequestConstraintBasedForecastWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("requestConstraintBasedForecastWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/resumeRequest(completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ResumeRequestWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("resumeRequestWithCompletion:"), completion)
}

// Command ResumeRequest
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/resumeRequest(with:completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) ResumeRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("resumeRequestWithParams:completion:"), params, completion)
}

// Command StartTimeAdjustRequest
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/startTimeAdjustRequest(with:completion:)
func (m_ MTRBaseClusterDeviceEnergyManagement) StartTimeAdjustRequestWithParamsCompletion(params unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("startTimeAdjustRequestWithParams:completion:"), params, completion)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/subscribeAttributeAbsMaxPower(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagement) SubscribeAttributeAbsMaxPowerWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAbsMaxPowerWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/subscribeAttributeAbsMinPower(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagement) SubscribeAttributeAbsMinPowerWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAbsMinPowerWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagement) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagement) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagement) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/subscribeAttributeESACanGenerate(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagement) SubscribeAttributeESACanGenerateWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeESACanGenerateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/subscribeAttributeESAState(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagement) SubscribeAttributeESAStateWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeESAStateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/subscribeAttributeESAType(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagement) SubscribeAttributeESATypeWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeESATypeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagement) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/subscribeAttributeForecast(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagement) SubscribeAttributeForecastWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeForecastWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagement) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/subscribeAttributeOptOutState(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagement) SubscribeAttributeOptOutStateWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOptOutStateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterDeviceEnergyManagement/subscribeAttributePowerAdjustmentCapability(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterDeviceEnergyManagement) SubscribeAttributePowerAdjustmentCapabilityWithParamsSubscriptionEstablishedReportHandler(params unsafe.Pointer, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePowerAdjustmentCapabilityWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}


