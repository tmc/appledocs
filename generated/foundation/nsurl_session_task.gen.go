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
	// properties:
	NSURLSessionTransferSizeUnknown() unsafe.Pointer
	CountOfBytesClientExpectsToReceive() unsafe.Pointer
	SetCountOfBytesClientExpectsToReceive(value unsafe.Pointer)
	CountOfBytesClientExpectsToSend() unsafe.Pointer
	SetCountOfBytesClientExpectsToSend(value unsafe.Pointer)
	CountOfBytesExpectedToReceive() unsafe.Pointer
	SetCountOfBytesExpectedToReceive(value unsafe.Pointer)
	CountOfBytesExpectedToSend() unsafe.Pointer
	SetCountOfBytesExpectedToSend(value unsafe.Pointer)
	CountOfBytesReceived() unsafe.Pointer
	SetCountOfBytesReceived(value unsafe.Pointer)
	CountOfBytesSent() unsafe.Pointer
	SetCountOfBytesSent(value unsafe.Pointer)
	CurrentRequest() IURLRequest
	SetCurrentRequest(value IURLRequest)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	EarliestBeginDate() IDate
	SetEarliestBeginDate(value IDate)
	Error() IError
	SetError(value IError)
	OriginalRequest() IURLRequest
	SetOriginalRequest(value IURLRequest)
	PrefersIncrementalDelivery() bool /* primitive/slice/pointer. */
	SetPrefersIncrementalDelivery(value bool /* primitive/slice/pointer. */)
	Priority() float32 /* primitive/slice/pointer. */
	SetPriority(value float32 /* primitive/slice/pointer. */)
	Progress() Progress /* not a class type */
	SetProgress(value Progress /* not a class type */)
	Response() IURLResponse
	SetResponse(value IURLResponse)
	TaskDescription() string /* primitive/slice/pointer. */
	SetTaskDescription(value string /* primitive/slice/pointer. */)
	TaskIdentifier() int /* primitive/slice/pointer. */
	SetTaskIdentifier(value int /* primitive/slice/pointer. */)
	// methods:
	Resume()
}

// A task, like downloading a specific resource, performed in a URL session.
//
// The class is the base class for tasks in a URL session. Tasks are always part of a session; you create a task by calling one of the task creation methods on a instance. The method you call determines the type of task. Use ‘s and related methods to create instances. Data tasks request a resource, returning the server’s response as one or more objects in memory. They are supported in default, ephemeral, and shared sessions, but are not supported in background sessions. Use ‘s and related methods to create instances. Upload tasks are like data tasks, except that they make it easier to provide a request body so you can upload data before retrieving the server’s response. Additionally, upload tasks are supported in background sessions. Use ’s and related methods to create instances. Download tasks download a resource directly to a file on disk. Download tasks are supported in any type of session. Use ’s or to create instances. Stream tasks establish a TCP/IP connection from a host name and port or a net service object. After you create a task, you start it by calling its method. The session then maintains a strong reference to the task until the request finishes or fails; you don’t need to maintain a reference to the task unless it’s useful for your app’s internal bookkeeping.


// A task, like downloading a specific resource, performed in a URL session.
//
// [Full Topic]
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



// Resumes the task, if it is suspended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/resume()
func (u_ URLSessionTask) Resume() {
	objc.Send[objc.ID](u_.ID, objc.Sel("resume"))
}


// The total size of the transfer cannot be determined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontransfersizeunknown
func (u_ URLSessionTask) NSURLSessionTransferSizeUnknown() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("NSURLSessionTransferSizeUnknown"))
	return rv
}


// A best-guess upper bound on the number of bytes the client expects to receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/countofbytesclientexpectstoreceive
func (u_ URLSessionTask) CountOfBytesClientExpectsToReceive() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfBytesClientExpectsToReceive"))
	return rv
}


// A best-guess upper bound on the number of bytes the client expects to receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/countofbytesclientexpectstoreceive
func (u_ URLSessionTask) SetCountOfBytesClientExpectsToReceive(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfBytesClientExpectsToReceive:"), value)
}


// A best-guess upper bound on the number of bytes the client expects to send.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/countofbytesclientexpectstosend
func (u_ URLSessionTask) CountOfBytesClientExpectsToSend() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfBytesClientExpectsToSend"))
	return rv
}


// A best-guess upper bound on the number of bytes the client expects to send.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/countofbytesclientexpectstosend
func (u_ URLSessionTask) SetCountOfBytesClientExpectsToSend(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfBytesClientExpectsToSend:"), value)
}


// The number of bytes that the task expects to receive in the response body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/countofbytesexpectedtoreceive
func (u_ URLSessionTask) CountOfBytesExpectedToReceive() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfBytesExpectedToReceive"))
	return rv
}


// The number of bytes that the task expects to receive in the response body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/countofbytesexpectedtoreceive
func (u_ URLSessionTask) SetCountOfBytesExpectedToReceive(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfBytesExpectedToReceive:"), value)
}


// The number of bytes that the task expects to send in the request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/countofbytesexpectedtosend
func (u_ URLSessionTask) CountOfBytesExpectedToSend() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfBytesExpectedToSend"))
	return rv
}


// The number of bytes that the task expects to send in the request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/countofbytesexpectedtosend
func (u_ URLSessionTask) SetCountOfBytesExpectedToSend(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfBytesExpectedToSend:"), value)
}


// The number of bytes that the task has received from the server in the response body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/countofbytesreceived
func (u_ URLSessionTask) CountOfBytesReceived() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfBytesReceived"))
	return rv
}


// The number of bytes that the task has received from the server in the response body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/countofbytesreceived
func (u_ URLSessionTask) SetCountOfBytesReceived(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfBytesReceived:"), value)
}


// The number of bytes that the task has sent to the server in the request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/countofbytessent
func (u_ URLSessionTask) CountOfBytesSent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfBytesSent"))
	return rv
}


// The number of bytes that the task has sent to the server in the request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/countofbytessent
func (u_ URLSessionTask) SetCountOfBytesSent(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfBytesSent:"), value)
}


// The URL request object currently being handled by the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/currentrequest
func (u_ URLSessionTask) CurrentRequest() IURLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("currentRequest"))
	return rv
}


// The URL request object currently being handled by the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/currentrequest
func (u_ URLSessionTask) SetCurrentRequest(value IURLRequest) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCurrentRequest:"), value)
}


// A delegate specific to the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/delegate
func (u_ URLSessionTask) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("delegate"))
	return rv
}


// A delegate specific to the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/delegate
func (u_ URLSessionTask) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDelegate:"), value)
}


// The earliest date at which the network load should begin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/earliestbegindate
func (u_ URLSessionTask) EarliestBeginDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("earliestBeginDate"))
	return rv
}


// The earliest date at which the network load should begin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/earliestbegindate
func (u_ URLSessionTask) SetEarliestBeginDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEarliestBeginDate:"), value)
}


// An error object that indicates why the task failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/error
func (u_ URLSessionTask) Error() IError {
	rv := objc.Send[Error](u_.ID, objc.Sel("error"))
	return rv
}


// An error object that indicates why the task failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/error
func (u_ URLSessionTask) SetError(value IError) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setError:"), value)
}


// The original request object passed when the task was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/originalrequest
func (u_ URLSessionTask) OriginalRequest() IURLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("originalRequest"))
	return rv
}


// The original request object passed when the task was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/originalrequest
func (u_ URLSessionTask) SetOriginalRequest(value IURLRequest) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setOriginalRequest:"), value)
}


// A Boolean value that determines whether to deliver a partial response body in increments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/prefersincrementaldelivery
func (u_ URLSessionTask) PrefersIncrementalDelivery() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("prefersIncrementalDelivery"))
	return rv
}


// A Boolean value that determines whether to deliver a partial response body in increments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/prefersincrementaldelivery
func (u_ URLSessionTask) SetPrefersIncrementalDelivery(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPrefersIncrementalDelivery:"), value)
}


// The relative priority at which you’d like a host to handle the task, specified as a floating point value between
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/priority
func (u_ URLSessionTask) Priority() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](u_.ID, objc.Sel("priority"))
	return rv
}


// The relative priority at which you’d like a host to handle the task, specified as a floating point value between
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/priority
func (u_ URLSessionTask) SetPriority(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPriority:"), value)
}


// A representation of the overall task progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/progress
func (u_ URLSessionTask) Progress() Progress /* not a class type */ {
	rv := objc.Send[Progress](u_.ID, objc.Sel("progress"))
	return rv
}


// A representation of the overall task progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/progress
func (u_ URLSessionTask) SetProgress(value Progress /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setProgress:"), value)
}


// The server’s response to the currently active request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/response
func (u_ URLSessionTask) Response() IURLResponse {
	rv := objc.Send[URLResponse](u_.ID, objc.Sel("response"))
	return rv
}


// The server’s response to the currently active request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/response
func (u_ URLSessionTask) SetResponse(value IURLResponse) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResponse:"), value)
}


// An app-provided string value for the current task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/taskdescription
func (u_ URLSessionTask) TaskDescription() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](u_.ID, objc.Sel("taskDescription"))
	return rv
}


// An app-provided string value for the current task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/taskdescription
func (u_ URLSessionTask) SetTaskDescription(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTaskDescription:"), objc.String(value))
}


// An identifier uniquely identifying the task within a given session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/taskidentifier
func (u_ URLSessionTask) TaskIdentifier() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](u_.ID, objc.Sel("taskIdentifier"))
	return rv
}


// An identifier uniquely identifying the task within a given session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/taskidentifier
func (u_ URLSessionTask) SetTaskIdentifier(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTaskIdentifier:"), value)
}



