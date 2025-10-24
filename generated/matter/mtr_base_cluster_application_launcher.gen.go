// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterApplicationLauncher */


/* debug [class_header]: Header for MTRBaseClusterApplicationLauncher */
// The class instance for the [MTRBaseClusterApplicationLauncher] class.
var (
	MTRBaseClusterApplicationLauncherClass     _MTRBaseClusterApplicationLauncherClass
	MTRBaseClusterApplicationLauncherClassOnce sync.Once
)

func getMTRBaseClusterApplicationLauncherClass() _MTRBaseClusterApplicationLauncherClass {
	MTRBaseClusterApplicationLauncherClassOnce.Do(func() {
		MTRBaseClusterApplicationLauncherClass = _MTRBaseClusterApplicationLauncherClass{objc.GetClass("MTRBaseClusterApplicationLauncher")}
	})
	return MTRBaseClusterApplicationLauncherClass
}

type _MTRBaseClusterApplicationLauncherClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterApplicationLauncher */
// An interface definition for the [MTRBaseClusterApplicationLauncher] class.
type IMTRBaseClusterApplicationLauncher interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterApplicationLauncher */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterApplicationLauncher */
	// methods:
	HideAppWithCompletion(completion unsafe.Pointer)
	HideAppWithParamsCompletion(params IMTRApplicationLauncherClusterHideAppParams, completion unsafe.Pointer)
	LaunchAppWithCompletion(completion unsafe.Pointer)
	LaunchAppWithParamsCompletion(params IMTRApplicationLauncherClusterLaunchAppParams, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeCatalogListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentAppWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	StopAppWithCompletion(completion unsafe.Pointer)
	StopAppWithParamsCompletion(params IMTRApplicationLauncherClusterStopAppParams, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCatalogListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentAppWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	WriteAttributeCurrentAppWithValueCompletion(value IMTRApplicationLauncherClusterApplicationEPStruct, completion unsafe.Pointer)
	WriteAttributeCurrentAppWithValueParamsCompletion(value IMTRApplicationLauncherClusterApplicationEPStruct, params IMTRWriteParams, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterApplicationLauncher */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterApplicationLauncherClass) Alloc() MTRBaseClusterApplicationLauncher {
	rv := objc.Send[MTRBaseClusterApplicationLauncher](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterApplicationLauncherClass) New() MTRBaseClusterApplicationLauncher {
	rv := objc.Send[MTRBaseClusterApplicationLauncher](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterApplicationLauncher) Init() MTRBaseClusterApplicationLauncher {
	rv := objc.Send[MTRBaseClusterApplicationLauncher](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterApplicationLauncher) Autorelease() MTRBaseClusterApplicationLauncher {
	rv := objc.Send[MTRBaseClusterApplicationLauncher](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterApplicationLauncher creates a new MTRBaseClusterApplicationLauncher instance.
func NewMTRBaseClusterApplicationLauncher() MTRBaseClusterApplicationLauncher {
	return getMTRBaseClusterApplicationLauncherClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterApplicationLauncher */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher
type MTRBaseClusterApplicationLauncher struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterApplicationLauncherFrom constructs a [MTRBaseClusterApplicationLauncher] from an unsafe.Pointer.
func MTRBaseClusterApplicationLauncherFrom(ptr unsafe.Pointer) MTRBaseClusterApplicationLauncher {
	return MTRBaseClusterApplicationLauncher{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterApplicationLauncher */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/init(device:endpointID:queue:)
func NewMTRBaseClusterApplicationLauncherWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterApplicationLauncher {
	instance := getMTRBaseClusterApplicationLauncherClass().Alloc()
	rv := objc.Send[MTRBaseClusterApplicationLauncher](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterApplicationLauncherWithDeviceEndpointIDQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/init(device:endpoint:queue:)
func NewMTRBaseClusterApplicationLauncherWithDeviceEndpointQueue(device IMTRBaseDevice, endpoint uint16 /* not a class type */, queue unsafe.Pointer) MTRBaseClusterApplicationLauncher {
	instance := getMTRBaseClusterApplicationLauncherClass().Alloc()
	rv := objc.Send[MTRBaseClusterApplicationLauncher](instance.ID, objc.Sel("initWithDevice:endpoint:queue:"), device, endpoint, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterApplicationLauncherWithDeviceEndpointQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterApplicationLauncher */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeAcceptedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeAttributeList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeCatalogList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeCatalogListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCatalogListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCatalogListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeCatalogList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeCatalogListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCatalogListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCatalogListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeClusterRevision(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeCurrentApp(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeCurrentAppWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentAppWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentAppWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeCurrentApp(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeCurrentAppWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentAppWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentAppWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeFeatureMap(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeGeneratedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterApplicationLauncherClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterApplicationLauncher */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterApplicationLauncher */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/hideApp(completion:)
func (m_ MTRBaseClusterApplicationLauncher) HideAppWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("hideAppWithCompletion:"), completion)
}/* debug [instance_methods/method]: HideAppWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/hideApp(with:completion:)
func (m_ MTRBaseClusterApplicationLauncher) HideAppWithParamsCompletion(params IMTRApplicationLauncherClusterHideAppParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("hideAppWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: HideAppWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/launchApp(completion:)
func (m_ MTRBaseClusterApplicationLauncher) LaunchAppWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("launchAppWithCompletion:"), completion)
}/* debug [instance_methods/method]: LaunchAppWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/launchApp(with:completion:)
func (m_ MTRBaseClusterApplicationLauncher) LaunchAppWithParamsCompletion(params IMTRApplicationLauncherClusterLaunchAppParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("launchAppWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: LaunchAppWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterApplicationLauncher) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterApplicationLauncher) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAttributeListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeCatalogList(completion:)
func (m_ MTRBaseClusterApplicationLauncher) ReadAttributeCatalogListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCatalogListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCatalogListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterApplicationLauncher) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeClusterRevisionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeCurrentApp(completion:)
func (m_ MTRBaseClusterApplicationLauncher) ReadAttributeCurrentAppWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentAppWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentAppWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterApplicationLauncher) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterApplicationLauncher) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeGeneratedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/stopApp(completion:)
func (m_ MTRBaseClusterApplicationLauncher) StopAppWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopAppWithCompletion:"), completion)
}/* debug [instance_methods/method]: StopAppWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/stopApp(with:completion:)
func (m_ MTRBaseClusterApplicationLauncher) StopAppWithParamsCompletion(params IMTRApplicationLauncherClusterStopAppParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopAppWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: StopAppWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterApplicationLauncher) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterApplicationLauncher) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/subscribeAttributeCatalogList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterApplicationLauncher) SubscribeAttributeCatalogListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCatalogListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCatalogListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterApplicationLauncher) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/subscribeAttributeCurrentApp(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterApplicationLauncher) SubscribeAttributeCurrentAppWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentAppWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentAppWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterApplicationLauncher) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterApplicationLauncher) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/writeAttributeCurrentApp(withValue:completion:)
func (m_ MTRBaseClusterApplicationLauncher) WriteAttributeCurrentAppWithValueCompletion(value IMTRApplicationLauncherClusterApplicationEPStruct, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeCurrentAppWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeCurrentAppWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher/writeAttributeCurrentApp(withValue:params:completion:)
func (m_ MTRBaseClusterApplicationLauncher) WriteAttributeCurrentAppWithValueParamsCompletion(value IMTRApplicationLauncherClusterApplicationEPStruct, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeCurrentAppWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeCurrentAppWithValueParamsCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterApplicationLauncher */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterApplicationLauncher */


