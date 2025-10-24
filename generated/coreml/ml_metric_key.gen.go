// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLMetricKey */


/* debug [class_header]: Header for MLMetricKey */
// The class instance for the [MetricKey] class.
var (
	MetricKeyClass     _MetricKeyClass
	MetricKeyClassOnce sync.Once
)

func getMetricKeyClass() _MetricKeyClass {
	MetricKeyClassOnce.Do(func() {
		MetricKeyClass = _MetricKeyClass{objc.GetClass("MLMetricKey")}
	})
	return MetricKeyClass
}

type _MetricKeyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricKey */
// An interface definition for the [MetricKey] class.
type IMetricKey interface {
	IKey
	
/* debug [class_interface_properties]: Properties for MetricKey */
	// properties:
	Metrics() IMLMetricKey
	SetMetrics(value IMLMetricKey)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricKey */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricKey */
// Alloc allocates a new instance without initialization.
func (mc _MetricKeyClass) Alloc() MetricKey {
	rv := objc.Send[MetricKey](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricKeyClass) New() MetricKey {
	rv := objc.Send[MetricKey](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricKey) Init() MetricKey {
	rv := objc.Send[MetricKey](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricKey) Autorelease() MetricKey {
	rv := objc.Send[MetricKey](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricKey creates a new MetricKey instance.
func NewMetricKey() MetricKey {
	return getMetricKeyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricKey */
// A key for the metrics dictionary in an update context.


// A key for the metrics dictionary in an update context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMetricKey
type MetricKey struct {
	Key
}

// MetricKeyFrom constructs a [MetricKey] from an unsafe.Pointer.
//
// A key for the metrics dictionary in an update context.
func MetricKeyFrom(ptr unsafe.Pointer) MetricKey {
	return MetricKey{
		Key: KeyFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricKey *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricKey */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricKey */

// The key you use to access the epoch index (an value).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMetricKey/epochIndex
func (mc _MetricKeyClass) EpochIndex() MetricKey {
	rv := objc.Send[MetricKey](objc.ID(mc.class), objc.Sel("epochIndex"))
	return rv
}/* debug [class_properties_class/property]: epochIndex */

// The key you use to access the current loss (a value).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMetricKey/lossValue
func (mc _MetricKeyClass) LossValue() MetricKey {
	rv := objc.Send[MetricKey](objc.ID(mc.class), objc.Sel("lossValue"))
	return rv
}/* debug [class_properties_class/property]: lossValue */

// The key you use to access the mini-batch index (an value) within an epoch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMetricKey/miniBatchIndex
func (mc _MetricKeyClass) MiniBatchIndex() MetricKey {
	rv := objc.Send[MetricKey](objc.ID(mc.class), objc.Sel("miniBatchIndex"))
	return rv
}/* debug [class_properties_class/property]: miniBatchIndex */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricKey */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricKey */

// The key you use to access the epoch index (an value).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMetricKey/epochIndex
func (m_ MetricKey) EpochIndex() IMLMetricKey {
	rv := objc.Send[MetricKey](m_.ID, objc.Sel("epochIndex"))
	return rv
}/* debug [instance_properties/getter]: epochIndex */


// The key you use to access the current loss (a value).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMetricKey/lossValue
func (m_ MetricKey) LossValue() IMLMetricKey {
	rv := objc.Send[MetricKey](m_.ID, objc.Sel("lossValue"))
	return rv
}/* debug [instance_properties/getter]: lossValue */


// The key you use to access the mini-batch index (an value) within an epoch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMetricKey/miniBatchIndex
func (m_ MetricKey) MiniBatchIndex() IMLMetricKey {
	rv := objc.Send[MetricKey](m_.ID, objc.Sel("miniBatchIndex"))
	return rv
}/* debug [instance_properties/getter]: miniBatchIndex */


// The training metrics of the model for the update task, contained in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/metrics
func (m_ MetricKey) Metrics() IMLMetricKey {
	rv := objc.Send[MetricKey](m_.ID, objc.Sel("metrics"))
	return rv
}/* debug [instance_properties/getter]: metrics */


// The training metrics of the model for the update task, contained in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/metrics
func (m_ MetricKey) SetMetrics(value IMLMetricKey) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetrics:"), value)
}/* debug [instance_properties/setter]: metrics */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLMetricKey */



