// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PURLSessionWebSocketDelegate is the NSURLSessionWebSocketDelegate protocol interface.
//
// A protocol that defines methods that URL session instances call on their delegates to handle task-level events specific to WebSocket tasks.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/URLSessionWebSocketDelegate
type PURLSessionWebSocketDelegate interface {
	// Optional methods
	URLSessionWebSocketTaskDidCloseWithCodeReason(session IURLSession, webSocketTask IURLSessionWebSocketTask, closeCode URLSessionWebSocketCloseCode, reason IData)
	HasURLSessionWebSocketTaskDidCloseWithCodeReason() bool
	URLSessionWebSocketTaskDidOpenWithProtocol(session IURLSession, webSocketTask IURLSessionWebSocketTask, protocol_ IString)
	HasURLSessionWebSocketTaskDidOpenWithProtocol() bool
}

// URLSessionWebSocketDelegate is a delegate implementation builder for the PURLSessionWebSocketDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type URLSessionWebSocketDelegate struct {
	_URLSessionWebSocketTaskDidCloseWithCodeReason func(session IURLSession, webSocketTask IURLSessionWebSocketTask, closeCode URLSessionWebSocketCloseCode, reason IData)
	_URLSessionWebSocketTaskDidOpenWithProtocol func(session IURLSession, webSocketTask IURLSessionWebSocketTask, protocol_ IString)
}

// SetURLSessionWebSocketTaskDidCloseWithCodeReason sets the handler for the URLSessionWebSocketTaskDidCloseWithCodeReason delegate method.
//
// Tells the delegate that the WebSocket task received a close frame from the server endpoint, optionally including a close code and reason from the server.
func (d *URLSessionWebSocketDelegate) SetURLSessionWebSocketTaskDidCloseWithCodeReason(f func(session IURLSession, webSocketTask IURLSessionWebSocketTask, closeCode URLSessionWebSocketCloseCode, reason IData)) {
	d._URLSessionWebSocketTaskDidCloseWithCodeReason = f
}

// SetURLSessionWebSocketTaskDidOpenWithProtocol sets the handler for the URLSessionWebSocketTaskDidOpenWithProtocol delegate method.
//
// Tells the delegate that the WebSocket task successfully negotiated the handshake with the endpoint, indicating the negotiated protocol.
func (d *URLSessionWebSocketDelegate) SetURLSessionWebSocketTaskDidOpenWithProtocol(f func(session IURLSession, webSocketTask IURLSessionWebSocketTask, protocol_ IString)) {
	d._URLSessionWebSocketTaskDidOpenWithProtocol = f
}

// URLSessionWebSocketTaskDidCloseWithCodeReason implements the PURLSessionWebSocketDelegate interface.
func (d *URLSessionWebSocketDelegate) URLSessionWebSocketTaskDidCloseWithCodeReason(session IURLSession, webSocketTask IURLSessionWebSocketTask, closeCode URLSessionWebSocketCloseCode, reason IData) {
	if d._URLSessionWebSocketTaskDidCloseWithCodeReason != nil {
		d._URLSessionWebSocketTaskDidCloseWithCodeReason(session, webSocketTask, closeCode, reason)
	}
}

// HasURLSessionWebSocketTaskDidCloseWithCodeReason returns true if a handler for URLSessionWebSocketTaskDidCloseWithCodeReason has been set.
func (d *URLSessionWebSocketDelegate) HasURLSessionWebSocketTaskDidCloseWithCodeReason() bool {
	return d._URLSessionWebSocketTaskDidCloseWithCodeReason != nil
}

// URLSessionWebSocketTaskDidOpenWithProtocol implements the PURLSessionWebSocketDelegate interface.
func (d *URLSessionWebSocketDelegate) URLSessionWebSocketTaskDidOpenWithProtocol(session IURLSession, webSocketTask IURLSessionWebSocketTask, protocol_ IString) {
	if d._URLSessionWebSocketTaskDidOpenWithProtocol != nil {
		d._URLSessionWebSocketTaskDidOpenWithProtocol(session, webSocketTask, protocol_)
	}
}

// HasURLSessionWebSocketTaskDidOpenWithProtocol returns true if a handler for URLSessionWebSocketTaskDidOpenWithProtocol has been set.
func (d *URLSessionWebSocketDelegate) HasURLSessionWebSocketTaskDidOpenWithProtocol() bool {
	return d._URLSessionWebSocketTaskDidOpenWithProtocol != nil
}
