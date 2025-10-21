// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MXHistogram] class.
type IMXHistogram interface {
	objectivec.IObject
}

// An object representing a histogram of data values of the same type of unit.
//
// A measures the number of times a data point for a variable falls into a specific range of possible values within a set of data. Usually, histograms are depicted as bar charts, in which each bar represents a range of values, and the height of each bar represents the number of times the value of the variable falls within a particular range. In this class, each bar is represented by a . A bucket holds the results for a series of measured values, such as all the events occurring between 3 and 5 seconds. MetricKit uses fixed-width buckets that are device-independent with intervals that are based on the type of metric. Use the and properties to find the start and end of an interval. The returned results contain only buckets with at least one item so may not return all intervals. For example, if the fixed width for the time to resume the app is 10 ms, then the sequence of buckets is: 0…9 ms, 10…19 ms, 20…29 ms, etc. If there’s data only in the 0…9 ms and 20…29 ms buckets, then the report skips the 10…19 ms bucket.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MXHistogramClass) Alloc() MXHistogram {
	rv := objc.Send[MXHistogram](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// An enumerator for the buckets containing the data in the histogram.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXHistogram/bucketEnumerator
func (m_ MXHistogram) BucketEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("bucketEnumerator"))
	return rv
}

// The total number of buckets in the histogram.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXHistogram/totalBucketCount
func (m_ MXHistogram) TotalBucketCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("totalBucketCount"))
	return rv
}

// Error domain for error values from app metrics.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxerrordomain
func (m_ MXHistogram) MXErrorDomain() string {
	rv := objc.Send[string](m_.ID, objc.Sel("MXErrorDomain"))
	return rv
}

// The value of the ending measurement for the bucket.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogrambucket/bucketend
func (m_ MXHistogram) BucketEnd() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("bucketEnd"))
	return rv
}


// SetBucketEnd sets the value of the bucketEnd property.
// The value of the ending measurement for the bucket.

//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogrambucket/bucketend
func (m_ MXHistogram) SetBucketEnd(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBucketEnd:"), value)
}

// The value of the starting measurement for the bucket.
//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogrambucket/bucketstart
func (m_ MXHistogram) BucketStart() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("bucketStart"))
	return rv
}


// SetBucketStart sets the value of the bucketStart property.
// The value of the starting measurement for the bucket.

//
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogrambucket/bucketstart
func (m_ MXHistogram) SetBucketStart(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBucketStart:"), value)
}



