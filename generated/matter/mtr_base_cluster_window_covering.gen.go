// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterWindowCovering */


/* debug [class_header]: Header for MTRBaseClusterWindowCovering */
// The class instance for the [MTRBaseClusterWindowCovering] class.
var (
	MTRBaseClusterWindowCoveringClass     _MTRBaseClusterWindowCoveringClass
	MTRBaseClusterWindowCoveringClassOnce sync.Once
)

func getMTRBaseClusterWindowCoveringClass() _MTRBaseClusterWindowCoveringClass {
	MTRBaseClusterWindowCoveringClassOnce.Do(func() {
		MTRBaseClusterWindowCoveringClass = _MTRBaseClusterWindowCoveringClass{objc.GetClass("MTRBaseClusterWindowCovering")}
	})
	return MTRBaseClusterWindowCoveringClass
}

type _MTRBaseClusterWindowCoveringClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterWindowCovering */
// An interface definition for the [MTRBaseClusterWindowCovering] class.
type IMTRBaseClusterWindowCovering interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterWindowCovering */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterWindowCovering */
	// methods:
	DownOrCloseWithCompletion(completion unsafe.Pointer)
	DownOrCloseWithParamsCompletion(params IMTRWindowCoveringClusterDownOrCloseParams, completion unsafe.Pointer)
	GoToLiftPercentageWithParamsCompletion(params IMTRWindowCoveringClusterGoToLiftPercentageParams, completion unsafe.Pointer)
	GoToLiftValueWithParamsCompletion(params IMTRWindowCoveringClusterGoToLiftValueParams, completion unsafe.Pointer)
	GoToTiltPercentageWithParamsCompletion(params IMTRWindowCoveringClusterGoToTiltPercentageParams, completion unsafe.Pointer)
	GoToTiltValueWithParamsCompletion(params IMTRWindowCoveringClusterGoToTiltValueParams, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeConfigStatusWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentPositionLiftWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentPositionLiftPercent100thsWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentPositionLiftPercentageWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentPositionTiltWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentPositionTiltPercent100thsWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentPositionTiltPercentageWithCompletion(completion unsafe.Pointer)
	ReadAttributeEndProductTypeWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeInstalledClosedLimitLiftWithCompletion(completion unsafe.Pointer)
	ReadAttributeInstalledClosedLimitTiltWithCompletion(completion unsafe.Pointer)
	ReadAttributeInstalledOpenLimitLiftWithCompletion(completion unsafe.Pointer)
	ReadAttributeInstalledOpenLimitTiltWithCompletion(completion unsafe.Pointer)
	ReadAttributeModeWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfActuationsLiftWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfActuationsTiltWithCompletion(completion unsafe.Pointer)
	ReadAttributeOperationalStatusWithCompletion(completion unsafe.Pointer)
	ReadAttributePhysicalClosedLimitLiftWithCompletion(completion unsafe.Pointer)
	ReadAttributePhysicalClosedLimitTiltWithCompletion(completion unsafe.Pointer)
	ReadAttributeSafetyStatusWithCompletion(completion unsafe.Pointer)
	ReadAttributeTargetPositionLiftPercent100thsWithCompletion(completion unsafe.Pointer)
	ReadAttributeTargetPositionTiltPercent100thsWithCompletion(completion unsafe.Pointer)
	ReadAttributeTypeWithCompletion(completion unsafe.Pointer)
	StopMotionWithCompletion(completion unsafe.Pointer)
	StopMotionWithParamsCompletion(params IMTRWindowCoveringClusterStopMotionParams, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeConfigStatusWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentPositionLiftWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentPositionLiftPercent100thsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentPositionLiftPercentageWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentPositionTiltWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentPositionTiltPercent100thsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentPositionTiltPercentageWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeEndProductTypeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeInstalledClosedLimitLiftWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeInstalledClosedLimitTiltWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeInstalledOpenLimitLiftWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeInstalledOpenLimitTiltWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfActuationsLiftWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfActuationsTiltWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOperationalStatusWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePhysicalClosedLimitLiftWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePhysicalClosedLimitTiltWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSafetyStatusWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeTargetPositionLiftPercent100thsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeTargetPositionTiltPercent100thsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeTypeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	UpOrOpenWithCompletion(completion unsafe.Pointer)
	UpOrOpenWithParamsCompletion(params IMTRWindowCoveringClusterUpOrOpenParams, completion unsafe.Pointer)
	WriteAttributeModeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeModeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterWindowCovering */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterWindowCoveringClass) Alloc() MTRBaseClusterWindowCovering {
	rv := objc.Send[MTRBaseClusterWindowCovering](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterWindowCoveringClass) New() MTRBaseClusterWindowCovering {
	rv := objc.Send[MTRBaseClusterWindowCovering](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterWindowCovering) Init() MTRBaseClusterWindowCovering {
	rv := objc.Send[MTRBaseClusterWindowCovering](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterWindowCovering) Autorelease() MTRBaseClusterWindowCovering {
	rv := objc.Send[MTRBaseClusterWindowCovering](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterWindowCovering creates a new MTRBaseClusterWindowCovering instance.
func NewMTRBaseClusterWindowCovering() MTRBaseClusterWindowCovering {
	return getMTRBaseClusterWindowCoveringClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterWindowCovering */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering
type MTRBaseClusterWindowCovering struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterWindowCoveringFrom constructs a [MTRBaseClusterWindowCovering] from an unsafe.Pointer.
func MTRBaseClusterWindowCoveringFrom(ptr unsafe.Pointer) MTRBaseClusterWindowCovering {
	return MTRBaseClusterWindowCovering{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterWindowCovering */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/init(device:endpointID:queue:)
func NewMTRBaseClusterWindowCoveringWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterWindowCovering {
	instance := getMTRBaseClusterWindowCoveringClass().Alloc()
	rv := objc.Send[MTRBaseClusterWindowCovering](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterWindowCoveringWithDeviceEndpointIDQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/init(device:endpoint:queue:)
func NewMTRBaseClusterWindowCoveringWithDeviceEndpointQueue(device IMTRBaseDevice, endpoint uint16 /* not a class type */, queue unsafe.Pointer) MTRBaseClusterWindowCovering {
	instance := getMTRBaseClusterWindowCoveringClass().Alloc()
	rv := objc.Send[MTRBaseClusterWindowCovering](instance.ID, objc.Sel("initWithDevice:endpoint:queue:"), device, endpoint, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterWindowCoveringWithDeviceEndpointQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterWindowCovering */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeAcceptedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeAttributeList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeClusterRevision(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeConfigStatus(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeConfigStatusWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeConfigStatusWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeConfigStatusWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeConfigStatus(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeConfigStatusWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeConfigStatusWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeConfigStatusWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionLift(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeCurrentPositionLiftWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentPositionLiftWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentPositionLiftWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionLift(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeCurrentPositionLiftWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentPositionLiftWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentPositionLiftWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionLiftPercent100ths(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeCurrentPositionLiftPercent100thsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentPositionLiftPercent100thsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentPositionLiftPercent100thsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionLiftPercent100ths(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeCurrentPositionLiftPercent100thsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentPositionLiftPercent100thsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentPositionLiftPercent100thsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionLiftPercentage(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeCurrentPositionLiftPercentageWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentPositionLiftPercentageWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentPositionLiftPercentageWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionLiftPercentage(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeCurrentPositionLiftPercentageWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentPositionLiftPercentageWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentPositionLiftPercentageWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionTilt(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeCurrentPositionTiltWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentPositionTiltWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentPositionTiltWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionTilt(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeCurrentPositionTiltWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentPositionTiltWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentPositionTiltWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionTiltPercent100ths(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeCurrentPositionTiltPercent100thsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentPositionTiltPercent100thsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentPositionTiltPercent100thsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionTiltPercent100ths(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeCurrentPositionTiltPercent100thsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentPositionTiltPercent100thsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentPositionTiltPercent100thsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionTiltPercentage(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeCurrentPositionTiltPercentageWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentPositionTiltPercentageWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentPositionTiltPercentageWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionTiltPercentage(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeCurrentPositionTiltPercentageWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentPositionTiltPercentageWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentPositionTiltPercentageWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeEndProductType(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeEndProductTypeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEndProductTypeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEndProductTypeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeEndProductType(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeEndProductTypeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEndProductTypeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEndProductTypeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeFeatureMap(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeGeneratedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeInstalledClosedLimitLift(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeInstalledClosedLimitLiftWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeInstalledClosedLimitLiftWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeInstalledClosedLimitLiftWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeInstalledClosedLimitLift(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeInstalledClosedLimitLiftWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeInstalledClosedLimitLiftWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeInstalledClosedLimitLiftWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeInstalledClosedLimitTilt(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeInstalledClosedLimitTiltWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeInstalledClosedLimitTiltWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeInstalledClosedLimitTiltWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeInstalledClosedLimitTilt(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeInstalledClosedLimitTiltWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeInstalledClosedLimitTiltWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeInstalledClosedLimitTiltWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeInstalledOpenLimitLift(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeInstalledOpenLimitLiftWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeInstalledOpenLimitLiftWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeInstalledOpenLimitLiftWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeInstalledOpenLimitLift(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeInstalledOpenLimitLiftWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeInstalledOpenLimitLiftWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeInstalledOpenLimitLiftWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeInstalledOpenLimitTilt(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeInstalledOpenLimitTiltWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeInstalledOpenLimitTiltWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeInstalledOpenLimitTiltWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeInstalledOpenLimitTilt(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeInstalledOpenLimitTiltWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeInstalledOpenLimitTiltWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeInstalledOpenLimitTiltWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeMode(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeModeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeModeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeModeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeMode(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeModeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeModeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeModeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeNumberOfActuationsLift(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeNumberOfActuationsLiftWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfActuationsLiftWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfActuationsLiftWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeNumberOfActuationsLift(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeNumberOfActuationsLiftWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfActuationsLiftWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfActuationsLiftWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeNumberOfActuationsTilt(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeNumberOfActuationsTiltWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfActuationsTiltWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfActuationsTiltWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeNumberOfActuationsTilt(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeNumberOfActuationsTiltWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfActuationsTiltWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfActuationsTiltWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeOperationalStatus(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeOperationalStatusWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOperationalStatusWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOperationalStatusWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeOperationalStatus(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeOperationalStatusWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOperationalStatusWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOperationalStatusWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributePhysicalClosedLimitLift(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributePhysicalClosedLimitLiftWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePhysicalClosedLimitLiftWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePhysicalClosedLimitLiftWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributePhysicalClosedLimitLift(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributePhysicalClosedLimitLiftWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePhysicalClosedLimitLiftWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePhysicalClosedLimitLiftWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributePhysicalClosedLimitTilt(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributePhysicalClosedLimitTiltWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePhysicalClosedLimitTiltWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePhysicalClosedLimitTiltWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributePhysicalClosedLimitTilt(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributePhysicalClosedLimitTiltWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePhysicalClosedLimitTiltWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePhysicalClosedLimitTiltWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeSafetyStatus(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeSafetyStatusWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSafetyStatusWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSafetyStatusWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeSafetyStatus(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeSafetyStatusWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSafetyStatusWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSafetyStatusWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeTargetPositionLiftPercent100ths(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeTargetPositionLiftPercent100thsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTargetPositionLiftPercent100thsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeTargetPositionLiftPercent100thsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeTargetPositionLiftPercent100ths(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeTargetPositionLiftPercent100thsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTargetPositionLiftPercent100thsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeTargetPositionLiftPercent100thsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeTargetPositionTiltPercent100ths(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeTargetPositionTiltPercent100thsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTargetPositionTiltPercent100thsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeTargetPositionTiltPercent100thsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeTargetPositionTiltPercent100ths(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeTargetPositionTiltPercent100thsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTargetPositionTiltPercent100thsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeTargetPositionTiltPercent100thsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeType(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeTypeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTypeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeTypeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeType(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterWindowCoveringClass) ReadAttributeTypeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeTypeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeTypeWithClusterStateCacheEndpointQueueCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterWindowCovering */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterWindowCovering */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/downOrClose(completion:)
func (m_ MTRBaseClusterWindowCovering) DownOrCloseWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("downOrCloseWithCompletion:"), completion)
}/* debug [instance_methods/method]: DownOrCloseWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/downOrClose(with:completion:)
func (m_ MTRBaseClusterWindowCovering) DownOrCloseWithParamsCompletion(params IMTRWindowCoveringClusterDownOrCloseParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("downOrCloseWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: DownOrCloseWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/goToLiftPercentage(with:completion:)
func (m_ MTRBaseClusterWindowCovering) GoToLiftPercentageWithParamsCompletion(params IMTRWindowCoveringClusterGoToLiftPercentageParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("goToLiftPercentageWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: GoToLiftPercentageWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/goToLiftValue(with:completion:)
func (m_ MTRBaseClusterWindowCovering) GoToLiftValueWithParamsCompletion(params IMTRWindowCoveringClusterGoToLiftValueParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("goToLiftValueWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: GoToLiftValueWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/goToTiltPercentage(with:completion:)
func (m_ MTRBaseClusterWindowCovering) GoToTiltPercentageWithParamsCompletion(params IMTRWindowCoveringClusterGoToTiltPercentageParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("goToTiltPercentageWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: GoToTiltPercentageWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/goToTiltValue(with:completion:)
func (m_ MTRBaseClusterWindowCovering) GoToTiltValueWithParamsCompletion(params IMTRWindowCoveringClusterGoToTiltValueParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("goToTiltValueWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: GoToTiltValueWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAttributeListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeClusterRevisionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeConfigStatus(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeConfigStatusWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeConfigStatusWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeConfigStatusWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionLift(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeCurrentPositionLiftWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentPositionLiftWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentPositionLiftWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionLiftPercent100ths(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeCurrentPositionLiftPercent100thsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentPositionLiftPercent100thsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentPositionLiftPercent100thsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionLiftPercentage(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeCurrentPositionLiftPercentageWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentPositionLiftPercentageWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentPositionLiftPercentageWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionTilt(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeCurrentPositionTiltWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentPositionTiltWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentPositionTiltWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionTiltPercent100ths(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeCurrentPositionTiltPercent100thsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentPositionTiltPercent100thsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentPositionTiltPercent100thsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeCurrentPositionTiltPercentage(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeCurrentPositionTiltPercentageWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentPositionTiltPercentageWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentPositionTiltPercentageWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeEndProductType(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeEndProductTypeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeEndProductTypeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeEndProductTypeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeGeneratedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeInstalledClosedLimitLift(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeInstalledClosedLimitLiftWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeInstalledClosedLimitLiftWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeInstalledClosedLimitLiftWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeInstalledClosedLimitTilt(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeInstalledClosedLimitTiltWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeInstalledClosedLimitTiltWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeInstalledClosedLimitTiltWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeInstalledOpenLimitLift(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeInstalledOpenLimitLiftWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeInstalledOpenLimitLiftWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeInstalledOpenLimitLiftWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeInstalledOpenLimitTilt(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeInstalledOpenLimitTiltWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeInstalledOpenLimitTiltWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeInstalledOpenLimitTiltWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeMode(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeModeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeModeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeModeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeNumberOfActuationsLift(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeNumberOfActuationsLiftWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfActuationsLiftWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfActuationsLiftWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeNumberOfActuationsTilt(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeNumberOfActuationsTiltWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfActuationsTiltWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfActuationsTiltWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeOperationalStatus(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeOperationalStatusWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOperationalStatusWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOperationalStatusWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributePhysicalClosedLimitLift(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributePhysicalClosedLimitLiftWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePhysicalClosedLimitLiftWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePhysicalClosedLimitLiftWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributePhysicalClosedLimitTilt(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributePhysicalClosedLimitTiltWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePhysicalClosedLimitTiltWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePhysicalClosedLimitTiltWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeSafetyStatus(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeSafetyStatusWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSafetyStatusWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeSafetyStatusWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeTargetPositionLiftPercent100ths(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeTargetPositionLiftPercent100thsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeTargetPositionLiftPercent100thsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeTargetPositionLiftPercent100thsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeTargetPositionTiltPercent100ths(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeTargetPositionTiltPercent100thsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeTargetPositionTiltPercent100thsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeTargetPositionTiltPercent100thsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/readAttributeType(completion:)
func (m_ MTRBaseClusterWindowCovering) ReadAttributeTypeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeTypeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeTypeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/stopMotion(completion:)
func (m_ MTRBaseClusterWindowCovering) StopMotionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopMotionWithCompletion:"), completion)
}/* debug [instance_methods/method]: StopMotionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/stopMotion(with:completion:)
func (m_ MTRBaseClusterWindowCovering) StopMotionWithParamsCompletion(params IMTRWindowCoveringClusterStopMotionParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopMotionWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: StopMotionWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeConfigStatus(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeConfigStatusWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeConfigStatusWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeConfigStatusWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeCurrentPositionLift(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeCurrentPositionLiftWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentPositionLiftWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentPositionLiftWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeCurrentPositionLiftPercent100ths(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeCurrentPositionLiftPercent100thsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentPositionLiftPercent100thsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentPositionLiftPercent100thsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeCurrentPositionLiftPercentage(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeCurrentPositionLiftPercentageWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentPositionLiftPercentageWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentPositionLiftPercentageWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeCurrentPositionTilt(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeCurrentPositionTiltWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentPositionTiltWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentPositionTiltWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeCurrentPositionTiltPercent100ths(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeCurrentPositionTiltPercent100thsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentPositionTiltPercent100thsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentPositionTiltPercent100thsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeCurrentPositionTiltPercentage(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeCurrentPositionTiltPercentageWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentPositionTiltPercentageWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentPositionTiltPercentageWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeEndProductType(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeEndProductTypeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeEndProductTypeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeEndProductTypeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeInstalledClosedLimitLift(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeInstalledClosedLimitLiftWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeInstalledClosedLimitLiftWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeInstalledClosedLimitLiftWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeInstalledClosedLimitTilt(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeInstalledClosedLimitTiltWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeInstalledClosedLimitTiltWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeInstalledClosedLimitTiltWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeInstalledOpenLimitLift(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeInstalledOpenLimitLiftWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeInstalledOpenLimitLiftWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeInstalledOpenLimitLiftWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeInstalledOpenLimitTilt(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeInstalledOpenLimitTiltWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeInstalledOpenLimitTiltWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeInstalledOpenLimitTiltWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeMode(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeModeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeModeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeNumberOfActuationsLift(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeNumberOfActuationsLiftWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfActuationsLiftWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfActuationsLiftWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeNumberOfActuationsTilt(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeNumberOfActuationsTiltWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfActuationsTiltWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfActuationsTiltWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeOperationalStatus(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeOperationalStatusWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOperationalStatusWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOperationalStatusWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributePhysicalClosedLimitLift(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributePhysicalClosedLimitLiftWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePhysicalClosedLimitLiftWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePhysicalClosedLimitLiftWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributePhysicalClosedLimitTilt(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributePhysicalClosedLimitTiltWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePhysicalClosedLimitTiltWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePhysicalClosedLimitTiltWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeSafetyStatus(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeSafetyStatusWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSafetyStatusWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeSafetyStatusWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeTargetPositionLiftPercent100ths(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeTargetPositionLiftPercent100thsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeTargetPositionLiftPercent100thsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeTargetPositionLiftPercent100thsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeTargetPositionTiltPercent100ths(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeTargetPositionTiltPercent100thsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeTargetPositionTiltPercent100thsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeTargetPositionTiltPercent100thsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/subscribeAttributeType(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterWindowCovering) SubscribeAttributeTypeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeTypeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeTypeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/upOrOpen(completion:)
func (m_ MTRBaseClusterWindowCovering) UpOrOpenWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("upOrOpenWithCompletion:"), completion)
}/* debug [instance_methods/method]: UpOrOpenWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/upOrOpen(with:completion:)
func (m_ MTRBaseClusterWindowCovering) UpOrOpenWithParamsCompletion(params IMTRWindowCoveringClusterUpOrOpenParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("upOrOpenWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: UpOrOpenWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/writeAttributeMode(withValue:completion:)
func (m_ MTRBaseClusterWindowCovering) WriteAttributeModeWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeModeWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeModeWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterWindowCovering/writeAttributeMode(withValue:params:completion:)
func (m_ MTRBaseClusterWindowCovering) WriteAttributeModeWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeModeWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeModeWithValueParamsCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterWindowCovering */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterWindowCovering */


