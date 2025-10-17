// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLSessionWebSocketTask] class.
var URLSessionWebSocketTaskClass = _URLSessionWebSocketTaskClass{objc.GetClass("NSURLSessionWebSocketTask")}

type _URLSessionWebSocketTaskClass struct {
	class objc.Class
}

type URLSessionWebSocketTask struct {
	objc.ID
}

func URLSessionWebSocketTaskFrom(ptr unsafe.Pointer) URLSessionWebSocketTask {
	return URLSessionWebSocketTask{
		ID: objc.ID(ptr),
	}
}




