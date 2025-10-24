// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKGeoJSONDecoder */


/* debug [class_header]: Header for MKGeoJSONDecoder */
// The class instance for the [MKGeoJSONDecoder] class.
var (
	MKGeoJSONDecoderClass     _MKGeoJSONDecoderClass
	MKGeoJSONDecoderClassOnce sync.Once
)

func getMKGeoJSONDecoderClass() _MKGeoJSONDecoderClass {
	MKGeoJSONDecoderClassOnce.Do(func() {
		MKGeoJSONDecoderClass = _MKGeoJSONDecoderClass{objc.GetClass("MKGeoJSONDecoder")}
	})
	return MKGeoJSONDecoderClass
}

type _MKGeoJSONDecoderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKGeoJSONDecoder */
// An interface definition for the [MKGeoJSONDecoder] class.
type IMKGeoJSONDecoder interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKGeoJSONDecoder */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKGeoJSONDecoder */
	// methods:
	GeoJSONObjectsWithDataError(data objc.IObject /* cross-framework: NSData */, errorPtr objectivec.IObject) []objc.ID
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKGeoJSONDecoder */
// Alloc allocates a new instance without initialization.
func (mc _MKGeoJSONDecoderClass) Alloc() MKGeoJSONDecoder {
	rv := objc.Send[MKGeoJSONDecoder](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKGeoJSONDecoderClass) New() MKGeoJSONDecoder {
	rv := objc.Send[MKGeoJSONDecoder](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKGeoJSONDecoder) Init() MKGeoJSONDecoder {
	rv := objc.Send[MKGeoJSONDecoder](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKGeoJSONDecoder) Autorelease() MKGeoJSONDecoder {
	rv := objc.Send[MKGeoJSONDecoder](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKGeoJSONDecoder creates a new MKGeoJSONDecoder instance.
func NewMKGeoJSONDecoder() MKGeoJSONDecoder {
	return getMKGeoJSONDecoderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKGeoJSONDecoder */
// An object that decodes GeoJSON objects into MapKit types.
//
// The GeoJSON decoder returns objects that conform to the protocol.


// An object that decodes GeoJSON objects into MapKit types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeoJSONDecoder
type MKGeoJSONDecoder struct {
	objectivec.Object
}

// MKGeoJSONDecoderFrom constructs a [MKGeoJSONDecoder] from an unsafe.Pointer.
//
// An object that decodes GeoJSON objects into MapKit types.
func MKGeoJSONDecoderFrom(ptr unsafe.Pointer) MKGeoJSONDecoder {
	return MKGeoJSONDecoder{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKGeoJSONDecoder *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKGeoJSONDecoder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKGeoJSONDecoder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKGeoJSONDecoder */

// Decodes the provided data into native MapKit types that a map can display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeoJSONDecoder/decode(_:)
func (m_ MKGeoJSONDecoder) GeoJSONObjectsWithDataError(data objc.IObject /* cross-framework: NSData */, errorPtr objectivec.IObject) []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("geoJSONObjectsWithData:error:"), data, errorPtr)
	return rv
}/* debug [instance_methods/method]: GeoJSONObjectsWithDataError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKGeoJSONDecoder */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKGeoJSONDecoder */



