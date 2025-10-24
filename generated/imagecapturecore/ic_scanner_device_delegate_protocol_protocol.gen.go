// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PICScannerDeviceDelegate is the ICScannerDeviceDelegate protocol interface.
//
// Methods for determining availability, selecting a functional unit, and performing scans on connected scanners.
//
// Availability:
//   - macOS 10.4+
//
// See: doc://com.apple.imagecapturecore/documentation/ImageCaptureCore/ICScannerDeviceDelegate
type PICScannerDeviceDelegate interface {
	// Required methods
	ScannerDeviceDidScanToURLData(scanner ICScannerDevice, url foundation.URL, data foundation.Data)/* debug [protocol_interface/required_method]: ScannerDeviceDidScanToURLData */
	// Optional methods
	ScannerDevice()
	HasScannerDevice() bool
	ScannerDeviceDidBecomeAvailable()
	HasScannerDeviceDidBecomeAvailable() bool
	ScannerDeviceDidCompleteOverviewScanWithError(scanner ICScannerDevice, error_ objc.IObject /* cross-framework: Error */)
	HasScannerDeviceDidCompleteOverviewScanWithError() bool
	ScannerDeviceDidCompleteScanWithError(scanner ICScannerDevice, error_ objc.IObject /* cross-framework: Error */)
	HasScannerDeviceDidCompleteScanWithError() bool
	ScannerDeviceDidScanToURL(scanner ICScannerDevice, url objc.IObject /* cross-framework: NSURL */)
	HasScannerDeviceDidScanToURL() bool
	ScannerDeviceDidScanToBandData(scanner ICScannerDevice, data ICScannerBandData)
	HasScannerDeviceDidScanToBandData() bool
	ScannerDeviceDidSelectFunctionalUnitError(scanner ICScannerDevice, functionalUnit ICScannerFunctionalUnit, error_ objc.IObject /* cross-framework: Error */)
	HasScannerDeviceDidSelectFunctionalUnitError() bool
}

// ICScannerDeviceDelegate is a delegate implementation builder for the PICScannerDeviceDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ICScannerDeviceDelegate struct {
	_ScannerDevice func()
	_ScannerDeviceDidBecomeAvailable func()
	_ScannerDeviceDidCompleteOverviewScanWithError func(scanner ICScannerDevice, error_ objc.IObject /* cross-framework: Error */)
	_ScannerDeviceDidCompleteScanWithError func(scanner ICScannerDevice, error_ objc.IObject /* cross-framework: Error */)
	_ScannerDeviceDidScanToURL func(scanner ICScannerDevice, url objc.IObject /* cross-framework: NSURL */)
	_ScannerDeviceDidScanToBandData func(scanner ICScannerDevice, data ICScannerBandData)
	_ScannerDeviceDidSelectFunctionalUnitError func(scanner ICScannerDevice, functionalUnit ICScannerFunctionalUnit, error_ objc.IObject /* cross-framework: Error */)
	_ScannerDeviceDidScanToURLData func(scanner ICScannerDevice, url foundation.URL, data foundation.Data)
}

// SetScannerDevice sets the handler for the ScannerDevice delegate method.
//
// Tells the client when a functional unit is selected on the scanner.
func (d *ICScannerDeviceDelegate) SetScannerDevice(f func()) {
	d._ScannerDevice = f
}

// SetScannerDeviceDidBecomeAvailable sets the handler for the ScannerDeviceDidBecomeAvailable delegate method.
//
// Tells the client when another client closes the current open session on the scanner.
func (d *ICScannerDeviceDelegate) SetScannerDeviceDidBecomeAvailable(f func()) {
	d._ScannerDeviceDidBecomeAvailable = f
}

// SetScannerDeviceDidCompleteOverviewScanWithError sets the handler for the ScannerDeviceDidCompleteOverviewScanWithError delegate method.
//
// Tells the client when the scanner completes an overview scan.
func (d *ICScannerDeviceDelegate) SetScannerDeviceDidCompleteOverviewScanWithError(f func(scanner ICScannerDevice, error_ objc.IObject /* cross-framework: Error */)) {
	d._ScannerDeviceDidCompleteOverviewScanWithError = f
}

// SetScannerDeviceDidCompleteScanWithError sets the handler for the ScannerDeviceDidCompleteScanWithError delegate method.
//
// Tells the client when the scanner completes a scan.
func (d *ICScannerDeviceDelegate) SetScannerDeviceDidCompleteScanWithError(f func(scanner ICScannerDevice, error_ objc.IObject /* cross-framework: Error */)) {
	d._ScannerDeviceDidCompleteScanWithError = f
}

// SetScannerDeviceDidScanToURL sets the handler for the ScannerDeviceDidScanToURL delegate method.
//
// Tells the client when the scanner receives the requested scan.
func (d *ICScannerDeviceDelegate) SetScannerDeviceDidScanToURL(f func(scanner ICScannerDevice, url objc.IObject /* cross-framework: NSURL */)) {
	d._ScannerDeviceDidScanToURL = f
}

// SetScannerDeviceDidScanToBandData sets the handler for the ScannerDeviceDidScanToBandData delegate method.
//
// Tells the client when the scanner receives the requested scan progress notification and a band of data is sent for each notification received.
func (d *ICScannerDeviceDelegate) SetScannerDeviceDidScanToBandData(f func(scanner ICScannerDevice, data ICScannerBandData)) {
	d._ScannerDeviceDidScanToBandData = f
}

// SetScannerDeviceDidSelectFunctionalUnitError sets the handler for the ScannerDeviceDidSelectFunctionalUnitError delegate method.
//
// Tells the client when a functional unit is selected on the scanner.
func (d *ICScannerDeviceDelegate) SetScannerDeviceDidSelectFunctionalUnitError(f func(scanner ICScannerDevice, functionalUnit ICScannerFunctionalUnit, error_ objc.IObject /* cross-framework: Error */)) {
	d._ScannerDeviceDidSelectFunctionalUnitError = f
}

// SetScannerDeviceDidScanToURLData sets the handler for the ScannerDeviceDidScanToURLData delegate method.
//
// Tells the client when the scanner device receives the requested scan. 
func (d *ICScannerDeviceDelegate) SetScannerDeviceDidScanToURLData(f func(scanner ICScannerDevice, url foundation.URL, data foundation.Data)) {
	d._ScannerDeviceDidScanToURLData = f
}

// ScannerDevice implements the PICScannerDeviceDelegate interface.
func (d *ICScannerDeviceDelegate) ScannerDevice() {
	if d._ScannerDevice != nil {
		d._ScannerDevice()
	}
}

// HasScannerDevice returns true if a handler for ScannerDevice has been set.
func (d *ICScannerDeviceDelegate) HasScannerDevice() bool {
	return d._ScannerDevice != nil
}

// ScannerDeviceDidBecomeAvailable implements the PICScannerDeviceDelegate interface.
func (d *ICScannerDeviceDelegate) ScannerDeviceDidBecomeAvailable() {
	if d._ScannerDeviceDidBecomeAvailable != nil {
		d._ScannerDeviceDidBecomeAvailable()
	}
}

// HasScannerDeviceDidBecomeAvailable returns true if a handler for ScannerDeviceDidBecomeAvailable has been set.
func (d *ICScannerDeviceDelegate) HasScannerDeviceDidBecomeAvailable() bool {
	return d._ScannerDeviceDidBecomeAvailable != nil
}

// ScannerDeviceDidCompleteOverviewScanWithError implements the PICScannerDeviceDelegate interface.
func (d *ICScannerDeviceDelegate) ScannerDeviceDidCompleteOverviewScanWithError(scanner ICScannerDevice, error_ objc.IObject /* cross-framework: Error */) {
	if d._ScannerDeviceDidCompleteOverviewScanWithError != nil {
		d._ScannerDeviceDidCompleteOverviewScanWithError(scanner, error_)
	}
}

// HasScannerDeviceDidCompleteOverviewScanWithError returns true if a handler for ScannerDeviceDidCompleteOverviewScanWithError has been set.
func (d *ICScannerDeviceDelegate) HasScannerDeviceDidCompleteOverviewScanWithError() bool {
	return d._ScannerDeviceDidCompleteOverviewScanWithError != nil
}

// ScannerDeviceDidCompleteScanWithError implements the PICScannerDeviceDelegate interface.
func (d *ICScannerDeviceDelegate) ScannerDeviceDidCompleteScanWithError(scanner ICScannerDevice, error_ objc.IObject /* cross-framework: Error */) {
	if d._ScannerDeviceDidCompleteScanWithError != nil {
		d._ScannerDeviceDidCompleteScanWithError(scanner, error_)
	}
}

// HasScannerDeviceDidCompleteScanWithError returns true if a handler for ScannerDeviceDidCompleteScanWithError has been set.
func (d *ICScannerDeviceDelegate) HasScannerDeviceDidCompleteScanWithError() bool {
	return d._ScannerDeviceDidCompleteScanWithError != nil
}

// ScannerDeviceDidScanToURL implements the PICScannerDeviceDelegate interface.
func (d *ICScannerDeviceDelegate) ScannerDeviceDidScanToURL(scanner ICScannerDevice, url objc.IObject /* cross-framework: NSURL */) {
	if d._ScannerDeviceDidScanToURL != nil {
		d._ScannerDeviceDidScanToURL(scanner, url)
	}
}

// HasScannerDeviceDidScanToURL returns true if a handler for ScannerDeviceDidScanToURL has been set.
func (d *ICScannerDeviceDelegate) HasScannerDeviceDidScanToURL() bool {
	return d._ScannerDeviceDidScanToURL != nil
}

// ScannerDeviceDidScanToBandData implements the PICScannerDeviceDelegate interface.
func (d *ICScannerDeviceDelegate) ScannerDeviceDidScanToBandData(scanner ICScannerDevice, data ICScannerBandData) {
	if d._ScannerDeviceDidScanToBandData != nil {
		d._ScannerDeviceDidScanToBandData(scanner, data)
	}
}

// HasScannerDeviceDidScanToBandData returns true if a handler for ScannerDeviceDidScanToBandData has been set.
func (d *ICScannerDeviceDelegate) HasScannerDeviceDidScanToBandData() bool {
	return d._ScannerDeviceDidScanToBandData != nil
}

// ScannerDeviceDidSelectFunctionalUnitError implements the PICScannerDeviceDelegate interface.
func (d *ICScannerDeviceDelegate) ScannerDeviceDidSelectFunctionalUnitError(scanner ICScannerDevice, functionalUnit ICScannerFunctionalUnit, error_ objc.IObject /* cross-framework: Error */) {
	if d._ScannerDeviceDidSelectFunctionalUnitError != nil {
		d._ScannerDeviceDidSelectFunctionalUnitError(scanner, functionalUnit, error_)
	}
}

// HasScannerDeviceDidSelectFunctionalUnitError returns true if a handler for ScannerDeviceDidSelectFunctionalUnitError has been set.
func (d *ICScannerDeviceDelegate) HasScannerDeviceDidSelectFunctionalUnitError() bool {
	return d._ScannerDeviceDidSelectFunctionalUnitError != nil
}

// ScannerDeviceDidScanToURLData implements the PICScannerDeviceDelegate interface.
func (d *ICScannerDeviceDelegate) ScannerDeviceDidScanToURLData(scanner ICScannerDevice, url foundation.URL, data foundation.Data) {
	if d._ScannerDeviceDidScanToURLData != nil {
		d._ScannerDeviceDidScanToURLData(scanner, url, data)
	}
}

// HasScannerDeviceDidScanToURLData returns true if a handler for ScannerDeviceDidScanToURLData has been set.
func (d *ICScannerDeviceDelegate) HasScannerDeviceDidScanToURLData() bool {
	return d._ScannerDeviceDidScanToURLData != nil
}
