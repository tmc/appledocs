// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLSession] class.
var URLSessionClass = _URLSessionClass{objc.GetClass("NSURLSession")}

type _URLSessionClass struct {
	class objc.Class
}

type URLSession struct {
	objc.ID
}

func URLSessionFrom(ptr unsafe.Pointer) URLSession {
	return URLSession{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _URLSessionClass) Alloc() URLSession {
	rv := objc.Send[URLSession](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (uc _URLSessionClass) New() URLSession {
	rv := objc.Send[URLSession](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSession) Init() URLSession {
	rv := objc.Send[URLSession](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSession) Autorelease() URLSession {
	rv := objc.Send[URLSession](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSession creates a new URLSession instance.
func NewURLSession() URLSession {
	return URLSessionClass.New()
}


// Creates a session with the specified session configuration. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/init(configuration:)
func (uc _URLSessionClass) SessionWithConfiguration(configuration unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("sessionWithConfiguration:"), configuration)
	return rv
}
// Creates a session with the specified session configuration, delegate, and operation queue. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/init(configuration:delegate:delegateQueue:)
func (uc _URLSessionClass) SessionWithConfigurationDelegateDelegateQueue(configuration unsafe.Pointer, delegate unsafe.Pointer, queue unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("sessionWithConfiguration:delegate:delegateQueue:"), configuration, delegate, queue)
	return rv
}
// Creates a task that retrieves the contents of the specified URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/dataTask(with:)-10dy7
func (u_ URLSession) DataTaskWithURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("dataTaskWithURL:"), url)
	return rv
}
// Creates a task that retrieves the contents of a URL based on the specified URL request object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/dataTask(with:)-7jpys
func (u_ URLSession) DataTaskWithRequest(request unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("dataTaskWithRequest:"), request)
	return rv
}
// Creates a task that retrieves the contents of the specified URL, then calls a handler upon completion. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/dataTask(with:completionHandler:)-52wk8
func (u_ URLSession) DataTaskWithURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("dataTaskWithURL:completionHandler:"), url, completionHandler)
	return rv
}
// Creates a task that retrieves the contents of a URL based on the specified URL request object, and calls a handler upon completion. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/dataTask(with:completionHandler:)-e6xv
func (u_ URLSession) DataTaskWithRequestCompletionHandler(request unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("dataTaskWithRequest:completionHandler:"), request, completionHandler)
	return rv
}
// Creates a download task that retrieves the contents of the specified URL and saves the results to a file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(with:)-1onj
func (u_ URLSession) DownloadTaskWithURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("downloadTaskWithURL:"), url)
	return rv
}
// Creates a download task that retrieves the contents of a URL based on the specified URL request object and saves the results to a file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(with:)-3fb7s
func (u_ URLSession) DownloadTaskWithRequest(request unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("downloadTaskWithRequest:"), request)
	return rv
}
// Creates a download task that retrieves the contents of a URL based on the specified URL request object, saves the results to a file, and calls a handler upon completion. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(with:completionHandler:)-4a84s
func (u_ URLSession) DownloadTaskWithRequestCompletionHandler(request unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("downloadTaskWithRequest:completionHandler:"), request, completionHandler)
	return rv
}
// Creates a download task that retrieves the contents of the specified URL, saves the results to a file, and calls a handler upon completion. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(with:completionHandler:)-7cuje
func (u_ URLSession) DownloadTaskWithURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("downloadTaskWithURL:completionHandler:"), url, completionHandler)
	return rv
}
// Creates a download task to resume a previously canceled or failed download. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(withResumeData:)
func (u_ URLSession) DownloadTaskWithResumeData(resumeData unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("downloadTaskWithResumeData:"), resumeData)
	return rv
}
// Creates a download task to resume a previously canceled or failed download and calls a handler upon completion. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(withResumeData:completionHandler:)
func (u_ URLSession) DownloadTaskWithResumeDataCompletionHandler(resumeData unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("downloadTaskWithResumeData:completionHandler:"), resumeData, completionHandler)
	return rv
}
// Invalidates the session, allowing any outstanding tasks to finish. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/finishTasksAndInvalidate()
func (u_ URLSession) FinishTasksAndInvalidate() {
	objc.Send[objc.ID](u_.ID, objc.Sel("finishTasksAndInvalidate"))
}
// Flushes cookies and credentials to disk, clears transient caches, and ensures that future requests occur on a new TCP connection. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/flush(completionHandler:)
func (u_ URLSession) FlushWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("flushWithCompletionHandler:"), completionHandler)
}
// Asynchronously calls a completion callback with all tasks in a session [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/getAllTasks(completionHandler:)
func (u_ URLSession) GetAllTasksWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getAllTasksWithCompletionHandler:"), completionHandler)
}
// Asynchronously calls a completion callback with all data, upload, and download tasks in a session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/getTasksWithCompletionHandler(_:)
func (u_ URLSession) GetTasksWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getTasksWithCompletionHandler:"), completionHandler)
}
// Cancels all outstanding tasks and then invalidates the session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/invalidateAndCancel()
func (u_ URLSession) InvalidateAndCancel() {
	objc.Send[objc.ID](u_.ID, objc.Sel("invalidateAndCancel"))
}
// Empties all cookies, caches and credential stores, removes disk files, flushes in-progress downloads to disk, and ensures that future requests occur on a new socket. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/reset(completionHandler:)
func (u_ URLSession) ResetWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("resetWithCompletionHandler:"), completionHandler)
}
// Creates a task that establishes a bidirectional TCP/IP connection using a specified network service. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/streamTask(with:)
func (u_ URLSession) StreamTaskWithNetService(service unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("streamTaskWithNetService:"), service)
	return rv
}
// Creates a task that establishes a bidirectional TCP/IP connection to a specified hostname and port. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/streamTask(withHostName:port:)
func (u_ URLSession) StreamTaskWithHostNamePort(hostname string, port int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("streamTaskWithHostName:port:"), hostname, port)
	return rv
}
// Creates a task that performs an HTTP request for the specified URL request object and uploads the provided data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(with:from:)
func (u_ URLSession) UploadTaskWithRequestFromData(request unsafe.Pointer, bodyData unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("uploadTaskWithRequest:fromData:"), request, bodyData)
	return rv
}
// Creates a task that performs an HTTP request for the specified URL request object, uploads the provided data, and calls a handler upon completion. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(with:from:completionHandler:)
func (u_ URLSession) UploadTaskWithRequestFromDataCompletionHandler(request unsafe.Pointer, bodyData unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("uploadTaskWithRequest:fromData:completionHandler:"), request, bodyData, completionHandler)
	return rv
}
// Creates a task that performs an HTTP request for uploading the specified file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(with:fromFile:)
func (u_ URLSession) UploadTaskWithRequestFromFile(request unsafe.Pointer, fileURL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("uploadTaskWithRequest:fromFile:"), request, fileURL)
	return rv
}
// Creates a task that performs an HTTP request for uploading the specified file, then calls a handler upon completion. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(with:fromFile:completionHandler:)
func (u_ URLSession) UploadTaskWithRequestFromFileCompletionHandler(request unsafe.Pointer, fileURL unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("uploadTaskWithRequest:fromFile:completionHandler:"), request, fileURL, completionHandler)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(withResumeData:)
func (u_ URLSession) UploadTaskWithResumeData(resumeData unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("uploadTaskWithResumeData:"), resumeData)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(withResumeData:completionHandler:)
func (u_ URLSession) UploadTaskWithResumeDataCompletionHandler(resumeData unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("uploadTaskWithResumeData:completionHandler:"), resumeData, completionHandler)
	return rv
}
// Creates a task that performs an HTTP request for uploading data based on the specified URL request. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(withStreamedRequest:)
func (u_ URLSession) UploadTaskWithStreamedRequest(request unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("uploadTaskWithStreamedRequest:"), request)
	return rv
}
// Creates a WebSocket task for the provided URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/webSocketTask(with:)-87ipz
func (u_ URLSession) WebSocketTaskWithURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("webSocketTaskWithURL:"), url)
	return rv
}
// Creates a WebSocket task for the provided URL request. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/webSocketTask(with:)-mtks
func (u_ URLSession) WebSocketTaskWithRequest(request unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("webSocketTaskWithRequest:"), request)
	return rv
}
// Creates a WebSocket task given a URL and an array of protocols. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/webSocketTask(with:protocols:)
func (u_ URLSession) WebSocketTaskWithURLProtocols(url unsafe.Pointer, protocols unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("webSocketTaskWithURL:protocols:"), url, protocols)
	return rv
}

