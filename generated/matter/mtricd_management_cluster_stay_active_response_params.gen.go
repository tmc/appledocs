// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRICDManagementClusterStayActiveResponseParams */


/* debug [class_header]: Header for MTRICDManagementClusterStayActiveResponseParams */
// The class instance for the [MTRICDManagementClusterStayActiveResponseParams] class.
var (
	MTRICDManagementClusterStayActiveResponseParamsClass     _MTRICDManagementClusterStayActiveResponseParamsClass
	MTRICDManagementClusterStayActiveResponseParamsClassOnce sync.Once
)

func getMTRICDManagementClusterStayActiveResponseParamsClass() _MTRICDManagementClusterStayActiveResponseParamsClass {
	MTRICDManagementClusterStayActiveResponseParamsClassOnce.Do(func() {
		MTRICDManagementClusterStayActiveResponseParamsClass = _MTRICDManagementClusterStayActiveResponseParamsClass{objc.GetClass("MTRICDManagementClusterStayActiveResponseParams")}
	})
	return MTRICDManagementClusterStayActiveResponseParamsClass
}

type _MTRICDManagementClusterStayActiveResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRICDManagementClusterStayActiveResponseParams */
// An interface definition for the [MTRICDManagementClusterStayActiveResponseParams] class.
type IMTRICDManagementClusterStayActiveResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRICDManagementClusterStayActiveResponseParams */
	// properties:
	PromisedActiveDuration() objc.IObject /* cross-framework: NSNumber */
	SetPromisedActiveDuration(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRICDManagementClusterStayActiveResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRICDManagementClusterStayActiveResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRICDManagementClusterStayActiveResponseParamsClass) Alloc() MTRICDManagementClusterStayActiveResponseParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRICDManagementClusterStayActiveResponseParamsClass) New() MTRICDManagementClusterStayActiveResponseParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRICDManagementClusterStayActiveResponseParams) Init() MTRICDManagementClusterStayActiveResponseParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRICDManagementClusterStayActiveResponseParams) Autorelease() MTRICDManagementClusterStayActiveResponseParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRICDManagementClusterStayActiveResponseParams creates a new MTRICDManagementClusterStayActiveResponseParams instance.
func NewMTRICDManagementClusterStayActiveResponseParams() MTRICDManagementClusterStayActiveResponseParams {
	return getMTRICDManagementClusterStayActiveResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRICDManagementClusterStayActiveResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveResponseParams
type MTRICDManagementClusterStayActiveResponseParams struct {
	objectivec.Object
}

// MTRICDManagementClusterStayActiveResponseParamsFrom constructs a [MTRICDManagementClusterStayActiveResponseParams] from an unsafe.Pointer.
func MTRICDManagementClusterStayActiveResponseParamsFrom(ptr unsafe.Pointer) MTRICDManagementClusterStayActiveResponseParams {
	return MTRICDManagementClusterStayActiveResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRICDManagementClusterStayActiveResponseParams */

// Initialize an MTRICDManagementClusterStayActiveResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveResponseParams/init(responseValue:)
func NewMTRICDManagementClusterStayActiveResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRICDManagementClusterStayActiveResponseParams {
	instance := getMTRICDManagementClusterStayActiveResponseParamsClass().Alloc()
	rv := objc.Send[MTRICDManagementClusterStayActiveResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRICDManagementClusterStayActiveResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRICDManagementClusterStayActiveResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRICDManagementClusterStayActiveResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRICDManagementClusterStayActiveResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRICDManagementClusterStayActiveResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterstayactiveresponseparams/promisedactiveduration
func (m_ MTRICDManagementClusterStayActiveResponseParams) PromisedActiveDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("promisedActiveDuration"))
	return rv
}/* debug [instance_properties/getter]: promisedActiveDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtricdmanagementclusterstayactiveresponseparams/promisedactiveduration
func (m_ MTRICDManagementClusterStayActiveResponseParams) SetPromisedActiveDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPromisedActiveDuration:"), value)
}/* debug [instance_properties/setter]: promisedActiveDuration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRICDManagementClusterStayActiveResponseParams */


