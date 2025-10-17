// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLSessionWebSocketMessage] class.
var uRLSessionWebSocketMessageClass = _URLSessionWebSocketMessageClass{objc.GetClass("NSURLSessionWebSocketMessage")}

type _URLSessionWebSocketMessageClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage

type URLSessionWebSocketMessage struct {
	objectivec.Object
}

// URLSessionWebSocketMessageFrom constructs a [URLSessionWebSocketMessage] from an unsafe.Pointer.
func URLSessionWebSocketMessageFrom(ptr unsafe.Pointer) URLSessionWebSocketMessage {
	return URLSessionWebSocketMessage{objectivec.Object{objc.ID(ptr)}}
}



