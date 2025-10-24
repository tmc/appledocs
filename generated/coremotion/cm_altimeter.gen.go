// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMAltimeter */


/* debug [class_header]: Header for CMAltimeter */
// The class instance for the [Altimeter] class.
var (
	AltimeterClass     _AltimeterClass
	AltimeterClassOnce sync.Once
)

func getAltimeterClass() _AltimeterClass {
	AltimeterClassOnce.Do(func() {
		AltimeterClass = _AltimeterClass{objc.GetClass("CMAltimeter")}
	})
	return AltimeterClass
}

type _AltimeterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Altimeter */
// An interface definition for the [Altimeter] class.
type IAltimeter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Altimeter */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Altimeter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Altimeter */
// Alloc allocates a new instance without initialization.
func (ac _AltimeterClass) Alloc() Altimeter {
	rv := objc.Send[Altimeter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AltimeterClass) New() Altimeter {
	rv := objc.Send[Altimeter](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Altimeter) Init() Altimeter {
	rv := objc.Send[Altimeter](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Altimeter) Autorelease() Altimeter {
	rv := objc.Send[Altimeter](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAltimeter creates a new Altimeter instance.
func NewAltimeter() Altimeter {
	return getAltimeterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Altimeter */
// An object that initiates the delivery of altitude-related changes.
//
// Altitude events report changes in both the relative and absolute altitude. For example, a hiking app could use this object to track the user’s elevation change over the course of a hike, or to report their current absolute altitude during the hike. Because altitude events may not be available on all devices, always call the method before starting relative altitude updates, and call before starting absolute altitude updates. After checking the availability of altitude data, call the method to start receiving relative altitude data, or call the method for absolute altitude data. Core Motion generates events at regular intervals (regardless of whether the data has changed) and delivers them to the block you specified. When you no longer need the event data, call the or methods respectively.


// An object that initiates the delivery of altitude-related changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter
type Altimeter struct {
	objectivec.Object
}

// AltimeterFrom constructs a [Altimeter] from an unsafe.Pointer.
//
// An object that initiates the delivery of altitude-related changes.
func AltimeterFrom(ptr unsafe.Pointer) Altimeter {
	return Altimeter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Altimeter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Altimeter */

// Returns a value indicating whether the app is authorized to retrieve altimeter data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/authorizationStatus()
func (ac _AltimeterClass) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](objc.ID(ac.class), objc.Sel("authorizationStatus"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AuthorizationStatus) */


// Returns a Boolean value indicating whether the current device reports changes in the absolute altitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/isAbsoluteAltitudeAvailable()
func (ac _AltimeterClass) IsAbsoluteAltitudeAvailable() bool {
	rv := objc.Send[bool](objc.ID(ac.class), objc.Sel("isAbsoluteAltitudeAvailable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsAbsoluteAltitudeAvailable) */


// Returns a Boolean value indicating whether the current device supports generating data for relative altitude changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltimeter/isRelativeAltitudeAvailable()
func (ac _AltimeterClass) IsRelativeAltitudeAvailable() bool {
	rv := objc.Send[bool](objc.ID(ac.class), objc.Sel("isRelativeAltitudeAvailable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsRelativeAltitudeAvailable) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Altimeter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Altimeter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Altimeter */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMAltimeter */


