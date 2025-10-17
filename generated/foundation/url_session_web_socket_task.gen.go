// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLSessionWebSocketTask] class.
var uRLSessionWebSocketTaskClass = _URLSessionWebSocketTaskClass{objc.GetClass("NSURLSessionWebSocketTask")}

type _URLSessionWebSocketTaskClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionWebSocketTask] class.
type IURLSessionWebSocketTask interface {
	IURLSessionTask
}

// A URL session task that communicates over the WebSockets protocol standard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask

type URLSessionWebSocketTask struct {
	URLSessionTask
}

// URLSessionWebSocketTaskFrom constructs a [URLSessionWebSocketTask] from an unsafe.Pointer.
//
// A URL session task that communicates over the WebSockets protocol standard.
func URLSessionWebSocketTaskFrom(ptr unsafe.Pointer) URLSessionWebSocketTask {
	return URLSessionWebSocketTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}



