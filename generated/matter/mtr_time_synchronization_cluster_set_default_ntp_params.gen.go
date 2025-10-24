// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterSetDefaultNTPParams */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterSetDefaultNTPParams */
// The class instance for the [MTRTimeSynchronizationClusterSetDefaultNTPParams] class.
var (
	MTRTimeSynchronizationClusterSetDefaultNTPParamsClass     _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass
	MTRTimeSynchronizationClusterSetDefaultNTPParamsClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterSetDefaultNTPParamsClass() _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass {
	MTRTimeSynchronizationClusterSetDefaultNTPParamsClassOnce.Do(func() {
		MTRTimeSynchronizationClusterSetDefaultNTPParamsClass = _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass{objc.GetClass("MTRTimeSynchronizationClusterSetDefaultNTPParams")}
	})
	return MTRTimeSynchronizationClusterSetDefaultNTPParamsClass
}

type _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterSetDefaultNTPParams */
// An interface definition for the [MTRTimeSynchronizationClusterSetDefaultNTPParams] class.
type IMTRTimeSynchronizationClusterSetDefaultNTPParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterSetDefaultNTPParams */
	// properties:
	DefaultNTP() objc.IObject /* cross-framework: NSString */
	SetDefaultNTP(value objc.IObject /* cross-framework: NSString */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterSetDefaultNTPParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterSetDefaultNTPParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass) Alloc() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDefaultNTPParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass) New() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDefaultNTPParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) Init() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDefaultNTPParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) Autorelease() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDefaultNTPParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterSetDefaultNTPParams creates a new MTRTimeSynchronizationClusterSetDefaultNTPParams instance.
func NewMTRTimeSynchronizationClusterSetDefaultNTPParams() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	return getMTRTimeSynchronizationClusterSetDefaultNTPParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterSetDefaultNTPParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams
type MTRTimeSynchronizationClusterSetDefaultNTPParams struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterSetDefaultNTPParamsFrom constructs a [MTRTimeSynchronizationClusterSetDefaultNTPParams] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterSetDefaultNTPParamsFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterSetDefaultNTPParams {
	return MTRTimeSynchronizationClusterSetDefaultNTPParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterSetDefaultNTPParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterSetDefaultNTPParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterSetDefaultNTPParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterSetDefaultNTPParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterSetDefaultNTPParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/defaultNTP
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) DefaultNTP() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("defaultNTP"))
	return rv
}/* debug [instance_properties/getter]: defaultNTP */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/defaultNTP
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) SetDefaultNTP(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultNTP:"), value)
}/* debug [instance_properties/setter]: defaultNTP */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetdefaultntpparams/serversideprocessingtimeout
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetdefaultntpparams/serversideprocessingtimeout
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetdefaultntpparams/timedinvoketimeoutms
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersetdefaultntpparams/timedinvoketimeoutms
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterSetDefaultNTPParams */



