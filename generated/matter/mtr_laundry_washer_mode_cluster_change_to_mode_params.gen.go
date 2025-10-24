// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRLaundryWasherModeClusterChangeToModeParams */


/* debug [class_header]: Header for MTRLaundryWasherModeClusterChangeToModeParams */
// The class instance for the [MTRLaundryWasherModeClusterChangeToModeParams] class.
var (
	MTRLaundryWasherModeClusterChangeToModeParamsClass     _MTRLaundryWasherModeClusterChangeToModeParamsClass
	MTRLaundryWasherModeClusterChangeToModeParamsClassOnce sync.Once
)

func getMTRLaundryWasherModeClusterChangeToModeParamsClass() _MTRLaundryWasherModeClusterChangeToModeParamsClass {
	MTRLaundryWasherModeClusterChangeToModeParamsClassOnce.Do(func() {
		MTRLaundryWasherModeClusterChangeToModeParamsClass = _MTRLaundryWasherModeClusterChangeToModeParamsClass{objc.GetClass("MTRLaundryWasherModeClusterChangeToModeParams")}
	})
	return MTRLaundryWasherModeClusterChangeToModeParamsClass
}

type _MTRLaundryWasherModeClusterChangeToModeParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRLaundryWasherModeClusterChangeToModeParams */
// An interface definition for the [MTRLaundryWasherModeClusterChangeToModeParams] class.
type IMTRLaundryWasherModeClusterChangeToModeParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRLaundryWasherModeClusterChangeToModeParams */
	// properties:
	NewMode() objc.IObject /* cross-framework: NSNumber */
	SetNewMode(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRLaundryWasherModeClusterChangeToModeParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRLaundryWasherModeClusterChangeToModeParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRLaundryWasherModeClusterChangeToModeParamsClass) Alloc() MTRLaundryWasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRLaundryWasherModeClusterChangeToModeParamsClass) New() MTRLaundryWasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) Init() MTRLaundryWasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) Autorelease() MTRLaundryWasherModeClusterChangeToModeParams {
	rv := objc.Send[MTRLaundryWasherModeClusterChangeToModeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLaundryWasherModeClusterChangeToModeParams creates a new MTRLaundryWasherModeClusterChangeToModeParams instance.
func NewMTRLaundryWasherModeClusterChangeToModeParams() MTRLaundryWasherModeClusterChangeToModeParams {
	return getMTRLaundryWasherModeClusterChangeToModeParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRLaundryWasherModeClusterChangeToModeParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeParams
type MTRLaundryWasherModeClusterChangeToModeParams struct {
	objectivec.Object
}

// MTRLaundryWasherModeClusterChangeToModeParamsFrom constructs a [MTRLaundryWasherModeClusterChangeToModeParams] from an unsafe.Pointer.
func MTRLaundryWasherModeClusterChangeToModeParamsFrom(ptr unsafe.Pointer) MTRLaundryWasherModeClusterChangeToModeParams {
	return MTRLaundryWasherModeClusterChangeToModeParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRLaundryWasherModeClusterChangeToModeParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRLaundryWasherModeClusterChangeToModeParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRLaundryWasherModeClusterChangeToModeParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRLaundryWasherModeClusterChangeToModeParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRLaundryWasherModeClusterChangeToModeParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeParams/newMode
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) NewMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newMode"))
	return rv
}/* debug [instance_properties/getter]: newMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterChangeToModeParams/newMode
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) SetNewMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewMode:"), value)
}/* debug [instance_properties/setter]: newMode */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlaundrywashermodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlaundrywashermodeclusterchangetomodeparams/serversideprocessingtimeout
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlaundrywashermodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrlaundrywashermodeclusterchangetomodeparams/timedinvoketimeoutms
func (m_ MTRLaundryWasherModeClusterChangeToModeParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRLaundryWasherModeClusterChangeToModeParams */



