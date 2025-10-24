// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CXSetGroupCallAction */


/* debug [class_header]: Header for CXSetGroupCallAction */
// The class instance for the [CXSetGroupCallAction] class.
var (
	CXSetGroupCallActionClass     _CXSetGroupCallActionClass
	CXSetGroupCallActionClassOnce sync.Once
)

func getCXSetGroupCallActionClass() _CXSetGroupCallActionClass {
	CXSetGroupCallActionClassOnce.Do(func() {
		CXSetGroupCallActionClass = _CXSetGroupCallActionClass{objc.GetClass("CXSetGroupCallAction")}
	})
	return CXSetGroupCallActionClass
}

type _CXSetGroupCallActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXSetGroupCallAction */
// An interface definition for the [CXSetGroupCallAction] class.
type ICXSetGroupCallAction interface {
	ICXCallAction
	
/* debug [class_interface_properties]: Properties for CXSetGroupCallAction */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXSetGroupCallAction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXSetGroupCallAction */
// Alloc allocates a new instance without initialization.
func (cc _CXSetGroupCallActionClass) Alloc() CXSetGroupCallAction {
	rv := objc.Send[CXSetGroupCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXSetGroupCallActionClass) New() CXSetGroupCallAction {
	rv := objc.Send[CXSetGroupCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXSetGroupCallAction) Init() CXSetGroupCallAction {
	rv := objc.Send[CXSetGroupCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXSetGroupCallAction) Autorelease() CXSetGroupCallAction {
	rv := objc.Send[CXSetGroupCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXSetGroupCallAction creates a new CXSetGroupCallAction instance.
func NewCXSetGroupCallAction() CXSetGroupCallAction {
	return getCXSetGroupCallActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXSetGroupCallAction */
// An encapsulation of the act of grouping or ungrouping calls.
//
// is a concrete subclass of . When the user or the system groups a call with another call, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed. A group call allows more than two recipients to simultaneously communicate with one another.


// An encapsulation of the act of grouping or ungrouping calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetGroupCallAction
type CXSetGroupCallAction struct {
	CXCallAction
}

// CXSetGroupCallActionFrom constructs a [CXSetGroupCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of grouping or ungrouping calls.
func CXSetGroupCallActionFrom(ptr unsafe.Pointer) CXSetGroupCallAction {
	return CXSetGroupCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXSetGroupCallAction */

// Initializes a new action for a call identified by a given UUID, as well as a call to group with identified by another UUID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetGroupCallAction/init(call:callUUIDToGroupWith:)
func NewCXSetGroupCallActionWithCallUUIDCallUUIDToGroupWith(callUUID foundation.UUID, callUUIDToGroupWith foundation.UUID) CXSetGroupCallAction {
	instance := getCXSetGroupCallActionClass().Alloc()
	rv := objc.Send[CXSetGroupCallAction](instance.ID, objc.Sel("initWithCallUUID:callUUIDToGroupWith:"), callUUID, callUUIDToGroupWith)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXSetGroupCallActionWithCallUUIDCallUUIDToGroupWith */


// Creates a new action to group calls with data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetGroupCallAction/init(coder:)
func NewCXSetGroupCallActionWithCoder(aDecoder foundation.Coder) CXSetGroupCallAction {
	instance := getCXSetGroupCallActionClass().Alloc()
	rv := objc.Send[CXSetGroupCallAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXSetGroupCallActionWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXSetGroupCallAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXSetGroupCallAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXSetGroupCallAction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXSetGroupCallAction */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXSetGroupCallAction */


