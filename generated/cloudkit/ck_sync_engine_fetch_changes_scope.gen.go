// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineFetchChangesScope */


/* debug [class_header]: Header for CKSyncEngineFetchChangesScope */
// The class instance for the [CKSyncEngineFetchChangesScope] class.
var (
	CKSyncEngineFetchChangesScopeClass     _CKSyncEngineFetchChangesScopeClass
	CKSyncEngineFetchChangesScopeClassOnce sync.Once
)

func getCKSyncEngineFetchChangesScopeClass() _CKSyncEngineFetchChangesScopeClass {
	CKSyncEngineFetchChangesScopeClassOnce.Do(func() {
		CKSyncEngineFetchChangesScopeClass = _CKSyncEngineFetchChangesScopeClass{objc.GetClass("CKSyncEngineFetchChangesScope")}
	})
	return CKSyncEngineFetchChangesScopeClass
}

type _CKSyncEngineFetchChangesScopeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineFetchChangesScope */
// An interface definition for the [CKSyncEngineFetchChangesScope] class.
type ICKSyncEngineFetchChangesScope interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineFetchChangesScope */
	// properties:
	ExcludedZoneIDs() unsafe.Pointer
	ZoneIDs() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineFetchChangesScope */
	// methods:
	ContainsZoneID(zoneID ICKRecordZoneID) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineFetchChangesScope */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFetchChangesScopeClass) Alloc() CKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineFetchChangesScopeClass) New() CKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFetchChangesScope) Init() CKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFetchChangesScope) Autorelease() CKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchChangesScope creates a new CKSyncEngineFetchChangesScope instance.
func NewCKSyncEngineFetchChangesScope() CKSyncEngineFetchChangesScope {
	return getCKSyncEngineFetchChangesScopeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineFetchChangesScope */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope
type CKSyncEngineFetchChangesScope struct {
	objectivec.Object
}

// CKSyncEngineFetchChangesScopeFrom constructs a [CKSyncEngineFetchChangesScope] from an unsafe.Pointer.
func CKSyncEngineFetchChangesScopeFrom(ptr unsafe.Pointer) CKSyncEngineFetchChangesScope {
	return CKSyncEngineFetchChangesScope{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineFetchChangesScope */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/initWithExcludedZoneIDs:
func NewCKSyncEngineFetchChangesScopeWithExcludedZoneIDs(zoneIDs unsafe.Pointer) CKSyncEngineFetchChangesScope {
	instance := getCKSyncEngineFetchChangesScopeClass().Alloc()
	rv := objc.Send[CKSyncEngineFetchChangesScope](instance.ID, objc.Sel("initWithExcludedZoneIDs:"), zoneIDs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSyncEngineFetchChangesScopeWithExcludedZoneIDs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/initWithZoneIDs:
func NewCKSyncEngineFetchChangesScopeWithZoneIDs(zoneIDs unsafe.Pointer) CKSyncEngineFetchChangesScope {
	instance := getCKSyncEngineFetchChangesScopeClass().Alloc()
	rv := objc.Send[CKSyncEngineFetchChangesScope](instance.ID, objc.Sel("initWithZoneIDs:"), zoneIDs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSyncEngineFetchChangesScopeWithZoneIDs */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineFetchChangesScope */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineFetchChangesScope */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineFetchChangesScope */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/containsZoneID:
func (c_ CKSyncEngineFetchChangesScope) ContainsZoneID(zoneID ICKRecordZoneID) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("containsZoneID:"), zoneID)
	return rv
}/* debug [instance_methods/method]: ContainsZoneID */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineFetchChangesScope */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/excludedZoneIDs
func (c_ CKSyncEngineFetchChangesScope) ExcludedZoneIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("excludedZoneIDs"))
	return rv
}/* debug [instance_properties/getter]: excludedZoneIDs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/zoneIDs
func (c_ CKSyncEngineFetchChangesScope) ZoneIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("zoneIDs"))
	return rv
}/* debug [instance_properties/getter]: zoneIDs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineFetchChangesScope */


