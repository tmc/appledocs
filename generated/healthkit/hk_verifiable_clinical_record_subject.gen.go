// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKVerifiableClinicalRecordSubject */


/* debug [class_header]: Header for HKVerifiableClinicalRecordSubject */
// The class instance for the [HKVerifiableClinicalRecordSubject] class.
var (
	HKVerifiableClinicalRecordSubjectClass     _HKVerifiableClinicalRecordSubjectClass
	HKVerifiableClinicalRecordSubjectClassOnce sync.Once
)

func getHKVerifiableClinicalRecordSubjectClass() _HKVerifiableClinicalRecordSubjectClass {
	HKVerifiableClinicalRecordSubjectClassOnce.Do(func() {
		HKVerifiableClinicalRecordSubjectClass = _HKVerifiableClinicalRecordSubjectClass{objc.GetClass("HKVerifiableClinicalRecordSubject")}
	})
	return HKVerifiableClinicalRecordSubjectClass
}

type _HKVerifiableClinicalRecordSubjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKVerifiableClinicalRecordSubject */
// An interface definition for the [HKVerifiableClinicalRecordSubject] class.
type IHKVerifiableClinicalRecordSubject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKVerifiableClinicalRecordSubject */
	// properties:
	DateOfBirthComponents() foundation.DateComponents
	FullName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKVerifiableClinicalRecordSubject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKVerifiableClinicalRecordSubject */
// Alloc allocates a new instance without initialization.
func (hc _HKVerifiableClinicalRecordSubjectClass) Alloc() HKVerifiableClinicalRecordSubject {
	rv := objc.Send[HKVerifiableClinicalRecordSubject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKVerifiableClinicalRecordSubjectClass) New() HKVerifiableClinicalRecordSubject {
	rv := objc.Send[HKVerifiableClinicalRecordSubject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKVerifiableClinicalRecordSubject) Init() HKVerifiableClinicalRecordSubject {
	rv := objc.Send[HKVerifiableClinicalRecordSubject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKVerifiableClinicalRecordSubject) Autorelease() HKVerifiableClinicalRecordSubject {
	rv := objc.Send[HKVerifiableClinicalRecordSubject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKVerifiableClinicalRecordSubject creates a new HKVerifiableClinicalRecordSubject instance.
func NewHKVerifiableClinicalRecordSubject() HKVerifiableClinicalRecordSubject {
	return getHKVerifiableClinicalRecordSubjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKVerifiableClinicalRecordSubject */
// The subject associated with a signed clinical record.
//
// objects contain data about the subject from a SMART Health Card. These cards combine both the user’s identity and clinical data into a cryptographically-signed bundle. To protect the subject’s privacy, SMART Health Cards provide the minimum required data. For more information, see .


// The subject associated with a signed clinical record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecordSubject
type HKVerifiableClinicalRecordSubject struct {
	objectivec.Object
}

// HKVerifiableClinicalRecordSubjectFrom constructs a [HKVerifiableClinicalRecordSubject] from an unsafe.Pointer.
//
// The subject associated with a signed clinical record.
func HKVerifiableClinicalRecordSubjectFrom(ptr unsafe.Pointer) HKVerifiableClinicalRecordSubject {
	return HKVerifiableClinicalRecordSubject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKVerifiableClinicalRecordSubject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKVerifiableClinicalRecordSubject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKVerifiableClinicalRecordSubject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKVerifiableClinicalRecordSubject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKVerifiableClinicalRecordSubject */

// The subject’s birthdate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecordSubject/dateOfBirthComponents
func (h_ HKVerifiableClinicalRecordSubject) DateOfBirthComponents() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](h_.ID, objc.Sel("dateOfBirthComponents"))
	return rv
}/* debug [instance_properties/getter]: dateOfBirthComponents */


// The subject’s full name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecordSubject/fullName
func (h_ HKVerifiableClinicalRecordSubject) FullName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("fullName"))
	return rv
}/* debug [instance_properties/getter]: fullName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKVerifiableClinicalRecordSubject */



