// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PCaptureTimecodeGeneratorDelegate is the AVCaptureTimecodeGeneratorDelegate protocol interface.
//
// A protocol for receiving real-time timecode updates and error notifications from a timecode generator.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureTimecodeGeneratorDelegate
type PCaptureTimecodeGeneratorDelegate interface {
	// Required methods
	TimecodeGeneratorDidReceiveUpdateFromSource(generator IAVCaptureTimecodeGenerator, timecode objc.IObject /* cross-framework: AVCaptureTimecode */, source IAVCaptureTimecodeSource)
	TimecodeGeneratorDidUpdateAvailableSources(generator IAVCaptureTimecodeGenerator, availableSources []CaptureTimecodeSource)
	TimecodeGeneratorTransitionedToSynchronizationStatusForSource(generator IAVCaptureTimecodeGenerator, synchronizationStatus CaptureTimecodeGeneratorSynchronizationStatus, source IAVCaptureTimecodeSource)
}

// CaptureTimecodeGeneratorDelegate is a delegate implementation builder for the PCaptureTimecodeGeneratorDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CaptureTimecodeGeneratorDelegate struct {
	_TimecodeGeneratorDidReceiveUpdateFromSource func(generator IAVCaptureTimecodeGenerator, timecode objc.IObject /* cross-framework: AVCaptureTimecode */, source IAVCaptureTimecodeSource)
	_TimecodeGeneratorDidUpdateAvailableSources func(generator IAVCaptureTimecodeGenerator, availableSources []CaptureTimecodeSource)
	_TimecodeGeneratorTransitionedToSynchronizationStatusForSource func(generator IAVCaptureTimecodeGenerator, synchronizationStatus CaptureTimecodeGeneratorSynchronizationStatus, source IAVCaptureTimecodeSource)
}

// SetTimecodeGeneratorDidReceiveUpdateFromSource sets the handler for the TimecodeGeneratorDidReceiveUpdateFromSource delegate method.
//
// Notifies the delegate when new, unaligned timecodes are parsed from the specified source.
func (d *CaptureTimecodeGeneratorDelegate) SetTimecodeGeneratorDidReceiveUpdateFromSource(f func(generator IAVCaptureTimecodeGenerator, timecode objc.IObject /* cross-framework: AVCaptureTimecode */, source IAVCaptureTimecodeSource)) {
	d._TimecodeGeneratorDidReceiveUpdateFromSource = f
}

// SetTimecodeGeneratorDidUpdateAvailableSources sets the handler for the TimecodeGeneratorDidUpdateAvailableSources delegate method.
//
// Notifies the delegate when the list of available timecode synchronization sources is updated.
func (d *CaptureTimecodeGeneratorDelegate) SetTimecodeGeneratorDidUpdateAvailableSources(f func(generator IAVCaptureTimecodeGenerator, availableSources []CaptureTimecodeSource)) {
	d._TimecodeGeneratorDidUpdateAvailableSources = f
}

// SetTimecodeGeneratorTransitionedToSynchronizationStatusForSource sets the handler for the TimecodeGeneratorTransitionedToSynchronizationStatusForSource delegate method.
//
// Notifies the delegate when the synchronization status of a timecode source changes.
func (d *CaptureTimecodeGeneratorDelegate) SetTimecodeGeneratorTransitionedToSynchronizationStatusForSource(f func(generator IAVCaptureTimecodeGenerator, synchronizationStatus CaptureTimecodeGeneratorSynchronizationStatus, source IAVCaptureTimecodeSource)) {
	d._TimecodeGeneratorTransitionedToSynchronizationStatusForSource = f
}

// TimecodeGeneratorDidReceiveUpdateFromSource implements the PCaptureTimecodeGeneratorDelegate interface.
func (d *CaptureTimecodeGeneratorDelegate) TimecodeGeneratorDidReceiveUpdateFromSource(generator IAVCaptureTimecodeGenerator, timecode objc.IObject /* cross-framework: AVCaptureTimecode */, source IAVCaptureTimecodeSource) {
	if d._TimecodeGeneratorDidReceiveUpdateFromSource != nil {
		d._TimecodeGeneratorDidReceiveUpdateFromSource(generator, timecode, source)
	}
}

// HasTimecodeGeneratorDidReceiveUpdateFromSource returns true if a handler for TimecodeGeneratorDidReceiveUpdateFromSource has been set.
func (d *CaptureTimecodeGeneratorDelegate) HasTimecodeGeneratorDidReceiveUpdateFromSource() bool {
	return d._TimecodeGeneratorDidReceiveUpdateFromSource != nil
}

// TimecodeGeneratorDidUpdateAvailableSources implements the PCaptureTimecodeGeneratorDelegate interface.
func (d *CaptureTimecodeGeneratorDelegate) TimecodeGeneratorDidUpdateAvailableSources(generator IAVCaptureTimecodeGenerator, availableSources []CaptureTimecodeSource) {
	if d._TimecodeGeneratorDidUpdateAvailableSources != nil {
		d._TimecodeGeneratorDidUpdateAvailableSources(generator, availableSources)
	}
}

// HasTimecodeGeneratorDidUpdateAvailableSources returns true if a handler for TimecodeGeneratorDidUpdateAvailableSources has been set.
func (d *CaptureTimecodeGeneratorDelegate) HasTimecodeGeneratorDidUpdateAvailableSources() bool {
	return d._TimecodeGeneratorDidUpdateAvailableSources != nil
}

// TimecodeGeneratorTransitionedToSynchronizationStatusForSource implements the PCaptureTimecodeGeneratorDelegate interface.
func (d *CaptureTimecodeGeneratorDelegate) TimecodeGeneratorTransitionedToSynchronizationStatusForSource(generator IAVCaptureTimecodeGenerator, synchronizationStatus CaptureTimecodeGeneratorSynchronizationStatus, source IAVCaptureTimecodeSource) {
	if d._TimecodeGeneratorTransitionedToSynchronizationStatusForSource != nil {
		d._TimecodeGeneratorTransitionedToSynchronizationStatusForSource(generator, synchronizationStatus, source)
	}
}

// HasTimecodeGeneratorTransitionedToSynchronizationStatusForSource returns true if a handler for TimecodeGeneratorTransitionedToSynchronizationStatusForSource has been set.
func (d *CaptureTimecodeGeneratorDelegate) HasTimecodeGeneratorTransitionedToSynchronizationStatusForSource() bool {
	return d._TimecodeGeneratorTransitionedToSynchronizationStatusForSource != nil
}
