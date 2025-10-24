// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMSensorDataList */


/* debug [class_header]: Header for CMSensorDataList */
// The class instance for the [SensorDataList] class.
var (
	SensorDataListClass     _SensorDataListClass
	SensorDataListClassOnce sync.Once
)

func getSensorDataListClass() _SensorDataListClass {
	SensorDataListClassOnce.Do(func() {
		SensorDataListClass = _SensorDataListClass{objc.GetClass("CMSensorDataList")}
	})
	return SensorDataListClass
}

type _SensorDataListClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SensorDataList */
// An interface definition for the [SensorDataList] class.
type ISensorDataList interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SensorDataList */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SensorDataList */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SensorDataList */
// Alloc allocates a new instance without initialization.
func (sc _SensorDataListClass) Alloc() SensorDataList {
	rv := objc.Send[SensorDataList](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SensorDataListClass) New() SensorDataList {
	rv := objc.Send[SensorDataList](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SensorDataList) Init() SensorDataList {
	rv := objc.Send[SensorDataList](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SensorDataList) Autorelease() SensorDataList {
	rv := objc.Send[SensorDataList](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSensorDataList creates a new SensorDataList instance.
func NewSensorDataList() SensorDataList {
	return getSensorDataListClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SensorDataList */
// A list of the accelerometer data recorded by the system.
//
// You do not create instances of this class directly. Instead, you receive one as the result of a query for accelerometer data from a object. You use a sensor data list object to enumerate over the accelerometer data as shown in the following example:


// A list of the accelerometer data recorded by the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMSensorDataList
type SensorDataList struct {
	objectivec.Object
}

// SensorDataListFrom constructs a [SensorDataList] from an unsafe.Pointer.
//
// A list of the accelerometer data recorded by the system.
func SensorDataListFrom(ptr unsafe.Pointer) SensorDataList {
	return SensorDataList{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SensorDataList *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SensorDataList */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SensorDataList */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SensorDataList */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SensorDataList */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMSensorDataList */



