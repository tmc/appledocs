// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PViewLayerContentScaleDelegate is the NSViewLayerContentScaleDelegate protocol interface.
//
// An optional layer delegate method for handling resolution changes.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSViewLayerContentScaleDelegate
type PViewLayerContentScaleDelegate interface {
	// Optional methods
	LayerShouldInheritContentsScaleFromWindow(layer objectivec.IObject, newScale float64, window IWindow) bool
	HasLayerShouldInheritContentsScaleFromWindow() bool
}

// ViewLayerContentScaleDelegate is a delegate implementation builder for the PViewLayerContentScaleDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ViewLayerContentScaleDelegate struct {
	_LayerShouldInheritContentsScaleFromWindow func(layer objectivec.IObject, newScale float64, window IWindow) bool
}

// SetLayerShouldInheritContentsScaleFromWindow sets the handler for the LayerShouldInheritContentsScaleFromWindow delegate method.
//
// Notifies you when a resolution changes occurs for the window that hosts the layer.
func (d *ViewLayerContentScaleDelegate) SetLayerShouldInheritContentsScaleFromWindow(f func(layer objectivec.IObject, newScale float64, window IWindow) bool) {
	d._LayerShouldInheritContentsScaleFromWindow = f
}

// LayerShouldInheritContentsScaleFromWindow implements the PViewLayerContentScaleDelegate interface.
func (d *ViewLayerContentScaleDelegate) LayerShouldInheritContentsScaleFromWindow(layer objectivec.IObject, newScale float64, window IWindow) bool {
	if d._LayerShouldInheritContentsScaleFromWindow != nil {
		return d._LayerShouldInheritContentsScaleFromWindow(layer, newScale, window)
	}
	var zero bool
	return zero
}

// HasLayerShouldInheritContentsScaleFromWindow returns true if a handler for LayerShouldInheritContentsScaleFromWindow has been set.
func (d *ViewLayerContentScaleDelegate) HasLayerShouldInheritContentsScaleFromWindow() bool {
	return d._LayerShouldInheritContentsScaleFromWindow != nil
}
