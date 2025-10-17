// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLSession] class.
var URLSessionClass objc.Class

func init() {
	URLSessionClass = objc.GetClass("NSURLSession")
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
func (uc URLSession) Alloc() URLSession {
	ret := objc.ID(URLSessionClass).Send(objc.RegisterName("alloc"))
	return URLSession{ret}
}

// Init initializes the instance.
func (u_ URLSession) Init() URLSession {
	ret := u_.ID.Send(objc.RegisterName("init"))
	return URLSession{ret}
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/init()
func NewURLSession() URLSession {
	instance := URLSession{}.Alloc()
	instance = instance.Init()
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Creates a session with the specified session configuration. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/init(configuration:)
func (uc URLSession) SessionWithConfiguration(configuration unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("sessionWithConfiguration:")
	ret := objc.ID(URLSessionClass).Send(sel, configuration)
	return unsafe.Pointer(ret)
}
// Creates a session with the specified session configuration, delegate, and operation queue. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/init(configuration:delegate:delegateQueue:)
func (uc URLSession) SessionWithConfigurationDelegateDelegateQueue(configuration unsafe.Pointer, delegate unsafe.Pointer, queue unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("sessionWithConfiguration:delegate:delegateQueue:")
	ret := objc.ID(URLSessionClass).Send(sel, configuration, delegate, queue)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/new()
func (uc URLSession) New() unsafe.Pointer {
	sel := objc.RegisterName("new")
	ret := objc.ID(URLSessionClass).Send(sel)
	return unsafe.Pointer(ret)
}
// Creates a task that retrieves the contents of the specified URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/dataTask(with:)-10dy7
func (u_ URLSession) DataTaskWithURL(url unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataTaskWithURL:")
	ret := u_.ID.Send(sel, url)
	return unsafe.Pointer(ret)
}
// Creates a task that retrieves the contents of a URL based on the specified URL request object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/dataTask(with:)-7jpys
func (u_ URLSession) DataTaskWithRequest(request unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataTaskWithRequest:")
	ret := u_.ID.Send(sel, request)
	return unsafe.Pointer(ret)
}
// Creates a task that retrieves the contents of the specified URL, then calls a handler upon completion. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/dataTask(with:completionHandler:)-52wk8
func (u_ URLSession) DataTaskWithURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataTaskWithURL:completionHandler:")
	ret := u_.ID.Send(sel, url, completionHandler)
	return unsafe.Pointer(ret)
}
// Creates a task that retrieves the contents of a URL based on the specified URL request object, and calls a handler upon completion. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/dataTask(with:completionHandler:)-e6xv
func (u_ URLSession) DataTaskWithRequestCompletionHandler(request unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataTaskWithRequest:completionHandler:")
	ret := u_.ID.Send(sel, request, completionHandler)
	return unsafe.Pointer(ret)
}
// Creates a download task that retrieves the contents of the specified URL and saves the results to a file. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/downloadTask(with:)-1onj
func (u_ URLSession) DownloadTaskWithURL(url unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("downloadTaskWithURL:")
	ret := u_.ID.Send(sel, url)
	return unsafe.Pointer(ret)
}
// Creates a download task that retrieves the contents of a URL based on the specified URL request object and saves the results to a file. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/downloadTask(with:)-3fb7s
func (u_ URLSession) DownloadTaskWithRequest(request unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("downloadTaskWithRequest:")
	ret := u_.ID.Send(sel, request)
	return unsafe.Pointer(ret)
}
// Creates a download task that retrieves the contents of a URL based on the specified URL request object, saves the results to a file, and calls a handler upon completion. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/downloadTask(with:completionHandler:)-4a84s
func (u_ URLSession) DownloadTaskWithRequestCompletionHandler(request unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("downloadTaskWithRequest:completionHandler:")
	ret := u_.ID.Send(sel, request, completionHandler)
	return unsafe.Pointer(ret)
}
// Creates a download task that retrieves the contents of the specified URL, saves the results to a file, and calls a handler upon completion. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/downloadTask(with:completionHandler:)-7cuje
func (u_ URLSession) DownloadTaskWithURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("downloadTaskWithURL:completionHandler:")
	ret := u_.ID.Send(sel, url, completionHandler)
	return unsafe.Pointer(ret)
}
// Creates a download task to resume a previously canceled or failed download. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/downloadTask(withResumeData:)
func (u_ URLSession) DownloadTaskWithResumeData(resumeData unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("downloadTaskWithResumeData:")
	ret := u_.ID.Send(sel, resumeData)
	return unsafe.Pointer(ret)
}
// Creates a download task to resume a previously canceled or failed download and calls a handler upon completion. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/downloadTask(withResumeData:completionHandler:)
func (u_ URLSession) DownloadTaskWithResumeDataCompletionHandler(resumeData unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("downloadTaskWithResumeData:completionHandler:")
	ret := u_.ID.Send(sel, resumeData, completionHandler)
	return unsafe.Pointer(ret)
}
// Invalidates the session, allowing any outstanding tasks to finish. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/finishTasksAndInvalidate()
func (u_ URLSession) FinishTasksAndInvalidate() {
	sel := objc.RegisterName("finishTasksAndInvalidate")
	u_.ID.Send(sel)
}
// Flushes cookies and credentials to disk, clears transient caches, and ensures that future requests occur on a new TCP connection. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/flush(completionHandler:)
func (u_ URLSession) FlushWithCompletionHandler(completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("flushWithCompletionHandler:")
	u_.ID.Send(sel, completionHandler)
}
// Asynchronously calls a completion callback with all tasks in a session [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/getAllTasks(completionHandler:)
func (u_ URLSession) GetAllTasksWithCompletionHandler(completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("getAllTasksWithCompletionHandler:")
	u_.ID.Send(sel, completionHandler)
}
// Asynchronously calls a completion callback with all data, upload, and download tasks in a session. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/getTasksWithCompletionHandler(_:)
func (u_ URLSession) GetTasksWithCompletionHandler(completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("getTasksWithCompletionHandler:")
	u_.ID.Send(sel, completionHandler)
}
// Cancels all outstanding tasks and then invalidates the session. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/invalidateAndCancel()
func (u_ URLSession) InvalidateAndCancel() {
	sel := objc.RegisterName("invalidateAndCancel")
	u_.ID.Send(sel)
}
// Empties all cookies, caches and credential stores, removes disk files, flushes in-progress downloads to disk, and ensures that future requests occur on a new socket. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/reset(completionHandler:)
func (u_ URLSession) ResetWithCompletionHandler(completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("resetWithCompletionHandler:")
	u_.ID.Send(sel, completionHandler)
}
// Creates a task that establishes a bidirectional TCP/IP connection using a specified network service. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/streamTask(with:)
func (u_ URLSession) StreamTaskWithNetService(service unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("streamTaskWithNetService:")
	ret := u_.ID.Send(sel, service)
	return unsafe.Pointer(ret)
}
// Creates a task that establishes a bidirectional TCP/IP connection to a specified hostname and port. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/streamTask(withHostName:port:)
func (u_ URLSession) StreamTaskWithHostNamePort(hostname string, port int) unsafe.Pointer {
	sel := objc.RegisterName("streamTaskWithHostName:port:")
	ret := u_.ID.Send(sel, hostname, port)
	return unsafe.Pointer(ret)
}
// Creates a task that performs an HTTP request for the specified URL request object and uploads the provided data. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/uploadTask(with:from:)
func (u_ URLSession) UploadTaskWithRequestFromData(request unsafe.Pointer, bodyData unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("uploadTaskWithRequest:fromData:")
	ret := u_.ID.Send(sel, request, bodyData)
	return unsafe.Pointer(ret)
}
// Creates a task that performs an HTTP request for the specified URL request object, uploads the provided data, and calls a handler upon completion. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/uploadTask(with:from:completionHandler:)
func (u_ URLSession) UploadTaskWithRequestFromDataCompletionHandler(request unsafe.Pointer, bodyData unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("uploadTaskWithRequest:fromData:completionHandler:")
	ret := u_.ID.Send(sel, request, bodyData, completionHandler)
	return unsafe.Pointer(ret)
}
// Creates a task that performs an HTTP request for uploading the specified file. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/uploadTask(with:fromFile:)
func (u_ URLSession) UploadTaskWithRequestFromFile(request unsafe.Pointer, fileURL unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("uploadTaskWithRequest:fromFile:")
	ret := u_.ID.Send(sel, request, fileURL)
	return unsafe.Pointer(ret)
}
// Creates a task that performs an HTTP request for uploading the specified file, then calls a handler upon completion. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/uploadTask(with:fromFile:completionHandler:)
func (u_ URLSession) UploadTaskWithRequestFromFileCompletionHandler(request unsafe.Pointer, fileURL unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("uploadTaskWithRequest:fromFile:completionHandler:")
	ret := u_.ID.Send(sel, request, fileURL, completionHandler)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/uploadTask(withResumeData:)
func (u_ URLSession) UploadTaskWithResumeData(resumeData unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("uploadTaskWithResumeData:")
	ret := u_.ID.Send(sel, resumeData)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/uploadTask(withResumeData:completionHandler:)
func (u_ URLSession) UploadTaskWithResumeDataCompletionHandler(resumeData unsafe.Pointer, completionHandler unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("uploadTaskWithResumeData:completionHandler:")
	ret := u_.ID.Send(sel, resumeData, completionHandler)
	return unsafe.Pointer(ret)
}
// Creates a task that performs an HTTP request for uploading data based on the specified URL request. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/uploadTask(withStreamedRequest:)
func (u_ URLSession) UploadTaskWithStreamedRequest(request unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("uploadTaskWithStreamedRequest:")
	ret := u_.ID.Send(sel, request)
	return unsafe.Pointer(ret)
}
// Creates a WebSocket task for the provided URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/webSocketTask(with:)-87ipz
func (u_ URLSession) WebSocketTaskWithURL(url unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("webSocketTaskWithURL:")
	ret := u_.ID.Send(sel, url)
	return unsafe.Pointer(ret)
}
// Creates a WebSocket task for the provided URL request. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/webSocketTask(with:)-mtks
func (u_ URLSession) WebSocketTaskWithRequest(request unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("webSocketTaskWithRequest:")
	ret := u_.ID.Send(sel, request)
	return unsafe.Pointer(ret)
}
// Creates a WebSocket task given a URL and an array of protocols. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSession/webSocketTask(with:protocols:)
func (u_ URLSession) WebSocketTaskWithURLProtocols(url unsafe.Pointer, protocols unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("webSocketTaskWithURL:protocols:")
	ret := u_.ID.Send(sel, url, protocols)
	return unsafe.Pointer(ret)
}

