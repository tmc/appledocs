// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKElectrocardiogramType */


/* debug [class_header]: Header for HKElectrocardiogramType */
// The class instance for the [HKElectrocardiogramType] class.
var (
	HKElectrocardiogramTypeClass     _HKElectrocardiogramTypeClass
	HKElectrocardiogramTypeClassOnce sync.Once
)

func getHKElectrocardiogramTypeClass() _HKElectrocardiogramTypeClass {
	HKElectrocardiogramTypeClassOnce.Do(func() {
		HKElectrocardiogramTypeClass = _HKElectrocardiogramTypeClass{objc.GetClass("HKElectrocardiogramType")}
	})
	return HKElectrocardiogramTypeClass
}

type _HKElectrocardiogramTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKElectrocardiogramType */
// An interface definition for the [HKElectrocardiogramType] class.
type IHKElectrocardiogramType interface {
	IHKSampleType
	
/* debug [class_interface_properties]: Properties for HKElectrocardiogramType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKElectrocardiogramType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKElectrocardiogramType */
// Alloc allocates a new instance without initialization.
func (hc _HKElectrocardiogramTypeClass) Alloc() HKElectrocardiogramType {
	rv := objc.Send[HKElectrocardiogramType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKElectrocardiogramTypeClass) New() HKElectrocardiogramType {
	rv := objc.Send[HKElectrocardiogramType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKElectrocardiogramType) Init() HKElectrocardiogramType {
	rv := objc.Send[HKElectrocardiogramType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKElectrocardiogramType) Autorelease() HKElectrocardiogramType {
	rv := objc.Send[HKElectrocardiogramType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKElectrocardiogramType creates a new HKElectrocardiogramType instance.
func NewHKElectrocardiogramType() HKElectrocardiogramType {
	return getHKElectrocardiogramTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKElectrocardiogramType */
// A type that identifies samples containing electrocardiogram data.
//
// The class is a concrete subclass of the class. To create an electrocardiogram type instance, use the object type’s convenience method. Use the electrocardiogram type to: Request permission to read electrocardiogram samples Query for electrocardiogram samples Electrocardiogram samples are read-only. You can request permission to read the samples using this identifier, but you can’t request authorization to share them. This means you can’t save new electrocardiogram samples to the HealthKit store. To add test data in iOS Simulator, open the Health app and select Browse > Heart > Electrocardiograms (ECG) > Add Data.


// A type that identifies samples containing electrocardiogram data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKElectrocardiogramType
type HKElectrocardiogramType struct {
	HKSampleType
}

// HKElectrocardiogramTypeFrom constructs a [HKElectrocardiogramType] from an unsafe.Pointer.
//
// A type that identifies samples containing electrocardiogram data.
func HKElectrocardiogramTypeFrom(ptr unsafe.Pointer) HKElectrocardiogramType {
	return HKElectrocardiogramType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKElectrocardiogramType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKElectrocardiogramType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKElectrocardiogramType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKElectrocardiogramType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKElectrocardiogramType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKElectrocardiogramType */



