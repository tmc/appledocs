// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKRecord */


/* debug [class_header]: Header for CKRecord */
// The class instance for the [CKRecord] class.
var (
	CKRecordClass     _CKRecordClass
	CKRecordClassOnce sync.Once
)

func getCKRecordClass() _CKRecordClass {
	CKRecordClassOnce.Do(func() {
		CKRecordClass = _CKRecordClass{objc.GetClass("CKRecord")}
	})
	return CKRecordClass
}

type _CKRecordClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKRecord */
// An interface definition for the [CKRecord] class.
type ICKRecord interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKRecord */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKRecord */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKRecord */
// Alloc allocates a new instance without initialization.
func (cc _CKRecordClass) Alloc() CKRecord {
	rv := objc.Send[CKRecord](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKRecordClass) New() CKRecord {
	rv := objc.Send[CKRecord](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKRecord) Init() CKRecord {
	rv := objc.Send[CKRecord](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKRecord) Autorelease() CKRecord {
	rv := objc.Send[CKRecord](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecord creates a new CKRecord instance.
func NewCKRecord() CKRecord {
	return getCKRecordClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKRecord */
// A parent class referenced by other CloudKit classes.


// A parent class referenced by other CloudKit classes. [Full Topic]
type CKRecord struct {
	objectivec.Object
}

// CKRecordFrom constructs a [CKRecord] from an unsafe.Pointer.
//
// A parent class referenced by other CloudKit classes.
func CKRecordFrom(ptr unsafe.Pointer) CKRecord {
	return CKRecord{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKRecord *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKRecord */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKRecord */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKRecord */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKRecord */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKRecord */



