// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PURLConnectionDataDelegate is the NSURLConnectionDataDelegate protocol interface.
//
// A protocol that most delegates of a URL connection implement to receive data associated with the connection.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSURLConnectionDataDelegate
type PURLConnectionDataDelegate interface {
	// Optional methods
	ConnectionDidReceiveData(connection IURLConnection, data IData)
	HasConnectionDidReceiveData() bool
	ConnectionDidReceiveResponse(connection IURLConnection, response IURLResponse)
	HasConnectionDidReceiveResponse() bool
	ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite(connection IURLConnection, bytesWritten int, totalBytesWritten int, totalBytesExpectedToWrite int)
	HasConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite() bool
	ConnectionNeedNewBodyStream(connection IURLConnection, request IURLRequest) IInputStream
	HasConnectionNeedNewBodyStream() bool
	ConnectionWillCacheResponse(connection IURLConnection, cachedResponse ICachedURLResponse) ICachedURLResponse
	HasConnectionWillCacheResponse() bool
	ConnectionWillSendRequestRedirectResponse(connection IURLConnection, request IURLRequest, response IURLResponse) IURLRequest
	HasConnectionWillSendRequestRedirectResponse() bool
	ConnectionDidFinishLoading(connection IURLConnection)
	HasConnectionDidFinishLoading() bool
}

// URLConnectionDataDelegate is a delegate implementation builder for the PURLConnectionDataDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type URLConnectionDataDelegate struct {
	_ConnectionDidReceiveData func(connection IURLConnection, data IData)
	_ConnectionDidReceiveResponse func(connection IURLConnection, response IURLResponse)
	_ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite func(connection IURLConnection, bytesWritten int, totalBytesWritten int, totalBytesExpectedToWrite int)
	_ConnectionNeedNewBodyStream func(connection IURLConnection, request IURLRequest) IInputStream
	_ConnectionWillCacheResponse func(connection IURLConnection, cachedResponse ICachedURLResponse) ICachedURLResponse
	_ConnectionWillSendRequestRedirectResponse func(connection IURLConnection, request IURLRequest, response IURLResponse) IURLRequest
	_ConnectionDidFinishLoading func(connection IURLConnection)
}

// SetConnectionDidReceiveData sets the handler for the ConnectionDidReceiveData delegate method.
//
// Sent as a connection loads data incrementally.
func (d *URLConnectionDataDelegate) SetConnectionDidReceiveData(f func(connection IURLConnection, data IData)) {
	d._ConnectionDidReceiveData = f
}

// SetConnectionDidReceiveResponse sets the handler for the ConnectionDidReceiveResponse delegate method.
//
// Sent when the connection has received sufficient data to construct the URL response for its request.
func (d *URLConnectionDataDelegate) SetConnectionDidReceiveResponse(f func(connection IURLConnection, response IURLResponse)) {
	d._ConnectionDidReceiveResponse = f
}

// SetConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite sets the handler for the ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite delegate method.
//
// Sent as the body (message data) of a request is transmitted (such as in an HTTP POST request).
func (d *URLConnectionDataDelegate) SetConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite(f func(connection IURLConnection, bytesWritten int, totalBytesWritten int, totalBytesExpectedToWrite int)) {
	d._ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite = f
}

// SetConnectionNeedNewBodyStream sets the handler for the ConnectionNeedNewBodyStream delegate method.
//
// Called when an   needs to retransmit a request that has a body stream to provide a new, unopened stream.
func (d *URLConnectionDataDelegate) SetConnectionNeedNewBodyStream(f func(connection IURLConnection, request IURLRequest) IInputStream) {
	d._ConnectionNeedNewBodyStream = f
}

// SetConnectionWillCacheResponse sets the handler for the ConnectionWillCacheResponse delegate method.
//
// Sent before the connection stores a cached response in the cache, to give the delegate an opportunity to alter it.
func (d *URLConnectionDataDelegate) SetConnectionWillCacheResponse(f func(connection IURLConnection, cachedResponse ICachedURLResponse) ICachedURLResponse) {
	d._ConnectionWillCacheResponse = f
}

// SetConnectionWillSendRequestRedirectResponse sets the handler for the ConnectionWillSendRequestRedirectResponse delegate method.
//
// Sent when the connection determines that it must change URLs in order to continue loading a request.
func (d *URLConnectionDataDelegate) SetConnectionWillSendRequestRedirectResponse(f func(connection IURLConnection, request IURLRequest, response IURLResponse) IURLRequest) {
	d._ConnectionWillSendRequestRedirectResponse = f
}

// SetConnectionDidFinishLoading sets the handler for the ConnectionDidFinishLoading delegate method.
//
// Sent when a connection has finished loading successfully.
func (d *URLConnectionDataDelegate) SetConnectionDidFinishLoading(f func(connection IURLConnection)) {
	d._ConnectionDidFinishLoading = f
}

// ConnectionDidReceiveData implements the PURLConnectionDataDelegate interface.
func (d *URLConnectionDataDelegate) ConnectionDidReceiveData(connection IURLConnection, data IData) {
	if d._ConnectionDidReceiveData != nil {
		d._ConnectionDidReceiveData(connection, data)
	}
}

// HasConnectionDidReceiveData returns true if a handler for ConnectionDidReceiveData has been set.
func (d *URLConnectionDataDelegate) HasConnectionDidReceiveData() bool {
	return d._ConnectionDidReceiveData != nil
}

// ConnectionDidReceiveResponse implements the PURLConnectionDataDelegate interface.
func (d *URLConnectionDataDelegate) ConnectionDidReceiveResponse(connection IURLConnection, response IURLResponse) {
	if d._ConnectionDidReceiveResponse != nil {
		d._ConnectionDidReceiveResponse(connection, response)
	}
}

// HasConnectionDidReceiveResponse returns true if a handler for ConnectionDidReceiveResponse has been set.
func (d *URLConnectionDataDelegate) HasConnectionDidReceiveResponse() bool {
	return d._ConnectionDidReceiveResponse != nil
}

// ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite implements the PURLConnectionDataDelegate interface.
func (d *URLConnectionDataDelegate) ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite(connection IURLConnection, bytesWritten int, totalBytesWritten int, totalBytesExpectedToWrite int) {
	if d._ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite != nil {
		d._ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite(connection, bytesWritten, totalBytesWritten, totalBytesExpectedToWrite)
	}
}

// HasConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite returns true if a handler for ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite has been set.
func (d *URLConnectionDataDelegate) HasConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite() bool {
	return d._ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite != nil
}

// ConnectionNeedNewBodyStream implements the PURLConnectionDataDelegate interface.
func (d *URLConnectionDataDelegate) ConnectionNeedNewBodyStream(connection IURLConnection, request IURLRequest) IInputStream {
	if d._ConnectionNeedNewBodyStream != nil {
		return d._ConnectionNeedNewBodyStream(connection, request)
	}
	var zero IInputStream
	return zero
}

// HasConnectionNeedNewBodyStream returns true if a handler for ConnectionNeedNewBodyStream has been set.
func (d *URLConnectionDataDelegate) HasConnectionNeedNewBodyStream() bool {
	return d._ConnectionNeedNewBodyStream != nil
}

// ConnectionWillCacheResponse implements the PURLConnectionDataDelegate interface.
func (d *URLConnectionDataDelegate) ConnectionWillCacheResponse(connection IURLConnection, cachedResponse ICachedURLResponse) ICachedURLResponse {
	if d._ConnectionWillCacheResponse != nil {
		return d._ConnectionWillCacheResponse(connection, cachedResponse)
	}
	var zero ICachedURLResponse
	return zero
}

// HasConnectionWillCacheResponse returns true if a handler for ConnectionWillCacheResponse has been set.
func (d *URLConnectionDataDelegate) HasConnectionWillCacheResponse() bool {
	return d._ConnectionWillCacheResponse != nil
}

// ConnectionWillSendRequestRedirectResponse implements the PURLConnectionDataDelegate interface.
func (d *URLConnectionDataDelegate) ConnectionWillSendRequestRedirectResponse(connection IURLConnection, request IURLRequest, response IURLResponse) IURLRequest {
	if d._ConnectionWillSendRequestRedirectResponse != nil {
		return d._ConnectionWillSendRequestRedirectResponse(connection, request, response)
	}
	var zero IURLRequest
	return zero
}

// HasConnectionWillSendRequestRedirectResponse returns true if a handler for ConnectionWillSendRequestRedirectResponse has been set.
func (d *URLConnectionDataDelegate) HasConnectionWillSendRequestRedirectResponse() bool {
	return d._ConnectionWillSendRequestRedirectResponse != nil
}

// ConnectionDidFinishLoading implements the PURLConnectionDataDelegate interface.
func (d *URLConnectionDataDelegate) ConnectionDidFinishLoading(connection IURLConnection) {
	if d._ConnectionDidFinishLoading != nil {
		d._ConnectionDidFinishLoading(connection)
	}
}

// HasConnectionDidFinishLoading returns true if a handler for ConnectionDidFinishLoading has been set.
func (d *URLConnectionDataDelegate) HasConnectionDidFinishLoading() bool {
	return d._ConnectionDidFinishLoading != nil
}

// URLConnectionDataDelegateObject wraps an existing Objective-C object that conforms to the PURLConnectionDataDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type URLConnectionDataDelegateObject struct {
	objectivec.Object
}

// NewURLConnectionDataDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSURLConnectionDataDelegate protocol.
func NewURLConnectionDataDelegateObject(obj objectivec.Object) *URLConnectionDataDelegateObject {
	return &URLConnectionDataDelegateObject{obj}
}

// Make sure URLConnectionDataDelegateObject implements PURLConnectionDataDelegate.
var _ PURLConnectionDataDelegate = (*URLConnectionDataDelegateObject)(nil)

// ConnectionDidReceiveData implements the PURLConnectionDataDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLConnectionDataDelegateObject) ConnectionDidReceiveData(connection IURLConnection, data IData) {
	objc.Send[objc.ID](o.ID, objc.Sel("connection:didReceiveData:"), connection, data)
}

// HasConnectionDidReceiveData returns true; this is a placeholder for optional method checks.
func (o *URLConnectionDataDelegateObject) HasConnectionDidReceiveData() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ConnectionDidReceiveResponse implements the PURLConnectionDataDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLConnectionDataDelegateObject) ConnectionDidReceiveResponse(connection IURLConnection, response IURLResponse) {
	objc.Send[objc.ID](o.ID, objc.Sel("connection:didReceiveResponse:"), connection, response)
}

// HasConnectionDidReceiveResponse returns true; this is a placeholder for optional method checks.
func (o *URLConnectionDataDelegateObject) HasConnectionDidReceiveResponse() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite implements the PURLConnectionDataDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLConnectionDataDelegateObject) ConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite(connection IURLConnection, bytesWritten int, totalBytesWritten int, totalBytesExpectedToWrite int) {
	objc.Send[objc.ID](o.ID, objc.Sel("connection:didSendBodyData:totalBytesWritten:totalBytesExpectedToWrite:"), connection, bytesWritten, totalBytesWritten, totalBytesExpectedToWrite)
}

// HasConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite returns true; this is a placeholder for optional method checks.
func (o *URLConnectionDataDelegateObject) HasConnectionDidSendBodyDataTotalBytesWrittenTotalBytesExpectedToWrite() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ConnectionNeedNewBodyStream implements the PURLConnectionDataDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLConnectionDataDelegateObject) ConnectionNeedNewBodyStream(connection IURLConnection, request IURLRequest) IInputStream {
	return objc.Send[IInputStream](o.ID, objc.Sel("connection:needNewBodyStream:"), connection, request)
}

// HasConnectionNeedNewBodyStream returns true; this is a placeholder for optional method checks.
func (o *URLConnectionDataDelegateObject) HasConnectionNeedNewBodyStream() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ConnectionWillCacheResponse implements the PURLConnectionDataDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLConnectionDataDelegateObject) ConnectionWillCacheResponse(connection IURLConnection, cachedResponse ICachedURLResponse) ICachedURLResponse {
	return objc.Send[ICachedURLResponse](o.ID, objc.Sel("connection:willCacheResponse:"), connection, cachedResponse)
}

// HasConnectionWillCacheResponse returns true; this is a placeholder for optional method checks.
func (o *URLConnectionDataDelegateObject) HasConnectionWillCacheResponse() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ConnectionWillSendRequestRedirectResponse implements the PURLConnectionDataDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLConnectionDataDelegateObject) ConnectionWillSendRequestRedirectResponse(connection IURLConnection, request IURLRequest, response IURLResponse) IURLRequest {
	return objc.Send[IURLRequest](o.ID, objc.Sel("connection:willSendRequest:redirectResponse:"), connection, request, response)
}

// HasConnectionWillSendRequestRedirectResponse returns true; this is a placeholder for optional method checks.
func (o *URLConnectionDataDelegateObject) HasConnectionWillSendRequestRedirectResponse() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ConnectionDidFinishLoading implements the PURLConnectionDataDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLConnectionDataDelegateObject) ConnectionDidFinishLoading(connection IURLConnection) {
	objc.Send[objc.ID](o.ID, objc.Sel("connectionDidFinishLoading:"), connection)
}

// HasConnectionDidFinishLoading returns true; this is a placeholder for optional method checks.
func (o *URLConnectionDataDelegateObject) HasConnectionDidFinishLoading() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
