// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [USBHostPipe] class.
var (
	USBHostPipeClass     _USBHostPipeClass
	USBHostPipeClassOnce sync.Once
)

func getUSBHostPipeClass() _USBHostPipeClass {
	USBHostPipeClassOnce.Do(func() {
		USBHostPipeClass = _USBHostPipeClass{objc.GetClass("IOUSBHostPipe")}
	})
	return USBHostPipeClass
}

type _USBHostPipeClass struct {
	class objc.Class
}

// An interface definition for the [USBHostPipe] class.
type IUSBHostPipe interface {
	IUSBHostIOSource
	CopyStreamWithStreamIDError(streamID uint, error_ unsafe.Pointer) USBHostStream
	DisableStreamsWithError(error_ unsafe.Pointer) bool
	EnableStreamsWithError(error_ unsafe.Pointer) bool
	EnqueueControlRequestDataCompletionTimeoutErrorCompletionHandler(request unsafe.Pointer, data foundation.IMutableData, completionTimeout foundation.ITimeInterval, error_ unsafe.Pointer, completionHandler unsafe.Pointer) bool
	SendControlRequestError(request unsafe.Pointer, error_ unsafe.Pointer) bool
	SendIORequestWithDataBytesTransferredCompletionTimeoutError(data foundation.IMutableData, bytesTransferred unsafe.Pointer, completionTimeout foundation.ITimeInterval, error_ unsafe.Pointer) bool
	IOUSBHostDefaultControlCompletionTimeout() unsafe.Pointer
	Descriptors() unsafe.Pointer
	SetDescriptors(value unsafe.Pointer)
	IdleTimeout() unsafe.Pointer
	SetIdleTimeout(value unsafe.Pointer)
	OriginalDescriptors() unsafe.Pointer
	SetOriginalDescriptors(value unsafe.Pointer)
}

// The class that sends control, bulk, interrupt, and isochronous input/output requests for function drivers, and manages stream capabilities.
//
// The client creates pipe objects using .


// The class that sends control, bulk, interrupt, and isochronous input/output requests for function drivers, and manages stream capabilities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe
type USBHostPipe struct {
	USBHostIOSource
}

// USBHostPipeFrom constructs a [USBHostPipe] from an unsafe.Pointer.
//
// The class that sends control, bulk, interrupt, and isochronous input/output requests for function drivers, and manages stream capabilities.
func USBHostPipeFrom(ptr unsafe.Pointer) USBHostPipe {
	return USBHostPipe{
		USBHostIOSource: USBHostIOSourceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _USBHostPipeClass) Alloc() USBHostPipe {
	rv := objc.Send[USBHostPipe](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _USBHostPipeClass) New() USBHostPipe {
	rv := objc.Send[USBHostPipe](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ USBHostPipe) Init() USBHostPipe {
	rv := objc.Send[USBHostPipe](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ USBHostPipe) Autorelease() USBHostPipe {
	rv := objc.Send[USBHostPipe](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUSBHostPipe creates a new USBHostPipe instance.
func NewUSBHostPipe() USBHostPipe {
	return getUSBHostPipeClass().New()
}



// Returns the stream for a stream ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/copyStream(withStreamID:)
func (u_ USBHostPipe) CopyStreamWithStreamIDError(streamID uint, error_ unsafe.Pointer) USBHostStream {
	rv := objc.Send[USBHostStream](u_.ID, objc.Sel("copyStreamWithStreamID:error:"), streamID, error_)
	return rv
}


// Disables streams for the pipe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/disableStreams()
func (u_ USBHostPipe) DisableStreamsWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("disableStreamsWithError:"), error_)
	return rv
}


// Enables streams for the pipe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/enableStreams()
func (u_ USBHostPipe) EnableStreamsWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enableStreamsWithError:"), error_)
	return rv
}


// Enqueues a request on a control endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/enqueueControlRequest:data:completionTimeout:error:completionHandler:
func (u_ USBHostPipe) EnqueueControlRequestDataCompletionTimeoutErrorCompletionHandler(request unsafe.Pointer, data foundation.IMutableData, completionTimeout foundation.ITimeInterval, error_ unsafe.Pointer, completionHandler unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueControlRequest:data:completionTimeout:error:completionHandler:"), request, data, completionTimeout, error_, completionHandler)
	return rv
}


// Sends a request on a control endpoint without a data phase and a default completion timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/sendControlRequest:error:
func (u_ USBHostPipe) SendControlRequestError(request unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendControlRequest:error:"), request, error_)
	return rv
}


// Sends an input/output request on the pipe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/sendIORequestWithData:bytesTransferred:completionTimeout:error:
func (u_ USBHostPipe) SendIORequestWithDataBytesTransferredCompletionTimeoutError(data foundation.IMutableData, bytesTransferred unsafe.Pointer, completionTimeout foundation.ITimeInterval, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendIORequestWithData:bytesTransferred:completionTimeout:error:"), data, bytesTransferred, completionTimeout, error_)
	return rv
}


// The default completion timeout for input/output requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostdefaultcontrolcompletiontimeout
func (u_ USBHostPipe) IOUSBHostDefaultControlCompletionTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("IOUSBHostDefaultControlCompletionTimeout"))
	return rv
}


// A property that retrieves the current endpoint descriptors controlling the endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostpipe/descriptors
func (u_ USBHostPipe) Descriptors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("descriptors"))
	return rv
}


// A property that retrieves the current endpoint descriptors controlling the endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostpipe/descriptors
func (u_ USBHostPipe) SetDescriptors(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDescriptors:"), value)
}


// A property that retrieves the current idle suspend timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostpipe/idletimeout
func (u_ USBHostPipe) IdleTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("idleTimeout"))
	return rv
}


// A property that retrieves the current idle suspend timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostpipe/idletimeout
func (u_ USBHostPipe) SetIdleTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIdleTimeout:"), value)
}


// A property that retrieves the original endpoint descriptors from the pipe at the point of creation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostpipe/originaldescriptors
func (u_ USBHostPipe) OriginalDescriptors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("originalDescriptors"))
	return rv
}


// A property that retrieves the original endpoint descriptors from the pipe at the point of creation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostpipe/originaldescriptors
func (u_ USBHostPipe) SetOriginalDescriptors(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOriginalDescriptors:"), value)
}



