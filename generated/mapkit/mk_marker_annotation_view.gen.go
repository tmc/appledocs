// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKMarkerAnnotationView] class.
var (
	MKMarkerAnnotationViewClass     _MKMarkerAnnotationViewClass
	MKMarkerAnnotationViewClassOnce sync.Once
)

func getMKMarkerAnnotationViewClass() _MKMarkerAnnotationViewClass {
	MKMarkerAnnotationViewClassOnce.Do(func() {
		MKMarkerAnnotationViewClass = _MKMarkerAnnotationViewClass{objc.GetClass("MKMarkerAnnotationView")}
	})
	return MKMarkerAnnotationViewClass
}

type _MKMarkerAnnotationViewClass struct {
	class objc.Class
}

// An interface definition for the [MKMarkerAnnotationView] class.
type IMKMarkerAnnotationView interface {
	IMKAnnotationView
}

// An annotation view that displays a balloon-shaped marker at the designated location.
//
// Return an instance of this class from the method of your map view delegate when you want to display the same types of markers used in the Maps app. The default for an instance of this class is .
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView
type MKMarkerAnnotationView struct {
	MKAnnotationView
}

// MKMarkerAnnotationViewFrom constructs a [MKMarkerAnnotationView] from an unsafe.Pointer.
//
// An annotation view that displays a balloon-shaped marker at the designated location.
func MKMarkerAnnotationViewFrom(ptr unsafe.Pointer) MKMarkerAnnotationView {
	return MKMarkerAnnotationView{
		MKAnnotationView: MKAnnotationViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMarkerAnnotationViewClass) Alloc() MKMarkerAnnotationView {
	rv := objc.Send[MKMarkerAnnotationView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMarkerAnnotationViewClass) New() MKMarkerAnnotationView {
	rv := objc.Send[MKMarkerAnnotationView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMarkerAnnotationView) Init() MKMarkerAnnotationView {
	rv := objc.Send[MKMarkerAnnotationView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMarkerAnnotationView) Autorelease() MKMarkerAnnotationView {
	rv := objc.Send[MKMarkerAnnotationView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMarkerAnnotationView creates a new MKMarkerAnnotationView instance.
func NewMKMarkerAnnotationView() MKMarkerAnnotationView {
	return getMKMarkerAnnotationViewClass().New()
}


// An image to display in the marker balloon.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/glyphImage
func (m_ MKMarkerAnnotationView) GlyphImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("glyphImage"))
	return rv
}


// SetGlyphImage sets the value of the glyphImage property.
// An image to display in the marker balloon.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/glyphImage
func (m_ MKMarkerAnnotationView) SetGlyphImage(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGlyphImage:"), value)
}
// The color to apply to the glyph text or image.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/glyphTintColor
func (m_ MKMarkerAnnotationView) GlyphTintColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("glyphTintColor"))
	return rv
}


// SetGlyphTintColor sets the value of the glyphTintColor property.
// The color to apply to the glyph text or image.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/glyphTintColor
func (m_ MKMarkerAnnotationView) SetGlyphTintColor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGlyphTintColor:"), value)
}
// The background color of the marker balloon.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/markerTintColor
func (m_ MKMarkerAnnotationView) MarkerTintColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("markerTintColor"))
	return rv
}


// SetMarkerTintColor sets the value of the markerTintColor property.
// The background color of the marker balloon.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/markerTintColor
func (m_ MKMarkerAnnotationView) SetMarkerTintColor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMarkerTintColor:"), value)
}


