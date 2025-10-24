// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEnginePendingRecordZoneChange */


/* debug [class_header]: Header for CKSyncEnginePendingRecordZoneChange */
// The class instance for the [CKSyncEnginePendingRecordZoneChange] class.
var (
	CKSyncEnginePendingRecordZoneChangeClass     _CKSyncEnginePendingRecordZoneChangeClass
	CKSyncEnginePendingRecordZoneChangeClassOnce sync.Once
)

func getCKSyncEnginePendingRecordZoneChangeClass() _CKSyncEnginePendingRecordZoneChangeClass {
	CKSyncEnginePendingRecordZoneChangeClassOnce.Do(func() {
		CKSyncEnginePendingRecordZoneChangeClass = _CKSyncEnginePendingRecordZoneChangeClass{objc.GetClass("CKSyncEnginePendingRecordZoneChange")}
	})
	return CKSyncEnginePendingRecordZoneChangeClass
}

type _CKSyncEnginePendingRecordZoneChangeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEnginePendingRecordZoneChange */
// An interface definition for the [CKSyncEnginePendingRecordZoneChange] class.
type ICKSyncEnginePendingRecordZoneChange interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEnginePendingRecordZoneChange */
	// properties:
	RecordID() ICKRecordID
	Type() CKSyncEnginePendingRecordZoneChangeType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEnginePendingRecordZoneChange */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEnginePendingRecordZoneChange */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEnginePendingRecordZoneChangeClass) Alloc() CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEnginePendingRecordZoneChangeClass) New() CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEnginePendingRecordZoneChange) Init() CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEnginePendingRecordZoneChange) Autorelease() CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEnginePendingRecordZoneChange creates a new CKSyncEnginePendingRecordZoneChange instance.
func NewCKSyncEnginePendingRecordZoneChange() CKSyncEnginePendingRecordZoneChange {
	return getCKSyncEnginePendingRecordZoneChangeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEnginePendingRecordZoneChange */
// Describes an unsent record modification.


// Describes an unsent record modification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange
type CKSyncEnginePendingRecordZoneChange struct {
	objectivec.Object
}

// CKSyncEnginePendingRecordZoneChangeFrom constructs a [CKSyncEnginePendingRecordZoneChange] from an unsafe.Pointer.
//
// Describes an unsent record modification.
func CKSyncEnginePendingRecordZoneChangeFrom(ptr unsafe.Pointer) CKSyncEnginePendingRecordZoneChange {
	return CKSyncEnginePendingRecordZoneChange{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEnginePendingRecordZoneChange */

// Creates a record zone change of the specified type for the given record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange/initWithRecordID:type:
func NewCKSyncEnginePendingRecordZoneChangeWithRecordIDType(recordID ICKRecordID, type_ CKSyncEnginePendingRecordZoneChangeType) CKSyncEnginePendingRecordZoneChange {
	instance := getCKSyncEnginePendingRecordZoneChangeClass().Alloc()
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](instance.ID, objc.Sel("initWithRecordID:type:"), recordID, type_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSyncEnginePendingRecordZoneChangeWithRecordIDType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEnginePendingRecordZoneChange */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEnginePendingRecordZoneChange */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEnginePendingRecordZoneChange */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEnginePendingRecordZoneChange */

// The identifier of the modified record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange/recordID
func (c_ CKSyncEnginePendingRecordZoneChange) RecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("recordID"))
	return rv
}/* debug [instance_properties/getter]: recordID */


// The type of change to make.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange/type
func (c_ CKSyncEnginePendingRecordZoneChange) Type() CKSyncEnginePendingRecordZoneChangeType {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChangeType](c_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEnginePendingRecordZoneChange */


