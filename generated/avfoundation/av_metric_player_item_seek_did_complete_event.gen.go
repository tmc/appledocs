// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetricPlayerItemSeekDidCompleteEvent */


/* debug [class_header]: Header for AVMetricPlayerItemSeekDidCompleteEvent */
// The class instance for the [MetricPlayerItemSeekDidCompleteEvent] class.
var (
	MetricPlayerItemSeekDidCompleteEventClass     _MetricPlayerItemSeekDidCompleteEventClass
	MetricPlayerItemSeekDidCompleteEventClassOnce sync.Once
)

func getMetricPlayerItemSeekDidCompleteEventClass() _MetricPlayerItemSeekDidCompleteEventClass {
	MetricPlayerItemSeekDidCompleteEventClassOnce.Do(func() {
		MetricPlayerItemSeekDidCompleteEventClass = _MetricPlayerItemSeekDidCompleteEventClass{objc.GetClass("AVMetricPlayerItemSeekDidCompleteEvent")}
	})
	return MetricPlayerItemSeekDidCompleteEventClass
}

type _MetricPlayerItemSeekDidCompleteEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricPlayerItemSeekDidCompleteEvent */
// An interface definition for the [MetricPlayerItemSeekDidCompleteEvent] class.
type IMetricPlayerItemSeekDidCompleteEvent interface {
	IMetricPlayerItemRateChangeEvent
	
/* debug [class_interface_properties]: Properties for MetricPlayerItemSeekDidCompleteEvent */
	// properties:
	DidSeekInBuffer() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricPlayerItemSeekDidCompleteEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricPlayerItemSeekDidCompleteEvent */
// Alloc allocates a new instance without initialization.
func (mc _MetricPlayerItemSeekDidCompleteEventClass) Alloc() MetricPlayerItemSeekDidCompleteEvent {
	rv := objc.Send[MetricPlayerItemSeekDidCompleteEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricPlayerItemSeekDidCompleteEventClass) New() MetricPlayerItemSeekDidCompleteEvent {
	rv := objc.Send[MetricPlayerItemSeekDidCompleteEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricPlayerItemSeekDidCompleteEvent) Init() MetricPlayerItemSeekDidCompleteEvent {
	rv := objc.Send[MetricPlayerItemSeekDidCompleteEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricPlayerItemSeekDidCompleteEvent) Autorelease() MetricPlayerItemSeekDidCompleteEvent {
	rv := objc.Send[MetricPlayerItemSeekDidCompleteEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricPlayerItemSeekDidCompleteEvent creates a new MetricPlayerItemSeekDidCompleteEvent instance.
func NewMetricPlayerItemSeekDidCompleteEvent() MetricPlayerItemSeekDidCompleteEvent {
	return getMetricPlayerItemSeekDidCompleteEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricPlayerItemSeekDidCompleteEvent */
// An event that represents when the playback seek completes.


// An event that represents when the playback seek completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemSeekDidCompleteEvent
type MetricPlayerItemSeekDidCompleteEvent struct {
	MetricPlayerItemRateChangeEvent
}

// MetricPlayerItemSeekDidCompleteEventFrom constructs a [MetricPlayerItemSeekDidCompleteEvent] from an unsafe.Pointer.
//
// An event that represents when the playback seek completes.
func MetricPlayerItemSeekDidCompleteEventFrom(ptr unsafe.Pointer) MetricPlayerItemSeekDidCompleteEvent {
	return MetricPlayerItemSeekDidCompleteEvent{
		MetricPlayerItemRateChangeEvent: MetricPlayerItemRateChangeEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricPlayerItemSeekDidCompleteEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricPlayerItemSeekDidCompleteEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricPlayerItemSeekDidCompleteEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricPlayerItemSeekDidCompleteEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricPlayerItemSeekDidCompleteEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemSeekDidCompleteEvent/didSeekInBuffer
func (m_ MetricPlayerItemSeekDidCompleteEvent) DidSeekInBuffer() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("didSeekInBuffer"))
	return rv
}/* debug [instance_properties/getter]: didSeekInBuffer */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricPlayerItemSeekDidCompleteEvent */



