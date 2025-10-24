// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MXMetric */


/* debug [class_header]: Header for MXMetric */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MXMetric */
// An interface definition for the [MXMetric] class.
type IMXMetric interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MXMetric */
	// properties:
	MXErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MXMetric */
	// methods:
	DictionaryRepresentation() foundation.Dictionary
	JSONRepresentation() foundation.Data
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MXMetric */
// Alloc allocates a new instance without initialization.
func (mc _MXMetricClass) Alloc() MXMetric {
	rv := objc.Send[MXMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MXMetric */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MXMetric *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MXMetric */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MXMetric */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MXMetric */

// Returns the contents of a metric as a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetric/dictionaryRepresentation()
func (m_ MXMetric) DictionaryRepresentation() foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](m_.ID, objc.Sel("dictionaryRepresentation"))
	return rv
}/* debug [instance_methods/method]: DictionaryRepresentation */


// Returns the contents of the metric in JSON format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXMetric/jsonRepresentation()
func (m_ MXMetric) JSONRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}/* debug [instance_methods/method]: JSONRepresentation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MXMetric */

// Error domain for error values from app metrics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metrickit/mxerrordomain
func (m_ MXMetric) MXErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MXErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MXErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MXMetric */


