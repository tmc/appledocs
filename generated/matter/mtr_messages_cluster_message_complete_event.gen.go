// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMessagesClusterMessageCompleteEvent */


/* debug [class_header]: Header for MTRMessagesClusterMessageCompleteEvent */
// The class instance for the [MTRMessagesClusterMessageCompleteEvent] class.
var (
	MTRMessagesClusterMessageCompleteEventClass     _MTRMessagesClusterMessageCompleteEventClass
	MTRMessagesClusterMessageCompleteEventClassOnce sync.Once
)

func getMTRMessagesClusterMessageCompleteEventClass() _MTRMessagesClusterMessageCompleteEventClass {
	MTRMessagesClusterMessageCompleteEventClassOnce.Do(func() {
		MTRMessagesClusterMessageCompleteEventClass = _MTRMessagesClusterMessageCompleteEventClass{objc.GetClass("MTRMessagesClusterMessageCompleteEvent")}
	})
	return MTRMessagesClusterMessageCompleteEventClass
}

type _MTRMessagesClusterMessageCompleteEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMessagesClusterMessageCompleteEvent */
// An interface definition for the [MTRMessagesClusterMessageCompleteEvent] class.
type IMTRMessagesClusterMessageCompleteEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMessagesClusterMessageCompleteEvent */
	// properties:
	ResponseID() objc.IObject /* cross-framework: NSNumber */
	SetResponseID(value objc.IObject /* cross-framework: NSNumber */)
	FutureMessagesPreference() objc.IObject /* cross-framework: NSNumber */
	SetFutureMessagesPreference(value objc.IObject /* cross-framework: NSNumber */)
	MessageID() foundation.Data
	SetMessageID(value foundation.Data)
	Reply() objc.IObject /* cross-framework: NSString */
	SetReply(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMessagesClusterMessageCompleteEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMessagesClusterMessageCompleteEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterMessageCompleteEventClass) Alloc() MTRMessagesClusterMessageCompleteEvent {
	rv := objc.Send[MTRMessagesClusterMessageCompleteEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMessagesClusterMessageCompleteEventClass) New() MTRMessagesClusterMessageCompleteEvent {
	rv := objc.Send[MTRMessagesClusterMessageCompleteEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterMessageCompleteEvent) Init() MTRMessagesClusterMessageCompleteEvent {
	rv := objc.Send[MTRMessagesClusterMessageCompleteEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterMessageCompleteEvent) Autorelease() MTRMessagesClusterMessageCompleteEvent {
	rv := objc.Send[MTRMessagesClusterMessageCompleteEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterMessageCompleteEvent creates a new MTRMessagesClusterMessageCompleteEvent instance.
func NewMTRMessagesClusterMessageCompleteEvent() MTRMessagesClusterMessageCompleteEvent {
	return getMTRMessagesClusterMessageCompleteEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMessagesClusterMessageCompleteEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent
type MTRMessagesClusterMessageCompleteEvent struct {
	objectivec.Object
}

// MTRMessagesClusterMessageCompleteEventFrom constructs a [MTRMessagesClusterMessageCompleteEvent] from an unsafe.Pointer.
func MTRMessagesClusterMessageCompleteEventFrom(ptr unsafe.Pointer) MTRMessagesClusterMessageCompleteEvent {
	return MTRMessagesClusterMessageCompleteEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMessagesClusterMessageCompleteEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMessagesClusterMessageCompleteEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMessagesClusterMessageCompleteEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMessagesClusterMessageCompleteEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMessagesClusterMessageCompleteEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/responseID
func (m_ MTRMessagesClusterMessageCompleteEvent) ResponseID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("responseID"))
	return rv
}/* debug [instance_properties/getter]: responseID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterMessageCompleteEvent/responseID
func (m_ MTRMessagesClusterMessageCompleteEvent) SetResponseID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResponseID:"), value)
}/* debug [instance_properties/setter]: responseID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagecompleteevent/futuremessagespreference
func (m_ MTRMessagesClusterMessageCompleteEvent) FutureMessagesPreference() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("futureMessagesPreference"))
	return rv
}/* debug [instance_properties/getter]: futureMessagesPreference */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagecompleteevent/futuremessagespreference
func (m_ MTRMessagesClusterMessageCompleteEvent) SetFutureMessagesPreference(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFutureMessagesPreference:"), value)
}/* debug [instance_properties/setter]: futureMessagesPreference */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagecompleteevent/messageid
func (m_ MTRMessagesClusterMessageCompleteEvent) MessageID() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("messageID"))
	return rv
}/* debug [instance_properties/getter]: messageID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagecompleteevent/messageid
func (m_ MTRMessagesClusterMessageCompleteEvent) SetMessageID(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageID:"), value)
}/* debug [instance_properties/setter]: messageID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagecompleteevent/reply
func (m_ MTRMessagesClusterMessageCompleteEvent) Reply() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("reply"))
	return rv
}/* debug [instance_properties/getter]: reply */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustermessagecompleteevent/reply
func (m_ MTRMessagesClusterMessageCompleteEvent) SetReply(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReply:"), value)
}/* debug [instance_properties/setter]: reply */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMessagesClusterMessageCompleteEvent */



