// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [HKVerifiableClinicalRecordSubject] class.
type IHKVerifiableClinicalRecordSubject interface {
	objectivec.IObject
	DateOfBirthComponents() foundation.DateComponents
	SetDateOfBirthComponents(value foundation.IDateComponents)
	FullName() string
	SetFullName(value string)
}

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

// Alloc allocates a new instance without initialization.
func (hc _HKVerifiableClinicalRecordSubjectClass) Alloc() HKVerifiableClinicalRecordSubject {
	rv := objc.Send[HKVerifiableClinicalRecordSubject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The subject’s birthdate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecordsubject/dateofbirthcomponents
func (h_ HKVerifiableClinicalRecordSubject) DateOfBirthComponents() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](h_.ID, objc.Sel("dateOfBirthComponents"))
	return rv
}


// The subject’s birthdate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecordsubject/dateofbirthcomponents
func (h_ HKVerifiableClinicalRecordSubject) SetDateOfBirthComponents(value foundation.IDateComponents) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDateOfBirthComponents:"), value)
}


// The subject’s full name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecordsubject/fullname
func (h_ HKVerifiableClinicalRecordSubject) FullName() string {
	rv := objc.Send[string](h_.ID, objc.Sel("fullName"))
	return rv
}


// The subject’s full name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecordsubject/fullname
func (h_ HKVerifiableClinicalRecordSubject) SetFullName(value string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setFullName:"), objc.String(value))
}



