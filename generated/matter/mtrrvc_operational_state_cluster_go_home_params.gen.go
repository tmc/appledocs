// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRRVCOperationalStateClusterGoHomeParams */


/* debug [class_header]: Header for MTRRVCOperationalStateClusterGoHomeParams */
// The class instance for the [MTRRVCOperationalStateClusterGoHomeParams] class.
var (
	MTRRVCOperationalStateClusterGoHomeParamsClass     _MTRRVCOperationalStateClusterGoHomeParamsClass
	MTRRVCOperationalStateClusterGoHomeParamsClassOnce sync.Once
)

func getMTRRVCOperationalStateClusterGoHomeParamsClass() _MTRRVCOperationalStateClusterGoHomeParamsClass {
	MTRRVCOperationalStateClusterGoHomeParamsClassOnce.Do(func() {
		MTRRVCOperationalStateClusterGoHomeParamsClass = _MTRRVCOperationalStateClusterGoHomeParamsClass{objc.GetClass("MTRRVCOperationalStateClusterGoHomeParams")}
	})
	return MTRRVCOperationalStateClusterGoHomeParamsClass
}

type _MTRRVCOperationalStateClusterGoHomeParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRRVCOperationalStateClusterGoHomeParams */
// An interface definition for the [MTRRVCOperationalStateClusterGoHomeParams] class.
type IMTRRVCOperationalStateClusterGoHomeParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRRVCOperationalStateClusterGoHomeParams */
	// properties:
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRRVCOperationalStateClusterGoHomeParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRRVCOperationalStateClusterGoHomeParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRRVCOperationalStateClusterGoHomeParamsClass) Alloc() MTRRVCOperationalStateClusterGoHomeParams {
	rv := objc.Send[MTRRVCOperationalStateClusterGoHomeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRRVCOperationalStateClusterGoHomeParamsClass) New() MTRRVCOperationalStateClusterGoHomeParams {
	rv := objc.Send[MTRRVCOperationalStateClusterGoHomeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCOperationalStateClusterGoHomeParams) Init() MTRRVCOperationalStateClusterGoHomeParams {
	rv := objc.Send[MTRRVCOperationalStateClusterGoHomeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCOperationalStateClusterGoHomeParams) Autorelease() MTRRVCOperationalStateClusterGoHomeParams {
	rv := objc.Send[MTRRVCOperationalStateClusterGoHomeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCOperationalStateClusterGoHomeParams creates a new MTRRVCOperationalStateClusterGoHomeParams instance.
func NewMTRRVCOperationalStateClusterGoHomeParams() MTRRVCOperationalStateClusterGoHomeParams {
	return getMTRRVCOperationalStateClusterGoHomeParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRRVCOperationalStateClusterGoHomeParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterGoHomeParams
type MTRRVCOperationalStateClusterGoHomeParams struct {
	objectivec.Object
}

// MTRRVCOperationalStateClusterGoHomeParamsFrom constructs a [MTRRVCOperationalStateClusterGoHomeParams] from an unsafe.Pointer.
func MTRRVCOperationalStateClusterGoHomeParamsFrom(ptr unsafe.Pointer) MTRRVCOperationalStateClusterGoHomeParams {
	return MTRRVCOperationalStateClusterGoHomeParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRRVCOperationalStateClusterGoHomeParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRRVCOperationalStateClusterGoHomeParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRRVCOperationalStateClusterGoHomeParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRRVCOperationalStateClusterGoHomeParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRRVCOperationalStateClusterGoHomeParams */

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterGoHomeParams/timedInvokeTimeoutMs
func (m_ MTRRVCOperationalStateClusterGoHomeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterGoHomeParams/timedInvokeTimeoutMs
func (m_ MTRRVCOperationalStateClusterGoHomeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclustergohomeparams/serversideprocessingtimeout
func (m_ MTRRVCOperationalStateClusterGoHomeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclustergohomeparams/serversideprocessingtimeout
func (m_ MTRRVCOperationalStateClusterGoHomeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRRVCOperationalStateClusterGoHomeParams */



