// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PNetServiceBrowserDelegate is the NSNetServiceBrowserDelegate protocol interface.
//
// The interface a net service browser uses to inform a delegate about the state of service discovery.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NetServiceBrowserDelegate
type PNetServiceBrowserDelegate interface {
	// Optional methods
	NetServiceBrowserDidFindServiceMoreComing(browser INetServiceBrowser, service INetService, moreComing bool)
	HasNetServiceBrowserDidFindServiceMoreComing() bool
	NetServiceBrowserDidFindDomainMoreComing(browser INetServiceBrowser, domainString IString, moreComing bool)
	HasNetServiceBrowserDidFindDomainMoreComing() bool
	NetServiceBrowserDidNotSearch(browser INetServiceBrowser, errorDict IDictionary)
	HasNetServiceBrowserDidNotSearch() bool
	NetServiceBrowserDidRemoveServiceMoreComing(browser INetServiceBrowser, service INetService, moreComing bool)
	HasNetServiceBrowserDidRemoveServiceMoreComing() bool
	NetServiceBrowserDidRemoveDomainMoreComing(browser INetServiceBrowser, domainString IString, moreComing bool)
	HasNetServiceBrowserDidRemoveDomainMoreComing() bool
	NetServiceBrowserDidStopSearch(browser INetServiceBrowser)
	HasNetServiceBrowserDidStopSearch() bool
	NetServiceBrowserWillSearch(browser INetServiceBrowser)
	HasNetServiceBrowserWillSearch() bool
}

// NetServiceBrowserDelegate is a delegate implementation builder for the PNetServiceBrowserDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type NetServiceBrowserDelegate struct {
	_NetServiceBrowserDidFindServiceMoreComing func(browser INetServiceBrowser, service INetService, moreComing bool)
	_NetServiceBrowserDidFindDomainMoreComing func(browser INetServiceBrowser, domainString IString, moreComing bool)
	_NetServiceBrowserDidNotSearch func(browser INetServiceBrowser, errorDict IDictionary)
	_NetServiceBrowserDidRemoveServiceMoreComing func(browser INetServiceBrowser, service INetService, moreComing bool)
	_NetServiceBrowserDidRemoveDomainMoreComing func(browser INetServiceBrowser, domainString IString, moreComing bool)
	_NetServiceBrowserDidStopSearch func(browser INetServiceBrowser)
	_NetServiceBrowserWillSearch func(browser INetServiceBrowser)
}

// SetNetServiceBrowserDidFindServiceMoreComing sets the handler for the NetServiceBrowserDidFindServiceMoreComing delegate method.
//
// Tells the delegate the sender found a service.
func (d *NetServiceBrowserDelegate) SetNetServiceBrowserDidFindServiceMoreComing(f func(browser INetServiceBrowser, service INetService, moreComing bool)) {
	d._NetServiceBrowserDidFindServiceMoreComing = f
}

// SetNetServiceBrowserDidFindDomainMoreComing sets the handler for the NetServiceBrowserDidFindDomainMoreComing delegate method.
//
// Tells the delegate the sender found a domain.
func (d *NetServiceBrowserDelegate) SetNetServiceBrowserDidFindDomainMoreComing(f func(browser INetServiceBrowser, domainString IString, moreComing bool)) {
	d._NetServiceBrowserDidFindDomainMoreComing = f
}

// SetNetServiceBrowserDidNotSearch sets the handler for the NetServiceBrowserDidNotSearch delegate method.
//
// Tells the delegate that a search was not successful.
func (d *NetServiceBrowserDelegate) SetNetServiceBrowserDidNotSearch(f func(browser INetServiceBrowser, errorDict IDictionary)) {
	d._NetServiceBrowserDidNotSearch = f
}

// SetNetServiceBrowserDidRemoveServiceMoreComing sets the handler for the NetServiceBrowserDidRemoveServiceMoreComing delegate method.
//
// Tells the delegate a service has disappeared or has become unavailable.
func (d *NetServiceBrowserDelegate) SetNetServiceBrowserDidRemoveServiceMoreComing(f func(browser INetServiceBrowser, service INetService, moreComing bool)) {
	d._NetServiceBrowserDidRemoveServiceMoreComing = f
}

// SetNetServiceBrowserDidRemoveDomainMoreComing sets the handler for the NetServiceBrowserDidRemoveDomainMoreComing delegate method.
//
// Tells the delegate the a domain has disappeared or has become unavailable.
func (d *NetServiceBrowserDelegate) SetNetServiceBrowserDidRemoveDomainMoreComing(f func(browser INetServiceBrowser, domainString IString, moreComing bool)) {
	d._NetServiceBrowserDidRemoveDomainMoreComing = f
}

// SetNetServiceBrowserDidStopSearch sets the handler for the NetServiceBrowserDidStopSearch delegate method.
//
// Tells the delegate that a search was stopped.
func (d *NetServiceBrowserDelegate) SetNetServiceBrowserDidStopSearch(f func(browser INetServiceBrowser)) {
	d._NetServiceBrowserDidStopSearch = f
}

// SetNetServiceBrowserWillSearch sets the handler for the NetServiceBrowserWillSearch delegate method.
//
// Tells the delegate that a search is commencing.
func (d *NetServiceBrowserDelegate) SetNetServiceBrowserWillSearch(f func(browser INetServiceBrowser)) {
	d._NetServiceBrowserWillSearch = f
}

// NetServiceBrowserDidFindServiceMoreComing implements the PNetServiceBrowserDelegate interface.
func (d *NetServiceBrowserDelegate) NetServiceBrowserDidFindServiceMoreComing(browser INetServiceBrowser, service INetService, moreComing bool) {
	if d._NetServiceBrowserDidFindServiceMoreComing != nil {
		d._NetServiceBrowserDidFindServiceMoreComing(browser, service, moreComing)
	}
}

// HasNetServiceBrowserDidFindServiceMoreComing returns true if a handler for NetServiceBrowserDidFindServiceMoreComing has been set.
func (d *NetServiceBrowserDelegate) HasNetServiceBrowserDidFindServiceMoreComing() bool {
	return d._NetServiceBrowserDidFindServiceMoreComing != nil
}

// NetServiceBrowserDidFindDomainMoreComing implements the PNetServiceBrowserDelegate interface.
func (d *NetServiceBrowserDelegate) NetServiceBrowserDidFindDomainMoreComing(browser INetServiceBrowser, domainString IString, moreComing bool) {
	if d._NetServiceBrowserDidFindDomainMoreComing != nil {
		d._NetServiceBrowserDidFindDomainMoreComing(browser, domainString, moreComing)
	}
}

// HasNetServiceBrowserDidFindDomainMoreComing returns true if a handler for NetServiceBrowserDidFindDomainMoreComing has been set.
func (d *NetServiceBrowserDelegate) HasNetServiceBrowserDidFindDomainMoreComing() bool {
	return d._NetServiceBrowserDidFindDomainMoreComing != nil
}

// NetServiceBrowserDidNotSearch implements the PNetServiceBrowserDelegate interface.
func (d *NetServiceBrowserDelegate) NetServiceBrowserDidNotSearch(browser INetServiceBrowser, errorDict IDictionary) {
	if d._NetServiceBrowserDidNotSearch != nil {
		d._NetServiceBrowserDidNotSearch(browser, errorDict)
	}
}

// HasNetServiceBrowserDidNotSearch returns true if a handler for NetServiceBrowserDidNotSearch has been set.
func (d *NetServiceBrowserDelegate) HasNetServiceBrowserDidNotSearch() bool {
	return d._NetServiceBrowserDidNotSearch != nil
}

// NetServiceBrowserDidRemoveServiceMoreComing implements the PNetServiceBrowserDelegate interface.
func (d *NetServiceBrowserDelegate) NetServiceBrowserDidRemoveServiceMoreComing(browser INetServiceBrowser, service INetService, moreComing bool) {
	if d._NetServiceBrowserDidRemoveServiceMoreComing != nil {
		d._NetServiceBrowserDidRemoveServiceMoreComing(browser, service, moreComing)
	}
}

// HasNetServiceBrowserDidRemoveServiceMoreComing returns true if a handler for NetServiceBrowserDidRemoveServiceMoreComing has been set.
func (d *NetServiceBrowserDelegate) HasNetServiceBrowserDidRemoveServiceMoreComing() bool {
	return d._NetServiceBrowserDidRemoveServiceMoreComing != nil
}

// NetServiceBrowserDidRemoveDomainMoreComing implements the PNetServiceBrowserDelegate interface.
func (d *NetServiceBrowserDelegate) NetServiceBrowserDidRemoveDomainMoreComing(browser INetServiceBrowser, domainString IString, moreComing bool) {
	if d._NetServiceBrowserDidRemoveDomainMoreComing != nil {
		d._NetServiceBrowserDidRemoveDomainMoreComing(browser, domainString, moreComing)
	}
}

// HasNetServiceBrowserDidRemoveDomainMoreComing returns true if a handler for NetServiceBrowserDidRemoveDomainMoreComing has been set.
func (d *NetServiceBrowserDelegate) HasNetServiceBrowserDidRemoveDomainMoreComing() bool {
	return d._NetServiceBrowserDidRemoveDomainMoreComing != nil
}

// NetServiceBrowserDidStopSearch implements the PNetServiceBrowserDelegate interface.
func (d *NetServiceBrowserDelegate) NetServiceBrowserDidStopSearch(browser INetServiceBrowser) {
	if d._NetServiceBrowserDidStopSearch != nil {
		d._NetServiceBrowserDidStopSearch(browser)
	}
}

// HasNetServiceBrowserDidStopSearch returns true if a handler for NetServiceBrowserDidStopSearch has been set.
func (d *NetServiceBrowserDelegate) HasNetServiceBrowserDidStopSearch() bool {
	return d._NetServiceBrowserDidStopSearch != nil
}

// NetServiceBrowserWillSearch implements the PNetServiceBrowserDelegate interface.
func (d *NetServiceBrowserDelegate) NetServiceBrowserWillSearch(browser INetServiceBrowser) {
	if d._NetServiceBrowserWillSearch != nil {
		d._NetServiceBrowserWillSearch(browser)
	}
}

// HasNetServiceBrowserWillSearch returns true if a handler for NetServiceBrowserWillSearch has been set.
func (d *NetServiceBrowserDelegate) HasNetServiceBrowserWillSearch() bool {
	return d._NetServiceBrowserWillSearch != nil
}
