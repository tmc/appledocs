// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var uRLSessionWebSocketMessageClass _URLSessionWebSocketMessageClass

func init() {
	uRLSessionWebSocketMessageClass = _URLSessionWebSocketMessageClass{objc.GetClass("NSURLSessionWebSocketMessage")}
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




