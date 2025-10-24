// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEFlowMetaData */


/* debug [class_header]: Header for NEFlowMetaData */
// The class instance for the [NEFlowMetaData] class.
var (
	NEFlowMetaDataClass     _NEFlowMetaDataClass
	NEFlowMetaDataClassOnce sync.Once
)

func getNEFlowMetaDataClass() _NEFlowMetaDataClass {
	NEFlowMetaDataClassOnce.Do(func() {
		NEFlowMetaDataClass = _NEFlowMetaDataClass{objc.GetClass("NEFlowMetaData")}
	})
	return NEFlowMetaDataClass
}

type _NEFlowMetaDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFlowMetaData */
// An interface definition for the [NEFlowMetaData] class.
type INEFlowMetaData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEFlowMetaData */
	// properties:
	FilterFlowIdentifier() foundation.UUID
	SourceAppAuditToken() objc.IObject /* cross-framework: NSData */
	SourceAppSigningIdentifier() objc.IObject /* cross-framework: NSString */
	SourceAppUniqueIdentifier() objc.IObject /* cross-framework: NSData */
	RoutingMethod() NETunnelProviderRoutingMethod
	SetRoutingMethod(value NETunnelProviderRoutingMethod)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFlowMetaData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFlowMetaData */
// Alloc allocates a new instance without initialization.
func (nc _NEFlowMetaDataClass) Alloc() NEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFlowMetaDataClass) New() NEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFlowMetaData) Init() NEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFlowMetaData) Autorelease() NEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFlowMetaData creates a new NEFlowMetaData instance.
func NewNEFlowMetaData() NEFlowMetaData {
	return getNEFlowMetaDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFlowMetaData */
// Additional information about data flowing through a per-app VPN provider.
//
// This metadata is only present for data flowing through per-app VPN providers, that is, app proxy providers and packet tunnel providers in per-app VPN mode, as indicated by the property.


// Additional information about data flowing through a per-app VPN provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData
type NEFlowMetaData struct {
	objectivec.Object
}

// NEFlowMetaDataFrom constructs a [NEFlowMetaData] from an unsafe.Pointer.
//
// Additional information about data flowing through a per-app VPN provider.
func NEFlowMetaDataFrom(ptr unsafe.Pointer) NEFlowMetaData {
	return NEFlowMetaData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFlowMetaData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFlowMetaData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFlowMetaData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFlowMetaData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFlowMetaData */

// The identifier of the content filter flow corresponding to this flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData/filterFlowIdentifier
func (n_ NEFlowMetaData) FilterFlowIdentifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](n_.ID, objc.Sel("filterFlowIdentifier"))
	return rv
}/* debug [instance_properties/getter]: filterFlowIdentifier */


// The audit token of the source application of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData/sourceAppAuditToken
func (n_ NEFlowMetaData) SourceAppAuditToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("sourceAppAuditToken"))
	return rv
}/* debug [instance_properties/getter]: sourceAppAuditToken */


// A string that contains the signing identifier of the source application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData/sourceAppSigningIdentifier
func (n_ NEFlowMetaData) SourceAppSigningIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("sourceAppSigningIdentifier"))
	return rv
}/* debug [instance_properties/getter]: sourceAppSigningIdentifier */


// A data instance that contains a unique hash value for the source application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData/sourceAppUniqueIdentifier
func (n_ NEFlowMetaData) SourceAppUniqueIdentifier() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("sourceAppUniqueIdentifier"))
	return rv
}/* debug [instance_properties/getter]: sourceAppUniqueIdentifier */


// The method by which network traffic is routed to the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovider/routingmethod
func (n_ NEFlowMetaData) RoutingMethod() NETunnelProviderRoutingMethod {
	rv := objc.Send[NETunnelProviderRoutingMethod](n_.ID, objc.Sel("routingMethod"))
	return rv
}/* debug [instance_properties/getter]: routingMethod */


// The method by which network traffic is routed to the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovider/routingmethod
func (n_ NEFlowMetaData) SetRoutingMethod(value NETunnelProviderRoutingMethod) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoutingMethod:"), value)
}/* debug [instance_properties/setter]: routingMethod */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFlowMetaData */



