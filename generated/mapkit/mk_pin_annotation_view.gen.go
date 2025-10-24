// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class MKPinAnnotationView */


/* debug [class_header]: Header for MKPinAnnotationView */
// The class instance for the [MKPinAnnotationView] class.
var (
	MKPinAnnotationViewClass     _MKPinAnnotationViewClass
	MKPinAnnotationViewClassOnce sync.Once
)

func getMKPinAnnotationViewClass() _MKPinAnnotationViewClass {
	MKPinAnnotationViewClassOnce.Do(func() {
		MKPinAnnotationViewClass = _MKPinAnnotationViewClass{objc.GetClass("MKPinAnnotationView")}
	})
	return MKPinAnnotationViewClass
}

type _MKPinAnnotationViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKPinAnnotationView */
// An interface definition for the [MKPinAnnotationView] class.
type IMKPinAnnotationView interface {
	IMKAnnotationView
	
/* debug [class_interface_properties]: Properties for MKPinAnnotationView */
	// properties:
	AnimatesDrop() bool
	SetAnimatesDrop(value bool)
	PinColor() MKPinAnnotationColor
	SetPinColor(value MKPinAnnotationColor)
	PinTintColor() appkit.Color
	SetPinTintColor(value appkit.Color)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKPinAnnotationView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKPinAnnotationView */
// Alloc allocates a new instance without initialization.
func (mc _MKPinAnnotationViewClass) Alloc() MKPinAnnotationView {
	rv := objc.Send[MKPinAnnotationView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKPinAnnotationViewClass) New() MKPinAnnotationView {
	rv := objc.Send[MKPinAnnotationView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPinAnnotationView) Init() MKPinAnnotationView {
	rv := objc.Send[MKPinAnnotationView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPinAnnotationView) Autorelease() MKPinAnnotationView {
	rv := objc.Send[MKPinAnnotationView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPinAnnotationView creates a new MKPinAnnotationView instance.
func NewMKPinAnnotationView() MKPinAnnotationView {
	return getMKPinAnnotationViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKPinAnnotationView */
// An annotation view that displays a pin image on the map.
//
// Return instances of this class from the method of your map view delegate when you want to display a pin for one of your annotations. The pins displayed by this view are the same ones found in the Maps application. You can specify the type of pin you want to display and whether you want the pin to be animated into place.


// An annotation view that displays a pin image on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationView
type MKPinAnnotationView struct {
	MKAnnotationView
}

// MKPinAnnotationViewFrom constructs a [MKPinAnnotationView] from an unsafe.Pointer.
//
// An annotation view that displays a pin image on the map.
func MKPinAnnotationViewFrom(ptr unsafe.Pointer) MKPinAnnotationView {
	return MKPinAnnotationView{
		MKAnnotationView: MKAnnotationViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKPinAnnotationView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKPinAnnotationView */

// Returns the standard color for green pins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationView/greenPinColor()
func (mc _MKPinAnnotationViewClass) GreenPinColor() appkit.Color {
	rv := objc.Send[appkit.Color](objc.ID(mc.class), objc.Sel("greenPinColor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GreenPinColor) */


// Returns the standard color for purple pins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationView/purplePinColor()
func (mc _MKPinAnnotationViewClass) PurplePinColor() appkit.Color {
	rv := objc.Send[appkit.Color](objc.ID(mc.class), objc.Sel("purplePinColor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PurplePinColor) */


// Returns the standard color for red pins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationView/redPinColor()
func (mc _MKPinAnnotationViewClass) RedPinColor() appkit.Color {
	rv := objc.Send[appkit.Color](objc.ID(mc.class), objc.Sel("redPinColor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RedPinColor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKPinAnnotationView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKPinAnnotationView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKPinAnnotationView */

// A Boolean value indicating whether the annotation view is animated onto the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationView/animatesDrop
func (m_ MKPinAnnotationView) AnimatesDrop() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("animatesDrop"))
	return rv
}/* debug [instance_properties/getter]: animatesDrop */


// A Boolean value indicating whether the annotation view is animated onto the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationView/animatesDrop
func (m_ MKPinAnnotationView) SetAnimatesDrop(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAnimatesDrop:"), value)
}/* debug [instance_properties/setter]: animatesDrop */


// The color of the pin head.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationView/pinColor
func (m_ MKPinAnnotationView) PinColor() MKPinAnnotationColor {
	rv := objc.Send[MKPinAnnotationColor](m_.ID, objc.Sel("pinColor"))
	return rv
}/* debug [instance_properties/getter]: pinColor */


// The color of the pin head.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationView/pinColor
func (m_ MKPinAnnotationView) SetPinColor(value MKPinAnnotationColor) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPinColor:"), value)
}/* debug [instance_properties/setter]: pinColor */


// The color of the pin head.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationView/pinTintColor
func (m_ MKPinAnnotationView) PinTintColor() appkit.Color {
	rv := objc.Send[appkit.Color](m_.ID, objc.Sel("pinTintColor"))
	return rv
}/* debug [instance_properties/getter]: pinTintColor */


// The color of the pin head.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationView/pinTintColor
func (m_ MKPinAnnotationView) SetPinTintColor(value appkit.Color) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPinTintColor:"), value)
}/* debug [instance_properties/setter]: pinTintColor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKPinAnnotationView */



