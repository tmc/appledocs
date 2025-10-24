// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PWaterSubmersionManagerDelegate is the CMWaterSubmersionManagerDelegate protocol interface.
//
// A delegate that receives updates about ambient pressure, water pressure, water temperature, and submersion events.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+
//
// See: doc://com.apple.coremotion/documentation/CoreMotion/CMWaterSubmersionManagerDelegate
type PWaterSubmersionManagerDelegate interface {
	// Required methods
	ManagerDidUpdateTemperature(manager ICMWaterSubmersionManager, measurement ICMWaterTemperature)/* debug [protocol_interface/required_method]: ManagerDidUpdateTemperature */
	ManagerDidUpdateEvent(manager ICMWaterSubmersionManager, event ICMWaterSubmersionEvent)/* debug [protocol_interface/required_method]: ManagerDidUpdateEvent */
	ManagerDidUpdateMeasurement(manager ICMWaterSubmersionManager, measurement ICMWaterSubmersionMeasurement)/* debug [protocol_interface/required_method]: ManagerDidUpdateMeasurement */
	ManagerErrorOccurred(manager ICMWaterSubmersionManager, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: ManagerErrorOccurred */
}

// WaterSubmersionManagerDelegate is a delegate implementation builder for the PWaterSubmersionManagerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type WaterSubmersionManagerDelegate struct {
	_ManagerDidUpdateTemperature func(manager ICMWaterSubmersionManager, measurement ICMWaterTemperature)
	_ManagerDidUpdateEvent func(manager ICMWaterSubmersionManager, event ICMWaterSubmersionEvent)
	_ManagerDidUpdateMeasurement func(manager ICMWaterSubmersionManager, measurement ICMWaterSubmersionMeasurement)
	_ManagerErrorOccurred func(manager ICMWaterSubmersionManager, error_ objc.IObject /* cross-framework: Error */)
}

// SetManagerDidUpdateTemperature sets the handler for the ManagerDidUpdateTemperature delegate method.
//
// Provides the delegate with updated water temperature data.
func (d *WaterSubmersionManagerDelegate) SetManagerDidUpdateTemperature(f func(manager ICMWaterSubmersionManager, measurement ICMWaterTemperature)) {
	d._ManagerDidUpdateTemperature = f
}

// SetManagerDidUpdateEvent sets the handler for the ManagerDidUpdateEvent delegate method.
//
// Tells the delegate when a water submersion event occurs.
func (d *WaterSubmersionManagerDelegate) SetManagerDidUpdateEvent(f func(manager ICMWaterSubmersionManager, event ICMWaterSubmersionEvent)) {
	d._ManagerDidUpdateEvent = f
}

// SetManagerDidUpdateMeasurement sets the handler for the ManagerDidUpdateMeasurement delegate method.
//
// Provides the delegate with a new set of pressure and depth measurements.
func (d *WaterSubmersionManagerDelegate) SetManagerDidUpdateMeasurement(f func(manager ICMWaterSubmersionManager, measurement ICMWaterSubmersionMeasurement)) {
	d._ManagerDidUpdateMeasurement = f
}

// SetManagerErrorOccurred sets the handler for the ManagerErrorOccurred delegate method.
//
// Tells the delegate when an error occurs.
func (d *WaterSubmersionManagerDelegate) SetManagerErrorOccurred(f func(manager ICMWaterSubmersionManager, error_ objc.IObject /* cross-framework: Error */)) {
	d._ManagerErrorOccurred = f
}

// ManagerDidUpdateTemperature implements the PWaterSubmersionManagerDelegate interface.
func (d *WaterSubmersionManagerDelegate) ManagerDidUpdateTemperature(manager ICMWaterSubmersionManager, measurement ICMWaterTemperature) {
	if d._ManagerDidUpdateTemperature != nil {
		d._ManagerDidUpdateTemperature(manager, measurement)
	}
}

// HasManagerDidUpdateTemperature returns true if a handler for ManagerDidUpdateTemperature has been set.
func (d *WaterSubmersionManagerDelegate) HasManagerDidUpdateTemperature() bool {
	return d._ManagerDidUpdateTemperature != nil
}

// ManagerDidUpdateEvent implements the PWaterSubmersionManagerDelegate interface.
func (d *WaterSubmersionManagerDelegate) ManagerDidUpdateEvent(manager ICMWaterSubmersionManager, event ICMWaterSubmersionEvent) {
	if d._ManagerDidUpdateEvent != nil {
		d._ManagerDidUpdateEvent(manager, event)
	}
}

// HasManagerDidUpdateEvent returns true if a handler for ManagerDidUpdateEvent has been set.
func (d *WaterSubmersionManagerDelegate) HasManagerDidUpdateEvent() bool {
	return d._ManagerDidUpdateEvent != nil
}

// ManagerDidUpdateMeasurement implements the PWaterSubmersionManagerDelegate interface.
func (d *WaterSubmersionManagerDelegate) ManagerDidUpdateMeasurement(manager ICMWaterSubmersionManager, measurement ICMWaterSubmersionMeasurement) {
	if d._ManagerDidUpdateMeasurement != nil {
		d._ManagerDidUpdateMeasurement(manager, measurement)
	}
}

// HasManagerDidUpdateMeasurement returns true if a handler for ManagerDidUpdateMeasurement has been set.
func (d *WaterSubmersionManagerDelegate) HasManagerDidUpdateMeasurement() bool {
	return d._ManagerDidUpdateMeasurement != nil
}

// ManagerErrorOccurred implements the PWaterSubmersionManagerDelegate interface.
func (d *WaterSubmersionManagerDelegate) ManagerErrorOccurred(manager ICMWaterSubmersionManager, error_ objc.IObject /* cross-framework: Error */) {
	if d._ManagerErrorOccurred != nil {
		d._ManagerErrorOccurred(manager, error_)
	}
}

// HasManagerErrorOccurred returns true if a handler for ManagerErrorOccurred has been set.
func (d *WaterSubmersionManagerDelegate) HasManagerErrorOccurred() bool {
	return d._ManagerErrorOccurred != nil
}
