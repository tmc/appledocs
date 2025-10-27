// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PURLSessionStreamDelegate is the NSURLSessionStreamDelegate protocol interface.
//
// A protocol that defines methods that URL session instances call on their delegates to handle task-level events specific to stream tasks.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/URLSessionStreamDelegate
type PURLSessionStreamDelegate interface {
	// Optional methods
	URLSessionBetterRouteDiscoveredForStreamTask(session IURLSession, streamTask IURLSessionStreamTask)
	HasURLSessionBetterRouteDiscoveredForStreamTask() bool
	URLSessionReadClosedForStreamTask(session IURLSession, streamTask IURLSessionStreamTask)
	HasURLSessionReadClosedForStreamTask() bool
	URLSessionStreamTaskDidBecomeInputStreamOutputStream(session IURLSession, streamTask IURLSessionStreamTask, inputStream IInputStream, outputStream IOutputStream)
	HasURLSessionStreamTaskDidBecomeInputStreamOutputStream() bool
	URLSessionWriteClosedForStreamTask(session IURLSession, streamTask IURLSessionStreamTask)
	HasURLSessionWriteClosedForStreamTask() bool
}

// URLSessionStreamDelegate is a delegate implementation builder for the PURLSessionStreamDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type URLSessionStreamDelegate struct {
	_URLSessionBetterRouteDiscoveredForStreamTask func(session IURLSession, streamTask IURLSessionStreamTask)
	_URLSessionReadClosedForStreamTask func(session IURLSession, streamTask IURLSessionStreamTask)
	_URLSessionStreamTaskDidBecomeInputStreamOutputStream func(session IURLSession, streamTask IURLSessionStreamTask, inputStream IInputStream, outputStream IOutputStream)
	_URLSessionWriteClosedForStreamTask func(session IURLSession, streamTask IURLSessionStreamTask)
}

// SetURLSessionBetterRouteDiscoveredForStreamTask sets the handler for the URLSessionBetterRouteDiscoveredForStreamTask delegate method.
//
// Tells the delegate that a better route to the host has been detected for the stream.
func (d *URLSessionStreamDelegate) SetURLSessionBetterRouteDiscoveredForStreamTask(f func(session IURLSession, streamTask IURLSessionStreamTask)) {
	d._URLSessionBetterRouteDiscoveredForStreamTask = f
}

// SetURLSessionReadClosedForStreamTask sets the handler for the URLSessionReadClosedForStreamTask delegate method.
//
// Tells the delegate that the read side of the underlying socket has been closed.
func (d *URLSessionStreamDelegate) SetURLSessionReadClosedForStreamTask(f func(session IURLSession, streamTask IURLSessionStreamTask)) {
	d._URLSessionReadClosedForStreamTask = f
}

// SetURLSessionStreamTaskDidBecomeInputStreamOutputStream sets the handler for the URLSessionStreamTaskDidBecomeInputStreamOutputStream delegate method.
//
// Tells the delegate that the stream task has been completed as a result of the stream task calling the   method.
func (d *URLSessionStreamDelegate) SetURLSessionStreamTaskDidBecomeInputStreamOutputStream(f func(session IURLSession, streamTask IURLSessionStreamTask, inputStream IInputStream, outputStream IOutputStream)) {
	d._URLSessionStreamTaskDidBecomeInputStreamOutputStream = f
}

// SetURLSessionWriteClosedForStreamTask sets the handler for the URLSessionWriteClosedForStreamTask delegate method.
//
// Tells the delegate that the write side of the underlying socket has been closed.
func (d *URLSessionStreamDelegate) SetURLSessionWriteClosedForStreamTask(f func(session IURLSession, streamTask IURLSessionStreamTask)) {
	d._URLSessionWriteClosedForStreamTask = f
}

// URLSessionBetterRouteDiscoveredForStreamTask implements the PURLSessionStreamDelegate interface.
func (d *URLSessionStreamDelegate) URLSessionBetterRouteDiscoveredForStreamTask(session IURLSession, streamTask IURLSessionStreamTask) {
	if d._URLSessionBetterRouteDiscoveredForStreamTask != nil {
		d._URLSessionBetterRouteDiscoveredForStreamTask(session, streamTask)
	}
}

// HasURLSessionBetterRouteDiscoveredForStreamTask returns true if a handler for URLSessionBetterRouteDiscoveredForStreamTask has been set.
func (d *URLSessionStreamDelegate) HasURLSessionBetterRouteDiscoveredForStreamTask() bool {
	return d._URLSessionBetterRouteDiscoveredForStreamTask != nil
}

// URLSessionReadClosedForStreamTask implements the PURLSessionStreamDelegate interface.
func (d *URLSessionStreamDelegate) URLSessionReadClosedForStreamTask(session IURLSession, streamTask IURLSessionStreamTask) {
	if d._URLSessionReadClosedForStreamTask != nil {
		d._URLSessionReadClosedForStreamTask(session, streamTask)
	}
}

// HasURLSessionReadClosedForStreamTask returns true if a handler for URLSessionReadClosedForStreamTask has been set.
func (d *URLSessionStreamDelegate) HasURLSessionReadClosedForStreamTask() bool {
	return d._URLSessionReadClosedForStreamTask != nil
}

// URLSessionStreamTaskDidBecomeInputStreamOutputStream implements the PURLSessionStreamDelegate interface.
func (d *URLSessionStreamDelegate) URLSessionStreamTaskDidBecomeInputStreamOutputStream(session IURLSession, streamTask IURLSessionStreamTask, inputStream IInputStream, outputStream IOutputStream) {
	if d._URLSessionStreamTaskDidBecomeInputStreamOutputStream != nil {
		d._URLSessionStreamTaskDidBecomeInputStreamOutputStream(session, streamTask, inputStream, outputStream)
	}
}

// HasURLSessionStreamTaskDidBecomeInputStreamOutputStream returns true if a handler for URLSessionStreamTaskDidBecomeInputStreamOutputStream has been set.
func (d *URLSessionStreamDelegate) HasURLSessionStreamTaskDidBecomeInputStreamOutputStream() bool {
	return d._URLSessionStreamTaskDidBecomeInputStreamOutputStream != nil
}

// URLSessionWriteClosedForStreamTask implements the PURLSessionStreamDelegate interface.
func (d *URLSessionStreamDelegate) URLSessionWriteClosedForStreamTask(session IURLSession, streamTask IURLSessionStreamTask) {
	if d._URLSessionWriteClosedForStreamTask != nil {
		d._URLSessionWriteClosedForStreamTask(session, streamTask)
	}
}

// HasURLSessionWriteClosedForStreamTask returns true if a handler for URLSessionWriteClosedForStreamTask has been set.
func (d *URLSessionStreamDelegate) HasURLSessionWriteClosedForStreamTask() bool {
	return d._URLSessionWriteClosedForStreamTask != nil
}

// URLSessionStreamDelegateObject wraps an existing Objective-C object that conforms to the PURLSessionStreamDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type URLSessionStreamDelegateObject struct {
	objectivec.Object
}

// NewURLSessionStreamDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSURLSessionStreamDelegate protocol.
func NewURLSessionStreamDelegateObject(obj objectivec.Object) *URLSessionStreamDelegateObject {
	return &URLSessionStreamDelegateObject{obj}
}

// Make sure URLSessionStreamDelegateObject implements PURLSessionStreamDelegate.
var _ PURLSessionStreamDelegate = (*URLSessionStreamDelegateObject)(nil)

// URLSessionBetterRouteDiscoveredForStreamTask implements the PURLSessionStreamDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionStreamDelegateObject) URLSessionBetterRouteDiscoveredForStreamTask(session IURLSession, streamTask IURLSessionStreamTask) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:betterRouteDiscoveredForStreamTask:"), session, streamTask)
}

// HasURLSessionBetterRouteDiscoveredForStreamTask returns true; this is a placeholder for optional method checks.
func (o *URLSessionStreamDelegateObject) HasURLSessionBetterRouteDiscoveredForStreamTask() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionReadClosedForStreamTask implements the PURLSessionStreamDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionStreamDelegateObject) URLSessionReadClosedForStreamTask(session IURLSession, streamTask IURLSessionStreamTask) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:readClosedForStreamTask:"), session, streamTask)
}

// HasURLSessionReadClosedForStreamTask returns true; this is a placeholder for optional method checks.
func (o *URLSessionStreamDelegateObject) HasURLSessionReadClosedForStreamTask() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionStreamTaskDidBecomeInputStreamOutputStream implements the PURLSessionStreamDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionStreamDelegateObject) URLSessionStreamTaskDidBecomeInputStreamOutputStream(session IURLSession, streamTask IURLSessionStreamTask, inputStream IInputStream, outputStream IOutputStream) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:streamTask:didBecomeInputStream:outputStream:"), session, streamTask, inputStream, outputStream)
}

// HasURLSessionStreamTaskDidBecomeInputStreamOutputStream returns true; this is a placeholder for optional method checks.
func (o *URLSessionStreamDelegateObject) HasURLSessionStreamTaskDidBecomeInputStreamOutputStream() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionWriteClosedForStreamTask implements the PURLSessionStreamDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionStreamDelegateObject) URLSessionWriteClosedForStreamTask(session IURLSession, streamTask IURLSessionStreamTask) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:writeClosedForStreamTask:"), session, streamTask)
}

// HasURLSessionWriteClosedForStreamTask returns true; this is a placeholder for optional method checks.
func (o *URLSessionStreamDelegateObject) HasURLSessionWriteClosedForStreamTask() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
