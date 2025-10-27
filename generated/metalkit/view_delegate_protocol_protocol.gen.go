// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (

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
	DrawInMTKView(view IMTKView)
	MtkViewDrawableSizeWillChange(view IMTKView, size corefoundation.CGSize)
}
