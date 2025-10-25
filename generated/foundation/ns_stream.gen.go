// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSStream */


/* debug [class_header]: Header for NSStream */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Stream */
// An interface definition for the [Stream] class.
type IStream interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Stream */
	// properties:
	NSStreamSOCKSErrorDomain() IString
	NSStreamSocketSSLErrorDomain() IString
	StreamError() objc.IObject
	SetStreamError(value objc.IObject)
	StreamStatus() objectivec.IObject
	SetStreamStatus(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Stream */
	// methods:
	Open()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Stream */
// Alloc allocates a new instance without initialization.
func (sc _StreamClass) Alloc() Stream {
	rv := objc.Send[Stream](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Stream */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Stream *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Stream */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Stream */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Stream */

// Opens the receiving stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/open()
func (s_ Stream) Open() {
	objc.Send[objc.ID](s_.ID, objc.Sel("open"))
}/* debug [instance_methods/method]: Open */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Stream */

// The error domain used by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamsockserrordomain
func (s_ Stream) NSStreamSOCKSErrorDomain() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("NSStreamSOCKSErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: NSStreamSOCKSErrorDomain */


// The error domain used by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsstreamsocketsslerrordomain
func (s_ Stream) NSStreamSocketSSLErrorDomain() IString {
	rv := objc.Send[String](s_.ID, objc.Sel("NSStreamSocketSSLErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: NSStreamSocketSSLErrorDomain */


// Returns an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/stream/streamerror
func (s_ Stream) StreamError() objc.IObject {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("streamError"))
	return rv
}/* debug [instance_properties/getter]: streamError */


// Returns an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/stream/streamerror
func (s_ Stream) SetStreamError(value objc.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStreamError:"), value)
}/* debug [instance_properties/setter]: streamError */


// Returns the receiver’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/stream/streamstatus
func (s_ Stream) StreamStatus() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("streamStatus"))
	return rv
}/* debug [instance_properties/getter]: streamStatus */


// Returns the receiver’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/stream/streamstatus
func (s_ Stream) SetStreamStatus(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStreamStatus:"), value)
}/* debug [instance_properties/setter]: streamStatus */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSStream */



