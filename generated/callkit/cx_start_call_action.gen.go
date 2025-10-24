// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CXStartCallAction */


/* debug [class_header]: Header for CXStartCallAction */
// The class instance for the [CXStartCallAction] class.
var (
	CXStartCallActionClass     _CXStartCallActionClass
	CXStartCallActionClassOnce sync.Once
)

func getCXStartCallActionClass() _CXStartCallActionClass {
	CXStartCallActionClassOnce.Do(func() {
		CXStartCallActionClass = _CXStartCallActionClass{objc.GetClass("CXStartCallAction")}
	})
	return CXStartCallActionClass
}

type _CXStartCallActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXStartCallAction */
// An interface definition for the [CXStartCallAction] class.
type ICXStartCallAction interface {
	ICXCallAction
	
/* debug [class_interface_properties]: Properties for CXStartCallAction */
	// properties:
	IsVideo() bool
	SetIsVideo(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXStartCallAction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXStartCallAction */
// Alloc allocates a new instance without initialization.
func (cc _CXStartCallActionClass) Alloc() CXStartCallAction {
	rv := objc.Send[CXStartCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXStartCallActionClass) New() CXStartCallAction {
	rv := objc.Send[CXStartCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXStartCallAction) Init() CXStartCallAction {
	rv := objc.Send[CXStartCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXStartCallAction) Autorelease() CXStartCallAction {
	rv := objc.Send[CXStartCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXStartCallAction creates a new CXStartCallAction instance.
func NewCXStartCallAction() CXStartCallAction {
	return getCXStartCallActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXStartCallAction */
// An encapsulation of the act of initiating an outgoing call.
//
// is a concrete subclass of . When the user initiates an outgoing call, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed. To indicate that the call started at a time other than the current time, you can instead call the .


// An encapsulation of the act of initiating an outgoing call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction
type CXStartCallAction struct {
	CXCallAction
}

// CXStartCallActionFrom constructs a [CXStartCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of initiating an outgoing call.
func CXStartCallActionFrom(ptr unsafe.Pointer) CXStartCallAction {
	return CXStartCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXStartCallAction */

// Initializes a new action to start a call with the specified UUID to a recipient with the specified handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/init(call:handle:)
func NewCXStartCallActionWithCallUUIDHandle(callUUID foundation.UUID, handle ICXHandle) CXStartCallAction {
	instance := getCXStartCallActionClass().Alloc()
	rv := objc.Send[CXStartCallAction](instance.ID, objc.Sel("initWithCallUUID:handle:"), callUUID, handle)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXStartCallActionWithCallUUIDHandle */


// Creates a new action to start a call with data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/init(coder:)
func NewCXStartCallActionWithCoder(aDecoder foundation.Coder) CXStartCallAction {
	instance := getCXStartCallActionClass().Alloc()
	rv := objc.Send[CXStartCallAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXStartCallActionWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXStartCallAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXStartCallAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXStartCallAction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXStartCallAction */

// A Boolean value that indicates whether the call is a video call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxstartcallaction/isvideo
func (c_ CXStartCallAction) IsVideo() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideo"))
	return rv
}/* debug [instance_properties/getter]: isVideo */


// A Boolean value that indicates whether the call is a video call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxstartcallaction/isvideo
func (c_ CXStartCallAction) SetIsVideo(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideo:"), value)
}/* debug [instance_properties/setter]: isVideo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXStartCallAction */


