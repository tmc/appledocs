// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMPedometerData */


/* debug [class_header]: Header for CMPedometerData */
// The class instance for the [PedometerData] class.
var (
	PedometerDataClass     _PedometerDataClass
	PedometerDataClassOnce sync.Once
)

func getPedometerDataClass() _PedometerDataClass {
	PedometerDataClassOnce.Do(func() {
		PedometerDataClass = _PedometerDataClass{objc.GetClass("CMPedometerData")}
	})
	return PedometerDataClass
}

type _PedometerDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PedometerData */
// An interface definition for the [PedometerData] class.
type IPedometerData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PedometerData */
	// properties:
	AverageActivePace() objc.IObject /* cross-framework: NSNumber */
	CurrentCadence() objc.IObject /* cross-framework: NSNumber */
	CurrentPace() objc.IObject /* cross-framework: NSNumber */
	Distance() objc.IObject /* cross-framework: NSNumber */
	EndDate() objc.IObject /* cross-framework: NSDate */
	FloorsAscended() objc.IObject /* cross-framework: NSNumber */
	FloorsDescended() objc.IObject /* cross-framework: NSNumber */
	NumberOfSteps() objc.IObject /* cross-framework: NSNumber */
	StartDate() objc.IObject /* cross-framework: NSDate */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PedometerData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PedometerData */
// Alloc allocates a new instance without initialization.
func (pc _PedometerDataClass) Alloc() PedometerData {
	rv := objc.Send[PedometerData](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PedometerDataClass) New() PedometerData {
	rv := objc.Send[PedometerData](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PedometerData) Init() PedometerData {
	rv := objc.Send[PedometerData](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PedometerData) Autorelease() PedometerData {
	rv := objc.Send[PedometerData](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPedometerData creates a new PedometerData instance.
func NewPedometerData() PedometerData {
	return getPedometerDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PedometerData */
// Information about the distance traveled by a user on foot.
//
// You do not create instances of this class yourself. Instead, you use a object to request pedometer data from the system. The data for each request is packaged into an instance of this class and delivered to the handlers you registered with the pedometer object.


// Information about the distance traveled by a user on foot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData
type PedometerData struct {
	objectivec.Object
}

// PedometerDataFrom constructs a [PedometerData] from an unsafe.Pointer.
//
// Information about the distance traveled by a user on foot.
func PedometerDataFrom(ptr unsafe.Pointer) PedometerData {
	return PedometerData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PedometerData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PedometerData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PedometerData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PedometerData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PedometerData */

// The average pace of the user, measured in seconds per meter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/averageActivePace
func (p_ PedometerData) AverageActivePace() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](p_.ID, objc.Sel("averageActivePace"))
	return rv
}/* debug [instance_properties/getter]: averageActivePace */


// The rate at which steps are taken, measured in steps per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/currentCadence
func (p_ PedometerData) CurrentCadence() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](p_.ID, objc.Sel("currentCadence"))
	return rv
}/* debug [instance_properties/getter]: currentCadence */


// The current pace of the user, measured in seconds per meter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/currentPace
func (p_ PedometerData) CurrentPace() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](p_.ID, objc.Sel("currentPace"))
	return rv
}/* debug [instance_properties/getter]: currentPace */


// The estimated distance (in meters) traveled by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/distance
func (p_ PedometerData) Distance() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](p_.ID, objc.Sel("distance"))
	return rv
}/* debug [instance_properties/getter]: distance */


// The end time for the pedometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/endDate
func (p_ PedometerData) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// The approximate number of floors ascended by walking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/floorsAscended
func (p_ PedometerData) FloorsAscended() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](p_.ID, objc.Sel("floorsAscended"))
	return rv
}/* debug [instance_properties/getter]: floorsAscended */


// The approximate number of floors descended by walking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/floorsDescended
func (p_ PedometerData) FloorsDescended() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](p_.ID, objc.Sel("floorsDescended"))
	return rv
}/* debug [instance_properties/getter]: floorsDescended */


// The number of steps taken by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/numberOfSteps
func (p_ PedometerData) NumberOfSteps() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](p_.ID, objc.Sel("numberOfSteps"))
	return rv
}/* debug [instance_properties/getter]: numberOfSteps */


// The start time for the pedometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerData/startDate
func (p_ PedometerData) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMPedometerData */



