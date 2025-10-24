// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRAccessControlClusterReviewFabricRestrictionsResponseParams */


/* debug [class_header]: Header for MTRAccessControlClusterReviewFabricRestrictionsResponseParams */
// The class instance for the [MTRAccessControlClusterReviewFabricRestrictionsResponseParams] class.
var (
	MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass     _MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass
	MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClassOnce sync.Once
)

func getMTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass() _MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass {
	MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClassOnce.Do(func() {
		MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass = _MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass{objc.GetClass("MTRAccessControlClusterReviewFabricRestrictionsResponseParams")}
	})
	return MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass
}

type _MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAccessControlClusterReviewFabricRestrictionsResponseParams */
// An interface definition for the [MTRAccessControlClusterReviewFabricRestrictionsResponseParams] class.
type IMTRAccessControlClusterReviewFabricRestrictionsResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRAccessControlClusterReviewFabricRestrictionsResponseParams */
	// properties:
	Token() objc.IObject /* cross-framework: NSNumber */
	SetToken(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAccessControlClusterReviewFabricRestrictionsResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAccessControlClusterReviewFabricRestrictionsResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass) Alloc() MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass) New() MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterReviewFabricRestrictionsResponseParams) Init() MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterReviewFabricRestrictionsResponseParams) Autorelease() MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterReviewFabricRestrictionsResponseParams creates a new MTRAccessControlClusterReviewFabricRestrictionsResponseParams instance.
func NewMTRAccessControlClusterReviewFabricRestrictionsResponseParams() MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	return getMTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAccessControlClusterReviewFabricRestrictionsResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsResponseParams
type MTRAccessControlClusterReviewFabricRestrictionsResponseParams struct {
	objectivec.Object
}

// MTRAccessControlClusterReviewFabricRestrictionsResponseParamsFrom constructs a [MTRAccessControlClusterReviewFabricRestrictionsResponseParams] from an unsafe.Pointer.
func MTRAccessControlClusterReviewFabricRestrictionsResponseParamsFrom(ptr unsafe.Pointer) MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	return MTRAccessControlClusterReviewFabricRestrictionsResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAccessControlClusterReviewFabricRestrictionsResponseParams */

// Initialize an MTRAccessControlClusterReviewFabricRestrictionsResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsResponseParams/init(responseValue:)
func NewMTRAccessControlClusterReviewFabricRestrictionsResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	instance := getMTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass().Alloc()
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRAccessControlClusterReviewFabricRestrictionsResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAccessControlClusterReviewFabricRestrictionsResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAccessControlClusterReviewFabricRestrictionsResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAccessControlClusterReviewFabricRestrictionsResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAccessControlClusterReviewFabricRestrictionsResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusterreviewfabricrestrictionsresponseparams/token
func (m_ MTRAccessControlClusterReviewFabricRestrictionsResponseParams) Token() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("token"))
	return rv
}/* debug [instance_properties/getter]: token */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusterreviewfabricrestrictionsresponseparams/token
func (m_ MTRAccessControlClusterReviewFabricRestrictionsResponseParams) SetToken(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setToken:"), value)
}/* debug [instance_properties/setter]: token */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAccessControlClusterReviewFabricRestrictionsResponseParams */


