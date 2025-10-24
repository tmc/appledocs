// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PGLKViewControllerDelegate is the GLKViewControllerDelegate protocol interface.
//
// Rendering loop callback methods for use with a   object.
//
// Availability:
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - tvOS 9.0+
//
// See: doc://com.apple.glkit/documentation/GLKit/GLKViewControllerDelegate
type PGLKViewControllerDelegate interface {
	// Required methods
	GlkViewControllerUpdate(controller IGLKViewController)/* debug [protocol_interface/required_method]: GlkViewControllerUpdate */
	// Optional methods
	GlkViewControllerWillPause(controller IGLKViewController, pause bool)
	HasGlkViewControllerWillPause() bool
}

// GLKViewControllerDelegate is a delegate implementation builder for the PGLKViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type GLKViewControllerDelegate struct {
	_GlkViewControllerWillPause func(controller IGLKViewController, pause bool)
	_GlkViewControllerUpdate func(controller IGLKViewController)
}

// SetGlkViewControllerWillPause sets the handler for the GlkViewControllerWillPause delegate method.
//
// Called before the rendering loop is paused or resumed.
func (d *GLKViewControllerDelegate) SetGlkViewControllerWillPause(f func(controller IGLKViewController, pause bool)) {
	d._GlkViewControllerWillPause = f
}

// SetGlkViewControllerUpdate sets the handler for the GlkViewControllerUpdate delegate method.
//
// Called before each frame is displayed.
func (d *GLKViewControllerDelegate) SetGlkViewControllerUpdate(f func(controller IGLKViewController)) {
	d._GlkViewControllerUpdate = f
}

// GlkViewControllerWillPause implements the PGLKViewControllerDelegate interface.
func (d *GLKViewControllerDelegate) GlkViewControllerWillPause(controller IGLKViewController, pause bool) {
	if d._GlkViewControllerWillPause != nil {
		d._GlkViewControllerWillPause(controller, pause)
	}
}

// HasGlkViewControllerWillPause returns true if a handler for GlkViewControllerWillPause has been set.
func (d *GLKViewControllerDelegate) HasGlkViewControllerWillPause() bool {
	return d._GlkViewControllerWillPause != nil
}

// GlkViewControllerUpdate implements the PGLKViewControllerDelegate interface.
func (d *GLKViewControllerDelegate) GlkViewControllerUpdate(controller IGLKViewController) {
	if d._GlkViewControllerUpdate != nil {
		d._GlkViewControllerUpdate(controller)
	}
}

// HasGlkViewControllerUpdate returns true if a handler for GlkViewControllerUpdate has been set.
func (d *GLKViewControllerDelegate) HasGlkViewControllerUpdate() bool {
	return d._GlkViewControllerUpdate != nil
}
