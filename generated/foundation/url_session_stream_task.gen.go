// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLSessionStreamTask] class.
var URLSessionStreamTaskClass = _URLSessionStreamTaskClass{objc.GetClass("NSURLSessionStreamTask")}

type _URLSessionStreamTaskClass struct {
	class objc.Class
}

type URLSessionStreamTask struct {
	objc.ID
}

func URLSessionStreamTaskFrom(ptr unsafe.Pointer) URLSessionStreamTask {
	return URLSessionStreamTask{
		ID: objc.ID(ptr),
	}
}


// Asynchronously writes the specified data to the stream, and calls a handler upon completion. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionStreamTask/write(_:timeout:completionHandler:)
func (u_ URLSessionStreamTask) WriteDataTimeoutCompletionHandler(data unsafe.Pointer, timeout TimeInterval, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("writeData:timeout:completionHandler:"), data, timeout, completionHandler)
}


