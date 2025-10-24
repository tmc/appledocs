// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKQueryDescriptor */


/* debug [class_header]: Header for HKQueryDescriptor */
// The class instance for the [HKQueryDescriptor] class.
var (
	HKQueryDescriptorClass     _HKQueryDescriptorClass
	HKQueryDescriptorClassOnce sync.Once
)

func getHKQueryDescriptorClass() _HKQueryDescriptorClass {
	HKQueryDescriptorClassOnce.Do(func() {
		HKQueryDescriptorClass = _HKQueryDescriptorClass{objc.GetClass("HKQueryDescriptor")}
	})
	return HKQueryDescriptorClass
}

type _HKQueryDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKQueryDescriptor */
// An interface definition for the [HKQueryDescriptor] class.
type IHKQueryDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKQueryDescriptor */
	// properties:
	Predicate() foundation.Predicate
	SampleType() IHKSampleType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKQueryDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKQueryDescriptor */
// Alloc allocates a new instance without initialization.
func (hc _HKQueryDescriptorClass) Alloc() HKQueryDescriptor {
	rv := objc.Send[HKQueryDescriptor](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKQueryDescriptorClass) New() HKQueryDescriptor {
	rv := objc.Send[HKQueryDescriptor](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQueryDescriptor) Init() HKQueryDescriptor {
	rv := objc.Send[HKQueryDescriptor](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQueryDescriptor) Autorelease() HKQueryDescriptor {
	rv := objc.Send[HKQueryDescriptor](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQueryDescriptor creates a new HKQueryDescriptor instance.
func NewHKQueryDescriptor() HKQueryDescriptor {
	return getHKQueryDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKQueryDescriptor */
// A descriptor that specifies a set of samples based on the data type and a predicate.
//
// Use descriptors to create queries that return multiple data types. You can use descriptors when creating , , or instances.


// A descriptor that specifies a set of samples based on the data type and a predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQueryDescriptor
type HKQueryDescriptor struct {
	objectivec.Object
}

// HKQueryDescriptorFrom constructs a [HKQueryDescriptor] from an unsafe.Pointer.
//
// A descriptor that specifies a set of samples based on the data type and a predicate.
func HKQueryDescriptorFrom(ptr unsafe.Pointer) HKQueryDescriptor {
	return HKQueryDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKQueryDescriptor */

// Creates a new descriptor for the data type and predicate you provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQueryDescriptor/init(sampleType:predicate:)
func NewHKQueryDescriptorWithSampleTypePredicate(sampleType IHKSampleType, predicate foundation.Predicate) HKQueryDescriptor {
	instance := getHKQueryDescriptorClass().Alloc()
	rv := objc.Send[HKQueryDescriptor](instance.ID, objc.Sel("initWithSampleType:predicate:"), sampleType, predicate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKQueryDescriptorWithSampleTypePredicate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKQueryDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKQueryDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKQueryDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKQueryDescriptor */

// The predicate that filters samples matching this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQueryDescriptor/predicate
func (h_ HKQueryDescriptor) Predicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](h_.ID, objc.Sel("predicate"))
	return rv
}/* debug [instance_properties/getter]: predicate */


// The data type of samples that match this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQueryDescriptor/sampleType
func (h_ HKQueryDescriptor) SampleType() IHKSampleType {
	rv := objc.Send[HKSampleType](h_.ID, objc.Sel("sampleType"))
	return rv
}/* debug [instance_properties/getter]: sampleType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKQueryDescriptor */


