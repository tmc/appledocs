// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
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
	URLSessionTaskDidReceiveChallengeCompletionHandler(session IURLSession, task IURLSessionTask, challenge IURLAuthenticationChallenge, completionHandler unsafe.Pointer)
	HasURLSessionTaskDidReceiveChallengeCompletionHandler() bool
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
	URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler(session IURLSession, task IURLSessionTask, response IHTTPURLResponse, request IURLRequest, completionHandler unsafe.Pointer)
	HasURLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler() bool
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
	_URLSessionTaskDidReceiveChallengeCompletionHandler func(session IURLSession, task IURLSessionTask, challenge IURLAuthenticationChallenge, completionHandler unsafe.Pointer)
	_URLSessionTaskDidReceiveInformationalResponse func(session IURLSession, task IURLSessionTask, response IHTTPURLResponse)
	_URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend func(session IURLSession, task IURLSessionTask, bytesSent int64, totalBytesSent int64, totalBytesExpectedToSend int64)
	_URLSessionTaskNeedNewBodyStream func(session IURLSession, task IURLSessionTask, completionHandler unsafe.Pointer)
	_URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler func(session IURLSession, task IURLSessionTask, offset int64, completionHandler unsafe.Pointer)
	_URLSessionTaskWillBeginDelayedRequestCompletionHandler func(session IURLSession, task IURLSessionTask, request IURLRequest, completionHandler unsafe.Pointer)
	_URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler func(session IURLSession, task IURLSessionTask, response IHTTPURLResponse, request IURLRequest, completionHandler unsafe.Pointer)
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

// SetURLSessionTaskDidReceiveChallengeCompletionHandler sets the handler for the URLSessionTaskDidReceiveChallengeCompletionHandler delegate method.
//
// Requests credentials from the delegate in response to an authentication request from the remote server.
func (d *URLSessionTaskDelegate) SetURLSessionTaskDidReceiveChallengeCompletionHandler(f func(session IURLSession, task IURLSessionTask, challenge IURLAuthenticationChallenge, completionHandler unsafe.Pointer)) {
	d._URLSessionTaskDidReceiveChallengeCompletionHandler = f
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

// SetURLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler sets the handler for the URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler delegate method.
//
// Tells the delegate that the remote server requested an HTTP redirect.
func (d *URLSessionTaskDelegate) SetURLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler(f func(session IURLSession, task IURLSessionTask, response IHTTPURLResponse, request IURLRequest, completionHandler unsafe.Pointer)) {
	d._URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler = f
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

// URLSessionTaskDidReceiveChallengeCompletionHandler implements the PURLSessionTaskDelegate interface.
func (d *URLSessionTaskDelegate) URLSessionTaskDidReceiveChallengeCompletionHandler(session IURLSession, task IURLSessionTask, challenge IURLAuthenticationChallenge, completionHandler unsafe.Pointer) {
	if d._URLSessionTaskDidReceiveChallengeCompletionHandler != nil {
		d._URLSessionTaskDidReceiveChallengeCompletionHandler(session, task, challenge, completionHandler)
	}
}

// HasURLSessionTaskDidReceiveChallengeCompletionHandler returns true if a handler for URLSessionTaskDidReceiveChallengeCompletionHandler has been set.
func (d *URLSessionTaskDelegate) HasURLSessionTaskDidReceiveChallengeCompletionHandler() bool {
	return d._URLSessionTaskDidReceiveChallengeCompletionHandler != nil
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

// URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler implements the PURLSessionTaskDelegate interface.
func (d *URLSessionTaskDelegate) URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler(session IURLSession, task IURLSessionTask, response IHTTPURLResponse, request IURLRequest, completionHandler unsafe.Pointer) {
	if d._URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler != nil {
		d._URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler(session, task, response, request, completionHandler)
	}
}

// HasURLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler returns true if a handler for URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler has been set.
func (d *URLSessionTaskDelegate) HasURLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler() bool {
	return d._URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler != nil
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

// URLSessionTaskDelegateObject wraps an existing Objective-C object that conforms to the PURLSessionTaskDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type URLSessionTaskDelegateObject struct {
	objectivec.Object
}

// NewURLSessionTaskDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSURLSessionTaskDelegate protocol.
func NewURLSessionTaskDelegateObject(obj objectivec.Object) *URLSessionTaskDelegateObject {
	return &URLSessionTaskDelegateObject{obj}
}

// Make sure URLSessionTaskDelegateObject implements PURLSessionTaskDelegate.
var _ PURLSessionTaskDelegate = (*URLSessionTaskDelegateObject)(nil)

// URLSessionDidCreateTask implements the PURLSessionTaskDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionTaskDelegateObject) URLSessionDidCreateTask(session IURLSession, task IURLSessionTask) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:didCreateTask:"), session, task)
}

// HasURLSessionDidCreateTask returns true; this is a placeholder for optional method checks.
func (o *URLSessionTaskDelegateObject) HasURLSessionDidCreateTask() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionTaskDidCompleteWithError implements the PURLSessionTaskDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionTaskDelegateObject) URLSessionTaskDidCompleteWithError(session IURLSession, task IURLSessionTask, error_ IError) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:task:didCompleteWithError:"), session, task, error_)
}

// HasURLSessionTaskDidCompleteWithError returns true; this is a placeholder for optional method checks.
func (o *URLSessionTaskDelegateObject) HasURLSessionTaskDidCompleteWithError() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionTaskDidFinishCollectingMetrics implements the PURLSessionTaskDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionTaskDelegateObject) URLSessionTaskDidFinishCollectingMetrics(session IURLSession, task IURLSessionTask, metrics IURLSessionTaskMetrics) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:task:didFinishCollectingMetrics:"), session, task, metrics)
}

// HasURLSessionTaskDidFinishCollectingMetrics returns true; this is a placeholder for optional method checks.
func (o *URLSessionTaskDelegateObject) HasURLSessionTaskDidFinishCollectingMetrics() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionTaskDidReceiveChallengeCompletionHandler implements the PURLSessionTaskDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionTaskDelegateObject) URLSessionTaskDidReceiveChallengeCompletionHandler(session IURLSession, task IURLSessionTask, challenge IURLAuthenticationChallenge, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:task:didReceiveChallenge:completionHandler:"), session, task, challenge, completionHandler)
}

// HasURLSessionTaskDidReceiveChallengeCompletionHandler returns true; this is a placeholder for optional method checks.
func (o *URLSessionTaskDelegateObject) HasURLSessionTaskDidReceiveChallengeCompletionHandler() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionTaskDidReceiveInformationalResponse implements the PURLSessionTaskDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionTaskDelegateObject) URLSessionTaskDidReceiveInformationalResponse(session IURLSession, task IURLSessionTask, response IHTTPURLResponse) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:task:didReceiveInformationalResponse:"), session, task, response)
}

// HasURLSessionTaskDidReceiveInformationalResponse returns true; this is a placeholder for optional method checks.
func (o *URLSessionTaskDelegateObject) HasURLSessionTaskDidReceiveInformationalResponse() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend implements the PURLSessionTaskDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionTaskDelegateObject) URLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend(session IURLSession, task IURLSessionTask, bytesSent int64, totalBytesSent int64, totalBytesExpectedToSend int64) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:task:didSendBodyData:totalBytesSent:totalBytesExpectedToSend:"), session, task, bytesSent, totalBytesSent, totalBytesExpectedToSend)
}

// HasURLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend returns true; this is a placeholder for optional method checks.
func (o *URLSessionTaskDelegateObject) HasURLSessionTaskDidSendBodyDataTotalBytesSentTotalBytesExpectedToSend() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionTaskNeedNewBodyStream implements the PURLSessionTaskDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionTaskDelegateObject) URLSessionTaskNeedNewBodyStream(session IURLSession, task IURLSessionTask, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:task:needNewBodyStream:"), session, task, completionHandler)
}

// HasURLSessionTaskNeedNewBodyStream returns true; this is a placeholder for optional method checks.
func (o *URLSessionTaskDelegateObject) HasURLSessionTaskNeedNewBodyStream() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler implements the PURLSessionTaskDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionTaskDelegateObject) URLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler(session IURLSession, task IURLSessionTask, offset int64, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:task:needNewBodyStreamFromOffset:completionHandler:"), session, task, offset, completionHandler)
}

// HasURLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler returns true; this is a placeholder for optional method checks.
func (o *URLSessionTaskDelegateObject) HasURLSessionTaskNeedNewBodyStreamFromOffsetCompletionHandler() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionTaskWillBeginDelayedRequestCompletionHandler implements the PURLSessionTaskDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionTaskDelegateObject) URLSessionTaskWillBeginDelayedRequestCompletionHandler(session IURLSession, task IURLSessionTask, request IURLRequest, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:task:willBeginDelayedRequest:completionHandler:"), session, task, request, completionHandler)
}

// HasURLSessionTaskWillBeginDelayedRequestCompletionHandler returns true; this is a placeholder for optional method checks.
func (o *URLSessionTaskDelegateObject) HasURLSessionTaskWillBeginDelayedRequestCompletionHandler() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler implements the PURLSessionTaskDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionTaskDelegateObject) URLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler(session IURLSession, task IURLSessionTask, response IHTTPURLResponse, request IURLRequest, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:task:willPerformHTTPRedirection:newRequest:completionHandler:"), session, task, response, request, completionHandler)
}

// HasURLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler returns true; this is a placeholder for optional method checks.
func (o *URLSessionTaskDelegateObject) HasURLSessionTaskWillPerformHTTPRedirectionNewRequestCompletionHandler() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionTaskIsWaitingForConnectivity implements the PURLSessionTaskDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionTaskDelegateObject) URLSessionTaskIsWaitingForConnectivity(session IURLSession, task IURLSessionTask) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:taskIsWaitingForConnectivity:"), session, task)
}

// HasURLSessionTaskIsWaitingForConnectivity returns true; this is a placeholder for optional method checks.
func (o *URLSessionTaskDelegateObject) HasURLSessionTaskIsWaitingForConnectivity() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
