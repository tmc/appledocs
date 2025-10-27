// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
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

// URLSessionWebSocketDelegateObject wraps an existing Objective-C object that conforms to the PURLSessionWebSocketDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type URLSessionWebSocketDelegateObject struct {
	objectivec.Object
}

// NewURLSessionWebSocketDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSURLSessionWebSocketDelegate protocol.
func NewURLSessionWebSocketDelegateObject(obj objectivec.Object) *URLSessionWebSocketDelegateObject {
	return &URLSessionWebSocketDelegateObject{obj}
}

// Make sure URLSessionWebSocketDelegateObject implements PURLSessionWebSocketDelegate.
var _ PURLSessionWebSocketDelegate = (*URLSessionWebSocketDelegateObject)(nil)

// URLSessionWebSocketTaskDidCloseWithCodeReason implements the PURLSessionWebSocketDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionWebSocketDelegateObject) URLSessionWebSocketTaskDidCloseWithCodeReason(session IURLSession, webSocketTask IURLSessionWebSocketTask, closeCode URLSessionWebSocketCloseCode, reason IData) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:webSocketTask:didCloseWithCode:reason:"), session, webSocketTask, closeCode, reason)
}

// HasURLSessionWebSocketTaskDidCloseWithCodeReason returns true; this is a placeholder for optional method checks.
func (o *URLSessionWebSocketDelegateObject) HasURLSessionWebSocketTaskDidCloseWithCodeReason() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionWebSocketTaskDidOpenWithProtocol implements the PURLSessionWebSocketDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionWebSocketDelegateObject) URLSessionWebSocketTaskDidOpenWithProtocol(session IURLSession, webSocketTask IURLSessionWebSocketTask, protocol_ IString) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:webSocketTask:didOpenWithProtocol:"), session, webSocketTask, protocol_)
}

// HasURLSessionWebSocketTaskDidOpenWithProtocol returns true; this is a placeholder for optional method checks.
func (o *URLSessionWebSocketDelegateObject) HasURLSessionWebSocketTaskDidOpenWithProtocol() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
