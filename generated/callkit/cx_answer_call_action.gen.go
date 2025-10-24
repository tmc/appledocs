// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CXAnswerCallAction */


/* debug [class_header]: Header for CXAnswerCallAction */
// The class instance for the [CXAnswerCallAction] class.
var (
	CXAnswerCallActionClass     _CXAnswerCallActionClass
	CXAnswerCallActionClassOnce sync.Once
)

func getCXAnswerCallActionClass() _CXAnswerCallActionClass {
	CXAnswerCallActionClassOnce.Do(func() {
		CXAnswerCallActionClass = _CXAnswerCallActionClass{objc.GetClass("CXAnswerCallAction")}
	})
	return CXAnswerCallActionClass
}

type _CXAnswerCallActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXAnswerCallAction */
// An interface definition for the [CXAnswerCallAction] class.
type ICXAnswerCallAction interface {
	ICXCallAction
	
/* debug [class_interface_properties]: Properties for CXAnswerCallAction */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXAnswerCallAction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXAnswerCallAction */
// Alloc allocates a new instance without initialization.
func (cc _CXAnswerCallActionClass) Alloc() CXAnswerCallAction {
	rv := objc.Send[CXAnswerCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXAnswerCallActionClass) New() CXAnswerCallAction {
	rv := objc.Send[CXAnswerCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXAnswerCallAction) Init() CXAnswerCallAction {
	rv := objc.Send[CXAnswerCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXAnswerCallAction) Autorelease() CXAnswerCallAction {
	rv := objc.Send[CXAnswerCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXAnswerCallAction creates a new CXAnswerCallAction instance.
func NewCXAnswerCallAction() CXAnswerCallAction {
	return getCXAnswerCallActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXAnswerCallAction */
// An encapsulation of the act of answering an incoming call.
//
// is a concrete subclass of . When an incoming call is allowed by the system and approved by the user, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed. To indicate that the call connected at a time other than the current time, you can instead call the .


// An encapsulation of the act of answering an incoming call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAnswerCallAction
type CXAnswerCallAction struct {
	CXCallAction
}

// CXAnswerCallActionFrom constructs a [CXAnswerCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of answering an incoming call.
func CXAnswerCallActionFrom(ptr unsafe.Pointer) CXAnswerCallAction {
	return CXAnswerCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXAnswerCallAction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXAnswerCallAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXAnswerCallAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXAnswerCallAction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXAnswerCallAction */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXAnswerCallAction */


