// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
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
func (m_ MKMarkerAnnotationView) GlyphImage() appkit.Image {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("glyphImage"))
	return rv
}


// SetGlyphImage sets the value of the glyphImage property.
// An image to display in the marker balloon.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/glyphImage
func (m_ MKMarkerAnnotationView) SetGlyphImage(value appkit.IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGlyphImage:"), value)
}

// The color to apply to the glyph text or image.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/glyphTintColor
func (m_ MKMarkerAnnotationView) GlyphTintColor() appkit.Color {
	rv := objc.Send[appkit.Color](m_.ID, objc.Sel("glyphTintColor"))
	return rv
}


// SetGlyphTintColor sets the value of the glyphTintColor property.
// The color to apply to the glyph text or image.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/glyphTintColor
func (m_ MKMarkerAnnotationView) SetGlyphTintColor(value appkit.IColor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGlyphTintColor:"), value)
}

// The background color of the marker balloon.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/markerTintColor
func (m_ MKMarkerAnnotationView) MarkerTintColor() appkit.Color {
	rv := objc.Send[appkit.Color](m_.ID, objc.Sel("markerTintColor"))
	return rv
}


// SetMarkerTintColor sets the value of the markerTintColor property.
// The background color of the marker balloon.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/markerTintColor
func (m_ MKMarkerAnnotationView) SetMarkerTintColor(value appkit.IColor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMarkerTintColor:"), value)
}

// The display priority of the annotation view.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/displaypriority
func (m_ MKMarkerAnnotationView) DisplayPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("displayPriority"))
	return rv
}


// SetDisplayPriority sets the value of the displayPriority property.
// The display priority of the annotation view.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/displaypriority
func (m_ MKMarkerAnnotationView) SetDisplayPriority(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisplayPriority:"), value)
}

// A Boolean that indicates whether the marker animates into position onscreen.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmarkerannotationview/animateswhenadded
func (m_ MKMarkerAnnotationView) AnimatesWhenAdded() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("animatesWhenAdded"))
	return rv
}


// SetAnimatesWhenAdded sets the value of the animatesWhenAdded property.
// A Boolean that indicates whether the marker animates into position onscreen.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmarkerannotationview/animateswhenadded
func (m_ MKMarkerAnnotationView) SetAnimatesWhenAdded(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAnimatesWhenAdded:"), value)
}

// The text to display in the marker balloon.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmarkerannotationview/glyphtext
func (m_ MKMarkerAnnotationView) GlyphText() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("glyphText"))
	return rv
}


// SetGlyphText sets the value of the glyphText property.
// The text to display in the marker balloon.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmarkerannotationview/glyphtext
func (m_ MKMarkerAnnotationView) SetGlyphText(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGlyphText:"), value)
}

// An image to display when the user selects the marker.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmarkerannotationview/selectedglyphimage
func (m_ MKMarkerAnnotationView) SelectedGlyphImage() appkit.Image {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("selectedGlyphImage"))
	return rv
}


// SetSelectedGlyphImage sets the value of the selectedGlyphImage property.
// An image to display when the user selects the marker.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmarkerannotationview/selectedglyphimage
func (m_ MKMarkerAnnotationView) SetSelectedGlyphImage(value appkit.IImage) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectedGlyphImage:"), value)
}

// The visibility of the subtitle text rendered beneath the marker balloon.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmarkerannotationview/subtitlevisibility
func (m_ MKMarkerAnnotationView) SubtitleVisibility() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("subtitleVisibility"))
	return rv
}


// SetSubtitleVisibility sets the value of the subtitleVisibility property.
// The visibility of the subtitle text rendered beneath the marker balloon.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmarkerannotationview/subtitlevisibility
func (m_ MKMarkerAnnotationView) SetSubtitleVisibility(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitleVisibility:"), value)
}

// The visibility of the title text rendered beneath the marker balloon.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmarkerannotationview/titlevisibility
func (m_ MKMarkerAnnotationView) TitleVisibility() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("titleVisibility"))
	return rv
}


// SetTitleVisibility sets the value of the titleVisibility property.
// The visibility of the title text rendered beneath the marker balloon.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmarkerannotationview/titlevisibility
func (m_ MKMarkerAnnotationView) SetTitleVisibility(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitleVisibility:"), value)
}



