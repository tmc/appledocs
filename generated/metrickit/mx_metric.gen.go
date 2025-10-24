// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MXMetric] class.
var (
	MXMetricClass     _MXMetricClass
	MXMetricClassOnce sync.Once
)

func getMXMetricClass() _MXMetricClass {
	MXMetricClassOnce.Do(func() {
		MXMetricClass = _MXMetricClass{objc.GetClass("MXMetric")}
	})
	return MXMetricClass
}

type _MXMetricClass struct {
	class objc.Class
}

// An interface definition for the [MXMetric] class.
type IMXMetric interface {
	objectivec.IObject
	// properties:
	MXErrorDomain() objc.IObject /* cross-framework: NSString */
	// methods:
	JSONRepresentation() objc.IObject /* cross-framework: Data */
}

// An abstract data class for a metric.


// An abstract data class for a metric.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetric
type MXMetric struct {
	objectivec.Object
}

// MXMetricFrom constructs a [MXMetric] from an unsafe.Pointer.
//
// An abstract data class for a metric.
func MXMetricFrom(ptr unsafe.Pointer) MXMetric {
	return MXMetric{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MXMetricClass) Alloc() MXMetric {
	rv := objc.Send[MXMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXMetricClass) New() MXMetric {
	rv := objc.Send[MXMetric](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXMetric) Init() MXMetric {
	rv := objc.Send[MXMetric](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXMetric) Autorelease() MXMetric {
	rv := objc.Send[MXMetric](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXMetric creates a new MXMetric instance.
func NewMXMetric() MXMetric {
	return getMXMetricClass().New()
}



// Returns the contents of the metric in JSON format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetric/jsonRepresentation()
func (m_ MXMetric) JSONRepresentation() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}


// Error domain for error values from app metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxerrordomain
func (m_ MXMetric) MXErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MXErrorDomain"))
	return rv
}


