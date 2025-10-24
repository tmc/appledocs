// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MXSignpostIntervalData */


/* debug [class_header]: Header for MXSignpostIntervalData */
// The class instance for the [MXSignpostIntervalData] class.
var (
	MXSignpostIntervalDataClass     _MXSignpostIntervalDataClass
	MXSignpostIntervalDataClassOnce sync.Once
)

func getMXSignpostIntervalDataClass() _MXSignpostIntervalDataClass {
	MXSignpostIntervalDataClassOnce.Do(func() {
		MXSignpostIntervalDataClass = _MXSignpostIntervalDataClass{objc.GetClass("MXSignpostIntervalData")}
	})
	return MXSignpostIntervalDataClass
}

type _MXSignpostIntervalDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXSignpostIntervalData */
// An interface definition for the [MXSignpostIntervalData] class.
type IMXSignpostIntervalData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MXSignpostIntervalData */
	// properties:
	AverageMemory() unsafe.Pointer
	CumulativeCPUTime() unsafe.Pointer
	CumulativeHitchTimeRatio() unsafe.Pointer
	CumulativeLogicalWrites() unsafe.Pointer
	HistogrammedSignpostDuration() unsafe.Pointer
	SignpostIntervalData() IMXSignpostIntervalData
	SetSignpostIntervalData(value IMXSignpostIntervalData)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXSignpostIntervalData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXSignpostIntervalData */
// Alloc allocates a new instance without initialization.
func (mc _MXSignpostIntervalDataClass) Alloc() MXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXSignpostIntervalDataClass) New() MXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXSignpostIntervalData) Init() MXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXSignpostIntervalData) Autorelease() MXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXSignpostIntervalData creates a new MXSignpostIntervalData instance.
func NewMXSignpostIntervalData() MXSignpostIntervalData {
	return getMXSignpostIntervalDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXSignpostIntervalData */
// A data object representing the captured data for a custom metric.


// A data object representing the captured data for a custom metric.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostIntervalData
type MXSignpostIntervalData struct {
	objectivec.Object
}

// MXSignpostIntervalDataFrom constructs a [MXSignpostIntervalData] from an unsafe.Pointer.
//
// A data object representing the captured data for a custom metric.
func MXSignpostIntervalDataFrom(ptr unsafe.Pointer) MXSignpostIntervalData {
	return MXSignpostIntervalData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXSignpostIntervalData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXSignpostIntervalData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXSignpostIntervalData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXSignpostIntervalData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXSignpostIntervalData */

// The average memory used during the logged intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostIntervalData/averageMemory
func (m_ MXSignpostIntervalData) AverageMemory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("averageMemory"))
	return rv
}/* debug [instance_properties/getter]: averageMemory */


// The total amount of CPU time used during the logged intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostIntervalData/cumulativeCPUTime
func (m_ MXSignpostIntervalData) CumulativeCPUTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeCPUTime"))
	return rv
}/* debug [instance_properties/getter]: cumulativeCPUTime */


// The ratio of the total time spent hitching to the total time spent animating during the logged intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostIntervalData/cumulativeHitchTimeRatio
func (m_ MXSignpostIntervalData) CumulativeHitchTimeRatio() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeHitchTimeRatio"))
	return rv
}/* debug [instance_properties/getter]: cumulativeHitchTimeRatio */


// The total amount of data written to disk or other long term storage during the logged intervals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostIntervalData/cumulativeLogicalWrites
func (m_ MXSignpostIntervalData) CumulativeLogicalWrites() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cumulativeLogicalWrites"))
	return rv
}/* debug [instance_properties/getter]: cumulativeLogicalWrites */


// A histogram of the different time intervals of a custom metric event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostIntervalData/histogrammedSignpostDuration
func (m_ MXSignpostIntervalData) HistogrammedSignpostDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("histogrammedSignpostDuration"))
	return rv
}/* debug [instance_properties/getter]: histogrammedSignpostDuration */


// The data captured for a custom metric.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostmetric/signpostintervaldata
func (m_ MXSignpostIntervalData) SignpostIntervalData() IMXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](m_.ID, objc.Sel("signpostIntervalData"))
	return rv
}/* debug [instance_properties/getter]: signpostIntervalData */


// The data captured for a custom metric.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxsignpostmetric/signpostintervaldata
func (m_ MXSignpostIntervalData) SetSignpostIntervalData(value IMXSignpostIntervalData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSignpostIntervalData:"), value)
}/* debug [instance_properties/setter]: signpostIntervalData */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXSignpostIntervalData */



