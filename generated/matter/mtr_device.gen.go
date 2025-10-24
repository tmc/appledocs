// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDevice */


/* debug [class_header]: Header for MTRDevice */
// The class instance for the [MTRDevice] class.
var (
	MTRDeviceClass     _MTRDeviceClass
	MTRDeviceClassOnce sync.Once
)

func getMTRDeviceClass() _MTRDeviceClass {
	MTRDeviceClassOnce.Do(func() {
		MTRDeviceClass = _MTRDeviceClass{objc.GetClass("MTRDevice")}
	})
	return MTRDeviceClass
}

type _MTRDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDevice */
// An interface definition for the [MTRDevice] class.
type IMTRDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDevice */
	// properties:
	DeviceCachePrimed() bool
	DeviceController() IMTRDeviceController
	EstimatedStartTime() objc.IObject /* cross-framework: NSDate */
	EstimatedSubscriptionLatency() objc.IObject /* cross-framework: NSNumber */
	NetworkCommissioningFeatures() unsafe.Pointer
	NodeID() objc.IObject /* cross-framework: NSNumber */
	ProductID() objc.IObject /* cross-framework: NSNumber */
	State() unsafe.Pointer
	VendorID() objc.IObject /* cross-framework: NSNumber */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDevice */
	// methods:
	AddDelegateQueue(delegate unsafe.Pointer, queue unsafe.Pointer)
	AddDelegateQueueInterestedPathsForAttributesInterestedPathsForEvents(delegate unsafe.Pointer, queue unsafe.Pointer, interestedPathsForAttributes objc.IObject /* cross-framework: NSArray */, interestedPathsForEvents objc.IObject /* cross-framework: NSArray */)
	DescriptorClusters() foundation.IDictionary
	DownloadLogOfTypeTimeoutQueueCompletion(type_ unsafe.Pointer, timeout float64, queue unsafe.Pointer, completion unsafe.Pointer)
	InvokeCommandWithEndpointIDClusterIDCommandIDCommandFieldsExpectedValuesExpectedValueIntervalQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, commandID objc.IObject /* cross-framework: NSNumber */, commandFields foundation.IDictionary, expectedValues foundation.IDictionary, expectedValueInterval objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer)
	InvokeCommandWithEndpointIDClusterIDCommandIDCommandFieldsExpectedValuesExpectedValueIntervalTimedInvokeTimeoutQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, commandID objc.IObject /* cross-framework: NSNumber */, commandFields objc.IObject, expectedValues foundation.IDictionary, expectedValueInterval objc.IObject /* cross-framework: NSNumber */, timeout objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer)
	InvokeCommandsQueueCompletion(commands []foundation.Array, queue unsafe.Pointer, completion unsafe.Pointer)
	OpenCommissioningWindowWithDiscriminatorDurationQueueCompletion(discriminator objc.IObject /* cross-framework: NSNumber */, duration objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer)
	OpenCommissioningWindowWithSetupPasscodeDiscriminatorDurationQueueCompletion(setupPasscode objc.IObject /* cross-framework: NSNumber */, discriminator objc.IObject /* cross-framework: NSNumber */, duration objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer)
	ReadAttributeWithEndpointIDClusterIDAttributeIDParams(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */, params IMTRReadParams) foundation.IDictionary
	ReadAttributePaths(attributePaths []MTRAttributeRequestPath) foundation.IDictionary
	RemoveDelegate(delegate unsafe.Pointer)
	WaitForAttributeValuesTimeoutQueueCompletion(values foundation.IDictionary, timeout float64, queue unsafe.Pointer, completion unsafe.Pointer) IMTRAttributeValueWaiter
	WriteAttributeWithEndpointIDClusterIDAttributeIDValueExpectedValueIntervalTimedWriteTimeout(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */, value objc.IObject, expectedValueInterval objc.IObject /* cross-framework: NSNumber */, timeout objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDevice */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceClass) Alloc() MTRDevice {
	rv := objc.Send[MTRDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceClass) New() MTRDevice {
	rv := objc.Send[MTRDevice](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDevice) Init() MTRDevice {
	rv := objc.Send[MTRDevice](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDevice) Autorelease() MTRDevice {
	rv := objc.Send[MTRDevice](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDevice creates a new MTRDevice instance.
func NewMTRDevice() MTRDevice {
	return getMTRDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice
type MTRDevice struct {
	objectivec.Object
}

// MTRDeviceFrom constructs a [MTRDevice] from an unsafe.Pointer.
func MTRDeviceFrom(ptr unsafe.Pointer) MTRDevice {
	return MTRDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDevice */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/init(nodeID:controller:)
func NewMTRDeviceWithNodeIDController(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController) MTRDevice {
	rv := objc.Send[MTRDevice](objc.ID(getMTRDeviceClass().class), objc.Sel("deviceWithNodeID:controller:"), nodeID, controller)
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDeviceWithNodeIDController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/init(nodeID:deviceController:)
func NewMTRDeviceWithNodeIDDeviceController(nodeID uint64, deviceController IMTRDeviceController) MTRDevice {
	rv := objc.Send[MTRDevice](objc.ID(getMTRDeviceClass().class), objc.Sel("deviceWithNodeID:deviceController:"), nodeID, deviceController)
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDeviceWithNodeIDDeviceController */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDevice */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/init(nodeID:controller:)
func (mc _MTRDeviceClass) DeviceWithNodeIDController(nodeID objc.IObject /* cross-framework: NSNumber */, controller IMTRDeviceController) MTRDevice {
	rv := objc.Send[MTRDevice](objc.ID(mc.class), objc.Sel("deviceWithNodeID:controller:"), nodeID, controller)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DeviceWithNodeIDController) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/init(nodeID:deviceController:)
func (mc _MTRDeviceClass) DeviceWithNodeIDDeviceController(nodeID uint64, deviceController IMTRDeviceController) MTRDevice {
	rv := objc.Send[MTRDevice](objc.ID(mc.class), objc.Sel("deviceWithNodeID:deviceController:"), nodeID, deviceController)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DeviceWithNodeIDDeviceController) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDevice */

// Adds a delegate to receive asynchronous callbacks about the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/add(_:queue:)
func (m_ MTRDevice) AddDelegateQueue(delegate unsafe.Pointer, queue unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addDelegate:queue:"), delegate, queue)
}/* debug [instance_methods/method]: AddDelegateQueue */


// Adds a delegate to receive asynchronous callbacks about the device, and limit attribute and/or event reports to a specific set of paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/add(_:queue:interestedPathsForAttributes:interestedPathsForEvents:)
func (m_ MTRDevice) AddDelegateQueueInterestedPathsForAttributesInterestedPathsForEvents(delegate unsafe.Pointer, queue unsafe.Pointer, interestedPathsForAttributes objc.IObject /* cross-framework: NSArray */, interestedPathsForEvents objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addDelegate:queue:interestedPathsForAttributes:interestedPathsForEvents:"), delegate, queue, interestedPathsForAttributes, interestedPathsForEvents)
}/* debug [instance_methods/method]: AddDelegateQueueInterestedPathsForAttributesInterestedPathsForEvents */


// Read all known attributes from descriptor clusters on all known endpoints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/descriptorClusters()
func (m_ MTRDevice) DescriptorClusters() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("descriptorClusters"))
	return rv
}/* debug [instance_methods/method]: DescriptorClusters */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/downloadLog(of:timeout:queue:completion:)
func (m_ MTRDevice) DownloadLogOfTypeTimeoutQueueCompletion(type_ unsafe.Pointer, timeout float64, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("downloadLogOfType:timeout:queue:completion:"), type_, timeout, queue, completion)
}/* debug [instance_methods/method]: DownloadLogOfTypeTimeoutQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/invokeCommand(withEndpointID:clusterID:commandID:commandFields:expectedValues:expectedValueInterval:queue:completion:)
func (m_ MTRDevice) InvokeCommandWithEndpointIDClusterIDCommandIDCommandFieldsExpectedValuesExpectedValueIntervalQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, commandID objc.IObject /* cross-framework: NSNumber */, commandFields foundation.IDictionary, expectedValues foundation.IDictionary, expectedValueInterval objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("invokeCommandWithEndpointID:clusterID:commandID:commandFields:expectedValues:expectedValueInterval:queue:completion:"), endpointID, clusterID, commandID, commandFields, expectedValues, expectedValueInterval, queue, completion)
}/* debug [instance_methods/method]: InvokeCommandWithEndpointIDClusterIDCommandIDCommandFieldsExpectedValuesExpectedValueIntervalQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/invokeCommand(withEndpointID:clusterID:commandID:commandFields:expectedValues:expectedValueInterval:timedInvokeTimeout:queue:completion:)
func (m_ MTRDevice) InvokeCommandWithEndpointIDClusterIDCommandIDCommandFieldsExpectedValuesExpectedValueIntervalTimedInvokeTimeoutQueueCompletion(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, commandID objc.IObject /* cross-framework: NSNumber */, commandFields objc.IObject, expectedValues foundation.IDictionary, expectedValueInterval objc.IObject /* cross-framework: NSNumber */, timeout objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("invokeCommandWithEndpointID:clusterID:commandID:commandFields:expectedValues:expectedValueInterval:timedInvokeTimeout:queue:completion:"), endpointID, clusterID, commandID, commandFields, expectedValues, expectedValueInterval, timeout, queue, completion)
}/* debug [instance_methods/method]: InvokeCommandWithEndpointIDClusterIDCommandIDCommandFieldsExpectedValuesExpectedValueIntervalTimedInvokeTimeoutQueueCompletion */


// Invoke one or more groups of commands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/invokeCommands(_:queue:completion:)
func (m_ MTRDevice) InvokeCommandsQueueCompletion(commands []foundation.Array, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("invokeCommands:queue:completion:"), commands, queue, completion)
}/* debug [instance_methods/method]: InvokeCommandsQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/openCommissioningWindow(withDiscriminator:duration:queue:completion:)
func (m_ MTRDevice) OpenCommissioningWindowWithDiscriminatorDurationQueueCompletion(discriminator objc.IObject /* cross-framework: NSNumber */, duration objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("openCommissioningWindowWithDiscriminator:duration:queue:completion:"), discriminator, duration, queue, completion)
}/* debug [instance_methods/method]: OpenCommissioningWindowWithDiscriminatorDurationQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/openCommissioningWindow(withSetupPasscode:discriminator:duration:queue:completion:)
func (m_ MTRDevice) OpenCommissioningWindowWithSetupPasscodeDiscriminatorDurationQueueCompletion(setupPasscode objc.IObject /* cross-framework: NSNumber */, discriminator objc.IObject /* cross-framework: NSNumber */, duration objc.IObject /* cross-framework: NSNumber */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("openCommissioningWindowWithSetupPasscode:discriminator:duration:queue:completion:"), setupPasscode, discriminator, duration, queue, completion)
}/* debug [instance_methods/method]: OpenCommissioningWindowWithSetupPasscodeDiscriminatorDurationQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/readAttribute(withEndpointID:clusterID:attributeID:params:)
func (m_ MTRDevice) ReadAttributeWithEndpointIDClusterIDAttributeIDParams(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */, params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeWithEndpointID:clusterID:attributeID:params:"), endpointID, clusterID, attributeID, params)
	return rv
}/* debug [instance_methods/method]: ReadAttributeWithEndpointIDClusterIDAttributeIDParams */


// Read the attributes identified by the provided attribute paths. The paths can include wildcards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/readAttributePaths(_:)
func (m_ MTRDevice) ReadAttributePaths(attributePaths []MTRAttributeRequestPath) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributePaths:"), attributePaths)
	return rv
}/* debug [instance_methods/method]: ReadAttributePaths */


// Removes the delegate from receiving callbacks about the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/remove(_:)
func (m_ MTRDevice) RemoveDelegate(delegate unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeDelegate:"), delegate)
}/* debug [instance_methods/method]: RemoveDelegate */


// Sets up the provided completion to be called when any of the following happens:
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/wait(forAttributeValues:timeout:queue:completion:)
func (m_ MTRDevice) WaitForAttributeValuesTimeoutQueueCompletion(values foundation.IDictionary, timeout float64, queue unsafe.Pointer, completion unsafe.Pointer) IMTRAttributeValueWaiter {
	rv := objc.Send[MTRAttributeValueWaiter](m_.ID, objc.Sel("waitForAttributeValues:timeout:queue:completion:"), values, timeout, queue, completion)
	return rv
}/* debug [instance_methods/method]: WaitForAttributeValuesTimeoutQueueCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/writeAttribute(withEndpointID:clusterID:attributeID:value:expectedValueInterval:timedWriteTimeout:)
func (m_ MTRDevice) WriteAttributeWithEndpointIDClusterIDAttributeIDValueExpectedValueIntervalTimedWriteTimeout(endpointID objc.IObject /* cross-framework: NSNumber */, clusterID objc.IObject /* cross-framework: NSNumber */, attributeID objc.IObject /* cross-framework: NSNumber */, value objc.IObject, expectedValueInterval objc.IObject /* cross-framework: NSNumber */, timeout objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("writeAttributeWithEndpointID:clusterID:attributeID:value:expectedValueInterval:timedWriteTimeout:"), endpointID, clusterID, attributeID, value, expectedValueInterval, timeout)
}/* debug [instance_methods/method]: WriteAttributeWithEndpointIDClusterIDAttributeIDValueExpectedValueIntervalTimedWriteTimeout */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDevice */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/deviceCachePrimed
func (m_ MTRDevice) DeviceCachePrimed() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("deviceCachePrimed"))
	return rv
}/* debug [instance_properties/getter]: deviceCachePrimed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/deviceController
func (m_ MTRDevice) DeviceController() IMTRDeviceController {
	rv := objc.Send[MTRDeviceController](m_.ID, objc.Sel("deviceController"))
	return rv
}/* debug [instance_properties/getter]: deviceController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/estimatedStartTime
func (m_ MTRDevice) EstimatedStartTime() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("estimatedStartTime"))
	return rv
}/* debug [instance_properties/getter]: estimatedStartTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/estimatedSubscriptionLatency
func (m_ MTRDevice) EstimatedSubscriptionLatency() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("estimatedSubscriptionLatency"))
	return rv
}/* debug [instance_properties/getter]: estimatedSubscriptionLatency */


// Network commissioning features supported by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/networkCommissioningFeatures
func (m_ MTRDevice) NetworkCommissioningFeatures() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("networkCommissioningFeatures"))
	return rv
}/* debug [instance_properties/getter]: networkCommissioningFeatures */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/nodeID
func (m_ MTRDevice) NodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeID"))
	return rv
}/* debug [instance_properties/getter]: nodeID */


// The Product Identifier associated with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/productID
func (m_ MTRDevice) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}/* debug [instance_properties/getter]: productID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/state
func (m_ MTRDevice) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The Vendor Identifier associated with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDevice/vendorID
func (m_ MTRDevice) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}/* debug [instance_properties/getter]: vendorID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDevice */


