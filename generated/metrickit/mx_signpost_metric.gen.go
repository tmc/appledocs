// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MXSignpostMetric] class.
var (
	MXSignpostMetricClass     _MXSignpostMetricClass
	MXSignpostMetricClassOnce sync.Once
)

func getMXSignpostMetricClass() _MXSignpostMetricClass {
	MXSignpostMetricClassOnce.Do(func() {
		MXSignpostMetricClass = _MXSignpostMetricClass{objc.GetClass("MXSignpostMetric")}
	})
	return MXSignpostMetricClass
}

type _MXSignpostMetricClass struct {
	class objc.Class
}

// An interface definition for the [MXSignpostMetric] class.
type IMXSignpostMetric interface {
	IMXMetric
	// properties:
	SignpostCategory() objc.IObject /* cross-framework: NSString */
	SignpostIntervalData() IMXSignpostIntervalData
	SignpostName() objc.IObject /* cross-framework: NSString */
	TotalCount() uint
	// methods:
}

// An object representing a custom metric.
//
// A custom metric is an event type with a developer-defined name and category. You can add custom metrics to daily reports to capture information specific to your app. Custom metrics are a type of signpost saved to custom OS logs created using . The daily report contains information about the number and duration of custom events, as well as the power and performance impact of those events. Only custom metric events logged using MetricKit utility functions capture additional power and performance data.


// An object representing a custom metric.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostMetric
type MXSignpostMetric struct {
	MXMetric
}

// MXSignpostMetricFrom constructs a [MXSignpostMetric] from an unsafe.Pointer.
//
// An object representing a custom metric.
func MXSignpostMetricFrom(ptr unsafe.Pointer) MXSignpostMetric {
	return MXSignpostMetric{
		MXMetric: MXMetricFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MXSignpostMetricClass) Alloc() MXSignpostMetric {
	rv := objc.Send[MXSignpostMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXSignpostMetricClass) New() MXSignpostMetric {
	rv := objc.Send[MXSignpostMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXSignpostMetric) Init() MXSignpostMetric {
	rv := objc.Send[MXSignpostMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXSignpostMetric) Autorelease() MXSignpostMetric {
	rv := objc.Send[MXSignpostMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXSignpostMetric creates a new MXSignpostMetric instance.
func NewMXSignpostMetric() MXSignpostMetric {
	return getMXSignpostMetricClass().New()
}



// The developer-specified category of the custom metric represented by the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostMetric/signpostCategory
func (m_ MXSignpostMetric) SignpostCategory() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("signpostCategory"))
	return rv
}


// The data captured for a custom metric.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostMetric/signpostIntervalData
func (m_ MXSignpostMetric) SignpostIntervalData() IMXSignpostIntervalData {
	rv := objc.Send[MXSignpostIntervalData](m_.ID, objc.Sel("signpostIntervalData"))
	return rv
}


// The developer-specified name of the custom metric represented by the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostMetric/signpostName
func (m_ MXSignpostMetric) SignpostName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("signpostName"))
	return rv
}


// The total number of occurrences of the captured custom metric.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXSignpostMetric/totalCount
func (m_ MXSignpostMetric) TotalCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("totalCount"))
	return rv
}



