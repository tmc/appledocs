// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKGeoJSONFeature */


/* debug [class_header]: Header for MKGeoJSONFeature */
// The class instance for the [MKGeoJSONFeature] class.
var (
	MKGeoJSONFeatureClass     _MKGeoJSONFeatureClass
	MKGeoJSONFeatureClassOnce sync.Once
)

func getMKGeoJSONFeatureClass() _MKGeoJSONFeatureClass {
	MKGeoJSONFeatureClassOnce.Do(func() {
		MKGeoJSONFeatureClass = _MKGeoJSONFeatureClass{objc.GetClass("MKGeoJSONFeature")}
	})
	return MKGeoJSONFeatureClass
}

type _MKGeoJSONFeatureClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKGeoJSONFeature */
// An interface definition for the [MKGeoJSONFeature] class.
type IMKGeoJSONFeature interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKGeoJSONFeature */
	// properties:
	Geometry() []MKShape
	Identifier() objc.IObject /* cross-framework: NSString */
	Properties() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKGeoJSONFeature */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKGeoJSONFeature */
// Alloc allocates a new instance without initialization.
func (mc _MKGeoJSONFeatureClass) Alloc() MKGeoJSONFeature {
	rv := objc.Send[MKGeoJSONFeature](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKGeoJSONFeatureClass) New() MKGeoJSONFeature {
	rv := objc.Send[MKGeoJSONFeature](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKGeoJSONFeature) Init() MKGeoJSONFeature {
	rv := objc.Send[MKGeoJSONFeature](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKGeoJSONFeature) Autorelease() MKGeoJSONFeature {
	rv := objc.Send[MKGeoJSONFeature](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKGeoJSONFeature creates a new MKGeoJSONFeature instance.
func NewMKGeoJSONFeature() MKGeoJSONFeature {
	return getMKGeoJSONFeatureClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKGeoJSONFeature */
// The decoded representation of a GeoJSON feature.
//
// A feature is an object with associated geometry and optional properties in JSON that you define. MapKit exposes these optional properties, but treats them as opaque. is one of the classes that the GeoJSON decoder ( ) can return. See the GeoJSON standards specification for more information about objects.


// The decoded representation of a GeoJSON feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeoJSONFeature
type MKGeoJSONFeature struct {
	objectivec.Object
}

// MKGeoJSONFeatureFrom constructs a [MKGeoJSONFeature] from an unsafe.Pointer.
//
// The decoded representation of a GeoJSON feature.
func MKGeoJSONFeatureFrom(ptr unsafe.Pointer) MKGeoJSONFeature {
	return MKGeoJSONFeature{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKGeoJSONFeature *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKGeoJSONFeature */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKGeoJSONFeature */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKGeoJSONFeature */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKGeoJSONFeature */

// The shape or shapes associated with the GeoJSON feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeoJSONFeature/geometry
func (m_ MKGeoJSONFeature) Geometry() []MKShape {
	rv := objc.Send[[]MKShape](m_.ID, objc.Sel("geometry"))
	return rv
}/* debug [instance_properties/getter]: geometry */


// An optional identifier the class returns as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeoJSONFeature/identifier
func (m_ MKGeoJSONFeature) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// Optional serialized JSON data that corresponds to the properties key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeoJSONFeature/properties
func (m_ MKGeoJSONFeature) Properties() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("properties"))
	return rv
}/* debug [instance_properties/getter]: properties */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKGeoJSONFeature */



