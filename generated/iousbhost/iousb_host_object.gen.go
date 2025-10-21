// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	Destroy()
	ReferenceMicroframeWithTimeError(time unsafe.Pointer, error_ unsafe.Pointer) uint64
	SendDeviceRequestDataBytesTransferredError(request unsafe.Pointer, data unsafe.Pointer, bytesTransferred unsafe.Pointer, error_ unsafe.Pointer) bool
}

// This class provides basic functionality for sending device requests and retrieving descriptors.
//
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
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/initWithIOService:options:queue:error:interestHandler:
func NewUSBHostObjectWithIOServiceOptionsQueueErrorInterestHandler(ioService unsafe.Pointer, options unsafe.Pointer, queue unsafe.Pointer, error_ unsafe.Pointer, interestHandler unsafe.Pointer) USBHostObject {
	instance := getUSBHostObjectClass().Alloc()
	rv := objc.Send[USBHostObject](instance.ID, objc.Sel("initWithIOService:options:queue:error:interestHandler:"), ioService, options, queue, error_, interestHandler)
	rv.Autorelease()
	return rv
}



// Creates a USB host object and sets up a default communication channel to the kernel.
//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/initWithIOService:queue:error:interestHandler:
func NewUSBHostObjectWithIOServiceQueueErrorInterestHandler(ioService unsafe.Pointer, queue unsafe.Pointer, error_ unsafe.Pointer, interestHandler unsafe.Pointer) USBHostObject {
	instance := getUSBHostObjectClass().Alloc()
	rv := objc.Send[USBHostObject](instance.ID, objc.Sel("initWithIOService:queue:error:interestHandler:"), ioService, queue, error_, interestHandler)
	rv.Autorelease()
	return rv
}


// Removes underlying allocations and connections from the USB host object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/destroy()
func (u_ USBHostObject) Destroy() {
	objc.Send[objc.ID](u_.ID, objc.Sel("destroy"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/referenceMicroframeWithTime:error:
func (u_ USBHostObject) ReferenceMicroframeWithTimeError(time unsafe.Pointer, error_ unsafe.Pointer) uint64 {
	rv := objc.Send[uint64](u_.ID, objc.Sel("referenceMicroframeWithTime:error:"), time, error_)
	return rv
}

// Sends a request on the default control endpoint with a default completion timeout.
//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/sendDeviceRequest:data:bytesTransferred:error:
func (u_ USBHostObject) SendDeviceRequestDataBytesTransferredError(request unsafe.Pointer, data unsafe.Pointer, bytesTransferred unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendDeviceRequest:data:bytesTransferred:error:"), request, data, bytesTransferred, error_)
	return rv
}

// A reference to the kernel object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/ioService
func (u_ USBHostObject) IoService() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("ioService"))
	return rv
}

// The queue for servicing input/output requests.
//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/queue
func (u_ USBHostObject) Queue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("queue"))
	return rv
}


