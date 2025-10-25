// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLSessionTask */


/* debug [class_header]: Header for NSURLSessionTask */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLSessionTask */
// An interface definition for the [URLSessionTask] class.
type IURLSessionTask interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URLSessionTask */
	// properties:
	CountOfBytesClientExpectsToReceive() int64
	SetCountOfBytesClientExpectsToReceive(value int64)
	CountOfBytesClientExpectsToSend() int64
	SetCountOfBytesClientExpectsToSend(value int64)
	CountOfBytesExpectedToReceive() int64
	CountOfBytesExpectedToSend() int64
	CountOfBytesReceived() int64
	CountOfBytesSent() int64
	CurrentRequest() IURLRequest
	EarliestBeginDate() IDate
	SetEarliestBeginDate(value IDate)
	Error() IError
	OriginalRequest() IURLRequest
	Priority() float32
	SetPriority(value float32)
	Progress() IProgress
	Response() IURLResponse
	State() URLSessionTaskState
	TaskDescription() IString
	SetTaskDescription(value IString)
	TaskIdentifier() uint
	NSURLSessionTransferSizeUnknown() objectivec.IObject
	PrefersIncrementalDelivery() bool
	SetPrefersIncrementalDelivery(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLSessionTask */
	// methods:
	Cancel()
	Resume()
	Suspend()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLSessionTask */
// Alloc allocates a new instance without initialization.
func (uc _URLSessionTaskClass) Alloc() URLSessionTask {
	rv := objc.Send[URLSessionTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLSessionTask */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLSessionTask *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLSessionTask */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLSessionTask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLSessionTask */

// Cancels the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/cancel()
func (u_ URLSessionTask) Cancel() {
	objc.Send[objc.ID](u_.ID, objc.Sel("cancel"))
}/* debug [instance_methods/method]: Cancel */


// Resumes the task, if it is suspended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/resume()
func (u_ URLSessionTask) Resume() {
	objc.Send[objc.ID](u_.ID, objc.Sel("resume"))
}/* debug [instance_methods/method]: Resume */


// Temporarily suspends a task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/suspend()
func (u_ URLSessionTask) Suspend() {
	objc.Send[objc.ID](u_.ID, objc.Sel("suspend"))
}/* debug [instance_methods/method]: Suspend */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLSessionTask */

// A best-guess upper bound on the number of bytes the client expects to receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesClientExpectsToReceive
func (u_ URLSessionTask) CountOfBytesClientExpectsToReceive() int64 {
	rv := objc.Send[int64](u_.ID, objc.Sel("countOfBytesClientExpectsToReceive"))
	return rv
}/* debug [instance_properties/getter]: countOfBytesClientExpectsToReceive */


// A best-guess upper bound on the number of bytes the client expects to receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesClientExpectsToReceive
func (u_ URLSessionTask) SetCountOfBytesClientExpectsToReceive(value int64) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfBytesClientExpectsToReceive:"), value)
}/* debug [instance_properties/setter]: countOfBytesClientExpectsToReceive */


// A best-guess upper bound on the number of bytes the client expects to send.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesClientExpectsToSend
func (u_ URLSessionTask) CountOfBytesClientExpectsToSend() int64 {
	rv := objc.Send[int64](u_.ID, objc.Sel("countOfBytesClientExpectsToSend"))
	return rv
}/* debug [instance_properties/getter]: countOfBytesClientExpectsToSend */


// A best-guess upper bound on the number of bytes the client expects to send.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesClientExpectsToSend
func (u_ URLSessionTask) SetCountOfBytesClientExpectsToSend(value int64) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfBytesClientExpectsToSend:"), value)
}/* debug [instance_properties/setter]: countOfBytesClientExpectsToSend */


// The number of bytes that the task expects to receive in the response body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesExpectedToReceive
func (u_ URLSessionTask) CountOfBytesExpectedToReceive() int64 {
	rv := objc.Send[int64](u_.ID, objc.Sel("countOfBytesExpectedToReceive"))
	return rv
}/* debug [instance_properties/getter]: countOfBytesExpectedToReceive */


// The number of bytes that the task expects to send in the request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesExpectedToSend
func (u_ URLSessionTask) CountOfBytesExpectedToSend() int64 {
	rv := objc.Send[int64](u_.ID, objc.Sel("countOfBytesExpectedToSend"))
	return rv
}/* debug [instance_properties/getter]: countOfBytesExpectedToSend */


// The number of bytes that the task has received from the server in the response body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesReceived
func (u_ URLSessionTask) CountOfBytesReceived() int64 {
	rv := objc.Send[int64](u_.ID, objc.Sel("countOfBytesReceived"))
	return rv
}/* debug [instance_properties/getter]: countOfBytesReceived */


// The number of bytes that the task has sent to the server in the request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/countOfBytesSent
func (u_ URLSessionTask) CountOfBytesSent() int64 {
	rv := objc.Send[int64](u_.ID, objc.Sel("countOfBytesSent"))
	return rv
}/* debug [instance_properties/getter]: countOfBytesSent */


// The URL request object currently being handled by the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/currentRequest
func (u_ URLSessionTask) CurrentRequest() IURLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("currentRequest"))
	return rv
}/* debug [instance_properties/getter]: currentRequest */


// The earliest date at which the network load should begin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/earliestBeginDate
func (u_ URLSessionTask) EarliestBeginDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("earliestBeginDate"))
	return rv
}/* debug [instance_properties/getter]: earliestBeginDate */


// The earliest date at which the network load should begin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/earliestBeginDate
func (u_ URLSessionTask) SetEarliestBeginDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEarliestBeginDate:"), value)
}/* debug [instance_properties/setter]: earliestBeginDate */


// An error object that indicates why the task failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/error
func (u_ URLSessionTask) Error() IError {
	rv := objc.Send[Error](u_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// The original request object passed when the task was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/originalRequest
func (u_ URLSessionTask) OriginalRequest() IURLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("originalRequest"))
	return rv
}/* debug [instance_properties/getter]: originalRequest */


// The relative priority at which you’d like a host to handle the task, specified as a floating point value between (lowest priority) and (highest priority).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/priority
func (u_ URLSessionTask) Priority() float32 {
	rv := objc.Send[float32](u_.ID, objc.Sel("priority"))
	return rv
}/* debug [instance_properties/getter]: priority */


// The relative priority at which you’d like a host to handle the task, specified as a floating point value between (lowest priority) and (highest priority).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/priority
func (u_ URLSessionTask) SetPriority(value float32) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPriority:"), value)
}/* debug [instance_properties/setter]: priority */


// A representation of the overall task progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/progress
func (u_ URLSessionTask) Progress() IProgress {
	rv := objc.Send[Progress](u_.ID, objc.Sel("progress"))
	return rv
}/* debug [instance_properties/getter]: progress */


// The server’s response to the currently active request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/response
func (u_ URLSessionTask) Response() IURLResponse {
	rv := objc.Send[URLResponse](u_.ID, objc.Sel("response"))
	return rv
}/* debug [instance_properties/getter]: response */


// The current state of the task—active, suspended, in the process of being canceled, or completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/state-swift.property
func (u_ URLSessionTask) State() URLSessionTaskState {
	rv := objc.Send[URLSessionTaskState](u_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// An app-provided string value for the current task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/taskDescription
func (u_ URLSessionTask) TaskDescription() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("taskDescription"))
	return rv
}/* debug [instance_properties/getter]: taskDescription */


// An app-provided string value for the current task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/taskDescription
func (u_ URLSessionTask) SetTaskDescription(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTaskDescription:"), value)
}/* debug [instance_properties/setter]: taskDescription */


// An identifier uniquely identifying the task within a given session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTask/taskIdentifier
func (u_ URLSessionTask) TaskIdentifier() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("taskIdentifier"))
	return rv
}/* debug [instance_properties/getter]: taskIdentifier */


// The total size of the transfer cannot be determined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessiontransfersizeunknown
func (u_ URLSessionTask) NSURLSessionTransferSizeUnknown() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("NSURLSessionTransferSizeUnknown"))
	return rv
}/* debug [instance_properties/getter]: NSURLSessionTransferSizeUnknown */


// A Boolean value that determines whether to deliver a partial response body in increments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/prefersincrementaldelivery
func (u_ URLSessionTask) PrefersIncrementalDelivery() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("prefersIncrementalDelivery"))
	return rv
}/* debug [instance_properties/getter]: prefersIncrementalDelivery */


// A Boolean value that determines whether to deliver a partial response body in increments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontask/prefersincrementaldelivery
func (u_ URLSessionTask) SetPrefersIncrementalDelivery(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPrefersIncrementalDelivery:"), value)
}/* debug [instance_properties/setter]: prefersIncrementalDelivery */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLSessionTask */



