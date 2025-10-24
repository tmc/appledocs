// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterOTASoftwareUpdateRequestor */


/* debug [class_header]: Header for MTRBaseClusterOTASoftwareUpdateRequestor */
// The class instance for the [MTRBaseClusterOTASoftwareUpdateRequestor] class.
var (
	MTRBaseClusterOTASoftwareUpdateRequestorClass     _MTRBaseClusterOTASoftwareUpdateRequestorClass
	MTRBaseClusterOTASoftwareUpdateRequestorClassOnce sync.Once
)

func getMTRBaseClusterOTASoftwareUpdateRequestorClass() _MTRBaseClusterOTASoftwareUpdateRequestorClass {
	MTRBaseClusterOTASoftwareUpdateRequestorClassOnce.Do(func() {
		MTRBaseClusterOTASoftwareUpdateRequestorClass = _MTRBaseClusterOTASoftwareUpdateRequestorClass{objc.GetClass("MTRBaseClusterOTASoftwareUpdateRequestor")}
	})
	return MTRBaseClusterOTASoftwareUpdateRequestorClass
}

type _MTRBaseClusterOTASoftwareUpdateRequestorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterOTASoftwareUpdateRequestor */
// An interface definition for the [MTRBaseClusterOTASoftwareUpdateRequestor] class.
type IMTRBaseClusterOTASoftwareUpdateRequestor interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterOTASoftwareUpdateRequestor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterOTASoftwareUpdateRequestor */
	// methods:
	AnnounceOTAProviderWithParamsCompletion(params IMTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeDefaultOTAProvidersWithParamsCompletion(params IMTRReadParams, completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeUpdatePossibleWithCompletion(completion unsafe.Pointer)
	ReadAttributeUpdateStateWithCompletion(completion unsafe.Pointer)
	ReadAttributeUpdateStateProgressWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeDefaultOTAProvidersWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeUpdatePossibleWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeUpdateStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeUpdateStateProgressWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	WriteAttributeDefaultOTAProvidersWithValueCompletion(value objc.IObject /* cross-framework: NSArray */, completion unsafe.Pointer)
	WriteAttributeDefaultOTAProvidersWithValueParamsCompletion(value objc.IObject /* cross-framework: NSArray */, params IMTRWriteParams, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterOTASoftwareUpdateRequestor */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterOTASoftwareUpdateRequestorClass) Alloc() MTRBaseClusterOTASoftwareUpdateRequestor {
	rv := objc.Send[MTRBaseClusterOTASoftwareUpdateRequestor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterOTASoftwareUpdateRequestorClass) New() MTRBaseClusterOTASoftwareUpdateRequestor {
	rv := objc.Send[MTRBaseClusterOTASoftwareUpdateRequestor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) Init() MTRBaseClusterOTASoftwareUpdateRequestor {
	rv := objc.Send[MTRBaseClusterOTASoftwareUpdateRequestor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) Autorelease() MTRBaseClusterOTASoftwareUpdateRequestor {
	rv := objc.Send[MTRBaseClusterOTASoftwareUpdateRequestor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterOTASoftwareUpdateRequestor creates a new MTRBaseClusterOTASoftwareUpdateRequestor instance.
func NewMTRBaseClusterOTASoftwareUpdateRequestor() MTRBaseClusterOTASoftwareUpdateRequestor {
	return getMTRBaseClusterOTASoftwareUpdateRequestorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterOTASoftwareUpdateRequestor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb
type MTRBaseClusterOTASoftwareUpdateRequestor struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterOTASoftwareUpdateRequestorFrom constructs a [MTRBaseClusterOTASoftwareUpdateRequestor] from an unsafe.Pointer.
func MTRBaseClusterOTASoftwareUpdateRequestorFrom(ptr unsafe.Pointer) MTRBaseClusterOTASoftwareUpdateRequestor {
	return MTRBaseClusterOTASoftwareUpdateRequestor{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterOTASoftwareUpdateRequestor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/init(device:endpointID:queue:)
func NewMTRBaseClusterOTASoftwareUpdateRequestorWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterOTASoftwareUpdateRequestor {
	instance := getMTRBaseClusterOTASoftwareUpdateRequestorClass().Alloc()
	rv := objc.Send[MTRBaseClusterOTASoftwareUpdateRequestor](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterOTASoftwareUpdateRequestorWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterOTASoftwareUpdateRequestor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateRequestorClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateRequestorClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateRequestorClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeDefaultOTAProviders(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateRequestorClass) ReadAttributeDefaultOTAProvidersWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeDefaultOTAProvidersWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeDefaultOTAProvidersWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateRequestorClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateRequestorClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeUpdatePossible(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateRequestorClass) ReadAttributeUpdatePossibleWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUpdatePossibleWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUpdatePossibleWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeUpdateState(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateRequestorClass) ReadAttributeUpdateStateWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUpdateStateWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUpdateStateWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeUpdateStateProgress(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateRequestorClass) ReadAttributeUpdateStateProgressWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeUpdateStateProgressWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeUpdateStateProgressWithClusterStateCacheEndpointQueueCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterOTASoftwareUpdateRequestor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterOTASoftwareUpdateRequestor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/announceOTAProvider(with:completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) AnnounceOTAProviderWithParamsCompletion(params IMTROTASoftwareUpdateRequestorClusterAnnounceOTAProviderParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("announceOTAProviderWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: AnnounceOTAProviderWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAttributeListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeClusterRevisionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeDefaultOTAProviders(with:completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) ReadAttributeDefaultOTAProvidersWithParamsCompletion(params IMTRReadParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeDefaultOTAProvidersWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: ReadAttributeDefaultOTAProvidersWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeGeneratedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeUpdatePossible(completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) ReadAttributeUpdatePossibleWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeUpdatePossibleWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeUpdatePossibleWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeUpdateState(completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) ReadAttributeUpdateStateWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeUpdateStateWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeUpdateStateWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/readAttributeUpdateStateProgress(completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) ReadAttributeUpdateStateProgressWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeUpdateStateProgressWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeUpdateStateProgressWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/subscribeAttributeDefaultOTAProviders(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) SubscribeAttributeDefaultOTAProvidersWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeDefaultOTAProvidersWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeDefaultOTAProvidersWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/subscribeAttributeUpdatePossible(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) SubscribeAttributeUpdatePossibleWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUpdatePossibleWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeUpdatePossibleWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/subscribeAttributeUpdateState(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) SubscribeAttributeUpdateStateWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUpdateStateWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeUpdateStateWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/subscribeAttributeUpdateStateProgress(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) SubscribeAttributeUpdateStateProgressWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUpdateStateProgressWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeUpdateStateProgressWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/writeAttributeDefaultOTAProviders(withValue:completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) WriteAttributeDefaultOTAProvidersWithValueCompletion(value objc.IObject /* cross-framework: NSArray */, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeDefaultOTAProvidersWithValue:completion:"), value, completion)
}/* debug [instance_methods/method]: WriteAttributeDefaultOTAProvidersWithValueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateRequestor-9n6nb/writeAttributeDefaultOTAProviders(withValue:params:completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateRequestor) WriteAttributeDefaultOTAProvidersWithValueParamsCompletion(value objc.IObject /* cross-framework: NSArray */, params IMTRWriteParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeDefaultOTAProvidersWithValue:params:completion:"), value, params, completion)
}/* debug [instance_methods/method]: WriteAttributeDefaultOTAProvidersWithValueParamsCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterOTASoftwareUpdateRequestor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterOTASoftwareUpdateRequestor */


