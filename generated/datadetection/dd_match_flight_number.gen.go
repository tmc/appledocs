// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DDMatchFlightNumber */


/* debug [class_header]: Header for DDMatchFlightNumber */
// The class instance for the [DDMatchFlightNumber] class.
var (
	DDMatchFlightNumberClass     _DDMatchFlightNumberClass
	DDMatchFlightNumberClassOnce sync.Once
)

func getDDMatchFlightNumberClass() _DDMatchFlightNumberClass {
	DDMatchFlightNumberClassOnce.Do(func() {
		DDMatchFlightNumberClass = _DDMatchFlightNumberClass{objc.GetClass("DDMatchFlightNumber")}
	})
	return DDMatchFlightNumberClass
}

type _DDMatchFlightNumberClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DDMatchFlightNumber */
// An interface definition for the [DDMatchFlightNumber] class.
type IDDMatchFlightNumber interface {
	IDDMatch
	
/* debug [class_interface_properties]: Properties for DDMatchFlightNumber */
	// properties:
	Airline() objc.IObject /* cross-framework: NSString */
	FlightNumber() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DDMatchFlightNumber */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DDMatchFlightNumber */
// Alloc allocates a new instance without initialization.
func (dc _DDMatchFlightNumberClass) Alloc() DDMatchFlightNumber {
	rv := objc.Send[DDMatchFlightNumber](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DDMatchFlightNumberClass) New() DDMatchFlightNumber {
	rv := objc.Send[DDMatchFlightNumber](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatchFlightNumber) Init() DDMatchFlightNumber {
	rv := objc.Send[DDMatchFlightNumber](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatchFlightNumber) Autorelease() DDMatchFlightNumber {
	rv := objc.Send[DDMatchFlightNumber](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatchFlightNumber creates a new DDMatchFlightNumber instance.
func NewDDMatchFlightNumber() DDMatchFlightNumber {
	return getDDMatchFlightNumberClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DDMatchFlightNumber */
// An object that contains a flight number that the data detection system matches.
//
// The DataDetection framework returns a flight number match in a object, which contains an airline name and flight number.


// An object that contains a flight number that the data detection system matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchFlightNumber
type DDMatchFlightNumber struct {
	DDMatch
}

// DDMatchFlightNumberFrom constructs a [DDMatchFlightNumber] from an unsafe.Pointer.
//
// An object that contains a flight number that the data detection system matches.
func DDMatchFlightNumberFrom(ptr unsafe.Pointer) DDMatchFlightNumber {
	return DDMatchFlightNumber{
		DDMatch: DDMatchFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DDMatchFlightNumber *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DDMatchFlightNumber */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DDMatchFlightNumber */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DDMatchFlightNumber */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DDMatchFlightNumber */

// The name of an airline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchFlightNumber/airline
func (d_ DDMatchFlightNumber) Airline() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("airline"))
	return rv
}/* debug [instance_properties/getter]: airline */


// A string that represents a flight number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchFlightNumber/flightNumber
func (d_ DDMatchFlightNumber) FlightNumber() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("flightNumber"))
	return rv
}/* debug [instance_properties/getter]: flightNumber */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DDMatchFlightNumber */



