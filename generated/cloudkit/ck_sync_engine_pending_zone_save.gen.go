// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEnginePendingZoneSave */


/* debug [class_header]: Header for CKSyncEnginePendingZoneSave */
// The class instance for the [CKSyncEnginePendingZoneSave] class.
var (
	CKSyncEnginePendingZoneSaveClass     _CKSyncEnginePendingZoneSaveClass
	CKSyncEnginePendingZoneSaveClassOnce sync.Once
)

func getCKSyncEnginePendingZoneSaveClass() _CKSyncEnginePendingZoneSaveClass {
	CKSyncEnginePendingZoneSaveClassOnce.Do(func() {
		CKSyncEnginePendingZoneSaveClass = _CKSyncEnginePendingZoneSaveClass{objc.GetClass("CKSyncEnginePendingZoneSave")}
	})
	return CKSyncEnginePendingZoneSaveClass
}

type _CKSyncEnginePendingZoneSaveClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEnginePendingZoneSave */
// An interface definition for the [CKSyncEnginePendingZoneSave] class.
type ICKSyncEnginePendingZoneSave interface {
	ICKSyncEnginePendingDatabaseChange
	
/* debug [class_interface_properties]: Properties for CKSyncEnginePendingZoneSave */
	// properties:
	Zone() ICKRecordZone
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEnginePendingZoneSave */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEnginePendingZoneSave */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEnginePendingZoneSaveClass) Alloc() CKSyncEnginePendingZoneSave {
	rv := objc.Send[CKSyncEnginePendingZoneSave](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEnginePendingZoneSaveClass) New() CKSyncEnginePendingZoneSave {
	rv := objc.Send[CKSyncEnginePendingZoneSave](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEnginePendingZoneSave) Init() CKSyncEnginePendingZoneSave {
	rv := objc.Send[CKSyncEnginePendingZoneSave](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEnginePendingZoneSave) Autorelease() CKSyncEnginePendingZoneSave {
	rv := objc.Send[CKSyncEnginePendingZoneSave](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEnginePendingZoneSave creates a new CKSyncEnginePendingZoneSave instance.
func NewCKSyncEnginePendingZoneSave() CKSyncEnginePendingZoneSave {
	return getCKSyncEnginePendingZoneSaveClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEnginePendingZoneSave */
// An object that describes an unsent record zone modification.


// An object that describes an unsent record zone modification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneSave
type CKSyncEnginePendingZoneSave struct {
	CKSyncEnginePendingDatabaseChange
}

// CKSyncEnginePendingZoneSaveFrom constructs a [CKSyncEnginePendingZoneSave] from an unsafe.Pointer.
//
// An object that describes an unsent record zone modification.
func CKSyncEnginePendingZoneSaveFrom(ptr unsafe.Pointer) CKSyncEnginePendingZoneSave {
	return CKSyncEnginePendingZoneSave{
		CKSyncEnginePendingDatabaseChange: CKSyncEnginePendingDatabaseChangeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEnginePendingZoneSave */

// Creates a pending zone save for the specified record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneSave/initWithZone:
func NewCKSyncEnginePendingZoneSaveWithZone(zone ICKRecordZone) CKSyncEnginePendingZoneSave {
	instance := getCKSyncEnginePendingZoneSaveClass().Alloc()
	rv := objc.Send[CKSyncEnginePendingZoneSave](instance.ID, objc.Sel("initWithZone:"), zone)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSyncEnginePendingZoneSaveWithZone */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEnginePendingZoneSave */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEnginePendingZoneSave */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEnginePendingZoneSave */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEnginePendingZoneSave */

// The record zone to save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneSave/zone
func (c_ CKSyncEnginePendingZoneSave) Zone() ICKRecordZone {
	rv := objc.Send[CKRecordZone](c_.ID, objc.Sel("zone"))
	return rv
}/* debug [instance_properties/getter]: zone */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEnginePendingZoneSave */


