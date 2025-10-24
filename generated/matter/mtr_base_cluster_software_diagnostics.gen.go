// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterSoftwareDiagnostics */


/* debug [class_header]: Header for MTRBaseClusterSoftwareDiagnostics */
// The class instance for the [MTRBaseClusterSoftwareDiagnostics] class.
var (
	MTRBaseClusterSoftwareDiagnosticsClass     _MTRBaseClusterSoftwareDiagnosticsClass
	MTRBaseClusterSoftwareDiagnosticsClassOnce sync.Once
)

func getMTRBaseClusterSoftwareDiagnosticsClass() _MTRBaseClusterSoftwareDiagnosticsClass {
	MTRBaseClusterSoftwareDiagnosticsClassOnce.Do(func() {
		MTRBaseClusterSoftwareDiagnosticsClass = _MTRBaseClusterSoftwareDiagnosticsClass{objc.GetClass("MTRBaseClusterSoftwareDiagnostics")}
	})
	return MTRBaseClusterSoftwareDiagnosticsClass
}

type _MTRBaseClusterSoftwareDiagnosticsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterSoftwareDiagnostics */
// An interface definition for the [MTRBaseClusterSoftwareDiagnostics] class.
type IMTRBaseClusterSoftwareDiagnostics interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterSoftwareDiagnostics */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterSoftwareDiagnostics */
	// methods:
	ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer)
	ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentHeapFreeWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentHeapHighWatermarkWithCompletion(completion unsafe.Pointer)
	ReadAttributeCurrentHeapUsedWithCompletion(completion unsafe.Pointer)
	ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer)
	ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer)
	ReadAttributeThreadMetricsWithCompletion(completion unsafe.Pointer)
	ResetWatermarksWithCompletion(completion unsafe.Pointer)
	ResetWatermarksWithParamsCompletion(params IMTRSoftwareDiagnosticsClusterResetWatermarksParams, completion unsafe.Pointer)
	SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentHeapFreeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentHeapHighWatermarkWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeCurrentHeapUsedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
	SubscribeAttributeThreadMetricsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterSoftwareDiagnostics */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) Alloc() MTRBaseClusterSoftwareDiagnostics {
	rv := objc.Send[MTRBaseClusterSoftwareDiagnostics](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) New() MTRBaseClusterSoftwareDiagnostics {
	rv := objc.Send[MTRBaseClusterSoftwareDiagnostics](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterSoftwareDiagnostics) Init() MTRBaseClusterSoftwareDiagnostics {
	rv := objc.Send[MTRBaseClusterSoftwareDiagnostics](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterSoftwareDiagnostics) Autorelease() MTRBaseClusterSoftwareDiagnostics {
	rv := objc.Send[MTRBaseClusterSoftwareDiagnostics](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterSoftwareDiagnostics creates a new MTRBaseClusterSoftwareDiagnostics instance.
func NewMTRBaseClusterSoftwareDiagnostics() MTRBaseClusterSoftwareDiagnostics {
	return getMTRBaseClusterSoftwareDiagnosticsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterSoftwareDiagnostics */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics
type MTRBaseClusterSoftwareDiagnostics struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterSoftwareDiagnosticsFrom constructs a [MTRBaseClusterSoftwareDiagnostics] from an unsafe.Pointer.
func MTRBaseClusterSoftwareDiagnosticsFrom(ptr unsafe.Pointer) MTRBaseClusterSoftwareDiagnostics {
	return MTRBaseClusterSoftwareDiagnostics{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterSoftwareDiagnostics */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/init(device:endpointID:queue:)
func NewMTRBaseClusterSoftwareDiagnosticsWithDeviceEndpointIDQueue(device IMTRBaseDevice, endpointID objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer) MTRBaseClusterSoftwareDiagnostics {
	instance := getMTRBaseClusterSoftwareDiagnosticsClass().Alloc()
	rv := objc.Send[MTRBaseClusterSoftwareDiagnostics](instance.ID, objc.Sel("initWithDevice:endpointID:queue:"), device, endpointID, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterSoftwareDiagnosticsWithDeviceEndpointIDQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/init(device:endpoint:queue:)
func NewMTRBaseClusterSoftwareDiagnosticsWithDeviceEndpointQueue(device IMTRBaseDevice, endpoint uint16 /* not a class type */, queue unsafe.Pointer) MTRBaseClusterSoftwareDiagnostics {
	instance := getMTRBaseClusterSoftwareDiagnosticsClass().Alloc()
	rv := objc.Send[MTRBaseClusterSoftwareDiagnostics](instance.ID, objc.Sel("initWithDevice:endpoint:queue:"), device, endpoint, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseClusterSoftwareDiagnosticsWithDeviceEndpointQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterSoftwareDiagnostics */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeAcceptedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeAcceptedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAcceptedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAcceptedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeAttributeList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeAttributeList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeAttributeListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeAttributeListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeClusterRevision(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeClusterRevision(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeClusterRevisionWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeClusterRevisionWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeCurrentHeapFree(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeCurrentHeapFreeWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentHeapFreeWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentHeapFreeWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeCurrentHeapFree(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeCurrentHeapFreeWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentHeapFreeWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentHeapFreeWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeCurrentHeapHighWatermark(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeCurrentHeapHighWatermarkWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentHeapHighWatermarkWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentHeapHighWatermarkWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeCurrentHeapHighWatermark(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeCurrentHeapHighWatermarkWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentHeapHighWatermarkWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentHeapHighWatermarkWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeCurrentHeapUsed(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeCurrentHeapUsedWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentHeapUsedWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentHeapUsedWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeCurrentHeapUsed(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeCurrentHeapUsedWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeCurrentHeapUsedWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeCurrentHeapUsedWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeFeatureMap(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeFeatureMap(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeFeatureMapWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeFeatureMapWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeGeneratedCommandList(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeGeneratedCommandList(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeGeneratedCommandListWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeGeneratedCommandListWithClusterStateCacheEndpointQueueCompletion) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeThreadMetrics(withAttributeCache:endpoint:queue:completionHandler:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeThreadMetricsWithAttributeCacheEndpointQueueCompletionHandler(attributeCacheContainer IMTRAttributeCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeThreadMetricsWithAttributeCache:endpoint:queue:completionHandler:"), attributeCacheContainer, endpoint, queue, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeThreadMetricsWithAttributeCacheEndpointQueueCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeThreadMetrics(withClusterStateCache:endpoint:queue:completion:)
func (mc _MTRBaseClusterSoftwareDiagnosticsClass) ReadAttributeThreadMetricsWithClusterStateCacheEndpointQueueCompletion(clusterStateCacheContainer IMTRClusterStateCacheContainer, endpoint objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("readAttributeThreadMetricsWithClusterStateCache:endpoint:queue:completion:"), clusterStateCacheContainer, endpoint, queue, completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReadAttributeThreadMetricsWithClusterStateCacheEndpointQueueCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterSoftwareDiagnostics */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterSoftwareDiagnostics */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeAcceptedCommandList(completion:)
func (m_ MTRBaseClusterSoftwareDiagnostics) ReadAttributeAcceptedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAcceptedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAcceptedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeAttributeList(completion:)
func (m_ MTRBaseClusterSoftwareDiagnostics) ReadAttributeAttributeListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeAttributeListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeAttributeListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeClusterRevision(completion:)
func (m_ MTRBaseClusterSoftwareDiagnostics) ReadAttributeClusterRevisionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeClusterRevisionWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeClusterRevisionWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeCurrentHeapFree(completion:)
func (m_ MTRBaseClusterSoftwareDiagnostics) ReadAttributeCurrentHeapFreeWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentHeapFreeWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentHeapFreeWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeCurrentHeapHighWatermark(completion:)
func (m_ MTRBaseClusterSoftwareDiagnostics) ReadAttributeCurrentHeapHighWatermarkWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentHeapHighWatermarkWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentHeapHighWatermarkWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeCurrentHeapUsed(completion:)
func (m_ MTRBaseClusterSoftwareDiagnostics) ReadAttributeCurrentHeapUsedWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeCurrentHeapUsedWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeCurrentHeapUsedWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeFeatureMap(completion:)
func (m_ MTRBaseClusterSoftwareDiagnostics) ReadAttributeFeatureMapWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeFeatureMapWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeFeatureMapWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeGeneratedCommandList(completion:)
func (m_ MTRBaseClusterSoftwareDiagnostics) ReadAttributeGeneratedCommandListWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeGeneratedCommandListWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/readAttributeThreadMetrics(completion:)
func (m_ MTRBaseClusterSoftwareDiagnostics) ReadAttributeThreadMetricsWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributeThreadMetricsWithCompletion:"), completion)
}/* debug [instance_methods/method]: ReadAttributeThreadMetricsWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/resetWatermarks(completion:)
func (m_ MTRBaseClusterSoftwareDiagnostics) ResetWatermarksWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("resetWatermarksWithCompletion:"), completion)
}/* debug [instance_methods/method]: ResetWatermarksWithCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/resetWatermarks(with:completion:)
func (m_ MTRBaseClusterSoftwareDiagnostics) ResetWatermarksWithParamsCompletion(params IMTRSoftwareDiagnosticsClusterResetWatermarksParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("resetWatermarksWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: ResetWatermarksWithParamsCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/subscribeAttributeAcceptedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterSoftwareDiagnostics) SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAcceptedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAcceptedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/subscribeAttributeAttributeList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterSoftwareDiagnostics) SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeAttributeListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeAttributeListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterSoftwareDiagnostics) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/subscribeAttributeCurrentHeapFree(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterSoftwareDiagnostics) SubscribeAttributeCurrentHeapFreeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentHeapFreeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentHeapFreeWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/subscribeAttributeCurrentHeapHighWatermark(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterSoftwareDiagnostics) SubscribeAttributeCurrentHeapHighWatermarkWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentHeapHighWatermarkWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentHeapHighWatermarkWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/subscribeAttributeCurrentHeapUsed(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterSoftwareDiagnostics) SubscribeAttributeCurrentHeapUsedWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeCurrentHeapUsedWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeCurrentHeapUsedWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/subscribeAttributeFeatureMap(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterSoftwareDiagnostics) SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeFeatureMapWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeFeatureMapWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/subscribeAttributeGeneratedCommandList(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterSoftwareDiagnostics) SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeGeneratedCommandListWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeGeneratedCommandListWithParamsSubscriptionEstablishedReportHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSoftwareDiagnostics/subscribeAttributeThreadMetrics(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterSoftwareDiagnostics) SubscribeAttributeThreadMetricsWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeThreadMetricsWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeThreadMetricsWithParamsSubscriptionEstablishedReportHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterSoftwareDiagnostics */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterSoftwareDiagnostics */


