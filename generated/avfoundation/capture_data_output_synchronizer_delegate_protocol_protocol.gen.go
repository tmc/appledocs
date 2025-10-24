// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PCaptureDataOutputSynchronizerDelegate is the AVCaptureDataOutputSynchronizerDelegate protocol interface.
//
// Methods for receiving captured data from multiple capture outputs synchronized to the same timestamp.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - tvOS 17.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureDataOutputSynchronizerDelegate
type PCaptureDataOutputSynchronizerDelegate interface {
	// Required methods
	DataOutputSynchronizerDidOutputSynchronizedDataCollection(synchronizer IAVCaptureDataOutputSynchronizer, synchronizedDataCollection IAVCaptureSynchronizedDataCollection)
}

// CaptureDataOutputSynchronizerDelegate is a delegate implementation builder for the PCaptureDataOutputSynchronizerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CaptureDataOutputSynchronizerDelegate struct {
	_DataOutputSynchronizerDidOutputSynchronizedDataCollection func(synchronizer IAVCaptureDataOutputSynchronizer, synchronizedDataCollection IAVCaptureSynchronizedDataCollection)
}

// SetDataOutputSynchronizerDidOutputSynchronizedDataCollection sets the handler for the DataOutputSynchronizerDidOutputSynchronizedDataCollection delegate method.
//
// Provides a collection of synchronized capture data to the delegate.
func (d *CaptureDataOutputSynchronizerDelegate) SetDataOutputSynchronizerDidOutputSynchronizedDataCollection(f func(synchronizer IAVCaptureDataOutputSynchronizer, synchronizedDataCollection IAVCaptureSynchronizedDataCollection)) {
	d._DataOutputSynchronizerDidOutputSynchronizedDataCollection = f
}

// DataOutputSynchronizerDidOutputSynchronizedDataCollection implements the PCaptureDataOutputSynchronizerDelegate interface.
func (d *CaptureDataOutputSynchronizerDelegate) DataOutputSynchronizerDidOutputSynchronizedDataCollection(synchronizer IAVCaptureDataOutputSynchronizer, synchronizedDataCollection IAVCaptureSynchronizedDataCollection) {
	if d._DataOutputSynchronizerDidOutputSynchronizedDataCollection != nil {
		d._DataOutputSynchronizerDidOutputSynchronizedDataCollection(synchronizer, synchronizedDataCollection)
	}
}

// HasDataOutputSynchronizerDidOutputSynchronizedDataCollection returns true if a handler for DataOutputSynchronizerDidOutputSynchronizedDataCollection has been set.
func (d *CaptureDataOutputSynchronizerDelegate) HasDataOutputSynchronizerDidOutputSynchronizedDataCollection() bool {
	return d._DataOutputSynchronizerDidOutputSynchronizedDataCollection != nil
}
