// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NIDLTDOAMeasurement */


/* debug [class_header]: Header for NIDLTDOAMeasurement */
// The class instance for the [NIDLTDOAMeasurement] class.
var (
	NIDLTDOAMeasurementClass     _NIDLTDOAMeasurementClass
	NIDLTDOAMeasurementClassOnce sync.Once
)

func getNIDLTDOAMeasurementClass() _NIDLTDOAMeasurementClass {
	NIDLTDOAMeasurementClassOnce.Do(func() {
		NIDLTDOAMeasurementClass = _NIDLTDOAMeasurementClass{objc.GetClass("NIDLTDOAMeasurement")}
	})
	return NIDLTDOAMeasurementClass
}

type _NIDLTDOAMeasurementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NIDLTDOAMeasurement */
// An interface definition for the [NIDLTDOAMeasurement] class.
type INIDLTDOAMeasurement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NIDLTDOAMeasurement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NIDLTDOAMeasurement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NIDLTDOAMeasurement */
// Alloc allocates a new instance without initialization.
func (nc _NIDLTDOAMeasurementClass) Alloc() NIDLTDOAMeasurement {
	rv := objc.Send[NIDLTDOAMeasurement](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NIDLTDOAMeasurementClass) New() NIDLTDOAMeasurement {
	rv := objc.Send[NIDLTDOAMeasurement](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NIDLTDOAMeasurement) Init() NIDLTDOAMeasurement {
	rv := objc.Send[NIDLTDOAMeasurement](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NIDLTDOAMeasurement) Autorelease() NIDLTDOAMeasurement {
	rv := objc.Send[NIDLTDOAMeasurement](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNIDLTDOAMeasurement creates a new NIDLTDOAMeasurement instance.
func NewNIDLTDOAMeasurement() NIDLTDOAMeasurement {
	return getNIDLTDOAMeasurementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NIDLTDOAMeasurement */
// Information from a Downlink Time-Difference-of-Arrival anchor that you use to derive a range estimate.
//
// Your app runs on a receiver device that fields messages from nearby physical base stations, or . The framework processes the messages into instances of this class and provides them to your app through the callback. Your app analyzes the measurements to calculate the receiver’s position relative to the anchors in the tracked area. Only sessions that run a receive Downlink Time-Difference-of-Arrival measurements.


// Information from a Downlink Time-Difference-of-Arrival anchor that you use to derive a range estimate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurement
type NIDLTDOAMeasurement struct {
	objectivec.Object
}

// NIDLTDOAMeasurementFrom constructs a [NIDLTDOAMeasurement] from an unsafe.Pointer.
//
// Information from a Downlink Time-Difference-of-Arrival anchor that you use to derive a range estimate.
func NIDLTDOAMeasurementFrom(ptr unsafe.Pointer) NIDLTDOAMeasurement {
	return NIDLTDOAMeasurement{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NIDLTDOAMeasurement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NIDLTDOAMeasurement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NIDLTDOAMeasurement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NIDLTDOAMeasurement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NIDLTDOAMeasurement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NIDLTDOAMeasurement */


