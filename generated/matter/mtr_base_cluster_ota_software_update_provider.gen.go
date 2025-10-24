// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterOTASoftwareUpdateProvider */


/* debug [class_header]: Header for MTRBaseClusterOTASoftwareUpdateProvider */
// The class instance for the [MTRBaseClusterOTASoftwareUpdateProvider] class.
var (
	MTRBaseClusterOTASoftwareUpdateProviderClass     _MTRBaseClusterOTASoftwareUpdateProviderClass
	MTRBaseClusterOTASoftwareUpdateProviderClassOnce sync.Once
)

func getMTRBaseClusterOTASoftwareUpdateProviderClass() _MTRBaseClusterOTASoftwareUpdateProviderClass {
	MTRBaseClusterOTASoftwareUpdateProviderClassOnce.Do(func() {
		MTRBaseClusterOTASoftwareUpdateProviderClass = _MTRBaseClusterOTASoftwareUpdateProviderClass{objc.GetClass("MTRBaseClusterOTASoftwareUpdateProvider")}
	})
	return MTRBaseClusterOTASoftwareUpdateProviderClass
}

type _MTRBaseClusterOTASoftwareUpdateProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterOTASoftwareUpdateProvider */
// An interface definition for the [MTRBaseClusterOTASoftwareUpdateProvider] class.
type IMTRBaseClusterOTASoftwareUpdateProvider interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterOTASoftwareUpdateProvider */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterOTASoftwareUpdateProvider */
	// methods:
	ApplyUpdateRequestWithParamsCompletion(params IMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams, completion unsafe.Pointer)
	NotifyUpdateAppliedWithParamsCompletion(params IMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams, completion unsafe.Pointer)
	QueryImageWithParamsCompletion(params IMTROTASoftwareUpdateProviderClusterQueryImageParams, completion unsafe.Pointer)
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterOTASoftwareUpdateProvider */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterOTASoftwareUpdateProviderClass) Alloc() MTRBaseClusterOTASoftwareUpdateProvider {
	rv := objc.Send[MTRBaseClusterOTASoftwareUpdateProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterOTASoftwareUpdateProviderClass) New() MTRBaseClusterOTASoftwareUpdateProvider {
	rv := objc.Send[MTRBaseClusterOTASoftwareUpdateProvider](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) Init() MTRBaseClusterOTASoftwareUpdateProvider {
	rv := objc.Send[MTRBaseClusterOTASoftwareUpdateProvider](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) Autorelease() MTRBaseClusterOTASoftwareUpdateProvider {
	rv := objc.Send[MTRBaseClusterOTASoftwareUpdateProvider](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterOTASoftwareUpdateProvider creates a new MTRBaseClusterOTASoftwareUpdateProvider instance.
func NewMTRBaseClusterOTASoftwareUpdateProvider() MTRBaseClusterOTASoftwareUpdateProvider {
	return getMTRBaseClusterOTASoftwareUpdateProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterOTASoftwareUpdateProvider */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit
type MTRBaseClusterOTASoftwareUpdateProvider struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterOTASoftwareUpdateProviderFrom constructs a [MTRBaseClusterOTASoftwareUpdateProvider] from an unsafe.Pointer.
func MTRBaseClusterOTASoftwareUpdateProviderFrom(ptr unsafe.Pointer) MTRBaseClusterOTASoftwareUpdateProvider {
	return MTRBaseClusterOTASoftwareUpdateProvider{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterOTASoftwareUpdateProvider */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/init(device:endpointID:queue:)
func NewMTRBaseClusterOTASoftwareUpdateProviderWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterOTASoftwareUpdateProvider {
	instance := getMTRBaseClusterOTASoftwareUpdateProviderClass().Alloc()
	rv := objc.Send[MTRBaseClusterOTASoftwareUpdateProvider](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterOTASoftwareUpdateProviderWithDeviceEndpointIDQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterOTASoftwareUpdateProvider */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateProviderClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateProviderClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateProviderClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateProviderClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterOTASoftwareUpdateProviderClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterOTASoftwareUpdateProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterOTASoftwareUpdateProvider */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/applyUpdateRequest(with:completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) ApplyUpdateRequestWithParamsCompletion(params IMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("applyUpdateRequestWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: ApplyUpdateRequestWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/notifyUpdateApplied(with:completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) NotifyUpdateAppliedWithParamsCompletion(params IMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("notifyUpdateAppliedWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: NotifyUpdateAppliedWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/queryImage(with:completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) QueryImageWithParamsCompletion(params IMTROTASoftwareUpdateProviderClusterQueryImageParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("queryImageWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: QueryImageWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAttributeListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeClusterRevisionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeGeneratedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOTASoftwareUpdateProvider-8bnit/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterOTASoftwareUpdateProvider) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterOTASoftwareUpdateProvider */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterOTASoftwareUpdateProvider */


