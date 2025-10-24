// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PLayoutManager is the CALayoutManager protocol interface.
//
// Methods that allow an object to manage the layout of a layer and its sublayers.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS +
//
// See: doc://com.apple.quartzcore/documentation/QuartzCore/CALayoutManager
type PLayoutManager interface {
	// Optional methods
	InvalidateLayoutOfLayer(layer ILayer)
	HasInvalidateLayoutOfLayer() bool
	LayoutSublayersOfLayer(layer ILayer)
	HasLayoutSublayersOfLayer() bool
	PreferredSizeOfLayer(layer ILayer) corefoundation.CGSize
	HasPreferredSizeOfLayer() bool
}
