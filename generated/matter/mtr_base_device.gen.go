// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRBaseDevice */


/* debug [class_header]: Header for MTRBaseDevice */
// The class instance for the [MTRBaseDevice] class.
var (
	MTRBaseDeviceClass     _MTRBaseDeviceClass
	MTRBaseDeviceClassOnce sync.Once
)

func getMTRBaseDeviceClass() _MTRBaseDeviceClass {
	MTRBaseDeviceClassOnce.Do(func() {
		MTRBaseDeviceClass = _MTRBaseDeviceClass{objc.GetClass("MTRBaseDevice")}
	})
	return MTRBaseDeviceClass
}

type _MTRBaseDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseDevice */
// An interface definition for the [MTRBaseDevice] class.
type IMTRBaseDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRBaseDevice */
	// properties:
	SessionTransportType() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseDevice */
	// methods:
	DeregisterReportHandlersWithQueueCompletion(queue unsafe.Pointer, completion unsafe.Pointer)
	DownloadLogOfTypeTimeoutQueueCompletion(type_ unsafe.Pointer, timeout float64, queue unsafe.Pointer, completion unsafe.Pointer)
	InvokeCommandWithEndpointIDClusterIDCommandIDCommandFieldsTimedInvokeTimeoutQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, commandID objc.IObject /* cross-framework: NSNumber */, commandFields objc.IObject, timeoutMs objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer)
	OpenCommissioningWindowWithDiscriminatorDurationQueueCompletion(discriminator objc.IObject /* cross-framework: NSNumber */, duration objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer)
	OpenCommissioningWindowWithSetupPasscodeDiscriminatorDurationQueueCompletion(setupPasscode objc.IObject /* cross-framework: NSNumber */, discriminator objc.IObject /* cross-framework: NSNumber */, duration objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributePathsEventPathsParamsQueueCompletion(attributePaths []MTRAttributeRequestPath, eventPaths []MTREventRequestPath, params IMTRReadParams, queue unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributesWithEndpointIDClusterIDAttributeIDParamsQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */, params IMTRReadParams, queue unsafe.Pointer, completion unsafe.Pointer)
	ReadEventsWithEndpointIDClusterIDEventIDParamsQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, eventID objc.IObject /* cross-framework: NSNumber */, params IMTRReadParams, queue unsafe.Pointer, completion unsafe.Pointer)
	SubscribeToAttributePathsEventPathsParamsQueueReportHandlerSubscriptionEstablishedResubscriptionScheduled(attributePaths []MTRAttributeRequestPath, eventPaths []MTREventRequestPath, params IMTRSubscribeParams, queue unsafe.Pointer, reportHandler unsafe.Pointer, subscriptionEstablished unsafe.Pointer, resubscriptionScheduled unsafe.Pointer)
	SubscribeWithQueueParamsClusterStateCacheContainerAttributeReportHandlerEventReportHandlerErrorHandlerSubscriptionEstablishedResubscriptionScheduled(queue unsafe.Pointer, params IMTRSubscribeParams, clusterStateCacheContainer IMTRClusterStateCacheContainer, attributeReportHandler unsafe.Pointer, eventReportHandler unsafe.Pointer, errorHandler unsafe.Pointer, subscriptionEstablished unsafe.Pointer, resubscriptionScheduled unsafe.Pointer)
	SubscribeToAttributesWithEndpointIDClusterIDAttributeIDParamsQueueReportHandlerSubscriptionEstablished(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */, params IMTRSubscribeParams, queue unsafe.Pointer, reportHandler unsafe.Pointer, subscriptionEstablished unsafe.Pointer)
	SubscribeToEventsWithEndpointIDClusterIDEventIDParamsQueueReportHandlerSubscriptionEstablished(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, eventID objc.IObject /* cross-framework: NSNumber */, params IMTRSubscribeParams, queue unsafe.Pointer, reportHandler unsafe.Pointer, subscriptionEstablished unsafe.Pointer)
	WriteAttributeWithEndpointIDClusterIDAttributeIDValueTimedWriteTimeoutQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */, value objc.IObject, timeoutMs objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseDevice */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseDeviceClass) Alloc() MTRBaseDevice {
	rv := objc.Send[MTRBaseDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseDeviceClass) New() MTRBaseDevice {
	rv := objc.Send[MTRBaseDevice](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseDevice) Init() MTRBaseDevice {
	rv := objc.Send[MTRBaseDevice](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseDevice) Autorelease() MTRBaseDevice {
	rv := objc.Send[MTRBaseDevice](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseDevice creates a new MTRBaseDevice instance.
func NewMTRBaseDevice() MTRBaseDevice {
	return getMTRBaseDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice
type MTRBaseDevice struct {
	objectivec.Object
}

// MTRBaseDeviceFrom constructs a [MTRBaseDevice] from an unsafe.Pointer.
func MTRBaseDeviceFrom(ptr unsafe.Pointer) MTRBaseDevice {
	return MTRBaseDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseDevice */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/init(nodeID:controller:)
func NewMTRBaseDeviceWithNodeIDController(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController) MTRBaseDevice {
	rv := objc.Send[MTRBaseDevice](objc.ID(getMTRBaseDeviceClass().class), objc.Sel("deviceWithNodeID:controller:"), nodeID, controller)
	return rv
}/* debug [class_init_methods/constructor]: NewMTRBaseDeviceWithNodeIDController */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseDevice */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/init(nodeID:controller:)
func (mc _MTRBaseDeviceClass) DeviceWithNodeIDController(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController) MTRBaseDevice {
	rv := objc.Send[MTRBaseDevice](objc.ID(mc.class), objc.Sel("deviceWithNodeID:controller:"), nodeID, controller)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DeviceWithNodeIDController) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseDevice */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/deregisterReportHandlers(with:completion:)
func (m_ MTRBaseDevice) DeregisterReportHandlersWithQueueCompletion(queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("deregisterReportHandlersWithQueue:completion:"), queue, completion)
}/* debug [instance_methods/method]: DeregisterReportHandlersWithQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/downloadLog(of:timeout:queue:completion:)
func (m_ MTRBaseDevice) DownloadLogOfTypeTimeoutQueueCompletion(type_ unsafe.Pointer, timeout float64, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("downloadLogOfType:timeout:queue:completion:"), type_, timeout, queue, completion)
}/* debug [instance_methods/method]: DownloadLogOfTypeTimeoutQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/invokeCommand(withEndpointID:clusterID:commandID:commandFields:timedInvokeTimeout:queue:completion:)
func (m_ MTRBaseDevice) InvokeCommandWithEndpointIDClusterIDCommandIDCommandFieldsTimedInvokeTimeoutQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, commandID objc.IObject /* cross-framework: NSNumber */, commandFields objc.IObject, timeoutMs objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("invokeCommandWithEndpointID:clusterID:commandID:commandFields:timedInvokeTimeout:queue:completion:"), endpointID, clusterID, commandID, commandFields, timeoutMs, queue, completion)
}/* debug [instance_methods/method]: InvokeCommandWithEndpointIDClusterIDCommandIDCommandFieldsTimedInvokeTimeoutQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/openCommissioningWindow(withDiscriminator:duration:queue:completion:)
func (m_ MTRBaseDevice) OpenCommissioningWindowWithDiscriminatorDurationQueueCompletion(discriminator objc.IObject /* cross-framework: NSNumber */, duration objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("openCommissioningWindowWithDiscriminator:duration:queue:completion:"), discriminator, duration, queue, completion)
}/* debug [instance_methods/method]: OpenCommissioningWindowWithDiscriminatorDurationQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/openCommissioningWindow(withSetupPasscode:discriminator:duration:queue:completion:)
func (m_ MTRBaseDevice) OpenCommissioningWindowWithSetupPasscodeDiscriminatorDurationQueueCompletion(setupPasscode objc.IObject /* cross-framework: NSNumber */, discriminator objc.IObject /* cross-framework: NSNumber */, duration objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("openCommissioningWindowWithSetupPasscode:discriminator:duration:queue:completion:"), setupPasscode, discriminator, duration, queue, completion)
}/* debug [instance_methods/method]: OpenCommissioningWindowWithSetupPasscodeDiscriminatorDurationQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/readAttributePaths(_:eventPaths:params:queue:completion:)
func (m_ MTRBaseDevice) ReadAttributePathsEventPathsParamsQueueCompletion(attributePaths []MTRAttributeRequestPath, eventPaths []MTREventRequestPath, params IMTRReadParams, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributePaths:eventPaths:params:queue:completion:"), attributePaths, eventPaths, params, queue, completion)
}/* debug [instance_methods/method]: ReadAttributePathsEventPathsParamsQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/readAttributes(withEndpointID:clusterID:attributeID:params:queue:completion:)
func (m_ MTRBaseDevice) ReadAttributesWithEndpointIDClusterIDAttributeIDParamsQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */, params IMTRReadParams, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readAttributesWithEndpointID:clusterID:attributeID:params:queue:completion:"), endpointID, clusterID, attributeID, params, queue, completion)
}/* debug [instance_methods/method]: ReadAttributesWithEndpointIDClusterIDAttributeIDParamsQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/readEvents(withEndpointID:clusterID:eventID:params:queue:completion:)
func (m_ MTRBaseDevice) ReadEventsWithEndpointIDClusterIDEventIDParamsQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, eventID objc.IObject /* cross-framework: NSNumber */, params IMTRReadParams, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("readEventsWithEndpointID:clusterID:eventID:params:queue:completion:"), endpointID, clusterID, eventID, params, queue, completion)
}/* debug [instance_methods/method]: ReadEventsWithEndpointIDClusterIDEventIDParamsQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/subscribe(toAttributePaths:eventPaths:params:queue:reportHandler:subscriptionEstablished:resubscriptionScheduled:)
func (m_ MTRBaseDevice) SubscribeToAttributePathsEventPathsParamsQueueReportHandlerSubscriptionEstablishedResubscriptionScheduled(attributePaths []MTRAttributeRequestPath, eventPaths []MTREventRequestPath, params IMTRSubscribeParams, queue unsafe.Pointer, reportHandler unsafe.Pointer, subscriptionEstablished unsafe.Pointer, resubscriptionScheduled unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeToAttributePaths:eventPaths:params:queue:reportHandler:subscriptionEstablished:resubscriptionScheduled:"), attributePaths, eventPaths, params, queue, reportHandler, subscriptionEstablished, resubscriptionScheduled)
}/* debug [instance_methods/method]: SubscribeToAttributePathsEventPathsParamsQueueReportHandlerSubscriptionEstablishedResubscriptionScheduled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/subscribe(with:params:clusterStateCacheContainer:attributeReportHandler:eventReportHandler:errorHandler:subscriptionEstablished:resubscriptionScheduled:)
func (m_ MTRBaseDevice) SubscribeWithQueueParamsClusterStateCacheContainerAttributeReportHandlerEventReportHandlerErrorHandlerSubscriptionEstablishedResubscriptionScheduled(queue unsafe.Pointer, params IMTRSubscribeParams, clusterStateCacheContainer IMTRClusterStateCacheContainer, attributeReportHandler unsafe.Pointer, eventReportHandler unsafe.Pointer, errorHandler unsafe.Pointer, subscriptionEstablished unsafe.Pointer, resubscriptionScheduled unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeWithQueue:params:clusterStateCacheContainer:attributeReportHandler:eventReportHandler:errorHandler:subscriptionEstablished:resubscriptionScheduled:"), queue, params, clusterStateCacheContainer, attributeReportHandler, eventReportHandler, errorHandler, subscriptionEstablished, resubscriptionScheduled)
}/* debug [instance_methods/method]: SubscribeWithQueueParamsClusterStateCacheContainerAttributeReportHandlerEventReportHandlerErrorHandlerSubscriptionEstablishedResubscriptionScheduled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/subscribeToAttributes(withEndpointID:clusterID:attributeID:params:queue:reportHandler:subscriptionEstablished:)
func (m_ MTRBaseDevice) SubscribeToAttributesWithEndpointIDClusterIDAttributeIDParamsQueueReportHandlerSubscriptionEstablished(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */, params IMTRSubscribeParams, queue unsafe.Pointer, reportHandler unsafe.Pointer, subscriptionEstablished unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeToAttributesWithEndpointID:clusterID:attributeID:params:queue:reportHandler:subscriptionEstablished:"), endpointID, clusterID, attributeID, params, queue, reportHandler, subscriptionEstablished)
}/* debug [instance_methods/method]: SubscribeToAttributesWithEndpointIDClusterIDAttributeIDParamsQueueReportHandlerSubscriptionEstablished */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/subscribeToEvents(withEndpointID:clusterID:eventID:params:queue:reportHandler:subscriptionEstablished:)
func (m_ MTRBaseDevice) SubscribeToEventsWithEndpointIDClusterIDEventIDParamsQueueReportHandlerSubscriptionEstablished(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, eventID objc.IObject /* cross-framework: NSNumber */, params IMTRSubscribeParams, queue unsafe.Pointer, reportHandler unsafe.Pointer, subscriptionEstablished unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeToEventsWithEndpointID:clusterID:eventID:params:queue:reportHandler:subscriptionEstablished:"), endpointID, clusterID, eventID, params, queue, reportHandler, subscriptionEstablished)
}/* debug [instance_methods/method]: SubscribeToEventsWithEndpointIDClusterIDEventIDParamsQueueReportHandlerSubscriptionEstablished */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/writeAttribute(withEndpointID:clusterID:attributeID:value:timedWriteTimeout:queue:completion:)
func (m_ MTRBaseDevice) WriteAttributeWithEndpointIDClusterIDAttributeIDValueTimedWriteTimeoutQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */, value objc.IObject, timeoutMs objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeWithEndpointID:clusterID:attributeID:value:timedWriteTimeout:queue:completion:"), endpointID, clusterID, attributeID, value, timeoutMs, queue, completion)
}/* debug [instance_methods/method]: WriteAttributeWithEndpointIDClusterIDAttributeIDValueTimedWriteTimeoutQueueCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseDevice */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseDevice/sessionTransportType
func (m_ MTRBaseDevice) SessionTransportType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sessionTransportType"))
	return rv
}/* debug [instance_properties/getter]: sessionTransportType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseDevice */


