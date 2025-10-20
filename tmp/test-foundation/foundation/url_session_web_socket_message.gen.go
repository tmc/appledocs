// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var URLSessionWebSocketMessageClass _URLSessionWebSocketMessageClass

func init() {
	URLSessionWebSocketMessageClass = _URLSessionWebSocketMessageClass{objc.GetClass("NSURLSessionWebSocketMessage")}
}

type _URLSessionWebSocketMessageClass struct {
	class objc.Class
}

type URLSessionWebSocketMessage struct {
	objc.ID
}

func URLSessionWebSocketMessageFrom(ptr unsafe.Pointer) URLSessionWebSocketMessage {
	return URLSessionWebSocketMessage{
		ID: objc.ID(ptr),
	}
}




