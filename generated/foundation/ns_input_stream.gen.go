// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [InputStream] class.
var (
	InputStreamClass     _InputStreamClass
	InputStreamClassOnce sync.Once
)

func getInputStreamClass() _InputStreamClass {
	InputStreamClassOnce.Do(func() {
		InputStreamClass = _InputStreamClass{objc.GetClass("NSInputStream")}
	})
	return InputStreamClass
}

type _InputStreamClass struct {
	class objc.Class
}

// An interface definition for the [InputStream] class.
type IInputStream interface {
	IStream
	GetBufferLength(buffer unsafe.Pointer, len_ unsafe.Pointer) bool
	ReadMaxLength(buffer unsafe.Pointer, len_ uint) int
}

// A stream that provides read-only stream functionality.
//
// is “toll-free bridged” with its Core Foundation counterpart, . For more information on toll-free bridging, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InputStream
type InputStream struct {
	Stream
}

// InputStreamFrom constructs a [InputStream] from an unsafe.Pointer.
//
// A stream that provides read-only stream functionality.
func InputStreamFrom(ptr unsafe.Pointer) InputStream {
	return InputStream{
		Stream: StreamFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _InputStreamClass) Alloc() InputStream {
	rv := objc.Send[InputStream](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InputStreamClass) New() InputStream {
	rv := objc.Send[InputStream](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InputStream) Init() InputStream {
	rv := objc.Send[InputStream](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InputStream) Autorelease() InputStream {
	rv := objc.Send[InputStream](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInputStream creates a new InputStream instance.
func NewInputStream() InputStream {
	return getInputStreamClass().New()
}




// Initializes and returns an object for reading from a given object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InputStream/init(data:)
func NewInputStreamWithData(data unsafe.Pointer) InputStream {
	instance := getInputStreamClass().Alloc()
	rv := objc.Send[InputStream](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}



// Initializes and returns an object that reads data from the file at a given path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InputStream/init(fileAtPath:)
func NewInputStreamWithFileAtPath(path string) InputStream {
	instance := getInputStreamClass().Alloc()
	rv := objc.Send[InputStream](instance.ID, objc.Sel("initWithFileAtPath:"), objc.String(path))
	rv.Autorelease()
	return rv
}



// Initializes and returns an object that reads data from the file at a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InputStream/init(url:)-1lfmj
func NewInputStreamWithURL(url URL) InputStream {
	instance := getInputStreamClass().Alloc()
	rv := objc.Send[InputStream](instance.ID, objc.Sel("initWithURL:"), url)
	rv.Autorelease()
	return rv
}


// Creates and returns an initialized object that reads data from the file at a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InputStream/init(URL:)-y5k
func (ic _InputStreamClass) InputStreamWithURL(url URL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("inputStreamWithURL:"), url)
	return rv
}

// Creates and returns an initialized object for reading from a given object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInputStream/inputStreamWithData:
func (ic _InputStreamClass) InputStreamWithData(data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("inputStreamWithData:"), data)
	return rv
}

// Creates and returns an initialized object that reads data from the file at a given path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInputStream/inputStreamWithFileAtPath:
func (ic _InputStreamClass) InputStreamWithFileAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("inputStreamWithFileAtPath:"), objc.String(path))
	return rv
}

// Returns by reference a pointer to a read buffer and, by reference, the number of bytes available, and returns a Boolean value that indicates whether the buffer is available.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InputStream/getBuffer(_:length:)
func (i_ InputStream) GetBufferLength(buffer unsafe.Pointer, len_ unsafe.Pointer) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("getBuffer:length:"), buffer, len_)
	return rv
}

// Reads up to a given number of bytes into a given buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InputStream/read(_:maxLength:)
func (i_ InputStream) ReadMaxLength(buffer unsafe.Pointer, len_ uint) int {
	rv := objc.Send[int](i_.ID, objc.Sel("read:maxLength:"), buffer, len_)
	return rv
}

// A Boolean value that indicates whether the receiver has bytes available to read.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InputStream/hasBytesAvailable
func (i_ InputStream) HasBytesAvailable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasBytesAvailable"))
	return rv
}


