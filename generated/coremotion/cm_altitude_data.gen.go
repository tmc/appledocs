// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CMAltitudeData */


/* debug [class_header]: Header for CMAltitudeData */
// The class instance for the [AltitudeData] class.
var (
	AltitudeDataClass     _AltitudeDataClass
	AltitudeDataClassOnce sync.Once
)

func getAltitudeDataClass() _AltitudeDataClass {
	AltitudeDataClassOnce.Do(func() {
		AltitudeDataClass = _AltitudeDataClass{objc.GetClass("CMAltitudeData")}
	})
	return AltitudeDataClass
}

type _AltitudeDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AltitudeData */
// An interface definition for the [AltitudeData] class.
type IAltitudeData interface {
	ILogItem
	
/* debug [class_interface_properties]: Properties for AltitudeData */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AltitudeData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AltitudeData */
// Alloc allocates a new instance without initialization.
func (ac _AltitudeDataClass) Alloc() AltitudeData {
	rv := objc.Send[AltitudeData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AltitudeDataClass) New() AltitudeData {
	rv := objc.Send[AltitudeData](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AltitudeData) Init() AltitudeData {
	rv := objc.Send[AltitudeData](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AltitudeData) Autorelease() AltitudeData {
	rv := objc.Send[AltitudeData](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAltitudeData creates a new AltitudeData instance.
func NewAltitudeData() AltitudeData {
	return getAltitudeDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AltitudeData */
// Data for a recorded change in altitude.
//
// You do not create instances of this class directly. When you want to receive altimeter changes, create an instance of the class and use that object to query for events or to start the delivery of events. The altimeter object creates new instances of this class at appropriate times and delivers them to the handler you specify.


// Data for a recorded change in altitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltitudeData
type AltitudeData struct {
	LogItem
}

// AltitudeDataFrom constructs a [AltitudeData] from an unsafe.Pointer.
//
// Data for a recorded change in altitude.
func AltitudeDataFrom(ptr unsafe.Pointer) AltitudeData {
	return AltitudeData{
		LogItem: LogItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AltitudeData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AltitudeData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AltitudeData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AltitudeData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AltitudeData */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMAltitudeData */


