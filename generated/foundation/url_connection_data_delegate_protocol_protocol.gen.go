// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
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
	ConnectionNeedNewBodyStream(connection IURLConnection, request IURLRequest) InputStream
	HasConnectionNeedNewBodyStream() bool
	ConnectionWillCacheResponse(connection IURLConnection, cachedResponse ICachedURLResponse) CachedURLResponse
	HasConnectionWillCacheResponse() bool
	ConnectionWillSendRequestRedirectResponse(connection IURLConnection, request IURLRequest, response IURLResponse) URLRequest
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
	_ConnectionNeedNewBodyStream func(connection IURLConnection, request IURLRequest) InputStream
	_ConnectionWillCacheResponse func(connection IURLConnection, cachedResponse ICachedURLResponse) CachedURLResponse
	_ConnectionWillSendRequestRedirectResponse func(connection IURLConnection, request IURLRequest, response IURLResponse) URLRequest
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
func (d *URLConnectionDataDelegate) SetConnectionNeedNewBodyStream(f func(connection IURLConnection, request IURLRequest) InputStream) {
	d._ConnectionNeedNewBodyStream = f
}

// SetConnectionWillCacheResponse sets the handler for the ConnectionWillCacheResponse delegate method.
//
// Sent before the connection stores a cached response in the cache, to give the delegate an opportunity to alter it.
func (d *URLConnectionDataDelegate) SetConnectionWillCacheResponse(f func(connection IURLConnection, cachedResponse ICachedURLResponse) CachedURLResponse) {
	d._ConnectionWillCacheResponse = f
}

// SetConnectionWillSendRequestRedirectResponse sets the handler for the ConnectionWillSendRequestRedirectResponse delegate method.
//
// Sent when the connection determines that it must change URLs in order to continue loading a request.
func (d *URLConnectionDataDelegate) SetConnectionWillSendRequestRedirectResponse(f func(connection IURLConnection, request IURLRequest, response IURLResponse) URLRequest) {
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
func (d *URLConnectionDataDelegate) ConnectionNeedNewBodyStream(connection IURLConnection, request IURLRequest) InputStream {
	if d._ConnectionNeedNewBodyStream != nil {
		return d._ConnectionNeedNewBodyStream(connection, request)
	}
	var zero InputStream
	return zero
}

// HasConnectionNeedNewBodyStream returns true if a handler for ConnectionNeedNewBodyStream has been set.
func (d *URLConnectionDataDelegate) HasConnectionNeedNewBodyStream() bool {
	return d._ConnectionNeedNewBodyStream != nil
}

// ConnectionWillCacheResponse implements the PURLConnectionDataDelegate interface.
func (d *URLConnectionDataDelegate) ConnectionWillCacheResponse(connection IURLConnection, cachedResponse ICachedURLResponse) CachedURLResponse {
	if d._ConnectionWillCacheResponse != nil {
		return d._ConnectionWillCacheResponse(connection, cachedResponse)
	}
	var zero CachedURLResponse
	return zero
}

// HasConnectionWillCacheResponse returns true if a handler for ConnectionWillCacheResponse has been set.
func (d *URLConnectionDataDelegate) HasConnectionWillCacheResponse() bool {
	return d._ConnectionWillCacheResponse != nil
}

// ConnectionWillSendRequestRedirectResponse implements the PURLConnectionDataDelegate interface.
func (d *URLConnectionDataDelegate) ConnectionWillSendRequestRedirectResponse(connection IURLConnection, request IURLRequest, response IURLResponse) URLRequest {
	if d._ConnectionWillSendRequestRedirectResponse != nil {
		return d._ConnectionWillSendRequestRedirectResponse(connection, request, response)
	}
	var zero URLRequest
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
