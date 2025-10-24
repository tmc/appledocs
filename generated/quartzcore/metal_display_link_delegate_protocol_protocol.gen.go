// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PMetalDisplayLinkDelegate is the CAMetalDisplayLinkDelegate protocol interface.
//
// A protocol your app implements to respond to callbacks from Core Animation for a Metal display link.
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.quartzcore/documentation/QuartzCore/CAMetalDisplayLinkDelegate
type PMetalDisplayLinkDelegate interface {
	// Required methods
	MetalDisplayLinkNeedsUpdate(link IMetalDisplayLink, update IMetalDisplayLinkUpdate)/* debug [protocol_interface/required_method]: MetalDisplayLinkNeedsUpdate */
}

// MetalDisplayLinkDelegate is a delegate implementation builder for the PMetalDisplayLinkDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MetalDisplayLinkDelegate struct {
	_MetalDisplayLinkNeedsUpdate func(link IMetalDisplayLink, update IMetalDisplayLinkUpdate)
}

// SetMetalDisplayLinkNeedsUpdate sets the handler for the MetalDisplayLinkNeedsUpdate delegate method.
//
// A method the system calls to notify your app when it plans to update the display.
func (d *MetalDisplayLinkDelegate) SetMetalDisplayLinkNeedsUpdate(f func(link IMetalDisplayLink, update IMetalDisplayLinkUpdate)) {
	d._MetalDisplayLinkNeedsUpdate = f
}

// MetalDisplayLinkNeedsUpdate implements the PMetalDisplayLinkDelegate interface.
func (d *MetalDisplayLinkDelegate) MetalDisplayLinkNeedsUpdate(link IMetalDisplayLink, update IMetalDisplayLinkUpdate) {
	if d._MetalDisplayLinkNeedsUpdate != nil {
		d._MetalDisplayLinkNeedsUpdate(link, update)
	}
}

// HasMetalDisplayLinkNeedsUpdate returns true if a handler for MetalDisplayLinkNeedsUpdate has been set.
func (d *MetalDisplayLinkDelegate) HasMetalDisplayLinkNeedsUpdate() bool {
	return d._MetalDisplayLinkNeedsUpdate != nil
}
