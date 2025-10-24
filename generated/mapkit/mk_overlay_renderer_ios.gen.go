//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKOverlayRenderer


// iOS-only properties

// The blend mode to apply to the overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayRenderer/blendMode
func (m_ MKOverlayRenderer) BlendMode() BlendMode /* not a class type */ {
	rv := objc.Send[BlendMode](m_.ID, objc.Sel("blendMode"))
	return rv
}
func (m_ MKOverlayRenderer) SetBlendMode(value BlendMode /* not a class type */) {
	m_.ID.Send(objc.RegisterName("setBlendMode:"), value)
}




