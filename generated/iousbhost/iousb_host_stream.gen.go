// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [USBHostStream] class.
type IUSBHostStream interface {
	IUSBHostIOSource
	// properties:
	StreamID() uint
	HostPipe() IOUSBHostPipe
	SetHostPipe(value IOUSBHostPipe)
	// methods:
	AbortWithError(error_ unsafe.Pointer) bool
	AbortWithOptionError(option unsafe.Pointer, error_ unsafe.Pointer) bool
	EnqueueIORequestWithDataErrorCompletionHandler(data foundation.MutableData, error_ unsafe.Pointer, completionHandler unsafe.Pointer) bool
}

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

// Alloc allocates a new instance without initialization.
func (uc _USBHostStreamClass) Alloc() USBHostStream {
	rv := objc.Send[USBHostStream](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Aborts pending input/output requests synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/abort()
func (u_ USBHostStream) AbortWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("abortWithError:"), error_)
	return rv
}


// Aborts pending input/output requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/abort(with:)
func (u_ USBHostStream) AbortWithOptionError(option unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("abortWithOption:error:"), option, error_)
	return rv
}


// Enqueues an input/output request on the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/enqueueIORequest(with:completionHandler:)
func (u_ USBHostStream) EnqueueIORequestWithDataErrorCompletionHandler(data foundation.MutableData, error_ unsafe.Pointer, completionHandler unsafe.Pointer) bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enqueueIORequestWithData:error:completionHandler:"), data, error_, completionHandler)
	return rv
}


// The ID for the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/streamID
func (u_ USBHostStream) StreamID() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("streamID"))
	return rv
}


// The pipe that creates the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhoststream/hostpipe
func (u_ USBHostStream) HostPipe() IOUSBHostPipe {
	rv := objc.Send[USBHostPipe](u_.ID, objc.Sel("hostPipe"))
	return rv
}


// The pipe that creates the stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/iousbhost/iousbhoststream/hostpipe
func (u_ USBHostStream) SetHostPipe(value IOUSBHostPipe) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHostPipe:"), value)
}




