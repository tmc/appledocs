// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKDocumentType */


/* debug [class_header]: Header for HKDocumentType */
// The class instance for the [HKDocumentType] class.
var (
	HKDocumentTypeClass     _HKDocumentTypeClass
	HKDocumentTypeClassOnce sync.Once
)

func getHKDocumentTypeClass() _HKDocumentTypeClass {
	HKDocumentTypeClassOnce.Do(func() {
		HKDocumentTypeClass = _HKDocumentTypeClass{objc.GetClass("HKDocumentType")}
	})
	return HKDocumentTypeClass
}

type _HKDocumentTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKDocumentType */
// An interface definition for the [HKDocumentType] class.
type IHKDocumentType interface {
	IHKSampleType
	
/* debug [class_interface_properties]: Properties for HKDocumentType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKDocumentType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKDocumentType */
// Alloc allocates a new instance without initialization.
func (hc _HKDocumentTypeClass) Alloc() HKDocumentType {
	rv := objc.Send[HKDocumentType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKDocumentTypeClass) New() HKDocumentType {
	rv := objc.Send[HKDocumentType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKDocumentType) Init() HKDocumentType {
	rv := objc.Send[HKDocumentType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKDocumentType) Autorelease() HKDocumentType {
	rv := objc.Send[HKDocumentType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKDocumentType creates a new HKDocumentType instance.
func NewHKDocumentType() HKDocumentType {
	return getHKDocumentTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKDocumentType */
// A sample type used to create queries for documents.
//
// To create a document type instance, use the class’s convenience method.


// A sample type used to create queries for documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDocumentType
type HKDocumentType struct {
	HKSampleType
}

// HKDocumentTypeFrom constructs a [HKDocumentType] from an unsafe.Pointer.
//
// A sample type used to create queries for documents.
func HKDocumentTypeFrom(ptr unsafe.Pointer) HKDocumentType {
	return HKDocumentType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKDocumentType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKDocumentType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKDocumentType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKDocumentType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKDocumentType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKDocumentType */



