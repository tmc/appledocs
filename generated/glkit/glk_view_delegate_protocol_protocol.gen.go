// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PGLKViewDelegate is the GLKViewDelegate protocol interface.
//
// Drawing callback methods for use with a   object.
//
// Availability:
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - tvOS 9.0+
//
// See: doc://com.apple.glkit/documentation/GLKit/GLKViewDelegate
type PGLKViewDelegate interface {
	// Required methods
	GlkViewDrawInRect(view IGLKView, rect corefoundation.CGRect)/* debug [protocol_interface/required_method]: GlkViewDrawInRect */
}

// GLKViewDelegate is a delegate implementation builder for the PGLKViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type GLKViewDelegate struct {
	_GlkViewDrawInRect func(view IGLKView, rect corefoundation.CGRect)
}

// SetGlkViewDrawInRect sets the handler for the GlkViewDrawInRect delegate method.
//
// Draws the view’s contents.
func (d *GLKViewDelegate) SetGlkViewDrawInRect(f func(view IGLKView, rect corefoundation.CGRect)) {
	d._GlkViewDrawInRect = f
}

// GlkViewDrawInRect implements the PGLKViewDelegate interface.
func (d *GLKViewDelegate) GlkViewDrawInRect(view IGLKView, rect corefoundation.CGRect) {
	if d._GlkViewDrawInRect != nil {
		d._GlkViewDrawInRect(view, rect)
	}
}

// HasGlkViewDrawInRect returns true if a handler for GlkViewDrawInRect has been set.
func (d *GLKViewDelegate) HasGlkViewDrawInRect() bool {
	return d._GlkViewDrawInRect != nil
}
