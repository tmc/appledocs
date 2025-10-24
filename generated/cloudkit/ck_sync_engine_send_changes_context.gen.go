// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineSendChangesContext */


/* debug [class_header]: Header for CKSyncEngineSendChangesContext */
// The class instance for the [CKSyncEngineSendChangesContext] class.
var (
	CKSyncEngineSendChangesContextClass     _CKSyncEngineSendChangesContextClass
	CKSyncEngineSendChangesContextClassOnce sync.Once
)

func getCKSyncEngineSendChangesContextClass() _CKSyncEngineSendChangesContextClass {
	CKSyncEngineSendChangesContextClassOnce.Do(func() {
		CKSyncEngineSendChangesContextClass = _CKSyncEngineSendChangesContextClass{objc.GetClass("CKSyncEngineSendChangesContext")}
	})
	return CKSyncEngineSendChangesContextClass
}

type _CKSyncEngineSendChangesContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineSendChangesContext */
// An interface definition for the [CKSyncEngineSendChangesContext] class.
type ICKSyncEngineSendChangesContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineSendChangesContext */
	// properties:
	Options() ICKSyncEngineSendChangesOptions
	Reason() CKSyncEngineSyncReason
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineSendChangesContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineSendChangesContext */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineSendChangesContextClass) Alloc() CKSyncEngineSendChangesContext {
	rv := objc.Send[CKSyncEngineSendChangesContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineSendChangesContextClass) New() CKSyncEngineSendChangesContext {
	rv := objc.Send[CKSyncEngineSendChangesContext](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineSendChangesContext) Init() CKSyncEngineSendChangesContext {
	rv := objc.Send[CKSyncEngineSendChangesContext](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineSendChangesContext) Autorelease() CKSyncEngineSendChangesContext {
	rv := objc.Send[CKSyncEngineSendChangesContext](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineSendChangesContext creates a new CKSyncEngineSendChangesContext instance.
func NewCKSyncEngineSendChangesContext() CKSyncEngineSendChangesContext {
	return getCKSyncEngineSendChangesContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineSendChangesContext */
// An object that describes a single attempt to send changes to the iCloud servers.
//
// A sync engine has two ways to send changes to iCloud — periodically, in cooperation with the system scheduler, and manually, whenever your app invokes the method. This object provides information about a single attempt to send changes that includes both the reason for the attempt and any additional options in use by the attempt.


// An object that describes a single attempt to send changes to the iCloud servers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesContext
type CKSyncEngineSendChangesContext struct {
	objectivec.Object
}

// CKSyncEngineSendChangesContextFrom constructs a [CKSyncEngineSendChangesContext] from an unsafe.Pointer.
//
// An object that describes a single attempt to send changes to the iCloud servers.
func CKSyncEngineSendChangesContextFrom(ptr unsafe.Pointer) CKSyncEngineSendChangesContext {
	return CKSyncEngineSendChangesContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineSendChangesContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineSendChangesContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineSendChangesContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineSendChangesContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineSendChangesContext */

// The additional options for the send operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesContext/options
func (c_ CKSyncEngineSendChangesContext) Options() ICKSyncEngineSendChangesOptions {
	rv := objc.Send[CKSyncEngineSendChangesOptions](c_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// The reason for the send operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesContext/reason
func (c_ CKSyncEngineSendChangesContext) Reason() CKSyncEngineSyncReason {
	rv := objc.Send[CKSyncEngineSyncReason](c_.ID, objc.Sel("reason"))
	return rv
}/* debug [instance_properties/getter]: reason */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineSendChangesContext */



