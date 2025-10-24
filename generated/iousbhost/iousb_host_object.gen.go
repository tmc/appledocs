// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [USBHostObject] class.
type IUSBHostObject interface {
	objectivec.IObject
	// properties:
	IOUSBHostDefaultControlCompletionTimeout() float64
	CapabilityDescriptors() USBBOSDescriptor /* not a class type */
	SetCapabilityDescriptors(value USBBOSDescriptor /* not a class type */)
	DeviceAddress() int
	SetDeviceAddress(value int)
	DeviceDescriptor() USBDeviceDescriptor /* not a class type */
	SetDeviceDescriptor(value USBDeviceDescriptor /* not a class type */)
	IoService() unsafe.Pointer
	SetIoService(value unsafe.Pointer)
	Queue() unsafe.Pointer
	SetQueue(value unsafe.Pointer)
	// methods:
	ReferenceMicroframeWithTimeError(time USBHostTime /* not a class type */, error_ unsafe.Pointer) uint64
	SendDeviceRequestDataBytesTransferredCompletionTimeoutError(request USBDeviceRequest /* not a class type */, data objc.IObject /* cross-framework: MutableData */, bytesTransferred uint, completionTimeout float64, error_ unsafe.Pointer) bool
	SendDeviceRequestDataBytesTransferredError(request USBDeviceRequest /* not a class type */, data objc.IObject /* cross-framework: MutableData */, bytesTransferred uint, error_ unsafe.Pointer) bool
}

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

// Alloc allocates a new instance without initialization.
func (uc _USBHostObjectClass) Alloc() USBHostObject {
	rv := objc.Send[USBHostObject](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a USB host object and sets up a communication channel to the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/initWithIOService:options:queue:error:interestHandler:
func NewUSBHostObjectWithIOServiceOptionsQueueErrorInterestHandler(ioService unsafe.Pointer, options USBHostObjectInitOptions /* not a class type */, queue unsafe.Pointer, error_ unsafe.Pointer, interestHandler USBHostInterestHandler /* not a class type */) USBHostObject {
	instance := getUSBHostObjectClass().Alloc()
	rv := objc.Send[USBHostObject](instance.ID, objc.Sel("initWithIOService:options:queue:error:interestHandler:"), ioService, options, queue, error_, interestHandler)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/referenceMicroframeWithTime:error:
func (u_ USBHostObject) ReferenceMicroframeWithTimeError(time USBHostTime /* not a class type */, error_ unsafe.Pointer) uint64 {
	rv := objc.Send[uint64](u_.ID, objc.Sel("referenceMicroframeWithTime:error:"), time, error_)
	return rv
}


// Sends a request on the default control endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/sendDeviceRequest:data:bytesTransferred:completionTimeout:error:
func (u_ USBHostObject) SendDeviceRequestDataBytesTransferredCompletionTimeoutError(request USBDeviceRequest /* not a class type */, data objc.IObject /* cross-framework: MutableData */, bytesTransferred uint, completionTimeout float64, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendDeviceRequest:data:bytesTransferred:completionTimeout:error:"), request, data, bytesTransferred, completionTimeout, error_)
	return rv
}


// Sends a request on the default control endpoint with a default completion timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/sendDeviceRequest:data:bytesTransferred:error:
func (u_ USBHostObject) SendDeviceRequestDataBytesTransferredError(request USBDeviceRequest /* not a class type */, data objc.IObject /* cross-framework: MutableData */, bytesTransferred uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendDeviceRequest:data:bytesTransferred:error:"), request, data, bytesTransferred, error_)
	return rv
}


// The default completion timeout for input/output requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostdefaultcontrolcompletiontimeout
func (u_ USBHostObject) IOUSBHostDefaultControlCompletionTimeout() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("IOUSBHostDefaultControlCompletionTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostobject/capabilitydescriptors
func (u_ USBHostObject) CapabilityDescriptors() USBBOSDescriptor /* not a class type */ {
	rv := objc.Send[USBBOSDescriptor](u_.ID, objc.Sel("capabilityDescriptors"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostobject/capabilitydescriptors
func (u_ USBHostObject) SetCapabilityDescriptors(value USBBOSDescriptor /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCapabilityDescriptors:"), value)
}


// The device’s bus address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostobject/deviceaddress
func (u_ USBHostObject) DeviceAddress() int {
	rv := objc.Send[int](u_.ID, objc.Sel("deviceAddress"))
	return rv
}


// The device’s bus address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostobject/deviceaddress
func (u_ USBHostObject) SetDeviceAddress(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeviceAddress:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostobject/devicedescriptor
func (u_ USBHostObject) DeviceDescriptor() USBDeviceDescriptor /* not a class type */ {
	rv := objc.Send[USBDeviceDescriptor](u_.ID, objc.Sel("deviceDescriptor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostobject/devicedescriptor
func (u_ USBHostObject) SetDeviceDescriptor(value USBDeviceDescriptor /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeviceDescriptor:"), value)
}


// A reference to the kernel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostobject/ioservice
func (u_ USBHostObject) IoService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("ioService"))
	return rv
}


// A reference to the kernel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostobject/ioservice
func (u_ USBHostObject) SetIoService(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIoService:"), value)
}


// The queue for servicing input/output requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostobject/queue
func (u_ USBHostObject) Queue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("queue"))
	return rv
}


// The queue for servicing input/output requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostobject/queue
func (u_ USBHostObject) SetQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setQueue:"), value)
}


