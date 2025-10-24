// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PMTRDevicePairingDelegate is the MTRDevicePairingDelegate protocol interface.
//
// Availability:
//   - Mac Catalyst 16.1+ (Deprecated in 16.4)
//   - iOS 16.1+ (Deprecated in 16.4)
//   - iPadOS 16.1+ (Deprecated in 16.4)
//   - macOS 13.0+ (Deprecated in 13.3)
//   - tvOS 16.1+ (Deprecated in 16.4)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 9.1+ (Deprecated in 9.4)
//
// See: doc://com.apple.matter/documentation/Matter/MTRDevicePairingDelegate
type PMTRDevicePairingDelegate interface {
	// Optional methods
	OnCommissioningComplete(error_ objc.IObject /* cross-framework: Error */)
	HasOnCommissioningComplete() bool
	OnPairingComplete(error_ objc.IObject /* cross-framework: Error */)
	HasOnPairingComplete() bool
	OnPairingDeleted(error_ objc.IObject /* cross-framework: Error */)
	HasOnPairingDeleted() bool
	OnStatusUpdate(status unsafe.Pointer)
	HasOnStatusUpdate() bool
}

// MTRDevicePairingDelegate is a delegate implementation builder for the PMTRDevicePairingDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MTRDevicePairingDelegate struct {
	_OnCommissioningComplete func(error_ objc.IObject /* cross-framework: Error */)
	_OnPairingComplete func(error_ objc.IObject /* cross-framework: Error */)
	_OnPairingDeleted func(error_ objc.IObject /* cross-framework: Error */)
	_OnStatusUpdate func(status unsafe.Pointer)
}

// SetOnCommissioningComplete sets the handler for the OnCommissioningComplete delegate method.
func (d *MTRDevicePairingDelegate) SetOnCommissioningComplete(f func(error_ objc.IObject /* cross-framework: Error */)) {
	d._OnCommissioningComplete = f
}

// SetOnPairingComplete sets the handler for the OnPairingComplete delegate method.
func (d *MTRDevicePairingDelegate) SetOnPairingComplete(f func(error_ objc.IObject /* cross-framework: Error */)) {
	d._OnPairingComplete = f
}

// SetOnPairingDeleted sets the handler for the OnPairingDeleted delegate method.
func (d *MTRDevicePairingDelegate) SetOnPairingDeleted(f func(error_ objc.IObject /* cross-framework: Error */)) {
	d._OnPairingDeleted = f
}

// SetOnStatusUpdate sets the handler for the OnStatusUpdate delegate method.
func (d *MTRDevicePairingDelegate) SetOnStatusUpdate(f func(status unsafe.Pointer)) {
	d._OnStatusUpdate = f
}

// OnCommissioningComplete implements the PMTRDevicePairingDelegate interface.
func (d *MTRDevicePairingDelegate) OnCommissioningComplete(error_ objc.IObject /* cross-framework: Error */) {
	if d._OnCommissioningComplete != nil {
		d._OnCommissioningComplete(error_)
	}
}

// HasOnCommissioningComplete returns true if a handler for OnCommissioningComplete has been set.
func (d *MTRDevicePairingDelegate) HasOnCommissioningComplete() bool {
	return d._OnCommissioningComplete != nil
}

// OnPairingComplete implements the PMTRDevicePairingDelegate interface.
func (d *MTRDevicePairingDelegate) OnPairingComplete(error_ objc.IObject /* cross-framework: Error */) {
	if d._OnPairingComplete != nil {
		d._OnPairingComplete(error_)
	}
}

// HasOnPairingComplete returns true if a handler for OnPairingComplete has been set.
func (d *MTRDevicePairingDelegate) HasOnPairingComplete() bool {
	return d._OnPairingComplete != nil
}

// OnPairingDeleted implements the PMTRDevicePairingDelegate interface.
func (d *MTRDevicePairingDelegate) OnPairingDeleted(error_ objc.IObject /* cross-framework: Error */) {
	if d._OnPairingDeleted != nil {
		d._OnPairingDeleted(error_)
	}
}

// HasOnPairingDeleted returns true if a handler for OnPairingDeleted has been set.
func (d *MTRDevicePairingDelegate) HasOnPairingDeleted() bool {
	return d._OnPairingDeleted != nil
}

// OnStatusUpdate implements the PMTRDevicePairingDelegate interface.
func (d *MTRDevicePairingDelegate) OnStatusUpdate(status unsafe.Pointer) {
	if d._OnStatusUpdate != nil {
		d._OnStatusUpdate(status)
	}
}

// HasOnStatusUpdate returns true if a handler for OnStatusUpdate has been set.
func (d *MTRDevicePairingDelegate) HasOnStatusUpdate() bool {
	return d._OnStatusUpdate != nil
}
