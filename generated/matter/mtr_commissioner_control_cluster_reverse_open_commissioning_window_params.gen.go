// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRCommissionerControlClusterReverseOpenCommissioningWindowParams */


/* debug [class_header]: Header for MTRCommissionerControlClusterReverseOpenCommissioningWindowParams */
// The class instance for the [MTRCommissionerControlClusterReverseOpenCommissioningWindowParams] class.
var (
	MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass     _MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass
	MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClassOnce sync.Once
)

func getMTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass() _MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass {
	MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClassOnce.Do(func() {
		MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass = _MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass{objc.GetClass("MTRCommissionerControlClusterReverseOpenCommissioningWindowParams")}
	})
	return MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass
}

type _MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRCommissionerControlClusterReverseOpenCommissioningWindowParams */
// An interface definition for the [MTRCommissionerControlClusterReverseOpenCommissioningWindowParams] class.
type IMTRCommissionerControlClusterReverseOpenCommissioningWindowParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRCommissionerControlClusterReverseOpenCommissioningWindowParams */
	// properties:
	CommissioningTimeout() objc.IObject /* cross-framework: NSNumber */
	SetCommissioningTimeout(value objc.IObject /* cross-framework: NSNumber */)
	Discriminator() objc.IObject /* cross-framework: NSNumber */
	SetDiscriminator(value objc.IObject /* cross-framework: NSNumber */)
	Iterations() objc.IObject /* cross-framework: NSNumber */
	SetIterations(value objc.IObject /* cross-framework: NSNumber */)
	PakePasscodeVerifier() foundation.Data
	SetPakePasscodeVerifier(value foundation.Data)
	Salt() foundation.Data
	SetSalt(value foundation.Data)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRCommissionerControlClusterReverseOpenCommissioningWindowParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRCommissionerControlClusterReverseOpenCommissioningWindowParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass) Alloc() MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	rv := objc.Send[MTRCommissionerControlClusterReverseOpenCommissioningWindowParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass) New() MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	rv := objc.Send[MTRCommissionerControlClusterReverseOpenCommissioningWindowParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) Init() MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	rv := objc.Send[MTRCommissionerControlClusterReverseOpenCommissioningWindowParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) Autorelease() MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	rv := objc.Send[MTRCommissionerControlClusterReverseOpenCommissioningWindowParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommissionerControlClusterReverseOpenCommissioningWindowParams creates a new MTRCommissionerControlClusterReverseOpenCommissioningWindowParams instance.
func NewMTRCommissionerControlClusterReverseOpenCommissioningWindowParams() MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	return getMTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRCommissionerControlClusterReverseOpenCommissioningWindowParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams
type MTRCommissionerControlClusterReverseOpenCommissioningWindowParams struct {
	objectivec.Object
}

// MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsFrom constructs a [MTRCommissionerControlClusterReverseOpenCommissioningWindowParams] from an unsafe.Pointer.
func MTRCommissionerControlClusterReverseOpenCommissioningWindowParamsFrom(ptr unsafe.Pointer) MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	return MTRCommissionerControlClusterReverseOpenCommissioningWindowParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRCommissionerControlClusterReverseOpenCommissioningWindowParams */

// Initialize an MTRCommissionerControlClusterReverseOpenCommissioningWindowParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissionerControlClusterReverseOpenCommissioningWindowParams/init(responseValue:)
func NewMTRCommissionerControlClusterReverseOpenCommissioningWindowParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRCommissionerControlClusterReverseOpenCommissioningWindowParams {
	instance := getMTRCommissionerControlClusterReverseOpenCommissioningWindowParamsClass().Alloc()
	rv := objc.Send[MTRCommissionerControlClusterReverseOpenCommissioningWindowParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRCommissionerControlClusterReverseOpenCommissioningWindowParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRCommissionerControlClusterReverseOpenCommissioningWindowParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRCommissionerControlClusterReverseOpenCommissioningWindowParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRCommissionerControlClusterReverseOpenCommissioningWindowParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRCommissionerControlClusterReverseOpenCommissioningWindowParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterreverseopencommissioningwindowparams/commissioningtimeout
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) CommissioningTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("commissioningTimeout"))
	return rv
}/* debug [instance_properties/getter]: commissioningTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterreverseopencommissioningwindowparams/commissioningtimeout
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) SetCommissioningTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommissioningTimeout:"), value)
}/* debug [instance_properties/setter]: commissioningTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterreverseopencommissioningwindowparams/discriminator
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) Discriminator() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("discriminator"))
	return rv
}/* debug [instance_properties/getter]: discriminator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterreverseopencommissioningwindowparams/discriminator
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) SetDiscriminator(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDiscriminator:"), value)
}/* debug [instance_properties/setter]: discriminator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterreverseopencommissioningwindowparams/iterations
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) Iterations() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("iterations"))
	return rv
}/* debug [instance_properties/getter]: iterations */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterreverseopencommissioningwindowparams/iterations
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) SetIterations(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIterations:"), value)
}/* debug [instance_properties/setter]: iterations */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterreverseopencommissioningwindowparams/pakepasscodeverifier
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) PakePasscodeVerifier() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("pakePasscodeVerifier"))
	return rv
}/* debug [instance_properties/getter]: pakePasscodeVerifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterreverseopencommissioningwindowparams/pakepasscodeverifier
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) SetPakePasscodeVerifier(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPakePasscodeVerifier:"), value)
}/* debug [instance_properties/setter]: pakePasscodeVerifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterreverseopencommissioningwindowparams/salt
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) Salt() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("salt"))
	return rv
}/* debug [instance_properties/getter]: salt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissionercontrolclusterreverseopencommissioningwindowparams/salt
func (m_ MTRCommissionerControlClusterReverseOpenCommissioningWindowParams) SetSalt(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSalt:"), value)
}/* debug [instance_properties/setter]: salt */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRCommissionerControlClusterReverseOpenCommissioningWindowParams */


