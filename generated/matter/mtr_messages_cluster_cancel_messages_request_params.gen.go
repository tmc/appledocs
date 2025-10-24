// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMessagesClusterCancelMessagesRequestParams */


/* debug [class_header]: Header for MTRMessagesClusterCancelMessagesRequestParams */
// The class instance for the [MTRMessagesClusterCancelMessagesRequestParams] class.
var (
	MTRMessagesClusterCancelMessagesRequestParamsClass     _MTRMessagesClusterCancelMessagesRequestParamsClass
	MTRMessagesClusterCancelMessagesRequestParamsClassOnce sync.Once
)

func getMTRMessagesClusterCancelMessagesRequestParamsClass() _MTRMessagesClusterCancelMessagesRequestParamsClass {
	MTRMessagesClusterCancelMessagesRequestParamsClassOnce.Do(func() {
		MTRMessagesClusterCancelMessagesRequestParamsClass = _MTRMessagesClusterCancelMessagesRequestParamsClass{objc.GetClass("MTRMessagesClusterCancelMessagesRequestParams")}
	})
	return MTRMessagesClusterCancelMessagesRequestParamsClass
}

type _MTRMessagesClusterCancelMessagesRequestParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMessagesClusterCancelMessagesRequestParams */
// An interface definition for the [MTRMessagesClusterCancelMessagesRequestParams] class.
type IMTRMessagesClusterCancelMessagesRequestParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMessagesClusterCancelMessagesRequestParams */
	// properties:
	MessageIDs() objc.IObject /* cross-framework: NSArray */
	SetMessageIDs(value objc.IObject /* cross-framework: NSArray */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMessagesClusterCancelMessagesRequestParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMessagesClusterCancelMessagesRequestParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRMessagesClusterCancelMessagesRequestParamsClass) Alloc() MTRMessagesClusterCancelMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterCancelMessagesRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMessagesClusterCancelMessagesRequestParamsClass) New() MTRMessagesClusterCancelMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterCancelMessagesRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMessagesClusterCancelMessagesRequestParams) Init() MTRMessagesClusterCancelMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterCancelMessagesRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMessagesClusterCancelMessagesRequestParams) Autorelease() MTRMessagesClusterCancelMessagesRequestParams {
	rv := objc.Send[MTRMessagesClusterCancelMessagesRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMessagesClusterCancelMessagesRequestParams creates a new MTRMessagesClusterCancelMessagesRequestParams instance.
func NewMTRMessagesClusterCancelMessagesRequestParams() MTRMessagesClusterCancelMessagesRequestParams {
	return getMTRMessagesClusterCancelMessagesRequestParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMessagesClusterCancelMessagesRequestParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterCancelMessagesRequestParams
type MTRMessagesClusterCancelMessagesRequestParams struct {
	objectivec.Object
}

// MTRMessagesClusterCancelMessagesRequestParamsFrom constructs a [MTRMessagesClusterCancelMessagesRequestParams] from an unsafe.Pointer.
func MTRMessagesClusterCancelMessagesRequestParamsFrom(ptr unsafe.Pointer) MTRMessagesClusterCancelMessagesRequestParams {
	return MTRMessagesClusterCancelMessagesRequestParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMessagesClusterCancelMessagesRequestParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMessagesClusterCancelMessagesRequestParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMessagesClusterCancelMessagesRequestParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMessagesClusterCancelMessagesRequestParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMessagesClusterCancelMessagesRequestParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterCancelMessagesRequestParams/messageIDs
func (m_ MTRMessagesClusterCancelMessagesRequestParams) MessageIDs() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("messageIDs"))
	return rv
}/* debug [instance_properties/getter]: messageIDs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMessagesClusterCancelMessagesRequestParams/messageIDs
func (m_ MTRMessagesClusterCancelMessagesRequestParams) SetMessageIDs(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageIDs:"), value)
}/* debug [instance_properties/setter]: messageIDs */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustercancelmessagesrequestparams/serversideprocessingtimeout
func (m_ MTRMessagesClusterCancelMessagesRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustercancelmessagesrequestparams/serversideprocessingtimeout
func (m_ MTRMessagesClusterCancelMessagesRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustercancelmessagesrequestparams/timedinvoketimeoutms
func (m_ MTRMessagesClusterCancelMessagesRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmessagesclustercancelmessagesrequestparams/timedinvoketimeoutms
func (m_ MTRMessagesClusterCancelMessagesRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMessagesClusterCancelMessagesRequestParams */



