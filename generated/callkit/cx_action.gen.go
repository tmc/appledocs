// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CXAction */


/* debug [class_header]: Header for CXAction */
// The class instance for the [CXAction] class.
var (
	CXActionClass     _CXActionClass
	CXActionClassOnce sync.Once
)

func getCXActionClass() _CXActionClass {
	CXActionClassOnce.Do(func() {
		CXActionClass = _CXActionClass{objc.GetClass("CXAction")}
	})
	return CXActionClass
}

type _CXActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXAction */
// An interface definition for the [CXAction] class.
type ICXAction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CXAction */
	// properties:
	IsComplete() bool
	SetIsComplete(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXAction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXAction */
// Alloc allocates a new instance without initialization.
func (cc _CXActionClass) Alloc() CXAction {
	rv := objc.Send[CXAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXActionClass) New() CXAction {
	rv := objc.Send[CXAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXAction) Init() CXAction {
	rv := objc.Send[CXAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXAction) Autorelease() CXAction {
	rv := objc.Send[CXAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXAction creates a new CXAction instance.
func NewCXAction() CXAction {
	return getCXActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXAction */
// An abstract class that declares a programmatic interface for objects that represent a telephony action.
//
// Each instance of is uniquely identified by a , which is generated on initialization. An action also tracks whether it has been completed or not. To perform one or more actions, you add them to a new object and pass the transaction to an instance of using the method. After each action is performed by the telephony provider, the provider’s delegate calls either the method, indicating that the action was successfully performed, or the method, to indicate that an error occurred; both of these methods set the property of the action to . The subclass is an abstract class that represents an action associated with a object. The CallKit framework provides several concrete subclasses to represent actions such as answering a call and putting a call on hold.


// An abstract class that declares a programmatic interface for objects that represent a telephony action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction
type CXAction struct {
	objectivec.Object
}

// CXActionFrom constructs a [CXAction] from an unsafe.Pointer.
//
// An abstract class that declares a programmatic interface for objects that represent a telephony action.
func CXActionFrom(ptr unsafe.Pointer) CXAction {
	return CXAction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXAction */

// Creates a new telephony action with data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction/init(coder:)
func NewCXActionWithCoder(aDecoder foundation.Coder) CXAction {
	instance := getCXActionClass().Alloc()
	rv := objc.Send[CXAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXActionWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXAction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXAction */

// A Boolean value that indicates whether the action has been performed by the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxaction/iscomplete
func (c_ CXAction) IsComplete() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isComplete"))
	return rv
}/* debug [instance_properties/getter]: isComplete */


// A Boolean value that indicates whether the action has been performed by the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxaction/iscomplete
func (c_ CXAction) SetIsComplete(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsComplete:"), value)
}/* debug [instance_properties/setter]: isComplete */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXAction */


