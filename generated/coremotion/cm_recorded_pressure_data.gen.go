// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CMRecordedPressureData */


/* debug [class_header]: Header for CMRecordedPressureData */
// The class instance for the [RecordedPressureData] class.
var (
	RecordedPressureDataClass     _RecordedPressureDataClass
	RecordedPressureDataClassOnce sync.Once
)

func getRecordedPressureDataClass() _RecordedPressureDataClass {
	RecordedPressureDataClassOnce.Do(func() {
		RecordedPressureDataClass = _RecordedPressureDataClass{objc.GetClass("CMRecordedPressureData")}
	})
	return RecordedPressureDataClass
}

type _RecordedPressureDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RecordedPressureData */
// An interface definition for the [RecordedPressureData] class.
type IRecordedPressureData interface {
	IAmbientPressureData
	
/* debug [class_interface_properties]: Properties for RecordedPressureData */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RecordedPressureData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RecordedPressureData */
// Alloc allocates a new instance without initialization.
func (rc _RecordedPressureDataClass) Alloc() RecordedPressureData {
	rv := objc.Send[RecordedPressureData](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RecordedPressureDataClass) New() RecordedPressureData {
	rv := objc.Send[RecordedPressureData](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecordedPressureData) Init() RecordedPressureData {
	rv := objc.Send[RecordedPressureData](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecordedPressureData) Autorelease() RecordedPressureData {
	rv := objc.Send[RecordedPressureData](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecordedPressureData creates a new RecordedPressureData instance.
func NewRecordedPressureData() RecordedPressureData {
	return getRecordedPressureDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RecordedPressureData */
// A recorded measurement of pressure data.
//
// Use SensorKit’s sensor to read ambient pressure data.


// A recorded measurement of pressure data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRecordedPressureData
type RecordedPressureData struct {
	AmbientPressureData
}

// RecordedPressureDataFrom constructs a [RecordedPressureData] from an unsafe.Pointer.
//
// A recorded measurement of pressure data.
func RecordedPressureDataFrom(ptr unsafe.Pointer) RecordedPressureData {
	return RecordedPressureData{
		AmbientPressureData: AmbientPressureDataFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RecordedPressureData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RecordedPressureData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RecordedPressureData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RecordedPressureData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RecordedPressureData */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMRecordedPressureData */


