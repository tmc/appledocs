// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CXSetMutedCallAction */


/* debug [class_header]: Header for CXSetMutedCallAction */
// The class instance for the [CXSetMutedCallAction] class.
var (
	CXSetMutedCallActionClass     _CXSetMutedCallActionClass
	CXSetMutedCallActionClassOnce sync.Once
)

func getCXSetMutedCallActionClass() _CXSetMutedCallActionClass {
	CXSetMutedCallActionClassOnce.Do(func() {
		CXSetMutedCallActionClass = _CXSetMutedCallActionClass{objc.GetClass("CXSetMutedCallAction")}
	})
	return CXSetMutedCallActionClass
}

type _CXSetMutedCallActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXSetMutedCallAction */
// An interface definition for the [CXSetMutedCallAction] class.
type ICXSetMutedCallAction interface {
	ICXCallAction
	
/* debug [class_interface_properties]: Properties for CXSetMutedCallAction */
	// properties:
	IsMuted() bool
	SetIsMuted(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXSetMutedCallAction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXSetMutedCallAction */
// Alloc allocates a new instance without initialization.
func (cc _CXSetMutedCallActionClass) Alloc() CXSetMutedCallAction {
	rv := objc.Send[CXSetMutedCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXSetMutedCallActionClass) New() CXSetMutedCallAction {
	rv := objc.Send[CXSetMutedCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXSetMutedCallAction) Init() CXSetMutedCallAction {
	rv := objc.Send[CXSetMutedCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXSetMutedCallAction) Autorelease() CXSetMutedCallAction {
	rv := objc.Send[CXSetMutedCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXSetMutedCallAction creates a new CXSetMutedCallAction instance.
func NewCXSetMutedCallAction() CXSetMutedCallAction {
	return getCXSetMutedCallActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXSetMutedCallAction */
// An encapsulation of the act of muting or unmuting a call.
//
// is a concrete subclass of . When the user or the system mutes a call, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed. When a caller mutes a call, that caller is unable to communicate with other callers until they unmute the call. A muted caller still receives communication from other unmuted callers.


// An encapsulation of the act of muting or unmuting a call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetMutedCallAction
type CXSetMutedCallAction struct {
	CXCallAction
}

// CXSetMutedCallActionFrom constructs a [CXSetMutedCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of muting or unmuting a call.
func CXSetMutedCallActionFrom(ptr unsafe.Pointer) CXSetMutedCallAction {
	return CXSetMutedCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXSetMutedCallAction */

// Initializes a new action for a call identified by a given UUID, as well as whether the call is muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetMutedCallAction/init(call:muted:)
func NewCXSetMutedCallActionWithCallUUIDMuted(callUUID foundation.UUID, muted bool) CXSetMutedCallAction {
	instance := getCXSetMutedCallActionClass().Alloc()
	rv := objc.Send[CXSetMutedCallAction](instance.ID, objc.Sel("initWithCallUUID:muted:"), callUUID, muted)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXSetMutedCallActionWithCallUUIDMuted */


// Creates a new action for a call with data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetMutedCallAction/init(coder:)
func NewCXSetMutedCallActionWithCoder(aDecoder foundation.Coder) CXSetMutedCallAction {
	instance := getCXSetMutedCallActionClass().Alloc()
	rv := objc.Send[CXSetMutedCallAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXSetMutedCallActionWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXSetMutedCallAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXSetMutedCallAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXSetMutedCallAction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXSetMutedCallAction */

// A Boolean value that indicates whether the call is muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxsetmutedcallaction/ismuted
func (c_ CXSetMutedCallAction) IsMuted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMuted"))
	return rv
}/* debug [instance_properties/getter]: isMuted */


// A Boolean value that indicates whether the call is muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxsetmutedcallaction/ismuted
func (c_ CXSetMutedCallAction) SetIsMuted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsMuted:"), value)
}/* debug [instance_properties/setter]: isMuted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXSetMutedCallAction */


