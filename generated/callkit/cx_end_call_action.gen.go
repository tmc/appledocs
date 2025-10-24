// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CXEndCallAction */


/* debug [class_header]: Header for CXEndCallAction */
// The class instance for the [CXEndCallAction] class.
var (
	CXEndCallActionClass     _CXEndCallActionClass
	CXEndCallActionClassOnce sync.Once
)

func getCXEndCallActionClass() _CXEndCallActionClass {
	CXEndCallActionClassOnce.Do(func() {
		CXEndCallActionClass = _CXEndCallActionClass{objc.GetClass("CXEndCallAction")}
	})
	return CXEndCallActionClass
}

type _CXEndCallActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXEndCallAction */
// An interface definition for the [CXEndCallAction] class.
type ICXEndCallAction interface {
	ICXCallAction
	
/* debug [class_interface_properties]: Properties for CXEndCallAction */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXEndCallAction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXEndCallAction */
// Alloc allocates a new instance without initialization.
func (cc _CXEndCallActionClass) Alloc() CXEndCallAction {
	rv := objc.Send[CXEndCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXEndCallActionClass) New() CXEndCallAction {
	rv := objc.Send[CXEndCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXEndCallAction) Init() CXEndCallAction {
	rv := objc.Send[CXEndCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXEndCallAction) Autorelease() CXEndCallAction {
	rv := objc.Send[CXEndCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXEndCallAction creates a new CXEndCallAction instance.
func NewCXEndCallAction() CXEndCallAction {
	return getCXEndCallActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXEndCallAction */
// An encapsulation of the act of ending a call.
//
// is a concrete subclass of . When the user initiates an outgoing call, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed. To indicate that the call ended at a time other than the current time, you can instead call the


// An encapsulation of the act of ending a call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXEndCallAction
type CXEndCallAction struct {
	CXCallAction
}

// CXEndCallActionFrom constructs a [CXEndCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of ending a call.
func CXEndCallActionFrom(ptr unsafe.Pointer) CXEndCallAction {
	return CXEndCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXEndCallAction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXEndCallAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXEndCallAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXEndCallAction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXEndCallAction */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXEndCallAction */


