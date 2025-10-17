// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLSessionTask] class.
var uRLSessionTaskClass = _URLSessionTaskClass{objc.GetClass("NSURLSessionTask")}

type _URLSessionTaskClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionTask] class.
type IURLSessionTask interface {
	objectivec.IObject
	Cancel()
	Resume()
}

// A task, like downloading a specific resource, performed in a URL session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask

type URLSessionTask struct {
	objectivec.Object
}

// URLSessionTaskFrom constructs a [URLSessionTask] from an unsafe.Pointer.
//
// A task, like downloading a specific resource, performed in a URL session.
func URLSessionTaskFrom(ptr unsafe.Pointer) URLSessionTask {
	return URLSessionTask{objectivec.Object{objc.ID(ptr)}}
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


