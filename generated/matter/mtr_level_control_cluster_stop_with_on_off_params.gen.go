// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRLevelControlClusterStopWithOnOffParams */


/* debug [class_header]: Header for MTRLevelControlClusterStopWithOnOffParams */
// The class instance for the [MTRLevelControlClusterStopWithOnOffParams] class.
var (
	MTRLevelControlClusterStopWithOnOffParamsClass     _MTRLevelControlClusterStopWithOnOffParamsClass
	MTRLevelControlClusterStopWithOnOffParamsClassOnce sync.Once
)

func getMTRLevelControlClusterStopWithOnOffParamsClass() _MTRLevelControlClusterStopWithOnOffParamsClass {
	MTRLevelControlClusterStopWithOnOffParamsClassOnce.Do(func() {
		MTRLevelControlClusterStopWithOnOffParamsClass = _MTRLevelControlClusterStopWithOnOffParamsClass{objc.GetClass("MTRLevelControlClusterStopWithOnOffParams")}
	})
	return MTRLevelControlClusterStopWithOnOffParamsClass
}

type _MTRLevelControlClusterStopWithOnOffParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRLevelControlClusterStopWithOnOffParams */
// An interface definition for the [MTRLevelControlClusterStopWithOnOffParams] class.
type IMTRLevelControlClusterStopWithOnOffParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRLevelControlClusterStopWithOnOffParams */
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

	
/* debug [class_interface_methods]: Methods for MTRLevelControlClusterStopWithOnOffParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRLevelControlClusterStopWithOnOffParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterStopWithOnOffParamsClass) Alloc() MTRLevelControlClusterStopWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStopWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRLevelControlClusterStopWithOnOffParamsClass) New() MTRLevelControlClusterStopWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStopWithOnOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterStopWithOnOffParams) Init() MTRLevelControlClusterStopWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStopWithOnOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterStopWithOnOffParams) Autorelease() MTRLevelControlClusterStopWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStopWithOnOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterStopWithOnOffParams creates a new MTRLevelControlClusterStopWithOnOffParams instance.
func NewMTRLevelControlClusterStopWithOnOffParams() MTRLevelControlClusterStopWithOnOffParams {
	return getMTRLevelControlClusterStopWithOnOffParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRLevelControlClusterStopWithOnOffParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStopWithOnOffParams
type MTRLevelControlClusterStopWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterStopWithOnOffParamsFrom constructs a [MTRLevelControlClusterStopWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterStopWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterStopWithOnOffParams {
	return MTRLevelControlClusterStopWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRLevelControlClusterStopWithOnOffParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRLevelControlClusterStopWithOnOffParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRLevelControlClusterStopWithOnOffParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRLevelControlClusterStopWithOnOffParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRLevelControlClusterStopWithOnOffParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStopWithOnOffParams/optionsMask
func (m_ MTRLevelControlClusterStopWithOnOffParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStopWithOnOffParams/optionsMask
func (m_ MTRLevelControlClusterStopWithOnOffParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStopWithOnOffParams/optionsOverride
func (m_ MTRLevelControlClusterStopWithOnOffParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStopWithOnOffParams/optionsOverride
func (m_ MTRLevelControlClusterStopWithOnOffParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStopWithOnOffParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterStopWithOnOffParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStopWithOnOffParams/serverSideProcessingTimeout
func (m_ MTRLevelControlClusterStopWithOnOffParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStopWithOnOffParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterStopWithOnOffParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStopWithOnOffParams/timedInvokeTimeoutMs
func (m_ MTRLevelControlClusterStopWithOnOffParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRLevelControlClusterStopWithOnOffParams */



