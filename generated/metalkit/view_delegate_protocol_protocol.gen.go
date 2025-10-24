// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PViewDelegate is the MTKViewDelegate protocol interface.
//
// Methods for responding to a MetalKit view’s drawing and resizing events.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metalkit/documentation/MetalKit/MTKViewDelegate
type PViewDelegate interface {
	// Required methods
	DrawInMTKView(view IMTKView)/* debug [protocol_interface/required_method]: DrawInMTKView */
	MtkViewDrawableSizeWillChange(view IMTKView, size corefoundation.CGSize)/* debug [protocol_interface/required_method]: MtkViewDrawableSizeWillChange */
}

// ViewDelegate is a delegate implementation builder for the PViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ViewDelegate struct {
	_DrawInMTKView func(view IMTKView)
	_MtkViewDrawableSizeWillChange func(view IMTKView, size corefoundation.CGSize)
}

// SetDrawInMTKView sets the handler for the DrawInMTKView delegate method.
//
// Draws the view’s contents.
func (d *ViewDelegate) SetDrawInMTKView(f func(view IMTKView)) {
	d._DrawInMTKView = f
}

// SetMtkViewDrawableSizeWillChange sets the handler for the MtkViewDrawableSizeWillChange delegate method.
//
// Updates the view’s contents upon receiving a change in layout, resolution, or size.
func (d *ViewDelegate) SetMtkViewDrawableSizeWillChange(f func(view IMTKView, size corefoundation.CGSize)) {
	d._MtkViewDrawableSizeWillChange = f
}

// DrawInMTKView implements the PViewDelegate interface.
func (d *ViewDelegate) DrawInMTKView(view IMTKView) {
	if d._DrawInMTKView != nil {
		d._DrawInMTKView(view)
	}
}

// HasDrawInMTKView returns true if a handler for DrawInMTKView has been set.
func (d *ViewDelegate) HasDrawInMTKView() bool {
	return d._DrawInMTKView != nil
}

// MtkViewDrawableSizeWillChange implements the PViewDelegate interface.
func (d *ViewDelegate) MtkViewDrawableSizeWillChange(view IMTKView, size corefoundation.CGSize) {
	if d._MtkViewDrawableSizeWillChange != nil {
		d._MtkViewDrawableSizeWillChange(view, size)
	}
}

// HasMtkViewDrawableSizeWillChange returns true if a handler for MtkViewDrawableSizeWillChange has been set.
func (d *ViewDelegate) HasMtkViewDrawableSizeWillChange() bool {
	return d._MtkViewDrawableSizeWillChange != nil
}
