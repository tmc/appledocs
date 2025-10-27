// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [URLConnection] class.
var (
	URLConnectionClass     _URLConnectionClass
	URLConnectionClassOnce sync.Once
)

func getURLConnectionClass() _URLConnectionClass {
	URLConnectionClassOnce.Do(func() {
		URLConnectionClass = _URLConnectionClass{objc.GetClass("NSURLConnection")}
	})
	return URLConnectionClass
}

type _URLConnectionClass struct {
	class objc.Class
}





// An interface definition for the [URLConnection] class.
type IURLConnection interface {
	objectivec.IObject
	

	// properties:
	CurrentRequest() IURLRequest
	OriginalRequest() IURLRequest


	

	// methods:
	Cancel()
	ScheduleInRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode)
	SetDelegateQueue(queue IOperationQueue)
	Start()
	UnscheduleFromRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode)


}





// Alloc allocates a new instance without initialization.
func (uc _URLConnectionClass) Alloc() URLConnection {
	rv := objc.Send[URLConnection](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _URLConnectionClass) New() URLConnection {
	rv := objc.Send[URLConnection](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLConnection) Init() URLConnection {
	rv := objc.Send[URLConnection](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLConnection) Autorelease() URLConnection {
	rv := objc.Send[URLConnection](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLConnection creates a new URLConnection instance.
func NewURLConnection() URLConnection {
	return getURLConnectionClass().New()
}





// An object that enables you to start and stop URL requests.
//
// An object lets you load the contents of a URL by providing a URL request object. The interface for is sparse, providing only the controls to start and cancel asynchronous loads of a URL request. You perform most of your configuration on the URL request object itself. The class provides convenience class methods to load URL requests both asynchronously using a callback block and synchronously. For greater control, you can create a URL connection object with a delegate object that conforms to the and protocols. The connection calls methods on that delegate to provide you with progress and status as the URL request is loaded asynchronously. The connection also calls delegate methods to let you override the connection’s default behavior (for example, specifying how a particular redirect should be handled). These delegate methods are called on the thread that initiated the asynchronous load operation. For more information about errors, see the header, , and URL Loading System Error Codes in .


// An object that enables you to start and stop URL requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection
type URLConnection struct {
	objectivec.Object
}

// URLConnectionFrom constructs a [URLConnection] from an unsafe.Pointer.
//
// An object that enables you to start and stop URL requests.
func URLConnectionFrom(ptr unsafe.Pointer) URLConnection {
	return URLConnection{objectivec.Object{objc.ID(ptr)}}
}






// Returns an initialized URL connection and begins to load the data for the URL request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection/init(request:delegate:)
func NewURLConnectionWithRequestDelegate(request IURLRequest, delegate objectivec.IObject) URLConnection {
	instance := getURLConnectionClass().Alloc()
	rv := objc.Send[URLConnection](instance.ID, objc.Sel("initWithRequest:delegate:"), request, delegate)
	rv.Autorelease()
	return rv
}


// Returns an initialized URL connection and begins to load the data for the URL request, if specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection/init(request:delegate:startImmediately:)
func NewURLConnectionWithRequestDelegateStartImmediately(request IURLRequest, delegate objectivec.IObject, startImmediately bool) URLConnection {
	instance := getURLConnectionClass().Alloc()
	rv := objc.Send[URLConnection](instance.ID, objc.Sel("initWithRequest:delegate:startImmediately:"), request, delegate, startImmediately)
	rv.Autorelease()
	return rv
}







// Returns whether a request can be handled based on a preflight evaluation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection/canHandle(_:)
func (uc _URLConnectionClass) CanHandleRequest(request IURLRequest) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("canHandleRequest:"), request)
	return rv
}


// Creates and returns an initialized URL connection and begins to load the data for the URL request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection/connectionWithRequest:delegate:
func (uc _URLConnectionClass) ConnectionWithRequestDelegate(request IURLRequest, delegate objectivec.IObject) IURLConnection {
	rv := objc.Send[URLConnection](objc.ID(uc.class), objc.Sel("connectionWithRequest:delegate:"), request, delegate)
	return rv
}


// Loads the data for a URL request and executes a handler block on an operation queue when the request completes or fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection/sendAsynchronousRequest(_:queue:completionHandler:)
func (uc _URLConnectionClass) SendAsynchronousRequestQueueCompletionHandler(request IURLRequest, queue IOperationQueue, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("sendAsynchronousRequest:queue:completionHandler:"), request, queue, handler)
}


// Performs a synchronous load of the specified URL request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection/sendSynchronousRequest(_:returning:)
func (uc _URLConnectionClass) SendSynchronousRequestReturningResponseError(request IURLRequest, response IURLResponse, error_ IError) IData {
	rv := objc.Send[Data](objc.ID(uc.class), objc.Sel("sendSynchronousRequest:returningResponse:error:"), request, response, error_)
	return rv
}












// Cancels an asynchronous load of a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection/cancel()
func (u_ URLConnection) Cancel() {
	objc.Send[objc.ID](u_.ID, objc.Sel("cancel"))
}


// Determines the run loop and mode that the connection uses to call methods on its delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection/schedule(in:forMode:)
func (u_ URLConnection) ScheduleInRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.Send[objc.ID](u_.ID, objc.Sel("scheduleInRunLoop:forMode:"), aRunLoop, mode)
}


// Determines the operation queue that is used to call methods on the connection’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection/setDelegateQueue(_:)
func (u_ URLConnection) SetDelegateQueue(queue IOperationQueue) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDelegateQueue:"), queue)
}


// Causes the connection to begin loading data, if it has not already.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection/start()
func (u_ URLConnection) Start() {
	objc.Send[objc.ID](u_.ID, objc.Sel("start"))
}


// Causes the connection to stop calling delegate methods in the specified run loop and mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection/unschedule(from:forMode:)
func (u_ URLConnection) UnscheduleFromRunLoopForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.Send[objc.ID](u_.ID, objc.Sel("unscheduleFromRunLoop:forMode:"), aRunLoop, mode)
}







// The current connection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection/currentRequest
func (u_ URLConnection) CurrentRequest() IURLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("currentRequest"))
	return rv
}


// A deep copy of the original connection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLConnection/originalRequest
func (u_ URLConnection) OriginalRequest() IURLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("originalRequest"))
	return rv
}







