// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLSessionStreamTask] class.
var URLSessionStreamTaskClass objc.Class

func init() {
	URLSessionStreamTaskClass = objc.GetClass("NSURLSessionStreamTask")
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
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSessionStreamTask/write(_:timeout:completionHandler:)
func (u_ URLSessionStreamTask) WriteDataTimeoutCompletionHandler(data unsafe.Pointer, timeout foundation.TimeInterval, completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("writeData:timeout:completionHandler:")
	u_.ID.Send(sel, data, timeout, completionHandler)
}

