// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRLevelControlClusterMoveToLevelWithOnOffParams */


/* debug [class_header]: Header for MTRLevelControlClusterMoveToLevelWithOnOffParams */
// The class instance for the [MTRLevelControlClusterMoveToLevelWithOnOffParams] class.
var (
	MTRLevelControlClusterMoveToLevelWithOnOffParamsClass     _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass
	MTRLevelControlClusterMoveToLevelWithOnOffParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveToLevelWithOnOffParamsClass() _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass {
	MTRLevelControlClusterMoveToLevelWithOnOffParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveToLevelWithOnOffParamsClass = _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass{objc.GetClass("MTRLevelControlClusterMoveToLevelWithOnOffParams")}
	})
	return MTRLevelControlClusterMoveToLevelWithOnOffParamsClass
}

type _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRLevelControlClusterMoveToLevelWithOnOffParams */
// An interface definition for the [MTRLevelControlClusterMoveToLevelWithOnOffParams] class.
type IMTRLevelControlClusterMoveToLevelWithOnOffParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRLevelControlClusterMoveToLevelWithOnOffParams */
	// properties:
	Level() objc.IObject /* cross-framework: NSNumber */
	SetLevel(value objc.IObject /* cross-framework: NSNumber */)
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TransitionTime() objc.IObject /* cross-framework: NSNumber */
	SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRLevelControlClusterMoveToLevelWithOnOffParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRLevelControlClusterMoveToLevelWithOnOffParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass) Alloc() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass) New() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) Init() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) Autorelease() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveToLevelWithOnOffParams creates a new MTRLevelControlClusterMoveToLevelWithOnOffParams instance.
func NewMTRLevelControlClusterMoveToLevelWithOnOffParams() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	return getMTRLevelControlClusterMoveToLevelWithOnOffParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRLevelControlClusterMoveToLevelWithOnOffParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams
type MTRLevelControlClusterMoveToLevelWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveToLevelWithOnOffParamsFrom constructs a [MTRLevelControlClusterMoveToLevelWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveToLevelWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveToLevelWithOnOffParams {
	return MTRLevelControlClusterMoveToLevelWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRLevelControlClusterMoveToLevelWithOnOffParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRLevelControlClusterMoveToLevelWithOnOffParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRLevelControlClusterMoveToLevelWithOnOffParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRLevelControlClusterMoveToLevelWithOnOffParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRLevelControlClusterMoveToLevelWithOnOffParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams/level
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) Level() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("level"))
	return rv
}/* debug [instance_properties/getter]: level */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams/level
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetLevel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLevel:"), value)
}/* debug [instance_properties/setter]: level */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams/optionsMask
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams/optionsMask
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams/optionsOverride
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams/optionsOverride
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams/transitionTime
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams/transitionTime
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRLevelControlClusterMoveToLevelWithOnOffParams */



