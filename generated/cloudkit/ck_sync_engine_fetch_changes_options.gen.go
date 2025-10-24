// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineFetchChangesOptions */


/* debug [class_header]: Header for CKSyncEngineFetchChangesOptions */
// The class instance for the [CKSyncEngineFetchChangesOptions] class.
var (
	CKSyncEngineFetchChangesOptionsClass     _CKSyncEngineFetchChangesOptionsClass
	CKSyncEngineFetchChangesOptionsClassOnce sync.Once
)

func getCKSyncEngineFetchChangesOptionsClass() _CKSyncEngineFetchChangesOptionsClass {
	CKSyncEngineFetchChangesOptionsClassOnce.Do(func() {
		CKSyncEngineFetchChangesOptionsClass = _CKSyncEngineFetchChangesOptionsClass{objc.GetClass("CKSyncEngineFetchChangesOptions")}
	})
	return CKSyncEngineFetchChangesOptionsClass
}

type _CKSyncEngineFetchChangesOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineFetchChangesOptions */
// An interface definition for the [CKSyncEngineFetchChangesOptions] class.
type ICKSyncEngineFetchChangesOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineFetchChangesOptions */
	// properties:
	OperationGroup() ICKOperationGroup
	SetOperationGroup(value ICKOperationGroup)
	PrioritizedZoneIDs() []CKRecordZoneID
	SetPrioritizedZoneIDs(value []CKRecordZoneID)
	Scope() ICKSyncEngineFetchChangesScope
	SetScope(value ICKSyncEngineFetchChangesScope)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineFetchChangesOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineFetchChangesOptions */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFetchChangesOptionsClass) Alloc() CKSyncEngineFetchChangesOptions {
	rv := objc.Send[CKSyncEngineFetchChangesOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineFetchChangesOptionsClass) New() CKSyncEngineFetchChangesOptions {
	rv := objc.Send[CKSyncEngineFetchChangesOptions](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFetchChangesOptions) Init() CKSyncEngineFetchChangesOptions {
	rv := objc.Send[CKSyncEngineFetchChangesOptions](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFetchChangesOptions) Autorelease() CKSyncEngineFetchChangesOptions {
	rv := objc.Send[CKSyncEngineFetchChangesOptions](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchChangesOptions creates a new CKSyncEngineFetchChangesOptions instance.
func NewCKSyncEngineFetchChangesOptions() CKSyncEngineFetchChangesOptions {
	return getCKSyncEngineFetchChangesOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineFetchChangesOptions */
// A set of options to use with a fetch operation.


// A set of options to use with a fetch operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions
type CKSyncEngineFetchChangesOptions struct {
	objectivec.Object
}

// CKSyncEngineFetchChangesOptionsFrom constructs a [CKSyncEngineFetchChangesOptions] from an unsafe.Pointer.
//
// A set of options to use with a fetch operation.
func CKSyncEngineFetchChangesOptionsFrom(ptr unsafe.Pointer) CKSyncEngineFetchChangesOptions {
	return CKSyncEngineFetchChangesOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineFetchChangesOptions */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions/initWithScope:
func NewCKSyncEngineFetchChangesOptionsWithScope(scope ICKSyncEngineFetchChangesScope) CKSyncEngineFetchChangesOptions {
	instance := getCKSyncEngineFetchChangesOptionsClass().Alloc()
	rv := objc.Send[CKSyncEngineFetchChangesOptions](instance.ID, objc.Sel("initWithScope:"), scope)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSyncEngineFetchChangesOptionsWithScope */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineFetchChangesOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineFetchChangesOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineFetchChangesOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineFetchChangesOptions */

// The operation group to use for the underlying CloudKit operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions/operationGroup
func (c_ CKSyncEngineFetchChangesOptions) OperationGroup() ICKOperationGroup {
	rv := objc.Send[CKOperationGroup](c_.ID, objc.Sel("operationGroup"))
	return rv
}/* debug [instance_properties/getter]: operationGroup */


// The operation group to use for the underlying CloudKit operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions/operationGroup
func (c_ CKSyncEngineFetchChangesOptions) SetOperationGroup(value ICKOperationGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOperationGroup:"), value)
}/* debug [instance_properties/setter]: operationGroup */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions/prioritizedZoneIDs
func (c_ CKSyncEngineFetchChangesOptions) PrioritizedZoneIDs() []CKRecordZoneID {
	rv := objc.Send[[]CKRecordZoneID](c_.ID, objc.Sel("prioritizedZoneIDs"))
	return rv
}/* debug [instance_properties/getter]: prioritizedZoneIDs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions/prioritizedZoneIDs
func (c_ CKSyncEngineFetchChangesOptions) SetPrioritizedZoneIDs(value []CKRecordZoneID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrioritizedZoneIDs:"), nsArray)
}/* debug [instance_properties/setter]: prioritizedZoneIDs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions/scope
func (c_ CKSyncEngineFetchChangesOptions) Scope() ICKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](c_.ID, objc.Sel("scope"))
	return rv
}/* debug [instance_properties/getter]: scope */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions/scope
func (c_ CKSyncEngineFetchChangesOptions) SetScope(value ICKSyncEngineFetchChangesScope) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScope:"), value)
}/* debug [instance_properties/setter]: scope */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineFetchChangesOptions */


