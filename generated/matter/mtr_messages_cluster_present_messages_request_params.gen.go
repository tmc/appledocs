// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMessagesClusterPresentMessagesRequestParams */


/* debug [class_header]: Header for MTRMessagesClusterPresentMessagesRequestParams */
// The class instance for the [MTRMessagesClusterPresentMessagesRequestParams] class.
var (
	MTRMessagesClusterPresentMessagesRequestParamsClass     _MTRMessagesClusterPresentMessagesRequestParamsClass
	MTRMessagesClusterPresentMessagesRequestParamsClassOnce sync.Once
)

func getMTRMessagesClusterPresentMessagesRequestParamsClass() _MTRMessagesClusterPresentMessagesRequestParamsClass {
	MTRMessagesClusterPresentMessagesRequestParamsClassOnce.Do(func() {
		MTRMessagesClusterPresentMessagesRequestParamsClass = _MTRMessagesClusterPresentMessagesRequestParamsClass{objc.GetClass("MTRMessagesClusterPresentMessagesRequestParams")}
	})
	return MTRMessagesClusterPresentMessagesRequestParamsClass
}

type _MTRMessagesClusterPresentMessagesRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMessagesClusterPresentMessagesRequestParams */
// An interface definition for the [MTRMessagesClusterPresentMessagesRequestParams] class.
type IMTRMessagesClusterPresentMessagesRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMessagesClusterPresentMessagesRequestParams */
	// properties:
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	MessageControl() objc.IObject /* cross-framework: NSNumber */
	SetMessageControl(value objc.IObject /* cross-framework: NSNumber */)
	MessageID() foundation.Data
	SetMessageID(value foundation.Data)
	MessageText() objc.IObject /* cross-framework: NSString */
	SetMessageText(value objc.IObject /* cross-framework: NSString */)
	Priority() objc.IObject /* cross-framework: NSNumber */
	SetPriority(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StartTime() objc.IObject /* cross-framework: NSNumber */
	SetStartTime(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMessagesClusterPresentMessagesRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMessagesClusterPresentMessagesRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterPresentMessagesRequestParamsClass) Alloc() MTRMessagesClusterPresentMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterPresentMessagesRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMessagesClusterPresentMessagesRequestParamsClass) New() MTRMessagesClusterPresentMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterPresentMessagesRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterPresentMessagesRequestParams) Init() MTRMessagesClusterPresentMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterPresentMessagesRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterPresentMessagesRequestParams) Autorelease() MTRMessagesClusterPresentMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterPresentMessagesRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterPresentMessagesRequestParams creates a new MTRMessagesClusterPresentMessagesRequestParams instance.
func NewMTRMessagesClusterPresentMessagesRequestParams() MTRMessagesClusterPresentMessagesRequestParams {
	return getMTRMessagesClusterPresentMessagesRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMessagesClusterPresentMessagesRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams
type MTRMessagesClusterPresentMessagesRequestParams struct {
	objectivec.Object
}

// MTRMessagesClusterPresentMessagesRequestParamsFrom constructs a [MTRMessagesClusterPresentMessagesRequestParams] from an unsafe.Pointer.
func MTRMessagesClusterPresentMessagesRequestParamsFrom(ptr unsafe.Pointer) MTRMessagesClusterPresentMessagesRequestParams {
	return MTRMessagesClusterPresentMessagesRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMessagesClusterPresentMessagesRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMessagesClusterPresentMessagesRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMessagesClusterPresentMessagesRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMessagesClusterPresentMessagesRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMessagesClusterPresentMessagesRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/duration
func (m_ MTRMessagesClusterPresentMessagesRequestParams) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterPresentMessagesRequestParams/duration
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/messagecontrol
func (m_ MTRMessagesClusterPresentMessagesRequestParams) MessageControl() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("messageControl"))
	return rv
}/* debug [instance_properties/getter]: messageControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/messagecontrol
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetMessageControl(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageControl:"), value)
}/* debug [instance_properties/setter]: messageControl */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/messageid
func (m_ MTRMessagesClusterPresentMessagesRequestParams) MessageID() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("messageID"))
	return rv
}/* debug [instance_properties/getter]: messageID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/messageid
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetMessageID(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageID:"), value)
}/* debug [instance_properties/setter]: messageID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/messagetext
func (m_ MTRMessagesClusterPresentMessagesRequestParams) MessageText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("messageText"))
	return rv
}/* debug [instance_properties/getter]: messageText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/messagetext
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetMessageText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageText:"), value)
}/* debug [instance_properties/setter]: messageText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/priority
func (m_ MTRMessagesClusterPresentMessagesRequestParams) Priority() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("priority"))
	return rv
}/* debug [instance_properties/getter]: priority */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/priority
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetPriority(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPriority:"), value)
}/* debug [instance_properties/setter]: priority */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/serversideprocessingtimeout
func (m_ MTRMessagesClusterPresentMessagesRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/serversideprocessingtimeout
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/starttime
func (m_ MTRMessagesClusterPresentMessagesRequestParams) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}/* debug [instance_properties/getter]: startTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/starttime
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}/* debug [instance_properties/setter]: startTime */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/timedinvoketimeoutms
func (m_ MTRMessagesClusterPresentMessagesRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclusterpresentmessagesrequestparams/timedinvoketimeoutms
func (m_ MTRMessagesClusterPresentMessagesRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMessagesClusterPresentMessagesRequestParams */



