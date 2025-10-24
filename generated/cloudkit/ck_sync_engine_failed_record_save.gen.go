// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineFailedRecordSave */


/* debug [class_header]: Header for CKSyncEngineFailedRecordSave */
// The class instance for the [CKSyncEngineFailedRecordSave] class.
var (
	CKSyncEngineFailedRecordSaveClass     _CKSyncEngineFailedRecordSaveClass
	CKSyncEngineFailedRecordSaveClassOnce sync.Once
)

func getCKSyncEngineFailedRecordSaveClass() _CKSyncEngineFailedRecordSaveClass {
	CKSyncEngineFailedRecordSaveClassOnce.Do(func() {
		CKSyncEngineFailedRecordSaveClass = _CKSyncEngineFailedRecordSaveClass{objc.GetClass("CKSyncEngineFailedRecordSave")}
	})
	return CKSyncEngineFailedRecordSaveClass
}

type _CKSyncEngineFailedRecordSaveClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineFailedRecordSave */
// An interface definition for the [CKSyncEngineFailedRecordSave] class.
type ICKSyncEngineFailedRecordSave interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineFailedRecordSave */
	// properties:
	Error() objc.IObject /* cross-framework: Error */
	Record() objc.IObject /* cross-framework: CKRecord */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineFailedRecordSave */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineFailedRecordSave */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFailedRecordSaveClass) Alloc() CKSyncEngineFailedRecordSave {
	rv := objc.Send[CKSyncEngineFailedRecordSave](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineFailedRecordSaveClass) New() CKSyncEngineFailedRecordSave {
	rv := objc.Send[CKSyncEngineFailedRecordSave](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFailedRecordSave) Init() CKSyncEngineFailedRecordSave {
	rv := objc.Send[CKSyncEngineFailedRecordSave](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFailedRecordSave) Autorelease() CKSyncEngineFailedRecordSave {
	rv := objc.Send[CKSyncEngineFailedRecordSave](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFailedRecordSave creates a new CKSyncEngineFailedRecordSave instance.
func NewCKSyncEngineFailedRecordSave() CKSyncEngineFailedRecordSave {
	return getCKSyncEngineFailedRecordSaveClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineFailedRecordSave */
// A type that describes an unsuccessful attempt to modify an individual record.


// A type that describes an unsuccessful attempt to modify an individual record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedRecordSave
type CKSyncEngineFailedRecordSave struct {
	objectivec.Object
}

// CKSyncEngineFailedRecordSaveFrom constructs a [CKSyncEngineFailedRecordSave] from an unsafe.Pointer.
//
// A type that describes an unsuccessful attempt to modify an individual record.
func CKSyncEngineFailedRecordSaveFrom(ptr unsafe.Pointer) CKSyncEngineFailedRecordSave {
	return CKSyncEngineFailedRecordSave{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineFailedRecordSave *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineFailedRecordSave */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineFailedRecordSave */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineFailedRecordSave */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineFailedRecordSave */

// A error that describes the reason for the unsuccessful attempt to modify the associated record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedRecordSave/error
func (c_ CKSyncEngineFailedRecordSave) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](c_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// The record that CloudKit is unable to modify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedRecordSave/record
func (c_ CKSyncEngineFailedRecordSave) Record() objc.IObject /* cross-framework: CKRecord */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("record"))
	return rv
}/* debug [instance_properties/getter]: record */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineFailedRecordSave */



