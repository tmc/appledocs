// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterStopMoveStepParams */


/* debug [class_header]: Header for MTRColorControlClusterStopMoveStepParams */
// The class instance for the [MTRColorControlClusterStopMoveStepParams] class.
var (
	MTRColorControlClusterStopMoveStepParamsClass     _MTRColorControlClusterStopMoveStepParamsClass
	MTRColorControlClusterStopMoveStepParamsClassOnce sync.Once
)

func getMTRColorControlClusterStopMoveStepParamsClass() _MTRColorControlClusterStopMoveStepParamsClass {
	MTRColorControlClusterStopMoveStepParamsClassOnce.Do(func() {
		MTRColorControlClusterStopMoveStepParamsClass = _MTRColorControlClusterStopMoveStepParamsClass{objc.GetClass("MTRColorControlClusterStopMoveStepParams")}
	})
	return MTRColorControlClusterStopMoveStepParamsClass
}

type _MTRColorControlClusterStopMoveStepParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterStopMoveStepParams */
// An interface definition for the [MTRColorControlClusterStopMoveStepParams] class.
type IMTRColorControlClusterStopMoveStepParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterStopMoveStepParams */
	// properties:
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterStopMoveStepParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterStopMoveStepParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterStopMoveStepParamsClass) Alloc() MTRColorControlClusterStopMoveStepParams {
	rv := objc.Send[MTRColorControlClusterStopMoveStepParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterStopMoveStepParamsClass) New() MTRColorControlClusterStopMoveStepParams {
	rv := objc.Send[MTRColorControlClusterStopMoveStepParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterStopMoveStepParams) Init() MTRColorControlClusterStopMoveStepParams {
	rv := objc.Send[MTRColorControlClusterStopMoveStepParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterStopMoveStepParams) Autorelease() MTRColorControlClusterStopMoveStepParams {
	rv := objc.Send[MTRColorControlClusterStopMoveStepParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterStopMoveStepParams creates a new MTRColorControlClusterStopMoveStepParams instance.
func NewMTRColorControlClusterStopMoveStepParams() MTRColorControlClusterStopMoveStepParams {
	return getMTRColorControlClusterStopMoveStepParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterStopMoveStepParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStopMoveStepParams
type MTRColorControlClusterStopMoveStepParams struct {
	objectivec.Object
}

// MTRColorControlClusterStopMoveStepParamsFrom constructs a [MTRColorControlClusterStopMoveStepParams] from an unsafe.Pointer.
func MTRColorControlClusterStopMoveStepParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterStopMoveStepParams {
	return MTRColorControlClusterStopMoveStepParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterStopMoveStepParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterStopMoveStepParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterStopMoveStepParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterStopMoveStepParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterStopMoveStepParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStopMoveStepParams/optionsMask
func (m_ MTRColorControlClusterStopMoveStepParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStopMoveStepParams/optionsMask
func (m_ MTRColorControlClusterStopMoveStepParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStopMoveStepParams/optionsOverride
func (m_ MTRColorControlClusterStopMoveStepParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStopMoveStepParams/optionsOverride
func (m_ MTRColorControlClusterStopMoveStepParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStopMoveStepParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterStopMoveStepParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStopMoveStepParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterStopMoveStepParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStopMoveStepParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterStopMoveStepParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterStopMoveStepParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterStopMoveStepParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterStopMoveStepParams */



