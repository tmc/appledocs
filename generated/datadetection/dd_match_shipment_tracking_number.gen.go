// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DDMatchShipmentTrackingNumber */


/* debug [class_header]: Header for DDMatchShipmentTrackingNumber */
// The class instance for the [DDMatchShipmentTrackingNumber] class.
var (
	DDMatchShipmentTrackingNumberClass     _DDMatchShipmentTrackingNumberClass
	DDMatchShipmentTrackingNumberClassOnce sync.Once
)

func getDDMatchShipmentTrackingNumberClass() _DDMatchShipmentTrackingNumberClass {
	DDMatchShipmentTrackingNumberClassOnce.Do(func() {
		DDMatchShipmentTrackingNumberClass = _DDMatchShipmentTrackingNumberClass{objc.GetClass("DDMatchShipmentTrackingNumber")}
	})
	return DDMatchShipmentTrackingNumberClass
}

type _DDMatchShipmentTrackingNumberClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DDMatchShipmentTrackingNumber */
// An interface definition for the [DDMatchShipmentTrackingNumber] class.
type IDDMatchShipmentTrackingNumber interface {
	IDDMatch
	
/* debug [class_interface_properties]: Properties for DDMatchShipmentTrackingNumber */
	// properties:
	Carrier() objc.IObject /* cross-framework: NSString */
	TrackingNumber() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DDMatchShipmentTrackingNumber */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DDMatchShipmentTrackingNumber */
// Alloc allocates a new instance without initialization.
func (dc _DDMatchShipmentTrackingNumberClass) Alloc() DDMatchShipmentTrackingNumber {
	rv := objc.Send[DDMatchShipmentTrackingNumber](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DDMatchShipmentTrackingNumberClass) New() DDMatchShipmentTrackingNumber {
	rv := objc.Send[DDMatchShipmentTrackingNumber](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatchShipmentTrackingNumber) Init() DDMatchShipmentTrackingNumber {
	rv := objc.Send[DDMatchShipmentTrackingNumber](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatchShipmentTrackingNumber) Autorelease() DDMatchShipmentTrackingNumber {
	rv := objc.Send[DDMatchShipmentTrackingNumber](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatchShipmentTrackingNumber creates a new DDMatchShipmentTrackingNumber instance.
func NewDDMatchShipmentTrackingNumber() DDMatchShipmentTrackingNumber {
	return getDDMatchShipmentTrackingNumberClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DDMatchShipmentTrackingNumber */
// An object that contains parcel tracking information that the data detection system matches.
//
// The DataDetection framework returns a shipment tracking number match in a object, which contains a carrier name and tracking identifier.


// An object that contains parcel tracking information that the data detection system matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchShipmentTrackingNumber
type DDMatchShipmentTrackingNumber struct {
	DDMatch
}

// DDMatchShipmentTrackingNumberFrom constructs a [DDMatchShipmentTrackingNumber] from an unsafe.Pointer.
//
// An object that contains parcel tracking information that the data detection system matches.
func DDMatchShipmentTrackingNumberFrom(ptr unsafe.Pointer) DDMatchShipmentTrackingNumber {
	return DDMatchShipmentTrackingNumber{
		DDMatch: DDMatchFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DDMatchShipmentTrackingNumber *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DDMatchShipmentTrackingNumber */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DDMatchShipmentTrackingNumber */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DDMatchShipmentTrackingNumber */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DDMatchShipmentTrackingNumber */

// The name of a parcel carrier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchShipmentTrackingNumber/carrier
func (d_ DDMatchShipmentTrackingNumber) Carrier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("carrier"))
	return rv
}/* debug [instance_properties/getter]: carrier */


// A string that represents a carrier’s tracking identifier for a parcel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchShipmentTrackingNumber/trackingNumber
func (d_ DDMatchShipmentTrackingNumber) TrackingNumber() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("trackingNumber"))
	return rv
}/* debug [instance_properties/getter]: trackingNumber */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DDMatchShipmentTrackingNumber */






