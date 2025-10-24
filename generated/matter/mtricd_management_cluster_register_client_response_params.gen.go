// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRICDManagementClusterRegisterClientResponseParams */


/* debug [class_header]: Header for MTRICDManagementClusterRegisterClientResponseParams */
// The class instance for the [MTRICDManagementClusterRegisterClientResponseParams] class.
var (
	MTRICDManagementClusterRegisterClientResponseParamsClass     _MTRICDManagementClusterRegisterClientResponseParamsClass
	MTRICDManagementClusterRegisterClientResponseParamsClassOnce sync.Once
)

func getMTRICDManagementClusterRegisterClientResponseParamsClass() _MTRICDManagementClusterRegisterClientResponseParamsClass {
	MTRICDManagementClusterRegisterClientResponseParamsClassOnce.Do(func() {
		MTRICDManagementClusterRegisterClientResponseParamsClass = _MTRICDManagementClusterRegisterClientResponseParamsClass{objc.GetClass("MTRICDManagementClusterRegisterClientResponseParams")}
	})
	return MTRICDManagementClusterRegisterClientResponseParamsClass
}

type _MTRICDManagementClusterRegisterClientResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRICDManagementClusterRegisterClientResponseParams */
// An interface definition for the [MTRICDManagementClusterRegisterClientResponseParams] class.
type IMTRICDManagementClusterRegisterClientResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRICDManagementClusterRegisterClientResponseParams */
	// properties:
	IcdCounter() objc.IObject /* cross-framework: NSNumber */
	SetIcdCounter(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRICDManagementClusterRegisterClientResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRICDManagementClusterRegisterClientResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRICDManagementClusterRegisterClientResponseParamsClass) Alloc() MTRICDManagementClusterRegisterClientResponseParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRICDManagementClusterRegisterClientResponseParamsClass) New() MTRICDManagementClusterRegisterClientResponseParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRICDManagementClusterRegisterClientResponseParams) Init() MTRICDManagementClusterRegisterClientResponseParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRICDManagementClusterRegisterClientResponseParams) Autorelease() MTRICDManagementClusterRegisterClientResponseParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRICDManagementClusterRegisterClientResponseParams creates a new MTRICDManagementClusterRegisterClientResponseParams instance.
func NewMTRICDManagementClusterRegisterClientResponseParams() MTRICDManagementClusterRegisterClientResponseParams {
	return getMTRICDManagementClusterRegisterClientResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRICDManagementClusterRegisterClientResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientResponseParams
type MTRICDManagementClusterRegisterClientResponseParams struct {
	objectivec.Object
}

// MTRICDManagementClusterRegisterClientResponseParamsFrom constructs a [MTRICDManagementClusterRegisterClientResponseParams] from an unsafe.Pointer.
func MTRICDManagementClusterRegisterClientResponseParamsFrom(ptr unsafe.Pointer) MTRICDManagementClusterRegisterClientResponseParams {
	return MTRICDManagementClusterRegisterClientResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRICDManagementClusterRegisterClientResponseParams */

// Initialize an MTRICDManagementClusterRegisterClientResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientResponseParams/init(responseValue:)
func NewMTRICDManagementClusterRegisterClientResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRICDManagementClusterRegisterClientResponseParams {
	instance := getMTRICDManagementClusterRegisterClientResponseParamsClass().Alloc()
	rv := objc.Send[MTRICDManagementClusterRegisterClientResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRICDManagementClusterRegisterClientResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRICDManagementClusterRegisterClientResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRICDManagementClusterRegisterClientResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRICDManagementClusterRegisterClientResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRICDManagementClusterRegisterClientResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientresponseparams/icdcounter
func (m_ MTRICDManagementClusterRegisterClientResponseParams) IcdCounter() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("icdCounter"))
	return rv
}/* debug [instance_properties/getter]: icdCounter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterregisterclientresponseparams/icdcounter
func (m_ MTRICDManagementClusterRegisterClientResponseParams) SetIcdCounter(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIcdCounter:"), value)
}/* debug [instance_properties/setter]: icdCounter */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRICDManagementClusterRegisterClientResponseParams */


