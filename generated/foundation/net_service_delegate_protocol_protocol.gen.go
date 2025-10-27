// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PNetServiceDelegate is the NSNetServiceDelegate protocol interface.
//
// The interface a net service uses to inform its delegate about the state of the service it offers.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NetServiceDelegate
type PNetServiceDelegate interface {
	// Optional methods
	NetServiceDidAcceptConnectionWithInputStreamOutputStream(sender INetService, inputStream IInputStream, outputStream IOutputStream)
	HasNetServiceDidAcceptConnectionWithInputStreamOutputStream() bool
	NetServiceDidNotPublish(sender INetService, errorDict IDictionary)
	HasNetServiceDidNotPublish() bool
	NetServiceDidNotResolve(sender INetService, errorDict IDictionary)
	HasNetServiceDidNotResolve() bool
	NetServiceDidUpdateTXTRecordData(sender INetService, data IData)
	HasNetServiceDidUpdateTXTRecordData() bool
	NetServiceDidPublish(sender INetService)
	HasNetServiceDidPublish() bool
	NetServiceDidResolveAddress(sender INetService)
	HasNetServiceDidResolveAddress() bool
	NetServiceDidStop(sender INetService)
	HasNetServiceDidStop() bool
	NetServiceWillPublish(sender INetService)
	HasNetServiceWillPublish() bool
	NetServiceWillResolve(sender INetService)
	HasNetServiceWillResolve() bool
}

// NetServiceDelegate is a delegate implementation builder for the PNetServiceDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type NetServiceDelegate struct {
	_NetServiceDidAcceptConnectionWithInputStreamOutputStream func(sender INetService, inputStream IInputStream, outputStream IOutputStream)
	_NetServiceDidNotPublish func(sender INetService, errorDict IDictionary)
	_NetServiceDidNotResolve func(sender INetService, errorDict IDictionary)
	_NetServiceDidUpdateTXTRecordData func(sender INetService, data IData)
	_NetServiceDidPublish func(sender INetService)
	_NetServiceDidResolveAddress func(sender INetService)
	_NetServiceDidStop func(sender INetService)
	_NetServiceWillPublish func(sender INetService)
	_NetServiceWillResolve func(sender INetService)
}

// SetNetServiceDidAcceptConnectionWithInputStreamOutputStream sets the handler for the NetServiceDidAcceptConnectionWithInputStreamOutputStream delegate method.
//
// Called when a client connects to a service managed by Bonjour.
func (d *NetServiceDelegate) SetNetServiceDidAcceptConnectionWithInputStreamOutputStream(f func(sender INetService, inputStream IInputStream, outputStream IOutputStream)) {
	d._NetServiceDidAcceptConnectionWithInputStreamOutputStream = f
}

// SetNetServiceDidNotPublish sets the handler for the NetServiceDidNotPublish delegate method.
//
// Notifies the delegate that a service could not be published.
func (d *NetServiceDelegate) SetNetServiceDidNotPublish(f func(sender INetService, errorDict IDictionary)) {
	d._NetServiceDidNotPublish = f
}

// SetNetServiceDidNotResolve sets the handler for the NetServiceDidNotResolve delegate method.
//
// Informs the delegate that an error occurred during resolution of a given service.
func (d *NetServiceDelegate) SetNetServiceDidNotResolve(f func(sender INetService, errorDict IDictionary)) {
	d._NetServiceDidNotResolve = f
}

// SetNetServiceDidUpdateTXTRecordData sets the handler for the NetServiceDidUpdateTXTRecordData delegate method.
//
// Notifies the delegate that the TXT record for a given service has been updated.
func (d *NetServiceDelegate) SetNetServiceDidUpdateTXTRecordData(f func(sender INetService, data IData)) {
	d._NetServiceDidUpdateTXTRecordData = f
}

// SetNetServiceDidPublish sets the handler for the NetServiceDidPublish delegate method.
//
// Notifies the delegate that a service was successfully published.
func (d *NetServiceDelegate) SetNetServiceDidPublish(f func(sender INetService)) {
	d._NetServiceDidPublish = f
}

// SetNetServiceDidResolveAddress sets the handler for the NetServiceDidResolveAddress delegate method.
//
// Informs the delegate that the address for a given service was resolved.
func (d *NetServiceDelegate) SetNetServiceDidResolveAddress(f func(sender INetService)) {
	d._NetServiceDidResolveAddress = f
}

// SetNetServiceDidStop sets the handler for the NetServiceDidStop delegate method.
//
// Informs the delegate that a   or   request was stopped.
func (d *NetServiceDelegate) SetNetServiceDidStop(f func(sender INetService)) {
	d._NetServiceDidStop = f
}

// SetNetServiceWillPublish sets the handler for the NetServiceWillPublish delegate method.
//
// Notifies the delegate that the network is ready to publish the service.
func (d *NetServiceDelegate) SetNetServiceWillPublish(f func(sender INetService)) {
	d._NetServiceWillPublish = f
}

// SetNetServiceWillResolve sets the handler for the NetServiceWillResolve delegate method.
//
// Notifies the delegate that the network is ready to resolve the service.
func (d *NetServiceDelegate) SetNetServiceWillResolve(f func(sender INetService)) {
	d._NetServiceWillResolve = f
}

// NetServiceDidAcceptConnectionWithInputStreamOutputStream implements the PNetServiceDelegate interface.
func (d *NetServiceDelegate) NetServiceDidAcceptConnectionWithInputStreamOutputStream(sender INetService, inputStream IInputStream, outputStream IOutputStream) {
	if d._NetServiceDidAcceptConnectionWithInputStreamOutputStream != nil {
		d._NetServiceDidAcceptConnectionWithInputStreamOutputStream(sender, inputStream, outputStream)
	}
}

// HasNetServiceDidAcceptConnectionWithInputStreamOutputStream returns true if a handler for NetServiceDidAcceptConnectionWithInputStreamOutputStream has been set.
func (d *NetServiceDelegate) HasNetServiceDidAcceptConnectionWithInputStreamOutputStream() bool {
	return d._NetServiceDidAcceptConnectionWithInputStreamOutputStream != nil
}

// NetServiceDidNotPublish implements the PNetServiceDelegate interface.
func (d *NetServiceDelegate) NetServiceDidNotPublish(sender INetService, errorDict IDictionary) {
	if d._NetServiceDidNotPublish != nil {
		d._NetServiceDidNotPublish(sender, errorDict)
	}
}

// HasNetServiceDidNotPublish returns true if a handler for NetServiceDidNotPublish has been set.
func (d *NetServiceDelegate) HasNetServiceDidNotPublish() bool {
	return d._NetServiceDidNotPublish != nil
}

// NetServiceDidNotResolve implements the PNetServiceDelegate interface.
func (d *NetServiceDelegate) NetServiceDidNotResolve(sender INetService, errorDict IDictionary) {
	if d._NetServiceDidNotResolve != nil {
		d._NetServiceDidNotResolve(sender, errorDict)
	}
}

// HasNetServiceDidNotResolve returns true if a handler for NetServiceDidNotResolve has been set.
func (d *NetServiceDelegate) HasNetServiceDidNotResolve() bool {
	return d._NetServiceDidNotResolve != nil
}

// NetServiceDidUpdateTXTRecordData implements the PNetServiceDelegate interface.
func (d *NetServiceDelegate) NetServiceDidUpdateTXTRecordData(sender INetService, data IData) {
	if d._NetServiceDidUpdateTXTRecordData != nil {
		d._NetServiceDidUpdateTXTRecordData(sender, data)
	}
}

// HasNetServiceDidUpdateTXTRecordData returns true if a handler for NetServiceDidUpdateTXTRecordData has been set.
func (d *NetServiceDelegate) HasNetServiceDidUpdateTXTRecordData() bool {
	return d._NetServiceDidUpdateTXTRecordData != nil
}

// NetServiceDidPublish implements the PNetServiceDelegate interface.
func (d *NetServiceDelegate) NetServiceDidPublish(sender INetService) {
	if d._NetServiceDidPublish != nil {
		d._NetServiceDidPublish(sender)
	}
}

// HasNetServiceDidPublish returns true if a handler for NetServiceDidPublish has been set.
func (d *NetServiceDelegate) HasNetServiceDidPublish() bool {
	return d._NetServiceDidPublish != nil
}

// NetServiceDidResolveAddress implements the PNetServiceDelegate interface.
func (d *NetServiceDelegate) NetServiceDidResolveAddress(sender INetService) {
	if d._NetServiceDidResolveAddress != nil {
		d._NetServiceDidResolveAddress(sender)
	}
}

// HasNetServiceDidResolveAddress returns true if a handler for NetServiceDidResolveAddress has been set.
func (d *NetServiceDelegate) HasNetServiceDidResolveAddress() bool {
	return d._NetServiceDidResolveAddress != nil
}

// NetServiceDidStop implements the PNetServiceDelegate interface.
func (d *NetServiceDelegate) NetServiceDidStop(sender INetService) {
	if d._NetServiceDidStop != nil {
		d._NetServiceDidStop(sender)
	}
}

// HasNetServiceDidStop returns true if a handler for NetServiceDidStop has been set.
func (d *NetServiceDelegate) HasNetServiceDidStop() bool {
	return d._NetServiceDidStop != nil
}

// NetServiceWillPublish implements the PNetServiceDelegate interface.
func (d *NetServiceDelegate) NetServiceWillPublish(sender INetService) {
	if d._NetServiceWillPublish != nil {
		d._NetServiceWillPublish(sender)
	}
}

// HasNetServiceWillPublish returns true if a handler for NetServiceWillPublish has been set.
func (d *NetServiceDelegate) HasNetServiceWillPublish() bool {
	return d._NetServiceWillPublish != nil
}

// NetServiceWillResolve implements the PNetServiceDelegate interface.
func (d *NetServiceDelegate) NetServiceWillResolve(sender INetService) {
	if d._NetServiceWillResolve != nil {
		d._NetServiceWillResolve(sender)
	}
}

// HasNetServiceWillResolve returns true if a handler for NetServiceWillResolve has been set.
func (d *NetServiceDelegate) HasNetServiceWillResolve() bool {
	return d._NetServiceWillResolve != nil
}

// NetServiceDelegateObject wraps an existing Objective-C object that conforms to the PNetServiceDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type NetServiceDelegateObject struct {
	objectivec.Object
}

// NewNetServiceDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSNetServiceDelegate protocol.
func NewNetServiceDelegateObject(obj objectivec.Object) *NetServiceDelegateObject {
	return &NetServiceDelegateObject{obj}
}

// Make sure NetServiceDelegateObject implements PNetServiceDelegate.
var _ PNetServiceDelegate = (*NetServiceDelegateObject)(nil)

// NetServiceDidAcceptConnectionWithInputStreamOutputStream implements the PNetServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *NetServiceDelegateObject) NetServiceDidAcceptConnectionWithInputStreamOutputStream(sender INetService, inputStream IInputStream, outputStream IOutputStream) {
	objc.Send[objc.ID](o.ID, objc.Sel("netService:didAcceptConnectionWithInputStream:outputStream:"), sender, inputStream, outputStream)
}

// HasNetServiceDidAcceptConnectionWithInputStreamOutputStream returns true; this is a placeholder for optional method checks.
func (o *NetServiceDelegateObject) HasNetServiceDidAcceptConnectionWithInputStreamOutputStream() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// NetServiceDidNotPublish implements the PNetServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *NetServiceDelegateObject) NetServiceDidNotPublish(sender INetService, errorDict IDictionary) {
	objc.Send[objc.ID](o.ID, objc.Sel("netService:didNotPublish:"), sender, errorDict)
}

// HasNetServiceDidNotPublish returns true; this is a placeholder for optional method checks.
func (o *NetServiceDelegateObject) HasNetServiceDidNotPublish() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// NetServiceDidNotResolve implements the PNetServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *NetServiceDelegateObject) NetServiceDidNotResolve(sender INetService, errorDict IDictionary) {
	objc.Send[objc.ID](o.ID, objc.Sel("netService:didNotResolve:"), sender, errorDict)
}

// HasNetServiceDidNotResolve returns true; this is a placeholder for optional method checks.
func (o *NetServiceDelegateObject) HasNetServiceDidNotResolve() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// NetServiceDidUpdateTXTRecordData implements the PNetServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *NetServiceDelegateObject) NetServiceDidUpdateTXTRecordData(sender INetService, data IData) {
	objc.Send[objc.ID](o.ID, objc.Sel("netService:didUpdateTXTRecordData:"), sender, data)
}

// HasNetServiceDidUpdateTXTRecordData returns true; this is a placeholder for optional method checks.
func (o *NetServiceDelegateObject) HasNetServiceDidUpdateTXTRecordData() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// NetServiceDidPublish implements the PNetServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *NetServiceDelegateObject) NetServiceDidPublish(sender INetService) {
	objc.Send[objc.ID](o.ID, objc.Sel("netServiceDidPublish:"), sender)
}

// HasNetServiceDidPublish returns true; this is a placeholder for optional method checks.
func (o *NetServiceDelegateObject) HasNetServiceDidPublish() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// NetServiceDidResolveAddress implements the PNetServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *NetServiceDelegateObject) NetServiceDidResolveAddress(sender INetService) {
	objc.Send[objc.ID](o.ID, objc.Sel("netServiceDidResolveAddress:"), sender)
}

// HasNetServiceDidResolveAddress returns true; this is a placeholder for optional method checks.
func (o *NetServiceDelegateObject) HasNetServiceDidResolveAddress() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// NetServiceDidStop implements the PNetServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *NetServiceDelegateObject) NetServiceDidStop(sender INetService) {
	objc.Send[objc.ID](o.ID, objc.Sel("netServiceDidStop:"), sender)
}

// HasNetServiceDidStop returns true; this is a placeholder for optional method checks.
func (o *NetServiceDelegateObject) HasNetServiceDidStop() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// NetServiceWillPublish implements the PNetServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *NetServiceDelegateObject) NetServiceWillPublish(sender INetService) {
	objc.Send[objc.ID](o.ID, objc.Sel("netServiceWillPublish:"), sender)
}

// HasNetServiceWillPublish returns true; this is a placeholder for optional method checks.
func (o *NetServiceDelegateObject) HasNetServiceWillPublish() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// NetServiceWillResolve implements the PNetServiceDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *NetServiceDelegateObject) NetServiceWillResolve(sender INetService) {
	objc.Send[objc.ID](o.ID, objc.Sel("netServiceWillResolve:"), sender)
}

// HasNetServiceWillResolve returns true; this is a placeholder for optional method checks.
func (o *NetServiceDelegateObject) HasNetServiceWillResolve() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
