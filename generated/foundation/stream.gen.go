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
	streamClass     _StreamClass
	streamClassOnce sync.Once
)

func getStreamClass() _StreamClass {
	streamClassOnce.Do(func() {
		streamClass = _StreamClass{objc.GetClass("NSStream")}
	})
	return streamClass
}

type _StreamClass struct {
	class objc.Class
}

// An interface definition for the [Stream] class.
type IStream interface {
	objectivec.IObject
}

// An abstract class representing a stream. [Full Topic]
//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/getStreamsToHost(withName:port:inputStream:outputStream:)
func (sc _StreamClass) GetStreamsToHostWithNamePortInputStreamOutputStream(hostname string, port int, inputStream unsafe.Pointer, outputStream unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getStreamsToHostWithName:port:inputStream:outputStream:"), objc.String(hostname), port, inputStream, outputStream)
}


