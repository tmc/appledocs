// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Stream] class.
var (
	StreamClass     _StreamClass
	StreamClassOnce sync.Once
)

func getStreamClass() _StreamClass {
	StreamClassOnce.Do(func() {
		StreamClass = _StreamClass{objc.GetClass("NSStream")}
	})
	return StreamClass
}

type _StreamClass struct {
	class objc.Class
}

// An interface definition for the [Stream] class.
type IStream interface {
	objectivec.IObject
	// properties:
	NSStreamSOCKSErrorDomain() IString
	NSStreamSocketSSLErrorDomain() IString
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	StreamError() objectivec.IObject
	SetStreamError(value objectivec.IObject)
	StreamStatus() unsafe.Pointer
	SetStreamStatus(value unsafe.Pointer)
	// methods:
	Open()
}

// An abstract class representing a stream.
//
// This class’s interface is common to all Cocoa stream classes, including its concrete subclasses and . objects provide an easy way to read and write data to and from a variety of media in a device-independent way. You can create stream objects for data located in memory, in a file, or on a network (using sockets), and you can use stream objects without loading all of the data into memory at once. By default, instances that aren’t file-based are non-seekable, one-way streams (although custom seekable subclasses are possible). After you provide or consume data, you can’t retrieve the data from the stream.


// An abstract class representing a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream
type Stream struct {
	objectivec.Object
}

// StreamFrom constructs a [Stream] from an unsafe.Pointer.
//
// An abstract class representing a stream.
func StreamFrom(ptr unsafe.Pointer) Stream {
	return Stream{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StreamClass) Alloc() Stream {
	rv := objc.Send[Stream](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StreamClass) New() Stream {
	rv := objc.Send[Stream](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Stream) Init() Stream {
	rv := objc.Send[Stream](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Stream) Autorelease() Stream {
	rv := objc.Send[Stream](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStream creates a new Stream instance.
func NewStream() Stream {
	return getStreamClass().New()
}



// Opens the receiving stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/open()
func (s_ Stream) Open() {
	objc.Send[objc.ID](s_.ID, objc.Sel("open"))
}


// The error domain used by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamsockserrordomain
func (s_ Stream) NSStreamSOCKSErrorDomain() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("NSStreamSOCKSErrorDomain"))
	return rv
}


// The error domain used by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamsocketsslerrordomain
func (s_ Stream) NSStreamSocketSSLErrorDomain() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("NSStreamSocketSSLErrorDomain"))
	return rv
}


// Sets the receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/stream/delegate
func (s_ Stream) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}


// Sets the receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/stream/delegate
func (s_ Stream) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}


// Returns an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/stream/streamerror
func (s_ Stream) StreamError() objectivec.IObject {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("streamError"))
	return rv
}


// Returns an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/stream/streamerror
func (s_ Stream) SetStreamError(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStreamError:"), value)
}


// Returns the receiver’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/stream/streamstatus
func (s_ Stream) StreamStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("streamStatus"))
	return rv
}


// Returns the receiver’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/stream/streamstatus
func (s_ Stream) SetStreamStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStreamStatus:"), value)
}



