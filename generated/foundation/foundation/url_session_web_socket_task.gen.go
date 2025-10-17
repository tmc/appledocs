// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLSessionWebSocketTask] class.
var URLSessionWebSocketTaskClass objc.Class

func init() {
	URLSessionWebSocketTaskClass = objc.GetClass("NSURLSessionWebSocketTask")
}

type URLSessionWebSocketTask struct {
	objc.ID
}

func URLSessionWebSocketTaskFrom(ptr unsafe.Pointer) URLSessionWebSocketTask {
	return URLSessionWebSocketTask{
		ID: objc.ID(ptr),
	}
}




