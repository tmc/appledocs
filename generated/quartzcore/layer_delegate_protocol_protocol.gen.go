// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

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
	ActionForLayerForKey(layer ILayer, event foundation.foundation.INSString) unsafe.Pointer
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
