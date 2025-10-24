// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MXHistogram */


/* debug [class_header]: Header for MXHistogram */
// The class instance for the [MXHistogram] class.
var (
	MXHistogramClass     _MXHistogramClass
	MXHistogramClassOnce sync.Once
)

func getMXHistogramClass() _MXHistogramClass {
	MXHistogramClassOnce.Do(func() {
		MXHistogramClass = _MXHistogramClass{objc.GetClass("MXHistogram")}
	})
	return MXHistogramClass
}

type _MXHistogramClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXHistogram */
// An interface definition for the [MXHistogram] class.
type IMXHistogram interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MXHistogram */
	// properties:
	BucketEnumerator() unsafe.Pointer
	TotalBucketCount() uint
	MXErrorDomain() objc.IObject /* cross-framework: NSString */
	BucketEnd() foundation.Measurement
	SetBucketEnd(value foundation.Measurement)
	BucketStart() foundation.Measurement
	SetBucketStart(value foundation.Measurement)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXHistogram */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXHistogram */
// Alloc allocates a new instance without initialization.
func (mc _MXHistogramClass) Alloc() MXHistogram {
	rv := objc.Send[MXHistogram](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXHistogramClass) New() MXHistogram {
	rv := objc.Send[MXHistogram](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXHistogram) Init() MXHistogram {
	rv := objc.Send[MXHistogram](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXHistogram) Autorelease() MXHistogram {
	rv := objc.Send[MXHistogram](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXHistogram creates a new MXHistogram instance.
func NewMXHistogram() MXHistogram {
	return getMXHistogramClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXHistogram */
// An object representing a histogram of data values of the same type of unit.
//
// A measures the number of times a data point for a variable falls into a specific range of possible values within a set of data. Usually, histograms are depicted as bar charts, in which each bar represents a range of values, and the height of each bar represents the number of times the value of the variable falls within a particular range. In this class, each bar is represented by a . A bucket holds the results for a series of measured values, such as all the events occurring between 3 and 5 seconds. MetricKit uses fixed-width buckets that are device-independent with intervals that are based on the type of metric. Use the and properties to find the start and end of an interval. The returned results contain only buckets with at least one item so may not return all intervals. For example, if the fixed width for the time to resume the app is 10 ms, then the sequence of buckets is: 0…9 ms, 10…19 ms, 20…29 ms, etc. If there’s data only in the 0…9 ms and 20…29 ms buckets, then the report skips the 10…19 ms bucket.


// An object representing a histogram of data values of the same type of unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXHistogram
type MXHistogram struct {
	objectivec.Object
}

// MXHistogramFrom constructs a [MXHistogram] from an unsafe.Pointer.
//
// An object representing a histogram of data values of the same type of unit.
func MXHistogramFrom(ptr unsafe.Pointer) MXHistogram {
	return MXHistogram{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXHistogram *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXHistogram */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXHistogram */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXHistogram */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXHistogram */

// An enumerator for the buckets containing the data in the histogram.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXHistogram/bucketEnumerator
func (m_ MXHistogram) BucketEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("bucketEnumerator"))
	return rv
}/* debug [instance_properties/getter]: bucketEnumerator */


// The total number of buckets in the histogram.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXHistogram/totalBucketCount
func (m_ MXHistogram) TotalBucketCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("totalBucketCount"))
	return rv
}/* debug [instance_properties/getter]: totalBucketCount */


// Error domain for error values from app metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxerrordomain
func (m_ MXHistogram) MXErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MXErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MXErrorDomain */


// The value of the ending measurement for the bucket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogrambucket/bucketend
func (m_ MXHistogram) BucketEnd() foundation.Measurement {
	rv := objc.Send[foundation.Measurement](m_.ID, objc.Sel("bucketEnd"))
	return rv
}/* debug [instance_properties/getter]: bucketEnd */


// The value of the ending measurement for the bucket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogrambucket/bucketend
func (m_ MXHistogram) SetBucketEnd(value foundation.Measurement) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBucketEnd:"), value)
}/* debug [instance_properties/setter]: bucketEnd */


// The value of the starting measurement for the bucket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogrambucket/bucketstart
func (m_ MXHistogram) BucketStart() foundation.Measurement {
	rv := objc.Send[foundation.Measurement](m_.ID, objc.Sel("bucketStart"))
	return rv
}/* debug [instance_properties/getter]: bucketStart */


// The value of the starting measurement for the bucket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogrambucket/bucketstart
func (m_ MXHistogram) SetBucketStart(value foundation.Measurement) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBucketStart:"), value)
}/* debug [instance_properties/setter]: bucketStart */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXHistogram */



