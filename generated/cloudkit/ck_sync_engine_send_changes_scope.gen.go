// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineSendChangesScope */


/* debug [class_header]: Header for CKSyncEngineSendChangesScope */
// The class instance for the [CKSyncEngineSendChangesScope] class.
var (
	CKSyncEngineSendChangesScopeClass     _CKSyncEngineSendChangesScopeClass
	CKSyncEngineSendChangesScopeClassOnce sync.Once
)

func getCKSyncEngineSendChangesScopeClass() _CKSyncEngineSendChangesScopeClass {
	CKSyncEngineSendChangesScopeClassOnce.Do(func() {
		CKSyncEngineSendChangesScopeClass = _CKSyncEngineSendChangesScopeClass{objc.GetClass("CKSyncEngineSendChangesScope")}
	})
	return CKSyncEngineSendChangesScopeClass
}

type _CKSyncEngineSendChangesScopeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineSendChangesScope */
// An interface definition for the [CKSyncEngineSendChangesScope] class.
type ICKSyncEngineSendChangesScope interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineSendChangesScope */
	// properties:
	ExcludedZoneIDs() unsafe.Pointer
	RecordIDs() unsafe.Pointer
	ZoneIDs() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineSendChangesScope */
	// methods:
	ContainsPendingRecordZoneChange(pendingRecordZoneChange ICKSyncEnginePendingRecordZoneChange) bool
	ContainsRecordID(recordID ICKRecordID) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineSendChangesScope */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineSendChangesScopeClass) Alloc() CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineSendChangesScopeClass) New() CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineSendChangesScope) Init() CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineSendChangesScope) Autorelease() CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineSendChangesScope creates a new CKSyncEngineSendChangesScope instance.
func NewCKSyncEngineSendChangesScope() CKSyncEngineSendChangesScope {
	return getCKSyncEngineSendChangesScopeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineSendChangesScope */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope
type CKSyncEngineSendChangesScope struct {
	objectivec.Object
}

// CKSyncEngineSendChangesScopeFrom constructs a [CKSyncEngineSendChangesScope] from an unsafe.Pointer.
func CKSyncEngineSendChangesScopeFrom(ptr unsafe.Pointer) CKSyncEngineSendChangesScope {
	return CKSyncEngineSendChangesScope{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineSendChangesScope */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/initWithExcludedZoneIDs:
func NewCKSyncEngineSendChangesScopeWithExcludedZoneIDs(excludedZoneIDs unsafe.Pointer) CKSyncEngineSendChangesScope {
	instance := getCKSyncEngineSendChangesScopeClass().Alloc()
	rv := objc.Send[CKSyncEngineSendChangesScope](instance.ID, objc.Sel("initWithExcludedZoneIDs:"), excludedZoneIDs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSyncEngineSendChangesScopeWithExcludedZoneIDs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/initWithRecordIDs:
func NewCKSyncEngineSendChangesScopeWithRecordIDs(recordIDs unsafe.Pointer) CKSyncEngineSendChangesScope {
	instance := getCKSyncEngineSendChangesScopeClass().Alloc()
	rv := objc.Send[CKSyncEngineSendChangesScope](instance.ID, objc.Sel("initWithRecordIDs:"), recordIDs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSyncEngineSendChangesScopeWithRecordIDs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/initWithZoneIDs:
func NewCKSyncEngineSendChangesScopeWithZoneIDs(zoneIDs unsafe.Pointer) CKSyncEngineSendChangesScope {
	instance := getCKSyncEngineSendChangesScopeClass().Alloc()
	rv := objc.Send[CKSyncEngineSendChangesScope](instance.ID, objc.Sel("initWithZoneIDs:"), zoneIDs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSyncEngineSendChangesScopeWithZoneIDs */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineSendChangesScope */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineSendChangesScope */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineSendChangesScope */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/containsPendingRecordZoneChange:
func (c_ CKSyncEngineSendChangesScope) ContainsPendingRecordZoneChange(pendingRecordZoneChange ICKSyncEnginePendingRecordZoneChange) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("containsPendingRecordZoneChange:"), pendingRecordZoneChange)
	return rv
}/* debug [instance_methods/method]: ContainsPendingRecordZoneChange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/containsRecordID:
func (c_ CKSyncEngineSendChangesScope) ContainsRecordID(recordID ICKRecordID) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("containsRecordID:"), recordID)
	return rv
}/* debug [instance_methods/method]: ContainsRecordID */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineSendChangesScope */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/excludedZoneIDs
func (c_ CKSyncEngineSendChangesScope) ExcludedZoneIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("excludedZoneIDs"))
	return rv
}/* debug [instance_properties/getter]: excludedZoneIDs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/recordIDs
func (c_ CKSyncEngineSendChangesScope) RecordIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordIDs"))
	return rv
}/* debug [instance_properties/getter]: recordIDs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/zoneIDs
func (c_ CKSyncEngineSendChangesScope) ZoneIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("zoneIDs"))
	return rv
}/* debug [instance_properties/getter]: zoneIDs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineSendChangesScope */


