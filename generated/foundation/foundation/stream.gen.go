// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Stream] class.
var StreamClass objc.Class

func init() {
	StreamClass = objc.GetClass("NSStream")
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
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Stream/getStreamsToHost(withName:port:inputStream:outputStream:)
func (sc Stream) GetStreamsToHostWithNamePortInputStreamOutputStream(hostname string, port int, inputStream unsafe.Pointer, outputStream unsafe.Pointer) {
	sel := objc.RegisterName("getStreamsToHostWithName:port:inputStream:outputStream:")
	objc.ID(StreamClass).Send(sel, hostname, port, inputStream, outputStream)
}


