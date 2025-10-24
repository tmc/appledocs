// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKDocumentSample */


/* debug [class_header]: Header for HKDocumentSample */
// The class instance for the [HKDocumentSample] class.
var (
	HKDocumentSampleClass     _HKDocumentSampleClass
	HKDocumentSampleClassOnce sync.Once
)

func getHKDocumentSampleClass() _HKDocumentSampleClass {
	HKDocumentSampleClassOnce.Do(func() {
		HKDocumentSampleClass = _HKDocumentSampleClass{objc.GetClass("HKDocumentSample")}
	})
	return HKDocumentSampleClass
}

type _HKDocumentSampleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKDocumentSample */
// An interface definition for the [HKDocumentSample] class.
type IHKDocumentSample interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKDocumentSample */
	// properties:
	DocumentType() IHKDocumentType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKDocumentSample */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKDocumentSample */
// Alloc allocates a new instance without initialization.
func (hc _HKDocumentSampleClass) Alloc() HKDocumentSample {
	rv := objc.Send[HKDocumentSample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKDocumentSampleClass) New() HKDocumentSample {
	rv := objc.Send[HKDocumentSample](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKDocumentSample) Init() HKDocumentSample {
	rv := objc.Send[HKDocumentSample](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKDocumentSample) Autorelease() HKDocumentSample {
	rv := objc.Send[HKDocumentSample](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKDocumentSample creates a new HKDocumentSample instance.
func NewHKDocumentSample() HKDocumentSample {
	return getHKDocumentSampleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKDocumentSample */
// An abstract class that represents a health document in the HealthKit store.
//
// You should never instantiate an object directly. Instead, you always work with a concrete subclass. In iOS 10 and watchOS 3, the only concrete class is the class. Document samples are immutable: You set the sample’s properties when you create it, and they cannot change.


// An abstract class that represents a health document in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDocumentSample
type HKDocumentSample struct {
	HKSample
}

// HKDocumentSampleFrom constructs a [HKDocumentSample] from an unsafe.Pointer.
//
// An abstract class that represents a health document in the HealthKit store.
func HKDocumentSampleFrom(ptr unsafe.Pointer) HKDocumentSample {
	return HKDocumentSample{
		HKSample: HKSampleFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKDocumentSample *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKDocumentSample */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKDocumentSample */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKDocumentSample */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKDocumentSample */

// The type of document represented by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDocumentSample/documentType
func (h_ HKDocumentSample) DocumentType() IHKDocumentType {
	rv := objc.Send[HKDocumentType](h_.ID, objc.Sel("documentType"))
	return rv
}/* debug [instance_properties/getter]: documentType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKDocumentSample */



