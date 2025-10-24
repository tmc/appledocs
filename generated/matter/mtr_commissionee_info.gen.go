// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRCommissioneeInfo */


/* debug [class_header]: Header for MTRCommissioneeInfo */
// The class instance for the [MTRCommissioneeInfo] class.
var (
	MTRCommissioneeInfoClass     _MTRCommissioneeInfoClass
	MTRCommissioneeInfoClassOnce sync.Once
)

func getMTRCommissioneeInfoClass() _MTRCommissioneeInfoClass {
	MTRCommissioneeInfoClassOnce.Do(func() {
		MTRCommissioneeInfoClass = _MTRCommissioneeInfoClass{objc.GetClass("MTRCommissioneeInfo")}
	})
	return MTRCommissioneeInfoClass
}

type _MTRCommissioneeInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRCommissioneeInfo */
// An interface definition for the [MTRCommissioneeInfo] class.
type IMTRCommissioneeInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRCommissioneeInfo */
	// properties:
	EndpointsById() foundation.IDictionary
	ProductIdentity() IMTRProductIdentity
	SetProductIdentity(value IMTRProductIdentity)
	RootEndpoint() IMTREndpointInfo
	SetRootEndpoint(value IMTREndpointInfo)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRCommissioneeInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRCommissioneeInfo */
// Alloc allocates a new instance without initialization.
func (mc _MTRCommissioneeInfoClass) Alloc() MTRCommissioneeInfo {
	rv := objc.Send[MTRCommissioneeInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRCommissioneeInfoClass) New() MTRCommissioneeInfo {
	rv := objc.Send[MTRCommissioneeInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommissioneeInfo) Init() MTRCommissioneeInfo {
	rv := objc.Send[MTRCommissioneeInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommissioneeInfo) Autorelease() MTRCommissioneeInfo {
	rv := objc.Send[MTRCommissioneeInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommissioneeInfo creates a new MTRCommissioneeInfo instance.
func NewMTRCommissioneeInfo() MTRCommissioneeInfo {
	return getMTRCommissioneeInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRCommissioneeInfo */
// Information read from the commissionee device during commissioning.


// Information read from the commissionee device during commissioning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioneeInfo
type MTRCommissioneeInfo struct {
	objectivec.Object
}

// MTRCommissioneeInfoFrom constructs a [MTRCommissioneeInfo] from an unsafe.Pointer.
//
// Information read from the commissionee device during commissioning.
func MTRCommissioneeInfoFrom(ptr unsafe.Pointer) MTRCommissioneeInfo {
	return MTRCommissioneeInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRCommissioneeInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRCommissioneeInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRCommissioneeInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRCommissioneeInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRCommissioneeInfo */

// Endpoint information for all endpoints of the commissionee. Will be present only if readEndpointInformation is set to YES on MTRCommissioningParameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommissioneeInfo/endpointsById
func (m_ MTRCommissioneeInfo) EndpointsById() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("endpointsById"))
	return rv
}/* debug [instance_properties/getter]: endpointsById */


// The product identity (VID / PID) of the commissionee.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioneeinfo/productidentity
func (m_ MTRCommissioneeInfo) ProductIdentity() IMTRProductIdentity {
	rv := objc.Send[MTRProductIdentity](m_.ID, objc.Sel("productIdentity"))
	return rv
}/* debug [instance_properties/getter]: productIdentity */


// The product identity (VID / PID) of the commissionee.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioneeinfo/productidentity
func (m_ MTRCommissioneeInfo) SetProductIdentity(value IMTRProductIdentity) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductIdentity:"), value)
}/* debug [instance_properties/setter]: productIdentity */


// Endpoint information for the root endpoint of the commissionee.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioneeinfo/rootendpoint
func (m_ MTRCommissioneeInfo) RootEndpoint() IMTREndpointInfo {
	rv := objc.Send[MTREndpointInfo](m_.ID, objc.Sel("rootEndpoint"))
	return rv
}/* debug [instance_properties/getter]: rootEndpoint */


// Endpoint information for the root endpoint of the commissionee.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommissioneeinfo/rootendpoint
func (m_ MTRCommissioneeInfo) SetRootEndpoint(value IMTREndpointInfo) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootEndpoint:"), value)
}/* debug [instance_properties/setter]: rootEndpoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRCommissioneeInfo */



