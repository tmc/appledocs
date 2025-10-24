// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterLevelControl */


/* debug [class_header]: Header for MTRBaseClusterLevelControl */
// The class instance for the [MTRBaseClusterLevelControl] class.
var (
	MTRBaseClusterLevelControlClass     _MTRBaseClusterLevelControlClass
	MTRBaseClusterLevelControlClassOnce sync.Once
)

func getMTRBaseClusterLevelControlClass() _MTRBaseClusterLevelControlClass {
	MTRBaseClusterLevelControlClassOnce.Do(func() {
		MTRBaseClusterLevelControlClass = _MTRBaseClusterLevelControlClass{objc.GetClass("MTRBaseClusterLevelControl")}
	})
	return MTRBaseClusterLevelControlClass
}

type _MTRBaseClusterLevelControlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterLevelControl */
// An interface definition for the [MTRBaseClusterLevelControl] class.
type IMTRBaseClusterLevelControl interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterLevelControl */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterLevelControl */
	// methods:
	MoveWithParamsCompletion(params IMTRLevelControlClusterMoveParams, completion unsafe.Pointer)
	MoveToClosestFrequencyWithParamsCompletion(params IMTRLevelControlClusterMoveToClosestFrequencyParams, completion unsafe.Pointer)
	MoveToLevelWithParamsCompletion(params IMTRLevelControlClusterMoveToLevelParams, completion unsafe.Pointer)
	MoveToLevelWithOnOffWithParamsCompletion(params IMTRLevelControlClusterMoveToLevelWithOnOffParams, completion unsafe.Pointer)
	MoveWithOnOffWithParamsCompletion(params IMTRLevelControlClusterMoveWithOnOffParams, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentFrequencyWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentLevelWithCompletion(completion unsafe.Pointer)
	ReadAttributeDefaultMoveRateWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeMaxFrequencyWithCompletion(completion unsafe.Pointer)
	ReadAttributeMaxLevelWithCompletion(completion unsafe.Pointer)
	ReadAttributeMinFrequencyWithCompletion(completion unsafe.Pointer)
	ReadAttributeMinLevelWithCompletion(completion unsafe.Pointer)
	ReadAttributeOffTransitionTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeOnLevelWithCompletion(completion unsafe.Pointer)
	ReadAttributeOnOffTransitionTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeOnTransitionTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeOptionsWithCompletion(completion unsafe.Pointer)
	ReadAttributeRemainingTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeStartUpCurrentLevelWithCompletion(completion unsafe.Pointer)
	StepWithParamsCompletion(params IMTRLevelControlClusterStepParams, completion unsafe.Pointer)
	StepWithOnOffWithParamsCompletion(params IMTRLevelControlClusterStepWithOnOffParams, completion unsafe.Pointer)
	StopWithParamsCompletion(params IMTRLevelControlClusterStopParams, completion unsafe.Pointer)
	StopWithOnOffWithParamsCompletion(params IMTRLevelControlClusterStopWithOnOffParams, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentFrequencyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentLevelWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeDefaultMoveRateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMaxFrequencyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMaxLevelWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMinFrequencyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeMinLevelWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOffTransitionTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOnLevelWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOnOffTransitionTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOnTransitionTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOptionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeRemainingTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeStartUpCurrentLevelWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	WriteAttributeDefaultMoveRateWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeDefaultMoveRateWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeOffTransitionTimeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeOffTransitionTimeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeOnLevelWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeOnLevelWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeOnOffTransitionTimeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeOnOffTransitionTimeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeOnTransitionTimeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeOnTransitionTimeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeOptionsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeOptionsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeStartUpCurrentLevelWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeStartUpCurrentLevelWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterLevelControl */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterLevelControlClass) Alloc() MTRBaseClusterLevelControl {
	rv := objc.Send[MTRBaseClusterLevelControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterLevelControlClass) New() MTRBaseClusterLevelControl {
	rv := objc.Send[MTRBaseClusterLevelControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterLevelControl) Init() MTRBaseClusterLevelControl {
	rv := objc.Send[MTRBaseClusterLevelControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterLevelControl) Autorelease() MTRBaseClusterLevelControl {
	rv := objc.Send[MTRBaseClusterLevelControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterLevelControl creates a new MTRBaseClusterLevelControl instance.
func NewMTRBaseClusterLevelControl() MTRBaseClusterLevelControl {
	return getMTRBaseClusterLevelControlClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterLevelControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl
type MTRBaseClusterLevelControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterLevelControlFrom constructs a [MTRBaseClusterLevelControl] from an unsafe.Pointer.
func MTRBaseClusterLevelControlFrom(ptr unsafe.Pointer) MTRBaseClusterLevelControl {
	return MTRBaseClusterLevelControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterLevelControl */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/init(device:endpointID:queue:)
func NewMTRBaseClusterLevelControlWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterLevelControl {
	instance := getMTRBaseClusterLevelControlClass().Alloc()
	rv := objc.Send[MTRBaseClusterLevelControl](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterLevelControlWithDeviceEndpointIDQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/init(device:endpoint:queue:)
func NewMTRBaseClusterLevelControlWithDeviceEndpointQueue(device IMTRBaseDevice, endpoint uint16 /* not a class type */, queue unsafe.Pointer) MTRBaseClusterLevelControl {
	instance := getMTRBaseClusterLevelControlClass().Alloc()
	rv := objc.Send[MTRBaseClusterLevelControl](instance.ID, objc.Sel("initWithDevice:endpoint:queue:"), device, endpoint, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterLevelControlWithDeviceEndpointQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterLevelControl */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeAcceptedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeAttributeList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeClusterRevision(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeCurrentFrequency(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeCurrentFrequencyWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentFrequencyWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentFrequencyWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeCurrentFrequency(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeCurrentFrequencyWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentFrequencyWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentFrequencyWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeCurrentLevel(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeCurrentLevelWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentLevelWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentLevelWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeCurrentLevel(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeCurrentLevelWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentLevelWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentLevelWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeDefaultMoveRate(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeDefaultMoveRateWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDefaultMoveRateWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeDefaultMoveRateWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeDefaultMoveRate(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeDefaultMoveRateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDefaultMoveRateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeDefaultMoveRateWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeFeatureMap(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeGeneratedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeMaxFrequency(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeMaxFrequencyWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxFrequencyWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMaxFrequencyWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeMaxFrequency(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeMaxFrequencyWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxFrequencyWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMaxFrequencyWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeMaxLevel(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeMaxLevelWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxLevelWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMaxLevelWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeMaxLevel(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeMaxLevelWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMaxLevelWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMaxLevelWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeMinFrequency(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeMinFrequencyWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinFrequencyWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinFrequencyWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeMinFrequency(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeMinFrequencyWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinFrequencyWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinFrequencyWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeMinLevel(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeMinLevelWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinLevelWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinLevelWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeMinLevel(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeMinLevelWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeMinLevelWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeMinLevelWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOffTransitionTime(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeOffTransitionTimeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOffTransitionTimeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOffTransitionTimeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOffTransitionTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeOffTransitionTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOffTransitionTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOffTransitionTimeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOnLevel(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeOnLevelWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOnLevelWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOnLevelWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOnLevel(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeOnLevelWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOnLevelWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOnLevelWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOnOffTransitionTime(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeOnOffTransitionTimeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOnOffTransitionTimeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOnOffTransitionTimeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOnOffTransitionTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeOnOffTransitionTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOnOffTransitionTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOnOffTransitionTimeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOnTransitionTime(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeOnTransitionTimeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOnTransitionTimeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOnTransitionTimeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOnTransitionTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeOnTransitionTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOnTransitionTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOnTransitionTimeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOptions(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeOptionsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOptionsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOptionsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOptions(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeOptionsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOptionsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOptionsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeRemainingTime(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeRemainingTimeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeRemainingTimeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeRemainingTimeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeRemainingTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeRemainingTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeRemainingTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeRemainingTimeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeStartUpCurrentLevel(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeStartUpCurrentLevelWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeStartUpCurrentLevelWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeStartUpCurrentLevelWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeStartUpCurrentLevel(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterLevelControlClass) ReadAttributeStartUpCurrentLevelWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeStartUpCurrentLevelWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeStartUpCurrentLevelWithClusterStateCacheEndpointQueueCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterLevelControl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterLevelControl */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/move(with:completion:)
func (m_ MTRBaseClusterLevelControl) MoveWithParamsCompletion(params IMTRLevelControlClusterMoveParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/moveToClosestFrequency(with:completion:)
func (m_ MTRBaseClusterLevelControl) MoveToClosestFrequencyWithParamsCompletion(params IMTRLevelControlClusterMoveToClosestFrequencyParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveToClosestFrequencyWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveToClosestFrequencyWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/moveToLevel(with:completion:)
func (m_ MTRBaseClusterLevelControl) MoveToLevelWithParamsCompletion(params IMTRLevelControlClusterMoveToLevelParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveToLevelWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveToLevelWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/moveToLevelWithOnOff(with:completion:)
func (m_ MTRBaseClusterLevelControl) MoveToLevelWithOnOffWithParamsCompletion(params IMTRLevelControlClusterMoveToLevelWithOnOffParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveToLevelWithOnOffWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveToLevelWithOnOffWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/moveWithOnOff(with:completion:)
func (m_ MTRBaseClusterLevelControl) MoveWithOnOffWithParamsCompletion(params IMTRLevelControlClusterMoveWithOnOffParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveWithOnOffWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveWithOnOffWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAttributeListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeClusterRevisionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeCurrentFrequency(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeCurrentFrequencyWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentFrequencyWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentFrequencyWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeCurrentLevel(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeCurrentLevelWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentLevelWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentLevelWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeDefaultMoveRate(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeDefaultMoveRateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeDefaultMoveRateWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeDefaultMoveRateWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeGeneratedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeMaxFrequency(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeMaxFrequencyWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMaxFrequencyWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeMaxFrequencyWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeMaxLevel(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeMaxLevelWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMaxLevelWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeMaxLevelWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeMinFrequency(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeMinFrequencyWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMinFrequencyWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeMinFrequencyWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeMinLevel(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeMinLevelWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeMinLevelWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeMinLevelWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOffTransitionTime(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeOffTransitionTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOffTransitionTimeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOffTransitionTimeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOnLevel(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeOnLevelWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOnLevelWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOnLevelWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOnOffTransitionTime(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeOnOffTransitionTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOnOffTransitionTimeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOnOffTransitionTimeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOnTransitionTime(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeOnTransitionTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOnTransitionTimeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOnTransitionTimeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeOptions(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeOptionsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOptionsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOptionsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeRemainingTime(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeRemainingTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeRemainingTimeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeRemainingTimeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/readAttributeStartUpCurrentLevel(completion:)
func (m_ MTRBaseClusterLevelControl) ReadAttributeStartUpCurrentLevelWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeStartUpCurrentLevelWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeStartUpCurrentLevelWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/step(with:completion:)
func (m_ MTRBaseClusterLevelControl) StepWithParamsCompletion(params IMTRLevelControlClusterStepParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stepWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: StepWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/stepWithOnOff(with:completion:)
func (m_ MTRBaseClusterLevelControl) StepWithOnOffWithParamsCompletion(params IMTRLevelControlClusterStepWithOnOffParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stepWithOnOffWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: StepWithOnOffWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/stop(with:completion:)
func (m_ MTRBaseClusterLevelControl) StopWithParamsCompletion(params IMTRLevelControlClusterStopParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: StopWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/stopWithOnOff(with:completion:)
func (m_ MTRBaseClusterLevelControl) StopWithOnOffWithParamsCompletion(params IMTRLevelControlClusterStopWithOnOffParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopWithOnOffWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: StopWithOnOffWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeCurrentFrequency(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeCurrentFrequencyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentFrequencyWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentFrequencyWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeCurrentLevel(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeCurrentLevelWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentLevelWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentLevelWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeDefaultMoveRate(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeDefaultMoveRateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeDefaultMoveRateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeDefaultMoveRateWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeMaxFrequency(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeMaxFrequencyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMaxFrequencyWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeMaxFrequencyWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeMaxLevel(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeMaxLevelWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMaxLevelWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeMaxLevelWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeMinFrequency(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeMinFrequencyWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMinFrequencyWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeMinFrequencyWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeMinLevel(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeMinLevelWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeMinLevelWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeMinLevelWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeOffTransitionTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeOffTransitionTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOffTransitionTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOffTransitionTimeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeOnLevel(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeOnLevelWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOnLevelWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOnLevelWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeOnOffTransitionTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeOnOffTransitionTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOnOffTransitionTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOnOffTransitionTimeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeOnTransitionTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeOnTransitionTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOnTransitionTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOnTransitionTimeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeOptions(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeOptionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOptionsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOptionsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeRemainingTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeRemainingTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeRemainingTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeRemainingTimeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/subscribeAttributeStartUpCurrentLevel(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterLevelControl) SubscribeAttributeStartUpCurrentLevelWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeStartUpCurrentLevelWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeStartUpCurrentLevelWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeDefaultMoveRate(withValue:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeDefaultMoveRateWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeDefaultMoveRateWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeDefaultMoveRateWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeDefaultMoveRate(withValue:params:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeDefaultMoveRateWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeDefaultMoveRateWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeDefaultMoveRateWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeOffTransitionTime(withValue:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeOffTransitionTimeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOffTransitionTimeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeOffTransitionTimeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeOffTransitionTime(withValue:params:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeOffTransitionTimeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOffTransitionTimeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeOffTransitionTimeWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeOnLevel(withValue:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeOnLevelWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOnLevelWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeOnLevelWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeOnLevel(withValue:params:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeOnLevelWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOnLevelWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeOnLevelWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeOnOffTransitionTime(withValue:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeOnOffTransitionTimeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOnOffTransitionTimeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeOnOffTransitionTimeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeOnOffTransitionTime(withValue:params:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeOnOffTransitionTimeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOnOffTransitionTimeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeOnOffTransitionTimeWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeOnTransitionTime(withValue:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeOnTransitionTimeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOnTransitionTimeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeOnTransitionTimeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeOnTransitionTime(withValue:params:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeOnTransitionTimeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOnTransitionTimeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeOnTransitionTimeWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeOptions(withValue:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeOptionsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOptionsWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeOptionsWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeOptions(withValue:params:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeOptionsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOptionsWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeOptionsWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeStartUpCurrentLevel(withValue:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeStartUpCurrentLevelWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeStartUpCurrentLevelWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeStartUpCurrentLevelWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterLevelControl/writeAttributeStartUpCurrentLevel(withValue:params:completion:)
func (m_ MTRBaseClusterLevelControl) WriteAttributeStartUpCurrentLevelWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeStartUpCurrentLevelWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeStartUpCurrentLevelWithValueParamsCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterLevelControl */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterLevelControl */


