// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLSessionWebSocketMessage] class.
var URLSessionWebSocketMessageClass objc.Class

func init() {
	URLSessionWebSocketMessageClass = objc.GetClass("NSURLSessionWebSocketMessage")
}

type URLSessionWebSocketMessage struct {
	objc.ID
}

func URLSessionWebSocketMessageFrom(ptr unsafe.Pointer) URLSessionWebSocketMessage {
	return URLSessionWebSocketMessage{
		ID: objc.ID(ptr),
	}
}




