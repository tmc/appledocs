// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRColorControlClusterMoveColorParams */


/* debug [class_header]: Header for MTRColorControlClusterMoveColorParams */
// The class instance for the [MTRColorControlClusterMoveColorParams] class.
var (
	MTRColorControlClusterMoveColorParamsClass     _MTRColorControlClusterMoveColorParamsClass
	MTRColorControlClusterMoveColorParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveColorParamsClass() _MTRColorControlClusterMoveColorParamsClass {
	MTRColorControlClusterMoveColorParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveColorParamsClass = _MTRColorControlClusterMoveColorParamsClass{objc.GetClass("MTRColorControlClusterMoveColorParams")}
	})
	return MTRColorControlClusterMoveColorParamsClass
}

type _MTRColorControlClusterMoveColorParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRColorControlClusterMoveColorParams */
// An interface definition for the [MTRColorControlClusterMoveColorParams] class.
type IMTRColorControlClusterMoveColorParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRColorControlClusterMoveColorParams */
	// properties:
	OptionsMask() objc.IObject /* cross-framework: NSNumber */
	SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */)
	OptionsOverride() objc.IObject /* cross-framework: NSNumber */
	SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */)
	RateX() objc.IObject /* cross-framework: NSNumber */
	SetRateX(value objc.IObject /* cross-framework: NSNumber */)
	RateY() objc.IObject /* cross-framework: NSNumber */
	SetRateY(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRColorControlClusterMoveColorParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRColorControlClusterMoveColorParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveColorParamsClass) Alloc() MTRColorControlClusterMoveColorParams {
	rv := objc.Send[MTRColorControlClusterMoveColorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRColorControlClusterMoveColorParamsClass) New() MTRColorControlClusterMoveColorParams {
	rv := objc.Send[MTRColorControlClusterMoveColorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveColorParams) Init() MTRColorControlClusterMoveColorParams {
	rv := objc.Send[MTRColorControlClusterMoveColorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveColorParams) Autorelease() MTRColorControlClusterMoveColorParams {
	rv := objc.Send[MTRColorControlClusterMoveColorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveColorParams creates a new MTRColorControlClusterMoveColorParams instance.
func NewMTRColorControlClusterMoveColorParams() MTRColorControlClusterMoveColorParams {
	return getMTRColorControlClusterMoveColorParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRColorControlClusterMoveColorParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams
type MTRColorControlClusterMoveColorParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveColorParamsFrom constructs a [MTRColorControlClusterMoveColorParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveColorParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveColorParams {
	return MTRColorControlClusterMoveColorParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRColorControlClusterMoveColorParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRColorControlClusterMoveColorParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRColorControlClusterMoveColorParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRColorControlClusterMoveColorParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRColorControlClusterMoveColorParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams/optionsMask
func (m_ MTRColorControlClusterMoveColorParams) OptionsMask() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsMask"))
	return rv
}/* debug [instance_properties/getter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams/optionsMask
func (m_ MTRColorControlClusterMoveColorParams) SetOptionsMask(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsMask:"), value)
}/* debug [instance_properties/setter]: optionsMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams/optionsOverride
func (m_ MTRColorControlClusterMoveColorParams) OptionsOverride() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionsOverride"))
	return rv
}/* debug [instance_properties/getter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams/optionsOverride
func (m_ MTRColorControlClusterMoveColorParams) SetOptionsOverride(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionsOverride:"), value)
}/* debug [instance_properties/setter]: optionsOverride */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams/rateX
func (m_ MTRColorControlClusterMoveColorParams) RateX() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rateX"))
	return rv
}/* debug [instance_properties/getter]: rateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams/rateX
func (m_ MTRColorControlClusterMoveColorParams) SetRateX(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRateX:"), value)
}/* debug [instance_properties/setter]: rateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams/rateY
func (m_ MTRColorControlClusterMoveColorParams) RateY() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rateY"))
	return rv
}/* debug [instance_properties/getter]: rateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams/rateY
func (m_ MTRColorControlClusterMoveColorParams) SetRateY(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRateY:"), value)
}/* debug [instance_properties/setter]: rateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveColorParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams/serverSideProcessingTimeout
func (m_ MTRColorControlClusterMoveColorParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveColorParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams/timedInvokeTimeoutMs
func (m_ MTRColorControlClusterMoveColorParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRColorControlClusterMoveColorParams */



