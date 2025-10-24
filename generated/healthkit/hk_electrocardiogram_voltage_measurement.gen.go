// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKElectrocardiogramVoltageMeasurement */


/* debug [class_header]: Header for HKElectrocardiogramVoltageMeasurement */
// The class instance for the [HKElectrocardiogramVoltageMeasurement] class.
var (
	HKElectrocardiogramVoltageMeasurementClass     _HKElectrocardiogramVoltageMeasurementClass
	HKElectrocardiogramVoltageMeasurementClassOnce sync.Once
)

func getHKElectrocardiogramVoltageMeasurementClass() _HKElectrocardiogramVoltageMeasurementClass {
	HKElectrocardiogramVoltageMeasurementClassOnce.Do(func() {
		HKElectrocardiogramVoltageMeasurementClass = _HKElectrocardiogramVoltageMeasurementClass{objc.GetClass("HKElectrocardiogramVoltageMeasurement")}
	})
	return HKElectrocardiogramVoltageMeasurementClass
}

type _HKElectrocardiogramVoltageMeasurementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKElectrocardiogramVoltageMeasurement */
// An interface definition for the [HKElectrocardiogramVoltageMeasurement] class.
type IHKElectrocardiogramVoltageMeasurement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKElectrocardiogramVoltageMeasurement */
	// properties:
	TimeSinceSampleStart() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKElectrocardiogramVoltageMeasurement */
	// methods:
	QuantityForLead(lead HKElectrocardiogramLead) IHKQuantity
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKElectrocardiogramVoltageMeasurement */
// Alloc allocates a new instance without initialization.
func (hc _HKElectrocardiogramVoltageMeasurementClass) Alloc() HKElectrocardiogramVoltageMeasurement {
	rv := objc.Send[HKElectrocardiogramVoltageMeasurement](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKElectrocardiogramVoltageMeasurementClass) New() HKElectrocardiogramVoltageMeasurement {
	rv := objc.Send[HKElectrocardiogramVoltageMeasurement](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKElectrocardiogramVoltageMeasurement) Init() HKElectrocardiogramVoltageMeasurement {
	rv := objc.Send[HKElectrocardiogramVoltageMeasurement](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKElectrocardiogramVoltageMeasurement) Autorelease() HKElectrocardiogramVoltageMeasurement {
	rv := objc.Send[HKElectrocardiogramVoltageMeasurement](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKElectrocardiogramVoltageMeasurement creates a new HKElectrocardiogramVoltageMeasurement instance.
func NewHKElectrocardiogramVoltageMeasurement() HKElectrocardiogramVoltageMeasurement {
	return getHKElectrocardiogramVoltageMeasurementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKElectrocardiogramVoltageMeasurement */
// The voltage for all leads at a single point in time.


// The voltage for all leads at a single point in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/VoltageMeasurement
type HKElectrocardiogramVoltageMeasurement struct {
	objectivec.Object
}

// HKElectrocardiogramVoltageMeasurementFrom constructs a [HKElectrocardiogramVoltageMeasurement] from an unsafe.Pointer.
//
// The voltage for all leads at a single point in time.
func HKElectrocardiogramVoltageMeasurementFrom(ptr unsafe.Pointer) HKElectrocardiogramVoltageMeasurement {
	return HKElectrocardiogramVoltageMeasurement{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKElectrocardiogramVoltageMeasurement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKElectrocardiogramVoltageMeasurement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKElectrocardiogramVoltageMeasurement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKElectrocardiogramVoltageMeasurement */

// Returns the voltage for the specified lead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/VoltageMeasurement/quantity(for:)
func (h_ HKElectrocardiogramVoltageMeasurement) QuantityForLead(lead HKElectrocardiogramLead) IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("quantityForLead:"), lead)
	return rv
}/* debug [instance_methods/method]: QuantityForLead */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKElectrocardiogramVoltageMeasurement */

// The time of the measurement relative to the sample’s start time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogram/VoltageMeasurement/timeSinceSampleStart
func (h_ HKElectrocardiogramVoltageMeasurement) TimeSinceSampleStart() float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("timeSinceSampleStart"))
	return rv
}/* debug [instance_properties/getter]: timeSinceSampleStart */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKElectrocardiogramVoltageMeasurement */



