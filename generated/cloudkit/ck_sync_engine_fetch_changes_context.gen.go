// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineFetchChangesContext */


/* debug [class_header]: Header for CKSyncEngineFetchChangesContext */
// The class instance for the [CKSyncEngineFetchChangesContext] class.
var (
	CKSyncEngineFetchChangesContextClass     _CKSyncEngineFetchChangesContextClass
	CKSyncEngineFetchChangesContextClassOnce sync.Once
)

func getCKSyncEngineFetchChangesContextClass() _CKSyncEngineFetchChangesContextClass {
	CKSyncEngineFetchChangesContextClassOnce.Do(func() {
		CKSyncEngineFetchChangesContextClass = _CKSyncEngineFetchChangesContextClass{objc.GetClass("CKSyncEngineFetchChangesContext")}
	})
	return CKSyncEngineFetchChangesContextClass
}

type _CKSyncEngineFetchChangesContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineFetchChangesContext */
// An interface definition for the [CKSyncEngineFetchChangesContext] class.
type ICKSyncEngineFetchChangesContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineFetchChangesContext */
	// properties:
	Options() ICKSyncEngineFetchChangesOptions
	Reason() CKSyncEngineSyncReason
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineFetchChangesContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineFetchChangesContext */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFetchChangesContextClass) Alloc() CKSyncEngineFetchChangesContext {
	rv := objc.Send[CKSyncEngineFetchChangesContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineFetchChangesContextClass) New() CKSyncEngineFetchChangesContext {
	rv := objc.Send[CKSyncEngineFetchChangesContext](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFetchChangesContext) Init() CKSyncEngineFetchChangesContext {
	rv := objc.Send[CKSyncEngineFetchChangesContext](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFetchChangesContext) Autorelease() CKSyncEngineFetchChangesContext {
	rv := objc.Send[CKSyncEngineFetchChangesContext](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchChangesContext creates a new CKSyncEngineFetchChangesContext instance.
func NewCKSyncEngineFetchChangesContext() CKSyncEngineFetchChangesContext {
	return getCKSyncEngineFetchChangesContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineFetchChangesContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesContext
type CKSyncEngineFetchChangesContext struct {
	objectivec.Object
}

// CKSyncEngineFetchChangesContextFrom constructs a [CKSyncEngineFetchChangesContext] from an unsafe.Pointer.
func CKSyncEngineFetchChangesContextFrom(ptr unsafe.Pointer) CKSyncEngineFetchChangesContext {
	return CKSyncEngineFetchChangesContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineFetchChangesContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineFetchChangesContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineFetchChangesContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineFetchChangesContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineFetchChangesContext */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesContext/options
func (c_ CKSyncEngineFetchChangesContext) Options() ICKSyncEngineFetchChangesOptions {
	rv := objc.Send[CKSyncEngineFetchChangesOptions](c_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesContext/reason
func (c_ CKSyncEngineFetchChangesContext) Reason() CKSyncEngineSyncReason {
	rv := objc.Send[CKSyncEngineSyncReason](c_.ID, objc.Sel("reason"))
	return rv
}/* debug [instance_properties/getter]: reason */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineFetchChangesContext */



