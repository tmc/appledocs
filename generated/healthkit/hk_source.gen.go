// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKSource */


/* debug [class_header]: Header for HKSource */
// The class instance for the [HKSource] class.
var (
	HKSourceClass     _HKSourceClass
	HKSourceClassOnce sync.Once
)

func getHKSourceClass() _HKSourceClass {
	HKSourceClassOnce.Do(func() {
		HKSourceClass = _HKSourceClass{objc.GetClass("HKSource")}
	})
	return HKSourceClass
}

type _HKSourceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKSource */
// An interface definition for the [HKSource] class.
type IHKSource interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKSource */
	// properties:
	BundleIdentifier() objc.IObject /* cross-framework: NSString */
	Name() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKSource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKSource */
// Alloc allocates a new instance without initialization.
func (hc _HKSourceClass) Alloc() HKSource {
	rv := objc.Send[HKSource](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKSourceClass) New() HKSource {
	rv := objc.Send[HKSource](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKSource) Init() HKSource {
	rv := objc.Send[HKSource](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKSource) Autorelease() HKSource {
	rv := objc.Send[HKSource](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKSource creates a new HKSource instance.
func NewHKSource() HKSource {
	return getHKSourceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKSource */
// An object indicating the app or device that created a HealthKit sample
//
// Sources include apps and devices that save data to the HealthKit store. Currently, HealthKit supports only the direct import of data from Bluetooth LE heart rate monitors. All other devices need a companion app to collect and save the data to HealthKit.


// An object indicating the app or device that created a HealthKit sample
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSource
type HKSource struct {
	objectivec.Object
}

// HKSourceFrom constructs a [HKSource] from an unsafe.Pointer.
//
// An object indicating the app or device that created a HealthKit sample
func HKSourceFrom(ptr unsafe.Pointer) HKSource {
	return HKSource{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKSource *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKSource */

// Returns a source object for the current app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSource/default()
func (hc _HKSourceClass) DefaultSource() HKSource {
	rv := objc.Send[HKSource](objc.ID(hc.class), objc.Sel("defaultSource"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultSource) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKSource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKSource */

// The source’s bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSource/bundleIdentifier
func (h_ HKSource) BundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("bundleIdentifier"))
	return rv
}/* debug [instance_properties/getter]: bundleIdentifier */


// The source’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSource/name
func (h_ HKSource) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKSource */



