// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterColorControl */


/* debug [class_header]: Header for MTRBaseClusterColorControl */
// The class instance for the [MTRBaseClusterColorControl] class.
var (
	MTRBaseClusterColorControlClass     _MTRBaseClusterColorControlClass
	MTRBaseClusterColorControlClassOnce sync.Once
)

func getMTRBaseClusterColorControlClass() _MTRBaseClusterColorControlClass {
	MTRBaseClusterColorControlClassOnce.Do(func() {
		MTRBaseClusterColorControlClass = _MTRBaseClusterColorControlClass{objc.GetClass("MTRBaseClusterColorControl")}
	})
	return MTRBaseClusterColorControlClass
}

type _MTRBaseClusterColorControlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterColorControl */
// An interface definition for the [MTRBaseClusterColorControl] class.
type IMTRBaseClusterColorControl interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterColorControl */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterColorControl */
	// methods:
	ColorLoopSetWithParamsCompletion(params IMTRColorControlClusterColorLoopSetParams, completion unsafe.Pointer)
	EnhancedMoveHueWithParamsCompletion(params IMTRColorControlClusterEnhancedMoveHueParams, completion unsafe.Pointer)
	EnhancedMoveToHueWithParamsCompletion(params IMTRColorControlClusterEnhancedMoveToHueParams, completion unsafe.Pointer)
	EnhancedMoveToHueAndSaturationWithParamsCompletion(params IMTRColorControlClusterEnhancedMoveToHueAndSaturationParams, completion unsafe.Pointer)
	EnhancedStepHueWithParamsCompletion(params IMTRColorControlClusterEnhancedStepHueParams, completion unsafe.Pointer)
	MoveColorWithParamsCompletion(params IMTRColorControlClusterMoveColorParams, completion unsafe.Pointer)
	MoveColorTemperatureWithParamsCompletion(params IMTRColorControlClusterMoveColorTemperatureParams, completion unsafe.Pointer)
	MoveHueWithParamsCompletion(params IMTRColorControlClusterMoveHueParams, completion unsafe.Pointer)
	MoveSaturationWithParamsCompletion(params IMTRColorControlClusterMoveSaturationParams, completion unsafe.Pointer)
	MoveToColorWithParamsCompletion(params IMTRColorControlClusterMoveToColorParams, completion unsafe.Pointer)
	MoveToColorTemperatureWithParamsCompletion(params IMTRColorControlClusterMoveToColorTemperatureParams, completion unsafe.Pointer)
	MoveToHueWithParamsCompletion(params IMTRColorControlClusterMoveToHueParams, completion unsafe.Pointer)
	MoveToHueAndSaturationWithParamsCompletion(params IMTRColorControlClusterMoveToHueAndSaturationParams, completion unsafe.Pointer)
	MoveToSaturationWithParamsCompletion(params IMTRColorControlClusterMoveToSaturationParams, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorCapabilitiesWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorLoopActiveWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorLoopDirectionWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorLoopStartEnhancedHueWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorLoopStoredEnhancedHueWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorLoopTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorModeWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorPointBIntensityWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorPointBXWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorPointBYWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorPointGIntensityWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorPointGXWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorPointGYWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorPointRIntensityWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorPointRXWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorPointRYWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorTemperatureMiredsWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorTempPhysicalMaxMiredsWithCompletion(completion unsafe.Pointer)
	ReadAttributeColorTempPhysicalMinMiredsWithCompletion(completion unsafe.Pointer)
	ReadAttributeCompensationTextWithCompletion(completion unsafe.Pointer)
	ReadAttributeCoupleColorTempToLevelMinMiredsWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentHueWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentSaturationWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentXWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentYWithCompletion(completion unsafe.Pointer)
	ReadAttributeDriftCompensationWithCompletion(completion unsafe.Pointer)
	ReadAttributeEnhancedColorModeWithCompletion(completion unsafe.Pointer)
	ReadAttributeEnhancedCurrentHueWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeNumberOfPrimariesWithCompletion(completion unsafe.Pointer)
	ReadAttributeOptionsWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary1IntensityWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary1XWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary1YWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary2IntensityWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary2XWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary2YWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary3IntensityWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary3XWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary3YWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary4IntensityWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary4XWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary4YWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary5IntensityWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary5XWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary5YWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary6IntensityWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary6XWithCompletion(completion unsafe.Pointer)
	ReadAttributePrimary6YWithCompletion(completion unsafe.Pointer)
	ReadAttributeRemainingTimeWithCompletion(completion unsafe.Pointer)
	ReadAttributeStartUpColorTemperatureMiredsWithCompletion(completion unsafe.Pointer)
	ReadAttributeWhitePointXWithCompletion(completion unsafe.Pointer)
	ReadAttributeWhitePointYWithCompletion(completion unsafe.Pointer)
	StepColorWithParamsCompletion(params IMTRColorControlClusterStepColorParams, completion unsafe.Pointer)
	StepColorTemperatureWithParamsCompletion(params IMTRColorControlClusterStepColorTemperatureParams, completion unsafe.Pointer)
	StepHueWithParamsCompletion(params IMTRColorControlClusterStepHueParams, completion unsafe.Pointer)
	StepSaturationWithParamsCompletion(params IMTRColorControlClusterStepSaturationParams, completion unsafe.Pointer)
	StopMoveStepWithParamsCompletion(params IMTRColorControlClusterStopMoveStepParams, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorCapabilitiesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorLoopActiveWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorLoopDirectionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorLoopStartEnhancedHueWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorLoopStoredEnhancedHueWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorLoopTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorPointBIntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorPointBXWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorPointBYWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorPointGIntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorPointGXWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorPointGYWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorPointRIntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorPointRXWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorPointRYWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorTemperatureMiredsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorTempPhysicalMaxMiredsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeColorTempPhysicalMinMiredsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCompensationTextWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCoupleColorTempToLevelMinMiredsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentHueWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentSaturationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentXWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentYWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeDriftCompensationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeEnhancedColorModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeEnhancedCurrentHueWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeNumberOfPrimariesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeOptionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary1IntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary1XWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary1YWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary2IntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary2XWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary2YWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary3IntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary3XWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary3YWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary4IntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary4XWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary4YWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary5IntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary5XWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary5YWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary6IntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary6XWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributePrimary6YWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeRemainingTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeStartUpColorTemperatureMiredsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeWhitePointXWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeWhitePointYWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	WriteAttributeOptionsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeOptionsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
	WriteAttributeStartUpColorTemperatureMiredsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeStartUpColorTemperatureMiredsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterColorControl */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterColorControlClass) Alloc() MTRBaseClusterColorControl {
	rv := objc.Send[MTRBaseClusterColorControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterColorControlClass) New() MTRBaseClusterColorControl {
	rv := objc.Send[MTRBaseClusterColorControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterColorControl) Init() MTRBaseClusterColorControl {
	rv := objc.Send[MTRBaseClusterColorControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterColorControl) Autorelease() MTRBaseClusterColorControl {
	rv := objc.Send[MTRBaseClusterColorControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterColorControl creates a new MTRBaseClusterColorControl instance.
func NewMTRBaseClusterColorControl() MTRBaseClusterColorControl {
	return getMTRBaseClusterColorControlClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterColorControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl
type MTRBaseClusterColorControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterColorControlFrom constructs a [MTRBaseClusterColorControl] from an unsafe.Pointer.
func MTRBaseClusterColorControlFrom(ptr unsafe.Pointer) MTRBaseClusterColorControl {
	return MTRBaseClusterColorControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterColorControl */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/init(device:endpointID:queue:)
func NewMTRBaseClusterColorControlWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterColorControl {
	instance := getMTRBaseClusterColorControlClass().Alloc()
	rv := objc.Send[MTRBaseClusterColorControl](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterColorControlWithDeviceEndpointIDQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/init(device:endpoint:queue:)
func NewMTRBaseClusterColorControlWithDeviceEndpointQueue(device IMTRBaseDevice, endpoint uint16 /* not a class type */, queue unsafe.Pointer) MTRBaseClusterColorControl {
	instance := getMTRBaseClusterColorControlClass().Alloc()
	rv := objc.Send[MTRBaseClusterColorControl](instance.ID, objc.Sel("initWithDevice:endpoint:queue:"), device, endpoint, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterColorControlWithDeviceEndpointQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterColorControl */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeAcceptedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeAttributeList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeClusterRevision(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorCapabilities(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorCapabilitiesWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorCapabilitiesWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorCapabilitiesWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorCapabilities(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorCapabilitiesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorCapabilitiesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorCapabilitiesWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopActive(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorLoopActiveWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorLoopActiveWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorLoopActiveWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopActive(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorLoopActiveWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorLoopActiveWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorLoopActiveWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopDirection(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorLoopDirectionWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorLoopDirectionWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorLoopDirectionWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopDirection(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorLoopDirectionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorLoopDirectionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorLoopDirectionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopStartEnhancedHue(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorLoopStartEnhancedHueWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorLoopStartEnhancedHueWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorLoopStartEnhancedHueWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopStartEnhancedHue(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorLoopStartEnhancedHueWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorLoopStartEnhancedHueWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorLoopStartEnhancedHueWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopStoredEnhancedHue(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorLoopStoredEnhancedHueWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorLoopStoredEnhancedHueWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorLoopStoredEnhancedHueWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopStoredEnhancedHue(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorLoopStoredEnhancedHueWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorLoopStoredEnhancedHueWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorLoopStoredEnhancedHueWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopTime(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorLoopTimeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorLoopTimeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorLoopTimeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorLoopTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorLoopTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorLoopTimeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorMode(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorModeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorModeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorModeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorMode(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorModeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorModeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorModeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointBIntensity(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointBIntensityWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointBIntensityWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointBIntensityWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointBIntensity(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointBIntensityWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointBIntensityWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointBIntensityWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointBX(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointBXWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointBXWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointBXWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointBX(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointBXWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointBXWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointBXWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointBY(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointBYWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointBYWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointBYWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointBY(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointBYWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointBYWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointBYWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointGIntensity(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointGIntensityWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointGIntensityWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointGIntensityWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointGIntensity(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointGIntensityWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointGIntensityWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointGIntensityWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointGX(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointGXWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointGXWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointGXWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointGX(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointGXWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointGXWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointGXWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointGY(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointGYWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointGYWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointGYWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointGY(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointGYWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointGYWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointGYWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointRIntensity(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointRIntensityWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointRIntensityWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointRIntensityWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointRIntensity(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointRIntensityWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointRIntensityWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointRIntensityWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointRX(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointRXWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointRXWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointRXWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointRX(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointRXWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointRXWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointRXWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointRY(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointRYWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointRYWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointRYWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointRY(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorPointRYWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorPointRYWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorPointRYWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorTemperatureMireds(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorTemperatureMiredsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorTemperatureMiredsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorTemperatureMiredsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorTemperatureMireds(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorTemperatureMiredsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorTemperatureMiredsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorTemperatureMiredsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorTempPhysicalMaxMireds(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorTempPhysicalMaxMiredsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorTempPhysicalMaxMiredsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorTempPhysicalMaxMiredsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorTempPhysicalMaxMireds(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorTempPhysicalMaxMiredsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorTempPhysicalMaxMiredsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorTempPhysicalMaxMiredsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorTempPhysicalMinMireds(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorTempPhysicalMinMiredsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorTempPhysicalMinMiredsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorTempPhysicalMinMiredsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorTempPhysicalMinMireds(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeColorTempPhysicalMinMiredsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeColorTempPhysicalMinMiredsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeColorTempPhysicalMinMiredsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCompensationText(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeCompensationTextWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCompensationTextWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCompensationTextWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCompensationText(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeCompensationTextWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCompensationTextWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCompensationTextWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCoupleColorTempToLevelMinMireds(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeCoupleColorTempToLevelMinMiredsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCoupleColorTempToLevelMinMiredsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCoupleColorTempToLevelMinMiredsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCoupleColorTempToLevelMinMireds(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeCoupleColorTempToLevelMinMiredsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCoupleColorTempToLevelMinMiredsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCoupleColorTempToLevelMinMiredsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCurrentHue(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeCurrentHueWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentHueWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentHueWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCurrentHue(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeCurrentHueWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentHueWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentHueWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCurrentSaturation(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeCurrentSaturationWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentSaturationWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentSaturationWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCurrentSaturation(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeCurrentSaturationWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentSaturationWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentSaturationWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCurrentX(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeCurrentXWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentXWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentXWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCurrentX(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeCurrentXWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentXWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentXWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCurrentY(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeCurrentYWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentYWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentYWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCurrentY(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeCurrentYWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentYWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentYWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeDriftCompensation(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeDriftCompensationWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDriftCompensationWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeDriftCompensationWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeDriftCompensation(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeDriftCompensationWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDriftCompensationWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeDriftCompensationWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeEnhancedColorMode(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeEnhancedColorModeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEnhancedColorModeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEnhancedColorModeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeEnhancedColorMode(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeEnhancedColorModeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEnhancedColorModeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEnhancedColorModeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeEnhancedCurrentHue(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeEnhancedCurrentHueWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEnhancedCurrentHueWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEnhancedCurrentHueWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeEnhancedCurrentHue(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeEnhancedCurrentHueWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeEnhancedCurrentHueWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeEnhancedCurrentHueWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeFeatureMap(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeGeneratedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeNumberOfPrimaries(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeNumberOfPrimariesWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfPrimariesWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfPrimariesWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeNumberOfPrimaries(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeNumberOfPrimariesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeNumberOfPrimariesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeNumberOfPrimariesWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeOptions(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeOptionsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOptionsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOptionsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeOptions(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeOptionsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeOptionsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeOptionsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary1Intensity(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary1IntensityWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary1IntensityWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary1IntensityWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary1Intensity(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary1IntensityWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary1IntensityWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary1IntensityWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary1X(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary1XWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary1XWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary1XWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary1X(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary1XWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary1XWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary1XWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary1Y(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary1YWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary1YWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary1YWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary1Y(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary1YWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary1YWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary1YWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary2Intensity(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary2IntensityWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary2IntensityWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary2IntensityWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary2Intensity(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary2IntensityWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary2IntensityWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary2IntensityWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary2X(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary2XWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary2XWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary2XWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary2X(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary2XWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary2XWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary2XWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary2Y(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary2YWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary2YWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary2YWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary2Y(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary2YWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary2YWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary2YWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary3Intensity(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary3IntensityWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary3IntensityWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary3IntensityWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary3Intensity(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary3IntensityWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary3IntensityWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary3IntensityWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary3X(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary3XWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary3XWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary3XWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary3X(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary3XWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary3XWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary3XWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary3Y(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary3YWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary3YWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary3YWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary3Y(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary3YWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary3YWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary3YWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary4Intensity(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary4IntensityWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary4IntensityWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary4IntensityWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary4Intensity(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary4IntensityWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary4IntensityWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary4IntensityWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary4X(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary4XWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary4XWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary4XWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary4X(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary4XWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary4XWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary4XWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary4Y(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary4YWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary4YWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary4YWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary4Y(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary4YWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary4YWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary4YWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary5Intensity(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary5IntensityWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary5IntensityWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary5IntensityWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary5Intensity(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary5IntensityWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary5IntensityWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary5IntensityWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary5X(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary5XWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary5XWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary5XWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary5X(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary5XWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary5XWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary5XWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary5Y(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary5YWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary5YWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary5YWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary5Y(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary5YWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary5YWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary5YWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary6Intensity(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary6IntensityWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary6IntensityWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary6IntensityWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary6Intensity(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary6IntensityWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary6IntensityWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary6IntensityWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary6X(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary6XWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary6XWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary6XWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary6X(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary6XWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary6XWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary6XWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary6Y(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary6YWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary6YWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary6YWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary6Y(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributePrimary6YWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributePrimary6YWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributePrimary6YWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeRemainingTime(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeRemainingTimeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeRemainingTimeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeRemainingTimeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeRemainingTime(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeRemainingTimeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeRemainingTimeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeRemainingTimeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeStartUpColorTemperatureMireds(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeStartUpColorTemperatureMiredsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeStartUpColorTemperatureMiredsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeStartUpColorTemperatureMiredsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeStartUpColorTemperatureMireds(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeStartUpColorTemperatureMiredsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeStartUpColorTemperatureMiredsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeStartUpColorTemperatureMiredsWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeWhitePointX(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeWhitePointXWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeWhitePointXWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeWhitePointXWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeWhitePointX(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeWhitePointXWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeWhitePointXWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeWhitePointXWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeWhitePointY(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeWhitePointYWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeWhitePointYWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeWhitePointYWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeWhitePointY(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterColorControlClass) ReadAttributeWhitePointYWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeWhitePointYWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeWhitePointYWithClusterStateCacheEndpointQueueCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterColorControl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterColorControl */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/colorLoopSet(with:completion:)
func (m_ MTRBaseClusterColorControl) ColorLoopSetWithParamsCompletion(params IMTRColorControlClusterColorLoopSetParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("colorLoopSetWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: ColorLoopSetWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/enhancedMoveHue(with:completion:)
func (m_ MTRBaseClusterColorControl) EnhancedMoveHueWithParamsCompletion(params IMTRColorControlClusterEnhancedMoveHueParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("enhancedMoveHueWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: EnhancedMoveHueWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/enhancedMoveToHue(with:completion:)
func (m_ MTRBaseClusterColorControl) EnhancedMoveToHueWithParamsCompletion(params IMTRColorControlClusterEnhancedMoveToHueParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("enhancedMoveToHueWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: EnhancedMoveToHueWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/enhancedMoveToHueAndSaturation(with:completion:)
func (m_ MTRBaseClusterColorControl) EnhancedMoveToHueAndSaturationWithParamsCompletion(params IMTRColorControlClusterEnhancedMoveToHueAndSaturationParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("enhancedMoveToHueAndSaturationWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: EnhancedMoveToHueAndSaturationWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/enhancedStepHue(with:completion:)
func (m_ MTRBaseClusterColorControl) EnhancedStepHueWithParamsCompletion(params IMTRColorControlClusterEnhancedStepHueParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("enhancedStepHueWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: EnhancedStepHueWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/moveColor(with:completion:)
func (m_ MTRBaseClusterColorControl) MoveColorWithParamsCompletion(params IMTRColorControlClusterMoveColorParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveColorWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveColorWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/moveColorTemperature(with:completion:)
func (m_ MTRBaseClusterColorControl) MoveColorTemperatureWithParamsCompletion(params IMTRColorControlClusterMoveColorTemperatureParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveColorTemperatureWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveColorTemperatureWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/moveHue(with:completion:)
func (m_ MTRBaseClusterColorControl) MoveHueWithParamsCompletion(params IMTRColorControlClusterMoveHueParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveHueWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveHueWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/moveSaturation(with:completion:)
func (m_ MTRBaseClusterColorControl) MoveSaturationWithParamsCompletion(params IMTRColorControlClusterMoveSaturationParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveSaturationWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveSaturationWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/moveToColor(with:completion:)
func (m_ MTRBaseClusterColorControl) MoveToColorWithParamsCompletion(params IMTRColorControlClusterMoveToColorParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveToColorWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveToColorWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/moveToColorTemperature(with:completion:)
func (m_ MTRBaseClusterColorControl) MoveToColorTemperatureWithParamsCompletion(params IMTRColorControlClusterMoveToColorTemperatureParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveToColorTemperatureWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveToColorTemperatureWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/moveToHue(with:completion:)
func (m_ MTRBaseClusterColorControl) MoveToHueWithParamsCompletion(params IMTRColorControlClusterMoveToHueParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveToHueWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveToHueWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/moveToHueAndSaturation(with:completion:)
func (m_ MTRBaseClusterColorControl) MoveToHueAndSaturationWithParamsCompletion(params IMTRColorControlClusterMoveToHueAndSaturationParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveToHueAndSaturationWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveToHueAndSaturationWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/moveToSaturation(with:completion:)
func (m_ MTRBaseClusterColorControl) MoveToSaturationWithParamsCompletion(params IMTRColorControlClusterMoveToSaturationParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("moveToSaturationWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: MoveToSaturationWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAttributeListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeClusterRevisionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorCapabilities(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorCapabilitiesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorCapabilitiesWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorCapabilitiesWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopActive(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorLoopActiveWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorLoopActiveWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorLoopActiveWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopDirection(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorLoopDirectionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorLoopDirectionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorLoopDirectionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopStartEnhancedHue(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorLoopStartEnhancedHueWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorLoopStartEnhancedHueWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorLoopStartEnhancedHueWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopStoredEnhancedHue(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorLoopStoredEnhancedHueWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorLoopStoredEnhancedHueWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorLoopStoredEnhancedHueWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorLoopTime(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorLoopTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorLoopTimeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorLoopTimeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorMode(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorModeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorModeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorModeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointBIntensity(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorPointBIntensityWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorPointBIntensityWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorPointBIntensityWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointBX(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorPointBXWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorPointBXWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorPointBXWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointBY(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorPointBYWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorPointBYWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorPointBYWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointGIntensity(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorPointGIntensityWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorPointGIntensityWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorPointGIntensityWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointGX(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorPointGXWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorPointGXWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorPointGXWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointGY(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorPointGYWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorPointGYWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorPointGYWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointRIntensity(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorPointRIntensityWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorPointRIntensityWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorPointRIntensityWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointRX(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorPointRXWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorPointRXWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorPointRXWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorPointRY(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorPointRYWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorPointRYWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorPointRYWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorTemperatureMireds(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorTemperatureMiredsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorTemperatureMiredsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorTemperatureMiredsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorTempPhysicalMaxMireds(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorTempPhysicalMaxMiredsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorTempPhysicalMaxMiredsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorTempPhysicalMaxMiredsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeColorTempPhysicalMinMireds(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeColorTempPhysicalMinMiredsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeColorTempPhysicalMinMiredsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeColorTempPhysicalMinMiredsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCompensationText(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeCompensationTextWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCompensationTextWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCompensationTextWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCoupleColorTempToLevelMinMireds(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeCoupleColorTempToLevelMinMiredsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCoupleColorTempToLevelMinMiredsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCoupleColorTempToLevelMinMiredsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCurrentHue(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeCurrentHueWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentHueWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentHueWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCurrentSaturation(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeCurrentSaturationWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentSaturationWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentSaturationWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCurrentX(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeCurrentXWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentXWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentXWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeCurrentY(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeCurrentYWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentYWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentYWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeDriftCompensation(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeDriftCompensationWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeDriftCompensationWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeDriftCompensationWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeEnhancedColorMode(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeEnhancedColorModeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeEnhancedColorModeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeEnhancedColorModeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeEnhancedCurrentHue(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeEnhancedCurrentHueWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeEnhancedCurrentHueWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeEnhancedCurrentHueWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeGeneratedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeNumberOfPrimaries(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeNumberOfPrimariesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeNumberOfPrimariesWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeNumberOfPrimariesWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeOptions(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeOptionsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeOptionsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeOptionsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary1Intensity(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary1IntensityWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary1IntensityWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary1IntensityWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary1X(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary1XWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary1XWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary1XWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary1Y(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary1YWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary1YWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary1YWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary2Intensity(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary2IntensityWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary2IntensityWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary2IntensityWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary2X(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary2XWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary2XWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary2XWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary2Y(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary2YWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary2YWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary2YWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary3Intensity(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary3IntensityWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary3IntensityWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary3IntensityWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary3X(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary3XWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary3XWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary3XWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary3Y(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary3YWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary3YWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary3YWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary4Intensity(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary4IntensityWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary4IntensityWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary4IntensityWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary4X(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary4XWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary4XWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary4XWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary4Y(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary4YWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary4YWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary4YWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary5Intensity(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary5IntensityWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary5IntensityWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary5IntensityWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary5X(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary5XWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary5XWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary5XWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary5Y(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary5YWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary5YWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary5YWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary6Intensity(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary6IntensityWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary6IntensityWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary6IntensityWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary6X(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary6XWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary6XWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary6XWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributePrimary6Y(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributePrimary6YWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePrimary6YWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributePrimary6YWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeRemainingTime(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeRemainingTimeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeRemainingTimeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeRemainingTimeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeStartUpColorTemperatureMireds(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeStartUpColorTemperatureMiredsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeStartUpColorTemperatureMiredsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeStartUpColorTemperatureMiredsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeWhitePointX(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeWhitePointXWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeWhitePointXWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeWhitePointXWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/readAttributeWhitePointY(completion:)
func (m_ MTRBaseClusterColorControl) ReadAttributeWhitePointYWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeWhitePointYWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeWhitePointYWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/stepColor(with:completion:)
func (m_ MTRBaseClusterColorControl) StepColorWithParamsCompletion(params IMTRColorControlClusterStepColorParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stepColorWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: StepColorWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/stepColorTemperature(with:completion:)
func (m_ MTRBaseClusterColorControl) StepColorTemperatureWithParamsCompletion(params IMTRColorControlClusterStepColorTemperatureParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stepColorTemperatureWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: StepColorTemperatureWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/stepHue(with:completion:)
func (m_ MTRBaseClusterColorControl) StepHueWithParamsCompletion(params IMTRColorControlClusterStepHueParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stepHueWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: StepHueWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/stepSaturation(with:completion:)
func (m_ MTRBaseClusterColorControl) StepSaturationWithParamsCompletion(params IMTRColorControlClusterStepSaturationParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stepSaturationWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: StepSaturationWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/stopMoveStep(with:completion:)
func (m_ MTRBaseClusterColorControl) StopMoveStepWithParamsCompletion(params IMTRColorControlClusterStopMoveStepParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopMoveStepWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: StopMoveStepWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorCapabilities(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorCapabilitiesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorCapabilitiesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorCapabilitiesWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorLoopActive(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorLoopActiveWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorLoopActiveWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorLoopActiveWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorLoopDirection(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorLoopDirectionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorLoopDirectionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorLoopDirectionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorLoopStartEnhancedHue(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorLoopStartEnhancedHueWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorLoopStartEnhancedHueWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorLoopStartEnhancedHueWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorLoopStoredEnhancedHue(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorLoopStoredEnhancedHueWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorLoopStoredEnhancedHueWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorLoopStoredEnhancedHueWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorLoopTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorLoopTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorLoopTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorLoopTimeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorMode(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorModeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorModeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorPointBIntensity(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorPointBIntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorPointBIntensityWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorPointBIntensityWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorPointBX(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorPointBXWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorPointBXWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorPointBXWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorPointBY(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorPointBYWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorPointBYWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorPointBYWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorPointGIntensity(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorPointGIntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorPointGIntensityWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorPointGIntensityWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorPointGX(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorPointGXWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorPointGXWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorPointGXWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorPointGY(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorPointGYWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorPointGYWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorPointGYWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorPointRIntensity(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorPointRIntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorPointRIntensityWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorPointRIntensityWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorPointRX(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorPointRXWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorPointRXWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorPointRXWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorPointRY(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorPointRYWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorPointRYWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorPointRYWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorTemperatureMireds(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorTemperatureMiredsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorTemperatureMiredsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorTemperatureMiredsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorTempPhysicalMaxMireds(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorTempPhysicalMaxMiredsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorTempPhysicalMaxMiredsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorTempPhysicalMaxMiredsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeColorTempPhysicalMinMireds(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeColorTempPhysicalMinMiredsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeColorTempPhysicalMinMiredsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeColorTempPhysicalMinMiredsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeCompensationText(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeCompensationTextWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCompensationTextWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCompensationTextWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeCoupleColorTempToLevelMinMireds(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeCoupleColorTempToLevelMinMiredsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCoupleColorTempToLevelMinMiredsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCoupleColorTempToLevelMinMiredsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeCurrentHue(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeCurrentHueWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentHueWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentHueWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeCurrentSaturation(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeCurrentSaturationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentSaturationWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentSaturationWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeCurrentX(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeCurrentXWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentXWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentXWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeCurrentY(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeCurrentYWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentYWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentYWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeDriftCompensation(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeDriftCompensationWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeDriftCompensationWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeDriftCompensationWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeEnhancedColorMode(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeEnhancedColorModeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeEnhancedColorModeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeEnhancedColorModeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeEnhancedCurrentHue(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeEnhancedCurrentHueWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeEnhancedCurrentHueWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeEnhancedCurrentHueWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeNumberOfPrimaries(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeNumberOfPrimariesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeNumberOfPrimariesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeNumberOfPrimariesWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeOptions(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeOptionsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeOptionsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeOptionsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary1Intensity(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary1IntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary1IntensityWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary1IntensityWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary1X(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary1XWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary1XWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary1XWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary1Y(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary1YWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary1YWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary1YWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary2Intensity(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary2IntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary2IntensityWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary2IntensityWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary2X(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary2XWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary2XWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary2XWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary2Y(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary2YWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary2YWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary2YWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary3Intensity(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary3IntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary3IntensityWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary3IntensityWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary3X(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary3XWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary3XWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary3XWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary3Y(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary3YWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary3YWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary3YWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary4Intensity(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary4IntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary4IntensityWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary4IntensityWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary4X(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary4XWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary4XWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary4XWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary4Y(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary4YWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary4YWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary4YWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary5Intensity(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary5IntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary5IntensityWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary5IntensityWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary5X(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary5XWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary5XWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary5XWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary5Y(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary5YWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary5YWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary5YWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary6Intensity(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary6IntensityWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary6IntensityWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary6IntensityWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary6X(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary6XWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary6XWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary6XWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributePrimary6Y(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributePrimary6YWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributePrimary6YWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributePrimary6YWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeRemainingTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeRemainingTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeRemainingTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeRemainingTimeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeStartUpColorTemperatureMireds(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeStartUpColorTemperatureMiredsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeStartUpColorTemperatureMiredsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeStartUpColorTemperatureMiredsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeWhitePointX(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeWhitePointXWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeWhitePointXWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeWhitePointXWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/subscribeAttributeWhitePointY(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterColorControl) SubscribeAttributeWhitePointYWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeWhitePointYWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeWhitePointYWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/writeAttributeOptions(withValue:completion:)
func (m_ MTRBaseClusterColorControl) WriteAttributeOptionsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOptionsWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeOptionsWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/writeAttributeOptions(withValue:params:completion:)
func (m_ MTRBaseClusterColorControl) WriteAttributeOptionsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeOptionsWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeOptionsWithValueParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/writeAttributeStartUpColorTemperatureMireds(withValue:completion:)
func (m_ MTRBaseClusterColorControl) WriteAttributeStartUpColorTemperatureMiredsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeStartUpColorTemperatureMiredsWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeStartUpColorTemperatureMiredsWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterColorControl/writeAttributeStartUpColorTemperatureMireds(withValue:params:completion:)
func (m_ MTRBaseClusterColorControl) WriteAttributeStartUpColorTemperatureMiredsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeStartUpColorTemperatureMiredsWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeStartUpColorTemperatureMiredsWithValueParamsCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterColorControl */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterColorControl */


