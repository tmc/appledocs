// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/imagecapturecore"
)

// PIKScannerDeviceViewDelegate is the IKScannerDeviceViewDelegate protocol interface.
//
// The   protocol defines the delegate protocol that the   delegate must conform to.
//
// Availability:
//   - macOS 10.4+
//
// See: doc://com.apple.quartz/documentation/Quartz/IKScannerDeviceViewDelegate
type PIKScannerDeviceViewDelegate interface {
	// Optional methods
	ScannerDeviceViewDidEncounterError(scannerDeviceView IKScannerDeviceView, error_ objc.IObject /* cross-framework: Error */)
	HasScannerDeviceViewDidEncounterError() bool
	ScannerDeviceViewDidScanToURLError(scannerDeviceView IKScannerDeviceView, url objc.IObject /* cross-framework: NSURL */, error_ objc.IObject /* cross-framework: Error */)
	HasScannerDeviceViewDidScanToURLError() bool
	ScannerDeviceViewDidScanToURLFileDataError(scannerDeviceView IKScannerDeviceView, url objc.IObject /* cross-framework: NSURL */, data objc.IObject /* cross-framework: NSData */, error_ objc.IObject /* cross-framework: Error */)
	HasScannerDeviceViewDidScanToURLFileDataError() bool
	ScannerDeviceViewDidScanToBandDataScanInfoError(scannerDeviceView IKScannerDeviceView, data objc.IObject, scanInfo objc.IObject /* cross-framework: NSDictionary */, error_ objc.IObject /* cross-framework: Error */)
	HasScannerDeviceViewDidScanToBandDataScanInfoError() bool
}

// IKScannerDeviceViewDelegate is a delegate implementation builder for the PIKScannerDeviceViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type IKScannerDeviceViewDelegate struct {
	_ScannerDeviceViewDidEncounterError func(scannerDeviceView IKScannerDeviceView, error_ objc.IObject /* cross-framework: Error */)
	_ScannerDeviceViewDidScanToURLError func(scannerDeviceView IKScannerDeviceView, url objc.IObject /* cross-framework: NSURL */, error_ objc.IObject /* cross-framework: Error */)
	_ScannerDeviceViewDidScanToURLFileDataError func(scannerDeviceView IKScannerDeviceView, url objc.IObject /* cross-framework: NSURL */, data objc.IObject /* cross-framework: NSData */, error_ objc.IObject /* cross-framework: Error */)
	_ScannerDeviceViewDidScanToBandDataScanInfoError func(scannerDeviceView IKScannerDeviceView, data objc.IObject, scanInfo objc.IObject /* cross-framework: NSDictionary */, error_ objc.IObject /* cross-framework: Error */)
}

// SetScannerDeviceViewDidEncounterError sets the handler for the ScannerDeviceViewDidEncounterError delegate method.
//
// Invoked whenever the scanner encounters an error.
func (d *IKScannerDeviceViewDelegate) SetScannerDeviceViewDidEncounterError(f func(scannerDeviceView IKScannerDeviceView, error_ objc.IObject /* cross-framework: Error */)) {
	d._ScannerDeviceViewDidEncounterError = f
}

// SetScannerDeviceViewDidScanToURLError sets the handler for the ScannerDeviceViewDidScanToURLError delegate method.
func (d *IKScannerDeviceViewDelegate) SetScannerDeviceViewDidScanToURLError(f func(scannerDeviceView IKScannerDeviceView, url objc.IObject /* cross-framework: NSURL */, error_ objc.IObject /* cross-framework: Error */)) {
	d._ScannerDeviceViewDidScanToURLError = f
}

// SetScannerDeviceViewDidScanToURLFileDataError sets the handler for the ScannerDeviceViewDidScanToURLFileDataError delegate method.
//
// Invoked when the scan has completed and the data is available.
func (d *IKScannerDeviceViewDelegate) SetScannerDeviceViewDidScanToURLFileDataError(f func(scannerDeviceView IKScannerDeviceView, url objc.IObject /* cross-framework: NSURL */, data objc.IObject /* cross-framework: NSData */, error_ objc.IObject /* cross-framework: Error */)) {
	d._ScannerDeviceViewDidScanToURLFileDataError = f
}

// SetScannerDeviceViewDidScanToBandDataScanInfoError sets the handler for the ScannerDeviceViewDidScanToBandDataScanInfoError delegate method.
func (d *IKScannerDeviceViewDelegate) SetScannerDeviceViewDidScanToBandDataScanInfoError(f func(scannerDeviceView IKScannerDeviceView, data objc.IObject, scanInfo objc.IObject /* cross-framework: NSDictionary */, error_ objc.IObject /* cross-framework: Error */)) {
	d._ScannerDeviceViewDidScanToBandDataScanInfoError = f
}

// ScannerDeviceViewDidEncounterError implements the PIKScannerDeviceViewDelegate interface.
func (d *IKScannerDeviceViewDelegate) ScannerDeviceViewDidEncounterError(scannerDeviceView IKScannerDeviceView, error_ objc.IObject /* cross-framework: Error */) {
	if d._ScannerDeviceViewDidEncounterError != nil {
		d._ScannerDeviceViewDidEncounterError(scannerDeviceView, error_)
	}
}

// HasScannerDeviceViewDidEncounterError returns true if a handler for ScannerDeviceViewDidEncounterError has been set.
func (d *IKScannerDeviceViewDelegate) HasScannerDeviceViewDidEncounterError() bool {
	return d._ScannerDeviceViewDidEncounterError != nil
}

// ScannerDeviceViewDidScanToURLError implements the PIKScannerDeviceViewDelegate interface.
func (d *IKScannerDeviceViewDelegate) ScannerDeviceViewDidScanToURLError(scannerDeviceView IKScannerDeviceView, url objc.IObject /* cross-framework: NSURL */, error_ objc.IObject /* cross-framework: Error */) {
	if d._ScannerDeviceViewDidScanToURLError != nil {
		d._ScannerDeviceViewDidScanToURLError(scannerDeviceView, url, error_)
	}
}

// HasScannerDeviceViewDidScanToURLError returns true if a handler for ScannerDeviceViewDidScanToURLError has been set.
func (d *IKScannerDeviceViewDelegate) HasScannerDeviceViewDidScanToURLError() bool {
	return d._ScannerDeviceViewDidScanToURLError != nil
}

// ScannerDeviceViewDidScanToURLFileDataError implements the PIKScannerDeviceViewDelegate interface.
func (d *IKScannerDeviceViewDelegate) ScannerDeviceViewDidScanToURLFileDataError(scannerDeviceView IKScannerDeviceView, url objc.IObject /* cross-framework: NSURL */, data objc.IObject /* cross-framework: NSData */, error_ objc.IObject /* cross-framework: Error */) {
	if d._ScannerDeviceViewDidScanToURLFileDataError != nil {
		d._ScannerDeviceViewDidScanToURLFileDataError(scannerDeviceView, url, data, error_)
	}
}

// HasScannerDeviceViewDidScanToURLFileDataError returns true if a handler for ScannerDeviceViewDidScanToURLFileDataError has been set.
func (d *IKScannerDeviceViewDelegate) HasScannerDeviceViewDidScanToURLFileDataError() bool {
	return d._ScannerDeviceViewDidScanToURLFileDataError != nil
}

// ScannerDeviceViewDidScanToBandDataScanInfoError implements the PIKScannerDeviceViewDelegate interface.
func (d *IKScannerDeviceViewDelegate) ScannerDeviceViewDidScanToBandDataScanInfoError(scannerDeviceView IKScannerDeviceView, data objc.IObject, scanInfo objc.IObject /* cross-framework: NSDictionary */, error_ objc.IObject /* cross-framework: Error */) {
	if d._ScannerDeviceViewDidScanToBandDataScanInfoError != nil {
		d._ScannerDeviceViewDidScanToBandDataScanInfoError(scannerDeviceView, data, scanInfo, error_)
	}
}

// HasScannerDeviceViewDidScanToBandDataScanInfoError returns true if a handler for ScannerDeviceViewDidScanToBandDataScanInfoError has been set.
func (d *IKScannerDeviceViewDelegate) HasScannerDeviceViewDidScanToBandDataScanInfoError() bool {
	return d._ScannerDeviceViewDidScanToBandDataScanInfoError != nil
}
