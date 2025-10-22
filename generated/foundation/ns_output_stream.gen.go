// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [OutputStream] class.
var (
	OutputStreamClass     _OutputStreamClass
	OutputStreamClassOnce sync.Once
)

func getOutputStreamClass() _OutputStreamClass {
	OutputStreamClassOnce.Do(func() {
		OutputStreamClass = _OutputStreamClass{objc.GetClass("NSOutputStream")}
	})
	return OutputStreamClass
}

type _OutputStreamClass struct {
	class objc.Class
}

// An interface definition for the [OutputStream] class.
type IOutputStream interface {
	IStream
	WriteMaxLength(buffer unsafe.Pointer, len_ uint) int
	HasSpaceAvailable() bool
}

// A stream that provides write-only stream functionality.
//
// is “toll-free bridged” with its Core Foundation counterpart, . For more information on toll-free bridging, see .


// A stream that provides write-only stream functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OutputStream
type OutputStream struct {
	Stream
}

// OutputStreamFrom constructs a [OutputStream] from an unsafe.Pointer.
//
// A stream that provides write-only stream functionality.
func OutputStreamFrom(ptr unsafe.Pointer) OutputStream {
	return OutputStream{
		Stream: StreamFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (oc _OutputStreamClass) Alloc() OutputStream {
	rv := objc.Send[OutputStream](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OutputStreamClass) New() OutputStream {
	rv := objc.Send[OutputStream](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OutputStream) Init() OutputStream {
	rv := objc.Send[OutputStream](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OutputStream) Autorelease() OutputStream {
	rv := objc.Send[OutputStream](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOutputStream creates a new OutputStream instance.
func NewOutputStream() OutputStream {
	return getOutputStreamClass().New()
}



// Returns an initialized output stream that can write to a provided buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OutputStream/init(toBuffer:capacity:)
func NewOutputStreamToBufferCapacity(buffer unsafe.Pointer, capacity uint) OutputStream {
	instance := getOutputStreamClass().Alloc()
	rv := objc.Send[OutputStream](instance.ID, objc.Sel("initToBuffer:capacity:"), buffer, capacity)
	rv.Autorelease()
	return rv
}


// Returns an initialized output stream for writing to a specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OutputStream/init(toFileAtPath:append:)
func NewOutputStreamToFileAtPathAppend(path string, shouldAppend bool) OutputStream {
	instance := getOutputStreamClass().Alloc()
	rv := objc.Send[OutputStream](instance.ID, objc.Sel("initToFileAtPath:append:"), objc.String(path), shouldAppend)
	rv.Autorelease()
	return rv
}


// Returns an initialized output stream that will write to memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OutputStream/init(toMemory:)
func NewOutputStreamToMemory() OutputStream {
	instance := getOutputStreamClass().Alloc()
	rv := objc.Send[OutputStream](instance.ID, objc.Sel("initToMemory"))
	rv.Autorelease()
	return rv
}


// Returns an initialized output stream for writing to a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OutputStream/init(url:append:)-5soau
func NewOutputStreamWithURLAppend(url IURL, shouldAppend bool) OutputStream {
	instance := getOutputStreamClass().Alloc()
	rv := objc.Send[OutputStream](instance.ID, objc.Sel("initWithURL:append:"), url, shouldAppend)
	rv.Autorelease()
	return rv
}



// Creates and returns an initialized output stream that can write to a provided buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOutputStream/outputStreamToBuffer:capacity:
func (oc _OutputStreamClass) OutputStreamToBufferCapacity(buffer unsafe.Pointer, capacity uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("outputStreamToBuffer:capacity:"), buffer, capacity)
	return rv
}


// Creates and returns an initialized output stream for writing to a specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOutputStream/outputStreamToFileAtPath:append:
func (oc _OutputStreamClass) OutputStreamToFileAtPathAppend(path string, shouldAppend bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("outputStreamToFileAtPath:append:"), objc.String(path), shouldAppend)
	return rv
}


// Creates and returns an initialized output stream for writing to a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OutputStream/init(URL:append:)-8e5le
func (oc _OutputStreamClass) OutputStreamWithURLAppend(url IURL, shouldAppend bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("outputStreamWithURL:append:"), url, shouldAppend)
	return rv
}


// Creates and returns an initialized output stream that will write stream data to memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OutputStream/toMemory()
func (oc _OutputStreamClass) OutputStreamToMemory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("outputStreamToMemory"))
	return rv
}


// Writes the contents of a provided data buffer to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OutputStream/write(_:maxLength:)
func (o_ OutputStream) WriteMaxLength(buffer unsafe.Pointer, len_ uint) int {
	rv := objc.Send[int](o_.ID, objc.Sel("write:maxLength:"), buffer, len_)
	return rv
}


// A boolean value that indicates whether the receiver can be written to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OutputStream/hasSpaceAvailable
func (o_ OutputStream) HasSpaceAvailable() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("hasSpaceAvailable"))
	return rv
}


