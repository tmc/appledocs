// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMetricMediaRendition */


/* debug [class_header]: Header for AVMetricMediaRendition */
// The class instance for the [MetricMediaRendition] class.
var (
	MetricMediaRenditionClass     _MetricMediaRenditionClass
	MetricMediaRenditionClassOnce sync.Once
)

func getMetricMediaRenditionClass() _MetricMediaRenditionClass {
	MetricMediaRenditionClassOnce.Do(func() {
		MetricMediaRenditionClass = _MetricMediaRenditionClass{objc.GetClass("AVMetricMediaRendition")}
	})
	return MetricMediaRenditionClass
}

type _MetricMediaRenditionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetricMediaRendition */
// An interface definition for the [MetricMediaRendition] class.
type IMetricMediaRendition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MetricMediaRendition */
	// properties:
	StableID() objc.IObject /* cross-framework: NSString */
	URL() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetricMediaRendition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetricMediaRendition */
// Alloc allocates a new instance without initialization.
func (mc _MetricMediaRenditionClass) Alloc() MetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetricMediaRenditionClass) New() MetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetricMediaRendition) Init() MetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetricMediaRendition) Autorelease() MetricMediaRendition {
	rv := objc.Send[MetricMediaRendition](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetricMediaRendition creates a new MetricMediaRendition instance.
func NewMetricMediaRendition() MetricMediaRendition {
	return getMetricMediaRenditionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetricMediaRendition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaRendition
type MetricMediaRendition struct {
	objectivec.Object
}

// MetricMediaRenditionFrom constructs a [MetricMediaRendition] from an unsafe.Pointer.
func MetricMediaRenditionFrom(ptr unsafe.Pointer) MetricMediaRendition {
	return MetricMediaRendition{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetricMediaRendition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetricMediaRendition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetricMediaRendition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetricMediaRendition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetricMediaRendition */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaRendition/stableID
func (m_ MetricMediaRendition) StableID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("stableID"))
	return rv
}/* debug [instance_properties/getter]: stableID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaRendition/url
func (m_ MetricMediaRendition) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetricMediaRendition */



