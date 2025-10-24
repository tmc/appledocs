// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEnginePendingDatabaseChange */


/* debug [class_header]: Header for CKSyncEnginePendingDatabaseChange */
// The class instance for the [CKSyncEnginePendingDatabaseChange] class.
var (
	CKSyncEnginePendingDatabaseChangeClass     _CKSyncEnginePendingDatabaseChangeClass
	CKSyncEnginePendingDatabaseChangeClassOnce sync.Once
)

func getCKSyncEnginePendingDatabaseChangeClass() _CKSyncEnginePendingDatabaseChangeClass {
	CKSyncEnginePendingDatabaseChangeClassOnce.Do(func() {
		CKSyncEnginePendingDatabaseChangeClass = _CKSyncEnginePendingDatabaseChangeClass{objc.GetClass("CKSyncEnginePendingDatabaseChange")}
	})
	return CKSyncEnginePendingDatabaseChangeClass
}

type _CKSyncEnginePendingDatabaseChangeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEnginePendingDatabaseChange */
// An interface definition for the [CKSyncEnginePendingDatabaseChange] class.
type ICKSyncEnginePendingDatabaseChange interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEnginePendingDatabaseChange */
	// properties:
	Type() CKSyncEnginePendingDatabaseChangeType
	ZoneID() ICKRecordZoneID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEnginePendingDatabaseChange */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEnginePendingDatabaseChange */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEnginePendingDatabaseChangeClass) Alloc() CKSyncEnginePendingDatabaseChange {
	rv := objc.Send[CKSyncEnginePendingDatabaseChange](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEnginePendingDatabaseChangeClass) New() CKSyncEnginePendingDatabaseChange {
	rv := objc.Send[CKSyncEnginePendingDatabaseChange](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEnginePendingDatabaseChange) Init() CKSyncEnginePendingDatabaseChange {
	rv := objc.Send[CKSyncEnginePendingDatabaseChange](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEnginePendingDatabaseChange) Autorelease() CKSyncEnginePendingDatabaseChange {
	rv := objc.Send[CKSyncEnginePendingDatabaseChange](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEnginePendingDatabaseChange creates a new CKSyncEnginePendingDatabaseChange instance.
func NewCKSyncEnginePendingDatabaseChange() CKSyncEnginePendingDatabaseChange {
	return getCKSyncEnginePendingDatabaseChangeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEnginePendingDatabaseChange */
// An object that describes an unsent database modification.


// An object that describes an unsent database modification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingDatabaseChange
type CKSyncEnginePendingDatabaseChange struct {
	objectivec.Object
}

// CKSyncEnginePendingDatabaseChangeFrom constructs a [CKSyncEnginePendingDatabaseChange] from an unsafe.Pointer.
//
// An object that describes an unsent database modification.
func CKSyncEnginePendingDatabaseChangeFrom(ptr unsafe.Pointer) CKSyncEnginePendingDatabaseChange {
	return CKSyncEnginePendingDatabaseChange{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEnginePendingDatabaseChange *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEnginePendingDatabaseChange */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEnginePendingDatabaseChange */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEnginePendingDatabaseChange */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEnginePendingDatabaseChange */

// The type of database change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingDatabaseChange/type
func (c_ CKSyncEnginePendingDatabaseChange) Type() CKSyncEnginePendingDatabaseChangeType {
	rv := objc.Send[CKSyncEnginePendingDatabaseChangeType](c_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The identifier of the record zone to change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingDatabaseChange/zoneID
func (c_ CKSyncEnginePendingDatabaseChange) ZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}/* debug [instance_properties/getter]: zoneID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEnginePendingDatabaseChange */



