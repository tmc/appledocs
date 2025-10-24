// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class IOUSBHostStream */


/* debug [class_header]: Header for IOUSBHostStream */
// The class instance for the [USBHostStream] class.
var (
	USBHostStreamClass     _USBHostStreamClass
	USBHostStreamClassOnce sync.Once
)

func getUSBHostStreamClass() _USBHostStreamClass {
	USBHostStreamClassOnce.Do(func() {
		USBHostStreamClass = _USBHostStreamClass{objc.GetClass("IOUSBHostStream")}
	})
	return USBHostStreamClass
}

type _USBHostStreamClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for USBHostStream */
// An interface definition for the [USBHostStream] class.
type IUSBHostStream interface {
	IUSBHostIOSource
	
/* debug [class_interface_properties]: Properties for USBHostStream */
	// properties:
	HostPipe() IOUSBHostPipe
	StreamID() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for USBHostStream */
	// methods:
	AbortWithError(error_ unsafe.Pointer) bool
	AbortWithOptionError(option USBHostAbortOption, error_ unsafe.Pointer) bool
	EnqueueIORequestWithDataErrorCompletionHandler(data foundation.MutableData, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool
	SendIORequestWithDataBytesTransferredError(data foundation.MutableData, bytesTransferred uint, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for USBHostStream */
// Alloc allocates a new instance without initialization.
func (uc _USBHostStreamClass) Alloc() USBHostStream {
	rv := objc.Send[USBHostStream](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _USBHostStreamClass) New() USBHostStream {
	rv := objc.Send[USBHostStream](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ USBHostStream) Init() USBHostStream {
	rv := objc.Send[USBHostStream](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ USBHostStream) Autorelease() USBHostStream {
	rv := objc.Send[USBHostStream](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUSBHostStream creates a new USBHostStream instance.
func NewUSBHostStream() USBHostStream {
	return getUSBHostStreamClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for USBHostStream */
// The class responsible for sending stream data for function drivers.
//
// The method creates stream objects.


// The class responsible for sending stream data for function drivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream
type USBHostStream struct {
	USBHostIOSource
}

// USBHostStreamFrom constructs a [USBHostStream] from an unsafe.Pointer.
//
// The class responsible for sending stream data for function drivers.
func USBHostStreamFrom(ptr unsafe.Pointer) USBHostStream {
	return USBHostStream{
		USBHostIOSource: USBHostIOSourceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for USBHostStream *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for USBHostStream */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for USBHostStream */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for USBHostStream */

// Aborts pending input/output requests synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/abort()
func (u_ USBHostStream) AbortWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("abortWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: AbortWithError */


// Aborts pending input/output requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/abort(with:)
func (u_ USBHostStream) AbortWithOptionError(option USBHostAbortOption, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("abortWithOption:error:"), option, error_)
	return rv
}/* debug [instance_methods/method]: AbortWithOptionError */


// Enqueues an input/output request on the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/enqueueIORequest(with:completionHandler:)
func (u_ USBHostStream) EnqueueIORequestWithDataErrorCompletionHandler(data foundation.MutableData, error_ unsafe.Pointer, completionHandler USBHostCompletionHandler /* not a class type */) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueIORequestWithData:error:completionHandler:"), data, error_, completionHandler)
	return rv
}/* debug [instance_methods/method]: EnqueueIORequestWithDataErrorCompletionHandler */


// Sends an input/output request on the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/sendIORequestWithData:bytesTransferred:error:
func (u_ USBHostStream) SendIORequestWithDataBytesTransferredError(data foundation.MutableData, bytesTransferred uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sendIORequestWithData:bytesTransferred:error:"), data, bytesTransferred, error_)
	return rv
}/* debug [instance_methods/method]: SendIORequestWithDataBytesTransferredError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for USBHostStream */

// The pipe that creates the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/hostPipe
func (u_ USBHostStream) HostPipe() IOUSBHostPipe {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("hostPipe"))
	return rv
}/* debug [instance_properties/getter]: hostPipe */


// The ID for the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/streamID
func (u_ USBHostStream) StreamID() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("streamID"))
	return rv
}/* debug [instance_properties/getter]: streamID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOUSBHostStream */






