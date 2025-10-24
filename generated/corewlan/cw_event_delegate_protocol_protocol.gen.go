// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PCWEventDelegate is the CWEventDelegate protocol interface.
//
// The interface a Wi-Fi client object uses to notify its delegate about Wi-Fi events.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.6+
//
// See: doc://com.apple.corewlan/documentation/CoreWLAN/CWEventDelegate
type PCWEventDelegate interface {
	// Optional methods
	BssidDidChangeForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */)
	HasBssidDidChangeForWiFiInterfaceWithName() bool
	ClientConnectionInterrupted()
	HasClientConnectionInterrupted() bool
	ClientConnectionInvalidated()
	HasClientConnectionInvalidated() bool
	CountryCodeDidChangeForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */)
	HasCountryCodeDidChangeForWiFiInterfaceWithName() bool
	LinkDidChangeForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */)
	HasLinkDidChangeForWiFiInterfaceWithName() bool
	LinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate(interfaceName objc.IObject /* cross-framework: NSString */, rssi int, transmitRate float64)
	HasLinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate() bool
	ModeDidChangeForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */)
	HasModeDidChangeForWiFiInterfaceWithName() bool
	PowerStateDidChangeForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */)
	HasPowerStateDidChangeForWiFiInterfaceWithName() bool
	ScanCacheUpdatedForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */)
	HasScanCacheUpdatedForWiFiInterfaceWithName() bool
	SsidDidChangeForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */)
	HasSsidDidChangeForWiFiInterfaceWithName() bool
}

// CWEventDelegate is a delegate implementation builder for the PCWEventDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CWEventDelegate struct {
	_BssidDidChangeForWiFiInterfaceWithName func(interfaceName objc.IObject /* cross-framework: NSString */)
	_ClientConnectionInterrupted func()
	_ClientConnectionInvalidated func()
	_CountryCodeDidChangeForWiFiInterfaceWithName func(interfaceName objc.IObject /* cross-framework: NSString */)
	_LinkDidChangeForWiFiInterfaceWithName func(interfaceName objc.IObject /* cross-framework: NSString */)
	_LinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate func(interfaceName objc.IObject /* cross-framework: NSString */, rssi int, transmitRate float64)
	_ModeDidChangeForWiFiInterfaceWithName func(interfaceName objc.IObject /* cross-framework: NSString */)
	_PowerStateDidChangeForWiFiInterfaceWithName func(interfaceName objc.IObject /* cross-framework: NSString */)
	_ScanCacheUpdatedForWiFiInterfaceWithName func(interfaceName objc.IObject /* cross-framework: NSString */)
	_SsidDidChangeForWiFiInterfaceWithName func(interfaceName objc.IObject /* cross-framework: NSString */)
}

// SetBssidDidChangeForWiFiInterfaceWithName sets the handler for the BssidDidChangeForWiFiInterfaceWithName delegate method.
//
// Tells the delegate that the current BSSID has changed.
func (d *CWEventDelegate) SetBssidDidChangeForWiFiInterfaceWithName(f func(interfaceName objc.IObject /* cross-framework: NSString */)) {
	d._BssidDidChangeForWiFiInterfaceWithName = f
}

// SetClientConnectionInterrupted sets the handler for the ClientConnectionInterrupted delegate method.
//
// Tells the delegate that the connection to the Wi-Fi subsystem is temporarily interrupted.
func (d *CWEventDelegate) SetClientConnectionInterrupted(f func()) {
	d._ClientConnectionInterrupted = f
}

// SetClientConnectionInvalidated sets the handler for the ClientConnectionInvalidated delegate method.
//
// Tells the delegate that the connection to the Wi-Fi subsystem is permanently invalidated.
func (d *CWEventDelegate) SetClientConnectionInvalidated(f func()) {
	d._ClientConnectionInvalidated = f
}

// SetCountryCodeDidChangeForWiFiInterfaceWithName sets the handler for the CountryCodeDidChangeForWiFiInterfaceWithName delegate method.
//
// Tells the delegate that the currently adopted country code has changed.
func (d *CWEventDelegate) SetCountryCodeDidChangeForWiFiInterfaceWithName(f func(interfaceName objc.IObject /* cross-framework: NSString */)) {
	d._CountryCodeDidChangeForWiFiInterfaceWithName = f
}

// SetLinkDidChangeForWiFiInterfaceWithName sets the handler for the LinkDidChangeForWiFiInterfaceWithName delegate method.
//
// Tells the delegate that the Wi-Fi link state changed.
func (d *CWEventDelegate) SetLinkDidChangeForWiFiInterfaceWithName(f func(interfaceName objc.IObject /* cross-framework: NSString */)) {
	d._LinkDidChangeForWiFiInterfaceWithName = f
}

// SetLinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate sets the handler for the LinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate delegate method.
//
// Tells the delegate that the link quality has changed.
func (d *CWEventDelegate) SetLinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate(f func(interfaceName objc.IObject /* cross-framework: NSString */, rssi int, transmitRate float64)) {
	d._LinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate = f
}

// SetModeDidChangeForWiFiInterfaceWithName sets the handler for the ModeDidChangeForWiFiInterfaceWithName delegate method.
//
// Tells the delegate that the operating mode has changed.
func (d *CWEventDelegate) SetModeDidChangeForWiFiInterfaceWithName(f func(interfaceName objc.IObject /* cross-framework: NSString */)) {
	d._ModeDidChangeForWiFiInterfaceWithName = f
}

// SetPowerStateDidChangeForWiFiInterfaceWithName sets the handler for the PowerStateDidChangeForWiFiInterfaceWithName delegate method.
//
// Tells the delegate that the Wi-Fi power state changed.
func (d *CWEventDelegate) SetPowerStateDidChangeForWiFiInterfaceWithName(f func(interfaceName objc.IObject /* cross-framework: NSString */)) {
	d._PowerStateDidChangeForWiFiInterfaceWithName = f
}

// SetScanCacheUpdatedForWiFiInterfaceWithName sets the handler for the ScanCacheUpdatedForWiFiInterfaceWithName delegate method.
//
// Tells the delegate that the Wi-Fi interface’s scan cache has been updated with new results.
func (d *CWEventDelegate) SetScanCacheUpdatedForWiFiInterfaceWithName(f func(interfaceName objc.IObject /* cross-framework: NSString */)) {
	d._ScanCacheUpdatedForWiFiInterfaceWithName = f
}

// SetSsidDidChangeForWiFiInterfaceWithName sets the handler for the SsidDidChangeForWiFiInterfaceWithName delegate method.
//
// Tells the delegate that the current SSID has changed.
func (d *CWEventDelegate) SetSsidDidChangeForWiFiInterfaceWithName(f func(interfaceName objc.IObject /* cross-framework: NSString */)) {
	d._SsidDidChangeForWiFiInterfaceWithName = f
}

// BssidDidChangeForWiFiInterfaceWithName implements the PCWEventDelegate interface.
func (d *CWEventDelegate) BssidDidChangeForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */) {
	if d._BssidDidChangeForWiFiInterfaceWithName != nil {
		d._BssidDidChangeForWiFiInterfaceWithName(interfaceName)
	}
}

// HasBssidDidChangeForWiFiInterfaceWithName returns true if a handler for BssidDidChangeForWiFiInterfaceWithName has been set.
func (d *CWEventDelegate) HasBssidDidChangeForWiFiInterfaceWithName() bool {
	return d._BssidDidChangeForWiFiInterfaceWithName != nil
}

// ClientConnectionInterrupted implements the PCWEventDelegate interface.
func (d *CWEventDelegate) ClientConnectionInterrupted() {
	if d._ClientConnectionInterrupted != nil {
		d._ClientConnectionInterrupted()
	}
}

// HasClientConnectionInterrupted returns true if a handler for ClientConnectionInterrupted has been set.
func (d *CWEventDelegate) HasClientConnectionInterrupted() bool {
	return d._ClientConnectionInterrupted != nil
}

// ClientConnectionInvalidated implements the PCWEventDelegate interface.
func (d *CWEventDelegate) ClientConnectionInvalidated() {
	if d._ClientConnectionInvalidated != nil {
		d._ClientConnectionInvalidated()
	}
}

// HasClientConnectionInvalidated returns true if a handler for ClientConnectionInvalidated has been set.
func (d *CWEventDelegate) HasClientConnectionInvalidated() bool {
	return d._ClientConnectionInvalidated != nil
}

// CountryCodeDidChangeForWiFiInterfaceWithName implements the PCWEventDelegate interface.
func (d *CWEventDelegate) CountryCodeDidChangeForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */) {
	if d._CountryCodeDidChangeForWiFiInterfaceWithName != nil {
		d._CountryCodeDidChangeForWiFiInterfaceWithName(interfaceName)
	}
}

// HasCountryCodeDidChangeForWiFiInterfaceWithName returns true if a handler for CountryCodeDidChangeForWiFiInterfaceWithName has been set.
func (d *CWEventDelegate) HasCountryCodeDidChangeForWiFiInterfaceWithName() bool {
	return d._CountryCodeDidChangeForWiFiInterfaceWithName != nil
}

// LinkDidChangeForWiFiInterfaceWithName implements the PCWEventDelegate interface.
func (d *CWEventDelegate) LinkDidChangeForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */) {
	if d._LinkDidChangeForWiFiInterfaceWithName != nil {
		d._LinkDidChangeForWiFiInterfaceWithName(interfaceName)
	}
}

// HasLinkDidChangeForWiFiInterfaceWithName returns true if a handler for LinkDidChangeForWiFiInterfaceWithName has been set.
func (d *CWEventDelegate) HasLinkDidChangeForWiFiInterfaceWithName() bool {
	return d._LinkDidChangeForWiFiInterfaceWithName != nil
}

// LinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate implements the PCWEventDelegate interface.
func (d *CWEventDelegate) LinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate(interfaceName objc.IObject /* cross-framework: NSString */, rssi int, transmitRate float64) {
	if d._LinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate != nil {
		d._LinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate(interfaceName, rssi, transmitRate)
	}
}

// HasLinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate returns true if a handler for LinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate has been set.
func (d *CWEventDelegate) HasLinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate() bool {
	return d._LinkQualityDidChangeForWiFiInterfaceWithNameRssiTransmitRate != nil
}

// ModeDidChangeForWiFiInterfaceWithName implements the PCWEventDelegate interface.
func (d *CWEventDelegate) ModeDidChangeForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */) {
	if d._ModeDidChangeForWiFiInterfaceWithName != nil {
		d._ModeDidChangeForWiFiInterfaceWithName(interfaceName)
	}
}

// HasModeDidChangeForWiFiInterfaceWithName returns true if a handler for ModeDidChangeForWiFiInterfaceWithName has been set.
func (d *CWEventDelegate) HasModeDidChangeForWiFiInterfaceWithName() bool {
	return d._ModeDidChangeForWiFiInterfaceWithName != nil
}

// PowerStateDidChangeForWiFiInterfaceWithName implements the PCWEventDelegate interface.
func (d *CWEventDelegate) PowerStateDidChangeForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */) {
	if d._PowerStateDidChangeForWiFiInterfaceWithName != nil {
		d._PowerStateDidChangeForWiFiInterfaceWithName(interfaceName)
	}
}

// HasPowerStateDidChangeForWiFiInterfaceWithName returns true if a handler for PowerStateDidChangeForWiFiInterfaceWithName has been set.
func (d *CWEventDelegate) HasPowerStateDidChangeForWiFiInterfaceWithName() bool {
	return d._PowerStateDidChangeForWiFiInterfaceWithName != nil
}

// ScanCacheUpdatedForWiFiInterfaceWithName implements the PCWEventDelegate interface.
func (d *CWEventDelegate) ScanCacheUpdatedForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */) {
	if d._ScanCacheUpdatedForWiFiInterfaceWithName != nil {
		d._ScanCacheUpdatedForWiFiInterfaceWithName(interfaceName)
	}
}

// HasScanCacheUpdatedForWiFiInterfaceWithName returns true if a handler for ScanCacheUpdatedForWiFiInterfaceWithName has been set.
func (d *CWEventDelegate) HasScanCacheUpdatedForWiFiInterfaceWithName() bool {
	return d._ScanCacheUpdatedForWiFiInterfaceWithName != nil
}

// SsidDidChangeForWiFiInterfaceWithName implements the PCWEventDelegate interface.
func (d *CWEventDelegate) SsidDidChangeForWiFiInterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */) {
	if d._SsidDidChangeForWiFiInterfaceWithName != nil {
		d._SsidDidChangeForWiFiInterfaceWithName(interfaceName)
	}
}

// HasSsidDidChangeForWiFiInterfaceWithName returns true if a handler for SsidDidChangeForWiFiInterfaceWithName has been set.
func (d *CWEventDelegate) HasSsidDidChangeForWiFiInterfaceWithName() bool {
	return d._SsidDidChangeForWiFiInterfaceWithName != nil
}
