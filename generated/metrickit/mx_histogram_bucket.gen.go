// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MXHistogramBucket] class.
var (
	MXHistogramBucketClass     _MXHistogramBucketClass
	MXHistogramBucketClassOnce sync.Once
)

func getMXHistogramBucketClass() _MXHistogramBucketClass {
	MXHistogramBucketClassOnce.Do(func() {
		MXHistogramBucketClass = _MXHistogramBucketClass{objc.GetClass("MXHistogramBucket")}
	})
	return MXHistogramBucketClass
}

type _MXHistogramBucketClass struct {
	class objc.Class
}

// An interface definition for the [MXHistogramBucket] class.
type IMXHistogramBucket interface {
	objectivec.IObject
	// properties:
	BucketEnd() unsafe.Pointer
	BucketStart() unsafe.Pointer
	BucketEnumerator() objc.IObject /* cross-framework: Enumerator */
	SetBucketEnumerator(value objc.IObject /* cross-framework: Enumerator */)
	TotalBucketCount() int
	SetTotalBucketCount(value int)
	BucketCount() int
	SetBucketCount(value int)
	// methods:
}

// An object representing a bucket of data in a histogram.


// An object representing a bucket of data in a histogram.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXHistogramBucket
type MXHistogramBucket struct {
	objectivec.Object
}

// MXHistogramBucketFrom constructs a [MXHistogramBucket] from an unsafe.Pointer.
//
// An object representing a bucket of data in a histogram.
func MXHistogramBucketFrom(ptr unsafe.Pointer) MXHistogramBucket {
	return MXHistogramBucket{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MXHistogramBucketClass) Alloc() MXHistogramBucket {
	rv := objc.Send[MXHistogramBucket](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXHistogramBucketClass) New() MXHistogramBucket {
	rv := objc.Send[MXHistogramBucket](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXHistogramBucket) Init() MXHistogramBucket {
	rv := objc.Send[MXHistogramBucket](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXHistogramBucket) Autorelease() MXHistogramBucket {
	rv := objc.Send[MXHistogramBucket](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXHistogramBucket creates a new MXHistogramBucket instance.
func NewMXHistogramBucket() MXHistogramBucket {
	return getMXHistogramBucketClass().New()
}



// The value of the ending measurement for the bucket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXHistogramBucket/bucketEnd
func (m_ MXHistogramBucket) BucketEnd() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("bucketEnd"))
	return rv
}


// The value of the starting measurement for the bucket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXHistogramBucket/bucketStart
func (m_ MXHistogramBucket) BucketStart() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("bucketStart"))
	return rv
}


// An enumerator for the buckets containing the data in the histogram.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogram/bucketenumerator
func (m_ MXHistogramBucket) BucketEnumerator() objc.IObject /* cross-framework: Enumerator */ {
	rv := objc.Send[foundation.Enumerator](m_.ID, objc.Sel("bucketEnumerator"))
	return rv
}


// An enumerator for the buckets containing the data in the histogram.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogram/bucketenumerator
func (m_ MXHistogramBucket) SetBucketEnumerator(value objc.IObject /* cross-framework: Enumerator */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBucketEnumerator:"), value)
}


// The total number of buckets in the histogram.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogram/totalbucketcount
func (m_ MXHistogramBucket) TotalBucketCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("totalBucketCount"))
	return rv
}


// The total number of buckets in the histogram.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogram/totalbucketcount
func (m_ MXHistogramBucket) SetTotalBucketCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTotalBucketCount:"), value)
}


// An integer representing the number of samples in the bucket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogrambucket/bucketcount
func (m_ MXHistogramBucket) BucketCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("bucketCount"))
	return rv
}


// An integer representing the number of samples in the bucket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxhistogrambucket/bucketcount
func (m_ MXHistogramBucket) SetBucketCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBucketCount:"), value)
}



