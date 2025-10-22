// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLSession] class.
var (
	URLSessionClass     _URLSessionClass
	URLSessionClassOnce sync.Once
)

func getURLSessionClass() _URLSessionClass {
	URLSessionClassOnce.Do(func() {
		URLSessionClass = _URLSessionClass{objc.GetClass("NSURLSession")}
	})
	return URLSessionClass
}

type _URLSessionClass struct {
	class objc.Class
}

// An interface definition for the [URLSession] class.
type IURLSession interface {
	objectivec.IObject
	DataTaskWithURL(url IURL) URLSessionDataTask
	DataTaskWithRequest(request IURLRequest) URLSessionDataTask
	DataTaskWithURLCompletionHandler(url IURL, completionHandler unsafe.Pointer) URLSessionDataTask
	DataTaskWithRequestCompletionHandler(request IURLRequest, completionHandler unsafe.Pointer) URLSessionDataTask
	DownloadTaskWithURL(url IURL) URLSessionDownloadTask
	DownloadTaskWithRequest(request IURLRequest) URLSessionDownloadTask
	DownloadTaskWithRequestCompletionHandler(request IURLRequest, completionHandler unsafe.Pointer) URLSessionDownloadTask
	DownloadTaskWithURLCompletionHandler(url IURL, completionHandler unsafe.Pointer) URLSessionDownloadTask
	DownloadTaskWithResumeData(resumeData IData) URLSessionDownloadTask
	DownloadTaskWithResumeDataCompletionHandler(resumeData IData, completionHandler unsafe.Pointer) URLSessionDownloadTask
	FinishTasksAndInvalidate()
	FlushWithCompletionHandler(completionHandler unsafe.Pointer)
	GetAllTasksWithCompletionHandler(completionHandler unsafe.Pointer)
	GetTasksWithCompletionHandler(completionHandler unsafe.Pointer)
	InvalidateAndCancel()
	ResetWithCompletionHandler(completionHandler unsafe.Pointer)
	StreamTaskWithNetService(service INetService) URLSessionStreamTask
	StreamTaskWithHostNamePort(hostname string, port int) URLSessionStreamTask
	UploadTaskWithRequestFromData(request IURLRequest, bodyData IData) URLSessionUploadTask
	UploadTaskWithRequestFromDataCompletionHandler(request IURLRequest, bodyData IData, completionHandler unsafe.Pointer) URLSessionUploadTask
	UploadTaskWithRequestFromFile(request IURLRequest, fileURL IURL) URLSessionUploadTask
	UploadTaskWithRequestFromFileCompletionHandler(request IURLRequest, fileURL IURL, completionHandler unsafe.Pointer) URLSessionUploadTask
	UploadTaskWithResumeData(resumeData IData) URLSessionUploadTask
	UploadTaskWithResumeDataCompletionHandler(resumeData IData, completionHandler unsafe.Pointer) URLSessionUploadTask
	UploadTaskWithStreamedRequest(request IURLRequest) URLSessionUploadTask
	WebSocketTaskWithURL(url IURL) URLSessionWebSocketTask
	WebSocketTaskWithRequest(request IURLRequest) URLSessionWebSocketTask
	WebSocketTaskWithURLProtocols(url IURL, protocols []string) URLSessionWebSocketTask
	Configuration() NSURLSessionConfiguration
	Delegate() objc.ID
	DelegateQueue() NSOperationQueue
	SessionDescription() string
	SetSessionDescription(value string)
}

// An object that coordinates a group of related, network data transfer tasks.
//
// The class and related classes provide an API for downloading data from and uploading data to endpoints indicated by URLs. Your app can also use this API to perform background downloads when your app isn’t running or, in iOS, while your app is suspended. You can use the related and to support authentication and receive events like redirection and task completion. Your app creates one or more instances, each of which coordinates a group of related data-transfer tasks. For example, if you’re creating a web browser, your app might create one session per tab or window, or one session for interactive use and another for background downloads. Within each session, your app adds a series of tasks, each of which represents a request for a specific URL (following HTTP redirects, if necessary).


// An object that coordinates a group of related, network data transfer tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession

type URLSession struct {
	objectivec.Object
}

// URLSessionFrom constructs a [URLSession] from an unsafe.Pointer.
//
// An object that coordinates a group of related, network data transfer tasks.
func URLSessionFrom(ptr unsafe.Pointer) URLSession {
	return URLSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLSessionClass) Alloc() URLSession {
	rv := objc.Send[URLSession](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getURLSessionClass().New()
}





// Creates a session with the specified session configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/init(configuration:)

func NewURLSessionWithConfiguration(configuration IURLSessionConfiguration) URLSession {
	rv := objc.Send[URLSession](objc.ID(getURLSessionClass().class), objc.Sel("sessionWithConfiguration:"), configuration)
	return rv
}




// Creates a session with the specified session configuration, delegate, and operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/init(configuration:delegate:delegateQueue:)

func NewURLSessionWithConfigurationDelegateDelegateQueue(configuration IURLSessionConfiguration, delegate objectivec.IObject, queue IOperationQueue) URLSession {
	rv := objc.Send[URLSession](objc.ID(getURLSessionClass().class), objc.Sel("sessionWithConfiguration:delegate:delegateQueue:"), configuration, delegate, queue)
	return rv
}



// Creates a session with the specified session configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/init(configuration:)

func (uc _URLSessionClass) SessionWithConfiguration(configuration IURLSessionConfiguration) URLSession {
	rv := objc.Send[URLSession](objc.ID(uc.class), objc.Sel("sessionWithConfiguration:"), configuration)
	return rv
}


// Creates a session with the specified session configuration, delegate, and operation queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/init(configuration:delegate:delegateQueue:)

func (uc _URLSessionClass) SessionWithConfigurationDelegateDelegateQueue(configuration IURLSessionConfiguration, delegate objectivec.IObject, queue IOperationQueue) URLSession {
	rv := objc.Send[URLSession](objc.ID(uc.class), objc.Sel("sessionWithConfiguration:delegate:delegateQueue:"), configuration, delegate, queue)
	return rv
}


// The shared singleton session object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/shared

func (uc _URLSessionClass) SharedSession() URLSession {
	rv := objc.Send[NSURLSession](objc.ID(uc.class), objc.Sel("sharedSession"))
	return rv
}

// Creates a task that retrieves the contents of the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/dataTask(with:)-10dy7

func (u_ URLSession) DataTaskWithURL(url IURL) URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](u_.ID, objc.Sel("dataTaskWithURL:"), url)
	return rv
}


// Creates a task that retrieves the contents of a URL based on the specified URL request object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/dataTask(with:)-7jpys

func (u_ URLSession) DataTaskWithRequest(request IURLRequest) URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](u_.ID, objc.Sel("dataTaskWithRequest:"), request)
	return rv
}


// Creates a task that retrieves the contents of the specified URL, then calls a handler upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/dataTask(with:completionHandler:)-52wk8

func (u_ URLSession) DataTaskWithURLCompletionHandler(url IURL, completionHandler unsafe.Pointer) URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](u_.ID, objc.Sel("dataTaskWithURL:completionHandler:"), url, completionHandler)
	return rv
}


// Creates a task that retrieves the contents of a URL based on the specified URL request object, and calls a handler upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/dataTask(with:completionHandler:)-e6xv

func (u_ URLSession) DataTaskWithRequestCompletionHandler(request IURLRequest, completionHandler unsafe.Pointer) URLSessionDataTask {
	rv := objc.Send[URLSessionDataTask](u_.ID, objc.Sel("dataTaskWithRequest:completionHandler:"), request, completionHandler)
	return rv
}


// Creates a download task that retrieves the contents of the specified URL and saves the results to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(with:)-1onj

func (u_ URLSession) DownloadTaskWithURL(url IURL) URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](u_.ID, objc.Sel("downloadTaskWithURL:"), url)
	return rv
}


// Creates a download task that retrieves the contents of a URL based on the specified URL request object and saves the results to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(with:)-3fb7s

func (u_ URLSession) DownloadTaskWithRequest(request IURLRequest) URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](u_.ID, objc.Sel("downloadTaskWithRequest:"), request)
	return rv
}


// Creates a download task that retrieves the contents of a URL based on the specified URL request object, saves the results to a file, and calls a handler upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(with:completionHandler:)-4a84s

func (u_ URLSession) DownloadTaskWithRequestCompletionHandler(request IURLRequest, completionHandler unsafe.Pointer) URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](u_.ID, objc.Sel("downloadTaskWithRequest:completionHandler:"), request, completionHandler)
	return rv
}


// Creates a download task that retrieves the contents of the specified URL, saves the results to a file, and calls a handler upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(with:completionHandler:)-7cuje

func (u_ URLSession) DownloadTaskWithURLCompletionHandler(url IURL, completionHandler unsafe.Pointer) URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](u_.ID, objc.Sel("downloadTaskWithURL:completionHandler:"), url, completionHandler)
	return rv
}


// Creates a download task to resume a previously canceled or failed download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(withResumeData:)

func (u_ URLSession) DownloadTaskWithResumeData(resumeData IData) URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](u_.ID, objc.Sel("downloadTaskWithResumeData:"), resumeData)
	return rv
}


// Creates a download task to resume a previously canceled or failed download and calls a handler upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/downloadTask(withResumeData:completionHandler:)

func (u_ URLSession) DownloadTaskWithResumeDataCompletionHandler(resumeData IData, completionHandler unsafe.Pointer) URLSessionDownloadTask {
	rv := objc.Send[URLSessionDownloadTask](u_.ID, objc.Sel("downloadTaskWithResumeData:completionHandler:"), resumeData, completionHandler)
	return rv
}


// Invalidates the session, allowing any outstanding tasks to finish.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/finishTasksAndInvalidate()

func (u_ URLSession) FinishTasksAndInvalidate() {
	objc.Send[objc.ID](u_.ID, objc.Sel("finishTasksAndInvalidate"))
}


// Flushes cookies and credentials to disk, clears transient caches, and ensures that future requests occur on a new TCP connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/flush(completionHandler:)

func (u_ URLSession) FlushWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("flushWithCompletionHandler:"), completionHandler)
}


// Asynchronously calls a completion callback with all tasks in a session
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/getAllTasks(completionHandler:)

func (u_ URLSession) GetAllTasksWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getAllTasksWithCompletionHandler:"), completionHandler)
}


// Asynchronously calls a completion callback with all data, upload, and download tasks in a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/getTasksWithCompletionHandler(_:)

func (u_ URLSession) GetTasksWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("getTasksWithCompletionHandler:"), completionHandler)
}


// Cancels all outstanding tasks and then invalidates the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/invalidateAndCancel()

func (u_ URLSession) InvalidateAndCancel() {
	objc.Send[objc.ID](u_.ID, objc.Sel("invalidateAndCancel"))
}


// Empties all cookies, caches and credential stores, removes disk files, flushes in-progress downloads to disk, and ensures that future requests occur on a new socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/reset(completionHandler:)

func (u_ URLSession) ResetWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("resetWithCompletionHandler:"), completionHandler)
}


// Creates a task that establishes a bidirectional TCP/IP connection using a specified network service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/streamTask(with:)

func (u_ URLSession) StreamTaskWithNetService(service INetService) URLSessionStreamTask {
	rv := objc.Send[URLSessionStreamTask](u_.ID, objc.Sel("streamTaskWithNetService:"), service)
	return rv
}


// Creates a task that establishes a bidirectional TCP/IP connection to a specified hostname and port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/streamTask(withHostName:port:)

func (u_ URLSession) StreamTaskWithHostNamePort(hostname string, port int) URLSessionStreamTask {
	rv := objc.Send[URLSessionStreamTask](u_.ID, objc.Sel("streamTaskWithHostName:port:"), objc.String(hostname), port)
	return rv
}


// Creates a task that performs an HTTP request for the specified URL request object and uploads the provided data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(with:from:)

func (u_ URLSession) UploadTaskWithRequestFromData(request IURLRequest, bodyData IData) URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](u_.ID, objc.Sel("uploadTaskWithRequest:fromData:"), request, bodyData)
	return rv
}


// Creates a task that performs an HTTP request for the specified URL request object, uploads the provided data, and calls a handler upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(with:from:completionHandler:)

func (u_ URLSession) UploadTaskWithRequestFromDataCompletionHandler(request IURLRequest, bodyData IData, completionHandler unsafe.Pointer) URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](u_.ID, objc.Sel("uploadTaskWithRequest:fromData:completionHandler:"), request, bodyData, completionHandler)
	return rv
}


// Creates a task that performs an HTTP request for uploading the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(with:fromFile:)

func (u_ URLSession) UploadTaskWithRequestFromFile(request IURLRequest, fileURL IURL) URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](u_.ID, objc.Sel("uploadTaskWithRequest:fromFile:"), request, fileURL)
	return rv
}


// Creates a task that performs an HTTP request for uploading the specified file, then calls a handler upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(with:fromFile:completionHandler:)

func (u_ URLSession) UploadTaskWithRequestFromFileCompletionHandler(request IURLRequest, fileURL IURL, completionHandler unsafe.Pointer) URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](u_.ID, objc.Sel("uploadTaskWithRequest:fromFile:completionHandler:"), request, fileURL, completionHandler)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(withResumeData:)

func (u_ URLSession) UploadTaskWithResumeData(resumeData IData) URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](u_.ID, objc.Sel("uploadTaskWithResumeData:"), resumeData)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(withResumeData:completionHandler:)

func (u_ URLSession) UploadTaskWithResumeDataCompletionHandler(resumeData IData, completionHandler unsafe.Pointer) URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](u_.ID, objc.Sel("uploadTaskWithResumeData:completionHandler:"), resumeData, completionHandler)
	return rv
}


// Creates a task that performs an HTTP request for uploading data based on the specified URL request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/uploadTask(withStreamedRequest:)

func (u_ URLSession) UploadTaskWithStreamedRequest(request IURLRequest) URLSessionUploadTask {
	rv := objc.Send[URLSessionUploadTask](u_.ID, objc.Sel("uploadTaskWithStreamedRequest:"), request)
	return rv
}


// Creates a WebSocket task for the provided URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/webSocketTask(with:)-87ipz

func (u_ URLSession) WebSocketTaskWithURL(url IURL) URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](u_.ID, objc.Sel("webSocketTaskWithURL:"), url)
	return rv
}


// Creates a WebSocket task for the provided URL request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/webSocketTask(with:)-mtks

func (u_ URLSession) WebSocketTaskWithRequest(request IURLRequest) URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](u_.ID, objc.Sel("webSocketTaskWithRequest:"), request)
	return rv
}


// Creates a WebSocket task given a URL and an array of protocols.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/webSocketTask(with:protocols:)

func (u_ URLSession) WebSocketTaskWithURLProtocols(url IURL, protocols []string) URLSessionWebSocketTask {
	rv := objc.Send[URLSessionWebSocketTask](u_.ID, objc.Sel("webSocketTaskWithURL:protocols:"), url, protocols)
	return rv
}


// A copy of the configuration object for this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/configuration

func (u_ URLSession) Configuration() NSURLSessionConfiguration {
	rv := objc.Send[NSURLSessionConfiguration](u_.ID, objc.Sel("configuration"))
	return rv
}


// The delegate assigned when this object was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/delegate

func (u_ URLSession) Delegate() objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("delegate"))
	return rv
}


// The operation queue provided when this object was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/delegateQueue

func (u_ URLSession) DelegateQueue() NSOperationQueue {
	rv := objc.Send[NSOperationQueue](u_.ID, objc.Sel("delegateQueue"))
	return rv
}


// An app-defined descriptive label for the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/sessionDescription

func (u_ URLSession) SessionDescription() string {
	rv := objc.Send[string](u_.ID, objc.Sel("sessionDescription"))
	return rv
}


// An app-defined descriptive label for the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/sessionDescription

func (u_ URLSession) SetSessionDescription(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSessionDescription:"), objc.String(value))
}


// The shared singleton session object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSession/shared

func (u_ URLSession) SharedSession() NSURLSession {
	rv := objc.Send[NSURLSession](u_.ID, objc.Sel("sharedSession"))
	return rv
}


