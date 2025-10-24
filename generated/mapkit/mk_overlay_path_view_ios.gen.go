//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for MKOverlayPathView


// iOS-only properties

// The fill color to use for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathView/fillColor
func (m_ MKOverlayPathView) FillColor() appkit.Color {
	rv := objc.Send[appkit.Color](m_.ID, objc.Sel("fillColor"))
	return rv
}
func (m_ MKOverlayPathView) SetFillColor(value appkit.Color) {
	m_.ID.Send(objc.RegisterName("setFillColor:"), value)
}

// The line cap style to apply to the open ends of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathView/lineCap
func (m_ MKOverlayPathView) LineCap() LineCap /* not a class type */ {
	rv := objc.Send[LineCap](m_.ID, objc.Sel("lineCap"))
	return rv
}
func (m_ MKOverlayPathView) SetLineCap(value LineCap /* not a class type */) {
	m_.ID.Send(objc.RegisterName("setLineCap:"), value)
}

// An array of numbers indicating the dash pattern for paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathView/lineDashPattern
func (m_ MKOverlayPathView) LineDashPattern() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("lineDashPattern"))
	return rv
}
func (m_ MKOverlayPathView) SetLineDashPattern(value objc.IObject /* cross-framework: NSArray */) {
	m_.ID.Send(objc.RegisterName("setLineDashPattern:"), value)
}

// The offset (in points) at which to start drawing the dash pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathView/lineDashPhase
func (m_ MKOverlayPathView) LineDashPhase() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineDashPhase"))
	return rv
}
func (m_ MKOverlayPathView) SetLineDashPhase(value float64) {
	m_.ID.Send(objc.RegisterName("setLineDashPhase:"), value)
}

// The line join style to apply to corners of the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathView/lineJoin
func (m_ MKOverlayPathView) LineJoin() LineJoin /* not a class type */ {
	rv := objc.Send[LineJoin](m_.ID, objc.Sel("lineJoin"))
	return rv
}
func (m_ MKOverlayPathView) SetLineJoin(value LineJoin /* not a class type */) {
	m_.ID.Send(objc.RegisterName("setLineJoin:"), value)
}

// The stroke width to use for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathView/lineWidth
func (m_ MKOverlayPathView) LineWidth() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("lineWidth"))
	return rv
}
func (m_ MKOverlayPathView) SetLineWidth(value float64) {
	m_.ID.Send(objc.RegisterName("setLineWidth:"), value)
}

// The limiting value that helps avoid spikes at junctions between connected line segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathView/miterLimit
func (m_ MKOverlayPathView) MiterLimit() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("miterLimit"))
	return rv
}
func (m_ MKOverlayPathView) SetMiterLimit(value float64) {
	m_.ID.Send(objc.RegisterName("setMiterLimit:"), value)
}

// The current path to use when drawing the overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathView/path
func (m_ MKOverlayPathView) Path() PathRef /* not a class type */ {
	rv := objc.Send[PathRef](m_.ID, objc.Sel("path"))
	return rv
}
func (m_ MKOverlayPathView) SetPath(value PathRef /* not a class type */) {
	m_.ID.Send(objc.RegisterName("setPath:"), value)
}

// The stroke color to use for the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathView/strokeColor
func (m_ MKOverlayPathView) StrokeColor() appkit.Color {
	rv := objc.Send[appkit.Color](m_.ID, objc.Sel("strokeColor"))
	return rv
}
func (m_ MKOverlayPathView) SetStrokeColor(value appkit.Color) {
	m_.ID.Send(objc.RegisterName("setStrokeColor:"), value)
}





