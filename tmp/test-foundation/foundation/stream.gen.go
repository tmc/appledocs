// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var StreamClass _StreamClass

func init() {
	StreamClass = _StreamClass{objc.GetClass("NSStream")}
}

type _StreamClass struct {
	class objc.Class
}

type Stream struct {
	objc.ID
}

func StreamFrom(ptr unsafe.Pointer) Stream {
	return Stream{
		ID: objc.ID(ptr),
	}
}


//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Stream/getStreamsToHost(withName:port:inputStream:outputStream:)
func (sc _StreamClass) GetStreamsToHostWithNamePortInputStreamOutputStream(hostname string, port int, inputStream unsafe.Pointer, outputStream unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getStreamsToHostWithName:port:inputStream:outputStream:"), hostname, port, inputStream, outputStream)
}


