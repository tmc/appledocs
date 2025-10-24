// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CMRecordedRotationRateData */


/* debug [class_header]: Header for CMRecordedRotationRateData */
// The class instance for the [RecordedRotationRateData] class.
var (
	RecordedRotationRateDataClass     _RecordedRotationRateDataClass
	RecordedRotationRateDataClassOnce sync.Once
)

func getRecordedRotationRateDataClass() _RecordedRotationRateDataClass {
	RecordedRotationRateDataClassOnce.Do(func() {
		RecordedRotationRateDataClass = _RecordedRotationRateDataClass{objc.GetClass("CMRecordedRotationRateData")}
	})
	return RecordedRotationRateDataClass
}

type _RecordedRotationRateDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RecordedRotationRateData */
// An interface definition for the [RecordedRotationRateData] class.
type IRecordedRotationRateData interface {
	IRotationRateData
	
/* debug [class_interface_properties]: Properties for RecordedRotationRateData */
	// properties:
	StartDate() objc.IObject /* cross-framework: NSDate */
	RotationRate() objc.IObject /* cross-framework: CMRotationRate */
	SetRotationRate(value objc.IObject /* cross-framework: CMRotationRate */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RecordedRotationRateData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RecordedRotationRateData */
// Alloc allocates a new instance without initialization.
func (rc _RecordedRotationRateDataClass) Alloc() RecordedRotationRateData {
	rv := objc.Send[RecordedRotationRateData](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RecordedRotationRateDataClass) New() RecordedRotationRateData {
	rv := objc.Send[RecordedRotationRateData](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecordedRotationRateData) Init() RecordedRotationRateData {
	rv := objc.Send[RecordedRotationRateData](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecordedRotationRateData) Autorelease() RecordedRotationRateData {
	rv := objc.Send[RecordedRotationRateData](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecordedRotationRateData creates a new RecordedRotationRateData instance.
func NewRecordedRotationRateData() RecordedRotationRateData {
	return getRecordedRotationRateDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RecordedRotationRateData */
// A data object that contains a single rotation-rate measurement at a specific time.


// A data object that contains a single rotation-rate measurement at a specific time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRecordedRotationRateData
type RecordedRotationRateData struct {
	RotationRateData
}

// RecordedRotationRateDataFrom constructs a [RecordedRotationRateData] from an unsafe.Pointer.
//
// A data object that contains a single rotation-rate measurement at a specific time.
func RecordedRotationRateDataFrom(ptr unsafe.Pointer) RecordedRotationRateData {
	return RecordedRotationRateData{
		RotationRateData: RotationRateDataFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RecordedRotationRateData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RecordedRotationRateData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RecordedRotationRateData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RecordedRotationRateData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RecordedRotationRateData */

// The time when the gyroscope measured the rotation data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRecordedRotationRateData/startDate
func (r_ RecordedRotationRateData) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](r_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */


// The rotation rate as measured by the device’s gyroscope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmgyrodata/rotationrate
func (r_ RecordedRotationRateData) RotationRate() objc.IObject /* cross-framework: CMRotationRate */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("rotationRate"))
	return rv
}/* debug [instance_properties/getter]: rotationRate */


// The rotation rate as measured by the device’s gyroscope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmgyrodata/rotationrate
func (r_ RecordedRotationRateData) SetRotationRate(value objc.IObject /* cross-framework: CMRotationRate */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRotationRate:"), value)
}/* debug [instance_properties/setter]: rotationRate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMRecordedRotationRateData */



