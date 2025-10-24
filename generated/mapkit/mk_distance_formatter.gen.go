// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MKDistanceFormatter */


/* debug [class_header]: Header for MKDistanceFormatter */
// The class instance for the [MKDistanceFormatter] class.
var (
	MKDistanceFormatterClass     _MKDistanceFormatterClass
	MKDistanceFormatterClassOnce sync.Once
)

func getMKDistanceFormatterClass() _MKDistanceFormatterClass {
	MKDistanceFormatterClassOnce.Do(func() {
		MKDistanceFormatterClass = _MKDistanceFormatterClass{objc.GetClass("MKDistanceFormatter")}
	})
	return MKDistanceFormatterClass
}

type _MKDistanceFormatterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKDistanceFormatter */
// An interface definition for the [MKDistanceFormatter] class.
type IMKDistanceFormatter interface {
	IFormatter
	
/* debug [class_interface_properties]: Properties for MKDistanceFormatter */
	// properties:
	Locale() foundation.Locale
	SetLocale(value foundation.Locale)
	Units() MKDistanceFormatterUnits
	SetUnits(value MKDistanceFormatterUnits)
	UnitStyle() MKDistanceFormatterUnitStyle
	SetUnitStyle(value MKDistanceFormatterUnitStyle)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKDistanceFormatter */
	// methods:
	DistanceFromString(distance objc.IObject /* cross-framework: NSString */) LocationDistance /* not a class type */
	StringFromDistance(distance LocationDistance /* not a class type */) foundation.String
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKDistanceFormatter */
// Alloc allocates a new instance without initialization.
func (mc _MKDistanceFormatterClass) Alloc() MKDistanceFormatter {
	rv := objc.Send[MKDistanceFormatter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKDistanceFormatterClass) New() MKDistanceFormatter {
	rv := objc.Send[MKDistanceFormatter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKDistanceFormatter) Init() MKDistanceFormatter {
	rv := objc.Send[MKDistanceFormatter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKDistanceFormatter) Autorelease() MKDistanceFormatter {
	rv := objc.Send[MKDistanceFormatter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKDistanceFormatter creates a new MKDistanceFormatter instance.
func NewMKDistanceFormatter() MKDistanceFormatter {
	return getMKDistanceFormatterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKDistanceFormatter */
// A utility object that converts between a geographic distance and a string-based expression of that distance.
//
// Use a distance formatter to display distances to the user or to parse user-specified text to obtain a numerical value for a distance. When formatting strings containing distances, a distance formatter object takes into account the user’s locale and language settings. You can also specify a custom locale or custom units for any distances that you format.


// A utility object that converts between a geographic distance and a string-based expression of that distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter
type MKDistanceFormatter struct {
	Formatter
}

// MKDistanceFormatterFrom constructs a [MKDistanceFormatter] from an unsafe.Pointer.
//
// A utility object that converts between a geographic distance and a string-based expression of that distance.
func MKDistanceFormatterFrom(ptr unsafe.Pointer) MKDistanceFormatter {
	return MKDistanceFormatter{
		Formatter: FormatterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKDistanceFormatter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKDistanceFormatter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKDistanceFormatter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKDistanceFormatter */

// Returns the distance value parsed from the specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/distance(from:)
func (m_ MKDistanceFormatter) DistanceFromString(distance objc.IObject /* cross-framework: NSString */) LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](m_.ID, objc.Sel("distanceFromString:"), distance)
	return rv
}/* debug [instance_methods/method]: DistanceFromString */


// Creates a string representation of the specified distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/string(fromDistance:)
func (m_ MKDistanceFormatter) StringFromDistance(distance LocationDistance /* not a class type */) foundation.String {
	rv := objc.Send[foundation.String](m_.ID, objc.Sel("stringFromDistance:"), distance)
	return rv
}/* debug [instance_methods/method]: StringFromDistance */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKDistanceFormatter */

// The locale to use when formatting strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/locale
func (m_ MKDistanceFormatter) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](m_.ID, objc.Sel("locale"))
	return rv
}/* debug [instance_properties/getter]: locale */


// The locale to use when formatting strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/locale
func (m_ MKDistanceFormatter) SetLocale(value foundation.Locale) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocale:"), value)
}/* debug [instance_properties/setter]: locale */


// The measuring system — imperial or metric — to use for units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/units-swift.property
func (m_ MKDistanceFormatter) Units() MKDistanceFormatterUnits {
	rv := objc.Send[MKDistanceFormatterUnits](m_.ID, objc.Sel("units"))
	return rv
}/* debug [instance_properties/getter]: units */


// The measuring system — imperial or metric — to use for units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/units-swift.property
func (m_ MKDistanceFormatter) SetUnits(value MKDistanceFormatterUnits) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUnits:"), value)
}/* debug [instance_properties/setter]: units */


// The preferred style for units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/unitStyle
func (m_ MKDistanceFormatter) UnitStyle() MKDistanceFormatterUnitStyle {
	rv := objc.Send[MKDistanceFormatterUnitStyle](m_.ID, objc.Sel("unitStyle"))
	return rv
}/* debug [instance_properties/getter]: unitStyle */


// The preferred style for units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/unitStyle
func (m_ MKDistanceFormatter) SetUnitStyle(value MKDistanceFormatterUnitStyle) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUnitStyle:"), value)
}/* debug [instance_properties/setter]: unitStyle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKDistanceFormatter */



