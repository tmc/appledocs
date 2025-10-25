// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PConnectionDelegate is the NSConnectionDelegate protocol interface.
//
// An interface for interacting with low-level, interprocess connections.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - macOS 10.0+ (Deprecated in 10.13)
//
// See: doc://com.apple.foundation/documentation/Foundation/NSConnectionDelegate
type PConnectionDelegate interface {
	// Required methods
	AuthenticateComponentsWithData(components IArray, signature IData) bool/* debug [protocol_interface/required_method]: AuthenticateComponentsWithData */
	AuthenticationDataForComponents(components IArray) Data/* debug [protocol_interface/required_method]: AuthenticationDataForComponents */
	ConnectionHandleRequest(connection IConnection, doreq IDistantObjectRequest) bool/* debug [protocol_interface/required_method]: ConnectionHandleRequest */
	ConnectionShouldMakeNewConnection(ancestor IConnection, conn IConnection) bool/* debug [protocol_interface/required_method]: ConnectionShouldMakeNewConnection */
	CreateConversationForConnection(conn IConnection) objc.ID/* debug [protocol_interface/required_method]: CreateConversationForConnection */
	MakeNewConnectionSender(conn IConnection, ancestor IConnection) bool/* debug [protocol_interface/required_method]: MakeNewConnectionSender */
}

// ConnectionDelegate is a delegate implementation builder for the PConnectionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ConnectionDelegate struct {
	_AuthenticateComponentsWithData func(components IArray, signature IData) bool
	_AuthenticationDataForComponents func(components IArray) Data
	_ConnectionHandleRequest func(connection IConnection, doreq IDistantObjectRequest) bool
	_ConnectionShouldMakeNewConnection func(ancestor IConnection, conn IConnection) bool
	_CreateConversationForConnection func(conn IConnection) objc.ID
	_MakeNewConnectionSender func(conn IConnection, ancestor IConnection) bool
}

// SetAuthenticateComponentsWithData sets the handler for the AuthenticateComponentsWithData delegate method.
//
// Returns a Boolean value that indicates whether given authentication data is valid for a given set of components.
func (d *ConnectionDelegate) SetAuthenticateComponentsWithData(f func(components IArray, signature IData) bool) {
	d._AuthenticateComponentsWithData = f
}

// SetAuthenticationDataForComponents sets the handler for the AuthenticationDataForComponents delegate method.
//
// Returns an   object to be used as an authentication stamp for an outgoing message.
func (d *ConnectionDelegate) SetAuthenticationDataForComponents(f func(components IArray) Data) {
	d._AuthenticationDataForComponents = f
}

// SetConnectionHandleRequest sets the handler for the ConnectionHandleRequest delegate method.
//
// This method should be implemented by   object delegates that want to intercept distant object requests.
func (d *ConnectionDelegate) SetConnectionHandleRequest(f func(connection IConnection, doreq IDistantObjectRequest) bool) {
	d._ConnectionHandleRequest = f
}

// SetConnectionShouldMakeNewConnection sets the handler for the ConnectionShouldMakeNewConnection delegate method.
//
// Returns a Boolean value that indicates whether the parent connection should allow a given new connection to be created.
func (d *ConnectionDelegate) SetConnectionShouldMakeNewConnection(f func(ancestor IConnection, conn IConnection) bool) {
	d._ConnectionShouldMakeNewConnection = f
}

// SetCreateConversationForConnection sets the handler for the CreateConversationForConnection delegate method.
//
// Returns an arbitrary object identifying a new conversation being created for the connection in the current thread.
func (d *ConnectionDelegate) SetCreateConversationForConnection(f func(conn IConnection) objc.ID) {
	d._CreateConversationForConnection = f
}

// SetMakeNewConnectionSender sets the handler for the MakeNewConnectionSender delegate method.
//
// Returns a Boolean value that indicates whether the parent should allow a given new connection to be created and configured.
func (d *ConnectionDelegate) SetMakeNewConnectionSender(f func(conn IConnection, ancestor IConnection) bool) {
	d._MakeNewConnectionSender = f
}

// AuthenticateComponentsWithData implements the PConnectionDelegate interface.
func (d *ConnectionDelegate) AuthenticateComponentsWithData(components IArray, signature IData) bool {
	if d._AuthenticateComponentsWithData != nil {
		return d._AuthenticateComponentsWithData(components, signature)
	}
	var zero bool
	return zero
}

// HasAuthenticateComponentsWithData returns true if a handler for AuthenticateComponentsWithData has been set.
func (d *ConnectionDelegate) HasAuthenticateComponentsWithData() bool {
	return d._AuthenticateComponentsWithData != nil
}

// AuthenticationDataForComponents implements the PConnectionDelegate interface.
func (d *ConnectionDelegate) AuthenticationDataForComponents(components IArray) Data {
	if d._AuthenticationDataForComponents != nil {
		return d._AuthenticationDataForComponents(components)
	}
	var zero Data
	return zero
}

// HasAuthenticationDataForComponents returns true if a handler for AuthenticationDataForComponents has been set.
func (d *ConnectionDelegate) HasAuthenticationDataForComponents() bool {
	return d._AuthenticationDataForComponents != nil
}

// ConnectionHandleRequest implements the PConnectionDelegate interface.
func (d *ConnectionDelegate) ConnectionHandleRequest(connection IConnection, doreq IDistantObjectRequest) bool {
	if d._ConnectionHandleRequest != nil {
		return d._ConnectionHandleRequest(connection, doreq)
	}
	var zero bool
	return zero
}

// HasConnectionHandleRequest returns true if a handler for ConnectionHandleRequest has been set.
func (d *ConnectionDelegate) HasConnectionHandleRequest() bool {
	return d._ConnectionHandleRequest != nil
}

// ConnectionShouldMakeNewConnection implements the PConnectionDelegate interface.
func (d *ConnectionDelegate) ConnectionShouldMakeNewConnection(ancestor IConnection, conn IConnection) bool {
	if d._ConnectionShouldMakeNewConnection != nil {
		return d._ConnectionShouldMakeNewConnection(ancestor, conn)
	}
	var zero bool
	return zero
}

// HasConnectionShouldMakeNewConnection returns true if a handler for ConnectionShouldMakeNewConnection has been set.
func (d *ConnectionDelegate) HasConnectionShouldMakeNewConnection() bool {
	return d._ConnectionShouldMakeNewConnection != nil
}

// CreateConversationForConnection implements the PConnectionDelegate interface.
func (d *ConnectionDelegate) CreateConversationForConnection(conn IConnection) objc.ID {
	if d._CreateConversationForConnection != nil {
		return d._CreateConversationForConnection(conn)
	}
	var zero objc.ID
	return zero
}

// HasCreateConversationForConnection returns true if a handler for CreateConversationForConnection has been set.
func (d *ConnectionDelegate) HasCreateConversationForConnection() bool {
	return d._CreateConversationForConnection != nil
}

// MakeNewConnectionSender implements the PConnectionDelegate interface.
func (d *ConnectionDelegate) MakeNewConnectionSender(conn IConnection, ancestor IConnection) bool {
	if d._MakeNewConnectionSender != nil {
		return d._MakeNewConnectionSender(conn, ancestor)
	}
	var zero bool
	return zero
}

// HasMakeNewConnectionSender returns true if a handler for MakeNewConnectionSender has been set.
func (d *ConnectionDelegate) HasMakeNewConnectionSender() bool {
	return d._MakeNewConnectionSender != nil
}
