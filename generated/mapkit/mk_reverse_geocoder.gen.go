// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKReverseGeocoder */


/* debug [class_header]: Header for MKReverseGeocoder */
// The class instance for the [MKReverseGeocoder] class.
var (
	MKReverseGeocoderClass     _MKReverseGeocoderClass
	MKReverseGeocoderClassOnce sync.Once
)

func getMKReverseGeocoderClass() _MKReverseGeocoderClass {
	MKReverseGeocoderClassOnce.Do(func() {
		MKReverseGeocoderClass = _MKReverseGeocoderClass{objc.GetClass("MKReverseGeocoder")}
	})
	return MKReverseGeocoderClass
}

type _MKReverseGeocoderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKReverseGeocoder */
// An interface definition for the [MKReverseGeocoder] class.
type IMKReverseGeocoder interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKReverseGeocoder */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKReverseGeocoder */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKReverseGeocoder */
// Alloc allocates a new instance without initialization.
func (mc _MKReverseGeocoderClass) Alloc() MKReverseGeocoder {
	rv := objc.Send[MKReverseGeocoder](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKReverseGeocoderClass) New() MKReverseGeocoder {
	rv := objc.Send[MKReverseGeocoder](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKReverseGeocoder) Init() MKReverseGeocoder {
	rv := objc.Send[MKReverseGeocoder](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKReverseGeocoder) Autorelease() MKReverseGeocoder {
	rv := objc.Send[MKReverseGeocoder](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKReverseGeocoder creates a new MKReverseGeocoder instance.
func NewMKReverseGeocoder() MKReverseGeocoder {
	return getMKReverseGeocoderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKReverseGeocoder */
// Provides services for converting a map coordinate (specified as a latitude/longitude pair) into information about that coordinate, such as the country or region, city, or street.
//
// A reverse geocoder object is a single-shot object that works with a network-based map service to look up placemark information for its specified coordinate value. The Google terms of service require that the reverse geocoding service be used in conjunction with a Google map; take this into account when designing your application’s user interface. Each Map Kit application has a limited amount of reverse geocoding capacity, so it is to your advantage to use reverse geocode requests sparingly. Here are some rules of thumb for using this class most effectively: Send at most one reverse-geocoding request for any one user action. If the user performs multiple actions that involve reverse-geocoding the same location, reuse the results from the initial reverse-geocoding request instead of starting individual requests for each action. When you want to update the location automatically (such as when the user is moving), reissue the reverse-geocoding request only when the user’s location has moved a significant distance and after a reasonable amount of time has passed. For example, in a typical situation, you should not send more than one reverse-geocode request per minute. Do not start a reverse-geocoding request at a time when the user will not see the results immediately. For example, do not start a request if your application recently resigned the active state (possibly because of an interruption such as a phone call) and is waiting to become active again. An iOS-based device must have access to the network in order for the reverse geocoder object to return valid information. The reverse geocoder returns information through its associated delegate object, which is an object that conforms to the protocol. If the reverse geocoder is unable to retrieve the requested information, it similarly reports the error to its delegate object. For more information on this protocol, see . This class is deprecated in iOS 5.0. Use the class instead.


// Provides services for converting a map coordinate (specified as a latitude/longitude pair) into information about that coordinate, such as the country or region, city, or street.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocoder
type MKReverseGeocoder struct {
	objectivec.Object
}

// MKReverseGeocoderFrom constructs a [MKReverseGeocoder] from an unsafe.Pointer.
//
// Provides services for converting a map coordinate (specified as a latitude/longitude pair) into information about that coordinate, such as the country or region, city, or street.
func MKReverseGeocoderFrom(ptr unsafe.Pointer) MKReverseGeocoder {
	return MKReverseGeocoder{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKReverseGeocoder */

// Initializes the reverse geocoder with the specified coordinate value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocoder/initWithCoordinate:
func NewMKReverseGeocoderWithCoordinate(coordinate LocationCoordinate2D /* not a class type */) MKReverseGeocoder {
	instance := getMKReverseGeocoderClass().Alloc()
	rv := objc.Send[MKReverseGeocoder](instance.ID, objc.Sel("initWithCoordinate:"), coordinate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKReverseGeocoderWithCoordinate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKReverseGeocoder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKReverseGeocoder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKReverseGeocoder */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKReverseGeocoder */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKReverseGeocoder */


