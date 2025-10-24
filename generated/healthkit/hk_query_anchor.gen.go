// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKQueryAnchor */


/* debug [class_header]: Header for HKQueryAnchor */
// The class instance for the [HKQueryAnchor] class.
var (
	HKQueryAnchorClass     _HKQueryAnchorClass
	HKQueryAnchorClassOnce sync.Once
)

func getHKQueryAnchorClass() _HKQueryAnchorClass {
	HKQueryAnchorClassOnce.Do(func() {
		HKQueryAnchorClass = _HKQueryAnchorClass{objc.GetClass("HKQueryAnchor")}
	})
	return HKQueryAnchorClass
}

type _HKQueryAnchorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKQueryAnchor */
// An interface definition for the [HKQueryAnchor] class.
type IHKQueryAnchor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKQueryAnchor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKQueryAnchor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKQueryAnchor */
// Alloc allocates a new instance without initialization.
func (hc _HKQueryAnchorClass) Alloc() HKQueryAnchor {
	rv := objc.Send[HKQueryAnchor](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKQueryAnchorClass) New() HKQueryAnchor {
	rv := objc.Send[HKQueryAnchor](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQueryAnchor) Init() HKQueryAnchor {
	rv := objc.Send[HKQueryAnchor](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQueryAnchor) Autorelease() HKQueryAnchor {
	rv := objc.Send[HKQueryAnchor](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQueryAnchor creates a new HKQueryAnchor instance.
func NewHKQueryAnchor() HKQueryAnchor {
	return getHKQueryAnchorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKQueryAnchor */
// An object used to identify all the samples previously returned by an anchored object query.
//
// The system returns objects in both the anchored object query’s results handler and it’s update handler. Use the anchors to query for samples added or deleted after the result or update.


// An object used to identify all the samples previously returned by an anchored object query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQueryAnchor
type HKQueryAnchor struct {
	objectivec.Object
}

// HKQueryAnchorFrom constructs a [HKQueryAnchor] from an unsafe.Pointer.
//
// An object used to identify all the samples previously returned by an anchored object query.
func HKQueryAnchorFrom(ptr unsafe.Pointer) HKQueryAnchor {
	return HKQueryAnchor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKQueryAnchor */

// Returns an anchor object from the provided anchor value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQueryAnchor/init(fromValue:)
func NewHKQueryAnchorFromValue(value uint) HKQueryAnchor {
	rv := objc.Send[HKQueryAnchor](objc.ID(getHKQueryAnchorClass().class), objc.Sel("anchorFromValue:"), value)
	return rv
}/* debug [class_init_methods/constructor]: NewHKQueryAnchorFromValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKQueryAnchor */

// Returns an anchor object from the provided anchor value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQueryAnchor/init(fromValue:)
func (hc _HKQueryAnchorClass) AnchorFromValue(value uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("anchorFromValue:"), value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AnchorFromValue) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKQueryAnchor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKQueryAnchor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKQueryAnchor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKQueryAnchor */


