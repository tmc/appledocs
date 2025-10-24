// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOUSBHostObject */


/* debug [class_header]: Header for IOUSBHostObject */
// The class instance for the [USBHostObject] class.
var (
	USBHostObjectClass     _USBHostObjectClass
	USBHostObjectClassOnce sync.Once
)

func getUSBHostObjectClass() _USBHostObjectClass {
	USBHostObjectClassOnce.Do(func() {
		USBHostObjectClass = _USBHostObjectClass{objc.GetClass("IOUSBHostObject")}
	})
	return USBHostObjectClass
}

type _USBHostObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for USBHostObject */
// An interface definition for the [USBHostObject] class.
type IUSBHostObject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for USBHostObject */
	// properties:
	CapabilityDescriptors() USBBOSDescriptor /* not a class type */
	DeviceAddress() uint
	DeviceDescriptor() USBDeviceDescriptor /* not a class type */
	IoService() unsafe.Pointer
	Queue() unsafe.Pointer
	IOUSBHostDefaultControlCompletionTimeout() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for USBHostObject */
	// methods:
	AbortDeviceRequestsWithError(error_ unsafe.Pointer) bool
	AbortDeviceRequestsWithOptionError(option USBHostAbortOption, error_ unsafe.Pointer) bool
	ConfigurationDescriptorWithIndexError(index uint, error_ unsafe.Pointer) USBConfigurationDescriptor /* not a class type */
	ConfigurationDescriptorWithConfigurationValueError(configurationValue uint, error_ unsafe.Pointer) USBConfigurationDescriptor /* not a class type */
	CurrentMicroframeWithTimeError(time USBHostTime /* typedef */, error_ unsafe.Pointer) uint64
	DescriptorWithTypeLengthError(type_ unsafe.Pointer, length uint, error_ unsafe.Pointer) USBDescriptor /* not a class type */
	DescriptorWithTypeLengthIndexLanguageIDError(type_ unsafe.Pointer, length uint, index uint, languageID uint, error_ unsafe.Pointer) USBDescriptor /* not a class type */
	DescriptorWithTypeLengthIndexLanguageIDRequestTypeRequestRecipientError(type_ unsafe.Pointer, length uint, index uint, languageID uint, requestType unsafe.Pointer, requestRecipient unsafe.Pointer, error_ unsafe.Pointer) USBDescriptor /* not a class type */
	Destroy()
	DestroyWithOptions(options USBHostObjectDestroyOptions)
	EnqueueDeviceRequestDataCompletionTimeoutErrorCompletionHandler(request USBDeviceRequest /* not a class type */, data foundation.MutableData, completionTimeout float64, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool
	EnqueueDeviceRequestDataErrorCompletionHandler(request USBDeviceRequest /* not a class type */, data foundation.MutableData, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool
	EnqueueDeviceRequestErrorCompletionHandler(request USBDeviceRequest /* not a class type */, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool
	FrameNumberWithTime(time USBHostTime /* typedef */) uint64
	IoDataWithCapacityError(capacity uint, error_ unsafe.Pointer) foundation.MutableData
	ReferenceMicroframeWithTimeError(time USBHostTime /* typedef */, error_ unsafe.Pointer) uint64
	SendDeviceRequestDataBytesTransferredCompletionTimeoutError(request USBDeviceRequest /* not a class type */, data foundation.MutableData, bytesTransferred uint, completionTimeout float64, error_ unsafe.Pointer) bool
	SendDeviceRequestDataBytesTransferredError(request USBDeviceRequest /* not a class type */, data foundation.MutableData, bytesTransferred uint, error_ unsafe.Pointer) bool
	SendDeviceRequestError(request USBDeviceRequest /* not a class type */, error_ unsafe.Pointer) bool
	StringWithIndexError(index uint, error_ unsafe.Pointer) foundation.String
	StringWithIndexLanguageIDError(index uint, languageID uint, error_ unsafe.Pointer) foundation.String
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for USBHostObject */
// Alloc allocates a new instance without initialization.
func (uc _USBHostObjectClass) Alloc() USBHostObject {
	rv := objc.Send[USBHostObject](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _USBHostObjectClass) New() USBHostObject {
	rv := objc.Send[USBHostObject](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ USBHostObject) Init() USBHostObject {
	rv := objc.Send[USBHostObject](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ USBHostObject) Autorelease() USBHostObject {
	rv := objc.Send[USBHostObject](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUSBHostObject creates a new USBHostObject instance.
func NewUSBHostObject() USBHostObject {
	return getUSBHostObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for USBHostObject */
// This class provides basic functionality for sending device requests and retrieving descriptors.


// This class provides basic functionality for sending device requests and retrieving descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject
type USBHostObject struct {
	objectivec.Object
}

// USBHostObjectFrom constructs a [USBHostObject] from an unsafe.Pointer.
//
// This class provides basic functionality for sending device requests and retrieving descriptors.
func USBHostObjectFrom(ptr unsafe.Pointer) USBHostObject {
	return USBHostObject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for USBHostObject */

// Creates a USB host object and sets up a communication channel to the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/initWithIOService:options:queue:error:interestHandler:
func NewUSBHostObjectWithIOServiceOptionsQueueErrorInterestHandler(ioService unsafe.Pointer, options USBHostObjectInitOptions, queue unsafe.Pointer, error_ unsafe.Pointer, interestHandler USBHostInterestHandler /* not a class type */) USBHostObject {
	instance := getUSBHostObjectClass().Alloc()
	rv := objc.Send[USBHostObject](instance.ID, objc.Sel("initWithIOService:options:queue:error:interestHandler:"), ioService, options, queue, error_, interestHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUSBHostObjectWithIOServiceOptionsQueueErrorInterestHandler */


// Creates a USB host object and sets up a default communication channel to the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/initWithIOService:queue:error:interestHandler:
func NewUSBHostObjectWithIOServiceQueueErrorInterestHandler(ioService unsafe.Pointer, queue unsafe.Pointer, error_ unsafe.Pointer, interestHandler USBHostInterestHandler /* not a class type */) USBHostObject {
	instance := getUSBHostObjectClass().Alloc()
	rv := objc.Send[USBHostObject](instance.ID, objc.Sel("initWithIOService:queue:error:interestHandler:"), ioService, queue, error_, interestHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUSBHostObjectWithIOServiceQueueErrorInterestHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for USBHostObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for USBHostObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for USBHostObject */

// Aborts device requests synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/abortDeviceRequestsWithError:
func (u_ USBHostObject) AbortDeviceRequestsWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("abortDeviceRequestsWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: AbortDeviceRequestsWithError */


// Aborts device requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/abortDeviceRequestsWithOption:error:
func (u_ USBHostObject) AbortDeviceRequestsWithOptionError(option USBHostAbortOption, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("abortDeviceRequestsWithOption:error:"), option, error_)
	return rv
}/* debug [instance_methods/method]: AbortDeviceRequestsWithOptionError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/configurationDescriptor(with:)
func (u_ USBHostObject) ConfigurationDescriptorWithIndexError(index uint, error_ unsafe.Pointer) USBConfigurationDescriptor /* not a class type */ {
	rv := objc.Send[USBConfigurationDescriptor](u_.ID, objc.Sel("configurationDescriptorWithIndex:error:"), index, error_)
	return rv
}/* debug [instance_methods/method]: ConfigurationDescriptorWithIndexError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/configurationDescriptor(withConfigurationValue:)
func (u_ USBHostObject) ConfigurationDescriptorWithConfigurationValueError(configurationValue uint, error_ unsafe.Pointer) USBConfigurationDescriptor /* not a class type */ {
	rv := objc.Send[USBConfigurationDescriptor](u_.ID, objc.Sel("configurationDescriptorWithConfigurationValue:error:"), configurationValue, error_)
	return rv
}/* debug [instance_methods/method]: ConfigurationDescriptorWithConfigurationValueError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/currentMicroframeWithTime:error:
func (u_ USBHostObject) CurrentMicroframeWithTimeError(time USBHostTime /* typedef */, error_ unsafe.Pointer) uint64 {
	rv := objc.Send[uint64](u_.ID, objc.Sel("currentMicroframeWithTime:error:"), time, error_)
	return rv
}/* debug [instance_methods/method]: CurrentMicroframeWithTimeError */


// Retrieves a descriptor with default arguments from the cache or the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/descriptorWithType:length:error:
func (u_ USBHostObject) DescriptorWithTypeLengthError(type_ unsafe.Pointer, length uint, error_ unsafe.Pointer) USBDescriptor /* not a class type */ {
	rv := objc.Send[USBDescriptor](u_.ID, objc.Sel("descriptorWithType:length:error:"), type_, length, error_)
	return rv
}/* debug [instance_methods/method]: DescriptorWithTypeLengthError */


// Retrieves a string descriptor from the cache or the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/descriptorWithType:length:index:languageID:error:
func (u_ USBHostObject) DescriptorWithTypeLengthIndexLanguageIDError(type_ unsafe.Pointer, length uint, index uint, languageID uint, error_ unsafe.Pointer) USBDescriptor /* not a class type */ {
	rv := objc.Send[USBDescriptor](u_.ID, objc.Sel("descriptorWithType:length:index:languageID:error:"), type_, length, index, languageID, error_)
	return rv
}/* debug [instance_methods/method]: DescriptorWithTypeLengthIndexLanguageIDError */


// Retrieves a descriptor from the cache or the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/descriptorWithType:length:index:languageID:requestType:requestRecipient:error:
func (u_ USBHostObject) DescriptorWithTypeLengthIndexLanguageIDRequestTypeRequestRecipientError(type_ unsafe.Pointer, length uint, index uint, languageID uint, requestType unsafe.Pointer, requestRecipient unsafe.Pointer, error_ unsafe.Pointer) USBDescriptor /* not a class type */ {
	rv := objc.Send[USBDescriptor](u_.ID, objc.Sel("descriptorWithType:length:index:languageID:requestType:requestRecipient:error:"), type_, length, index, languageID, requestType, requestRecipient, error_)
	return rv
}/* debug [instance_methods/method]: DescriptorWithTypeLengthIndexLanguageIDRequestTypeRequestRecipientError */


// Removes underlying allocations and connections from the USB host object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/destroy()
func (u_ USBHostObject) Destroy() {
	objc.Send[objc.ID](u_.ID, objc.Sel("destroy"))
}/* debug [instance_methods/method]: Destroy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/destroy(options:)
func (u_ USBHostObject) DestroyWithOptions(options USBHostObjectDestroyOptions) {
	objc.Send[objc.ID](u_.ID, objc.Sel("destroyWithOptions:"), options)
}/* debug [instance_methods/method]: DestroyWithOptions */


// Enqueues a request on the default control endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/enqueueDeviceRequest:data:completionTimeout:error:completionHandler:
func (u_ USBHostObject) EnqueueDeviceRequestDataCompletionTimeoutErrorCompletionHandler(request USBDeviceRequest /* not a class type */, data foundation.MutableData, completionTimeout float64, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueDeviceRequest:data:completionTimeout:error:completionHandler:"), request, data, completionTimeout, error_, completionHandler)
	return rv
}/* debug [instance_methods/method]: EnqueueDeviceRequestDataCompletionTimeoutErrorCompletionHandler */


// Enqueues a request on the default control endpoint with a default completion timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/enqueueDeviceRequest:data:error:completionHandler:
func (u_ USBHostObject) EnqueueDeviceRequestDataErrorCompletionHandler(request USBDeviceRequest /* not a class type */, data foundation.MutableData, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueDeviceRequest:data:error:completionHandler:"), request, data, error_, completionHandler)
	return rv
}/* debug [instance_methods/method]: EnqueueDeviceRequestDataErrorCompletionHandler */


// Enqueues a request on the default control endpoint without a data phase and a default timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/enqueueDeviceRequest:error:completionHandler:
func (u_ USBHostObject) EnqueueDeviceRequestErrorCompletionHandler(request USBDeviceRequest /* not a class type */, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueDeviceRequest:error:completionHandler:"), request, error_, completionHandler)
	return rv
}/* debug [instance_methods/method]: EnqueueDeviceRequestErrorCompletionHandler */


// Returns the current frame number of the USB controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/frameNumberWithTime:
func (u_ USBHostObject) FrameNumberWithTime(time USBHostTime /* typedef */) uint64 {
	rv := objc.Send[uint64](u_.ID, objc.Sel("frameNumberWithTime:"), time)
	return rv
}/* debug [instance_methods/method]: FrameNumberWithTime */


// Allocates a buffer for input/output requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/ioData(withCapacity:)
func (u_ USBHostObject) IoDataWithCapacityError(capacity uint, error_ unsafe.Pointer) foundation.MutableData {
	rv := objc.Send[foundation.MutableData](u_.ID, objc.Sel("ioDataWithCapacity:error:"), capacity, error_)
	return rv
}/* debug [instance_methods/method]: IoDataWithCapacityError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/referenceMicroframeWithTime:error:
func (u_ USBHostObject) ReferenceMicroframeWithTimeError(time USBHostTime /* typedef */, error_ unsafe.Pointer) uint64 {
	rv := objc.Send[uint64](u_.ID, objc.Sel("referenceMicroframeWithTime:error:"), time, error_)
	return rv
}/* debug [instance_methods/method]: ReferenceMicroframeWithTimeError */


// Sends a request on the default control endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/sendDeviceRequest:data:bytesTransferred:completionTimeout:error:
func (u_ USBHostObject) SendDeviceRequestDataBytesTransferredCompletionTimeoutError(request USBDeviceRequest /* not a class type */, data foundation.MutableData, bytesTransferred uint, completionTimeout float64, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendDeviceRequest:data:bytesTransferred:completionTimeout:error:"), request, data, bytesTransferred, completionTimeout, error_)
	return rv
}/* debug [instance_methods/method]: SendDeviceRequestDataBytesTransferredCompletionTimeoutError */


// Sends a request on the default control endpoint with a default completion timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/sendDeviceRequest:data:bytesTransferred:error:
func (u_ USBHostObject) SendDeviceRequestDataBytesTransferredError(request USBDeviceRequest /* not a class type */, data foundation.MutableData, bytesTransferred uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendDeviceRequest:data:bytesTransferred:error:"), request, data, bytesTransferred, error_)
	return rv
}/* debug [instance_methods/method]: SendDeviceRequestDataBytesTransferredError */


// Sends a request on the default control endpoint without a data phase and default completion timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/sendDeviceRequest:error:
func (u_ USBHostObject) SendDeviceRequestError(request USBDeviceRequest /* not a class type */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendDeviceRequest:error:"), request, error_)
	return rv
}/* debug [instance_methods/method]: SendDeviceRequestError */


// Retrieves an English-language string from a string descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/stringWithIndex:error:
func (u_ USBHostObject) StringWithIndexError(index uint, error_ unsafe.Pointer) foundation.String {
	rv := objc.Send[foundation.String](u_.ID, objc.Sel("stringWithIndex:error:"), index, error_)
	return rv
}/* debug [instance_methods/method]: StringWithIndexError */


// Retrieves a string from a string descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/stringWithIndex:languageID:error:
func (u_ USBHostObject) StringWithIndexLanguageIDError(index uint, languageID uint, error_ unsafe.Pointer) foundation.String {
	rv := objc.Send[foundation.String](u_.ID, objc.Sel("stringWithIndex:languageID:error:"), index, languageID, error_)
	return rv
}/* debug [instance_methods/method]: StringWithIndexLanguageIDError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for USBHostObject */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/capabilityDescriptors
func (u_ USBHostObject) CapabilityDescriptors() USBBOSDescriptor /* not a class type */ {
	rv := objc.Send[USBBOSDescriptor](u_.ID, objc.Sel("capabilityDescriptors"))
	return rv
}/* debug [instance_properties/getter]: capabilityDescriptors */


// The device’s bus address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/deviceAddress
func (u_ USBHostObject) DeviceAddress() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("deviceAddress"))
	return rv
}/* debug [instance_properties/getter]: deviceAddress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/deviceDescriptor
func (u_ USBHostObject) DeviceDescriptor() USBDeviceDescriptor /* not a class type */ {
	rv := objc.Send[USBDeviceDescriptor](u_.ID, objc.Sel("deviceDescriptor"))
	return rv
}/* debug [instance_properties/getter]: deviceDescriptor */


// A reference to the kernel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/ioService
func (u_ USBHostObject) IoService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("ioService"))
	return rv
}/* debug [instance_properties/getter]: ioService */


// The queue for servicing input/output requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/queue
func (u_ USBHostObject) Queue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("queue"))
	return rv
}/* debug [instance_properties/getter]: queue */


// The default completion timeout for input/output requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostdefaultcontrolcompletiontimeout
func (u_ USBHostObject) IOUSBHostDefaultControlCompletionTimeout() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("IOUSBHostDefaultControlCompletionTimeout"))
	return rv
}/* debug [instance_properties/getter]: IOUSBHostDefaultControlCompletionTimeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOUSBHostObject */


