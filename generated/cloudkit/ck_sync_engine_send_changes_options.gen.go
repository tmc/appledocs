// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineSendChangesOptions */


/* debug [class_header]: Header for CKSyncEngineSendChangesOptions */
// The class instance for the [CKSyncEngineSendChangesOptions] class.
var (
	CKSyncEngineSendChangesOptionsClass     _CKSyncEngineSendChangesOptionsClass
	CKSyncEngineSendChangesOptionsClassOnce sync.Once
)

func getCKSyncEngineSendChangesOptionsClass() _CKSyncEngineSendChangesOptionsClass {
	CKSyncEngineSendChangesOptionsClassOnce.Do(func() {
		CKSyncEngineSendChangesOptionsClass = _CKSyncEngineSendChangesOptionsClass{objc.GetClass("CKSyncEngineSendChangesOptions")}
	})
	return CKSyncEngineSendChangesOptionsClass
}

type _CKSyncEngineSendChangesOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineSendChangesOptions */
// An interface definition for the [CKSyncEngineSendChangesOptions] class.
type ICKSyncEngineSendChangesOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineSendChangesOptions */
	// properties:
	OperationGroup() ICKOperationGroup
	SetOperationGroup(value ICKOperationGroup)
	Scope() ICKSyncEngineSendChangesScope
	SetScope(value ICKSyncEngineSendChangesScope)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineSendChangesOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineSendChangesOptions */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineSendChangesOptionsClass) Alloc() CKSyncEngineSendChangesOptions {
	rv := objc.Send[CKSyncEngineSendChangesOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineSendChangesOptionsClass) New() CKSyncEngineSendChangesOptions {
	rv := objc.Send[CKSyncEngineSendChangesOptions](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineSendChangesOptions) Init() CKSyncEngineSendChangesOptions {
	rv := objc.Send[CKSyncEngineSendChangesOptions](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineSendChangesOptions) Autorelease() CKSyncEngineSendChangesOptions {
	rv := objc.Send[CKSyncEngineSendChangesOptions](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineSendChangesOptions creates a new CKSyncEngineSendChangesOptions instance.
func NewCKSyncEngineSendChangesOptions() CKSyncEngineSendChangesOptions {
	return getCKSyncEngineSendChangesOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineSendChangesOptions */
// A set of options to use with a send operation.


// A set of options to use with a send operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions
type CKSyncEngineSendChangesOptions struct {
	objectivec.Object
}

// CKSyncEngineSendChangesOptionsFrom constructs a [CKSyncEngineSendChangesOptions] from an unsafe.Pointer.
//
// A set of options to use with a send operation.
func CKSyncEngineSendChangesOptionsFrom(ptr unsafe.Pointer) CKSyncEngineSendChangesOptions {
	return CKSyncEngineSendChangesOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineSendChangesOptions */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/initWithScope:
func NewCKSyncEngineSendChangesOptionsWithScope(scope ICKSyncEngineSendChangesScope) CKSyncEngineSendChangesOptions {
	instance := getCKSyncEngineSendChangesOptionsClass().Alloc()
	rv := objc.Send[CKSyncEngineSendChangesOptions](instance.ID, objc.Sel("initWithScope:"), scope)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSyncEngineSendChangesOptionsWithScope */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineSendChangesOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineSendChangesOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineSendChangesOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineSendChangesOptions */

// The operation group to use for the underlying CloudKit operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/operationGroup
func (c_ CKSyncEngineSendChangesOptions) OperationGroup() ICKOperationGroup {
	rv := objc.Send[CKOperationGroup](c_.ID, objc.Sel("operationGroup"))
	return rv
}/* debug [instance_properties/getter]: operationGroup */


// The operation group to use for the underlying CloudKit operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/operationGroup
func (c_ CKSyncEngineSendChangesOptions) SetOperationGroup(value ICKOperationGroup) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOperationGroup:"), value)
}/* debug [instance_properties/setter]: operationGroup */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/scope
func (c_ CKSyncEngineSendChangesOptions) Scope() ICKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](c_.ID, objc.Sel("scope"))
	return rv
}/* debug [instance_properties/getter]: scope */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesOptions/scope
func (c_ CKSyncEngineSendChangesOptions) SetScope(value ICKSyncEngineSendChangesScope) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setScope:"), value)
}/* debug [instance_properties/setter]: scope */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineSendChangesOptions */


