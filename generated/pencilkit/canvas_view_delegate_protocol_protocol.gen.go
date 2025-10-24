// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PCanvasViewDelegate is the PKCanvasViewDelegate protocol interface.
//
// Methods for monitoring drawing related changes in a canvas view.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.pencilkit/documentation/PencilKit/PKCanvasViewDelegate
type PCanvasViewDelegate interface {
	// Optional methods
	CanvasViewDidBeginUsingTool(canvasView IPKCanvasView)
	HasCanvasViewDidBeginUsingTool() bool
	CanvasViewDidEndUsingTool(canvasView IPKCanvasView)
	HasCanvasViewDidEndUsingTool() bool
	CanvasViewDidFinishRendering(canvasView IPKCanvasView)
	HasCanvasViewDidFinishRendering() bool
	CanvasViewDrawingDidChange(canvasView IPKCanvasView)
	HasCanvasViewDrawingDidChange() bool
}

// CanvasViewDelegate is a delegate implementation builder for the PCanvasViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CanvasViewDelegate struct {
	_CanvasViewDidBeginUsingTool func(canvasView IPKCanvasView)
	_CanvasViewDidEndUsingTool func(canvasView IPKCanvasView)
	_CanvasViewDidFinishRendering func(canvasView IPKCanvasView)
	_CanvasViewDrawingDidChange func(canvasView IPKCanvasView)
}

// SetCanvasViewDidBeginUsingTool sets the handler for the CanvasViewDidBeginUsingTool delegate method.
//
// Tells the delegate that the user started a new drawing sequence with the currently selected tool.
func (d *CanvasViewDelegate) SetCanvasViewDidBeginUsingTool(f func(canvasView IPKCanvasView)) {
	d._CanvasViewDidBeginUsingTool = f
}

// SetCanvasViewDidEndUsingTool sets the handler for the CanvasViewDidEndUsingTool delegate method.
//
// Tells the delegate that the user ended a drawing sequence with the tool they were using.
func (d *CanvasViewDelegate) SetCanvasViewDidEndUsingTool(f func(canvasView IPKCanvasView)) {
	d._CanvasViewDidEndUsingTool = f
}

// SetCanvasViewDidFinishRendering sets the handler for the CanvasViewDidFinishRendering delegate method.
//
// Tells the delegate that the previously drawn content is ready to display.
func (d *CanvasViewDelegate) SetCanvasViewDidFinishRendering(f func(canvasView IPKCanvasView)) {
	d._CanvasViewDidFinishRendering = f
}

// SetCanvasViewDrawingDidChange sets the handler for the CanvasViewDrawingDidChange delegate method.
//
// Tells the delegate that the contents of the current drawing changed.
func (d *CanvasViewDelegate) SetCanvasViewDrawingDidChange(f func(canvasView IPKCanvasView)) {
	d._CanvasViewDrawingDidChange = f
}

// CanvasViewDidBeginUsingTool implements the PCanvasViewDelegate interface.
func (d *CanvasViewDelegate) CanvasViewDidBeginUsingTool(canvasView IPKCanvasView) {
	if d._CanvasViewDidBeginUsingTool != nil {
		d._CanvasViewDidBeginUsingTool(canvasView)
	}
}

// HasCanvasViewDidBeginUsingTool returns true if a handler for CanvasViewDidBeginUsingTool has been set.
func (d *CanvasViewDelegate) HasCanvasViewDidBeginUsingTool() bool {
	return d._CanvasViewDidBeginUsingTool != nil
}

// CanvasViewDidEndUsingTool implements the PCanvasViewDelegate interface.
func (d *CanvasViewDelegate) CanvasViewDidEndUsingTool(canvasView IPKCanvasView) {
	if d._CanvasViewDidEndUsingTool != nil {
		d._CanvasViewDidEndUsingTool(canvasView)
	}
}

// HasCanvasViewDidEndUsingTool returns true if a handler for CanvasViewDidEndUsingTool has been set.
func (d *CanvasViewDelegate) HasCanvasViewDidEndUsingTool() bool {
	return d._CanvasViewDidEndUsingTool != nil
}

// CanvasViewDidFinishRendering implements the PCanvasViewDelegate interface.
func (d *CanvasViewDelegate) CanvasViewDidFinishRendering(canvasView IPKCanvasView) {
	if d._CanvasViewDidFinishRendering != nil {
		d._CanvasViewDidFinishRendering(canvasView)
	}
}

// HasCanvasViewDidFinishRendering returns true if a handler for CanvasViewDidFinishRendering has been set.
func (d *CanvasViewDelegate) HasCanvasViewDidFinishRendering() bool {
	return d._CanvasViewDidFinishRendering != nil
}

// CanvasViewDrawingDidChange implements the PCanvasViewDelegate interface.
func (d *CanvasViewDelegate) CanvasViewDrawingDidChange(canvasView IPKCanvasView) {
	if d._CanvasViewDrawingDidChange != nil {
		d._CanvasViewDrawingDidChange(canvasView)
	}
}

// HasCanvasViewDrawingDidChange returns true if a handler for CanvasViewDrawingDidChange has been set.
func (d *CanvasViewDelegate) HasCanvasViewDrawingDidChange() bool {
	return d._CanvasViewDrawingDidChange != nil
}
