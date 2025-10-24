// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PURLSessionTaskDelegate is the NSURLSessionTaskDelegate protocol interface.
//
// A protocol that defines methods that URL session instances call on their delegates to handle task-level events.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/URLSessionTaskDelegate
type PURLSessionTaskDelegate interface {
	// Optional methods
	URLSessionDidCreateTask(session IURLSession, task IURLSessionTask)
	HasURLSessionDidCreateTask() bool
	URLSessionTaskDidCompleteWithError(session IURLSession, task IURLSessionTask, error_ IError)
	HasURLSessionTaskDidCompleteWithError() bool
	URLSessionTaskDidFinishCollectingMetrics(session IURLSession, task IURLSessionTask, metrics IURLSessionTaskMetrics)
	HasURLSessionTaskDidFinishCollectingMetrics() bool
	URLSessionTaskDidReceiveInformationalResponse(session IURLSession, task IURLSessionTask, response IHTTPURLResponse)
	HasURLSessionTaskDidReceiveInformationalResponse() bool
	URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend(session IURLSession, task IURLSessionTask, bytesSent int64, totalBytesSent int64, totalBytesExpectedToSend int64)
	HasURLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend() bool
	URLSessionTaskNeedNewBodyStream(session IURLSession, task IURLSessionTask, completionHandler unsafe.Pointer)
	HasURLSessionTaskNeedNewBodyStream() bool
	URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler(session IURLSession, task IURLSessionTask, offset int64, completionHandler unsafe.Pointer)
	HasURLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler() bool
	URLSessionTaskWillBeginDelayedRequestCompletionHandler(session IURLSession, task IURLSessionTask, request IURLRequest, completionHandler unsafe.Pointer)
	HasURLSessionTaskWillBeginDelayedRequestCompletionHandler() bool
	URLSessionTaskIsWaitingForConnectivity(session IURLSession, task IURLSessionTask)
	HasURLSessionTaskIsWaitingForConnectivity() bool
}

// URLSessionTaskDelegate is a delegate implementation builder for the PURLSessionTaskDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type URLSessionTaskDelegate struct {
	_URLSessionDidCreateTask func(session IURLSession, task IURLSessionTask)
	_URLSessionTaskDidCompleteWithError func(session IURLSession, task IURLSessionTask, error_ IError)
	_URLSessionTaskDidFinishCollectingMetrics func(session IURLSession, task IURLSessionTask, metrics IURLSessionTaskMetrics)
	_URLSessionTaskDidReceiveInformationalResponse func(session IURLSession, task IURLSessionTask, response IHTTPURLResponse)
	_URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend func(session IURLSession, task IURLSessionTask, bytesSent int64, totalBytesSent int64, totalBytesExpectedToSend int64)
	_URLSessionTaskNeedNewBodyStream func(session IURLSession, task IURLSessionTask, completionHandler unsafe.Pointer)
	_URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler func(session IURLSession, task IURLSessionTask, offset int64, completionHandler unsafe.Pointer)
	_URLSessionTaskWillBeginDelayedRequestCompletionHandler func(session IURLSession, task IURLSessionTask, request IURLRequest, completionHandler unsafe.Pointer)
	_URLSessionTaskIsWaitingForConnectivity func(session IURLSession, task IURLSessionTask)
}

// SetURLSessionDidCreateTask sets the handler for the URLSessionDidCreateTask delegate method.
func (d *URLSessionTaskDelegate) SetURLSessionDidCreateTask(f func(session IURLSession, task IURLSessionTask)) {
	d._URLSessionDidCreateTask = f
}

// SetURLSessionTaskDidCompleteWithError sets the handler for the URLSessionTaskDidCompleteWithError delegate method.
//
// Tells the delegate that the task finished transferring data.
func (d *URLSessionTaskDelegate) SetURLSessionTaskDidCompleteWithError(f func(session IURLSession, task IURLSessionTask, error_ IError)) {
	d._URLSessionTaskDidCompleteWithError = f
}

// SetURLSessionTaskDidFinishCollectingMetrics sets the handler for the URLSessionTaskDidFinishCollectingMetrics delegate method.
//
// Tells the delegate that the session finished collecting metrics for the task.
func (d *URLSessionTaskDelegate) SetURLSessionTaskDidFinishCollectingMetrics(f func(session IURLSession, task IURLSessionTask, metrics IURLSessionTaskMetrics)) {
	d._URLSessionTaskDidFinishCollectingMetrics = f
}

// SetURLSessionTaskDidReceiveInformationalResponse sets the handler for the URLSessionTaskDidReceiveInformationalResponse delegate method.
func (d *URLSessionTaskDelegate) SetURLSessionTaskDidReceiveInformationalResponse(f func(session IURLSession, task IURLSessionTask, response IHTTPURLResponse)) {
	d._URLSessionTaskDidReceiveInformationalResponse = f
}

// SetURLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend sets the handler for the URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend delegate method.
//
// Periodically informs the delegate of the progress of sending body content to the server.
func (d *URLSessionTaskDelegate) SetURLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend(f func(session IURLSession, task IURLSessionTask, bytesSent int64, totalBytesSent int64, totalBytesExpectedToSend int64)) {
	d._URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend = f
}

// SetURLSessionTaskNeedNewBodyStream sets the handler for the URLSessionTaskNeedNewBodyStream delegate method.
//
// Tells the delegate when a task requires a new request body stream to send to the remote server.
func (d *URLSessionTaskDelegate) SetURLSessionTaskNeedNewBodyStream(f func(session IURLSession, task IURLSessionTask, completionHandler unsafe.Pointer)) {
	d._URLSessionTaskNeedNewBodyStream = f
}

// SetURLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler sets the handler for the URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler delegate method.
func (d *URLSessionTaskDelegate) SetURLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler(f func(session IURLSession, task IURLSessionTask, offset int64, completionHandler unsafe.Pointer)) {
	d._URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler = f
}

// SetURLSessionTaskWillBeginDelayedRequestCompletionHandler sets the handler for the URLSessionTaskWillBeginDelayedRequestCompletionHandler delegate method.
//
// Tells the delegate that a delayed URL session task will now begin loading.
func (d *URLSessionTaskDelegate) SetURLSessionTaskWillBeginDelayedRequestCompletionHandler(f func(session IURLSession, task IURLSessionTask, request IURLRequest, completionHandler unsafe.Pointer)) {
	d._URLSessionTaskWillBeginDelayedRequestCompletionHandler = f
}

// SetURLSessionTaskIsWaitingForConnectivity sets the handler for the URLSessionTaskIsWaitingForConnectivity delegate method.
//
// Tells the delegate that the task is waiting until suitable connectivity is available before beginning the network load.
func (d *URLSessionTaskDelegate) SetURLSessionTaskIsWaitingForConnectivity(f func(session IURLSession, task IURLSessionTask)) {
	d._URLSessionTaskIsWaitingForConnectivity = f
}

// URLSessionDidCreateTask implements the PURLSessionTaskDelegate interface.
func (d *URLSessionTaskDelegate) URLSessionDidCreateTask(session IURLSession, task IURLSessionTask) {
	if d._URLSessionDidCreateTask != nil {
		d._URLSessionDidCreateTask(session, task)
	}
}

// HasURLSessionDidCreateTask returns true if a handler for URLSessionDidCreateTask has been set.
func (d *URLSessionTaskDelegate) HasURLSessionDidCreateTask() bool {
	return d._URLSessionDidCreateTask != nil
}

// URLSessionTaskDidCompleteWithError implements the PURLSessionTaskDelegate interface.
func (d *URLSessionTaskDelegate) URLSessionTaskDidCompleteWithError(session IURLSession, task IURLSessionTask, error_ IError) {
	if d._URLSessionTaskDidCompleteWithError != nil {
		d._URLSessionTaskDidCompleteWithError(session, task, error_)
	}
}

// HasURLSessionTaskDidCompleteWithError returns true if a handler for URLSessionTaskDidCompleteWithError has been set.
func (d *URLSessionTaskDelegate) HasURLSessionTaskDidCompleteWithError() bool {
	return d._URLSessionTaskDidCompleteWithError != nil
}

// URLSessionTaskDidFinishCollectingMetrics implements the PURLSessionTaskDelegate interface.
func (d *URLSessionTaskDelegate) URLSessionTaskDidFinishCollectingMetrics(session IURLSession, task IURLSessionTask, metrics IURLSessionTaskMetrics) {
	if d._URLSessionTaskDidFinishCollectingMetrics != nil {
		d._URLSessionTaskDidFinishCollectingMetrics(session, task, metrics)
	}
}

// HasURLSessionTaskDidFinishCollectingMetrics returns true if a handler for URLSessionTaskDidFinishCollectingMetrics has been set.
func (d *URLSessionTaskDelegate) HasURLSessionTaskDidFinishCollectingMetrics() bool {
	return d._URLSessionTaskDidFinishCollectingMetrics != nil
}

// URLSessionTaskDidReceiveInformationalResponse implements the PURLSessionTaskDelegate interface.
func (d *URLSessionTaskDelegate) URLSessionTaskDidReceiveInformationalResponse(session IURLSession, task IURLSessionTask, response IHTTPURLResponse) {
	if d._URLSessionTaskDidReceiveInformationalResponse != nil {
		d._URLSessionTaskDidReceiveInformationalResponse(session, task, response)
	}
}

// HasURLSessionTaskDidReceiveInformationalResponse returns true if a handler for URLSessionTaskDidReceiveInformationalResponse has been set.
func (d *URLSessionTaskDelegate) HasURLSessionTaskDidReceiveInformationalResponse() bool {
	return d._URLSessionTaskDidReceiveInformationalResponse != nil
}

// URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend implements the PURLSessionTaskDelegate interface.
func (d *URLSessionTaskDelegate) URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend(session IURLSession, task IURLSessionTask, bytesSent int64, totalBytesSent int64, totalBytesExpectedToSend int64) {
	if d._URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend != nil {
		d._URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend(session, task, bytesSent, totalBytesSent, totalBytesExpectedToSend)
	}
}

// HasURLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend returns true if a handler for URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend has been set.
func (d *URLSessionTaskDelegate) HasURLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend() bool {
	return d._URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend != nil
}

// URLSessionTaskNeedNewBodyStream implements the PURLSessionTaskDelegate interface.
func (d *URLSessionTaskDelegate) URLSessionTaskNeedNewBodyStream(session IURLSession, task IURLSessionTask, completionHandler unsafe.Pointer) {
	if d._URLSessionTaskNeedNewBodyStream != nil {
		d._URLSessionTaskNeedNewBodyStream(session, task, completionHandler)
	}
}

// HasURLSessionTaskNeedNewBodyStream returns true if a handler for URLSessionTaskNeedNewBodyStream has been set.
func (d *URLSessionTaskDelegate) HasURLSessionTaskNeedNewBodyStream() bool {
	return d._URLSessionTaskNeedNewBodyStream != nil
}

// URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler implements the PURLSessionTaskDelegate interface.
func (d *URLSessionTaskDelegate) URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler(session IURLSession, task IURLSessionTask, offset int64, completionHandler unsafe.Pointer) {
	if d._URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler != nil {
		d._URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler(session, task, offset, completionHandler)
	}
}

// HasURLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler returns true if a handler for URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler has been set.
func (d *URLSessionTaskDelegate) HasURLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler() bool {
	return d._URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler != nil
}

// URLSessionTaskWillBeginDelayedRequestCompletionHandler implements the PURLSessionTaskDelegate interface.
func (d *URLSessionTaskDelegate) URLSessionTaskWillBeginDelayedRequestCompletionHandler(session IURLSession, task IURLSessionTask, request IURLRequest, completionHandler unsafe.Pointer) {
	if d._URLSessionTaskWillBeginDelayedRequestCompletionHandler != nil {
		d._URLSessionTaskWillBeginDelayedRequestCompletionHandler(session, task, request, completionHandler)
	}
}

// HasURLSessionTaskWillBeginDelayedRequestCompletionHandler returns true if a handler for URLSessionTaskWillBeginDelayedRequestCompletionHandler has been set.
func (d *URLSessionTaskDelegate) HasURLSessionTaskWillBeginDelayedRequestCompletionHandler() bool {
	return d._URLSessionTaskWillBeginDelayedRequestCompletionHandler != nil
}

// URLSessionTaskIsWaitingForConnectivity implements the PURLSessionTaskDelegate interface.
func (d *URLSessionTaskDelegate) URLSessionTaskIsWaitingForConnectivity(session IURLSession, task IURLSessionTask) {
	if d._URLSessionTaskIsWaitingForConnectivity != nil {
		d._URLSessionTaskIsWaitingForConnectivity(session, task)
	}
}

// HasURLSessionTaskIsWaitingForConnectivity returns true if a handler for URLSessionTaskIsWaitingForConnectivity has been set.
func (d *URLSessionTaskDelegate) HasURLSessionTaskIsWaitingForConnectivity() bool {
	return d._URLSessionTaskIsWaitingForConnectivity != nil
}
