// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Stream] class.
var streamClass = _StreamClass{objc.GetClass("NSStream")}

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

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/getStreamsToHost(withName:port:inputStream:outputStream:)
func (sc _StreamClass) GetStreamsToHostWithNamePortInputStreamOutputStream(hostname string, port int, inputStream unsafe.Pointer, outputStream unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getStreamsToHostWithName:port:inputStream:outputStream:"), hostname, port, inputStream, outputStream)
}


