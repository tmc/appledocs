// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var URLSessionTaskClass _URLSessionTaskClass

func init() {
	URLSessionTaskClass = _URLSessionTaskClass{objc.GetClass("NSURLSessionTask")}
}

type _URLSessionTaskClass struct {
	class objc.Class
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/cancel()
func (u_ URLSessionTask) Cancel() {
	objc.Send[objc.ID](u_.ID, objc.Sel("cancel"))
}
// Resumes the task, if it is suspended. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/resume()
func (u_ URLSessionTask) Resume() {
	objc.Send[objc.ID](u_.ID, objc.Sel("resume"))
}


