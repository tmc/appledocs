// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRLevelControlClusterMoveToLevelParams */


/* debug [class_header]: Header for MTRLevelControlClusterMoveToLevelParams */
// The class instance for the [MTRLevelControlClusterMoveToLevelParams] class.
var (
	MTRLevelControlClusterMoveToLevelParamsClass     _MTRLevelControlClusterMoveToLevelParamsClass
	MTRLevelControlClusterMoveToLevelParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveToLevelParamsClass() _MTRLevelControlClusterMoveToLevelParamsClass {
	MTRLevelControlClusterMoveToLevelParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveToLevelParamsClass = _MTRLevelControlClusterMoveToLevelParamsClass{objc.GetClass("MTRLevelControlClusterMoveToLevelParams")}
	})
	return MTRLevelControlClusterMoveToLevelParamsClass
}

type _MTRLevelControlClusterMoveToLevelParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRLevelControlClusterMoveToLevelParams */
// An interface definition for the [MTRLevelControlClusterMoveToLevelParams] class.
type IMTRLevelControlClusterMoveToLevelParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRLevelControlClusterMoveToLevelParams */
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

	
/* debug [class_interface_methods]: Methods for MTRLevelControlClusterMoveToLevelParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRLevelControlClusterMoveToLevelParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveToLevelParamsClass) Alloc() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRLevelControlClusterMoveToLevelParamsClass) New() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveToLevelParams) Init() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveToLevelParams) Autorelease() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveToLevelParams creates a new MTRLevelControlClusterMoveToLevelParams instance.
func NewMTRLevelControlClusterMoveToLevelParams() MTRLevelControlClusterMoveToLevelParams {
	return getMTRLevelControlClusterMoveToLevelParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRLevelControlClusterMoveToLevelParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams
type MTRLevelControlClusterMoveToLevelParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveToLevelParamsFrom constructs a [MTRLevelControlClusterMoveToLevelParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveToLevelParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveToLevelParams {
	return MTRLevelControlClusterMoveToLevelParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRLevelControlClusterMoveToLevelParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRLevelControlClusterMoveToLevelParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRLevelControlClusterMoveToLevelParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRLevelControlClusterMoveToLevelParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRLevelControlClusterMoveToLevelParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams/level
func (m_ MTRLevelControlClusterMoveToLevelParams) Level() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("level"))
	return rv
}/* debug [instance_properties/getter]: level */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams/level
func (m_ MTRLevelControlClusterMoveToLevelParams) SetLevel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLevel:"), value)
}/* debug [instance_properties/setter]: level */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams/optionsMask
func (m_ MTRLevelControlClusterMoveToLevelParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams/optionsMask
func (m_ MTRLevelControlClusterMoveToLevelParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams/optionsOverride
func (m_ MTRLevelControlClusterMoveToLevelParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams/optionsOverride
func (m_ MTRLevelControlClusterMoveToLevelParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterMoveToLevelParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterMoveToLevelParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterMoveToLevelParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterMoveToLevelParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams/transitionTime
func (m_ MTRLevelControlClusterMoveToLevelParams) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams/transitionTime
func (m_ MTRLevelControlClusterMoveToLevelParams) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRLevelControlClusterMoveToLevelParams */



