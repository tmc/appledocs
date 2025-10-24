// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKAudiogramSampleType */


/* debug [class_header]: Header for HKAudiogramSampleType */
// The class instance for the [HKAudiogramSampleType] class.
var (
	HKAudiogramSampleTypeClass     _HKAudiogramSampleTypeClass
	HKAudiogramSampleTypeClassOnce sync.Once
)

func getHKAudiogramSampleTypeClass() _HKAudiogramSampleTypeClass {
	HKAudiogramSampleTypeClassOnce.Do(func() {
		HKAudiogramSampleTypeClass = _HKAudiogramSampleTypeClass{objc.GetClass("HKAudiogramSampleType")}
	})
	return HKAudiogramSampleTypeClass
}

type _HKAudiogramSampleTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKAudiogramSampleType */
// An interface definition for the [HKAudiogramSampleType] class.
type IHKAudiogramSampleType interface {
	IHKSampleType
	
/* debug [class_interface_properties]: Properties for HKAudiogramSampleType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKAudiogramSampleType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKAudiogramSampleType */
// Alloc allocates a new instance without initialization.
func (hc _HKAudiogramSampleTypeClass) Alloc() HKAudiogramSampleType {
	rv := objc.Send[HKAudiogramSampleType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKAudiogramSampleTypeClass) New() HKAudiogramSampleType {
	rv := objc.Send[HKAudiogramSampleType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAudiogramSampleType) Init() HKAudiogramSampleType {
	rv := objc.Send[HKAudiogramSampleType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAudiogramSampleType) Autorelease() HKAudiogramSampleType {
	rv := objc.Send[HKAudiogramSampleType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAudiogramSampleType creates a new HKAudiogramSampleType instance.
func NewHKAudiogramSampleType() HKAudiogramSampleType {
	return getHKAudiogramSampleTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKAudiogramSampleType */
// A type that identifies samples that contain audiogram data.
//
// The class is a concrete subclass of the class. To create an audiogram sample type instance, use the object type’s convenience method. Use audiogram sample types to: Request permission to read or write audiogram samples. Create and share audiogram samples. Query for audiogram samples.


// A type that identifies samples that contain audiogram data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAudiogramSampleType
type HKAudiogramSampleType struct {
	HKSampleType
}

// HKAudiogramSampleTypeFrom constructs a [HKAudiogramSampleType] from an unsafe.Pointer.
//
// A type that identifies samples that contain audiogram data.
func HKAudiogramSampleTypeFrom(ptr unsafe.Pointer) HKAudiogramSampleType {
	return HKAudiogramSampleType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKAudiogramSampleType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKAudiogramSampleType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKAudiogramSampleType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKAudiogramSampleType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKAudiogramSampleType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKAudiogramSampleType */



