// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLSessionTask] class.
var URLSessionTaskClass objc.Class

func init() {
	URLSessionTaskClass = objc.GetClass("NSURLSessionTask")
}

type URLSessionTask struct {
	objc.ID
}

func URLSessionTaskFrom(ptr unsafe.Pointer) URLSessionTask {
	return URLSessionTask{
		ID: objc.ID(ptr),
	}
}


// Cancels the task. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSessionTask/cancel()
func (u_ URLSessionTask) Cancel() {
	sel := objc.RegisterName("cancel")
	u_.ID.Send(sel)
}
// Resumes the task, if it is suspended. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSessionTask/resume()
func (u_ URLSessionTask) Resume() {
	sel := objc.RegisterName("resume")
	u_.ID.Send(sel)
}

