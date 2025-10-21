// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLSessionTask] class.
var (
	URLSessionTaskClass     _URLSessionTaskClass
	URLSessionTaskClassOnce sync.Once
)

func getURLSessionTaskClass() _URLSessionTaskClass {
	URLSessionTaskClassOnce.Do(func() {
		URLSessionTaskClass = _URLSessionTaskClass{objc.GetClass("NSURLSessionTask")}
	})
	return URLSessionTaskClass
}

type _URLSessionTaskClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionTask] class.
type IURLSessionTask interface {
	objectivec.IObject
	Cancel()
	Resume()
	Suspend()
}

// A task, like downloading a specific resource, performed in a URL session.
//
// The class is the base class for tasks in a URL session. Tasks are always part of a session; you create a task by calling one of the task creation methods on a instance. The method you call determines the type of task. Use ‘s and related methods to create instances. Data tasks request a resource, returning the server’s response as one or more objects in memory. They are supported in default, ephemeral, and shared sessions, but are not supported in background sessions. Use ‘s and related methods to create instances. Upload tasks are like data tasks, except that they make it easier to provide a request body so you can upload data before retrieving the server’s response. Additionally, upload tasks are supported in background sessions. Use ’s and related methods to create instances. Download tasks download a resource directly to a file on disk. Download tasks are supported in any type of session. Use ’s or to create instances. Stream tasks establish a TCP/IP connection from a host name and port or a net service object. After you create a task, you start it by calling its method. The session then maintains a strong reference to the task until the request finishes or fails; you don’t need to maintain a reference to the task unless it’s useful for your app’s internal bookkeeping.
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

// Alloc allocates a new instance without initialization.
func (uc _URLSessionTaskClass) Alloc() URLSessionTask {
	rv := objc.Send[URLSessionTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLSessionTaskClass) New() URLSessionTask {
	rv := objc.Send[URLSessionTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionTask) Init() URLSessionTask {
	rv := objc.Send[URLSessionTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionTask) Autorelease() URLSessionTask {
	rv := objc.Send[URLSessionTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionTask creates a new URLSessionTask instance.
func NewURLSessionTask() URLSessionTask {
	return getURLSessionTaskClass().New()
}



// Cancels the task.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/cancel()
func (u_ URLSessionTask) Cancel() {
	objc.Send[objc.ID](u_.ID, objc.Sel("cancel"))
}

// Resumes the task, if it is suspended.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/resume()
func (u_ URLSessionTask) Resume() {
	objc.Send[objc.ID](u_.ID, objc.Sel("resume"))
}

// Temporarily suspends a task.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/suspend()
func (u_ URLSessionTask) Suspend() {
	objc.Send[objc.ID](u_.ID, objc.Sel("suspend"))
}

// A best-guess upper bound on the number of bytes the client expects to receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesClientExpectsToReceive
func (u_ URLSessionTask) CountOfBytesClientExpectsToReceive() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfBytesClientExpectsToReceive"))
	return rv
}


// SetCountOfBytesClientExpectsToReceive sets the value of the countOfBytesClientExpectsToReceive property.
// A best-guess upper bound on the number of bytes the client expects to receive.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesClientExpectsToReceive
func (u_ URLSessionTask) SetCountOfBytesClientExpectsToReceive(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfBytesClientExpectsToReceive:"), value)
}

// A best-guess upper bound on the number of bytes the client expects to send.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesClientExpectsToSend
func (u_ URLSessionTask) CountOfBytesClientExpectsToSend() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfBytesClientExpectsToSend"))
	return rv
}


// SetCountOfBytesClientExpectsToSend sets the value of the countOfBytesClientExpectsToSend property.
// A best-guess upper bound on the number of bytes the client expects to send.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesClientExpectsToSend
func (u_ URLSessionTask) SetCountOfBytesClientExpectsToSend(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfBytesClientExpectsToSend:"), value)
}

// The number of bytes that the task expects to receive in the response body.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesExpectedToReceive
func (u_ URLSessionTask) CountOfBytesExpectedToReceive() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfBytesExpectedToReceive"))
	return rv
}

// The number of bytes that the task expects to send in the request body.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesExpectedToSend
func (u_ URLSessionTask) CountOfBytesExpectedToSend() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfBytesExpectedToSend"))
	return rv
}

// The number of bytes that the task has received from the server in the response body.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesReceived
func (u_ URLSessionTask) CountOfBytesReceived() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfBytesReceived"))
	return rv
}

// The number of bytes that the task has sent to the server in the request body.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesSent
func (u_ URLSessionTask) CountOfBytesSent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfBytesSent"))
	return rv
}

// The URL request object currently being handled by the task.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/currentRequest
func (u_ URLSessionTask) CurrentRequest() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("currentRequest"))
	return rv
}

// A delegate specific to the task.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/delegate
func (u_ URLSessionTask) Delegate() objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate specific to the task.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/delegate
func (u_ URLSessionTask) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDelegate:"), value)
}

// The earliest date at which the network load should begin.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/earliestBeginDate
func (u_ URLSessionTask) EarliestBeginDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("earliestBeginDate"))
	return rv
}


// SetEarliestBeginDate sets the value of the earliestBeginDate property.
// The earliest date at which the network load should begin.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/earliestBeginDate
func (u_ URLSessionTask) SetEarliestBeginDate(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEarliestBeginDate:"), value)
}

// An error object that indicates why the task failed.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/error
func (u_ URLSessionTask) Error() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("error"))
	return rv
}

// The original request object passed when the task was created.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/originalRequest
func (u_ URLSessionTask) OriginalRequest() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("originalRequest"))
	return rv
}

// A Boolean value that determines whether to deliver a partial response body in increments.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/prefersIncrementalDelivery
func (u_ URLSessionTask) PrefersIncrementalDelivery() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("prefersIncrementalDelivery"))
	return rv
}


// SetPrefersIncrementalDelivery sets the value of the prefersIncrementalDelivery property.
// A Boolean value that determines whether to deliver a partial response body in increments.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/prefersIncrementalDelivery
func (u_ URLSessionTask) SetPrefersIncrementalDelivery(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPrefersIncrementalDelivery:"), value)
}

// The relative priority at which you’d like a host to handle the task, specified as a floating point value between (lowest priority) and (highest priority).
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/priority
func (u_ URLSessionTask) Priority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("priority"))
	return rv
}


// SetPriority sets the value of the priority property.
// The relative priority at which you’d like a host to handle the task, specified as a floating point value between (lowest priority) and (highest priority).

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/priority
func (u_ URLSessionTask) SetPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPriority:"), value)
}

// A representation of the overall task progress.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/progress
func (u_ URLSessionTask) Progress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("progress"))
	return rv
}

// The server’s response to the currently active request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/response
func (u_ URLSessionTask) Response() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("response"))
	return rv
}

// The current state of the task—active, suspended, in the process of being canceled, or completed.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/state-swift.property
func (u_ URLSessionTask) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("state"))
	return rv
}

// An app-provided string value for the current task.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/taskDescription
func (u_ URLSessionTask) TaskDescription() string {
	rv := objc.Send[string](u_.ID, objc.Sel("taskDescription"))
	return rv
}


// SetTaskDescription sets the value of the taskDescription property.
// An app-provided string value for the current task.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/taskDescription
func (u_ URLSessionTask) SetTaskDescription(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTaskDescription:"), objc.String(value))
}

// An identifier uniquely identifying the task within a given session.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/taskIdentifier
func (u_ URLSessionTask) TaskIdentifier() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("taskIdentifier"))
	return rv
}

// The total size of the transfer cannot be determined.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontransfersizeunknown
func (u_ URLSessionTask) NSURLSessionTransferSizeUnknown() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("NSURLSessionTransferSizeUnknown"))
	return rv
}


