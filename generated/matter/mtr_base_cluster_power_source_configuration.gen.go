// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterPowerSourceConfiguration */


/* debug [class_header]: Header for MTRBaseClusterPowerSourceConfiguration */
// The class instance for the [MTRBaseClusterPowerSourceConfiguration] class.
var (
	MTRBaseClusterPowerSourceConfigurationClass     _MTRBaseClusterPowerSourceConfigurationClass
	MTRBaseClusterPowerSourceConfigurationClassOnce sync.Once
)

func getMTRBaseClusterPowerSourceConfigurationClass() _MTRBaseClusterPowerSourceConfigurationClass {
	MTRBaseClusterPowerSourceConfigurationClassOnce.Do(func() {
		MTRBaseClusterPowerSourceConfigurationClass = _MTRBaseClusterPowerSourceConfigurationClass{objc.GetClass("MTRBaseClusterPowerSourceConfiguration")}
	})
	return MTRBaseClusterPowerSourceConfigurationClass
}

type _MTRBaseClusterPowerSourceConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterPowerSourceConfiguration */
// An interface definition for the [MTRBaseClusterPowerSourceConfiguration] class.
type IMTRBaseClusterPowerSourceConfiguration interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterPowerSourceConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterPowerSourceConfiguration */
	// methods:
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeSourcesWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSourcesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterPowerSourceConfiguration */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterPowerSourceConfigurationClass) Alloc() MTRBaseClusterPowerSourceConfiguration {
	rv := objc.Send[MTRBaseClusterPowerSourceConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterPowerSourceConfigurationClass) New() MTRBaseClusterPowerSourceConfiguration {
	rv := objc.Send[MTRBaseClusterPowerSourceConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterPowerSourceConfiguration) Init() MTRBaseClusterPowerSourceConfiguration {
	rv := objc.Send[MTRBaseClusterPowerSourceConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterPowerSourceConfiguration) Autorelease() MTRBaseClusterPowerSourceConfiguration {
	rv := objc.Send[MTRBaseClusterPowerSourceConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterPowerSourceConfiguration creates a new MTRBaseClusterPowerSourceConfiguration instance.
func NewMTRBaseClusterPowerSourceConfiguration() MTRBaseClusterPowerSourceConfiguration {
	return getMTRBaseClusterPowerSourceConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterPowerSourceConfiguration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration
type MTRBaseClusterPowerSourceConfiguration struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterPowerSourceConfigurationFrom constructs a [MTRBaseClusterPowerSourceConfiguration] from an unsafe.Pointer.
func MTRBaseClusterPowerSourceConfigurationFrom(ptr unsafe.Pointer) MTRBaseClusterPowerSourceConfiguration {
	return MTRBaseClusterPowerSourceConfiguration{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterPowerSourceConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/init(device:endpointID:queue:)
func NewMTRBaseClusterPowerSourceConfigurationWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterPowerSourceConfiguration {
	instance := getMTRBaseClusterPowerSourceConfigurationClass().Alloc()
	rv := objc.Send[MTRBaseClusterPowerSourceConfiguration](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterPowerSourceConfigurationWithDeviceEndpointIDQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/init(device:endpoint:queue:)
func NewMTRBaseClusterPowerSourceConfigurationWithDeviceEndpointQueue(device IMTRBaseDevice, endpoint uint16 /* not a class type */, queue unsafe.Pointer) MTRBaseClusterPowerSourceConfiguration {
	instance := getMTRBaseClusterPowerSourceConfigurationClass().Alloc()
	rv := objc.Send[MTRBaseClusterPowerSourceConfiguration](instance.ID, objc.Sel("initWithDevice:endpoint:queue:"), device, endpoint, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterPowerSourceConfigurationWithDeviceEndpointQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterPowerSourceConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeAcceptedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterPowerSourceConfigurationClass) ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterPowerSourceConfigurationClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeAttributeList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterPowerSourceConfigurationClass) ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterPowerSourceConfigurationClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeClusterRevision(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterPowerSourceConfigurationClass) ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterPowerSourceConfigurationClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeFeatureMap(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterPowerSourceConfigurationClass) ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterPowerSourceConfigurationClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeGeneratedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterPowerSourceConfigurationClass) ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterPowerSourceConfigurationClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeSources(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterPowerSourceConfigurationClass) ReadAttributeSourcesWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSourcesWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSourcesWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeSources(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterPowerSourceConfigurationClass) ReadAttributeSourcesWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSourcesWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSourcesWithClusterStateCacheEndpointQueueCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterPowerSourceConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterPowerSourceConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterPowerSourceConfiguration) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterPowerSourceConfiguration) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAttributeListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterPowerSourceConfiguration) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeClusterRevisionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterPowerSourceConfiguration) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterPowerSourceConfiguration) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeGeneratedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/readAttributeSources(completion:)
func (m_ MTRBaseClusterPowerSourceConfiguration) ReadAttributeSourcesWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSourcesWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeSourcesWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterPowerSourceConfiguration) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterPowerSourceConfiguration) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterPowerSourceConfiguration) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterPowerSourceConfiguration) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterPowerSourceConfiguration) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPowerSourceConfiguration/subscribeAttributeSources(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterPowerSourceConfiguration) SubscribeAttributeSourcesWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSourcesWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeSourcesWithParamsSubscriptionEstablishedReportHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterPowerSourceConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterPowerSourceConfiguration */


