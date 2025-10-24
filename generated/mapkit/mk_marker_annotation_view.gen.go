// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MKMarkerAnnotationView */


/* debug [class_header]: Header for MKMarkerAnnotationView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMarkerAnnotationView */
// An interface definition for the [MKMarkerAnnotationView] class.
type IMKMarkerAnnotationView interface {
	IMKAnnotationView
	
/* debug [class_interface_properties]: Properties for MKMarkerAnnotationView */
	// properties:
	AnimatesWhenAdded() bool
	SetAnimatesWhenAdded(value bool)
	GlyphImage() appkit.Image
	SetGlyphImage(value appkit.Image)
	GlyphText() objc.IObject /* cross-framework: NSString */
	SetGlyphText(value objc.IObject /* cross-framework: NSString */)
	GlyphTintColor() appkit.Color
	SetGlyphTintColor(value appkit.Color)
	MarkerTintColor() appkit.Color
	SetMarkerTintColor(value appkit.Color)
	SelectedGlyphImage() appkit.Image
	SetSelectedGlyphImage(value appkit.Image)
	SubtitleVisibility() MKFeatureVisibility
	SetSubtitleVisibility(value MKFeatureVisibility)
	TitleVisibility() MKFeatureVisibility
	SetTitleVisibility(value MKFeatureVisibility)
	DisplayPriority() MKFeatureDisplayPriority /* typedef */
	SetDisplayPriority(value MKFeatureDisplayPriority /* typedef */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMarkerAnnotationView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMarkerAnnotationView */
// Alloc allocates a new instance without initialization.
func (mc _MKMarkerAnnotationViewClass) Alloc() MKMarkerAnnotationView {
	rv := objc.Send[MKMarkerAnnotationView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMarkerAnnotationView */
// An annotation view that displays a balloon-shaped marker at the designated location.
//
// Return an instance of this class from the method of your map view delegate when you want to display the same types of markers used in the Maps app. The default for an instance of this class is .


// An annotation view that displays a balloon-shaped marker at the designated location.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMarkerAnnotationView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMarkerAnnotationView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMarkerAnnotationView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMarkerAnnotationView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMarkerAnnotationView */

// A Boolean that indicates whether the marker animates into position onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/animatesWhenAdded
func (m_ MKMarkerAnnotationView) AnimatesWhenAdded() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("animatesWhenAdded"))
	return rv
}/* debug [instance_properties/getter]: animatesWhenAdded */


// A Boolean that indicates whether the marker animates into position onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/animatesWhenAdded
func (m_ MKMarkerAnnotationView) SetAnimatesWhenAdded(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAnimatesWhenAdded:"), value)
}/* debug [instance_properties/setter]: animatesWhenAdded */


// An image to display in the marker balloon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/glyphImage
func (m_ MKMarkerAnnotationView) GlyphImage() appkit.Image {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("glyphImage"))
	return rv
}/* debug [instance_properties/getter]: glyphImage */


// An image to display in the marker balloon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/glyphImage
func (m_ MKMarkerAnnotationView) SetGlyphImage(value appkit.Image) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGlyphImage:"), value)
}/* debug [instance_properties/setter]: glyphImage */


// The text to display in the marker balloon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/glyphText
func (m_ MKMarkerAnnotationView) GlyphText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("glyphText"))
	return rv
}/* debug [instance_properties/getter]: glyphText */


// The text to display in the marker balloon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/glyphText
func (m_ MKMarkerAnnotationView) SetGlyphText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGlyphText:"), value)
}/* debug [instance_properties/setter]: glyphText */


// The color to apply to the glyph text or image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/glyphTintColor
func (m_ MKMarkerAnnotationView) GlyphTintColor() appkit.Color {
	rv := objc.Send[appkit.Color](m_.ID, objc.Sel("glyphTintColor"))
	return rv
}/* debug [instance_properties/getter]: glyphTintColor */


// The color to apply to the glyph text or image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/glyphTintColor
func (m_ MKMarkerAnnotationView) SetGlyphTintColor(value appkit.Color) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGlyphTintColor:"), value)
}/* debug [instance_properties/setter]: glyphTintColor */


// The background color of the marker balloon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/markerTintColor
func (m_ MKMarkerAnnotationView) MarkerTintColor() appkit.Color {
	rv := objc.Send[appkit.Color](m_.ID, objc.Sel("markerTintColor"))
	return rv
}/* debug [instance_properties/getter]: markerTintColor */


// The background color of the marker balloon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/markerTintColor
func (m_ MKMarkerAnnotationView) SetMarkerTintColor(value appkit.Color) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMarkerTintColor:"), value)
}/* debug [instance_properties/setter]: markerTintColor */


// An image to display when the user selects the marker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/selectedGlyphImage
func (m_ MKMarkerAnnotationView) SelectedGlyphImage() appkit.Image {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("selectedGlyphImage"))
	return rv
}/* debug [instance_properties/getter]: selectedGlyphImage */


// An image to display when the user selects the marker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/selectedGlyphImage
func (m_ MKMarkerAnnotationView) SetSelectedGlyphImage(value appkit.Image) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSelectedGlyphImage:"), value)
}/* debug [instance_properties/setter]: selectedGlyphImage */


// The visibility of the subtitle text rendered beneath the marker balloon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/subtitleVisibility
func (m_ MKMarkerAnnotationView) SubtitleVisibility() MKFeatureVisibility {
	rv := objc.Send[MKFeatureVisibility](m_.ID, objc.Sel("subtitleVisibility"))
	return rv
}/* debug [instance_properties/getter]: subtitleVisibility */


// The visibility of the subtitle text rendered beneath the marker balloon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/subtitleVisibility
func (m_ MKMarkerAnnotationView) SetSubtitleVisibility(value MKFeatureVisibility) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubtitleVisibility:"), value)
}/* debug [instance_properties/setter]: subtitleVisibility */


// The visibility of the title text rendered beneath the marker balloon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/titleVisibility
func (m_ MKMarkerAnnotationView) TitleVisibility() MKFeatureVisibility {
	rv := objc.Send[MKFeatureVisibility](m_.ID, objc.Sel("titleVisibility"))
	return rv
}/* debug [instance_properties/getter]: titleVisibility */


// The visibility of the title text rendered beneath the marker balloon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMarkerAnnotationView/titleVisibility
func (m_ MKMarkerAnnotationView) SetTitleVisibility(value MKFeatureVisibility) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTitleVisibility:"), value)
}/* debug [instance_properties/setter]: titleVisibility */


// The display priority of the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/displaypriority
func (m_ MKMarkerAnnotationView) DisplayPriority() MKFeatureDisplayPriority /* typedef */ {
	rv := objc.Send[float32](m_.ID, objc.Sel("displayPriority"))
	return rv
}/* debug [instance_properties/getter]: displayPriority */


// The display priority of the annotation view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/displaypriority
func (m_ MKMarkerAnnotationView) SetDisplayPriority(value MKFeatureDisplayPriority /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisplayPriority:"), value)
}/* debug [instance_properties/setter]: displayPriority */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMarkerAnnotationView */



