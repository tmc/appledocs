// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PEAWiFiUnconfiguredAccessoryBrowserDelegate is the EAWiFiUnconfiguredAccessoryBrowserDelegate protocol interface.
//
// A protocol you use to manage the search and configuration processes for an unconfigured accessory browser.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.13+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.externalaccessory/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowserDelegate
type PEAWiFiUnconfiguredAccessoryBrowserDelegate interface {
	// Required methods
	AccessoryBrowserDidFindUnconfiguredAccessories(browser IEAWiFiUnconfiguredAccessoryBrowser, accessories unsafe.Pointer)/* debug [protocol_interface/required_method]: AccessoryBrowserDidFindUnconfiguredAccessories */
	AccessoryBrowserDidFinishConfiguringAccessoryWithStatus(browser IEAWiFiUnconfiguredAccessoryBrowser, accessory IEAWiFiUnconfiguredAccessory, status EAWiFiUnconfiguredAccessoryConfigurationStatus)/* debug [protocol_interface/required_method]: AccessoryBrowserDidFinishConfiguringAccessoryWithStatus */
	AccessoryBrowserDidRemoveUnconfiguredAccessories(browser IEAWiFiUnconfiguredAccessoryBrowser, accessories unsafe.Pointer)/* debug [protocol_interface/required_method]: AccessoryBrowserDidRemoveUnconfiguredAccessories */
	AccessoryBrowserDidUpdateState(browser IEAWiFiUnconfiguredAccessoryBrowser, state EAWiFiUnconfiguredAccessoryBrowserState)/* debug [protocol_interface/required_method]: AccessoryBrowserDidUpdateState */
}

// EAWiFiUnconfiguredAccessoryBrowserDelegate is a delegate implementation builder for the PEAWiFiUnconfiguredAccessoryBrowserDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type EAWiFiUnconfiguredAccessoryBrowserDelegate struct {
	_AccessoryBrowserDidFindUnconfiguredAccessories func(browser IEAWiFiUnconfiguredAccessoryBrowser, accessories unsafe.Pointer)
	_AccessoryBrowserDidFinishConfiguringAccessoryWithStatus func(browser IEAWiFiUnconfiguredAccessoryBrowser, accessory IEAWiFiUnconfiguredAccessory, status EAWiFiUnconfiguredAccessoryConfigurationStatus)
	_AccessoryBrowserDidRemoveUnconfiguredAccessories func(browser IEAWiFiUnconfiguredAccessoryBrowser, accessories unsafe.Pointer)
	_AccessoryBrowserDidUpdateState func(browser IEAWiFiUnconfiguredAccessoryBrowser, state EAWiFiUnconfiguredAccessoryBrowserState)
}

// SetAccessoryBrowserDidFindUnconfiguredAccessories sets the handler for the AccessoryBrowserDidFindUnconfiguredAccessories delegate method.
//
// Indicates that the browser found a new unconfigured accessory that matches the filter predicate defined at the start of the search.
func (d *EAWiFiUnconfiguredAccessoryBrowserDelegate) SetAccessoryBrowserDidFindUnconfiguredAccessories(f func(browser IEAWiFiUnconfiguredAccessoryBrowser, accessories unsafe.Pointer)) {
	d._AccessoryBrowserDidFindUnconfiguredAccessories = f
}

// SetAccessoryBrowserDidFinishConfiguringAccessoryWithStatus sets the handler for the AccessoryBrowserDidFinishConfiguringAccessoryWithStatus delegate method.
//
// Indicates that the browser has completed configuring the specified accessory.
func (d *EAWiFiUnconfiguredAccessoryBrowserDelegate) SetAccessoryBrowserDidFinishConfiguringAccessoryWithStatus(f func(browser IEAWiFiUnconfiguredAccessoryBrowser, accessory IEAWiFiUnconfiguredAccessory, status EAWiFiUnconfiguredAccessoryConfigurationStatus)) {
	d._AccessoryBrowserDidFinishConfiguringAccessoryWithStatus = f
}

// SetAccessoryBrowserDidRemoveUnconfiguredAccessories sets the handler for the AccessoryBrowserDidRemoveUnconfiguredAccessories delegate method.
//
// Indicates that the browser removed an unconfigured accessory from the search results.
func (d *EAWiFiUnconfiguredAccessoryBrowserDelegate) SetAccessoryBrowserDidRemoveUnconfiguredAccessories(f func(browser IEAWiFiUnconfiguredAccessoryBrowser, accessories unsafe.Pointer)) {
	d._AccessoryBrowserDidRemoveUnconfiguredAccessories = f
}

// SetAccessoryBrowserDidUpdateState sets the handler for the AccessoryBrowserDidUpdateState delegate method.
//
// Indicates that the browser’s state has changed.
func (d *EAWiFiUnconfiguredAccessoryBrowserDelegate) SetAccessoryBrowserDidUpdateState(f func(browser IEAWiFiUnconfiguredAccessoryBrowser, state EAWiFiUnconfiguredAccessoryBrowserState)) {
	d._AccessoryBrowserDidUpdateState = f
}

// AccessoryBrowserDidFindUnconfiguredAccessories implements the PEAWiFiUnconfiguredAccessoryBrowserDelegate interface.
func (d *EAWiFiUnconfiguredAccessoryBrowserDelegate) AccessoryBrowserDidFindUnconfiguredAccessories(browser IEAWiFiUnconfiguredAccessoryBrowser, accessories unsafe.Pointer) {
	if d._AccessoryBrowserDidFindUnconfiguredAccessories != nil {
		d._AccessoryBrowserDidFindUnconfiguredAccessories(browser, accessories)
	}
}

// HasAccessoryBrowserDidFindUnconfiguredAccessories returns true if a handler for AccessoryBrowserDidFindUnconfiguredAccessories has been set.
func (d *EAWiFiUnconfiguredAccessoryBrowserDelegate) HasAccessoryBrowserDidFindUnconfiguredAccessories() bool {
	return d._AccessoryBrowserDidFindUnconfiguredAccessories != nil
}

// AccessoryBrowserDidFinishConfiguringAccessoryWithStatus implements the PEAWiFiUnconfiguredAccessoryBrowserDelegate interface.
func (d *EAWiFiUnconfiguredAccessoryBrowserDelegate) AccessoryBrowserDidFinishConfiguringAccessoryWithStatus(browser IEAWiFiUnconfiguredAccessoryBrowser, accessory IEAWiFiUnconfiguredAccessory, status EAWiFiUnconfiguredAccessoryConfigurationStatus) {
	if d._AccessoryBrowserDidFinishConfiguringAccessoryWithStatus != nil {
		d._AccessoryBrowserDidFinishConfiguringAccessoryWithStatus(browser, accessory, status)
	}
}

// HasAccessoryBrowserDidFinishConfiguringAccessoryWithStatus returns true if a handler for AccessoryBrowserDidFinishConfiguringAccessoryWithStatus has been set.
func (d *EAWiFiUnconfiguredAccessoryBrowserDelegate) HasAccessoryBrowserDidFinishConfiguringAccessoryWithStatus() bool {
	return d._AccessoryBrowserDidFinishConfiguringAccessoryWithStatus != nil
}

// AccessoryBrowserDidRemoveUnconfiguredAccessories implements the PEAWiFiUnconfiguredAccessoryBrowserDelegate interface.
func (d *EAWiFiUnconfiguredAccessoryBrowserDelegate) AccessoryBrowserDidRemoveUnconfiguredAccessories(browser IEAWiFiUnconfiguredAccessoryBrowser, accessories unsafe.Pointer) {
	if d._AccessoryBrowserDidRemoveUnconfiguredAccessories != nil {
		d._AccessoryBrowserDidRemoveUnconfiguredAccessories(browser, accessories)
	}
}

// HasAccessoryBrowserDidRemoveUnconfiguredAccessories returns true if a handler for AccessoryBrowserDidRemoveUnconfiguredAccessories has been set.
func (d *EAWiFiUnconfiguredAccessoryBrowserDelegate) HasAccessoryBrowserDidRemoveUnconfiguredAccessories() bool {
	return d._AccessoryBrowserDidRemoveUnconfiguredAccessories != nil
}

// AccessoryBrowserDidUpdateState implements the PEAWiFiUnconfiguredAccessoryBrowserDelegate interface.
func (d *EAWiFiUnconfiguredAccessoryBrowserDelegate) AccessoryBrowserDidUpdateState(browser IEAWiFiUnconfiguredAccessoryBrowser, state EAWiFiUnconfiguredAccessoryBrowserState) {
	if d._AccessoryBrowserDidUpdateState != nil {
		d._AccessoryBrowserDidUpdateState(browser, state)
	}
}

// HasAccessoryBrowserDidUpdateState returns true if a handler for AccessoryBrowserDidUpdateState has been set.
func (d *EAWiFiUnconfiguredAccessoryBrowserDelegate) HasAccessoryBrowserDidUpdateState() bool {
	return d._AccessoryBrowserDidUpdateState != nil
}
