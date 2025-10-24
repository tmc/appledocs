// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams */


/* debug [class_header]: Header for MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams */
// The class instance for the [MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams] class.
var (
	MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass     _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass
	MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClassOnce sync.Once
)

func getMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass() _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass {
	MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClassOnce.Do(func() {
		MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass = _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass{objc.GetClass("MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams")}
	})
	return MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass
}

type _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams */
// An interface definition for the [MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams] class.
type IMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams */
	// properties:
	Action() objc.IObject /* cross-framework: NSNumber */
	SetAction(value objc.IObject /* cross-framework: NSNumber */)
	DelayedActionTime() objc.IObject /* cross-framework: NSNumber */
	SetDelayedActionTime(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass) Alloc() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass) New() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) Init() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) Autorelease() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams creates a new MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams instance.
func NewMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams() MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	return getMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams-36zc9
type MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom constructs a [MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	return MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams-36zc9/init(responseValue:)
func NewMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams {
	instance := getMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsClass().Alloc()
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROTASoftwareUpdateProviderClusterApplyUpdateResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams-36zc9/action
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) Action() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams-36zc9/action
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) SetAction(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams-36zc9/delayedActionTime
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) DelayedActionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("delayedActionTime"))
	return rv
}/* debug [instance_properties/getter]: delayedActionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams-36zc9/delayedActionTime
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) SetDelayedActionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelayedActionTime:"), value)
}/* debug [instance_properties/setter]: delayedActionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams-36zc9/timedInvokeTimeoutMs
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams-36zc9/timedInvokeTimeoutMs
func (m_ MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROTASoftwareUpdateProviderClusterApplyUpdateResponseParams */


