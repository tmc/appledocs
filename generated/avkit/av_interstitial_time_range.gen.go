// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVInterstitialTimeRange */


/* debug [class_header]: Header for AVInterstitialTimeRange */
// The class instance for the [InterstitialTimeRange] class.
var (
	InterstitialTimeRangeClass     _InterstitialTimeRangeClass
	InterstitialTimeRangeClassOnce sync.Once
)

func getInterstitialTimeRangeClass() _InterstitialTimeRangeClass {
	InterstitialTimeRangeClassOnce.Do(func() {
		InterstitialTimeRangeClass = _InterstitialTimeRangeClass{objc.GetClass("AVInterstitialTimeRange")}
	})
	return InterstitialTimeRangeClass
}

type _InterstitialTimeRangeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InterstitialTimeRange */
// An interface definition for the [InterstitialTimeRange] class.
type IInterstitialTimeRange interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for InterstitialTimeRange */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InterstitialTimeRange */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InterstitialTimeRange */
// Alloc allocates a new instance without initialization.
func (ic _InterstitialTimeRangeClass) Alloc() InterstitialTimeRange {
	rv := objc.Send[InterstitialTimeRange](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _InterstitialTimeRangeClass) New() InterstitialTimeRange {
	rv := objc.Send[InterstitialTimeRange](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InterstitialTimeRange) Init() InterstitialTimeRange {
	rv := objc.Send[InterstitialTimeRange](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InterstitialTimeRange) Autorelease() InterstitialTimeRange {
	rv := objc.Send[InterstitialTimeRange](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInterstitialTimeRange creates a new InterstitialTimeRange instance.
func NewInterstitialTimeRange() InterstitialTimeRange {
	return getInterstitialTimeRangeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InterstitialTimeRange */
// A time range in an audiovisual presentation for content with an interstitial designation, such as advertisements or legal notices.
//
// When you associate interstitial time ranges with an you present with an , you can customize or restrict the presentation of interstitial content. For example, you can allow the user to skip advertisements or prohibit skipping of a legal notice.


// A time range in an audiovisual presentation for content with an interstitial designation, such as advertisements or legal notices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInterstitialTimeRange
type InterstitialTimeRange struct {
	objectivec.Object
}

// InterstitialTimeRangeFrom constructs a [InterstitialTimeRange] from an unsafe.Pointer.
//
// A time range in an audiovisual presentation for content with an interstitial designation, such as advertisements or legal notices.
func InterstitialTimeRangeFrom(ptr unsafe.Pointer) InterstitialTimeRange {
	return InterstitialTimeRange{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InterstitialTimeRange */

// Initializes an interstitial time range object with the specified time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInterstitialTimeRange/init(timeRange:)
func NewInterstitialTimeRangeWithTimeRange(timeRange TimeRange /* not a class type */) InterstitialTimeRange {
	instance := getInterstitialTimeRangeClass().Alloc()
	rv := objc.Send[InterstitialTimeRange](instance.ID, objc.Sel("initWithTimeRange:"), timeRange)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewInterstitialTimeRangeWithTimeRange */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InterstitialTimeRange */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InterstitialTimeRange */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InterstitialTimeRange */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InterstitialTimeRange */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVInterstitialTimeRange */


