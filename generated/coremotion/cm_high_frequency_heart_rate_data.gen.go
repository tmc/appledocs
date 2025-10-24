// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CMHighFrequencyHeartRateData */


/* debug [class_header]: Header for CMHighFrequencyHeartRateData */
// The class instance for the [HighFrequencyHeartRateData] class.
var (
	HighFrequencyHeartRateDataClass     _HighFrequencyHeartRateDataClass
	HighFrequencyHeartRateDataClassOnce sync.Once
)

func getHighFrequencyHeartRateDataClass() _HighFrequencyHeartRateDataClass {
	HighFrequencyHeartRateDataClassOnce.Do(func() {
		HighFrequencyHeartRateDataClass = _HighFrequencyHeartRateDataClass{objc.GetClass("CMHighFrequencyHeartRateData")}
	})
	return HighFrequencyHeartRateDataClass
}

type _HighFrequencyHeartRateDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HighFrequencyHeartRateData */
// An interface definition for the [HighFrequencyHeartRateData] class.
type IHighFrequencyHeartRateData interface {
	ILogItem
	
/* debug [class_interface_properties]: Properties for HighFrequencyHeartRateData */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HighFrequencyHeartRateData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HighFrequencyHeartRateData */
// Alloc allocates a new instance without initialization.
func (hc _HighFrequencyHeartRateDataClass) Alloc() HighFrequencyHeartRateData {
	rv := objc.Send[HighFrequencyHeartRateData](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HighFrequencyHeartRateDataClass) New() HighFrequencyHeartRateData {
	rv := objc.Send[HighFrequencyHeartRateData](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HighFrequencyHeartRateData) Init() HighFrequencyHeartRateData {
	rv := objc.Send[HighFrequencyHeartRateData](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HighFrequencyHeartRateData) Autorelease() HighFrequencyHeartRateData {
	rv := objc.Send[HighFrequencyHeartRateData](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHighFrequencyHeartRateData creates a new HighFrequencyHeartRateData instance.
func NewHighFrequencyHeartRateData() HighFrequencyHeartRateData {
	return getHighFrequencyHeartRateDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HighFrequencyHeartRateData */
// A class that represents heart rate data collected at 1 Hz.
//
// Use the property to get the data, and the property for the accuracy.


// A class that represents heart rate data collected at 1 Hz.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateData
type HighFrequencyHeartRateData struct {
	LogItem
}

// HighFrequencyHeartRateDataFrom constructs a [HighFrequencyHeartRateData] from an unsafe.Pointer.
//
// A class that represents heart rate data collected at 1 Hz.
func HighFrequencyHeartRateDataFrom(ptr unsafe.Pointer) HighFrequencyHeartRateData {
	return HighFrequencyHeartRateData{
		LogItem: LogItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HighFrequencyHeartRateData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HighFrequencyHeartRateData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HighFrequencyHeartRateData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HighFrequencyHeartRateData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HighFrequencyHeartRateData */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMHighFrequencyHeartRateData */


