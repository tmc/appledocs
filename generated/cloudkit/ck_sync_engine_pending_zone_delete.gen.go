// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEnginePendingZoneDelete */


/* debug [class_header]: Header for CKSyncEnginePendingZoneDelete */
// The class instance for the [CKSyncEnginePendingZoneDelete] class.
var (
	CKSyncEnginePendingZoneDeleteClass     _CKSyncEnginePendingZoneDeleteClass
	CKSyncEnginePendingZoneDeleteClassOnce sync.Once
)

func getCKSyncEnginePendingZoneDeleteClass() _CKSyncEnginePendingZoneDeleteClass {
	CKSyncEnginePendingZoneDeleteClassOnce.Do(func() {
		CKSyncEnginePendingZoneDeleteClass = _CKSyncEnginePendingZoneDeleteClass{objc.GetClass("CKSyncEnginePendingZoneDelete")}
	})
	return CKSyncEnginePendingZoneDeleteClass
}

type _CKSyncEnginePendingZoneDeleteClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEnginePendingZoneDelete */
// An interface definition for the [CKSyncEnginePendingZoneDelete] class.
type ICKSyncEnginePendingZoneDelete interface {
	ICKSyncEnginePendingDatabaseChange
	
/* debug [class_interface_properties]: Properties for CKSyncEnginePendingZoneDelete */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEnginePendingZoneDelete */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEnginePendingZoneDelete */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEnginePendingZoneDeleteClass) Alloc() CKSyncEnginePendingZoneDelete {
	rv := objc.Send[CKSyncEnginePendingZoneDelete](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEnginePendingZoneDeleteClass) New() CKSyncEnginePendingZoneDelete {
	rv := objc.Send[CKSyncEnginePendingZoneDelete](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEnginePendingZoneDelete) Init() CKSyncEnginePendingZoneDelete {
	rv := objc.Send[CKSyncEnginePendingZoneDelete](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEnginePendingZoneDelete) Autorelease() CKSyncEnginePendingZoneDelete {
	rv := objc.Send[CKSyncEnginePendingZoneDelete](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEnginePendingZoneDelete creates a new CKSyncEnginePendingZoneDelete instance.
func NewCKSyncEnginePendingZoneDelete() CKSyncEnginePendingZoneDelete {
	return getCKSyncEnginePendingZoneDeleteClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEnginePendingZoneDelete */
// An object that describes an unsent record zone deletion.


// An object that describes an unsent record zone deletion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneDelete
type CKSyncEnginePendingZoneDelete struct {
	CKSyncEnginePendingDatabaseChange
}

// CKSyncEnginePendingZoneDeleteFrom constructs a [CKSyncEnginePendingZoneDelete] from an unsafe.Pointer.
//
// An object that describes an unsent record zone deletion.
func CKSyncEnginePendingZoneDeleteFrom(ptr unsafe.Pointer) CKSyncEnginePendingZoneDelete {
	return CKSyncEnginePendingZoneDelete{
		CKSyncEnginePendingDatabaseChange: CKSyncEnginePendingDatabaseChangeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEnginePendingZoneDelete */

// Creates a pending zone delete for the specified record zone identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneDelete/initWithZoneID:
func NewCKSyncEnginePendingZoneDeleteWithZoneID(zoneID ICKRecordZoneID) CKSyncEnginePendingZoneDelete {
	instance := getCKSyncEnginePendingZoneDeleteClass().Alloc()
	rv := objc.Send[CKSyncEnginePendingZoneDelete](instance.ID, objc.Sel("initWithZoneID:"), zoneID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSyncEnginePendingZoneDeleteWithZoneID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEnginePendingZoneDelete */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEnginePendingZoneDelete */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEnginePendingZoneDelete */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEnginePendingZoneDelete */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEnginePendingZoneDelete */


