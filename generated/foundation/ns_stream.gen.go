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
	Close()
	Open()
	PropertyForKey(key unsafe.Pointer) objc.ID
	RemoveFromRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode)
	ScheduleInRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode)
	SetPropertyForKey(property objectivec.IObject, key unsafe.Pointer) bool
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	StreamError() NSError
	StreamStatus() StreamStatus
	NSStreamSOCKSErrorDomain() string
	NSStreamSocketSSLErrorDomain() string
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



// Creates and returns by reference a bound pair of input and output streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/getBoundStreams(withBufferSize:inputStream:outputStream:)

func (sc _StreamClass) GetBoundStreamsWithBufferSizeInputStreamOutputStream(bufferSize uint, inputStream IInputStream, outputStream IOutputStream) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getBoundStreamsWithBufferSize:inputStream:outputStream:"), bufferSize, inputStream, outputStream)
}


//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/getStreamsToHost(withName:port:inputStream:outputStream:)

func (sc _StreamClass) GetStreamsToHostWithNamePortInputStreamOutputStream(hostname string, port int, inputStream IInputStream, outputStream IOutputStream) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getStreamsToHostWithName:port:inputStream:outputStream:"), objc.String(hostname), port, inputStream, outputStream)
}


// Closes the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/close()

func (s_ Stream) Close() {
	objc.Send[objc.ID](s_.ID, objc.Sel("close"))
}


// Opens the receiving stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/open()

func (s_ Stream) Open() {
	objc.Send[objc.ID](s_.ID, objc.Sel("open"))
}


// Returns the receiver’s property for a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/property(forKey:)

func (s_ Stream) PropertyForKey(key unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("propertyForKey:"), key)
	return rv
}


// Removes the receiver from a given run loop running in a given mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/remove(from:forMode:)

func (s_ Stream) RemoveFromRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeFromRunLoop:forMode:"), aRunLoop, mode)
}


// Schedules the receiver on a given run loop in a given mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/schedule(in:forMode:)

func (s_ Stream) ScheduleInRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.Send[objc.ID](s_.ID, objc.Sel("scheduleInRunLoop:forMode:"), aRunLoop, mode)
}


// Attempts to set the value of a given property of the receiver and returns a Boolean value that indicates whether the value is accepted by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/setProperty(_:forKey:)

func (s_ Stream) SetPropertyForKey(property objectivec.IObject, key unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("setProperty:forKey:"), property, key)
	return rv
}


// Sets the receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/delegate

func (s_ Stream) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// Sets the receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/delegate

func (s_ Stream) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}


// Returns an object representing the stream error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/streamError

func (s_ Stream) StreamError() NSError {
	rv := objc.Send[NSError](s_.ID, objc.Sel("streamError"))
	return rv
}


// Returns the receiver’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/streamStatus

func (s_ Stream) StreamStatus() StreamStatus {
	rv := objc.Send[StreamStatus](s_.ID, objc.Sel("streamStatus"))
	return rv
}


// The error domain used by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamsockserrordomain

func (s_ Stream) NSStreamSOCKSErrorDomain() string {
	rv := objc.Send[string](s_.ID, objc.Sel("NSStreamSOCKSErrorDomain"))
	return rv
}


// The error domain used by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamsocketsslerrordomain

func (s_ Stream) NSStreamSocketSSLErrorDomain() string {
	rv := objc.Send[string](s_.ID, objc.Sel("NSStreamSocketSSLErrorDomain"))
	return rv
}



