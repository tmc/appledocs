// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MXDiskSpaceUsageMetric */


/* debug [class_header]: Header for MXDiskSpaceUsageMetric */
// The class instance for the [MXDiskSpaceUsageMetric] class.
var (
	MXDiskSpaceUsageMetricClass     _MXDiskSpaceUsageMetricClass
	MXDiskSpaceUsageMetricClassOnce sync.Once
)

func getMXDiskSpaceUsageMetricClass() _MXDiskSpaceUsageMetricClass {
	MXDiskSpaceUsageMetricClassOnce.Do(func() {
		MXDiskSpaceUsageMetricClass = _MXDiskSpaceUsageMetricClass{objc.GetClass("MXDiskSpaceUsageMetric")}
	})
	return MXDiskSpaceUsageMetricClass
}

type _MXDiskSpaceUsageMetricClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXDiskSpaceUsageMetric */
// An interface definition for the [MXDiskSpaceUsageMetric] class.
type IMXDiskSpaceUsageMetric interface {
	IMXMetric
	
/* debug [class_interface_properties]: Properties for MXDiskSpaceUsageMetric */
	// properties:
	TotalBinaryFileCount() int
	TotalBinaryFileSize() unsafe.Pointer
	TotalCacheFolderSize() unsafe.Pointer
	TotalCloneSize() unsafe.Pointer
	TotalDataFileCount() int
	TotalDataFileSize() unsafe.Pointer
	TotalDiskSpaceCapacity() unsafe.Pointer
	TotalDiskSpaceUsedSize() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXDiskSpaceUsageMetric */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXDiskSpaceUsageMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXDiskSpaceUsageMetricClass) Alloc() MXDiskSpaceUsageMetric {
	rv := objc.Send[MXDiskSpaceUsageMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MXDiskSpaceUsageMetricClass) New() MXDiskSpaceUsageMetric {
	rv := objc.Send[MXDiskSpaceUsageMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXDiskSpaceUsageMetric) Init() MXDiskSpaceUsageMetric {
	rv := objc.Send[MXDiskSpaceUsageMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXDiskSpaceUsageMetric) Autorelease() MXDiskSpaceUsageMetric {
	rv := objc.Send[MXDiskSpaceUsageMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXDiskSpaceUsageMetric creates a new MXDiskSpaceUsageMetric instance.
func NewMXDiskSpaceUsageMetric() MXDiskSpaceUsageMetric {
	return getMXDiskSpaceUsageMetricClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXDiskSpaceUsageMetric */
// An object representing metrics about your app’s disk space usage.


// An object representing metrics about your app’s disk space usage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskSpaceUsageMetric
type MXDiskSpaceUsageMetric struct {
	MXMetric
}

// MXDiskSpaceUsageMetricFrom constructs a [MXDiskSpaceUsageMetric] from an unsafe.Pointer.
//
// An object representing metrics about your app’s disk space usage.
func MXDiskSpaceUsageMetricFrom(ptr unsafe.Pointer) MXDiskSpaceUsageMetric {
	return MXDiskSpaceUsageMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXDiskSpaceUsageMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXDiskSpaceUsageMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXDiskSpaceUsageMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXDiskSpaceUsageMetric */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXDiskSpaceUsageMetric */

// The total number of your app’s binary files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskSpaceUsageMetric/totalBinaryFileCount
func (m_ MXDiskSpaceUsageMetric) TotalBinaryFileCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("totalBinaryFileCount"))
	return rv
}/* debug [instance_properties/getter]: totalBinaryFileCount */


// The total size of disk space your app’s binary files occupy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskSpaceUsageMetric/totalBinaryFileSize
func (m_ MXDiskSpaceUsageMetric) TotalBinaryFileSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalBinaryFileSize"))
	return rv
}/* debug [instance_properties/getter]: totalBinaryFileSize */


// The total size of your application’s cache folder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskSpaceUsageMetric/totalCacheFolderSize
func (m_ MXDiskSpaceUsageMetric) TotalCacheFolderSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalCacheFolderSize"))
	return rv
}/* debug [instance_properties/getter]: totalCacheFolderSize */


// The total size of all clone files that are attributed to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskSpaceUsageMetric/totalCloneSize
func (m_ MXDiskSpaceUsageMetric) TotalCloneSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalCloneSize"))
	return rv
}/* debug [instance_properties/getter]: totalCloneSize */


// The total number of data files in your app’s container(s).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskSpaceUsageMetric/totalDataFileCount
func (m_ MXDiskSpaceUsageMetric) TotalDataFileCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("totalDataFileCount"))
	return rv
}/* debug [instance_properties/getter]: totalDataFileCount */


// The total size of disk space your app uses for storing data files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskSpaceUsageMetric/totalDataFileSize
func (m_ MXDiskSpaceUsageMetric) TotalDataFileSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalDataFileSize"))
	return rv
}/* debug [instance_properties/getter]: totalDataFileSize */


// The total disk space capacity of the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskSpaceUsageMetric/totalDiskSpaceCapacity
func (m_ MXDiskSpaceUsageMetric) TotalDiskSpaceCapacity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalDiskSpaceCapacity"))
	return rv
}/* debug [instance_properties/getter]: totalDiskSpaceCapacity */


// The total amount of used disk storage on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXDiskSpaceUsageMetric/totalDiskSpaceUsedSize
func (m_ MXDiskSpaceUsageMetric) TotalDiskSpaceUsedSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalDiskSpaceUsedSize"))
	return rv
}/* debug [instance_properties/getter]: totalDiskSpaceUsedSize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXDiskSpaceUsageMetric */



