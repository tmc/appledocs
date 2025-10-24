// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineFailedZoneSave */


/* debug [class_header]: Header for CKSyncEngineFailedZoneSave */
// The class instance for the [CKSyncEngineFailedZoneSave] class.
var (
	CKSyncEngineFailedZoneSaveClass     _CKSyncEngineFailedZoneSaveClass
	CKSyncEngineFailedZoneSaveClassOnce sync.Once
)

func getCKSyncEngineFailedZoneSaveClass() _CKSyncEngineFailedZoneSaveClass {
	CKSyncEngineFailedZoneSaveClassOnce.Do(func() {
		CKSyncEngineFailedZoneSaveClass = _CKSyncEngineFailedZoneSaveClass{objc.GetClass("CKSyncEngineFailedZoneSave")}
	})
	return CKSyncEngineFailedZoneSaveClass
}

type _CKSyncEngineFailedZoneSaveClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineFailedZoneSave */
// An interface definition for the [CKSyncEngineFailedZoneSave] class.
type ICKSyncEngineFailedZoneSave interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineFailedZoneSave */
	// properties:
	Error() objc.IObject /* cross-framework: Error */
	RecordZone() ICKRecordZone
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineFailedZoneSave */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineFailedZoneSave */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFailedZoneSaveClass) Alloc() CKSyncEngineFailedZoneSave {
	rv := objc.Send[CKSyncEngineFailedZoneSave](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineFailedZoneSaveClass) New() CKSyncEngineFailedZoneSave {
	rv := objc.Send[CKSyncEngineFailedZoneSave](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFailedZoneSave) Init() CKSyncEngineFailedZoneSave {
	rv := objc.Send[CKSyncEngineFailedZoneSave](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFailedZoneSave) Autorelease() CKSyncEngineFailedZoneSave {
	rv := objc.Send[CKSyncEngineFailedZoneSave](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFailedZoneSave creates a new CKSyncEngineFailedZoneSave instance.
func NewCKSyncEngineFailedZoneSave() CKSyncEngineFailedZoneSave {
	return getCKSyncEngineFailedZoneSaveClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineFailedZoneSave */
// An object that describes an unsuccessful attempt to modify a single record zone.


// An object that describes an unsuccessful attempt to modify a single record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedZoneSave
type CKSyncEngineFailedZoneSave struct {
	objectivec.Object
}

// CKSyncEngineFailedZoneSaveFrom constructs a [CKSyncEngineFailedZoneSave] from an unsafe.Pointer.
//
// An object that describes an unsuccessful attempt to modify a single record zone.
func CKSyncEngineFailedZoneSaveFrom(ptr unsafe.Pointer) CKSyncEngineFailedZoneSave {
	return CKSyncEngineFailedZoneSave{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineFailedZoneSave *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineFailedZoneSave */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineFailedZoneSave */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineFailedZoneSave */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineFailedZoneSave */

// A error that describes the reason for the unsuccessful attempt to modify the associated record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedZoneSave/error
func (c_ CKSyncEngineFailedZoneSave) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](c_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// The record zone that CloudKit is unable to modify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedZoneSave/recordZone
func (c_ CKSyncEngineFailedZoneSave) RecordZone() ICKRecordZone {
	rv := objc.Send[CKRecordZone](c_.ID, objc.Sel("recordZone"))
	return rv
}/* debug [instance_properties/getter]: recordZone */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineFailedZoneSave */



