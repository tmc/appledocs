// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterContentLauncher */


/* debug [class_header]: Header for MTRBaseClusterContentLauncher */
// The class instance for the [MTRBaseClusterContentLauncher] class.
var (
	MTRBaseClusterContentLauncherClass     _MTRBaseClusterContentLauncherClass
	MTRBaseClusterContentLauncherClassOnce sync.Once
)

func getMTRBaseClusterContentLauncherClass() _MTRBaseClusterContentLauncherClass {
	MTRBaseClusterContentLauncherClassOnce.Do(func() {
		MTRBaseClusterContentLauncherClass = _MTRBaseClusterContentLauncherClass{objc.GetClass("MTRBaseClusterContentLauncher")}
	})
	return MTRBaseClusterContentLauncherClass
}

type _MTRBaseClusterContentLauncherClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterContentLauncher */
// An interface definition for the [MTRBaseClusterContentLauncher] class.
type IMTRBaseClusterContentLauncher interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterContentLauncher */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterContentLauncher */
	// methods:
	LaunchContentWithParamsCompletion(params objc.IObject /* cross-framework: MTRContentLauncherClusterLaunchContentParams */, completion unsafe.Pointer)
	LaunchURLWithParamsCompletion(params objc.IObject /* cross-framework: MTRContentLauncherClusterLaunchURLParams */, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAcceptHeaderWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeSupportedStreamingProtocolsWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAcceptHeaderWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeSupportedStreamingProtocolsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	WriteAttributeSupportedStreamingProtocolsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer)
	WriteAttributeSupportedStreamingProtocolsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterContentLauncher */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterContentLauncherClass) Alloc() MTRBaseClusterContentLauncher {
	rv := objc.Send[MTRBaseClusterContentLauncher](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterContentLauncherClass) New() MTRBaseClusterContentLauncher {
	rv := objc.Send[MTRBaseClusterContentLauncher](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterContentLauncher) Init() MTRBaseClusterContentLauncher {
	rv := objc.Send[MTRBaseClusterContentLauncher](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterContentLauncher) Autorelease() MTRBaseClusterContentLauncher {
	rv := objc.Send[MTRBaseClusterContentLauncher](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterContentLauncher creates a new MTRBaseClusterContentLauncher instance.
func NewMTRBaseClusterContentLauncher() MTRBaseClusterContentLauncher {
	return getMTRBaseClusterContentLauncherClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterContentLauncher */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher
type MTRBaseClusterContentLauncher struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterContentLauncherFrom constructs a [MTRBaseClusterContentLauncher] from an unsafe.Pointer.
func MTRBaseClusterContentLauncherFrom(ptr unsafe.Pointer) MTRBaseClusterContentLauncher {
	return MTRBaseClusterContentLauncher{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterContentLauncher */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/init(device:endpointID:queue:)
func NewMTRBaseClusterContentLauncherWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterContentLauncher {
	instance := getMTRBaseClusterContentLauncherClass().Alloc()
	rv := objc.Send[MTRBaseClusterContentLauncher](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterContentLauncherWithDeviceEndpointIDQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/init(device:endpoint:queue:)
func NewMTRBaseClusterContentLauncherWithDeviceEndpointQueue(device IMTRBaseDevice, endpoint uint16 /* not a class type */, queue unsafe.Pointer) MTRBaseClusterContentLauncher {
	instance := getMTRBaseClusterContentLauncherClass().Alloc()
	rv := objc.Send[MTRBaseClusterContentLauncher](instance.ID, objc.Sel("initWithDevice:endpoint:queue:"), device, endpoint, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterContentLauncherWithDeviceEndpointQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterContentLauncher */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeAcceptedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeAcceptHeader(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeAcceptHeaderWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptHeaderWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptHeaderWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeAcceptHeader(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeAcceptHeaderWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptHeaderWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptHeaderWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeAttributeList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeClusterRevision(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeFeatureMap(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeGeneratedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeSupportedStreamingProtocols(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeSupportedStreamingProtocolsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedStreamingProtocolsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSupportedStreamingProtocolsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeSupportedStreamingProtocols(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterContentLauncherClass) ReadAttributeSupportedStreamingProtocolsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeSupportedStreamingProtocolsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeSupportedStreamingProtocolsWithClusterStateCacheEndpointQueueCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterContentLauncher */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterContentLauncher */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/launchContent(with:completion:)
func (m_ MTRBaseClusterContentLauncher) LaunchContentWithParamsCompletion(params objc.IObject /* cross-framework: MTRContentLauncherClusterLaunchContentParams */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("launchContentWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: LaunchContentWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/launchURL(with:completion:)
func (m_ MTRBaseClusterContentLauncher) LaunchURLWithParamsCompletion(params objc.IObject /* cross-framework: MTRContentLauncherClusterLaunchURLParams */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("launchURLWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: LaunchURLWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterContentLauncher) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeAcceptHeader(completion:)
func (m_ MTRBaseClusterContentLauncher) ReadAttributeAcceptHeaderWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptHeaderWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAcceptHeaderWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterContentLauncher) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAttributeListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterContentLauncher) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeClusterRevisionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterContentLauncher) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterContentLauncher) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeGeneratedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/readAttributeSupportedStreamingProtocols(completion:)
func (m_ MTRBaseClusterContentLauncher) ReadAttributeSupportedStreamingProtocolsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeSupportedStreamingProtocolsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeSupportedStreamingProtocolsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterContentLauncher) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/subscribeAttributeAcceptHeader(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterContentLauncher) SubscribeAttributeAcceptHeaderWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptHeaderWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAcceptHeaderWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterContentLauncher) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterContentLauncher) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterContentLauncher) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterContentLauncher) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/subscribeAttributeSupportedStreamingProtocols(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterContentLauncher) SubscribeAttributeSupportedStreamingProtocolsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeSupportedStreamingProtocolsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeSupportedStreamingProtocolsWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/writeAttributeSupportedStreamingProtocols(withValue:completion:)
func (m_ MTRBaseClusterContentLauncher) WriteAttributeSupportedStreamingProtocolsWithValueCompletion(value objc.IObject /* cross-framework: NSNumber */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSupportedStreamingProtocolsWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeSupportedStreamingProtocolsWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher/writeAttributeSupportedStreamingProtocols(withValue:params:completion:)
func (m_ MTRBaseClusterContentLauncher) WriteAttributeSupportedStreamingProtocolsWithValueParamsCompletion(value objc.IObject /* cross-framework: NSNumber */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeSupportedStreamingProtocolsWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeSupportedStreamingProtocolsWithValueParamsCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterContentLauncher */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterContentLauncher */


