// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class IOUSBHostPipe */


/* debug [class_header]: Header for IOUSBHostPipe */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for USBHostPipe */
// An interface definition for the [USBHostPipe] class.
type IUSBHostPipe interface {
	IUSBHostIOSource
	
/* debug [class_interface_properties]: Properties for USBHostPipe */
	// properties:
	Descriptors() IOUSBHostIOSourceDescriptors
	IdleTimeout() float64
	OriginalDescriptors() IOUSBHostIOSourceDescriptors
	IOUSBHostDefaultControlCompletionTimeout() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for USBHostPipe */
	// methods:
	AbortWithError(error_ unsafe.Pointer) bool
	AbortWithOptionError(option USBHostAbortOption, error_ unsafe.Pointer) bool
	AdjustPipeWithDescriptorsError(descriptors USBHostIOSourceDescriptors, error_ unsafe.Pointer) bool
	ClearStallWithError(error_ unsafe.Pointer) bool
	CopyStreamWithStreamIDError(streamID uint, error_ unsafe.Pointer) IUSBHostStream
	DisableStreamsWithError(error_ unsafe.Pointer) bool
	EnableStreamsWithError(error_ unsafe.Pointer) bool
	EnqueueControlRequestDataCompletionTimeoutErrorCompletionHandler(request USBDeviceRequest /* not a class type */, data foundation.MutableData, completionTimeout float64, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool
	EnqueueControlRequestDataErrorCompletionHandler(request USBDeviceRequest /* not a class type */, data foundation.MutableData, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool
	EnqueueControlRequestErrorCompletionHandler(request USBDeviceRequest /* not a class type */, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool
	EnqueueIORequestWithDataCompletionTimeoutErrorCompletionHandler(data foundation.MutableData, completionTimeout float64, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool
	EnqueueIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsErrorCompletionHandler(data foundation.MutableData, transactionList USBHostIsochronousTransaction, transactionListCount uint, firstFrameNumber uint64, options USBHostIsochronousTransferOptions, error_ unsafe.Pointer, completionHandler USBHostIsochronousTransactionCompletionHandler /* not a class type */) bool
	SendControlRequestDataBytesTransferredCompletionTimeoutError(request USBDeviceRequest /* not a class type */, data foundation.MutableData, bytesTransferred uint, completionTimeout float64, error_ unsafe.Pointer) bool
	SendControlRequestDataBytesTransferredError(request USBDeviceRequest /* not a class type */, data foundation.MutableData, bytesTransferred uint, error_ unsafe.Pointer) bool
	SendControlRequestError(request USBDeviceRequest /* not a class type */, error_ unsafe.Pointer) bool
	SendIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsError(data foundation.MutableData, transactionList USBHostIsochronousTransaction, transactionListCount uint, firstFrameNumber uint64, options USBHostIsochronousTransferOptions, error_ unsafe.Pointer) bool
	SendIORequestWithDataBytesTransferredCompletionTimeoutError(data foundation.MutableData, bytesTransferred uint, completionTimeout float64, error_ unsafe.Pointer) bool
	SetIdleTimeoutError(idleTimeout float64, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for USBHostPipe */
// Alloc allocates a new instance without initialization.
func (uc _USBHostPipeClass) Alloc() USBHostPipe {
	rv := objc.Send[USBHostPipe](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for USBHostPipe */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for USBHostPipe *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for USBHostPipe */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for USBHostPipe */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for USBHostPipe */

// Aborts pending input/output requests synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/abortWithError:
func (u_ USBHostPipe) AbortWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("abortWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: AbortWithError */


// Aborts pending input/output requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/abortWithOption:error:
func (u_ USBHostPipe) AbortWithOptionError(option USBHostAbortOption, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("abortWithOption:error:"), option, error_)
	return rv
}/* debug [instance_methods/method]: AbortWithOptionError */


// Adjusts the behavior of periodic endpoints to consume a different amount of bus bandwidth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/adjust(with:)
func (u_ USBHostPipe) AdjustPipeWithDescriptorsError(descriptors USBHostIOSourceDescriptors, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("adjustPipeWithDescriptors:error:"), descriptors, error_)
	return rv
}/* debug [instance_methods/method]: AdjustPipeWithDescriptorsError */


// Clears the halt condition of the pipe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/clearStall()
func (u_ USBHostPipe) ClearStallWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("clearStallWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: ClearStallWithError */


// Returns the stream for a stream ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/copyStream(withStreamID:)
func (u_ USBHostPipe) CopyStreamWithStreamIDError(streamID uint, error_ unsafe.Pointer) IUSBHostStream {
	rv := objc.Send[USBHostStream](u_.ID, objc.Sel("copyStreamWithStreamID:error:"), streamID, error_)
	return rv
}/* debug [instance_methods/method]: CopyStreamWithStreamIDError */


// Disables streams for the pipe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/disableStreams()
func (u_ USBHostPipe) DisableStreamsWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("disableStreamsWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: DisableStreamsWithError */


// Enables streams for the pipe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/enableStreams()
func (u_ USBHostPipe) EnableStreamsWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enableStreamsWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: EnableStreamsWithError */


// Enqueues a request on a control endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/enqueueControlRequest:data:completionTimeout:error:completionHandler:
func (u_ USBHostPipe) EnqueueControlRequestDataCompletionTimeoutErrorCompletionHandler(request USBDeviceRequest /* not a class type */, data foundation.MutableData, completionTimeout float64, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueControlRequest:data:completionTimeout:error:completionHandler:"), request, data, completionTimeout, error_, completionHandler)
	return rv
}/* debug [instance_methods/method]: EnqueueControlRequestDataCompletionTimeoutErrorCompletionHandler */


// Enqueues a request on a control endpoint with a default completion timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/enqueueControlRequest:data:error:completionHandler:
func (u_ USBHostPipe) EnqueueControlRequestDataErrorCompletionHandler(request USBDeviceRequest /* not a class type */, data foundation.MutableData, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueControlRequest:data:error:completionHandler:"), request, data, error_, completionHandler)
	return rv
}/* debug [instance_methods/method]: EnqueueControlRequestDataErrorCompletionHandler */


// Enqueues a request on a control endpoint without a data phase and a default completion timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/enqueueControlRequest:error:completionHandler:
func (u_ USBHostPipe) EnqueueControlRequestErrorCompletionHandler(request USBDeviceRequest /* not a class type */, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueControlRequest:error:completionHandler:"), request, error_, completionHandler)
	return rv
}/* debug [instance_methods/method]: EnqueueControlRequestErrorCompletionHandler */


// Enqueues an input/output request on the pipe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/enqueueIORequest(with:completionTimeout:completionHandler:)
func (u_ USBHostPipe) EnqueueIORequestWithDataCompletionTimeoutErrorCompletionHandler(data foundation.MutableData, completionTimeout float64, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueIORequestWithData:completionTimeout:error:completionHandler:"), data, completionTimeout, error_, completionHandler)
	return rv
}/* debug [instance_methods/method]: EnqueueIORequestWithDataCompletionTimeoutErrorCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/enqueueIORequest(with:transactionList:transactionListCount:firstFrameNumber:options:completionHandler:)
func (u_ USBHostPipe) EnqueueIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsErrorCompletionHandler(data foundation.MutableData, transactionList USBHostIsochronousTransaction, transactionListCount uint, firstFrameNumber uint64, options USBHostIsochronousTransferOptions, error_ unsafe.Pointer, completionHandler USBHostIsochronousTransactionCompletionHandler /* not a class type */) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueIORequestWithData:transactionList:transactionListCount:firstFrameNumber:options:error:completionHandler:"), data, transactionList, transactionListCount, firstFrameNumber, options, error_, completionHandler)
	return rv
}/* debug [instance_methods/method]: EnqueueIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsErrorCompletionHandler */


// Sends a request on a control endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/sendControlRequest:data:bytesTransferred:completionTimeout:error:
func (u_ USBHostPipe) SendControlRequestDataBytesTransferredCompletionTimeoutError(request USBDeviceRequest /* not a class type */, data foundation.MutableData, bytesTransferred uint, completionTimeout float64, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendControlRequest:data:bytesTransferred:completionTimeout:error:"), request, data, bytesTransferred, completionTimeout, error_)
	return rv
}/* debug [instance_methods/method]: SendControlRequestDataBytesTransferredCompletionTimeoutError */


// Sends a request on a control endpoint with a default timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/sendControlRequest:data:bytesTransferred:error:
func (u_ USBHostPipe) SendControlRequestDataBytesTransferredError(request USBDeviceRequest /* not a class type */, data foundation.MutableData, bytesTransferred uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendControlRequest:data:bytesTransferred:error:"), request, data, bytesTransferred, error_)
	return rv
}/* debug [instance_methods/method]: SendControlRequestDataBytesTransferredError */


// Sends a request on a control endpoint without a data phase and a default completion timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/sendControlRequest:error:
func (u_ USBHostPipe) SendControlRequestError(request USBDeviceRequest /* not a class type */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendControlRequest:error:"), request, error_)
	return rv
}/* debug [instance_methods/method]: SendControlRequestError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/sendIORequest(with:transactionList:transactionListCount:firstFrameNumber:options:)
func (u_ USBHostPipe) SendIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsError(data foundation.MutableData, transactionList USBHostIsochronousTransaction, transactionListCount uint, firstFrameNumber uint64, options USBHostIsochronousTransferOptions, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendIORequestWithData:transactionList:transactionListCount:firstFrameNumber:options:error:"), data, transactionList, transactionListCount, firstFrameNumber, options, error_)
	return rv
}/* debug [instance_methods/method]: SendIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsError */


// Sends an input/output request on the pipe.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/sendIORequestWithData:bytesTransferred:completionTimeout:error:
func (u_ USBHostPipe) SendIORequestWithDataBytesTransferredCompletionTimeoutError(data foundation.MutableData, bytesTransferred uint, completionTimeout float64, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendIORequestWithData:bytesTransferred:completionTimeout:error:"), data, bytesTransferred, completionTimeout, error_)
	return rv
}/* debug [instance_methods/method]: SendIORequestWithDataBytesTransferredCompletionTimeoutError */


// Sets the desired idle suspend timeout for the interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/setIdleTimeout(_:)
func (u_ USBHostPipe) SetIdleTimeoutError(idleTimeout float64, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("setIdleTimeout:error:"), idleTimeout, error_)
	return rv
}/* debug [instance_methods/method]: SetIdleTimeoutError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for USBHostPipe */

// A property that retrieves the current endpoint descriptors controlling the endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/descriptors
func (u_ USBHostPipe) Descriptors() IOUSBHostIOSourceDescriptors {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("descriptors"))
	return rv
}/* debug [instance_properties/getter]: descriptors */


// A property that retrieves the current idle suspend timeout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/idleTimeout
func (u_ USBHostPipe) IdleTimeout() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("idleTimeout"))
	return rv
}/* debug [instance_properties/getter]: idleTimeout */


// A property that retrieves the original endpoint descriptors from the pipe at the point of creation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/originalDescriptors
func (u_ USBHostPipe) OriginalDescriptors() IOUSBHostIOSourceDescriptors {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("originalDescriptors"))
	return rv
}/* debug [instance_properties/getter]: originalDescriptors */


// The default completion timeout for input/output requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhostdefaultcontrolcompletiontimeout
func (u_ USBHostPipe) IOUSBHostDefaultControlCompletionTimeout() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("IOUSBHostDefaultControlCompletionTimeout"))
	return rv
}/* debug [instance_properties/getter]: IOUSBHostDefaultControlCompletionTimeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOUSBHostPipe */



