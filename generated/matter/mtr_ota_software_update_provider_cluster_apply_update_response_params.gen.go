// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams */


/* debug [class_header]: Header for MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams */
// The class instance for the [MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams] class.
var (
	MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass     _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass
	MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClassOnce sync.Once
)

func getMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass() _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass {
	MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClassOnce.Do(func() {
		MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass = _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass{objc.GetClass("MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams")}
	})
	return MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass
}

type _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams */
// An interface definition for the [MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams] class.
type IMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams interface {
	IMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams
	
/* debug [class_interface_properties]: Properties for MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams */
	// properties:
	Action() objc.IObject /* cross-framework: NSNumber */
	SetAction(value objc.IObject /* cross-framework: NSNumber */)
	DelayedActionTime() objc.IObject /* cross-framework: NSNumber */
	SetDelayedActionTime(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass) New() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) Init() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) Autorelease() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams creates a new MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams instance.
func NewMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams() MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	return getMTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams-92als
type MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams struct {
	MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams
}

// MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom constructs a [MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams] from an unsafe.Pointer.
func MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams {
	return MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams{
		MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams: MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams-92als/action
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) Action() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams-92als/action
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) SetAction(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams-92als/delayedActionTime
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) DelayedActionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("delayedActionTime"))
	return rv
}/* debug [instance_properties/getter]: delayedActionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams-92als/delayedActionTime
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) SetDelayedActionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelayedActionTime:"), value)
}/* debug [instance_properties/setter]: delayedActionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams-92als/timedInvokeTimeoutMs
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams-92als/timedInvokeTimeoutMs
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROtaSoftwareUpdateProviderClusterApplyUpdateResponseParams */



