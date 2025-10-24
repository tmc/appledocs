// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEMessageActionDecision */


/* debug [class_header]: Header for MEMessageActionDecision */
// The class instance for the [MEMessageActionDecision] class.
var (
	MEMessageActionDecisionClass     _MEMessageActionDecisionClass
	MEMessageActionDecisionClassOnce sync.Once
)

func getMEMessageActionDecisionClass() _MEMessageActionDecisionClass {
	MEMessageActionDecisionClassOnce.Do(func() {
		MEMessageActionDecisionClass = _MEMessageActionDecisionClass{objc.GetClass("MEMessageActionDecision")}
	})
	return MEMessageActionDecisionClass
}

type _MEMessageActionDecisionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEMessageActionDecision */
// An interface definition for the [MEMessageActionDecision] class.
type IMEMessageActionDecision interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEMessageActionDecision */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEMessageActionDecision */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEMessageActionDecision */
// Alloc allocates a new instance without initialization.
func (mc _MEMessageActionDecisionClass) Alloc() MEMessageActionDecision {
	rv := objc.Send[MEMessageActionDecision](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEMessageActionDecisionClass) New() MEMessageActionDecision {
	rv := objc.Send[MEMessageActionDecision](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEMessageActionDecision) Init() MEMessageActionDecision {
	rv := objc.Send[MEMessageActionDecision](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEMessageActionDecision) Autorelease() MEMessageActionDecision {
	rv := objc.Send[MEMessageActionDecision](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEMessageActionDecision creates a new MEMessageActionDecision instance.
func NewMEMessageActionDecision() MEMessageActionDecision {
	return getMEMessageActionDecisionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEMessageActionDecision */
// The action that the system performs on a message, or a request to ask the action handler again later when the message content is available.


// The action that the system performs on a message, or a request to ask the action handler again later when the message content is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageActionDecision
type MEMessageActionDecision struct {
	objectivec.Object
}

// MEMessageActionDecisionFrom constructs a [MEMessageActionDecision] from an unsafe.Pointer.
//
// The action that the system performs on a message, or a request to ask the action handler again later when the message content is available.
func MEMessageActionDecisionFrom(ptr unsafe.Pointer) MEMessageActionDecision {
	return MEMessageActionDecision{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEMessageActionDecision *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEMessageActionDecision */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageActionDecision/action(_:)
func (mc _MEMessageActionDecisionClass) DecisionApplyingAction(action IMEMessageAction) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("decisionApplyingAction:"), action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecisionApplyingAction) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageActionDecision/actions(_:)
func (mc _MEMessageActionDecisionClass) DecisionApplyingActions(actions []MEMessageAction) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("decisionApplyingActions:"), actions)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecisionApplyingActions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEMessageActionDecision */

// An object that indicates the handler needs the message content before it can decide what action to take on a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageActionDecision/invokeAgainWithBody
func (mc _MEMessageActionDecisionClass) InvokeAgainWithBody() MEMessageActionDecision {
	rv := objc.Send[MEMessageActionDecision](objc.ID(mc.class), objc.Sel("invokeAgainWithBody"))
	return rv
}/* debug [class_properties_class/property]: invokeAgainWithBody */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEMessageActionDecision */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEMessageActionDecision */

// An object that indicates the handler needs the message content before it can decide what action to take on a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEMessageActionDecision/invokeAgainWithBody
func (m_ MEMessageActionDecision) InvokeAgainWithBody() IMEMessageActionDecision {
	rv := objc.Send[MEMessageActionDecision](m_.ID, objc.Sel("invokeAgainWithBody"))
	return rv
}/* debug [instance_properties/getter]: invokeAgainWithBody */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEMessageActionDecision */



