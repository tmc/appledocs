// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [MetricKey] class.
type IMetricKey interface {
	IKey
}

// A key for the metrics dictionary in an update context.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MetricKeyClass) Alloc() MetricKey {
	rv := objc.Send[MetricKey](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The key you use to access the epoch index (an value).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMetricKey/epochIndex
func (mc _MetricKeyClass) EpochIndex() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("epochIndex"))
	return rv
}
// The key you use to access the current loss (a value).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMetricKey/lossValue
func (mc _MetricKeyClass) LossValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("lossValue"))
	return rv
}
// The key you use to access the mini-batch index (an value) within an epoch.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMetricKey/miniBatchIndex
func (mc _MetricKeyClass) MiniBatchIndex() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("miniBatchIndex"))
	return rv
}
// The training metrics of the model for the update task, contained in a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/metrics
func (m_ MetricKey) Metrics() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("metrics"))
	return rv
}


// SetMetrics sets the value of the metrics property.
// The training metrics of the model for the update task, contained in a dictionary.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlupdatecontext/metrics
func (m_ MetricKey) SetMetrics(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetrics:"), value)
}

// The key you use to access the epoch index (an value).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMetricKey/epochIndex
func (m_ MetricKey) EpochIndex() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("epochIndex"))
	return rv
}

// The key you use to access the current loss (a value).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMetricKey/lossValue
func (m_ MetricKey) LossValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("lossValue"))
	return rv
}

// The key you use to access the mini-batch index (an value) within an epoch.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLMetricKey/miniBatchIndex
func (m_ MetricKey) MiniBatchIndex() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("miniBatchIndex"))
	return rv
}



