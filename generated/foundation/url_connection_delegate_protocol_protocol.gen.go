// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PURLConnectionDelegate is the NSURLConnectionDelegate protocol interface.
//
// A protocol that delegates of a URL connection implement to receive status about and provide feedback to the connection object.
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
// See: doc://com.apple.foundation/documentation/Foundation/NSURLConnectionDelegate
type PURLConnectionDelegate interface {
	// Optional methods
	ConnectionCanAuthenticateAgainstProtectionSpace(connection IURLConnection, protectionSpace IURLProtectionSpace) bool
	HasConnectionCanAuthenticateAgainstProtectionSpace() bool
	ConnectionDidCancelAuthenticationChallenge(connection IURLConnection, challenge IURLAuthenticationChallenge)
	HasConnectionDidCancelAuthenticationChallenge() bool
	ConnectionDidFailWithError(connection IURLConnection, error_ IError)
	HasConnectionDidFailWithError() bool
	ConnectionDidReceiveAuthenticationChallenge(connection IURLConnection, challenge IURLAuthenticationChallenge)
	HasConnectionDidReceiveAuthenticationChallenge() bool
	ConnectionWillSendRequestForAuthenticationChallenge(connection IURLConnection, challenge IURLAuthenticationChallenge)
	HasConnectionWillSendRequestForAuthenticationChallenge() bool
	ConnectionShouldUseCredentialStorage(connection IURLConnection) bool
	HasConnectionShouldUseCredentialStorage() bool
}

// URLConnectionDelegate is a delegate implementation builder for the PURLConnectionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type URLConnectionDelegate struct {
	_ConnectionCanAuthenticateAgainstProtectionSpace func(connection IURLConnection, protectionSpace IURLProtectionSpace) bool
	_ConnectionDidCancelAuthenticationChallenge func(connection IURLConnection, challenge IURLAuthenticationChallenge)
	_ConnectionDidFailWithError func(connection IURLConnection, error_ IError)
	_ConnectionDidReceiveAuthenticationChallenge func(connection IURLConnection, challenge IURLAuthenticationChallenge)
	_ConnectionWillSendRequestForAuthenticationChallenge func(connection IURLConnection, challenge IURLAuthenticationChallenge)
	_ConnectionShouldUseCredentialStorage func(connection IURLConnection) bool
}

// SetConnectionCanAuthenticateAgainstProtectionSpace sets the handler for the ConnectionCanAuthenticateAgainstProtectionSpace delegate method.
//
// Sent to determine whether the delegate is able to respond to a protection space’s form of authentication.
func (d *URLConnectionDelegate) SetConnectionCanAuthenticateAgainstProtectionSpace(f func(connection IURLConnection, protectionSpace IURLProtectionSpace) bool) {
	d._ConnectionCanAuthenticateAgainstProtectionSpace = f
}

// SetConnectionDidCancelAuthenticationChallenge sets the handler for the ConnectionDidCancelAuthenticationChallenge delegate method.
//
// Sent when a connection cancels an authentication challenge.
func (d *URLConnectionDelegate) SetConnectionDidCancelAuthenticationChallenge(f func(connection IURLConnection, challenge IURLAuthenticationChallenge)) {
	d._ConnectionDidCancelAuthenticationChallenge = f
}

// SetConnectionDidFailWithError sets the handler for the ConnectionDidFailWithError delegate method.
//
// Sent when a connection fails to load its request successfully.
func (d *URLConnectionDelegate) SetConnectionDidFailWithError(f func(connection IURLConnection, error_ IError)) {
	d._ConnectionDidFailWithError = f
}

// SetConnectionDidReceiveAuthenticationChallenge sets the handler for the ConnectionDidReceiveAuthenticationChallenge delegate method.
//
// Sent when a connection must authenticate a challenge in order to download its request.
func (d *URLConnectionDelegate) SetConnectionDidReceiveAuthenticationChallenge(f func(connection IURLConnection, challenge IURLAuthenticationChallenge)) {
	d._ConnectionDidReceiveAuthenticationChallenge = f
}

// SetConnectionWillSendRequestForAuthenticationChallenge sets the handler for the ConnectionWillSendRequestForAuthenticationChallenge delegate method.
//
// Tells the delegate that the connection will send a request for an authentication challenge.
func (d *URLConnectionDelegate) SetConnectionWillSendRequestForAuthenticationChallenge(f func(connection IURLConnection, challenge IURLAuthenticationChallenge)) {
	d._ConnectionWillSendRequestForAuthenticationChallenge = f
}

// SetConnectionShouldUseCredentialStorage sets the handler for the ConnectionShouldUseCredentialStorage delegate method.
//
// Sent to determine whether the URL loader should use the credential storage for authenticating the connection.
func (d *URLConnectionDelegate) SetConnectionShouldUseCredentialStorage(f func(connection IURLConnection) bool) {
	d._ConnectionShouldUseCredentialStorage = f
}

// ConnectionCanAuthenticateAgainstProtectionSpace implements the PURLConnectionDelegate interface.
func (d *URLConnectionDelegate) ConnectionCanAuthenticateAgainstProtectionSpace(connection IURLConnection, protectionSpace IURLProtectionSpace) bool {
	if d._ConnectionCanAuthenticateAgainstProtectionSpace != nil {
		return d._ConnectionCanAuthenticateAgainstProtectionSpace(connection, protectionSpace)
	}
	var zero bool
	return zero
}

// HasConnectionCanAuthenticateAgainstProtectionSpace returns true if a handler for ConnectionCanAuthenticateAgainstProtectionSpace has been set.
func (d *URLConnectionDelegate) HasConnectionCanAuthenticateAgainstProtectionSpace() bool {
	return d._ConnectionCanAuthenticateAgainstProtectionSpace != nil
}

// ConnectionDidCancelAuthenticationChallenge implements the PURLConnectionDelegate interface.
func (d *URLConnectionDelegate) ConnectionDidCancelAuthenticationChallenge(connection IURLConnection, challenge IURLAuthenticationChallenge) {
	if d._ConnectionDidCancelAuthenticationChallenge != nil {
		d._ConnectionDidCancelAuthenticationChallenge(connection, challenge)
	}
}

// HasConnectionDidCancelAuthenticationChallenge returns true if a handler for ConnectionDidCancelAuthenticationChallenge has been set.
func (d *URLConnectionDelegate) HasConnectionDidCancelAuthenticationChallenge() bool {
	return d._ConnectionDidCancelAuthenticationChallenge != nil
}

// ConnectionDidFailWithError implements the PURLConnectionDelegate interface.
func (d *URLConnectionDelegate) ConnectionDidFailWithError(connection IURLConnection, error_ IError) {
	if d._ConnectionDidFailWithError != nil {
		d._ConnectionDidFailWithError(connection, error_)
	}
}

// HasConnectionDidFailWithError returns true if a handler for ConnectionDidFailWithError has been set.
func (d *URLConnectionDelegate) HasConnectionDidFailWithError() bool {
	return d._ConnectionDidFailWithError != nil
}

// ConnectionDidReceiveAuthenticationChallenge implements the PURLConnectionDelegate interface.
func (d *URLConnectionDelegate) ConnectionDidReceiveAuthenticationChallenge(connection IURLConnection, challenge IURLAuthenticationChallenge) {
	if d._ConnectionDidReceiveAuthenticationChallenge != nil {
		d._ConnectionDidReceiveAuthenticationChallenge(connection, challenge)
	}
}

// HasConnectionDidReceiveAuthenticationChallenge returns true if a handler for ConnectionDidReceiveAuthenticationChallenge has been set.
func (d *URLConnectionDelegate) HasConnectionDidReceiveAuthenticationChallenge() bool {
	return d._ConnectionDidReceiveAuthenticationChallenge != nil
}

// ConnectionWillSendRequestForAuthenticationChallenge implements the PURLConnectionDelegate interface.
func (d *URLConnectionDelegate) ConnectionWillSendRequestForAuthenticationChallenge(connection IURLConnection, challenge IURLAuthenticationChallenge) {
	if d._ConnectionWillSendRequestForAuthenticationChallenge != nil {
		d._ConnectionWillSendRequestForAuthenticationChallenge(connection, challenge)
	}
}

// HasConnectionWillSendRequestForAuthenticationChallenge returns true if a handler for ConnectionWillSendRequestForAuthenticationChallenge has been set.
func (d *URLConnectionDelegate) HasConnectionWillSendRequestForAuthenticationChallenge() bool {
	return d._ConnectionWillSendRequestForAuthenticationChallenge != nil
}

// ConnectionShouldUseCredentialStorage implements the PURLConnectionDelegate interface.
func (d *URLConnectionDelegate) ConnectionShouldUseCredentialStorage(connection IURLConnection) bool {
	if d._ConnectionShouldUseCredentialStorage != nil {
		return d._ConnectionShouldUseCredentialStorage(connection)
	}
	var zero bool
	return zero
}

// HasConnectionShouldUseCredentialStorage returns true if a handler for ConnectionShouldUseCredentialStorage has been set.
func (d *URLConnectionDelegate) HasConnectionShouldUseCredentialStorage() bool {
	return d._ConnectionShouldUseCredentialStorage != nil
}

// URLConnectionDelegateObject wraps an existing Objective-C object that conforms to the PURLConnectionDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type URLConnectionDelegateObject struct {
	objectivec.Object
}

// NewURLConnectionDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSURLConnectionDelegate protocol.
func NewURLConnectionDelegateObject(obj objectivec.Object) *URLConnectionDelegateObject {
	return &URLConnectionDelegateObject{obj}
}

// Make sure URLConnectionDelegateObject implements PURLConnectionDelegate.
var _ PURLConnectionDelegate = (*URLConnectionDelegateObject)(nil)

// ConnectionCanAuthenticateAgainstProtectionSpace implements the PURLConnectionDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLConnectionDelegateObject) ConnectionCanAuthenticateAgainstProtectionSpace(connection IURLConnection, protectionSpace IURLProtectionSpace) bool {
	return objc.Send[bool](o.ID, objc.Sel("connection:canAuthenticateAgainstProtectionSpace:"), connection, protectionSpace)
}

// HasConnectionCanAuthenticateAgainstProtectionSpace returns true; this is a placeholder for optional method checks.
func (o *URLConnectionDelegateObject) HasConnectionCanAuthenticateAgainstProtectionSpace() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ConnectionDidCancelAuthenticationChallenge implements the PURLConnectionDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLConnectionDelegateObject) ConnectionDidCancelAuthenticationChallenge(connection IURLConnection, challenge IURLAuthenticationChallenge) {
	objc.Send[objc.ID](o.ID, objc.Sel("connection:didCancelAuthenticationChallenge:"), connection, challenge)
}

// HasConnectionDidCancelAuthenticationChallenge returns true; this is a placeholder for optional method checks.
func (o *URLConnectionDelegateObject) HasConnectionDidCancelAuthenticationChallenge() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ConnectionDidFailWithError implements the PURLConnectionDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLConnectionDelegateObject) ConnectionDidFailWithError(connection IURLConnection, error_ IError) {
	objc.Send[objc.ID](o.ID, objc.Sel("connection:didFailWithError:"), connection, error_)
}

// HasConnectionDidFailWithError returns true; this is a placeholder for optional method checks.
func (o *URLConnectionDelegateObject) HasConnectionDidFailWithError() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ConnectionDidReceiveAuthenticationChallenge implements the PURLConnectionDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLConnectionDelegateObject) ConnectionDidReceiveAuthenticationChallenge(connection IURLConnection, challenge IURLAuthenticationChallenge) {
	objc.Send[objc.ID](o.ID, objc.Sel("connection:didReceiveAuthenticationChallenge:"), connection, challenge)
}

// HasConnectionDidReceiveAuthenticationChallenge returns true; this is a placeholder for optional method checks.
func (o *URLConnectionDelegateObject) HasConnectionDidReceiveAuthenticationChallenge() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ConnectionWillSendRequestForAuthenticationChallenge implements the PURLConnectionDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLConnectionDelegateObject) ConnectionWillSendRequestForAuthenticationChallenge(connection IURLConnection, challenge IURLAuthenticationChallenge) {
	objc.Send[objc.ID](o.ID, objc.Sel("connection:willSendRequestForAuthenticationChallenge:"), connection, challenge)
}

// HasConnectionWillSendRequestForAuthenticationChallenge returns true; this is a placeholder for optional method checks.
func (o *URLConnectionDelegateObject) HasConnectionWillSendRequestForAuthenticationChallenge() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ConnectionShouldUseCredentialStorage implements the PURLConnectionDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLConnectionDelegateObject) ConnectionShouldUseCredentialStorage(connection IURLConnection) bool {
	return objc.Send[bool](o.ID, objc.Sel("connectionShouldUseCredentialStorage:"), connection)
}

// HasConnectionShouldUseCredentialStorage returns true; this is a placeholder for optional method checks.
func (o *URLConnectionDelegateObject) HasConnectionShouldUseCredentialStorage() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
