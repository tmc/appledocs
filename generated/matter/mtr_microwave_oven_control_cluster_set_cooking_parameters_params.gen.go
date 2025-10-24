// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMicrowaveOvenControlClusterSetCookingParametersParams */


/* debug [class_header]: Header for MTRMicrowaveOvenControlClusterSetCookingParametersParams */
// The class instance for the [MTRMicrowaveOvenControlClusterSetCookingParametersParams] class.
var (
	MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass     _MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass
	MTRMicrowaveOvenControlClusterSetCookingParametersParamsClassOnce sync.Once
)

func getMTRMicrowaveOvenControlClusterSetCookingParametersParamsClass() _MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass {
	MTRMicrowaveOvenControlClusterSetCookingParametersParamsClassOnce.Do(func() {
		MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass = _MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass{objc.GetClass("MTRMicrowaveOvenControlClusterSetCookingParametersParams")}
	})
	return MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass
}

type _MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMicrowaveOvenControlClusterSetCookingParametersParams */
// An interface definition for the [MTRMicrowaveOvenControlClusterSetCookingParametersParams] class.
type IMTRMicrowaveOvenControlClusterSetCookingParametersParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMicrowaveOvenControlClusterSetCookingParametersParams */
	// properties:
	PowerSetting() objc.IObject /* cross-framework: NSNumber */
	SetPowerSetting(value objc.IObject /* cross-framework: NSNumber */)
	CookMode() objc.IObject /* cross-framework: NSNumber */
	SetCookMode(value objc.IObject /* cross-framework: NSNumber */)
	CookTime() objc.IObject /* cross-framework: NSNumber */
	SetCookTime(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StartAfterSetting() objc.IObject /* cross-framework: NSNumber */
	SetStartAfterSetting(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMicrowaveOvenControlClusterSetCookingParametersParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMicrowaveOvenControlClusterSetCookingParametersParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass) Alloc() MTRMicrowaveOvenControlClusterSetCookingParametersParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterSetCookingParametersParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass) New() MTRMicrowaveOvenControlClusterSetCookingParametersParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterSetCookingParametersParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) Init() MTRMicrowaveOvenControlClusterSetCookingParametersParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterSetCookingParametersParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) Autorelease() MTRMicrowaveOvenControlClusterSetCookingParametersParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterSetCookingParametersParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMicrowaveOvenControlClusterSetCookingParametersParams creates a new MTRMicrowaveOvenControlClusterSetCookingParametersParams instance.
func NewMTRMicrowaveOvenControlClusterSetCookingParametersParams() MTRMicrowaveOvenControlClusterSetCookingParametersParams {
	return getMTRMicrowaveOvenControlClusterSetCookingParametersParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMicrowaveOvenControlClusterSetCookingParametersParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams
type MTRMicrowaveOvenControlClusterSetCookingParametersParams struct {
	objectivec.Object
}

// MTRMicrowaveOvenControlClusterSetCookingParametersParamsFrom constructs a [MTRMicrowaveOvenControlClusterSetCookingParametersParams] from an unsafe.Pointer.
func MTRMicrowaveOvenControlClusterSetCookingParametersParamsFrom(ptr unsafe.Pointer) MTRMicrowaveOvenControlClusterSetCookingParametersParams {
	return MTRMicrowaveOvenControlClusterSetCookingParametersParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMicrowaveOvenControlClusterSetCookingParametersParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMicrowaveOvenControlClusterSetCookingParametersParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMicrowaveOvenControlClusterSetCookingParametersParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMicrowaveOvenControlClusterSetCookingParametersParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMicrowaveOvenControlClusterSetCookingParametersParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/powerSetting
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) PowerSetting() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("powerSetting"))
	return rv
}/* debug [instance_properties/getter]: powerSetting */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/powerSetting
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) SetPowerSetting(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPowerSetting:"), value)
}/* debug [instance_properties/setter]: powerSetting */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclustersetcookingparametersparams/cookmode
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) CookMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cookMode"))
	return rv
}/* debug [instance_properties/getter]: cookMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclustersetcookingparametersparams/cookmode
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) SetCookMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCookMode:"), value)
}/* debug [instance_properties/setter]: cookMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclustersetcookingparametersparams/cooktime
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) CookTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cookTime"))
	return rv
}/* debug [instance_properties/getter]: cookTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclustersetcookingparametersparams/cooktime
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) SetCookTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCookTime:"), value)
}/* debug [instance_properties/setter]: cookTime */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclustersetcookingparametersparams/serversideprocessingtimeout
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclustersetcookingparametersparams/serversideprocessingtimeout
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclustersetcookingparametersparams/startaftersetting
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) StartAfterSetting() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startAfterSetting"))
	return rv
}/* debug [instance_properties/getter]: startAfterSetting */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclustersetcookingparametersparams/startaftersetting
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) SetStartAfterSetting(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartAfterSetting:"), value)
}/* debug [instance_properties/setter]: startAfterSetting */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclustersetcookingparametersparams/timedinvoketimeoutms
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovencontrolclustersetcookingparametersparams/timedinvoketimeoutms
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMicrowaveOvenControlClusterSetCookingParametersParams */



