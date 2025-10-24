// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PLayerDelegate is the CALayerDelegate protocol interface.
//
// Methods your app can implement to respond to layer-related events.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.quartzcore/documentation/QuartzCore/CALayerDelegate
type PLayerDelegate interface {
	// Optional methods
	ActionForLayerForKey(layer ILayer, event objc.IObject /* cross-framework: NSString */) unsafe.Pointer
	HasActionForLayerForKey() bool
	DisplayLayer(layer ILayer)
	HasDisplayLayer() bool
	DrawLayerInContext(layer ILayer, ctx ContextRef /* not a class type */)
	HasDrawLayerInContext() bool
	LayerWillDraw(layer ILayer)
	HasLayerWillDraw() bool
	LayoutSublayersOfLayer(layer ILayer)
	HasLayoutSublayersOfLayer() bool
}

// LayerDelegate is a delegate implementation builder for the PLayerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type LayerDelegate struct {
	_ActionForLayerForKey func(layer ILayer, event objc.IObject /* cross-framework: NSString */) unsafe.Pointer
	_DisplayLayer func(layer ILayer)
	_DrawLayerInContext func(layer ILayer, ctx ContextRef /* not a class type */)
	_LayerWillDraw func(layer ILayer)
	_LayoutSublayersOfLayer func(layer ILayer)
}

// SetActionForLayerForKey sets the handler for the ActionForLayerForKey delegate method.
//
// Returns the default action of the   method.
func (d *LayerDelegate) SetActionForLayerForKey(f func(layer ILayer, event objc.IObject /* cross-framework: NSString */) unsafe.Pointer) {
	d._ActionForLayerForKey = f
}

// SetDisplayLayer sets the handler for the DisplayLayer delegate method.
//
// Tells the delegate to implement the display process.
func (d *LayerDelegate) SetDisplayLayer(f func(layer ILayer)) {
	d._DisplayLayer = f
}

// SetDrawLayerInContext sets the handler for the DrawLayerInContext delegate method.
//
// Tells the delegate to implement the display process using the layer’s context.
func (d *LayerDelegate) SetDrawLayerInContext(f func(layer ILayer, ctx ContextRef /* not a class type */)) {
	d._DrawLayerInContext = f
}

// SetLayerWillDraw sets the handler for the LayerWillDraw delegate method.
//
// Notifies the delegate of an imminent draw.
func (d *LayerDelegate) SetLayerWillDraw(f func(layer ILayer)) {
	d._LayerWillDraw = f
}

// SetLayoutSublayersOfLayer sets the handler for the LayoutSublayersOfLayer delegate method.
//
// Tells the delegate a layer’s bounds have changed.
func (d *LayerDelegate) SetLayoutSublayersOfLayer(f func(layer ILayer)) {
	d._LayoutSublayersOfLayer = f
}

// ActionForLayerForKey implements the PLayerDelegate interface.
func (d *LayerDelegate) ActionForLayerForKey(layer ILayer, event objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	if d._ActionForLayerForKey != nil {
		return d._ActionForLayerForKey(layer, event)
	}
	var zero unsafe.Pointer
	return zero
}

// HasActionForLayerForKey returns true if a handler for ActionForLayerForKey has been set.
func (d *LayerDelegate) HasActionForLayerForKey() bool {
	return d._ActionForLayerForKey != nil
}

// DisplayLayer implements the PLayerDelegate interface.
func (d *LayerDelegate) DisplayLayer(layer ILayer) {
	if d._DisplayLayer != nil {
		d._DisplayLayer(layer)
	}
}

// HasDisplayLayer returns true if a handler for DisplayLayer has been set.
func (d *LayerDelegate) HasDisplayLayer() bool {
	return d._DisplayLayer != nil
}

// DrawLayerInContext implements the PLayerDelegate interface.
func (d *LayerDelegate) DrawLayerInContext(layer ILayer, ctx ContextRef /* not a class type */) {
	if d._DrawLayerInContext != nil {
		d._DrawLayerInContext(layer, ctx)
	}
}

// HasDrawLayerInContext returns true if a handler for DrawLayerInContext has been set.
func (d *LayerDelegate) HasDrawLayerInContext() bool {
	return d._DrawLayerInContext != nil
}

// LayerWillDraw implements the PLayerDelegate interface.
func (d *LayerDelegate) LayerWillDraw(layer ILayer) {
	if d._LayerWillDraw != nil {
		d._LayerWillDraw(layer)
	}
}

// HasLayerWillDraw returns true if a handler for LayerWillDraw has been set.
func (d *LayerDelegate) HasLayerWillDraw() bool {
	return d._LayerWillDraw != nil
}

// LayoutSublayersOfLayer implements the PLayerDelegate interface.
func (d *LayerDelegate) LayoutSublayersOfLayer(layer ILayer) {
	if d._LayoutSublayersOfLayer != nil {
		d._LayoutSublayersOfLayer(layer)
	}
}

// HasLayoutSublayersOfLayer returns true if a handler for LayoutSublayersOfLayer has been set.
func (d *LayerDelegate) HasLayoutSublayersOfLayer() bool {
	return d._LayoutSublayersOfLayer != nil
}
