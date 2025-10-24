// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PURLSessionDelegate is the NSURLSessionDelegate protocol interface.
//
// A protocol that defines methods that URL session instances call on their delegates to handle session-level events, like session life cycle changes.
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
// See: doc://com.apple.foundation/documentation/Foundation/URLSessionDelegate
type PURLSessionDelegate interface {
	// Optional methods
	URLSessionDidBecomeInvalidWithError(session IURLSession, error_ IError)
	HasURLSessionDidBecomeInvalidWithError() bool
	URLSessionDidReceiveChallengeCompletionHandler(session IURLSession, challenge IURLAuthenticationChallenge, completionHandler unsafe.Pointer)
	HasURLSessionDidReceiveChallengeCompletionHandler() bool
	URLSessionDidFinishEventsForBackgroundURLSession(session IURLSession)
	HasURLSessionDidFinishEventsForBackgroundURLSession() bool
}

// URLSessionDelegate is a delegate implementation builder for the PURLSessionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type URLSessionDelegate struct {
	_URLSessionDidBecomeInvalidWithError func(session IURLSession, error_ IError)
	_URLSessionDidReceiveChallengeCompletionHandler func(session IURLSession, challenge IURLAuthenticationChallenge, completionHandler unsafe.Pointer)
	_URLSessionDidFinishEventsForBackgroundURLSession func(session IURLSession)
}

// SetURLSessionDidBecomeInvalidWithError sets the handler for the URLSessionDidBecomeInvalidWithError delegate method.
//
// Tells the URL session that the session has been invalidated.
func (d *URLSessionDelegate) SetURLSessionDidBecomeInvalidWithError(f func(session IURLSession, error_ IError)) {
	d._URLSessionDidBecomeInvalidWithError = f
}

// SetURLSessionDidReceiveChallengeCompletionHandler sets the handler for the URLSessionDidReceiveChallengeCompletionHandler delegate method.
//
// Requests credentials from the delegate in response to a session-level authentication request from the remote server.
func (d *URLSessionDelegate) SetURLSessionDidReceiveChallengeCompletionHandler(f func(session IURLSession, challenge IURLAuthenticationChallenge, completionHandler unsafe.Pointer)) {
	d._URLSessionDidReceiveChallengeCompletionHandler = f
}

// SetURLSessionDidFinishEventsForBackgroundURLSession sets the handler for the URLSessionDidFinishEventsForBackgroundURLSession delegate method.
//
// Tells the delegate that all messages enqueued for a session have been delivered.
func (d *URLSessionDelegate) SetURLSessionDidFinishEventsForBackgroundURLSession(f func(session IURLSession)) {
	d._URLSessionDidFinishEventsForBackgroundURLSession = f
}

// URLSessionDidBecomeInvalidWithError implements the PURLSessionDelegate interface.
func (d *URLSessionDelegate) URLSessionDidBecomeInvalidWithError(session IURLSession, error_ IError) {
	if d._URLSessionDidBecomeInvalidWithError != nil {
		d._URLSessionDidBecomeInvalidWithError(session, error_)
	}
}

// HasURLSessionDidBecomeInvalidWithError returns true if a handler for URLSessionDidBecomeInvalidWithError has been set.
func (d *URLSessionDelegate) HasURLSessionDidBecomeInvalidWithError() bool {
	return d._URLSessionDidBecomeInvalidWithError != nil
}

// URLSessionDidReceiveChallengeCompletionHandler implements the PURLSessionDelegate interface.
func (d *URLSessionDelegate) URLSessionDidReceiveChallengeCompletionHandler(session IURLSession, challenge IURLAuthenticationChallenge, completionHandler unsafe.Pointer) {
	if d._URLSessionDidReceiveChallengeCompletionHandler != nil {
		d._URLSessionDidReceiveChallengeCompletionHandler(session, challenge, completionHandler)
	}
}

// HasURLSessionDidReceiveChallengeCompletionHandler returns true if a handler for URLSessionDidReceiveChallengeCompletionHandler has been set.
func (d *URLSessionDelegate) HasURLSessionDidReceiveChallengeCompletionHandler() bool {
	return d._URLSessionDidReceiveChallengeCompletionHandler != nil
}

// URLSessionDidFinishEventsForBackgroundURLSession implements the PURLSessionDelegate interface.
func (d *URLSessionDelegate) URLSessionDidFinishEventsForBackgroundURLSession(session IURLSession) {
	if d._URLSessionDidFinishEventsForBackgroundURLSession != nil {
		d._URLSessionDidFinishEventsForBackgroundURLSession(session)
	}
}

// HasURLSessionDidFinishEventsForBackgroundURLSession returns true if a handler for URLSessionDidFinishEventsForBackgroundURLSession has been set.
func (d *URLSessionDelegate) HasURLSessionDidFinishEventsForBackgroundURLSession() bool {
	return d._URLSessionDidFinishEventsForBackgroundURLSession != nil
}
