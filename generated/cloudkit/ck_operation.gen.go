// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKOperation */


/* debug [class_header]: Header for CKOperation */
// The class instance for the [CKOperation] class.
var (
	CKOperationClass     _CKOperationClass
	CKOperationClassOnce sync.Once
)

func getCKOperationClass() _CKOperationClass {
	CKOperationClassOnce.Do(func() {
		CKOperationClass = _CKOperationClass{objc.GetClass("CKOperation")}
	})
	return CKOperationClass
}

type _CKOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKOperation */
// An interface definition for the [CKOperation] class.
type ICKOperation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKOperation */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKOperationClass) Alloc() CKOperation {
	rv := objc.Send[CKOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKOperationClass) New() CKOperation {
	rv := objc.Send[CKOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKOperation) Init() CKOperation {
	rv := objc.Send[CKOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKOperation) Autorelease() CKOperation {
	rv := objc.Send[CKOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKOperation creates a new CKOperation instance.
func NewCKOperation() CKOperation {
	return getCKOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKOperation */
// A parent class referenced by other CloudKit classes.


// A parent class referenced by other CloudKit classes. [Full Topic]
type CKOperation struct {
	objectivec.Object
}

// CKOperationFrom constructs a [CKOperation] from an unsafe.Pointer.
//
// A parent class referenced by other CloudKit classes.
func CKOperationFrom(ptr unsafe.Pointer) CKOperation {
	return CKOperation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKOperation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKOperation */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKOperation */



